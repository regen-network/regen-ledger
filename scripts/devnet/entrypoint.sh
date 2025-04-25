#!/bin/bash
set -e

# Constants
BASE_PATH=${BASE_PATH:-/mnt/nvme}
SHARED_DIR=${BASE_PATH}/shared
GENTX_DIR="$SHARED_DIR/gentxs"
INITIAL_GENESIS_READY="$SHARED_DIR/initial_genesis_ready"
FINAL_GENESIS_READY="$SHARED_DIR/final_genesis_ready"
CHAIN_ID="regen-devnet"

# Colors and Emojis
GREEN='\033[0;32m'
NC='\033[0m'
INFO="${GREEN}ℹ️${NC}"
SUCCESS="${GREEN}✅${NC}"
WAIT="${GREEN}⏳${NC}"

log() { echo -e "${1} ${2}"; }

mkdir -p "$GENTX_DIR"

NODE_COUNT="${NODE_COUNT:-3}"
NODE_NAMES=($(for i in $(seq 1 "$NODE_COUNT"); do echo "regen-node$i"; done))
HOME_DIR="${BASE_PATH}/.regen/${NODE_NAME}"

P2P_PORT=${P2P_PORT:-26656}
RPC_PORT=${RPC_PORT:-26657}
GRPC_PORT=${GRPC_PORT:-9090}

configure_rpc_and_p2p() {
  CONFIG_FILE="$HOME_DIR/config/config.toml"
  APP_FILE="$HOME_DIR/config/app.toml"

  log "$INFO" "🔍 Verifying configuration of RPC, P2P, and gRPC for ${NODE_NAME}..."

  if [ -f "$CONFIG_FILE" ]; then
    sed -i "/\[rpc\]/,/^\[.*\]/ s|^laddr *=.*|laddr = \"tcp://0.0.0.0:$RPC_PORT\"|" "$CONFIG_FILE"
    sed -i "/\[p2p\]/,/^\[.*\]/ s|^laddr *=.*|laddr = \"tcp://0.0.0.0:$P2P_PORT\"|" "$CONFIG_FILE"
    sed -i "/\[p2p\]/,/^\[.*\]/ s|^external_address *=.*|external_address = \"tcp://0.0.0.0:$P2P_PORT\"|" "$CONFIG_FILE"
    log "$SUCCESS" "✅ Configured RPC on $RPC_PORT and P2P on $P2P_PORT in $CONFIG_FILE."
  fi

  if [ -f "$APP_FILE" ]; then
    sed -i "/\[grpc\]/,/^\[.*\]/ s|^address *=.*|address = \"localhost:$GRPC_PORT\"|" "$APP_FILE"
    log "$SUCCESS" "✅ Configured gRPC on $GRPC_PORT and API enabled in $APP_FILE."
  fi
}

fetch_environment_variables() {
  NODE_ENV_NAME=$(echo "${NODE_NAME^^}" | tr '-' '_')
  VALIDATOR_MNEMONIC_VAR="${NODE_ENV_NAME}_VALIDATOR_MNEMONIC"
  VALIDATOR_ADDRESS_VAR="${NODE_ENV_NAME}_VALIDATOR_ADDRESS"

  VALIDATOR_MNEMONIC="${!VALIDATOR_MNEMONIC_VAR}"
  VALIDATOR_ADDRESS="${!VALIDATOR_ADDRESS_VAR}"

  if [ -z "$VALIDATOR_MNEMONIC" ] || [ -z "$VALIDATOR_ADDRESS" ]; then
    log "$WAIT" "Mnemonic or address not found for ${NODE_NAME}!"
    exit 1
  fi
  log "$SUCCESS" "Fetched mnemonic and address for ${NODE_NAME}."
}

initialize_node() {
  if [ ! -f "$HOME_DIR/config/node_key.json" ]; then
    regen init "$NODE_NAME" --chain-id "$CHAIN_ID" --home "$HOME_DIR"
    log "$SUCCESS" "Initialized ${NODE_NAME}."
  fi
}

save_node_id() {
  NODE_ID=$(regen tendermint show-node-id --home "$HOME_DIR")
  echo "$NODE_ID" > "$SHARED_DIR/${NODE_NAME}_id"
  log "$SUCCESS" "Saved Node ID for ${NODE_NAME}: $NODE_ID"
}

add_validator_accounts_to_genesis() {
  log "$INFO" "Adding validator accounts to genesis..."
  for NODE in "${NODE_NAMES[@]}"; do
    NODE_ENV=$(echo "${NODE^^}" | tr '-' '_')
    ADDR_VAR="${NODE_ENV}_VALIDATOR_ADDRESS"
    ADDRESS="${!ADDR_VAR}"
    regen genesis add-genesis-account "$ADDRESS" 100000000uregen --home "$HOME_DIR"
  done
  log "$SUCCESS" "All validator accounts added to genesis."
}

generate_gentx() {
  log "$INFO" "Generating gentx for ${NODE_NAME}..."
  echo "$VALIDATOR_MNEMONIC" | regen keys add my_validator --recover --keyring-backend test --home "$HOME_DIR"
  regen genesis gentx my_validator 50000000uregen --keyring-backend test --chain-id "$CHAIN_ID" --home "$HOME_DIR"
  cp "$HOME_DIR/config/gentx/"*.json "$GENTX_DIR/${NODE_NAME}_gentx.json"
  log "$SUCCESS" "Generated gentx for ${NODE_NAME}."
}

wait_for_gentx_files() {
  log "$INFO" "Waiting for all gentx files..."
  for NODE in "${NODE_NAMES[@]}"; do
    if [ "$NODE" != "$NODE_NAME" ]; then
      while [ ! -f "$GENTX_DIR/${NODE}_gentx.json" ]; do
        log "$WAIT" "Waiting for gentx from ${NODE}..."
        sleep 2
      done
      log "$SUCCESS" "Received gentx from ${NODE}."
    fi
  done
}

