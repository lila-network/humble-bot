# Set the default Go build flags
#GOFLAGS = -ldflags='-w -s -X cmd.Version=$(VERSION)'
GOFLAGS = -ldflags='-w -s'

# Build the application
build:
	go build $(GOFLAGS) -o bin/humble-bot main.go 

# Run tests
test:
	go test -v ./...

# Clean the build artifacts
clean:
	rm -rf bin

# Set a version for the build
VERSION := $(shell git describe --tags --always)
