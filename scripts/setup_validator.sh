#!/bin/bash
set -e

if [ $# -lt 2 ]; then
	echo "Usage: $0 <account> <staking_amount>"
	exit 1
else
	VALIDATOR=$1
	STAKING_AMOUNT=$2
fi

NODE_HOME=$HOME/.fiamma
CHAIN_BINARY=fiammad
: ${CHAIN_ID:="fiamma-testnet-1"}

: ${PEER_ADDR:="18.182.20.173"}

VALIDATOR_KEY=$($CHAIN_BINARY tendermint show-validator)
MONIKER=$($CHAIN_BINARY config get config moniker)

cat << EOF > $NODE_HOME/config/validator.json
{
	"pubkey": $VALIDATOR_KEY,
	"amount": "$STAKING_AMOUNT",
	"moniker": $MONIKER,
	"commission-rate": "0.1",
	"commission-max-rate": "0.2",
	"commission-max-change-rate": "0.01",
	"min-self-delegation": "1"
}
EOF

# You should ensure that the fiammad binary files have been correctly installed.
$CHAIN_BINARY tx staking create-validator $NODE_HOME/config/validator.json \
	--from $VALIDATOR --keyring-backend test --chain-id $CHAIN_ID \
	--fees 2000ufia \
	--node tcp://$(echo $PEER_ADDR | cut -d, -f1):26657 
