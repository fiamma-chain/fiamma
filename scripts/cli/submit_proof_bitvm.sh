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
: ${FEES:=2000ufia}
: ${GAS:=80000000}
: ${PROOF_FILE:=../../prover_examples/bitvm/proof.bitvm}
: ${PUBLIC_INPUT_FILE:=../../prover_examples/bitvm/public_input.bitvm}
: ${VK_FILE:=../../prover_examples/bitvm/vk.bitvm}
: ${PROOF_SYSTEM:="Groth16Bn254_BitVM"}

NEW_PROOF=$(xxd -p -c 256 $PROOF_FILE | tr -d '\n')

NEW_PUBLIC_INPUT=$(xxd -p -c 256 $PUBLIC_INPUT_FILE | tr -d '\n')

NEW_VK=$(xxd -p -c 256 $VK_FILE | tr -d '\n')


fiammad tx zkpverify submit-proof \
  --from $ACCOUNT --chain-id $CHAIN_ID  \
  --gas $GAS --fees $FEES \
  --node $NODE \
  --keyring-backend test \
  $PROOF_SYSTEM \
  $NEW_PROOF \
	$NEW_PUBLIC_INPUT \
	$NEW_VK
