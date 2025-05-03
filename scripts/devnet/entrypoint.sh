#!/bin/bash
set -e

# Constants
BASE_PATH=${BASE_PATH:-/mnt/nvme}
SHARED_DIR=${BASE_PATH}/shared
GENTX_DIR="$SHARED_DIR/gentxs"
INITIAL_GENESIS_READY="$SHARED_DIR/initial_genesis_ready"
FINAL_GENESIS_READY="$SHARED_DIR/final_genesis_ready"
CHAIN_ID="regen-devnet"
HOME_DIR="${BASE_PATH}/.regen/${NODE_NAME}"

# Styling
GREEN='\033[0;32m'
NC='\033[0m'
INFO="${GREEN}ℹ️${NC}"
SUCCESS="${GREEN}✅${NC}"
WAIT="${GREEN}⏳${NC}"
log() { echo -e "$1 $2"; }

mkdir -p "$GENTX_DIR"

NODE_COUNT="${NODE_COUNT:-3}"
NODE_NAMES=($(for i in $(seq 1 "$NODE_COUNT"); do echo "regen-node$i"; done))

# Detect CLI version
use_new_cli=false
if [[ "$REGEN_VERSION_MAJOR" == "v6" || "$REGEN_VERSION_MAJOR" == "v7" ]]; then
  use_new_cli=true
fi

# CLI wrapper functions
add_genesis_account() {
  local address="$1"
  local amount="1000000000uregen"  # Pre-fund with 1,000,000,000 uregen for devnet purposes

  if [ "$use_new_cli" = true ]; then
    regen genesis add-genesis-account "$address" "$amount" --home "$HOME_DIR"
  else
    regen add-genesis-account "$address" "$amount" --home "$HOME_DIR"
  fi
}

generate_gentx() {
  echo "$VALIDATOR_MNEMONIC" > "$HOME_DIR/mnemonic.txt"
  chmod 600 "$HOME_DIR/mnemonic.txt"

  expect <<EOF
spawn regen keys add my_validator --recover --keyring-backend=test --home=$HOME_DIR
expect "Enter your bip39 mnemonic"
send "$(cat $HOME_DIR/mnemonic.txt)\r"
expect eof
EOF

  if [ "$use_new_cli" = true ]; then
    regen genesis gentx my_validator 50000000uregen --keyring-backend test --chain-id "$CHAIN_ID" --home "$HOME_DIR"
  else
    regen gentx my_validator 50000000uregen --keyring-backend test --chain-id "$CHAIN_ID" --home "$HOME_DIR"
  fi
  cp "$HOME_DIR/config/gentx/"*.json "$GENTX_DIR/${NODE_NAME}_gentx.json"
  log "$SUCCESS" "Generated gentx for ${NODE_NAME}."
}

collect_gentxs() {
  if [ "$use_new_cli" = true ]; then
    regen genesis collect-gentxs --gentx-dir "$GENTX_DIR" --home "$HOME_DIR"
    regen genesis validate-genesis --home "$HOME_DIR"
  else
    regen collect-gentxs --gentx-dir "$GENTX_DIR" --home "$HOME_DIR"
    regen validate-genesis --home "$HOME_DIR"
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
  log "$SUCCESS" "✅ Fetched mnemonic and address for ${NODE_NAME}."
}

initialize_node() {
  if [ ! -f "$HOME_DIR/config/node_key.json" ]; then
    regen init "$NODE_NAME" --chain-id "$CHAIN_ID" --home "$HOME_DIR"
    log "$SUCCESS" "✅ Initialized ${NODE_NAME}."
  fi
}

setup_cosmovisor_layout() {
  mkdir -p "$HOME_DIR/cosmovisor/genesis/bin"
  mkdir -p "$HOME_DIR/cosmovisor/upgrades/regen-v6-upgrade/bin"

  cp /upgrade-binaries/regen-v5 "$HOME_DIR/cosmovisor/genesis/bin/regen"
  cp /upgrade-binaries/regen-v6 "$HOME_DIR/cosmovisor/upgrades/regen-v6-upgrade/bin/regen"

  chmod +x "$HOME_DIR"/cosmovisor/**/bin/regen

  log "$SUCCESS" "✅ Cosmovisor layout set up for ${NODE_NAME}"
}

save_node_id() {
  NODE_ID=$(regen tendermint show-node-id --home "$HOME_DIR")
  echo "$NODE_ID" > "$SHARED_DIR/${NODE_NAME}_id"
  log "$SUCCESS" "✅ Saved Node ID for ${NODE_NAME}: $NODE_ID"
}

configure_rpc_and_p2p() {
  CONFIG_FILE="$HOME_DIR/config/config.toml"
  APP_FILE="$HOME_DIR/config/app.toml"

  sed -i "/\[rpc\]/,/^\[.*\]/ s|^laddr *=.*|laddr = \"tcp://0.0.0.0:$RPC_PORT\"|" "$CONFIG_FILE"
  sed -i "/\[p2p\]/,/^\[.*\]/ s|^laddr *=.*|laddr = \"tcp://0.0.0.0:$P2P_PORT\"|" "$CONFIG_FILE"
  sed -i "/\[p2p\]/,/^\[.*\]/ s|^external_address *=.*|external_address = \"tcp://0.0.0.0:$P2P_PORT\"|" "$CONFIG_FILE"
  sed -i "/\[grpc\]/,/^\[.*\]/ s|^address *=.*|address = \"localhost:$GRPC_PORT\"|" "$APP_FILE"

  log "$SUCCESS" "✅ Configured ports for ${NODE_NAME}"
}

