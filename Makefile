build:
	docker-compose build

run:
	docker-compose up

test:
	go test -v ./...

api:
	docker-compose up go

postgres:
	docker-compose up postgres

cache:
	docker-compose up cache

migrate:
	docker-compose up migrate
