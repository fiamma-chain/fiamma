#!/bin/bash
set -e

if [ $# -lt 1 ]; then
	echo "Usage: $0 <moniker>"
	exit 1
else
    MONIKER=$1
fi

NODE_HOME=$HOME/.fiamma
CHAIN_BINARY=fiammad
: ${CHAIN_ID:="fiamma-testnet-1"}
: ${MINIMUM_GAS_PRICES="0.0001fiamma"}

: ${PEER_ADDR="35.73.202.182,18.182.20.173", "35.74.243.172", "18.179.17.155"}

PEER_ARRAY=(${PEER_ADDR//,/ })

# You should ensure that the fiammad binary files have been correctly installed.
$CHAIN_BINARY comet unsafe-reset-all
$CHAIN_BINARY init $MONIKER \
    --chain-id $CHAIN_ID --overwrite

for ADDR in "${PEER_ARRAY[@]}"; do
    GENESIS=$(curl -f "$ADDR:26657/genesis" | jq '.result.genesis')
    if [ -n "$GENESIS" ]; then
        echo "$GENESIS" > $NODE_HOME/config/genesis.json;
        break;
    fi
done

PERSISTENT_PEERS=()

for ADDR in "${PEER_ARRAY[@]}"; do
    PEER_ID=$(curl -s "$ADDR:26657/status" | jq -r '.result.node_info.id')
    if [ -n "$PEER_ID" ]; then
        PERSISTENT_PEERS+=("$PEER_ID@$ADDR:26656")
    fi
done

CONFIG_STRING=$(IFS=,; echo "${PERSISTENT_PEERS[*]}")

$CHAIN_BINARY config set config p2p.persistent_peers "$CONFIG_STRING" --skip-validate

$CHAIN_BINARY config set app minimum-gas-prices "$MINIMUM_GAS_PRICES" --skip-validate