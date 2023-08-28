include .env

LOCAL_DB_CONTAINER_NAME=bank-transfer-postgres
GO_BANK_TRANSFER_IMAGE=go-bank-transfer_bank-transfer:latest
POSTGRES_IMAGE=postgres:15.4
MIGRATE_IMAGE=migrate/migrate:v4.15.2

.PHONY: run stop docker-build docker-run clean-docker rm-all-images rm-volume tests swagger db-up db-down db-create db-drop db-migrate

# Run on docker-compose
run:
	docker-compose up

stop:
	docker-compose down

## Delete all docker data
clean-docker: stop rm-all-images rm-volume

rm-all-images:
	docker rmi $(GO_BANK_TRANSFER_IMAGE) && docker rmi $(POSTGRES_IMAGE) && docker rmi $(MIGRATE_IMAGE)

rm-volume:
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
	docker run -d --rm --name $(LOCAL_DB_CONTAINER_NAME) -p5432:5432 -v $(CURDIR)/db/data:/var/lib/postgresql/data -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASS) $(POSTGRES_IMAGE)
	sleep 2

db-down:
	docker stop $(LOCAL_DB_CONTAINER_NAME)

db-delete-volume:
	sudo rm -rf $(CURDIR)/db/data

db-create:
	docker exec -it $(LOCAL_DB_CONTAINER_NAME) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

db-drop:
	docker exec -it $(LOCAL_DB_CONTAINER_NAME) dropdb --username=$(DB_USER) $(DB_NAME)

## 'Migrate' required, see README - run migration over db to apply new changes
db-migrate:
	migrate -source file://db/migrations -database postgres://$(DB_USER):$(DB_PASS)@localhost:5432/$(DB_NAME)?sslmode=disable up

## 'Sqlc' required, see README - converts queries in .sql files to golang files
db-queries:
	sqlc generate