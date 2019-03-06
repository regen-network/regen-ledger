export GO111MODULE=on

install:
	go install ./cmd/xrnd
	go install ./cmd/xrncli

test:
	go test ./... -godog.strict

integration_test:
	@go test -p 4 `go list github.com/cosmos/cosmos-sdk/cmd/gaia/cli_test` -tags=cli_test

lint:
	go get -u golang.org/x/lint/golint
	${GOPATH}/bin/golint -set_exit_status ./...
