.PHONY: run build test clean lint

# Default target
all: build

# Build the application
build:
	go build -o bin/derbyops main.go

# Run the application
run:
	go run main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f derbyops

# Run linter
lint:
	golangci-lint run

# Install dependencies
deps:
	go mod download

# Generate API documentation
docs:
	swag init

# Run database migrations
migrate:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

# Rollback database migrations
rollback:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

# Create a new migration
migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name 