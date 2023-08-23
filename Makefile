DB_CONTAINER = msfelipe-postgres
DB_USER = master
DB_PASS = greenbeans
DB_NAME = bank-transfer

.PHONY: db-up db-down db-create db-drop db-migrate

# Database commands
db-up:
	docker run --rm -it --name $(DB_CONTAINER) -p5432:5432 -v $(CURDIR)/db/data:/var/lib/postgresql/data -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASS) postgres:15.4

db-down:
	docker stop $(DB_CONTAINER)

db-create:
	docker exec -it $(DB_CONTAINER) createdb --username=master --owner=master $(DB_NAME)

db-drop:
	docker exec -it $(DB_CONTAINER) dropdb --username=master $(DB_NAME)

db-migrate:
	migrate -source file://db/migrations -database postgres://$(DB_USER):$(DB_PASS)@localhost:5432/$(DB_NAME)?sslmode=disable up
