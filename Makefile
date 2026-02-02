# Makefile for crud-go project

.PHONY: help build run clean test lint docker-up docker-down

# Default target
help:
	@echo "Available targets:"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  clean       - Clean build artifacts"
	@echo "  test        - Run tests"
	@echo "  lint        - Run linters"
	@echo "  docker-up   - Start MongoDB with Docker"
	@echo "  docker-down - Stop MongoDB"

# Build the application
build:
	@echo "Building application..."
	@go build -o bin/api.exe ./cmd/api

# Run the application
run:
	@echo "Starting application..."
	@go run ./cmd/api/main.go

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

# Run linters
lint:
	@echo "Running linters..."
	@golangci-lint run

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

# Start MongoDB with Docker
docker-up:
	@echo "Starting MongoDB..."
	@docker run -d -p 27017:27017 --name mongodb mongo:latest

# Stop MongoDB
docker-down:
	@echo "Stopping MongoDB..."
	@docker stop mongodb
	@docker rm mongodb

# Run the application with hot reload (requires air)
dev:
	@echo "Starting development server with hot reload..."
	@air
