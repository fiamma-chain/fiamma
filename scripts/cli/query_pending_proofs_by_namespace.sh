#!/bin/bash

set -e

if [ $# -ne 1 ]; then
  echo "Usage: $0 <namespace> " 
  echo "accepts 1 arg, received $#"
  exit 1
else
  NAMESPACE=$1
fi

: ${NODE:="http://127.0.0.1:26657"}

fiammad query zkpverify pending-proof-by-namespace $NAMESPACE --node $NODE
