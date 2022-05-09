#!/usr/bin/make -f

export GO111MODULE=on

BIN_DIR ?= $(GOPATH)/bin
BUILD_DIR ?= $(CURDIR)/build
REGEN_DIR := $(CURDIR)/app/regen

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

SDK_VERSION := $(shell go list -m github.com/cosmos/cosmos-sdk | sed 's:.* ::')
TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')

LEDGER_ENABLED ?= true

###############################################################################
###                            Build Tags/Flags                             ###
###############################################################################

# process build tags

build_tags = netgo

ifeq ($(EXPERIMENTAL),true)
	build_tags += experimental
endif

ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(REGEN_BUILD_OPTIONS)))
  build_tags += gcc
  build_tags += cleveldb
endif

ifeq (boltdb,$(findstring boltdb,$(REGEN_BUILD_OPTIONS)))
  build_tags += boltdb
endif

ifeq (rocksdb,$(findstring rocksdb,$(REGEN_BUILD_OPTIONS)))
  build_tags += rocksdb
endif

ifeq (badgerdb,$(findstring badgerdb,$(REGEN_BUILD_OPTIONS)))
  build_tags += badgerdb
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

# process linker flags

empty :=
whitespace := $(empty) $(empty)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=regen \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=regen \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)

ifeq (cleveldb,$(findstring cleveldb,$(build_tags)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif

ifeq (boltdb,$(findstring boltdb,$(build_tags)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb
endif

ifeq (rocksdb,$(findstring rocksdb,$(build_tags)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif

ifeq (badgerdb,$(findstring badgerdb,$(build_tags)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb
endif

ldflags += -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)

ifeq (,$(findstring nostrip,$(REGEN_BUILD_OPTIONS)))
  ldflags += -w -s
endif

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

# set build flags

BUILD_FLAGS := -tags '$(build_tags)' -ldflags '$(ldflags)'

ifeq (,$(findstring nostrip,$(REGEN_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

###############################################################################
###                             Build / Install                             ###
###############################################################################

all: build

build:
	mkdir -p $(BUILD_DIR)
	go build -mod=readonly -o $(BUILD_DIR) $(BUILD_FLAGS) $(REGEN_DIR)

build-linux:
	GOOS=linux GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

install:
	go install -mod=readonly $(BUILD_FLAGS) $(REGEN_DIR)

.PHONY: build build-linux install

###############################################################################
###                               Go Modules                                ###
###############################################################################

verify:
	@find . -name 'go.mod' -type f -execdir go mod verify \;

tidy:
	@find . -name 'go.mod' -type f -execdir go mod tidy \;

generate:
	@find . -name 'go.mod' -type f -execdir go generate ./... \;

.PHONY: verify tidy generate

###############################################################################
###                                  Tools                                  ###
###############################################################################

clean:
	rm -rf $(BUILD_DIR) artifacts

.PHONY: clean

include contrib/devtools/Makefile

###############################################################################
###                              Documentation                              ###
###############################################################################

docs-dev:
	@echo "Starting regen-ledger static documentation site..."
	@cd docs && yarn && yarn dev

docs-build:
	@echo "Building regen-ledger static documentation site..."
	@cd docs && yarn && yarn build

godocs:
	@echo "Wait a few seconds and then visit http://localhost:6060/pkg/github.com/regen-network/regen-ledger/v3/"
	godoc -http=:6060

.PHONY: docs-dev docs-build godocs

###############################################################################
###                                Swagger                                  ###
###############################################################################

swagger: statik proto-update-deps
	./scripts/protoc-swagger-gen.sh

.PHONY: swagger

###############################################################################
###                               Simulation                                ###
###############################################################################

include sims.mk

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

###############################################################################
###                             Lint / Format                               ###
###############################################################################

lint:
	golangci-lint run --out-format=tab

lint-fix:
	golangci-lint run --fix --out-format=tab --issues-exit-code=0

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs goimports -w -local github.com/cosmos/cosmos-sdk

.PHONY: lint lint-fix format

###############################################################################
###                                Protobuf                                 ###
###############################################################################

containerProtoVer=v0.7
containerProtoImage=tendermintdev/sdk-proto-gen:$(containerProtoVer)
containerProtoGen=regen-ledger-proto-gen-$(containerProtoVer)
containerProtoFmt=regen-ledger-proto-fmt-$(containerProtoVer)

proto-all: proto-gen proto-lint proto-check-breaking proto-format

proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) \
		sh ./scripts/protocgen.sh; fi

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; ; fi

DOCKER_BUF := docker run -v $(shell pwd):/workspace --workdir /workspace bufbuild/buf:1.0.0-rc11

proto-lint:
	@$(DOCKER_BUF) lint --error-format=json

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against https://github.com/regen-network/regen-ledger.git#branch=master

GOGO_PROTO_URL           = https://raw.githubusercontent.com/regen-network/protobuf/cosmos
GOOGLE_PROTO_URL         = https://raw.githubusercontent.com/googleapis/googleapis/master
REGEN_COSMOS_PROTO_URL   = https://raw.githubusercontent.com/regen-network/cosmos-proto/master
COSMOS_PROTO_URL         = https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.45.4/proto/cosmos
COSMOS_ORM_PROTO_URL     = https://raw.githubusercontent.com/cosmos/cosmos-sdk/orm/v1.0.0-alpha.10/proto/cosmos

GOGO_PROTO_TYPES         = third_party/proto/gogoproto
GOOGLE_PROTO_TYPES       = third_party/proto/google
REGEN_COSMOS_PROTO_TYPES = third_party/proto/cosmos_proto
COSMOS_PROTO_TYPES       = third_party/proto/cosmos

proto-update-deps:
	@mkdir -p $(GOGO_PROTO_TYPES)
	@curl -sSL $(GOGO_PROTO_URL)/gogoproto/gogo.proto > $(GOGO_PROTO_TYPES)/gogo.proto

	@mkdir -p $(GOOGLE_PROTO_TYPES)/api/
	@curl -sSL $(GOOGLE_PROTO_URL)/google/api/annotations.proto > $(GOOGLE_PROTO_TYPES)/api/annotations.proto
	@curl -sSL $(GOOGLE_PROTO_URL)/google/api/http.proto > $(GOOGLE_PROTO_TYPES)/api/http.proto

	@mkdir -p $(REGEN_COSMOS_PROTO_TYPES)
	@curl -sSL $(REGEN_COSMOS_PROTO_URL)/cosmos.proto > $(REGEN_COSMOS_PROTO_TYPES)/cosmos.proto

	@mkdir -p $(COSMOS_PROTO_TYPES)/base/v1beta1/
	@curl -sSL $(COSMOS_PROTO_URL)/base/v1beta1/coin.proto > $(COSMOS_PROTO_TYPES)/base/v1beta1/coin.proto

	@mkdir -p $(COSMOS_PROTO_TYPES)/base/query/v1beta1/
	@curl -sSL $(COSMOS_PROTO_URL)/base/query/v1beta1/pagination.proto > $(COSMOS_PROTO_TYPES)/base/query/v1beta1/pagination.proto

	@mkdir -p $(COSMOS_PROTO_TYPES)/orm/v1alpha1/
	@curl -sSL $(COSMOS_ORM_PROTO_URL)/orm/v1alpha1/orm.proto > $(COSMOS_PROTO_TYPES)/orm/v1alpha1/orm.proto

.PHONY: proto-all proto-gen proto-format proto-lint proto-check-breaking proto-update-deps

###############################################################################
###                                Localnet                                 ###
###############################################################################

DOCKER := $(shell which docker)

localnet-build-env:
	$(MAKE) -C contrib/images regen-env

localnet-build-nodes:
	$(DOCKER) run --rm -v $(CURDIR)/.testnets:/data regenledger/regen-env \
			  testnet init-files --v 4 -o /data --starting-ip-address 192.168.10.2 --keyring-backend=test
	docker-compose up -d

# localnet-start will run a 4-node testnet locally. The nodes are
# based off the docker images in: ./contrib/images/regen-env
localnet-start: localnet-stop localnet-build-env localnet-build-nodes

localnet-stop:
	docker-compose down -v

.PHONY: localnet-start localnet-stop localnet-build-nodes localnet-build-env
