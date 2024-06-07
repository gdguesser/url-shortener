.PHONY: all test test-coverage run build clean

# Variables
APP_NAME := url-shortener
PKG := ./...

# Default target
all: test build

# Run tests with coverage
test:
	go test -v $(PKG)

test-coverage:
	go test -coverprofile=coverage.out $(PKG)
	go tool cover -html=coverage.out -o coverage.html

# Run the application
run:
	go run cmd/main.go

# Build the application
build:
	go build -o bin/$(APP_NAME) cmd/main.go

# Clean up generated files
clean:
	go clean
	rm -f coverage.out coverage.html
	rm -rf bin
