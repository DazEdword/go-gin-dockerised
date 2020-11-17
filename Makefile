up:
	docker-compose up --build -d

down:
	docker-compose down

.PHONY: db
db:
	docker-compose up --build -d postgres
