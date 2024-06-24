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
RUN apt-get update && apt-get install -y make git bash gcc curl jq pkg-config libssl-dev 

# Install Rust
RUN curl https://sh.rustup.rs -sSf | bash -s -- -y 
    
ENV PATH="/root/.cargo/bin:${PATH}"

# Build
WORKDIR /go/src/github.com/fiamma-network/fiamma
# First cache dependencies
COPY go.mod go.sum /go/src/github.com/fiamma-network/fiamma/
RUN go mod download
# Then copy everything else
COPY ./ /go/src/github.com/fiamma-network/fiamma/
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
# Create a user
RUN addgroup --gid 1137 --system fiamma && adduser --uid 1137 --gid 1137 --system --home /home/fiamma fiamma
RUN apt-get update && apt-get install -y bash curl jq wget pkg-config libssl-dev

# Label should match your github repo
LABEL org.opencontainers.image.source="https://github.com/fiamma-network/fiammad:${VERSION}"

COPY --from=build-env /go/src/github.com/fiamma-network/fiamma/build/fiammad /bin/fiammad

# Set home directory and user
WORKDIR /home/fiamma
RUN chown -R fiamma /home/fiamma
RUN chmod g+s /home/fiamma
USER fiamma
