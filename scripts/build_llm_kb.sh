#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUTPUT_DIR="${1:-$REPO_ROOT/artifacts/llm-kb}"
STAMP="$(date +%Y%m%d_%H%M%S)"
OUT_JSONL="$OUTPUT_DIR/usp-kb-$STAMP.jsonl"
OUT_META="$OUTPUT_DIR/usp-kb-$STAMP.meta.json"
TMP_LIST="$(mktemp)"
TMP_FILTERED="$(mktemp)"
TMP_RECORD="$(mktemp)"

cleanup() {
  rm -f "$TMP_LIST" "$TMP_FILTERED" "$TMP_RECORD"
}
trap cleanup EXIT

require_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "Erro: comando obrigatorio ausente: $1" >&2
    exit 1
  fi
}

require_cmd git
require_cmd jq
require_cmd sed
require_cmd grep
require_cmd wc
require_cmd file
require_cmd stat

mkdir -p "$OUTPUT_DIR"

# Fonte de verdade: apenas arquivos monitorados pelo git.
git -C "$REPO_ROOT" ls-files >"$TMP_LIST"

# Inclusao por escopo aprovado (codigo/docs proprios).
INCLUDE_REGEX='^(README\.md|docker-compose\.yaml|\.gitignore|docs/.*|packages/contracts/src/.*|packages/contracts/test/.*|packages/contracts/script/.*|packages/contracts/abi/.*|packages/contracts/foundry\.toml|packages/contracts/remappings\.txt|packages/dapp/usp-coin-interface/src/.*|packages/dapp/usp-coin-interface/[^/]+\.(js|ts|json|html|md|yml|yaml)|packages/relayer-bundler/cmd/.*|packages/relayer-bundler/internal/.*|packages/relayer-bundler/pkg/.*|packages/relayer-bundler/abi/.*|packages/relayer-bundler/[^/]+\.(go|mod|sum|md|yaml|yml)|packages/relayer-bundler/Dockerfile)$'

# Exclusoes de vendor/build/cache/sensiveis.
EXCLUDE_REGEX='(^|/)(lib/openzeppelin-contracts/|lib/account-abstraction/|lib/forge-std/|node_modules/|dist/|dist-ssr/|cache/|broadcast/|out/)|(^|/)\.env($|\.)|(^|/)annotations\.txt$|(^|/)setup-paymaster\.js$|(^|/)mint-test\.js$|(^|/)package-lock\.json$|(^|/)foundry\.lock$|(^|/)test-build$'

grep -E "$INCLUDE_REGEX" "$TMP_LIST" | grep -Ev "$EXCLUDE_REGEX" >"$TMP_FILTERED" || true

FILE_COUNT="$(wc -l < "$TMP_FILTERED" | tr -d ' ')"
if [[ "$FILE_COUNT" -eq 0 ]]; then
  echo "Erro: nenhum arquivo selecionado apos filtros." >&2
  exit 1
fi

: >"$OUT_JSONL"

total_source_bytes=0
redacted_entries=0

language_from_path() {
  local p="$1"
  case "$p" in
    *.sol) echo "solidity" ;;
    *.go) echo "go" ;;
    *.ts) echo "typescript" ;;
    *.tsx) echo "typescriptreact" ;;
    *.js) echo "javascript" ;;
    *.jsx) echo "javascriptreact" ;;
    *.json) echo "json" ;;
    *.md) echo "markdown" ;;
    *.yaml|*.yml) echo "yaml" ;;
    *.toml) echo "toml" ;;
    *.html) echo "html" ;;
    *.css) echo "css" ;;
    *) echo "text" ;;
  esac
}

sanitize_content() {
  sed -E \
    -e 's/(PRIVATE_KEY\s*[=:]\s*)[^"'\''[:space:]]+/\1[REDACTED]/g' \
    -e 's/(MNEMONIC\s*[=:]\s*)[^\r\n]+/\1[REDACTED]/g' \
    -e 's/(API_KEY\s*[=:]\s*)[^"'\''[:space:]]+/\1[REDACTED]/g' \
    -e 's/(TOKEN\s*[=:]\s*)[^"'\''[:space:]]+/\1[REDACTED]/g' \
    -e 's/(SECRET\s*[=:]\s*)[^"'\''[:space:]]+/\1[REDACTED]/g' \
    -e 's/0x[a-fA-F0-9]{64}/0x[REDACTED_64_HEX]/g'
}

while IFS= read -r rel_path; do
  abs_path="$REPO_ROOT/$rel_path"
  [[ -f "$abs_path" ]] || continue

  # Pular binarios para evitar ruido no corpus.
  if file --brief --mime "$abs_path" | grep -qi 'charset=binary'; then
    continue
  fi

  size_bytes="$(stat -c%s "$abs_path" 2>/dev/null || stat -f%z "$abs_path")"
  total_source_bytes=$((total_source_bytes + size_bytes))

  language="$(language_from_path "$rel_path")"

  raw_content="$(cat "$abs_path")"
  sanitized_content="$(printf '%s' "$raw_content" | sanitize_content)"

  redacted="false"
  if [[ "$raw_content" != "$sanitized_content" ]]; then
    redacted="true"
    redacted_entries=$((redacted_entries + 1))
  fi

  printf '%s' "$sanitized_content" >"$TMP_RECORD"
  jq -cn \
    --arg path "$rel_path" \
    --arg language "$language" \
    --argjson size_bytes "$size_bytes" \
    --argjson line_count "$(wc -l < "$abs_path" | tr -d ' ')" \
    --argjson redacted "$redacted" \
    --rawfile content "$TMP_RECORD" \
    '{path:$path,language:$language,size_bytes:$size_bytes,line_count:$line_count,redacted:$redacted,content:$content}' \
    >>"$OUT_JSONL"
done <"$TMP_FILTERED"

jsonl_entries="$(jq -s 'length' "$OUT_JSONL")"
jsonl_size_bytes="$(stat -c%s "$OUT_JSONL" 2>/dev/null || stat -f%z "$OUT_JSONL")"

# Distribuicao por linguagem.
lang_distribution="$(jq -s 'group_by(.language) | map({language:.[0].language,count:length})' "$OUT_JSONL")"

jq -n \
  --arg generated_at "$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
  --arg repo_root "$REPO_ROOT" \
  --argjson tracked_files_total "$(wc -l < "$TMP_LIST" | tr -d ' ')" \
  --argjson selected_files_total "$FILE_COUNT" \
  --argjson exported_entries "$jsonl_entries" \
  --argjson total_source_bytes "$total_source_bytes" \
  --argjson jsonl_size_bytes "$jsonl_size_bytes" \
  --argjson redacted_entries "$redacted_entries" \
  --argjson languages "$lang_distribution" \
  --arg include_regex "$INCLUDE_REGEX" \
  --arg exclude_regex "$EXCLUDE_REGEX" \
  '{
    generated_at:$generated_at,
    repo_root:$repo_root,
    tracked_files_total:$tracked_files_total,
    selected_files_total:$selected_files_total,
    exported_entries:$exported_entries,
    total_source_bytes:$total_source_bytes,
    jsonl_size_bytes:$jsonl_size_bytes,
    redacted_entries:$redacted_entries,
    include_regex:$include_regex,
    exclude_regex:$exclude_regex,
    languages:$languages
  }' >"$OUT_META"

# Validacao simples de parse jsonl.
jq -c . "$OUT_JSONL" >/dev/null

echo "KB JSONL: $OUT_JSONL"
echo "KB META:  $OUT_META"
echo "Entradas: $jsonl_entries"
echo "Redacoes: $redacted_entries"
