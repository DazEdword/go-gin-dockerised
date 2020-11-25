# Gin + Postgres Dockerised
A multi-container Gin app template with Postgres instance defined via docker-compose.

## Basic commands
See Makefile:
```
make up         # Bring up the stack
make down       # Stop the stack
make db         # Bring up the Postgres service online (e.g. to run or test the app directly, undockerised)
make test       # Run tests
```

Bring all services up with `make up` and then visit `http://localhost:8080/`.

## Database setup
The user and database `gogin` will be created by the Postgres' init script automatically.

## Migrations
Go Migrate and CLI required. See more details [here](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md).

### Create migration
```
migrate create -ext sql -dir db/migrations -seq <your-migration-name>
```

### Run migration
Migrations are set up to run automatically within the `start.sh` script. They can also be run manually:

```
source .env     # Not required when running inside Docker.
POSTGRESQL_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```
