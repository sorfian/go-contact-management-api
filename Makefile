.PHONY: help dev build run test test-unit test-integration clean wire install

# Default target
help:
	@echo "Available commands:"
	@echo "  make install           - Install all dependencies"
	@echo "  make wire              - Generate Wire dependency injection code"
	@echo "  make dev               - Run application in development mode"
	@echo "  make build             - Build the application binary"
	@echo "  make run               - Run the compiled binary"
	@echo "  make test              - Run all tests"
	@echo "  make test-unit         - Run unit tests"
	@echo "  make test-integration  - Run integration tests"
	@echo "  make clean             - Clean build artifacts"

# Install dependencies
install:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy
	@echo "Dependencies installed successfully!"

# Generate Wire code
wire:
	@echo "Generating Wire dependency injection code..."
	go run github.com/google/wire/cmd/wire
	cd test && go run github.com/google/wire/cmd/wire
	@echo "Wire code generated successfully!"

# Run in development mode
dev:
	@echo "Running in development mode..."
	go run .

# Build the application
build:
	@echo "Building application..."
	go build -o bin/app .
	@echo "Build completed! Binary: bin/app"

# Run the compiled binary
run: build
	@echo "Running application..."
	./bin/app

# Run all tests
test:
	@echo "Running all tests..."
	go test -v ./...

# Run unit tests only
test-unit:
	@echo "Running unit tests..."
	go test -v -short ./...

# Run integration tests
test-integration:
	@echo "Running integration tests..."
	cd test && go test -v

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -f app.exe
	@echo "Clean completed!"
