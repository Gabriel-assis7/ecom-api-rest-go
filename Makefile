build:
	@go build -o bin/ecom-api-rest-go cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom-api-rest-go
