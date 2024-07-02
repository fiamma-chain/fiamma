FROM golang:1.21 AS build-env

# Customize to your build env

# TARGETPLATFORM should be one of linux/amd64 or linux/arm64.
ARG TARGETPLATFORM="linux/amd64"
# Version to build. Default is empty
ARG VERSION

ARG LEDGER_ENABLED="false"
# Cosmos build options
ARG COSMOS_BUILD_OPTIONS=""

# Install cli tools for building and final image
RUN apt-get update && apt-get install -y make git bash gcc curl jq pkg-config openssl libssl-dev 

# Install Rust
RUN curl https://sh.rustup.rs -sSf | bash -s -- -y 
    
ENV PATH="/root/.cargo/bin:${PATH}"

USER root

# Build
WORKDIR /root/fiamma

# First cache dependencies
COPY go.mod go.sum .
RUN go mod download
# Then copy everything else
COPY . .
# If version is set, then checkout this version
RUN if [ -n "${VERSION}" ]; then \
        git checkout -f ${VERSION}; \
    fi

RUN LEDGER_ENABLED=$LEDGER_ENABLED \
    BUILD_TAGS=$BUILD_TAGS \
    COSMOS_BUILD_OPTIONS=$COSMOS_BUILD_OPTIONS \
    LINK_STATICALLY=false \
    make build

FROM debian:bookworm-slim AS run

RUN apt-get update && apt-get install -y bash curl jq wget pkg-config openssl libssl-dev

# Label should match your github repo
LABEL org.opencontainers.image.source="https://github.com/fiamma-chain/fiammad:${VERSION}"

# Install libraries
# Cosmwasm - Download correct libwasmvm version
COPY --from=build-env /root/fiamma/go.mod /tmp
RUN WASMVM_VERSION=$(grep github.com/CosmWasm/wasmvm /tmp/go.mod | cut -d' ' -f2) && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm.$(uname -m).so \
        -O /lib/libwasmvm.$(uname -m).so && \
    # verify checksum
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt -O /tmp/checksums.txt && \
    sha256sum /lib/libwasmvm.$(uname -m).so | grep $(cat /tmp/checksums.txt | grep libwasmvm.$(uname -m) | cut -d ' ' -f 1)
RUN rm -f /tmp/go.mod

COPY --from=build-env /root/fiamma/build/fiammad /bin/fiammad

# COPY --from=build-env /root/fiamma/x/zkpverify/verifiers/sp1/lib/libsp1_verifier.so \
#     /root/fiamma/x/zkpverify/verifiers/sp1/lib/libsp1_verifier.so

ENTRYPOINT [ "fiammad" ]

