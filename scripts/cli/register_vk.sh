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
: ${NODE:="https://testnet-rpc.fiammachain.io"}

: ${GAS:=70000}
: ${FEE:=140ufia}
: ${VK_FILE:=../../prover_examples/bitvm/vk.bitvm}
: ${VK_FILE_CHALLENGE:=../../prover_examples/bitvm_challenge/vk.bitvm}

fiammad tx bitvmstaker register-vk \
  --from $ACCOUNT --keyring-backend test --chain-id $CHAIN_ID  \
  --gas $GAS  \
  --fees $FEE \
  --node $NODE \
  $VK_FILE

sleep 2

fiammad tx bitvmstaker register-vk \
  --from $ACCOUNT --keyring-backend test --chain-id $CHAIN_ID  \
  --gas $GAS  \
  --fees $FEE \
  --node $NODE \
  $VK_FILE_CHALLENGE
