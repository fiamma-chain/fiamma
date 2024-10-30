#!/bin/bash

set -e


: ${CHAIN_ID:="fiamma-testnet-1"}
: ${NODE:="http://127.0.0.1:26657"}

fiammad query zkpverify get-da-submission-queue --node $NODE
