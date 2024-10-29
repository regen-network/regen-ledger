#!/bin/bash

set -e

# 🎨 Colors for better visibility
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# 🚀 Start the setup
echo -e "${GREEN}🚀 Starting the Regen network setup...${NC}"

# Check if node count is provided; default to 3
NODE_COUNT=${1:-3}
echo -e "🔢 Setting up $NODE_COUNT nodes."

# 🧹 Clean up existing directories
echo -e "🧹 Cleaning up existing directories..."
rm -rf ./shared ./node*_data .env docker-compose.yaml

# Temporary directory for key generation
TEMP_DIR=$(mktemp -d)
echo -e "🔑 Generating validator keys in temporary directory: $TEMP_DIR"

NODE_NAMES=()
NODE_ADDRESSES=()
NODE_MNEMONICS=()

# 📦 Generate keys for each node
for i in $(seq 1 "$NODE_COUNT"); do
  NODE="regen-node$i"
  NODE_NAMES+=("$NODE")
  echo -e "🔐 Generating keys for ${GREEN}$NODE${NC}..."

  NODE_HOME="$TEMP_DIR/$NODE"
  mkdir -p "$NODE_HOME"
  KEY_OUTPUT=$(regen keys add my_validator --keyring-backend test --home "$NODE_HOME" --output json)

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
  NODE_ENV_NAME=$(echo "${NODE^^}" | tr '-' '_')

  echo "${NODE_ENV_NAME}_VALIDATOR_ADDRESS=${ADDRESS}" >> .env
  echo "${NODE_ENV_NAME}_VALIDATOR_MNEMONIC=\"${MNEMONIC}\"" >> .env
done

# 📝 Generate `docker-compose.yaml`
echo -e "📝 Generating ${GREEN}docker-compose.yaml${NC} file..."
cat <<EOF > docker-compose.yaml
services:
EOF

# ⚙️ Assign non-overlapping ports for each node
BASE_PORT=26000  # Start from a clean base to avoid collisions
for i in $(seq 0 $((NODE_COUNT - 1))); do
  P2P_PORT=$((BASE_PORT + i * 3))      # P2P port
  RPC_PORT=$((BASE_PORT + i * 3 + 1))  # RPC port
  GRPC_PORT=$((BASE_PORT + i * 3 + 2)) # gRPC port


  NODE="regen-node$((i + 1))"

  cat <<EOF >> docker-compose.yaml
  $NODE:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: $NODE
    environment:
      - NODE_NAME=$NODE
      - NODE_COUNT=$NODE_COUNT
      - P2P_PORT=$P2P_PORT
      - RPC_PORT=$RPC_PORT
      - GRPC_PORT=$GRPC_PORT
$(for j in $(seq 1 "$NODE_COUNT"); do
  PEER_NODE="regen-node$j"
  PEER_ENV=$(echo "${PEER_NODE^^}" | tr '-' '_')
  echo "      - ${PEER_ENV}_VALIDATOR_ADDRESS=\${${PEER_ENV}_VALIDATOR_ADDRESS}"
  echo "      - ${PEER_ENV}_VALIDATOR_MNEMONIC=\${${PEER_ENV}_VALIDATOR_MNEMONIC}"
done)
    volumes:
      - ./shared/node:/mnt/nvme/shared
      - ./shared/node$i-conf:/mnt/nvme/.regen
      - ./entrypoint.sh:/entrypoint.sh
    networks:
      - regen-network
    ports:
      - "${RPC_PORT}:${RPC_PORT}"
      - ":${P2P_PORT}:${P2P_PORT}"
      - ":${P2P_PORT}:${P2P_PORT}"
    entrypoint: ["/bin/bash", "/entrypoint.sh"]

EOF
done

cat <<EOF >> docker-compose.yaml
networks:
  regen-network:
    driver: bridge
EOF

echo -e "${GREEN}✅ docker-compose.yaml${NC} generated."

# 🐳 Start the Docker containers
echo -e "${GREEN}🐳 Starting the Regen network with Docker Compose...${NC}"
docker compose up --build

# 🧹 Clean up temporary files
rm -rf "$TEMP_DIR"
echo -e "🧹 Cleaned up temporary files."
