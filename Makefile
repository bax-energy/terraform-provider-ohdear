# Makefile for Terraform Provider farsight

# Build the provider
build:
	go build -o terraform-provider-farsight

# Run tests
test:
	go test ./...

# Clean up build artifacts
clean:
	go clean
	rm -f terraform-provider-farsight

# Install the provider
install: build
	mv terraform-provider-farsight %APPDATA%/.terraform.d/plugins/

# Format the code
fmt:
	go fmt ./...

# Lint the code
lint:
	golangci-lint run

# Default target
all: build test lint