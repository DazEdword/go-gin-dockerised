up:
	docker-compose up --build -d

down:
	docker-compose down

.PHONY: db
db:
	docker-compose up --build -d gogin-postgres

.PHONY: test
test:
	docker-compose run --rm goginapp go test /go/src/app/... -cover
