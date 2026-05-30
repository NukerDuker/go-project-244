# Bash
.PHONY: build test lint golangci-lint

build:
	go build -o bin/gendiff ./cmd/gendiff

test:
	go test -v -cover ./...

lint:
	golangci-lint run