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
: ${FEES:=2000ufia}
: ${GAS:=80000000}
: ${PROOF_FILE:=../../prover_examples/bitvm_challenge/proof.bitvm}
: ${PUBLIC_INPUT_FILE:=../../prover_examples/bitvm_challenge/public_input.bitvm}
: ${VK_FILE:=../../prover_examples/bitvm_challenge/vk.bitvm}
: ${PROOF_SYSTEM:="GROTH16_BN254_BITVM"}
: ${NAMESPACE:="TEST"}

NEW_NAMESPACE=$(echo -n $NAMESPACE | xxd -p)
# Concatenate the proof, public input, and vk
allDataHex="${NEW_NAMESPACE}${NEW_PROOF_SYSTEM}${NEW_PROOF}${NEW_PUBLIC_INPUT}${NEW_VK}"

proof_id=$(echo -n "$allDataHex" | xxd -r -p | sha256sum | awk '{print $1}')


fiammad tx zkpverify submit-proof \
  --from $ACCOUNT --chain-id $CHAIN_ID  \
  --gas $GAS --fees $FEES \
  --node $NODE \
  --keyring-backend test \
  $NAMESPACE \
  $PROOF_SYSTEM \
  $PROOF_FILE \
	$PUBLIC_INPUT_FILE \
	$VK_FILE
