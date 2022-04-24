
CONTAINER_NAME := template-backend
PORT := 5432
DB_NAME := template
DB_URL := postgres://root:secret@localhost:$(PORT)/$(DB_NAME)?sslmode=disable


postgres:
	docker run --name $(CONTAINER_NAME) -p $(PORT):$(PORT) -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it $(CONTAINER_NAME) createdb --username=root --owner=root $(DB_NAME)

dropdb:
	docker exec -it $(CONTAINER_NAME) dropdb --username=root --owner=root $(DB_NAME)

startdb:
	docker start $(CONTAINER_NAME)

stopdb:
	docker stop $(CONTAINER_NAME)

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/tornvallalexander/go-backend-template/db/sqlc Store

# migrate create -ext sql -dir db/migration -seq init_schema

.PHONY: postgres startdb stopdb createdb dropdb migrateup migratedown sqlc test db_docs db_schema server mock