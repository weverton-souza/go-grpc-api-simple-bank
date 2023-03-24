DB_USER=root
DB_PASSWORD=root

mysql:
	docker run --name mysql-simple-bank -p 3306:3306 -e MYSQL_ROOT_PASSWORD="${DB_PASSWORD}" -d mysql:8.0-debian
create-db:
	docker exec -it mysql-simple-bank mysql --user="${DB_USER}" --password="${DB_PASSWORD}" -e 'create database simple_bank'

drop-db:
	docker exec -it mysql-simple-bank mysql --user="${DB_USER}" --password="${DB_PASSWORD}" -e 'drop database simple_bank'

migrate-up:
	migrate -path db/migration -database "mysql://root:root@tcp(localhost)/simple_bank" -verbose up

migrate-down:
	migrate -path db/migration -database "mysql://root:root@tcp(localhost)/simple_bank" -verbose down

.PHONY: postgres createdb dropdb
