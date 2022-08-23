#!/usr/bin/make -f

###############################################################################
###                                Protobuf                                 ###
###############################################################################

containerProtoVer=v0.7
containerProtoImage=tendermintdev/sdk-proto-gen:$(containerProtoVer)
containerProtoGen=regen-ledger-proto-gen-$(containerProtoVer)
containerProtoFmt=regen-ledger-proto-fmt-$(containerProtoVer)
containerProtoGenSwagger=regen-ledger-proto-gen-swagger-$(containerProtoVer)

proto-all: proto-lint proto-format proto-gen proto-check-breaking

proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) \
		sh ./scripts/protocgen.sh; fi

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find .  -name "*.proto" -not -path "./third_party/*" -exec clang-format -i {} \; ; fi

DOCKER_BUF := docker run -v $(shell pwd):/workspace --workdir /workspace bufbuild/buf:1.0.0-rc11

proto-lint:
	@$(DOCKER_BUF) lint --error-format=json
	@protolint .

proto-lint-fix:
	@protolint -fix .

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against https://github.com/regen-network/regen-ledger.git#branch=main

GOGO_PROTO_URL           = https://raw.githubusercontent.com/regen-network/protobuf/cosmos
GOOGLE_PROTO_URL         = https://raw.githubusercontent.com/googleapis/googleapis/master
REGEN_COSMOS_PROTO_URL   = https://raw.githubusercontent.com/regen-network/cosmos-proto/master
COSMOS_PROTO_URL         = https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.45.4/proto/cosmos
COSMOS_ORM_PROTO_URL     = https://raw.githubusercontent.com/cosmos/cosmos-sdk/orm/v1.0.0-alpha.10/proto/cosmos

GOGO_PROTO_TYPES         = third_party/proto/gogoproto
GOOGLE_PROTO_TYPES       = third_party/proto/google
REGEN_COSMOS_PROTO_TYPES = third_party/proto/cosmos_proto
COSMOS_PROTO_TYPES       = third_party/proto/cosmos

proto-swagger-deps:
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

proto-swagger-gen: proto-swagger-deps
	@echo "Generating Protobuf Swagger"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGenSwagger}$$"; then docker start -a $(containerProtoGenSwagger); else docker run --name $(containerProtoGenSwagger) -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) \
		sh ./scripts/protoc-swagger-gen.sh; fi

.PHONY: proto-all proto-gen proto-format proto-lint proto-check-breaking proto-swagger-deps proto-swagger-gen
