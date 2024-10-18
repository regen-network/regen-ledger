#!/bin/bash

# setup.sh - Automate the setup of a multi-node Regen network with Docker Compose

set -e

# 🎨 Colors for better visibility
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# 🚀 Starting the setup
echo -e "${GREEN}🚀 Starting the Regen network setup...${NC}"


# 🧹 Clean up existing data directories
echo -e "🧹 Cleaning up existing data directories..."
rm -rf ./shared
rm -rf ./node1_data
rm -rf ./node2_data
rm -rf ./node3_data
# 🛠 Create a temporary directory for key generation
TEMP_DIR=$(mktemp -d)
echo -e "🔑 Generating validator keys in temporary directory: $TEMP_DIR"

# 📦 Node configurations
NODE_NAMES=("regen-node1" "regen-node2" "regen-node3")
NODE_ADDRESSES=()
NODE_MNEMONICS=()

# 📝 Generate keys for each node
for NODE in "${NODE_NAMES[@]}"; do
  echo -e "🔐 Generating keys for ${GREEN}$NODE${NC}..."
  NODE_HOME="$TEMP_DIR/$NODE"
  mkdir -p "$NODE_HOME"
  KEY_OUTPUT=$(regen keys add my_validator --keyring-backend test --home "$NODE_HOME" --output json)

  # Extract address and mnemonic
  ADDRESS=$(echo "$KEY_OUTPUT" | jq -r '.address')
  MNEMONIC=$(echo "$KEY_OUTPUT" | jq -r '.mnemonic')

  NODE_ADDRESSES+=("$ADDRESS")
  NODE_MNEMONICS+=("$MNEMONIC")

  echo -e "📬 Address for ${GREEN}$NODE${NC}: ${ADDRESS}"
done

# 📝 Write the .env file
echo -e "📝 Writing ${GREEN}.env${NC} file..."
rm -f .env
for i in "${!NODE_NAMES[@]}"; do
  NODE="${NODE_NAMES[$i]}"
  ADDRESS="${NODE_ADDRESSES[$i]}"
  MNEMONIC="${NODE_MNEMONICS[$i]}"

  # Replace hyphens with underscores for environment variable names
  NODE_ENV_NAME=$(echo "${NODE^^}" | tr '-' '_')

  echo "${NODE_ENV_NAME}_VALIDATOR_ADDRESS=${ADDRESS}" >> .env
  echo "${NODE_ENV_NAME}_VALIDATOR_MNEMONIC=\"${MNEMONIC}\"" >> .env
done

echo -e "✅ ${GREEN}.env${NC} file has been written."

# 🐳 Starting Docker Compose
echo -e "${GREEN}🐳 Starting the Regen network with Docker Compose...${NC}"
docker compose up --build

echo -e "${GREEN}🎉 Regen network setup complete!${NC}"

# 🧹 Clean up temporary directory
rm -rf "$TEMP_DIR"
echo -e "🧹 Cleaned up temporary files."
