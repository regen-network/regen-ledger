#!/usr/bin/make -f

###############################################################################
###                                 Tests                                   ###
###############################################################################

CURRENT_DIR=$(shell pwd)
GO_MODULES=$(shell find . -type f -name 'go.mod' -print0 | xargs -0 -n1 dirname | sort)

test: test-all

test-all:
	@for module in $(GO_MODULES); do \
		echo "Testing Module $$module"; \
		cd ${CURRENT_DIR}/$$module; \
		go test ./... -tags=nosimulation; \
	done

test-app:
	@echo "Testing Module ."
	@go test ./... -tags=nointegration,nosimulation \
		-coverprofile=coverage-app.out -covermode=atomic

test-types:
	@echo "Testing Module types"
	@cd types && go test ./... \
		-coverprofile=${CURRENT_DIR}/coverage-types.out -covermode=atomic

test-x-data:
	@echo "Testing Module x/data"
	@cd x/data && go test ./... \
		-coverprofile=${CURRENT_DIR}/coverage-x-data.out -covermode=atomic

test-x-ecocredit:
	@echo "Testing Module x/ecocredit"
	@cd x/ecocredit && go test ./... \
		-coverprofile=${CURRENT_DIR}/coverage-x-ecocredit.out -covermode=atomic

test-x-intertx:
	@echo "Testing Module x/intertx"
	@cd x/intertx && go test ./... \
		-coverprofile=${CURRENT_DIR}/coverage-x-intertx.out -covermode=atomic

test-integration:
	@echo "Testing Integration"
	@go test ./app/testsuite/... \
		-coverpkg=./... -coverprofile=coverage-integration.out -covermode=atomic

test-coverage:
	@cat coverage*.out | grep -v "mode: atomic" >> coverage.txt

test-clean:
	@go clean -testcache
	@find . -name 'coverage.txt' -delete
	@find . -name 'coverage*.out' -delete

.PHONY: test test-all test-app test-types test-x-data test-x-ecocredit \
	test-x-intertx test-integration test-coverage test-clean
