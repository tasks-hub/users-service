# Usa la imagen oficial de Golang como base
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Instala Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copia los archivos necesarios para las migraciones
COPY . .

# Establece la variable de entorno para indicar la ubicación del archivo de configuración de la base de datos
ENV GOOSE_DRIVER=postgres \
    GOOSE_DBSTRING="user=admin password=admin_password host=db database=users_db sslmode=disable" \
    GOOSE_DIR=migrations

# Comando por defecto para ejecutar migraciones (ajusta según tus necesidades)
CMD ["goose", "up"]
