# ==========================================
# Estágio 1: Builder (Compilação)
# ==========================================
FROM golang:1.25-alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos de dependência primeiro.
# Isso é um truque para aproveitar o cache do Docker: se o go.mod não mudar,
# o Docker não baixa as dependências de novo a cada build.
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código fonte para dentro do container
COPY . .

# Compila o binário. 
# Apontamos para o caminho exato do seu main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o relayer-api ./cmd/server/main.go

# ==========================================
# Estágio 2: Runner (Execução)
# ==========================================
FROM alpine:latest

# Atualiza os pacotes e instala certificados CA para reduzir vulnerabilidades
RUN apk update && apk upgrade && apk --no-cache add ca-certificates

WORKDIR /app

# Copia apenas o binário compilado do estágio 1 para este novo container
COPY --from=builder /app/relayer-api .

# Exposição de porta (documentação para quem for ler o Dockerfile)
EXPOSE 8080

# Comando que será executado quando o container subir
CMD ["./relayer-api"]