#!/bin/bash
set -e

# 🎨 Colors
GREEN='\033[0;32m'
NC='\033[0m'

echo -e "${GREEN}🚀 Starting the Regen network setup...${NC}"

# 🔧 Default: 3 nodes
NODE_COUNT=${1:-3}
echo -e "🔢 Setting up $NODE_COUNT nodes."

# 🌱 Regen version (affects CLI syntax)
REGEN_VERSION=${REGEN_VERSION:-v5.1.4}
REGEN_VERSION_MAJOR=$(echo "$REGEN_VERSION" | cut -d. -f1) # → "v5"
echo -e "📦 Detected Regen version: ${GREEN}$REGEN_VERSION${NC} ($REGEN_VERSION_MAJOR)"

# 🧹 Clean up
echo -e "🧹 Cleaning up..."
rm -rf ./shared ./node*_data .env docker-compose.yaml

# 🔐 Temp dir for keys
TEMP_DIR=$(mktemp -d)
echo -e "🔑 Generating validator keys in: $TEMP_DIR"

NODE_NAMES=()
NODE_ADDRESSES=()
NODE_MNEMONICS=()

for i in $(seq 1 "$NODE_COUNT"); do
  NODE="regen-node$i"
  NODE_NAMES+=("$NODE")
  echo -e "🔐 Creating keys for ${GREEN}$NODE${NC}..."

  NODE_HOME="$TEMP_DIR/$NODE"
  mkdir -p "$NODE_HOME"
  KEY_OUTPUT=$(regen keys add my_validator --keyring-backend test --home "$NODE_HOME" --output json)

  ADDRESS=$(echo "$KEY_OUTPUT" | jq -r '.address')
  MNEMONIC=$(echo "$KEY_OUTPUT" | jq -r '.mnemonic')

  NODE_ADDRESSES+=("$ADDRESS")
  NODE_MNEMONICS+=("$MNEMONIC")

  echo -e "📬 Address for ${GREEN}$NODE${NC}: $ADDRESS"
done

# 📝 Write .env
echo -e "📝 Writing ${GREEN}.env${NC} file..."
rm -f .env
echo "REGEN_VERSION_MAJOR=${REGEN_VERSION_MAJOR}" >> .env
for i in "${!NODE_NAMES[@]}"; do
  NODE="${NODE_NAMES[$i]}"
  ADDRESS="${NODE_ADDRESSES[$i]}"
  MNEMONIC="${NODE_MNEMONICS[$i]}"
  NODE_ENV_NAME=$(echo "${NODE^^}" | tr '-' '_')

  echo "${NODE_ENV_NAME}_VALIDATOR_ADDRESS=${ADDRESS}" >> .env
  echo "${NODE_ENV_NAME}_VALIDATOR_MNEMONIC=\"${MNEMONIC}\"" >> .env
done

# 🧱 docker-compose.yaml
echo -e "📝 Generating ${GREEN}docker-compose.yaml${NC}..."
cat <<EOF > docker-compose.yaml
services:
EOF

BASE_PORT=26000
for i in $(seq 0 $((NODE_COUNT - 1))); do
  P2P_PORT=$((BASE_PORT + i * 3))
  RPC_PORT=$((BASE_PORT + i * 3 + 1))
  GRPC_PORT=$((BASE_PORT + i * 3 + 2))
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
      - REGEN_VERSION_MAJOR=${REGEN_VERSION_MAJOR}
$(for j in $(seq 1 "$NODE_COUNT"); do
  PEER_NODE="regen-node$j"
  PEER_ENV=$(echo "${PEER_NODE^^}" | tr '-' '_')
  echo "      - ${PEER_ENV}_VALIDATOR_ADDRESS=\${${PEER_ENV}_VALIDATOR_ADDRESS}"
  echo "      - ${PEER_ENV}_VALIDATOR_MNEMONIC=\${${PEER_ENV}_VALIDATOR_MNEMONIC}"
done)
    volumes:
      - ./shared/node:/root/shared
      - ./shared/node$i-conf:/root/.regen
      - ./entrypoint.sh:/entrypoint.sh
      - ./upgrade-binaries:/upgrade-binaries
    networks:
      - regen-network
    ports:
      - "${RPC_PORT}:${RPC_PORT}"
    entrypoint: ["/bin/bash", "/entrypoint.sh"]

EOF
done

cat <<EOF >> docker-compose.yaml
networks:
  regen-network:
    driver: bridge
EOF

echo -e "${GREEN}✅ docker-compose.yaml${NC} generated."
echo -e "${GREEN}🐳 Launching devnet...${NC}"
docker compose up --build

# 🧹 Cleanup
rm -rf "$TEMP_DIR"
echo -e "🧹 Removed temp keys directory."
