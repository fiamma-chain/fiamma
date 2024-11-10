#!/bin/bash

set -e

if [ $# -ne 1 ]; then
  echo "Usage: $0 <namespace> " 
  echo "accepts 1 arg, received $#"
  exit 1
else
  NAMESPACE=$1
fi


: ${NODE:="https://testnet-rpc.fiammachain.io"}


fiammad query zkpverify get-verify-results-by-namespace $NAMESPACE --node $NODE
