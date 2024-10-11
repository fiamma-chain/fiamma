#!/bin/bash

set -e


: ${CHAIN_ID:="fiamma-testnet-1"}
: ${NODE:="http://54.65.75.57:26657"}
: ${PROOF_FILE:=../../prover_examples/bitvm_challenge/proof.bitvm}
: ${PUBLIC_INPUT_FILE:=../../prover_examples/bitvm_challenge/public_input.bitvm}
: ${VK_FILE:=../../prover_examples/bitvm_challenge/vk.bitvm}
: ${PROOF_SYSTEM:="GROTH16_BN254_BITVM"}
: ${NAMESPACE:="TEST"}

NEW_NAMESPACE=$(echo -n $NAMESPACE | xxd -p)
NEW_PROOF_SYSTEM=$(echo -n $PROOF_SYSTEM | xxd -p)

NEW_PROOF=$(xxd -p -c 256 $PROOF_FILE | tr -d '\n')

NEW_PUBLIC_INPUT=$(xxd -p -c 256 $PUBLIC_INPUT_FILE | tr -d '\n')

NEW_VK=$(xxd -p -c 256 $VK_FILE | tr -d '\n')

# Concatenate the proof, public input, and vk
allDataHex="${NEW_NAMESPACE}${NEW_PROOF_SYSTEM}${NEW_PROOF}${NEW_PUBLIC_INPUT}${NEW_VK}"

proof_id=$(echo -n "$allDataHex" | xxd -r -p | sha256sum | awk '{print $1}')

fiammad query zkpverify get-bitvm-challenge-data $proof_id --node $NODE
 