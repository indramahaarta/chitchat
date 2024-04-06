run-postgres: 
	docker run --name chitchat-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:12.17-alpine3.19

postgres-up: 
	docker start chitchat-db

postgres-down: 
	docker stop chitchat-db

create-db: 
	docker exec -it chitchat-db createdb --username=root --owner=root chitchat_db

drop-db: 
	docker exec -it chitchat-db dropdb -U root chitchat_db

migrate-up: 
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/chitchat_db?sslmode=disable" --verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/chitchat_db?sslmode=disable" --verbose down

sqlc:
	sqlc generate

server-dev:
	air

server-prod:
	go run main.go

swagger:
	swag i

test:
	go test ./... -cover -v