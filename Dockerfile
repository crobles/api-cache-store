# Etapa 1: Construcción
FROM golang:1.20 AS builder

# Establece el directorio de trabajo
WORKDIR /app

# Copia go.mod y go.sum e instala las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el código fuente de la aplicación
COPY . .

# Copila la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Etapa 2: Imagen final
FROM alpine:latest

# Instala ca-certificates para permitir conexiones HTTPS
RUN apk --no-cache add ca-certificates

# Establece el directorio de trabajo
WORKDIR /root/

# Copia el binario desde la etapa de construcción
COPY --from=builder /app/myapp .

# Copia el archivo .env
COPY .env .env

# Instala `bash` para permitir la ejecución de scripts y carga de variables de entorno
RUN apk --no-cache add bash

# Expone el puerto que utiliza la aplicación (cambiar según sea necesario)
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["sh", "-c", "source .env && ./myapp"]