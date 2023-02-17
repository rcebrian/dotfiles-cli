SHELL = /bin/bash
MAKEFLAGS += --silent

.PHONY: build
build: clean
	go build -o ./build/dotfiles-cli ./main.go

run:
	go run ./main.go

lint:
	golangci-lint run

lint-fix:
	gofmt -w -s .
	goimports -w .

clean:
	rm -rf ./build/*

test:
	go test ./...

test-coverage:
	go test ./... -cover
