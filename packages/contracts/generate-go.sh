#!/bin/bash

# Verifica se o usuário passou o nome do contrato como argumento
if [ -z "$1" ]; then
  echo "⚠️ Erro: Você precisa informar o nome do contrato."
  echo "💡 Uso: ./generate-go.sh <NomeDoContrato>"
  echo "Exemplo: ./generate-go.sh USPCertificate"
  exit 1
fi

CONTRACT_NAME=$1
OUT_FILE="out/${CONTRACT_NAME}.sol/${CONTRACT_NAME}.json"
DEST_DIR="../relayer-bundler/pkg/contracts"

echo "🚀 Iniciando geração de bindings Go para: $CONTRACT_NAME"

# 1. Compila os contratos para garantir que estamos usando a última versão
echo "🔨 Compilando contratos com Foundry..."
forge build --silent
if [ $? -ne 0 ]; then
  echo "❌ Erro na compilação do Foundry. Verifique seus arquivos .sol."
  exit 1
fi

# Verifica se o artefato foi gerado
if [ ! -f "$OUT_FILE" ]; then
  echo "❌ Artefato não encontrado: $OUT_FILE"
  echo "O nome do contrato está correto? (Sensível a maiúsculas/minúsculas)"
  exit 1
fi

# 2. Extrai o ABI e o Bytecode limpos usando jq
echo "📄 Extraindo ABI e Bytecode puros..."
jq '.abi' "$OUT_FILE" > "${CONTRACT_NAME}.abi"
jq -r '.bytecode.object' "$OUT_FILE" > "${CONTRACT_NAME}.bin"

# 3. Roda o abigen
echo "⚙️ Rodando abigen..."
abigen --abi "${CONTRACT_NAME}.abi" --bin "${CONTRACT_NAME}.bin" --pkg contracts --type "$CONTRACT_NAME" --out "${CONTRACT_NAME}.go"

if [ $? -ne 0 ]; then
  echo "❌ Erro na geração do abigen."
  rm -f "${CONTRACT_NAME}.abi" "${CONTRACT_NAME}.bin"
  exit 1
fi

# 4. Move o arquivo para o backend em Go
echo "🚚 Movendo ${CONTRACT_NAME}.go para a pasta do Relayer..."
mkdir -p "$DEST_DIR"
mv "${CONTRACT_NAME}.go" "$DEST_DIR/"

# 5. Limpa os arquivos temporários
echo "🧹 Limpando arquivos temporários..."
rm -f "${CONTRACT_NAME}.abi" "${CONTRACT_NAME}.bin"

echo "✅ Sucesso! O arquivo ${CONTRACT_NAME}.go está pronto para uso no seu backend."