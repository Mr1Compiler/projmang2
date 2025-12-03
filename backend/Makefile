.PHONY: build run test clean dev

# Build the application
build:
	go build -o bin/api cmd/api/main.go

# Run the application
run:
	go run cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -cover ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Run with hot reload (requires air: go install github.com/air-verse/air@latest)
dev:
	air

# Tidy dependencies
tidy:
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run
