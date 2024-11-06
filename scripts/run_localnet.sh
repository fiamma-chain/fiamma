#!/bin/bash

set -e

# Basic settings
CHAIN_ID="fiamma-testnet-1"
token="ufia"
initial_balance="1000000000000000"
initial_stake="50000000000000"
minimum_gas_price="0.002"
committee_address=""
staker_addresses=()

babylonContractAddr=fiamma14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sgx3jav
btcStakingContractAddr=fiamma1nc5tatafv6eyq7llkr2gv50ff9e22mnf70qgjlv737ktmt4eswrqyn5sl2

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
sed -i '' 's/"babylon_contract_address": ""/"babylon_contract_address": "'"$babylonContractAddr"'"/g' "$DATA_DIR/config/genesis.json"
sed -i '' 's/"btc_staking_contract_address": ""/"btc_staking_contract_address": "'"$btcStakingContractAddr"'"/g' "$DATA_DIR/config/genesis.json"

fiammad config set app minimum-gas-prices "$minimum_gas_price$token"
fiammad config set app pruning "nothing"

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

echo "Setting zkpverify da_submitter in genesis..."
jq '.app_state.zkpverify.da_submitter = "'$committee_address'"' ~/.fiamma/config/genesis.json > ~/.fiamma/config/genesis.json.tmp
mv ~/.fiamma/config/genesis.json.tmp ~/.fiamma/config/genesis.json

# Set staker addresses
val_operator=$(fiammad keys show $node --keyring-backend test -a --bech val)
staker_addresses+=($val_operator)
echo "Adding initial staker addresses to genesis..."
jq --argjson staker_addresses "$(printf '%s\n' "${staker_addresses[@]}" | jq -R . | jq -s .)" '.app_state.bitvmstaker.staker_addresses = $staker_addresses' ~/.fiamma/config/genesis.json > ~/.fiamma/config/genesis.json.tmp
mv ~/.fiamma/config/genesis.json.tmp ~/.fiamma/config/genesis.json

# Setup genesis
fiammad genesis add-genesis-account $val_address $initial_balance$token

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