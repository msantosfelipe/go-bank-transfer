include .env

.PHONY: run stop docker-build docker-run remove-volume tests swagger db-up db-down db-create db-drop db-migrate

# Run on docker-compose
run:
	docker-compose up

stop:
	docker-compose down

## Delete all docker data
clean-docker: stop remove-all-images remove-volume

remove-all-images:
	docker rmi go-bank-transfer_bank-transfer:latest && docker rmi postgres:14.2 && docker rmi migrate/migrate:v4.15.2

remove-volume:
	docker volume rm go-bank-transfer_postgres-data

# Development:
tests:
	go test ./app/...

swagger:
	swag init -o infrastructure/swagger

## Run go app on docker
docker-build:
	docker build -t go-bank-transfer:latest .

docker-run:
	docker run -it --name bank-transfer -p 8080:8080 go-bank-transfer:latest

## Development database commands
init-dev-db: db-up db-create db-migrate

db-up:
	docker run -d --rm --name $(DB_DOCKER_CONTAINER_NAME) -p5432:5432 -v $(CURDIR)/db/data:/var/lib/postgresql/data -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASS) postgres:15.4
	sleep 1

db-down:
	docker stop $(DB_DOCKER_CONTAINER_NAME)

db-create:
	docker exec -it $(DB_DOCKER_CONTAINER_NAME) createdb --username=master --owner=master $(DB_NAME)

db-drop:
	docker exec -it $(DB_DOCKER_CONTAINER_NAME) dropdb --username=master $(DB_NAME)

## 'Migrate' required, see README - run migration over db to apply new changes
db-migrate:
	migrate -source file://db/migrations -database postgres://$(DB_USER):$(DB_PASS)@localhost:5432/$(DB_NAME)?sslmode=disable up

## 'Sqlc' required, see README - converts queries in .sql files to golang files
db-queries:
	sqlc generate