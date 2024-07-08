#!/usr/bin/make -f

PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
BINDIR ?= $(GOPATH)/bin
PROJECT_NAME ?= fiamma
BUILDDIR ?= $(CURDIR)/build
DOCKER := $(shell which docker)
SIMAPP = ./app
GORELEASER_CROSS_VERSION = v1.21.9
GORELEASER_VERSION = v1.21.0
HTTPS_GIT := https://github.com/fiamma-chain/fiamma.git

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

GOLANGCI_LINT := $(shell which golangci-lint)
MISSPELL := $(shell which misspell)

LIBWASM_VERSION = $(shell go list -m -f '{{ .Version }}' github.com/CosmWasm/wasmvm)

# Release environment variable
RELEASE ?= false
GORELEASER_SKIP_VALIDATE ?= false


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


ifeq (cleveldb,$(findstring cleveldb,$(FIAMMA_BUILD_OPTIONS)))
  build_tags += gcc
endif

ifeq (secp,$(findstring secp,$(FIAMMA_BUILD_OPTIONS)))
  build_tags += libsecp256k1_sdk
endif

whitespace :=
whitespace := $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=fiamma \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=fiammad \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(FIAMMA_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (badgerdb,$(findstring badgerdb,$(FIAMMA_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb
  BUILD_TAGS += badgerdb
endif
# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(FIAMMA_BUILD_OPTIONS)))
  CGO_ENABLED=1
  BUILD_TAGS += rocksdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(FIAMMA_BUILD_OPTIONS)))
  BUILD_TAGS += boltdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb
endif

ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif

ifeq (,$(findstring nostrip,$(FIAMMA_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(FIAMMA_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

# Update changelog vars
ifneq (,$(SINCE_TAG))
	since_tag := --since-tag $(SINCE_TAG)
endif
ifneq (,$(UPCOMING_TAG))
	upcoming_tag := --upcoming-tag $(UPCOMING_TAG)
endif

all: tools build lint test

###############################################################################
###                          Formatting & Linting                           ###
###############################################################################

install-golangci-lint: 
ifndef GOLANGCI_LINT
	@echo "Installing golangci-lint..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
endif

lint: install-golangci-lint 
	golangci-lint run  --timeout 5m
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s

lint-fix: install-golangci-lint
	golangci-lint run --fix --out-format=tab --issues-exit-code=0

install-misspell:
ifndef MISSPELL
	@echo "Installing misspell..."
	go install github.com/client9/misspell/cmd/misspell@latest
endif

format: install-misspell
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs goimports -w -local fiamma
.PHONY: lint lint-fix format

###############################################################################
###                                  Build                                  ###
###############################################################################

# build-sp1-ffi:
# 	@cd ./x/zkpverify/verifiers/sp1/lib \
# 		&& cargo build --release \
# 		&& cp target/release/libsp1_verifier_ffi.a ./libsp1_verifier.a 

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./cmd/fiammad

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

.PHONY: build install



###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=0.14.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-all: proto-gen proto-format proto-lint

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh

proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

proto-lint:
	@$(protoImage) buf lint --error-format=json

.PHONY: proto-gen proto-format prot-lint



###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

test:
	@echo "Running tests..."
	@go test -cover -mod=readonly ./x/...
	@echo "Completed tests!"

.PHONY: test


###############################################################################
###                                Docker                                   ###
###############################################################################

build-docker:
	docker build \
	-t fiammachain/fiammad \
	--build-arg GIT_VERSION=$(VERSION) \
	--build-arg GIT_COMMIT=$(COMMIT) \
	-f Dockerfile .

docker-rmi: 
	docker rmi fiammachain/fiammad 2>/dev/null; true

.PHONY: build-docker docker-rmi