postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=aswedD4321 -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root gold_bank

dropdb:
	docker exec -it postgres17 dropdb gold_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:aswedD4321@localhost:5432/gold_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:aswedD4321@localhost:5432/gold_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:aswedD4321@localhost:5432/gold_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:aswedD4321@localhost:5432/gold_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go -imports "gomock=go.uber.org/mock/gomock" gitlab.com/xfx1/goldbank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
