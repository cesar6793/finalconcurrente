# Usamos la imagen base de Alpine
FROM alpine:latest

# Instalamos las dependencias necesarias
RUN apk update && apk add --no-cache ca-certificates

# Configuramos el directorio de trabajo
WORKDIR /app

# Copiamos el ejecutable al contenedor
COPY kmeans-api .

# Definimos el comando por defecto al iniciar el contenedor
CMD ["./kmeans-api"]
