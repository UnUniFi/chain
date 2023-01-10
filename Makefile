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
PACKAGES_UNIT=$(shell go list ./... | grep -E -v 'tests/simulator|e2e')
PACKAGES_E2E=$(shell go list ./... | grep '/e2e')

test: test-unit
test-unit:
	@VERSION=$(VERSION) go test  ./... 

# test-e2e runs a full e2e test suite
# deletes any pre-existing Osmosis containers before running.
#
# Attempts to delete Docker resources at the end.
# May fail to do so if stopped mid way.
# In that case, run `make e2e-remove-resources`
# manually.
# Utilizes Go cache.
test-e2e-check:
	echo $(PACKAGES_E2E)

test-e2e: e2e-setup test-e2e-ci

# test-e2e-ci runs a full e2e test suite
# does not do any validation about the state of the Docker environment
# As a result, avoid using this locally.
test-e2e-ci:
	@VERSION=$(VERSION) OSMOSIS_E2E=True OSMOSIS_E2E_DEBUG_LOG=False OSMOSIS_E2E_UPGRADE_VERSION=$(E2E_UPGRADE_VERSION)  go test -mod=readonly -timeout=25m -v $(PACKAGES_E2E)

# test-e2e-debug runs a full e2e test suite but does
# not attempt to delete Docker resources at the end.
test-e2e-debug: e2e-setup
	@VERSION=$(VERSION) OSMOSIS_E2E=True OSMOSIS_E2E_DEBUG_LOG=True OSMOSIS_E2E_UPGRADE_VERSION=$(E2E_UPGRADE_VERSION) OSMOSIS_E2E_SKIP_CLEANUP=True OSMOSIS_E2E_SKIP_UPGRADE=True OSMOSIS_E2E_SKIP_IBC=True go test -mod=readonly -timeout=25m -v $(PACKAGES_E2E) -count=1

# test-e2e-short runs the e2e test with only short tests.
# Does not delete any of the containers after running.
# Deletes any existing containers before running.
# Does not use Go cache.
test-e2e-short: e2e-setup
	@VERSION=$(VERSION) OSMOSIS_E2E=True OSMOSIS_E2E_DEBUG_LOG=True OSMOSIS_E2E_SKIP_UPGRADE=True OSMOSIS_E2E_SKIP_IBC=True OSMOSIS_E2E_SKIP_STATE_SYNC=True OSMOSIS_E2E_SKIP_CLEANUP=True go test -mod=readonly -timeout=25m -v $(PACKAGES_E2E) -count=1

test-mutation:
	@bash scripts/mutation-test.sh $(MODULES)

benchmark:
	@go test -mod=readonly -bench=. $(PACKAGES_UNIT)

build-e2e-script:
	mkdir -p $(BUILDDIR)
	go build -mod=readonly $(BUILD_FLAGS) -o $(BUILDDIR)/ ./tests/e2e/initialization/$(E2E_SCRIPT_NAME)

docker-build-debug:
	@DOCKER_BUILDKIT=1 docker build -t ununifi:${COMMIT} --build-arg BASE_IMG_TAG=debug -f Dockerfile .
	@DOCKER_BUILDKIT=1 docker tag ununifi:${COMMIT} ununifi:debug

docker-build-debug-alpine:
	@DOCKER_BUILDKIT=1 docker build -t ununifi:${COMMIT} --build-arg BASE_IMG_TAG=debug --build-arg RUNNER_IMAGE=$(RUNNER_BASE_IMAGE_ALPINE) -f Dockerfile .
	@DOCKER_BUILDKIT=1 docker tag ununifi:${COMMIT} ununifi:debug

docker-build-e2e-init-chain:
	@DOCKER_BUILDKIT=1 docker build -t ununifi-e2e-init-chain:debug --build-arg E2E_SCRIPT_NAME=chain --platform=linux/x86_64 -f tests/e2e/initialization/init.Dockerfile .

docker-build-e2e-init-node:
	@DOCKER_BUILDKIT=1 docker build -t ununifi-e2e-init-node:debug --build-arg E2E_SCRIPT_NAME=node --platform=linux/x86_64 -f tests/e2e/initialization/init.Dockerfile .

e2e-setup: e2e-check-image-sha e2e-remove-resources
	@echo Finished e2e environment setup, ready to start the test

e2e-check-image-sha:
	tests/e2e/scripts/run/check_image_sha.sh

e2e-remove-resources:
	tests/e2e/scripts/run/remove_stale_resources.sh

.PHONY: test-mutation

###############################################################################
###                                Protobuf                                 ###
###############################################################################

proto-gen:
	@echo "Generating Protobuf files"
	# need buf and proto plugin
	# cd docs/devtools
	# make buf-tools
	./proto/gen.sh