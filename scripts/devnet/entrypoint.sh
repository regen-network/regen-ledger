#!/bin/bash
set -e

# Constants
BASE_PATH=${BASE_PATH:-/mnt/nvme}
HOME_DIR=${BASE_PATH}/.regen
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

# Helper: Print message with emoji
log() { echo -e "${1} ${2}"; }

# Ensure shared directories exist
mkdir -p "$GENTX_DIR"

# Determine the number of nodes
NODE_COUNT="${NODE_COUNT:-3}"
NODE_NAMES=($(for i in $(seq 1 "$NODE_COUNT"); do echo "regen-node$i"; done))

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
    log "$INFO" "Adding ${NODE}'s address: ${ADDRESS}"
    regen add-genesis-account "$ADDRESS" 100000000uregen --home "$HOME_DIR"
  done
  log "$SUCCESS" "All validator accounts added to genesis."
}

generate_gentx() {
  log "$INFO" "Generating gentx for ${NODE_NAME}..."
  echo "$VALIDATOR_MNEMONIC" | regen keys add my_validator --recover --keyring-backend test --home "$HOME_DIR"
  regen gentx my_validator 50000000uregen --keyring-backend test --chain-id "$CHAIN_ID" --home "$HOME_DIR"
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
  log "$INFO" "Collecting and validating gentx files..."
  regen collect-gentxs --gentx-dir "$GENTX_DIR" --home "$HOME_DIR"
  regen validate-genesis --home "$HOME_DIR"
  cp "$HOME_DIR/config/genesis.json" "$SHARED_DIR/genesis.json"
  touch "$FINAL_GENESIS_READY"
  log "$SUCCESS" "Finalized genesis.json saved."
}

configure_peers() {
  log "$INFO" "Configuring persistent peers..."
  PERSISTENT_PEERS=""
  for NODE in "${NODE_NAMES[@]}"; do
    if [ "$NODE" != "$NODE_NAME" ]; then
      PEER_ID=$(cat "$SHARED_DIR/${NODE}_id")
      PERSISTENT_PEERS+="$PEER_ID@$NODE:26656,"
    fi
  done
  PERSISTENT_PEERS=${PERSISTENT_PEERS%,}
  sed -i "s/^persistent_peers = .*/persistent_peers = \"$PERSISTENT_PEERS\"/" "$HOME_DIR/config/config.toml"
  log "$SUCCESS" "Peers configured."
}

# Main Setup Logic
log "$INFO" "Starting setup for ${NODE_NAME}..."

configure_rpc() {
  CONFIG_FILE="$HOME_DIR/config/config.toml"
  log "$INFO" "Configuring RPC endpoint..."

  # Update RPC listener address in config.toml
  sed -i 's/^laddr *=.*/laddr = "tcp:\/\/0.0.0.0:26657"/' "$CONFIG_FILE"

  # Ensure other RPC parameters are correct
  sed -i 's/^grpc_laddr *=.*/grpc_laddr = ""/' "$CONFIG_FILE"
  sed -i 's/^rpc_servers *=.*/rpc_servers = ""/' "$CONFIG_FILE"

  log "$SUCCESS" "RPC endpoint configured."
}


fetch_environment_variables
initialize_node
save_node_id

if [ "$NODE_NAME" == "regen-node1" ]; then
  jq '.app_state.staking.params.bond_denom = "uregen"' "$HOME_DIR/config/genesis.json" > "$HOME_DIR/config/genesis_tmp.json"
  mv "$HOME_DIR/config/genesis_tmp.json" "$HOME_DIR/config/genesis.json"
  log "$SUCCESS" "Modified genesis.json for ${NODE_NAME}."

  add_validator_accounts_to_genesis

  # Save initial genesis.json and notify other nodes
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

configure_peers
configure_rpc

log "$INFO" "Starting ${NODE_NAME}..."
exec regen start --home "$HOME_DIR" --minimum-gas-prices="0.025uregen"
