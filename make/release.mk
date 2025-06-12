#!/usr/bin/make -f
###############################################################################
###                                Release                                  ###
###############################################################################
release-help:
	@echo "release subcommands"
	@echo ""
	@echo "Usage:"
	@echo "  make release-[command]"
	@echo ""
	@echo "Available Commands:"
	@echo "  dry-run                   Perform a dry run release"
	@echo "  snapshot                  Create a snapshot release"

# Set a GoReleaser image tag. v1.23.0 is a recent v1 tag.
# If your .goreleaser.yml requires v2, use a v2 tag and update flags below.
# You can also define GO_VERSION elsewhere and use v$(GO_VERSION) if preferred.
GORELEASER_IMAGE := ghcr.io/goreleaser/goreleaser-cross:v1.22.0
COSMWASM_VERSION := $(shell go list -m github.com/CosmWasm/wasmvm | sed 's/.* //')
GO_MOD_NAME      := $(shell go list -m 2>/dev/null) # Get the go module name dynamically

release-dry-run:
	docker run \
		--rm \
		-e COSMWASM_VERSION=$(COSMWASM_VERSION) \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(GO_MOD_NAME) \
		-w /go/src/$(GO_MOD_NAME) \
		$(GORELEASER_IMAGE) \
		release \
		--clean \
		--skip=publish,validate 

release-snapshot:
	docker run \
		--rm \
		-e COSMWASM_VERSION=$(COSMWASM_VERSION) \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(GO_MOD_NAME) \
		-w /go/src/$(GO_MOD_NAME) \
		$(GORELEASER_IMAGE) \
		release \
		--clean \
		--snapshot \
		--skip=publish,validate 

# New target for actual publishing
release-publish:
	docker run \
		--rm \
		-e COSMWASM_VERSION=$(COSMWASM_VERSION) \
		-e GITHUB_TOKEN=$(GITHUB_TOKEN) \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(GO_MOD_NAME) \
		-w /go/src/$(GO_MOD_NAME) \
		$(GORELEASER_IMAGE) \
		release \
		--clean # Run goreleaser release command without skips

.PHONY: release-help release-dry-run release-snapshot release-publish
