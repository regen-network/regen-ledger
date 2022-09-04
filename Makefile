#!/usr/bin/make -f

export GO111MODULE=on

BUILD_DIR ?= $(CURDIR)/build
REGEN_CMD := $(CURDIR)/cmd/regen

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
DB_BACKEND ?= goleveldb

###############################################################################
###                            Build Tags/Flags                             ###
###############################################################################

# process build tags

build_tags = netgo

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

ifeq ($(DB_BACKEND), goleveldb)
  build_tags += goleveldb
endif

ifeq ($(DB_BACKEND), cleveldb)
  build_tags += gcc
  build_tags += cleveldb
endif

ifeq ($(DB_BACKEND), boltdb)
  build_tags += boltdb
endif

ifeq ($(DB_BACKEND), rocksdb)
  build_tags += rocksdb
endif

ifeq ($(DB_BACKEND), badgerdb)
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

ifeq ($(DB_BACKEND), goleveldb)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=goleveldb
endif

ifeq ($(DB_BACKEND), cleveldb)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif

ifeq ($(DB_BACKEND), boltdb)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb
endif

ifeq ($(DB_BACKEND), rocksdb)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif

ifeq ($(DB_BACKEND), badgerdb)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb
endif

ldflags += -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)

ifeq ($(NO_STRIP),false)
  ldflags += -w -s
endif

ldflags += $(LD_FLAGS)
ldflags := $(strip $(ldflags))

# set build flags

BUILD_FLAGS := -tags '$(build_tags)' -ldflags '$(ldflags)'

ifeq ($(NO_STRIP),false)
  BUILD_FLAGS += -trimpath
endif

###############################################################################
###                             Build / Install                             ###
###############################################################################

all: build

build: go.sum go-version
	@mkdir -p $(BUILD_DIR)
	go build -mod=readonly -o $(BUILD_DIR) $(BUILD_FLAGS) $(REGEN_CMD)

build-linux: go.sum go-version
	GOOS=linux GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

install: go.sum go-version
	go install -mod=readonly $(BUILD_FLAGS) $(REGEN_CMD)

.PHONY: build build-linux install

###############################################################################
###                               Go Version                                ###
###############################################################################

GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
MIN_GO_MAJOR_VERSION = 1
MIN_GO_MINOR_VERSION = 18
GO_VERSION_ERROR = Golang version $(GO_MAJOR_VERSION).$(GO_MINOR_VERSION) is not supported, \
please update to at least $(MIN_GO_MAJOR_VERSION).$(MIN_GO_MINOR_VERSION)

go-version:
	@echo "Verifying go version..."
	@if [ $(GO_MAJOR_VERSION) -gt $(MIN_GO_MAJOR_VERSION) ]; then \
		exit 0; \
	elif [ $(GO_MAJOR_VERSION) -lt $(MIN_GO_MAJOR_VERSION) ]; then \
		echo $(GO_VERSION_ERROR); \
		exit 1; \
	elif [ $(GO_MINOR_VERSION) -lt $(MIN_GO_MINOR_VERSION) ]; then \
		echo $(GO_VERSION_ERROR); \
		exit 1; \
	fi

.PHONY: go-version

###############################################################################
###                               Go Modules                                ###
###############################################################################

go.sum: go.mod
	@echo "Ensuring app dependencies have not been modified..."
	go mod verify
	go mod tidy

verify:
	@echo "Verifying all go module dependencies..."
	@find . -name 'go.mod' -type f -execdir go mod verify \;

tidy:
	@echo "Cleaning up all go module dependencies..."
	@find . -name 'go.mod' -type f -execdir go mod tidy \;

.PHONY: verify tidy

###############################################################################
###                               Go Generate                               ###
###############################################################################

generate:
	@echo "Generating source files from directives..."
	@find . -name 'go.mod' -type f -execdir go generate ./... \;

.PHONY: generate

###############################################################################
###                             Lint / Format                               ###
###############################################################################

lint:
	@echo "Linting all go modules..."
	@find . -name 'go.mod' -type f -execdir golangci-lint run --out-format=tab \;

lint-fix: format
	@echo "Attempting to fix lint errors in all go modules..."
	@find . -name 'go.mod' -type f -execdir golangci-lint run --fix --out-format=tab --issues-exit-code=0 \;

format_filter = -name '*.go' -type f \
	-not -name '*.pb.go' \
	-not -name '*.gw.go' \
	-not -name '*.pulsar.go' \
	-not -name '*.cosmos_orm.go' \
	-not -name '*statik.go'

format_local = \
	github.com/tendermint/tendermint \
	github.com/cosmos/cosmos-sdk \
	github.com/cosmos/ibc-go \
	github.com/regen-network/regen-ledger

format:
	@echo "Formatting all go modules..."
	@find . $(format_filter) | xargs gofmt -s -w
	@find . $(format_filter) | xargs goimports -w -local $(subst $(whitespace),$(comma),$(format_local))
	@find . $(format_filter) | xargs misspell -w

.PHONY: lint lint-fix format

###############################################################################
###                                  Tools                                  ###
###############################################################################

tools: go-version
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/client9/misspell/cmd/misspell@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/yoheimuta/protolint/cmd/protolint@latest

.PHONY: tools

###############################################################################
###                                Protobuf                                 ###
###############################################################################

include make/proto.mk

###############################################################################
###                                  Tests                                  ###
###############################################################################

include make/tests.mk

###############################################################################
###                               Simulations                               ###
###############################################################################

include make/sims.mk

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
	@echo "Wait a few seconds and then visit http://localhost:6060/pkg/github.com/regen-network/regen-ledger/v4/"
	godoc -http=:6060

.PHONY: docs-dev docs-build godocs

###############################################################################
###                                Swagger                                  ###
###############################################################################

swagger: proto-swagger-gen
	@./scripts/generate_swagger_docs.sh

.PHONY: swagger

###############################################################################
###                                Localnet                                 ###
###############################################################################

DOCKER := $(shell which docker)
LOCALNET_DIR = $(CURDIR)/.testnets

localnet-build-env:
	$(MAKE) -C contrib/images regen-env

localnet-build-nodes:
	$(DOCKER) run --rm -v $(CURDIR)/.testnets:/data regen-ledger/regen-env \
			  testnet init-files --v 4 -o /data --starting-ip-address 192.168.10.2 --keyring-backend=test
	docker-compose up -d

# localnet-start will run a 4-node testnet locally. The nodes are
# based off the docker images in: ./contrib/images/regen-env
localnet-start: localnet-stop localnet-build-env localnet-build-nodes

localnet-stop:
	docker-compose down -v

.PHONY: localnet-start localnet-stop localnet-build-nodes localnet-build-env

###############################################################################
###                                 Clean                                   ###
###############################################################################

clean: test-clean
	rm -rf $(BUILD_DIR)
	rm -rf $(LOCALNET_DIR)

.PHONY: clean
