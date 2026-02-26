DB_URL=postgres://postgres:postgres@localhost:5432/golang_todo?sslmode=disable

migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" down

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

swagger:
	swag init -g cmd/main.go