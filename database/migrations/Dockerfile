# Usa la imagen oficial de Golang como base
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Instala Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

ARG POSTGRES_DB
ARG POSTGRES_USER
ARG POSTGRES_PASSWORD

# Copia los archivos necesarios para las migraciones
COPY . .

# Establece la variable de entorno para indicar la ubicación del archivo de configuración de la base de datos
ENV GOOSE_DRIVER=postgres \
    GOOSE_DBSTRING="user=$POSTGRES_USER password=$POSTGRES_PASSWORD host=db database=$POSTGRES_DB sslmode=disable" \
    GOOSE_DIR=migrations

# Comando por defecto para ejecutar migraciones (ajusta según tus necesidades)
CMD ["goose", "up"]
