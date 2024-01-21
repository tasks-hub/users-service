#!/bin/bash

# Read values from files
POSTGRES_HOST=$(cat "$POSTGRES_HOST_FILE")
POSTGRES_DB=$(cat "$POSTGRES_DB_FILE")
POSTGRES_USER=$(cat "$POSTGRES_USER_FILE")
POSTGRES_PASSWORD=$(cat "$POSTGRES_PASSWORD_FILE")

# Export environment variables
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="user=$POSTGRES_USER password=$POSTGRES_PASSWORD host=$POSTGRES_HOST database=$POSTGRES_DB sslmode=disable"
export GOOSE_DIR=migrations

# Run goose migration
goose up
