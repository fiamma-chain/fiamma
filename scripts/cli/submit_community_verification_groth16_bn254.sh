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
: ${PROOF_FILE:=../../prover_examples/gnark_groth16/example/proof}
: ${PUBLIC_INPUT_FILE:=../../prover_examples/gnark_groth16/example/public_input}
: ${VK_FILE:=../../prover_examples/gnark_groth16/example/vk}
: ${PROOF_SYSTEM:="Groth16Bn254"}
: ${FEES:=2000ufia}
: ${GAS:=20000000}


NEW_PROOF_SYSTEM=$(echo -n $PROOF_SYSTEM | xxd -p)

NEW_PROOF=$(cat $PROOF_FILE)

NEW_PUBLIC_INPUT=$(cat $PUBLIC_INPUT_FILE)

NEW_VK=$(cat $VK_FILE)

# Concatenate the proof, public input, and vk
allDataHex="${NEW_PROOF_SYSTEM}${NEW_PROOF}${NEW_PUBLIC_INPUT}${NEW_VK}"

proof_id=$(echo -n "$allDataHex" | xxd -r -p | sha256sum | awk '{print $1}')

fiammad tx zkpverify submit-community-verification \
  --from $ACCOUNT --chain-id $CHAIN_ID  \
  --gas $GAS --fees $FEES \
  --node $NODE \
  --keyring-backend test \
  $proof_id true

