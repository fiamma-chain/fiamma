![](https://github.com/fiamma-chain/fiamma/blob/main/banner.png)

# Fiamma

[![Version](https://img.shields.io/github/v/tag/fiamma-chain/fiamma.svg?sort=semver&style=flat-square)](https://github.com/fiamma-chain/fiamma/releases/latest)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://pkg.go.dev/github.com/fiamma-chain/fiamma)
[![Go Report Card](https://goreportcard.com/badge/github.com/fiamma-chain/fiamma)](https://goreportcard.com/report/github.com/fiamma-chain/fiamma)
[![codecov](https://codecov.io/gh/fiamma-chain/fiamma/branch/master/graph/badge.svg)](https://codecov.io/gh/fiamma-chain/fiamma)
[![License:Apache-2.0](https://img.shields.io/github/license/fiamma-chain/fiamma.svg?style=flat-square)](https://github.com/fiamma-chain/fiamma/LICENSE)

Fiamma is a BitVM-Powered ZKP (Zero-Knowledge Proofs) Verification Layer, and additional foundational tools to support a thriving, diverse Bitcoin ecosystem.

Our infrastructures empower BTC holders to earn trustless yields and explore limitless ecosystems like never before. We’re unlocking native, trustless access to Bitcoin’s immense liquidity for all ecosystems, creating unprecedented opportunities. Developers can now innovate freely on Bitcoin beyond EVM or SVM, shaping a future where Bitcoin serves as the foundation for boundless potential.


## System Requirements

The following specifications have been found to work well:

- Quad Core or larger AMD or Intel (amd64) CPU
- 32GB RAM;
- 1TB NVMe SSD Storage (disk i/o is crucial);
- 100Mbps bi-directional Internet connection;

## Software Dependencies

The following software should be installed on the target system:

- The Go Programming Language (<https://go.dev>)
- Git Distributed Version Control (<https://git-scm.com>)
- Docker (<https://www.docker.com>)
- GNU Make (<https://www.gnu.org/software/make>)
- Openssl <https://www.openssl.org/>
- jq (https://jqlang.github.io/jq/)


## Build from Source

[Clone the repository](https://github.com/fiamma-chain/fiamma), checkout the `<release version>` branch and install:

```sh
cd fiamma
git checkout <release version>
make install
```

This will install the `fiammad` binary to your `GOPATH`.

## Dockerized Containers

A docker image for production purposes (no shell access):

[Packages: fiammad](https://github.com/fiamma-chain/fiamma/pkgs/container/fiamma)


## Joining the testnet

Please follow the instructions on the [Developer Guides](https://docs.fiammachain.io/developer-guides/set-up-a-node).


## Documentation

To learn more, please [visit the official fiamma documentation](https://docs.fiammachain.io).
