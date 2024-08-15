.PHONY: build
build:
	go build -o http_rest_api -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
