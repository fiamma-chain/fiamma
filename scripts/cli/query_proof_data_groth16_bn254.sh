#!/bin/bash

set -e


: ${CHAIN_ID:="fiamma-testnet-1"}
: ${NODE:="https://testnet-rpc.fiammachain.io"}
: ${PROOF_FILE:=../../prover_examples/gnark_groth16/example/proof}
: ${PUBLIC_INPUT_FILE:=../../prover_examples/gnark_groth16/example/public_input}
: ${VK_FILE:=../../prover_examples/gnark_groth16/example/vk}
: ${PROOF_SYSTEM:="GROTH16_BN254"}


NEW_PROOF_SYSTEM=$(echo -n $PROOF_SYSTEM | xxd -p)

NEW_PROOF=$(cat $PROOF_FILE)

NEW_PUBLIC_INPUT=$(cat $PUBLIC_INPUT_FILE)

NEW_VK=$(cat $VK_FILE)

# Concatenate the proof, public input, and vk
allDataHex="${NEW_PROOF_SYSTEM}${NEW_PROOF}${NEW_PUBLIC_INPUT}${NEW_VK}"

proof_id=$(echo -n "$allDataHex" | xxd -r -p | sha256sum | awk '{print $1}')

fiammad query zkpverify get-proof-data $proof_id --node $NODE
