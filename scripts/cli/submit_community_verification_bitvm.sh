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
: ${NODE:="http://54.65.137.66:26657"}
: ${PROOF_FILE:=../../prover_examples/bitvm/proof.bitvm}
: ${PUBLIC_INPUT_FILE:=../../prover_examples/bitvm/public_input.bitvm}
: ${VK_FILE:=../../prover_examples/bitvm/vk.bitvm}
: ${NAMESPACE:="TEST"}
: ${PROOF_SYSTEM:="GROTH16_BN254_BITVM"}
: ${FEES:=2000ufia}
: ${GAS:=80000000}


NEW_NAMESPACE=$(echo -n $NAMESPACE | xxd -p)
NEW_PROOF_SYSTEM=$(echo -n $PROOF_SYSTEM | xxd -p)

NEW_PROOF=$(xxd -p -c 256 $PROOF_FILE | tr -d '\n')

NEW_PUBLIC_INPUT=$(xxd -p -c 256 $PUBLIC_INPUT_FILE | tr -d '\n')

NEW_VK=$(xxd -p -c 256 $VK_FILE | tr -d '\n')

# Concatenate the proof, public input, and vk
allDataHex="${NEW_NAMESPACE}${NEW_PROOF_SYSTEM}${NEW_PROOF}${NEW_PUBLIC_INPUT}${NEW_VK}"

proof_id=$(echo -n "$allDataHex" | xxd -r -p | sha256sum | awk '{print $1}')


fiammad tx zkpverify submit-community-verification \
  --from $ACCOUNT --chain-id $CHAIN_ID  \
  --gas $GAS --fees $FEES \
  --node $NODE \
  --keyring-backend test \
  $proof_id true
