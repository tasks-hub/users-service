# Usa la imagen oficial de Golang como base
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Instala Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

ARG POSTGRES_HOST_FILE
ARG POSTGRES_DB_FILE
ARG POSTGRES_USER_FILE
ARG POSTGRES_PASSWORD_FILE

# Set environment variables directly
RUN POSTGRES_HOST=$(cat $POSTGRES_HOST_FILE) && \
    POSTGRES_DB=$(cat $POSTGRES_DB_FILE) && \
    POSTGRES_USER=$(cat $POSTGRES_USER_FILE) && \
    POSTGRES_PASSWORD=$(cat $POSTGRES_PASSWORD_FILE) && \
    GOOSE_DRIVER=postgres && \
    GOOSE_DBSTRING="user=$POSTGRES_USER password=$POSTGRES_PASSWORD host=$POSTGRES_HOST database=$POSTGRES_DB sslmode=disable" && \
    GOOSE_DIR=migrations

# Copia los archivos necesarios para las migraciones
COPY . .


# Comando por defecto para ejecutar migraciones (ajusta según tus necesidades)
ENTRYPOINT ["tail", "-f", "/dev/null"]

