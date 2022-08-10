#!/usr/bin/make -f

###############################################################################
###                                 Tests                                   ###
###############################################################################

TEST_PACKAGES=./...
TEST_TARGETS := test-unit test-unit-amino test-ledger test-ledger-mock test-race test-race

PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')

# Test runs-specific rules. To add a new test target, just add
# a new rule, customise ARGS or TEST_PACKAGES ad libitum, and
# append the new rule to the TEST_TARGETS list.
UNIT_TEST_ARGS		= cgo ledger test_ledger_mock norace
AMINO_TEST_ARGS		= ledger test_ledger_mock test_amino norace
LEDGER_TEST_ARGS	= cgo ledger norace
LEDGER_MOCK_ARGS	= ledger test_ledger_mock norace
TEST_RACE_ARGS		= cgo ledger test_ledger_mock
ifeq ($(EXPERIMENTAL),true)
	UNIT_TEST_ARGS		+= experimental
	AMINO_TEST_ARGS		+= experimental
	LEDGER_TEST_ARGS	+= experimental
	LEDGER_MOCK_ARGS	+= experimental
	TEST_RACE_ARGS		+= experimental
endif

test: test-unit

test-all: test-unit test-ledger-mock test-race test-cover

test-unit: ARGS=-tags='$(UNIT_TEST_ARGS)'
test-unit-amino: ARGS=-tags='${AMINO_TEST_ARGS}'
test-ledger: ARGS=-tags='${LEDGER_TEST_ARGS}'
test-ledger-mock: ARGS=-tags='${LEDGER_MOCK_ARGS}'
test-race: ARGS=-race -tags='${TEST_RACE_ARGS}'
test-race: TEST_PACKAGES=$(PACKAGES_NOSIMULATION)

$(TEST_TARGETS): run-tests

SUB_MODULES = $(shell find . -type f -name 'go.mod' -print0 | xargs -0 -n1 dirname | sort)
CURRENT_DIR = $(shell pwd)

run-tests:
ifneq (,$(shell which tparse 2>/dev/null))
	@echo "Unit tests"; \
	for module in $(SUB_MODULES); do \
		echo "Testing Module $$module"; \
		cd ${CURRENT_DIR}/$$module; \
		go test -mod=readonly -json $(ARGS) $(TEST_PACKAGES) ./... | tparse; \
	done
else
	@echo "Unit tests"; \
	for module in $(SUB_MODULES); do \
		echo "Testing Module $$module"; \
		cd ${CURRENT_DIR}/$$module; \
		go test -mod=readonly $(ARGS) $(TEST_PACKAGES) ./... ; \
	done
endif

test-cover:
	@export VERSION=$(VERSION);
	@bash scripts/test_cover.sh

benchmark:
	@go test -mod=readonly -bench=. $(PACKAGES_NOSIMULATION)

.PHONY: run-tests test test-all $(TEST_TARGETS) test-cover benchmark