wait_for_initial_genesis() {
  log "$INFO" "${NODE_NAME} waiting for initial genesis.json..."
  while [ ! -f "$INITIAL_GENESIS_READY" ]; do
    sleep 2
  done
  cp "$SHARED_DIR/genesis.json" "$HOME_DIR/config/genesis.json"
  log "$SUCCESS" "Initial genesis.json received for ${NODE_NAME}."
}

wait_for_final_genesis() {
  log "$INFO" "Waiting for finalized genesis.json..."
  while [ ! -f "$FINAL_GENESIS_READY" ]; do
    sleep 2
  done
  cp "$SHARED_DIR/genesis.json" "$HOME_DIR/config/genesis.json"
  log "$SUCCESS" "Finalized genesis.json received for ${NODE_NAME}."
}

collect_and_finalize_genesis() {
  regen genesis collect-gentxs --gentx-dir "$GENTX_DIR" --home "$HOME_DIR"
  regen genesis validate-genesis --home "$HOME_DIR"
  cp "$HOME_DIR/config/genesis.json" "$SHARED_DIR/genesis.json"
  touch "$FINAL_GENESIS_READY"
  log "$SUCCESS" "Finalized genesis.json saved."
}

configure_peers() {
  CONFIG_FILE="$HOME_DIR/config/config.toml"
  log "$INFO" "🔗 Configuring persistent peers..."

  PEERS=""
  for NODE in "${NODE_NAMES[@]}"; do
    if [ "$NODE" != "$NODE_NAME" ]; then
      while [ ! -f "$SHARED_DIR/${NODE}_id" ]; do
        log "$WAIT" "Waiting for ${NODE}'s node ID..."
        sleep 2
      done
      NODE_ID=$(cat "$SHARED_DIR/${NODE}_id")

      PORT_VAR=$(echo "${NODE^^}_P2P_PORT" | tr '-' '_')
      PORT="${!PORT_VAR:-26000}"

      PEERS+="$NODE_ID@$NODE:$PORT,"
    fi
  done

  PEERS="${PEERS%,}"
  sed -i "s/^persistent_peers =.*/persistent_peers = \"$PEERS\"/" "$CONFIG_FILE"
  log "$SUCCESS" "✅ Configured persistent peers: $PEERS"
}

wait_for_chain_ready() {
  local max_tries=40
  local try=0

  log "$INFO" "⏳ Waiting for chain to reach first block at RPC port $RPC_PORT..."

  until [ "$(curl -sf http://localhost:${RPC_PORT}/status | jq -r '.result.sync_info.latest_block_height')" -gt 0 ]; do
    try=$((try+1))
    if [ $try -ge $max_tries ]; then
      log "$WAIT" "❌ Chain did not reach block height > 0 in time on port ${RPC_PORT}"
      exit 1
    fi
    sleep 2
  done

  log "$SUCCESS" "🎯 Chain is live with block height > 0 on RPC port $RPC_PORT"
}

create_validator_tx() {
  if [ "$NODE_NAME" != "regen-node1" ]; then
    log "$INFO" "⛏ Creating validator for ${NODE_NAME}..."

    echo "$VALIDATOR_MNEMONIC" | regen keys add my_validator --recover --keyring-backend test --home "$HOME_DIR" || true

    if regen query staking validator "$VALIDATOR_ADDRESS" --node "tcp://localhost:${RPC_PORT}" --output json | jq -e '.validator' > /dev/null 2>&1; then
      log "$WAIT" "Validator already exists for ${NODE_NAME}. Skipping creation."
      return
    fi

    regen tx staking create-validator \
      --amount=50000000uregen \
      --pubkey="$(regen tendermint show-validator --home "$HOME_DIR")" \
      --moniker="$NODE_NAME" \
      --chain-id="$CHAIN_ID" \
      --commission-rate="0.10" \
      --commission-max-rate="0.20" \
      --commission-max-change-rate="0.01" \
      --min-self-delegation="1" \
      --from=my_validator \
      --gas=auto \
      --gas-adjustment=1.5 \
      --yes \
      --keyring-backend=test \
      --home="$HOME_DIR" \
      --broadcast-mode=block \
      --node "tcp://localhost:${RPC_PORT}"

    log "$SUCCESS" "Validator created for ${NODE_NAME}."
  fi
}

### 🔧 Main Logic
log "$INFO" "Starting setup for ${NODE_NAME}..."

fetch_environment_variables
initialize_node
save_node_id
configure_rpc_and_p2p

if [ "$NODE_NAME" == "regen-node1" ]; then
  jq '.app_state.staking.params.bond_denom = "uregen"' "$HOME_DIR/config/genesis.json" > "$HOME_DIR/config/genesis_tmp.json"
  mv "$HOME_DIR/config/genesis_tmp.json" "$HOME_DIR/config/genesis.json"
  log "$SUCCESS" "Modified genesis.json for ${NODE_NAME}."

  add_validator_accounts_to_genesis
  cp "$HOME_DIR/config/genesis.json" "$SHARED_DIR/genesis.json"
  touch "$INITIAL_GENESIS_READY"
  log "$SUCCESS" "Initial genesis.json saved."

  wait_for_gentx_files
  collect_and_finalize_genesis
else
  wait_for_initial_genesis
  generate_gentx
  wait_for_final_genesis
fi

log "$INFO" "Starting ${NODE_NAME}..."
configure_peers

regen start --home "$HOME_DIR" --minimum-gas-prices="0.025uregen" &

wait_for_chain_ready
create_validator_tx
wait