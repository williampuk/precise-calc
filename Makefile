# Makefile for precise calculator

.PHONY: fmt vet lint build test clean

# Format code
fmt:
	gofmt -w .

# Run go vet for static analysis
vet:
	go vet ./...

# Lint code (fmt + vet)
lint: fmt vet

# Build the calculator
build:
	go build -o calculator ./cmd/calculator

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -f calculator
	go clean ./...