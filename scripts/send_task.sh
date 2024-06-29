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
: ${FEES:=2000fiamma}
: ${GAS:=20000000}

fiammad tx zkpverify send-task \
  --from $ACCOUNT --chain-id $CHAIN_ID  \
  --gas $GAS --fees $FEES \
  --node $NODE \
  "Groth16Bn254" \
  $(cat ../prover_examples/gnark_groth16/example/proof) \
	$(cat ../prover_examples/gnark_groth16/example/public_input) \
	$(cat ../prover_examples/gnark_groth16/example/vk)
