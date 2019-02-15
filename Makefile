export GO111MODULE=on

install:
	go install ./cmd/xrnd
	go install ./cmd/xrncli

test:
	go test ./... -godog.strict

lint:
	[ -x $(which golint) ] || go get -u golang.org/x/lint/golint
	golint ./...
