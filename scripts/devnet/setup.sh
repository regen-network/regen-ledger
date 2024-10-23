#!/bin/bash

set -e

# 🎨 Colors for better visibility
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# 🚀 Starting the setup
echo -e "${GREEN}🚀 Starting the Regen network setup...${NC}"

# Check if node count is provided, default to 3
NODE_COUNT=${1:-3}
echo -e "🔢 Setting up $NODE_COUNT nodes."

# 🧹 Clean up existing data directories
echo -e "🧹 Cleaning up existing data directories..."
rm -rf ./shared ./node*_data .env docker-compose.yaml

# 🛠 Create a temporary directory for key generation
TEMP_DIR=$(mktemp -d)
echo -e "🔑 Generating validator keys in temporary directory: $TEMP_DIR"

NODE_NAMES=()
NODE_ADDRESSES=()
NODE_MNEMONICS=()

# 📦 Generate keys for each node dynamically
for i in $(seq 1 "$NODE_COUNT"); do
  NODE="regen-node$i"
  NODE_NAMES+=("$NODE")
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

# 📝 Write the .env file dynamically
echo -e "📝 Writing ${GREEN}.env${NC} file..."
rm -f .env
for i in "${!NODE_NAMES[@]}"; do
  NODE="${NODE_NAMES[$i]}"
  ADDRESS="${NODE_ADDRESSES[$i]}"
  MNEMONIC="${NODE_MNEMONICS[$i]}"

  # Replace hyphens with underscores for valid environment variable names
  NODE_ENV_NAME=$(echo "${NODE^^}" | tr '-' '_')

  echo "${NODE_ENV_NAME}_VALIDATOR_ADDRESS=${ADDRESS}" >> .env
  echo "${NODE_ENV_NAME}_VALIDATOR_MNEMONIC=\"${MNEMONIC}\"" >> .env
done

# 📝 Generate the `docker-compose.yaml` file dynamically
echo -e "📝 Generating ${GREEN}docker-compose.yaml${NC} file..."
cat <<EOF > docker-compose.yaml
version: '3.8'

services:
EOF

for i in $(seq 1 "$NODE_COUNT"); do
  RPC_PORT=$((26657 + (i - 1) * 10))
  P2P_PORT=$((26656 + (i - 1) * 10))
  NODE="regen-node$i"

  cat <<EOF >> docker-compose.yaml
  $NODE:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: $NODE
    environment:
      - NODE_NAME=$NODE
      - NODE_COUNT=$NODE_COUNT
      - RPC_PORT=$RPC_PORT
      - P2P_PORT=$P2P_PORT
$(for j in $(seq 1 "$NODE_COUNT"); do
  PEER_NODE="regen-node$j"
  PEER_ENV=$(echo "${PEER_NODE^^}" | tr '-' '_')
  echo "      - ${PEER_ENV}_VALIDATOR_ADDRESS=\${${PEER_ENV}_VALIDATOR_ADDRESS}"
  echo "      - ${PEER_ENV}_VALIDATOR_MNEMONIC=\${${PEER_ENV}_VALIDATOR_MNEMONIC}"
done)
    volumes:
      - ./shared:/mnt/nvme/shared
      - ./entrypoint.sh:/entrypoint.sh
    ports:
      - "$P2P_PORT:$P2P_PORT"
      - "$RPC_PORT:$RPC_PORT"
    entrypoint: ["/bin/bash", "/entrypoint.sh"]
    networks:
      - regen-network
EOF
done

cat <<EOF >> docker-compose.yaml
networks:
  regen-network:
    driver: bridge
EOF

echo -e "${GREEN}✅ docker-compose.yaml${NC} file generated."

# 🐳 Start Docker Compose
echo -e "${GREEN}🐳 Starting the Regen network with Docker Compose...${NC}"
docker compose up --build

# 🧹 Clean up temporary files
rm -rf "$TEMP_DIR"
echo -e "🧹 Cleaned up temporary files."
