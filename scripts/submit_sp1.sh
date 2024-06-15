#!/bin/bash

set -e

if [ $# -ne 5 ]; then
  echo "Usage: $0 <account-name> <proof-id> <proof-file> <elf-file> <meta-data>" 
  echo "accepts 5 arg, received $#"
  exit 1
else
  ACCOUNT=$1
  PROOF_ID=$2
  PROOF_FILE=$3
  ELF_FILE=$4
  META_DATA=$5
fi

: ${CHAIN_ID:="fiamma-testnet-1"}

: ${NODE:="tcp://localhost:26657"}
: ${FEES:=2000stake}
: ${GAS:=20000000}

NEW_PROOF_FILE=$(mktemp)
base64 -i $PROOF_FILE | tr -d '\n' > $NEW_PROOF_FILE

NEW_ELF_FILE=$(mktemp)
base64 -i $ELF_FILE | tr -d '\n' > $NEW_ELF_FILE

TRANSACTION=$(mktemp)
fiammad tx zkproof submit-sp1 $PROOF_ID, "PLACEHOLDER" "PLACEHOLDER" $META_DATA \
  --from $ACCOUNT --chain-id $CHAIN_ID --generate-only \
  --gas $GAS --fees $FEES \
  | jq '.body.messages[0].proof=$proof' --rawfile proof $NEW_PROOF_FILE \
  | jq '.body.messages[0].elf=$elf' --rawfile elf $NEW_ELF_FILE \
  > $TRANSACTION

SIGNED=$(mktemp)
fiammad tx sign $TRANSACTION \
  --from $ACCOUNT --node $NODE \
  > $SIGNED

fiammad tx broadcast $SIGNED --node $NODE
