#!/usr/bin/make -f

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::') # grab everything after the space in "github.com/tendermint/tendermint v0.34.7"
DOCKER := $(shell which docker)
BUILDDIR ?= $(CURDIR)/build
TEST_DOCKER_REPO=jackzampolin/ununifitest

export GO111MODULE = on

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

ifeq (cleveldb,$(findstring cleveldb,$(JPYX_BUILD_OPTIONS)))
  build_tags += gcc cleveldb
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=ununifi \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=ununifid \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
			-X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)

ifeq (cleveldb,$(findstring cleveldb,$(JPYX_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (,$(findstring nostrip,$(JPYX_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(JPYX_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

#$(info $$BUILD_FLAGS is [$(BUILD_FLAGS)])

# The below include contains the tools target.
include docs/devtools/Makefile

###############################################################################
###                              Documentation                              ###
###############################################################################

all: install lint test

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

build-reproducible: go.sum
	$(DOCKER) rm latest-build || true
	$(DOCKER) run --volume=$(CURDIR):/sources:ro \
        --env TARGET_PLATFORMS='linux/amd64 darwin/amd64 linux/arm64 windows/amd64' \
        --env APP=ununifid \
        --env VERSION=$(VERSION) \
        --env COMMIT=$(COMMIT) \
        --env LEDGER_ENABLED=$(LEDGER_ENABLED) \
        --name latest-build cosmossdk/rbuilder:latest
	$(DOCKER) cp -a latest-build:/home/builder/artifacts/ $(CURDIR)/

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

build-contract-tests-hooks:
	mkdir -p $(BUILDDIR)
	go build -mod=readonly $(BUILD_FLAGS) -o $(BUILDDIR)/ ./cmd/contract_tests

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

draw-deps:
	@# requires brew install graphviz or apt-get install graphviz
	go get github.com/RobotsAndPencils/goviz
	@goviz -i ./cmd/ununifid -d 2 | dot -Tpng -o dependency-graph.png

clean:
	rm -rf $(BUILDDIR)/ artifacts/

distclean: clean
	rm -rf vendor/

###############################################################################
###                                 Tests                                   ###
###############################################################################

test: test-unit-all

test-unit-all:
	@VERSION=$(VERSION) go test  ./... 

NAMES = nftmarket auction cdp nftmint

test-unit: $(addprefix test-unit-, $(NAMES))

test-unit-%:
	@go test ./x/${@:test-unit-%=%}/...

mocks: $(MOCKS_DIR)
	docker run -v "$(CURDIR):/app"  -u `stat -c "%u:%g" $(CURDIR)` -w /app ekofr/gomock:go-1.17 /bin/sh scripts/utils/mockgen.sh

.PHONY: mocks

###############################################################################
###                                Protobuf                                 ###
###############################################################################

# Ensure following files are up-to-date based cosmos-sdk version
# https://github.com/cosmos/cosmos-sdk/blob/main/buf.work.yaml
# https://github.com/cosmos/cosmos-sdk/blob/main/scripts/protocgen.sh
# https://github.com/cosmos/cosmos-sdk/tree/main/proto/buf.*

protoVer=0.11.6
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh


proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@$(protoImage) sh ./scripts/protoc-swagger-gen.sh

proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

proto-lint:
	@$(protoImage) buf lint --error-format=json

proto-check-breaking:
	@$(protoImage) buf breaking --against $(HTTPS_GIT)#branch=main

CMT_URL              = https://raw.githubusercontent.com/cometbft/cometbft/v0.37.0/proto/tendermint

CMT_CRYPTO_TYPES     = proto/tendermint/crypto
CMT_ABCI_TYPES       = proto/tendermint/abci
CMT_TYPES            = proto/tendermint/types
CMT_VERSION          = proto/tendermint/version
CMT_LIBS             = proto/tendermint/libs/bits
CMT_P2P              = proto/tendermint/p2p

proto-update-deps:
	@echo "Updating Protobuf dependencies"

	@mkdir -p $(CMT_ABCI_TYPES)
	@curl -sSL $(CMT_URL)/abci/types.proto > $(CMT_ABCI_TYPES)/types.proto

	@mkdir -p $(CMT_VERSION)
	@curl -sSL $(CMT_URL)/version/types.proto > $(CMT_VERSION)/types.proto

	@mkdir -p $(CMT_TYPES)
	@curl -sSL $(CMT_URL)/types/types.proto > $(CMT_TYPES)/types.proto
	@curl -sSL $(CMT_URL)/types/evidence.proto > $(CMT_TYPES)/evidence.proto
	@curl -sSL $(CMT_URL)/types/params.proto > $(CMT_TYPES)/params.proto
	@curl -sSL $(CMT_URL)/types/validator.proto > $(CMT_TYPES)/validator.proto
	@curl -sSL $(CMT_URL)/types/block.proto > $(CMT_TYPES)/block.proto

	@mkdir -p $(CMT_CRYPTO_TYPES)
	@curl -sSL $(CMT_URL)/crypto/proof.proto > $(CMT_CRYPTO_TYPES)/proof.proto
	@curl -sSL $(CMT_URL)/crypto/keys.proto > $(CMT_CRYPTO_TYPES)/keys.proto

	@mkdir -p $(CMT_LIBS)
	@curl -sSL $(CMT_URL)/libs/bits/types.proto > $(CMT_LIBS)/types.proto

	@mkdir -p $(CMT_P2P)
	@curl -sSL $(CMT_URL)/p2p/types.proto > $(CMT_P2P)/types.proto

	$(DOCKER) run --rm -v $(CURDIR)/proto:/workspace --workdir /workspace $(protoImageName) buf mod update

.PHONY: proto-all proto-gen proto-swagger-gen proto-format proto-lint proto-check-breaking proto-update-deps
