postgres:
	docker run --name postgres12 -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

create-db:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

drop-db:
	docker exec -it postgres12 dropdb simple_bank

migrate-up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/simple_bank?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb
