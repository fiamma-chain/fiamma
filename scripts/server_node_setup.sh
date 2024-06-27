#!/bin/bash

set -e

if [ $# -lt 2 ]; then
    echo "Usage: $0 [prod|test] binary_release_tag"
    exit 1
fi

if [ "$1" = "prod" ]; then
    nodes=("node0" "node1")
    nodes_ips=("13.112.47.16" "13.113.190.239")
    servers=("ubuntu@13.112.47.16" "ubuntu@13.113.190.239")

    read -p "Are you sure you want to deploy in production? (y/n): " answer
    if [ "$answer" != "y" ]; then
        exit 0
    fi
elif [ "$1" = "test" ]; then
    nodes=("node0" "node1")
    nodes_ips=("13.112.47.16" "13.113.190.239")
    servers=("ubuntu@13.112.47.16" "ubuntu@13.113.190.239")
else
    echo "Usage: $0 [prod|test] binary_release_tag"
    exit 1
fi

rm -rf server-setup

# echo "Downloading source code into servers..."
# for server in "${servers[@]}"; do
#     ssh $server "rm -rf /home/ubuntu/fiamma"
#     ssh $server "git clone https://github.com/fiamma-network/fiamma.git /home/ubuntu/fiamma"
#     ssh $server "cd /home/ubuntu/fiamma && git checkout $2 && source /home/ubuntu/.profile && make install"
#     echo "Source code downloaded into $server successfully"
# done

mkdir -p server-setup
cd server-setup


echo "Calling setup script..."
bash ../multi_node_setup.sh "${nodes[@]}"

echo "Setting node addresses in config..."
for i in "${!nodes[@]}"; do 
    echo $(pwd)
    seeds=$(docker run -v "$(pwd)/testnet-nodes/${nodes[$i]}:/root/.fiamma" -it fiammachain/fiammad  config get config p2p.persistent_peers)
    for j in "${!nodes[@]}"; do  
        seeds=${seeds//${nodes[$j]}/${nodes_ips[$j]}}
    done
    
    docker run -v "$(pwd)/testnet-nodes/${nodes[$i]}:/root/.fiamma" -it fiammachain/fiammad  config set config p2p.persistent_peers $seeds --skip-validate    
done

echo "Sending directories to servers..."
for i in "${!servers[@]}"; do  
    ssh ${servers[$i]} "rm -rf /home/ubuntu/.fiamma"
    scp -r "testnet-nodes/${nodes[$i]}" "${servers[$i]}:/home/ubuntu/.fiamma"
done


cd ..

