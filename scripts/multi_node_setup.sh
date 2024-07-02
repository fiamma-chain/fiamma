#!/bin/bash

set -e

: "${CHAIN_ID:=fiamma-testnet-1}"
: "${PASSWORD:=password}"
token="ufia"
initial_balance=10000000000000
initial_faucet_balance=10000000000
initial_stake=10000000
minimum_gas_price=0

# nubit da rpc url and authkey

rpc="http://172.31.46.84:26658"
authkey="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJwdWJsaWMiLCJyZWFkIiwid3JpdGUiLCJhZG1pbiJdfQ.z8uQPBAWnOTKS8C1BrT29O0it38o66sXSWHhyfGejKk"


if [ $# -lt 1 ]; then
    echo "Usage: $0 <node1> [<node2> ...]"
    exit 1
fi

echo "Creating directories for nodes..."
rm -rf testnet-nodes
for node in "$@"; do
    mkdir -p testnet-nodes/$node
done

node_ids=()

for node in "$@"; do
    echo "Initializing $node..."
    docker run --rm -v $(pwd)/testnet-nodes/$node:/root/.fiamma -it fiammachain/fiammad init fiamma_$node --chain-id $CHAIN_ID  > /dev/null
    
    docker run --rm -it -v $(pwd)/testnet-nodes/$node:/root/.fiamma --entrypoint sed fiammachain/fiammad -i 's/"stake"/"'$token'"/g' /root/.fiamma/config/genesis.json 
    docker run --rm -v $(pwd)/testnet-nodes/$node:/root/.fiamma -it fiammachain/fiammad config set app minimum-gas-prices "$minimum_gas_price$token"
    docker run --rm -v $(pwd)/testnet-nodes/$node:/root/.fiamma -it fiammachain/fiammad config set app pruning "nothing" 

    docker run --rm -v $(pwd)/testnet-nodes/$node:/root/.fiamma -it fiammachain/fiammad config set app da-config.rpc "$rpc"
    docker run --rm -v $(pwd)/testnet-nodes/$node:/root/.fiamma -it fiammachain/fiammad config set app da-config.authkey "$authkey"

    node_id=$(docker run --rm -i -v $(pwd)/testnet-nodes/$node:/root/.fiamma fiammachain/fiammad tendermint show-node-id)
    node_ids+=($node_id)

    echo "Node ID for $node: $node_id"
done


for (( i=1; i <= "$#"; i++ )); do
    echo "Creating key for ${!i} user..."
    printf "$PASSWORD\n$PASSWORD\n" | docker run --rm -i -v $(pwd)/testnet-nodes/${!i}:/root/.fiamma fiammachain/fiammad keys --keyring-backend file --keyring-dir /root/.fiamma/keys add val_${!i} > /dev/null 2> $(pwd)/testnet-nodes/${!i}/mnemonic.txt

    val_address=$(echo $PASSWORD | docker run --rm -i -v $(pwd)/testnet-nodes/${!i}:/root/.fiamma fiammachain/fiammad keys --keyring-backend file --keyring-dir /root/.fiamma/keys show val_${!i} --address)
    echo "val_${!i} address: $val_address"
    echo "val_${!i} mnemonic: $(cat $(pwd)/testnet-nodes/${!i}/mnemonic.txt)"

    echo "Giving val_${!i} some tokens..."
    if [ $i -eq 1 ]; then
        faucet_initial_balance=$((initial_faucet_balance + initial_stake))
        docker run --rm -it -v $(pwd)/testnet-nodes/${!i}:/root/.fiamma fiammachain/fiammad genesis add-genesis-account $val_address $faucet_initial_balance$token
    else
        docker run --rm -it -v $(pwd)/testnet-nodes/${!i}:/root/.fiamma fiammachain/fiammad genesis add-genesis-account $val_address $initial_balance$token
    fi

    if [ $((i+1)) -le "$#" ]; then
        j=$((i+1))
        cp $(pwd)/testnet-nodes/${!i}/config/genesis.json $(pwd)/testnet-nodes/${!j}/config/genesis.json
    elif [ $# != 1 ] && [ $((i+1)) -gt $# ]; then
        cp $(pwd)/testnet-nodes/${!i}/config/genesis.json $(pwd)/testnet-nodes/$1/config/genesis.json
    fi
done



for (( i=1; i <= "$#"; i++ )); do
    echo "Giving val_${!i} some stake..."
    echo $PASSWORD | docker run --rm -i -v $(pwd)/testnet-nodes/${!i}:/root/.fiamma fiammachain/fiammad genesis gentx val_${!i} $initial_stake$token --keyring-backend file --keyring-dir /root/.fiamma/keys --account-number 0 --sequence 0 --chain-id $CHAIN_ID --gas 1000000 --gas-prices $minimum_gas_price$token

    if [ $i -gt 1 ]; then
        cp $(pwd)/testnet-nodes/${!i}/config/gentx/* $(pwd)/testnet-nodes/$1/config/gentx/
    fi

    if [ $((i+1)) -le "$#" ]; then
        j=$((i+1))
        cp $(pwd)/testnet-nodes/${!i}/config/genesis.json $(pwd)/testnet-nodes/${!j}/config/genesis.json
    elif [ $# != 1 ] && [ $((i+1)) -gt $# ]; then
        cp $(pwd)/testnet-nodes/${!i}/config/genesis.json $(pwd)/testnet-nodes/$1/config/genesis.json
    fi
done

echo "Collecting genesis transactions..."
docker run --rm -it -v $(pwd)/testnet-nodes/$1:/root/.fiamma fiammachain/fiammad genesis collect-gentxs > /dev/null

if ! docker run --rm -it -v $(pwd)/testnet-nodes/$1:/root/.fiamma fiammachain/fiammad genesis validate-genesis; then
    echo "Invalid genesis"
    exit 1
fi

echo "Copying genesis file to other nodes..."
for node in "${@:2}"; do
    cp $(pwd)/testnet-nodes/$1/config/genesis.json $(pwd)/testnet-nodes/$node/config/genesis.json
done

echo "Setting node addresses in config..."
for (( i=1; i <= "$#"; i++ )); do
    other_addresses=()
    for (( j=1; j <= "$#"; j++ )); do
        if [ $j -ne $i ]; then
            other_addresses+=("${node_ids[$j - 1]}@${!j}:26656")
        fi
    done
    other_addresses=$(IFS=,; echo "${other_addresses[*]}")
    #Peer configuration
    docker run --rm -v $(pwd)/testnet-nodes/${!i}:/root/.fiamma -it fiammachain/fiammad config set config p2p.persistent_peers "$other_addresses" --skip-validate
    #RPC configuration
    docker run --rm -v $(pwd)/testnet-nodes/${!i}:/root/.fiamma -it fiammachain/fiammad config set config rpc.laddr "tcp://0.0.0.0:26657" --skip-validate
    #Explorer configuration
    docker run --rm -v $(pwd)/testnet-nodes/${!i}:/root/.fiamma -it fiammachain/fiammad config set config rpc.cors_allowed_origins '["*"]' --skip-validate 
    docker run --rm -v $(pwd)/testnet-nodes/${!i}:/root/.fiamma -it fiammachain/fiammad config set app api.enable true --skip-validate 
    docker run --rm -v $(pwd)/testnet-nodes/${!i}:/root/.fiamma -it fiammachain/fiammad config set app api.enabled-unsafe-cors true --skip-validate 
    docker run --rm -v $(pwd)/testnet-nodes/${!i}:/root/.fiamma -it fiammachain/fiammad config set app api.address "tcp://0.0.0.0:1317" --skip-validate
done


echo "Setting up docker compose..."
rm -f $(pwd)/testnet-nodes/docker-compose.yml
printf "version: '3.7'\nnetworks:\n  net-public:\nservices:\n" > $(pwd)/testnet-nodes/docker-compose.yml
for node in "$@"; do
    printf "  fiammad-$node:\n    command: start\n    image: fiammachain/fiammad\n    container_name: $node\n    volumes:\n      - ./$node:/root/.fiamma\n    networks:\n      - net-public\n" >> $(pwd)/testnet-nodes/docker-compose.yml
    if [ $node == "$1" ]; then
        printf "    ports:\n      - 0.0.0.0:26657:26657\n" >> $(pwd)/testnet-nodes/docker-compose.yml
    fi
    printf "\n" >> $(pwd)/testnet-nodes/docker-compose.yml
done
