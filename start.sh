#! /usr/bin/env bash

# Migrate db
echo "Migrating db..."
POSTGRESQL_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

migrate -database ${POSTGRESQL_URL} -path db/migrations up

# Start app
echo "Starting app..."
go run ./cmd/goginapp/app.go
