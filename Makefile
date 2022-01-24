#!/usr/bin/make -f

PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags))
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
BINDIR ?= $(GOPATH)/bin
BUILDDIR ?= $(CURDIR)/build
APP_DIR = ./app
MOCKS_DIR = $(CURDIR)/tests/mocks
HTTPS_GIT := https://github.com/regen-network/regen-ledger.git
DOCKER_BUF := docker run -v $(shell pwd):/workspace --workdir /workspace bufbuild/buf:1.0.0-rc11
PROJECT_NAME = $(shell git remote get-url origin | xargs basename -s .git)
DOCKER := $(shell which docker)

export GO111MODULE = on

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

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += gcc
endif


whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=regen \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=regen \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (badgerdb,$(findstring badgerdb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb
  BUILD_TAGS += badgerdb
endif
# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(TENDERMINT_BUILD_OPTIONS)))
  CGO_ENABLED=1
  BUILD_TAGS += rocksdb
endif
# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(TENDERMINT_BUILD_OPTIONS)))
  BUILD_TAGS += boltdb
endif

ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

all: tools build lint test

# The below include contains the tools and runsim targets.
include contrib/devtools/Makefile

###############################################################################
###                                  Build                                  ###
###############################################################################

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/
build-linux:
	GOOS=linux GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

build-regen-all: go.sum
	$(if $(shell docker inspect -f '{{ .Id }}' tendermint/xrnnode 2>/dev/null),$(info found image tendermint/xrnnode),docker pull tendermint/xrnnode:latest)
	docker rm latest-build || true
	docker run --volume=$(CURDIR):/sources:ro \
        --env TARGET_OS='darwin linux windows' \
        --env APP=regen \
        --env VERSION=$(VERSION) \
        --env COMMIT=$(COMMIT) \
        --env LEDGER_ENABLED=$(LEDGER_ENABLED) \
        --name latest-build tendermint/xrnnode:latest
	docker cp -a latest-build:/home/builder/artifacts/ $(CURDIR)/

build-regen-linux: go.sum $(BUILDDIR)/
	$(if $(shell docker inspect -f '{{ .Id }}' tendermint/xrnnode 2>/dev/null),$(info found image tendermint/xrnnode),docker pull tendermint/xrnnode:latest)
	docker rm latest-build || true
	docker run --volume=$(CURDIR):/sources:ro \
        --env TARGET_OS='linux' \
        --env APP=regen \
        --env VERSION=$(VERSION) \
        --env COMMIT=$(COMMIT) \
        --env LEDGER_ENABLED=false \
        --name latest-build tendermint/xrnnode:latest
	docker cp -a latest-build:/home/builder/artifacts/ $(CURDIR)/
	cp artifacts/regen-*-linux-amd64 $(BUILDDIR)/regen

.PHONY: build build-linux build-regen-all build-regen-linux

mockgen_cmd=go run github.com/golang/mock/mockgen

mocks: $(MOCKS_DIR)
	$(mockgen_cmd) -source=client/account_retriever.go -package mocks -destination tests/mocks/account_retriever.go
	$(mockgen_cmd) -package mocks -destination tests/mocks/tendermint_tm_db_DB.go github.com/tendermint/tm-db DB
	$(mockgen_cmd) -source=types/module/module.go -package mocks -destination tests/mocks/types_module_module.go
	$(mockgen_cmd) -source=types/invariant.go -package mocks -destination tests/mocks/types_invariant.go
	$(mockgen_cmd) -source=types/router.go -package mocks -destination tests/mocks/types_router.go
	$(mockgen_cmd) -package mocks -destination tests/mocks/grpc_server.go github.com/gogo/protobuf/grpc Server
	$(mockgen_cmd) -package mocks -destination tests/mocks/tendermint_tendermint_libs_log_DB.go github.com/tendermint/tendermint/libs/log Logger
.PHONY: mocks

$(MOCKS_DIR):
	mkdir -p $(MOCKS_DIR)

distclean: clean tools-clean
clean:
	rm -rf \
    $(BUILDDIR)/ \
    artifacts/ \
    tmp-swagger-gen/

.PHONY: distclean clean

###############################################################################
###                          Tools & Dependencies                           ###
###############################################################################

go.sum: go.mod
	echo "Ensure dependencies have not been modified ..." >&2
	go mod verify
	go mod tidy

###############################################################################
###                              Documentation                              ###
###############################################################################

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGenSwagger}$$"; then docker start -a $(containerProtoGenSwagger); else docker run --name $(containerProtoGenSwagger) -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) \
		sh ./scripts/protoc-swagger-gen.sh; fi
	
update-swagger-docs: statik
	$(BINDIR)/statik -src=client/docs/swagger-ui -dest=client/docs -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
		echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
		exit 1;\
	else \
		echo "\033[92mSwagger docs are in sync\033[0m";\
	fi
.PHONY: update-swagger-docs

godocs:
	@echo "--> Wait a few seconds and visit http://localhost:6060/pkg/github.com/cosmos/cosmos-sdk/types"
	godoc -http=:6060

# This builds a docs site for each branch/tag in `./docs/versions`
# and copies each site to a version prefixed path. The last entry inside
# the `versions` file will be the default root index.html.
build-docs:
	@cd docs && \
	while read -a p; do \
		branch=$${p[0]} ; \
		path_prefix=$${p[1]} ; \
		(git checkout $${branch} && npm install && VUEPRESS_BASE="/$${path_prefix}/" npm run build) ; \
		mkdir -p ~/output/$${path_prefix} ; \
		cp -r .vuepress/dist/* ~/output/$${path_prefix}/ ; \
		cp ~/output/$${path_prefix}/index.html ~/output ; \
	done < versions ;

sync-docs:
	cd ~/output && \
	echo "role_arn = ${DEPLOYMENT_ROLE_ARN}" >> /root/.aws/config ; \
	echo "CI job = ${CIRCLE_BUILD_URL}" >> version.html ; \
	aws s3 sync . s3://${WEBSITE_BUCKET} --profile terraform --delete ; \
	aws cloudfront create-invalidation --distribution-id ${CF_DISTRIBUTION_ID} --profile terraform --path "/*" ;
.PHONY: sync-docs

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

test: test-unit
test-all: test-unit test-ledger-mock test-race test-cover

TEST_PACKAGES=./...
TEST_TARGETS := test-unit test-unit-amino test-unit-proto test-ledger-mock test-race test-ledger test-race

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
		cd ${CURRENT_DIR}/$$module; \
		go test -mod=readonly -json $(ARGS) $(TEST_PACKAGES) ./... | tparse; \
	done
else
	@echo "Unit tests"; \
	for module in $(SUB_MODULES); do \
		cd ${CURRENT_DIR}/$$module; \
		go test -mod=readonly $(ARGS) $(TEST_PACKAGES) ./... ; \
	done
endif

.PHONY: run-tests test test-all $(TEST_TARGETS)

test-cover:
	@export VERSION=$(VERSION);
	@bash scripts/test_cover.sh
.PHONY: test-cover

benchmark:
	@go test -mod=readonly -bench=. $(PACKAGES_NOSIMULATION)
.PHONY: benchmark

###############################################################################
###                                Linting                                  ###
###############################################################################

lint:
	golangci-lint run --out-format=tab

lint-fix:
	golangci-lint run --fix --out-format=tab --issues-exit-code=0
.PHONY: lint lint-fix

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs goimports -w -local github.com/cosmos/cosmos-sdk
.PHONY: format

###############################################################################
###                                 Devdoc                                  ###
###############################################################################

DEVDOC_SAVE = docker commit `docker ps -a -n 1 -q` devdoc:local

devdoc-init:
	docker run -it -v "$(CURDIR):/go/src/github.com/cosmos/cosmos-sdk" -w "/go/src/github.com/cosmos/cosmos-sdk" tendermint/devdoc echo
	# TODO make this safer
	$(call DEVDOC_SAVE)

devdoc:
	docker run -it -v "$(CURDIR):/go/src/github.com/cosmos/cosmos-sdk" -w "/go/src/github.com/cosmos/cosmos-sdk" devdoc:local bash

devdoc-save:
	# TODO make this safer
	$(call DEVDOC_SAVE)

devdoc-clean:
	docker rmi -f $$(docker images -f "dangling=true" -q)

devdoc-update:
	docker pull tendermint/devdoc

.PHONY: devdoc devdoc-clean devdoc-init devdoc-save devdoc-update

###############################################################################
###                                Protobuf                                 ###
###############################################################################

containerProtoVer=v0.2
containerProtoImage=tendermintdev/sdk-proto-gen:$(containerProtoVer)
containerProtoGen=${PROJECT_NAME}-proto-gen-$(containerProtoVer)
containerProtoFmt=${PROJECT_NAME}-proto-fmt-$(containerProtoVer)
containerProtoGenSwagger=${PROJECT_NAME}-proto-gen-swagger-$(containerProtoVer)

proto-all: proto-gen proto-lint proto-check-breaking proto-format
.PHONY: proto-all proto-gen proto-gen-docker proto-lint proto-check-breaking proto-format

proto-gen:
	@echo "Generating Protobuf files"
	@echo "If you're having trouble with this command, you need to install the latest buf, protoc-gen-gocosmos, protoc-gen-grpc-gateway, protoc-gen-go-pulsar, and protoc-gen-go-grpc locally"
#	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) \
#		sh ./scripts/protocgen.sh; fi
	./scripts/protocgen.sh

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; ; fi

proto-format-direct:
	find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \;

proto-lint:
	@$(DOCKER_BUF) lint --error-format=json

proto-lint-direct:
	@buf lint --error-format=json

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=master

proto-check-breaking-direct:
	@buf breaking --against '.git#branch=master'


GOGO_PROTO_URL   = https://raw.githubusercontent.com/regen-network/protobuf/cosmos
REGEN_COSMOS_PROTO_URL = https://raw.githubusercontent.com/regen-network/cosmos-proto/master
COSMOS_PROTO_URL   = https://raw.githubusercontent.com/cosmos/cosmos-sdk/master/proto/cosmos

GOGO_PROTO_TYPES    = third_party/proto/gogoproto
REGEN_COSMOS_PROTO_TYPES  = third_party/proto/cosmos_proto
COSMOS_PROTO_TYPES    = third_party/proto/cosmos

proto-update-deps:
	@mkdir -p $(GOGO_PROTO_TYPES)
	@curl -sSL $(GOGO_PROTO_URL)/gogoproto/gogo.proto > $(GOGO_PROTO_TYPES)/gogo.proto

	@mkdir -p $(REGEN_COSMOS_PROTO_TYPES)
	@curl -sSL $(REGEN_COSMOS_PROTO_URL)/cosmos.proto > $(REGEN_COSMOS_PROTO_TYPES)/cosmos.proto

	@mkdir -p $(COSMOS_PROTO_TYPES)/base/query/v1beta1/
	@curl -sSL $(COSMOS_PROTO_URL)/base/query/v1beta1/pagination.proto > $(COSMOS_PROTO_TYPES)/base/query/v1beta1/pagination.proto
	@curl -sSL $(COSMOS_PROTO_URL)/base/v1beta1/coin.proto > $(COSMOS_PROTO_TYPES)/base/v1beta1/coin.proto


###############################################################################
###                                Localnet                                 ###
###############################################################################

# Run a 4-node testnet locally via docker compose
localnet-start: build-linux localnet-stop
	$(if $(shell $(DOCKER) inspect -f '{{ .Id }}' regenledger/regen-env 2>/dev/null),$(info found image regenledger/regen-env),$(MAKE) -C contrib/images regen-env)
	if ! test -f build/node0/regen/config/genesis.json; then $(DOCKER) run --rm \
		--user $(shell id -u):$(shell id -g) \
		-v $(BUILDDIR):/regen:Z \
		-v /etc/group:/etc/group:ro \
		-v /etc/passwd:/etc/passwd:ro \
		-v /etc/shadow:/etc/shadow:ro \
		regenledger/regen-env testnet init-files --v 4 -o . --starting-ip-address 192.168.10.2 --keyring-backend=test ; fi
	docker-compose up -d

localnet-stop:
	docker-compose down -v 

.PHONY: localnet-start localnet-stop


include sims.mk
