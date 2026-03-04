#!/bin/bash

echo "🚀 Iniciando testes automatizados via Docker..."

# Usamos a imagem do golang:1.25-alpine para manter a consistência exata com o Dockerfile.
# Mapeamos o diretório atual ($PWD) para /app dentro do container.
docker run --rm -v "$PWD":/app -w /app golang:1.25-alpine \
  sh -c "go mod tidy && go test -v ./..."

echo "✅ Processo de testes finalizado."