configure_peers() {
  CONFIG_FILE="$HOME_DIR/config/config.toml"
  PEERS=""
  for NODE in "${NODE_NAMES[@]}"; do
    if [ "$NODE" != "$NODE_NAME" ]; then
      while [ ! -f "$SHARED_DIR/${NODE}_id" ]; do sleep 2; done
      NODE_ID=$(cat "$SHARED_DIR/${NODE}_id")
      INDEX=$(echo "$NODE" | grep -o '[0-9]\+$')
      PEER_PORT=$((26000 + (INDEX - 1) * 3))
      PEERS+="$NODE_ID@$NODE:$PEER_PORT,"
    fi
  done
  PEERS="${PEERS%,}"
  sed -i "s/^persistent_peers =.*/persistent_peers = \"$PEERS\"/" "$CONFIG_FILE"
  log "$SUCCESS" "✅ Configured persistent peers: $PEERS"
}

wait_for_gentx_files() {
  for NODE in "${NODE_NAMES[@]}"; do
    if [ "$NODE" != "$NODE_NAME" ]; then
      while [ ! -f "$GENTX_DIR/${NODE}_gentx.json" ]; do sleep 2; done
      log "$SUCCESS" "✅ Received gentx from ${NODE}."
    fi
  done
}

wait_for_initial_genesis() {
  while [ ! -f "$INITIAL_GENESIS_READY" ]; do sleep 2; done
  cp "$SHARED_DIR/genesis.json" "$HOME_DIR/config/genesis.json"
  log "$SUCCESS" "✅ Initial genesis received."
}

wait_for_final_genesis() {
  while [ ! -f "$FINAL_GENESIS_READY" ]; do sleep 2; done
  cp "$SHARED_DIR/genesis.json" "$HOME_DIR/config/genesis.json"
  log "$SUCCESS" "✅ Final genesis received."
}

wait_for_chain_ready() {
  local max_tries=40 try=0
  until [ "$(curl -sf http://localhost:${RPC_PORT}/status | jq -r '.result.sync_info.latest_block_height')" -gt 0 ]; do
    try=$((try+1))
    if [ $try -ge $max_tries ]; then
      log "$WAIT" "❌ Chain failed to start on port ${RPC_PORT}"
      exit 1
    fi
    sleep 2
  done
  log "$SUCCESS" "🌟 Chain is live on $RPC_PORT"
}

create_validator_tx() {
  if [ "$NODE_NAME" != "regen-node1" ]; then
    echo "$VALIDATOR_MNEMONIC" | regen keys add my_validator --recover --keyring-backend test --home "$HOME_DIR" || true
    if regen query staking validator "$VALIDATOR_ADDRESS" --node "tcp://localhost:${RPC_PORT}" --output json | jq -e '.validator' > /dev/null 2>&1; then
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
  fi
}
add_validator_accounts_to_genesis() {
  log "$INFO" "Adding validator accounts to genesis..."
  for NODE in "${NODE_NAMES[@]}"; do
    NODE_ENV=$(echo "${NODE^^}" | tr '-' '_')
    ADDR_VAR="${NODE_ENV}_VALIDATOR_ADDRESS"
    ADDRESS="${!ADDR_VAR}"
    add_genesis_account "$ADDRESS"
  done
  log "$SUCCESS" "✅ Added validator accounts to genesis."
}

# Main
log "$INFO" "🛠️ Starting setup for ${NODE_NAME}..."
fetch_environment_variables
initialize_node
setup_cosmovisor_layout
save_node_id
configure_rpc_and_p2p

if [ "$NODE_NAME" == "regen-node1" ]; then
  # Ensure staking denom is set to uregen
  jq '.app_state.staking.params.bond_denom = "uregen"' "$HOME_DIR/config/genesis.json" > "$HOME_DIR/config/genesis_tmp.json"
  mv "$HOME_DIR/config/genesis_tmp.json" "$HOME_DIR/config/genesis.json"
  log "$SUCCESS" "✅ Updated staking bond_denom to uregen"

  # Ensure gov min_deposit denom matches staking denom
  jq '.app_state.gov.deposit_params.min_deposit[0].denom = "uregen"' "$HOME_DIR/config/genesis.json" > "$HOME_DIR/config/genesis_tmp.json"
  mv "$HOME_DIR/config/genesis_tmp.json" "$HOME_DIR/config/genesis.json"
  log "$SUCCESS" "✅ Updated gov min_deposit denom to uregen"

  jq '.app_state.gov.voting_params.voting_period = "60s"' "$HOME_DIR/config/genesis.json" > "$HOME_DIR/config/genesis_tmp.json" && \
  mv "$HOME_DIR/config/genesis_tmp.json" "$HOME_DIR/config/genesis.json"

  add_validator_accounts_to_genesis
  cp "$HOME_DIR/config/genesis.json" "$SHARED_DIR/genesis.json"
  touch "$INITIAL_GENESIS_READY"
  wait_for_gentx_files
  collect_gentxs
  cp "$HOME_DIR/config/genesis.json" "$SHARED_DIR/genesis.json"
  touch "$FINAL_GENESIS_READY"
else
  wait_for_initial_genesis
  generate_gentx
  wait_for_final_genesis
fi

configure_peers

export DAEMON_HOME="$HOME_DIR"
export DAEMON_NAME="regen"
export DAEMON_ALLOW_DOWNLOAD_BINARIES=false
export DAEMON_RESTART_AFTER_UPGRADE=true
export UNSAFE_SKIP_BACKUP=true

cosmovisor run start --home "$HOME_DIR" --minimum-gas-prices="0.025uregen" &
wait_for_chain_ready
create_validator_tx
wait
