#!/bin/bash

set -e

SERVERS=("13.112.47.16" "13.113.190.239")

SSH_USER="ubuntu"


for server in "${SERVERS[@]}"; do
    echo "Deploying to ${server}"
    
    scp env_setup.sh ${SSH_USER}@${server}:~/

    ssh ${SSH_USER}@${server} "bash ~/env_setup.sh"
    
    ssh ${SSH_USER}@${server} "rm ~/env_setup.sh"
    
    echo "Environment setup completed on ${server}"
done

echo "Deployment completed on all servers."
