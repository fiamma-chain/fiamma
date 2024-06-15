#!/bin/bash

set -e

if [ $# -ne 5 ]; then
  echo "Usage: $0 <account-name> <proof-id> <result>" 
  echo "accepts 3 arg, received $#"
  exit 1
else
  ACCOUNT=$1
  PROOF_ID=$2
  RESULT=$3
  
fi

: ${CHAIN_ID:="fiamma-testnet-1"}

: ${NODE:="tcp://localhost:26657"}
: ${FEES:=2000stake}
: ${GAS:=20000000}

TRANSACTION=$(mktemp)
fiammad tx zkproof verify-proof $PROOF_ID, $RESULT \
  --from $ACCOUNT --chain-id $CHAIN_ID --generate-only \
  --gas $GAS --fees $FEES \
  > $TRANSACTION

SIGNED=$(mktemp)
fiammad tx sign $TRANSACTION \
  --from $ACCOUNT --node $NODE \
  > $SIGNED

fiammad tx broadcast $SIGNED --node $NODE
