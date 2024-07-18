#!/bin/bash

set -e

if [ $# -lt 2 ]; then
    echo "Usage: $0 [prod|test] binary_release_tag"
    exit 1
fi

if [ "$1" = "prod" ]; then
    nodes=("node1" "node2" "node3" "node4")
    nodes_ips=("172.31.31.70" "172.31.27.65" "172.31.17.72" "172.31.26.39") 
    servers=("ubuntu@18.182.20.173" "ubuntu@35.73.202.182" "ubuntu@35.74.243.172" "ubuntu@18.179.17.155")
    read -p "Are you sure you want to deploy in production? (y/n): " answer
    if [ "$answer" != "y" ]; then
        exit 0
    fi
elif [ "$1" = "test" ]; then
    nodes=("node1" "node2" "node3" "node4")
    nodes_ips=("172.31.31.70" "172.31.27.65" "172.31.17.72" "172.31.26.39") 
    servers=("ubuntu@18.182.20.173" "ubuntu@35.73.202.182" "ubuntu@35.74.243.172" "ubuntu@18.179.17.155")
else
    echo "Usage: $0 [prod|test] binary_release_tag"
    exit 1
fi


echo "Downloading source code into servers..."
for server in "${servers[@]}"; do
    ssh $server "rm -rf /home/ubuntu/fiamma"
    ssh $server "git clone https://github.com/fiamma-chain/fiamma.git /home/ubuntu/fiamma"
    ssh $server "cd /home/ubuntu/fiamma && git checkout $2 && source /home/ubuntu/.profile && make install"
    echo "Source code downloaded into $server successfully"
done


echo "Sending directories to servers..."
for i in "${!servers[@]}"; do  

    ## Config Cosmovisor for chain upgrade
    ssh ${servers[$i]} "cp /home/ubuntu/go/bin/fiammad /home/ubuntu/.fiamma/cosmovisor/genesis/bin/fiammad"
done


cd ..


echo "Starting new fiamma services..."
for server in "${servers[@]}"; do
    ssh $server "sudo systemctl start fiamma"
    echo "Started fiamma service on $server"
done