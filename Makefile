.PHONY: all install clean build test

# Default target
all: build

# Build the binary
build:
	go build -o dist/sai ./src/cmd/sai

# Install the binary to /usr/local/bin
install: build
	cp dist/sai /usr/local/bin/
	chmod +x /usr/local/bin/sai

# Clean build artifacts
clean:
	rm -rf dist
	go clean

# Run tests
test:
	go test ./...

# Format code
format:
	go fmt ./...

# Check code style
lint:
	go vet ./... 