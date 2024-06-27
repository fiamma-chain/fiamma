#!/bin/bash

set -e

GO_VERSION="1.22.3"  # Set the desired Go version

# Update system packages
echo "Updating system packages..."
sudo apt-get update

# Install common dependencies
echo "Installing common dependencies..."
sudo apt-get install -y make git bash gcc curl jq pkg-config openssl libssl-dev

# Check if Go is installed and install if not
if go version &>/dev/null; then
    echo "Go is already installed. Skipping installation."
else
    echo "Go is not installed. Installing Go..."
    wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
    echo 'export GOROOT=/usr/local/go' >> ~/.profile
    echo 'export GOPATH=$HOME/go' >> ~/.profile
    echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> ~/.profile
    source ~/.profile
fi


# Check if Rust is installed and install if not
if rustc --version &>/dev/null; then
    echo "Rust is already installed. Skipping installation."
else
    echo "Rust is not installed. Installing Rust..."
    curl https://sh.rustup.rs -sSf | bash -s -- -y 
    source $HOME/.cargo/env
    echo "source \$HOME/.cargo/env" >> ~/.profile
fi
    

# Check installations
echo "Checking installations..."
go version
rustc --version
echo "Environment setup complete."
