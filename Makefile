.PHONY: build test

build:
	go build -o procstackql

test:
	go test ./...
