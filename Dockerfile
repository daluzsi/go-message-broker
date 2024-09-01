# Usar a imagem base do Golang
FROM golang:1.21 AS builder

# Definir o diretório de trabalho
WORKDIR /app

# Copiar o código fonte para o contêiner
COPY . .

# Compilar o aplicativo
RUN GOOS=linux GOARCH=amd64 go build -o app main.go

# Usar uma imagem base mínima para a execução
FROM debian:latest

# Instalar dependências necessárias
RUN apt-get update && apt-get install -y ca-certificates

# Definir o diretório de trabalho
WORKDIR /root/

# Copiar o aplicativo compilado e o arquivo de configuração para o contêiner
COPY --from=builder /app/app .
COPY --from=builder /app/application-properties.yaml .

# Garantir que o binário tenha permissões de execução
RUN chmod +x ./app

# Definir o comando de inicialização do contêiner
CMD ["./app"]
