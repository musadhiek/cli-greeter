APP_NAME := greeter
BIN_DIR := bin
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-X main.Version=$(VERSION)"

.PHONY: all run test lint build clean install fmt vet check

all: check build

run:
	go run .

test:
	go test ./... -v -race -coverprofile=coverage.out

lint:
	golangci-lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

check: fmt vet lint test

build:
	mkdir -p $(BIN_DIR)
	go build $(LDFLAGS) -o $(BIN_DIR)/$(APP_NAME)

install:
	go install $(LDFLAGS)

clean:
	rm -rf $(BIN_DIR)
	rm -f coverage.out

.DEFAULT_GOAL := all
