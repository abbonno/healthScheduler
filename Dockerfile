# Usar la imagen base de Go
FROM golang:latest

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar archivos de configuración de Go
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY *.go ./

# Instalar task runner si es necesario
RUN go install github.com/go-task/task/v2/cmd/task@latest

# Comando para ejecutar las pruebas
CMD ["go", "test", "./..."]