%:
	@:

build:
	@go build -o bin/ecom-api-rest-go cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom-api-rest-go

migration:
	@migrate create -ext sql -dir db/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run db/migrations/main.go up

migrate-down:
	@go run db/migrations/main.go down
