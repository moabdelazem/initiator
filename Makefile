APP_NAME=initiator

VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "0.1.0")
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
PKG := github.com/moabdelazem/initiator
LDFLAGS := -ldflags "-X '$(PKG)/internal/version.Version=$(VERSION)' -X '$(PKG)/internal/version.Commit=$(COMMIT)' -X '$(PKG)/internal/version.BuildDate=$(BUILD_DATE)'"

.PHONY: build
# Build the binary
build:
	go build $(LDFLAGS) -o bin/$(APP_NAME)

.PHONY: install
install:
	go install $(LDFLAGS)

.PHONY: start
# Build the binary & run the tests & run the binary
start:
	go build $(LDFLAGS) -o bin/$(APP_NAME)
	go test -v ./...
	./bin/$(APP_NAME)

.PHONY: test
# Run the tests
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf bin/