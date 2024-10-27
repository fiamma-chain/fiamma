#!/bin/bash

set -e

if [ $# -ne 1 ]; then
  echo "Usage: $0 <account> " 
  echo "accepts 1 arg, received $#"
  exit 1
else
  ACCOUNT=$1
fi

: ${CHAIN_ID:="fiamma-testnet-1"}
: ${NODE:="http://127.0.0.1:26657"}

: ${GAS:=80000000}
: ${VK_FILE:=../../prover_examples/bitvm/vk.bitvm}
: ${VK_FILE_CHALLENGE:=../../prover_examples/bitvm_challenge/vk.bitvm}

fiammad tx bitvmstaker register-vk \
  --from $ACCOUNT --chain-id $CHAIN_ID  \
  --gas $GAS  \
  --node $NODE \
  $VK_FILE

sleep 2

fiammad tx bitvmstaker register-vk \
  --from $ACCOUNT --chain-id $CHAIN_ID  \
  --gas $GAS  \
  --node $NODE \
  $VK_FILE_CHALLENGE
