SHELL=/bin/bash
export SHELL

PHONY: build test

build:
	go build
test:
	go test
