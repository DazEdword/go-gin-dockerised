# Gogin Dockerised
A multi-container Gin app template with Postgres instance defined via docker-compose.

## Database setup

The user and database `gogin` will be created by the Postgres' init script automatically.

## Migrations

Go Migrate and CLI required. See more details [here](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md).

### Create migration
```
migrate create -ext sql -dir db/migrations -seq <your-migration-name>
```

### Run migration
```
source .env
POSTGRESQL_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```
