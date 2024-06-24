#!/usr/bin/env bash

set -eo pipefail


# Directory containing proto files relative to the script's location
PROTO_DIR="./proto"

# Path to the buf template, assuming it's in the proto directory
TEMPLATE_PATH="${PROTO_DIR}/buf.gen.gogo.yaml"
# Create a temporary directory for the output
OUTPUT_DIR="."

# Find and process all proto files
find "${PROTO_DIR}" -name '*.proto' -print0 | while IFS= read -r -d '' proto_file; do
    if grep go_package $proto_file &>/dev/null; then
      buf generate --template "${TEMPLATE_PATH}" --output "${OUTPUT_DIR}" --error-format=json --log-format=json "${proto_file}"
      if [ $? -ne 0 ]; then
          echo "Failed to process ${proto_file}"
          exit 1
      fi
    fi  
done


# move proto files to the right places
cp -r fiamma/* ./
rm -rf fiamma

go mod tidy 

echo "Proto file generation completed successfully."