.PHONY: build run clean test docker-up docker-down

# Build the application
build:
	go build -o bin/api ./cmd

# Run the application
run:
	go run ./cmd/main.go

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Run tests
test:
	go test -v ./...

# Download dependencies
deps:
	go mod download
	go mod tidy

# Start MongoDB with Docker
docker-up:
	docker run -d -p 27017:27017 --name mongodb mongo:latest

# Stop MongoDB
docker-down:
	docker stop mongodb
	docker rm mongodb
