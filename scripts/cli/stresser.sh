#!/bin/bash

#
# This script sends dummy transactions from an <account> with a test keyring. It should be run from the repository root.
#

if [ $# -ne 1 ]; then
  echo "Usage: $0 <account>"
  echo "accepts 1 arg, received $#"
  exit 1
else
  ACCOUNT=$1
fi

: ${CHAIN_ID:="fiamma-testnet-1"}
: ${NODE_RPC:="https://testnet-rpc.fiammachain.io"}
: ${NODE_API:="https://testnet-api.fiammachain.io"}

# New elements can be added to the array to send more transactions
PROOFS=(

)

ADDRESS=$(
  fiammad keys show $ACCOUNT \
    --keyring-backend test --output json \
    | jq .address | tr -d \"
)

ACCOUNT_INFO= $(curl -s "$NODE_API/cosmos/auth/v1beta1/accounts/$ADDRESS" -o account_info.json)

ACCOUNT=$(jq -r '.account.address' account_info.json)
ACCOUNT_NUMBER=$(jq -r '.account.account_number' account_info.json)
SEQUENCE=$(jq -r '.account.sequence' account_info.json)

rm account_info.json
echo "Account: $ACCOUNT"
echo "Account Number: $ACCOUNT_NUMBER"
echo "Sequence: $SEQUENCE"

for (( i=0; i<10000; i++ ))
do
   fiammad tx zkpverify submit-proof \
    --keyring-backend test --from $ACCOUNT \
    --chain-id $CHAIN_ID \
    --fees 20ufia \
    --offline \
    --gas 20000000 \
    --sequence $SEQUENCE \
    --account-number $ACCOUNT_NUMBER \
    --node $NODE_RPC \
    --yes \
    PLONK_BN254 \
    $(cat ../../prover_examples/gnark_plonk/example/proof) \
	  $(cat ../../prover_examples/gnark_plonk/example/public_input) \
	  $(cat ../../prover_examples/gnark_plonk/example/vk)
  let SEQUENCE++
done
