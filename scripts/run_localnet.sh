#!/bin/bash

set -e

# Basic settings
CHAIN_ID="fiamma-testnet-1"
PASSWORD="password"
token="ufia"
initial_balance="10000000000000"
initial_faucet_balance="10000000000"
initial_stake="10000000"
minimum_gas_price="0"
committee_address=""
staker_addresses=()

# RPC and Auth Key settings
rpc="http://127.0.0.1:26658"
authkey=""

# Ensure that node name is provided
if [ $# -lt 1 ]; then
    echo "Usage: $0 <node_name>"
    exit 1
fi

node=$1

# Default data directory for fiammad
DATA_DIR="$HOME/.fiamma"

# Check if the data directory already exists
if [ -d "$DATA_DIR" ]; then
    echo "An existing data directory was found at '$DATA_DIR'."
    # Prompt user for confirmation
    read -p "Do you want to remove this directory and all of its contents? (y/n) " user_confirm

    # Check user input
    case $user_confirm in
        [Yy]* ) 
            echo "Removing existing data directory..."
            rm -rf "$DATA_DIR"
            echo "Directory removed."
            ;;
        [Nn]* )
            echo "Operation aborted by the user."
            exit 1
            ;;
        * ) 
            echo "Invalid input. Please answer 'yes' or 'no'."
            exit 1
            ;;
    esac
fi


echo "Initializing $node..."
fiammad init $node --chain-id $CHAIN_ID > /dev/null

# Configuration adjustments
perl -pi -e 's/"stake"/"'$token'"/g' "$DATA_DIR/config/genesis.json"
fiammad config set app minimum-gas-prices "$minimum_gas_price$token"
fiammad config set app pruning "nothing"
fiammad config set app da-config.rpc "$rpc"
fiammad config set app da-config.authkey "$authkey"

# Key generation and setup
echo "Creating key for $node..."
fiammad keys add $node --keyring-backend test > $DATA_DIR/mnemonic.txt 2>&1
val_address=$(fiammad keys --keyring-backend test show $node --address)
echo "Validator address: $val_address"
echo "Validator mnemonic: $(cat mnemonic.txt)"


# Set committee_address to the created account address
committee_address=$val_address
jq '.app_state.bitvmstaker.committee_address = "'$committee_address'"' ~/.fiamma/config/genesis.json > ~/.fiamma/config/genesis.json.tmp
mv ~/.fiamma/config/genesis.json.tmp ~/.fiamma/config/genesis.json


# Set staker addresses
val_operator=$(fiammad keys show $node --keyring-backend test -a --bech val)
staker_addresses+=($val_operator)
echo "Adding initial staker addresses to genesis..."
jq --argjson staker_addresses "$(printf '%s\n' "${staker_addresses[@]}" | jq -R . | jq -s .)" '.app_state.bitvmstaker.staker_addresses = $staker_addresses' ~/.fiamma/config/genesis.json > ~/.fiamma/config/genesis.json.tmp
mv ~/.fiamma/config/genesis.json.tmp ~/.fiamma/config/genesis.json

# Setup genesis
faucet_initial_balance=$((initial_faucet_balance + initial_stake))
fiammad genesis add-genesis-account $val_address $faucet_initial_balance$token

# Configure staking and genesis
fiammad genesis gentx $node $initial_stake$token --keyring-backend test --chain-id $CHAIN_ID --gas 1000000 --gas-prices $minimum_gas_price$token

fiammad genesis collect-gentxs
if ! fiammad genesis validate-genesis; then
    echo "Invalid genesis"
    exit 1
fi


#RPC configuration
fiammad config set config rpc.laddr "tcp://0.0.0.0:26657" --skip-validate
#Explorer configuration
fiammad config set config rpc.cors_allowed_origins '["*"]' --skip-validate 
fiammad config set app api.enable true --skip-validate 
fiammad config set app api.enabled-unsafe-cors true --skip-validate 
fiammad config set app api.address "tcp://0.0.0.0:1317" --skip-validate
fiammad config set app grpc.address "0.0.0.0:9090" --skip-validate

echo "Node $node is set up and ready to start..."

fiammad start