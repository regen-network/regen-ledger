
GORELEASER_SKIP_VALIDATE ?= false
GORELEASER_DEBUG         ?= false
GORELEASER_RELEASE       ?= false
GORELEASER_IMAGE         := ghcr.io/goreleaser/goreleaser-cross:v1.19.0
GO_MOD                 ?= readonly
BUILD_TAGS             ?= osusergo,netgo,ledger,static_build
GO_LINKMODE            ?= external
RELEASE_TAG           ?= $(shell git describe --tags --abbrev=0)
GORELEASER_MOUNT_CONFIG  ?= false


ifeq ($(GORELEASER_MOUNT_CONFIG),true)
	GORELEASER_IMAGE := -v $(HOME)/.docker/config.json:/root/.docker/config.json $(GORELEASER_IMAGE)
endif


ifeq ($(IS_MAINNET), true)
	ifeq ($(IS_PREREL), false)
		GORELEASER_HOMEBREW_NAME   := regen
		GORELEASER_HOMEBREW_CUSTOM :=
		IS_STABLE                  := true
	else
		GORELEASER_HOMEBREW_NAME   := regen-test
		GORELEASER_HOMEBREW_CUSTOM := keg_only :unneeded, \"This is testnet release. Run brew install regen-network/tap/regen to install mainnet version\"
	endif
else
	GORELEASER_HOMEBREW_NAME       := regen-edge
	GORELEASER_HOMEBREW_CUSTOM     := keg_only :unneeded, \"This is edgenet release. Run brew install regen-network/tap/regen to install mainnet version\"
endif

GORELEASER_BUILD_VARS := \
-X github.com/cosmos/cosmos-sdk/version.Name=regen \
-X github.com/cosmos/cosmos-sdk/version.AppName=regen \
-X github.com/cosmos/cosmos-sdk/version.BuildTags=\"$(BUILD_TAGS)\" \
-X github.com/cosmos/cosmos-sdk/version.Version=$(RELEASE_TAG) \
-X github.com/cosmos/cosmos-sdk/version.Commit=$(GIT_HEAD_COMMIT_LONG)

ldflags = -linkmode=$(GO_LINKMODE) -X github.com/cosmos/cosmos-sdk/version.Name=regen \
-X github.com/cosmos/cosmos-sdk/version.AppName=regen \
-X github.com/cosmos/cosmos-sdk/version.BuildTags="$(BUILD_TAGS)" \
-X github.com/cosmos/cosmos-sdk/version.Version=$(shell git describe --tags | sed 's/^v//') \
-X github.com/cosmos/cosmos-sdk/version.Commit=$(GIT_HEAD_COMMIT_LONG)

# check for nostrip option
ifeq (,$(findstring nostrip,$(BUILD_OPTIONS)))
	ldflags                += -s -w
	GORELEASER_STRIP_FLAGS += -s -w
endif

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -mod=$(GO_MOD) -tags='$(BUILD_TAGS)' -ldflags '$(ldflags)'


ifeq ($(GORELEASER_RELEASE),true)
	GORELEASER_SKIP_VALIDATE := false
	GORELEASER_SKIP_PUBLISH  := release --skip-publish=false
else
	GORELEASER_SKIP_PUBLISH  := --skip-publish=true
	GORELEASER_SKIP_VALIDATE ?= false
	GITHUB_TOKEN=
endif

$(AKASH): modvendor
	$(GO) build -o $@ $(BUILD_FLAGS) ./cmd/regen

.PHONY: release

GO_MOD_NAME                  := $(shell go list -m 2>/dev/null)

release:
	docker run \
		--rm \
		-e STABLE=false \
		-e MOD="$(GO_MOD)" \
		-e BUILD_TAGS="$(BUILD_TAGS)" \
		-e BUILD_VARS="$(GORELEASER_BUILD_VARS)" \
		-e STRIP_FLAGS="$(GORELEASER_STRIP_FLAGS)" \
		-e LINKMODE="$(GO_LINKMODE)" \
		-e HOMEBREW_NAME="$(GORELEASER_HOMEBREW_NAME)" \
		-e HOMEBREW_CUSTOM="$(GORELEASER_HOMEBREW_CUSTOM)" \
		-e GITHUB_TOKEN="$(GITHUB_TOKEN)" \
		-e GORELEASER_CURRENT_TAG="$(RELEASE_TAG)" \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(GO_MOD_NAME) \
		-w /go/src/$(GO_MOD_NAME) \
		$(GORELEASER_IMAGE) \
		-f "$(GORELEASER_CONFIG)" \
		$(GORELEASER_SKIP_PUBLISH) \
		--skip-validate=$(GORELEASER_SKIP_VALIDATE) \
		--debug=$(GORELEASER_DEBUG) \
		--rm-dist \
