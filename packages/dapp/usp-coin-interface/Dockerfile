# ==========================================
# Estágio 1: Build da aplicação (Node.js)
# ==========================================
FROM node:20-alpine AS build

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos de dependência e instala
COPY package.json package-lock.json ./
RUN npm install

# Copia o resto do código do projeto
COPY . .

# Roda o comando do Vite para gerar a pasta /dist com os arquivos otimizados
RUN npm run build

# ==========================================
# Estágio 2: Servir a aplicação (Nginx)
# ==========================================
FROM nginx:alpine

# Copia o arquivo de configuração customizado do Nginx que criamos
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copia os arquivos estáticos gerados no Estágio 1 para a pasta pública do Nginx
COPY --from=build /app/dist /usr/share/nginx/html

# Expõe a porta 80 do container
EXPOSE 80

# Inicia o Nginx
CMD ["nginx", "-g", "daemon off;"]