.PHONY: run test lint

-include .env
export

run:
	go run ./cmd

test:
	go test ./...

lint:
	go vet ./...
