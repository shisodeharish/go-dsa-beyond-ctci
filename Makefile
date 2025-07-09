.DEFAULT_GOAL := test

build:
	go build ./...

test:
	go test ./...

fmt:
	go fmt ./...
