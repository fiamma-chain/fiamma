#!/usr/bin/make -f


DOCKER := $(shell which docker)
GOPATH=$(shell go env GOPATH)



###############################################################################
###                                  Build                                  ###
###############################################################################


build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/layerd-darwin-amd64 ./cmd/fiammad
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ./build/layerd-darwin-arm64 ./cmd/fiammad

build-all: build-darwin



###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=0.14.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-all: proto-gen proto-swagger-gen

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

.PHONY: proto-gen proto-swagger-gen proto-format prot-lint