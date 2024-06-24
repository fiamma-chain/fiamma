#!/usr/bin/make -f


DOCKER := $(shell which docker)
GOPATH=$(shell go env GOPATH)
BUILDDIR ?= $(CURDIR)/build
GOLANGCI_LINT := $(shell which golangci-lint)
MISSPELL := $(shell which misspell)


###############################################################################
###                          Formatting & Linting                           ###
###############################################################################

containerMarkdownLintImage=tmknom/markdownlint
containerMarkdownLint=cosmos-sdk-markdownlint
containerMarkdownLintFix=cosmos-sdk-markdownlint-fix

golangci_lint_cmd=golangci-lint

install-golangci-lint: 
ifndef GOLANGCI_LINT
	@echo "Installing golangci-lint..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
endif

lint: lint-go
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerMarkdownLint}$$"; then docker start -a $(containerMarkdownLint); else docker run --name $(containerMarkdownLint) -i -v "$(CURDIR):/work" $(markdownLintImage); fi

lint-fix: install-golangci-lint
	$(golangci_lint_cmd) run --fix --out-format=tab --issues-exit-code=0
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerMarkdownLintFix}$$"; then docker start -a $(containerMarkdownLintFix); else docker run --name $(containerMarkdownLintFix) -i -v "$(CURDIR):/work" $(markdownLintImage) . --fix; fi

lint-go: install-golangci-lint
	echo $(GIT_DIFF)
	$(golangci_lint_cmd) run --out-format=tab $(GIT_DIFF)

.PHONY: lint lint-fix

install-misspell:
ifndef MISSPELL
	@echo "Installing misspell..."
	go install github.com/client9/misspell/cmd/misspell@latest
endif

format: install-misspell
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs goimports -w -local fiamma
.PHONY: format

###############################################################################
###                                  Build                                  ###
###############################################################################


BUILD_TARGETS := build install

build-sp1-ffi-macos:
	@cd ./x/zkpverify/verifiers/sp1/lib \
		&& cargo build --release \
		&& cp target/release/libsp1_verifier_ffi.dylib ./libsp1_verifier.dylib \
		&& cp target/release/libsp1_verifier_ffi.a ./libsp1_verifier.a

build-sp1-ffi-linux:
	@cd ./x/zkpverify/verifiers/sp1/lib \
		&& cargo build --release \
		&& cp target/release/libsp1_verifier_ffi.so ./libsp1_verifier.so \
		&& cp target/release/libsp1_verifier_ffi.a ./libsp1_verifier.a

ifeq ($(shell uname), Darwin)
    build-sp1-ffi: build-sp1-ffi-macos
else
    build-sp1-ffi: build-sp1-ffi-linux
endif

build: build-sp1-ffi go-build

install: build-sp1-ffi go-install

go-build: $(BUILDDIR)/ go.sum
	go build -mod=readonly -o $(BUILDDIR)/ ./cmd/fiammad

go-install: $(BUILDDIR)/ go.sum
	go install -mod=readonly ./cmd/fiammad

build-linux: build-sp1-ffi-linux go-build-linux

go-build-linux: build-sp1-ffi-linux go-build-linux
	GOOS=linux GOARCH=$(if $(findstring aarch64,$(shell uname -m)) || $(findstring arm64,$(shell uname -m)),arm64,amd64) LEDGER_ENABLED=false $(MAKE) build

# $(BUILD_TARGETS): go.sum $(BUILDDIR)/
# 	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./cmd/fiammad

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

.PHONY: build build-linux


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
	$(MAKE) -C contrib/images fiammad

.PHONY: build-docker