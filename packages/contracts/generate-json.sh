#!/bin/bash

if [ -z "$1" ]; then
  echo "⚠️ Erro: Você precisa informar o nome do contrato."
  echo "💡 Uso: ./generate-go.sh <NomeDoContrato>"
  exit 1
fi

CONTRACT_NAME=$1
OUT_FILE="out/${CONTRACT_NAME}.sol/${CONTRACT_NAME}.json"

# --- CONFIGURAÇÃO DE PASTAS ---
# Ajuste estes caminhos conforme a estrutura do seu projeto
GO_DEST_DIR="../relayer-bundler/pkg/contracts"
FRONTEND_ABI_DIR="../dapp/usp-coin-interface/src/abi" # <-- AJUSTE AQUI PARA A SUA PASTA DO REACT

echo "🚀 Iniciando atualização de bindings e ABIs para: $CONTRACT_NAME"

echo "🔨 Compilando contratos com Foundry..."
forge build --silent
if [ $? -ne 0 ]; then
  echo "❌ Erro na compilação do Foundry."
  exit 1
fi

if [ ! -f "$OUT_FILE" ]; then
  echo "❌ Artefato não encontrado: $OUT_FILE"
  exit 1
fi

echo "📄 Extraindo dados..."
jq '.abi' "$OUT_FILE" > "${CONTRACT_NAME}.abi"
jq -r '.bytecode.object' "$OUT_FILE" > "${CONTRACT_NAME}.bin"

# Exportando para o GO
echo "⚙️ Gerando código Go..."
abigen --abi "${CONTRACT_NAME}.abi" --bin "${CONTRACT_NAME}.bin" --pkg contracts --type "$CONTRACT_NAME" --out "${CONTRACT_NAME}.go"

if [ $? -eq 0 ]; then
  mkdir -p "$GO_DEST_DIR"
  mv "${CONTRACT_NAME}.go" "$GO_DEST_DIR/"
  echo "✅ Go Binding atualizado em $GO_DEST_DIR"
else
  echo "❌ Erro na geração do abigen."
fi

# Exportando para o FRONTEND
echo "🌐 Atualizando ABI do Frontend..."
mkdir -p "$FRONTEND_ABI_DIR"
# Usamos a mesma extração do .abi, mas salvamos como .json lá no frontend
jq '.abi' "$OUT_FILE" > "${FRONTEND_ABI_DIR}/${CONTRACT_NAME}.json"
echo "✅ ABI JSON copiado para $FRONTEND_ABI_DIR"

echo "🧹 Limpando arquivos temporários..."
rm -f "${CONTRACT_NAME}.abi" "${CONTRACT_NAME}.bin"

echo "🎉 Processo concluído com sucesso!"