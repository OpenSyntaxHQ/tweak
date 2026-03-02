.PHONY: verify test race generate build

generate:
	go generate ./...

build:
	go build ./...

test:
	go test -count=1 ./...

race:
	go test -race -count=1 ./...

verify: generate build test race
