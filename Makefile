DB_USER=root
DB_PASSWORD=root
DB_SERVER=localhost
DB_DATABASE=simple_bank

mysql:
	docker run --name mysql-simple-bank -p 3306:3306 -e MYSQL_ROOT_PASSWORD="${DB_PASSWORD}" -d mysql:8.0-debian

create-db:
	docker exec -it mysql-simple-bank mysql --user="${DB_USER}" --password="${DB_PASSWORD}" -e 'create database ${simple_bank}'

drop-db:
	docker exec -it mysql-simple-bank mysql --user="${DB_USER}" --password="${DB_PASSWORD}" -e 'drop database ${simple_bank}'

migrate-up:
	migrate -path db/migration -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_SERVER})/${simple_bank}?parseTime=true" -verbose up

migrate-down:
	migrate -path db/migration -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_SERVER})/${simple_bank}?parseTime=true" -verbose down

sqlc:
	sqlc generate

test-clear:
	go clean -testcache

test:
	go test -v -cover ./...

.PHONY: mysql create-db drop-db migrate-up migrate-down sqlc test-clear test
