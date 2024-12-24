include .env
export

up:
	docker compose up --build -d

down:
	docker compose down

.PHONY: db
db:
	docker compose up --build -d gogin-postgres

.PHONY: test
test:
	docker compose run --rm --entrypoint="go test ./... -cover" goginapp
