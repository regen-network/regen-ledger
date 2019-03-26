export GO111MODULE=on

install:
	go install ./cmd/xrnd
	go install ./cmd/xrncli

test:
	go test ./...

test_cover:
	bash -x tests/test_cover.sh

lint:
	go get -u golang.org/x/lint/golint
	${GOPATH}/bin/golint -set_exit_status ./...
