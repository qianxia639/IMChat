DB_URL=postgresql://root:root@localhost:5432/imchat?sslmode=disable

proto: 
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		proto/*.proto

server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go

postgres:
	docker run -d -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root --restart=always --name postgres postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root imchat

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migrateup1:
	migrate -path db/migration -database "${DB_URL}" -verbose up 1

migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

migratedown1:
	migrate -path db/migration -database "${DB_URL}" -verbose down 1

new_migrate:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

evans:
	evans --host localhost --port 9090 -r repl

mock:
	mockgen -package mockdb -destination db/mock/store.go IMChat/db/sqlc Store

test:
	go test -cover -race ./... -count=1

.PHONY: proto server client postgres createdb migrateup migrateup1 migratedown migratedown1 new_migrate sqlc evans mock test