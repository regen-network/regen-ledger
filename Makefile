export GO111MODULE=on

install:
	go install ./cmd/xrnd
	go install ./cmd/xrncli

test:
	go test ./... -godog.strict

lint:
	go get -u golang.org/x/lint/golint
	${GOPATH}/bin/golint -set_exit_status ./...
