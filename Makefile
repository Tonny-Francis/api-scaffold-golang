.PHONY: run build test

run:
	GO111MODULE=on air

build:
	GO111MODULE=on go build -o tmp/main

test:
	GO111MODULE=on go test