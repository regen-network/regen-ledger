#!/bin/bash
# Entry script to initialize and start the Regen node with persistent peers

set -e

# Set the base path for the node's home directory
BASE_PATH=${BASE_PATH:-/mnt/nvme}
HOME_DIR=${BASE_PATH}/.regen

# Create a directory for shared Node IDs and genesis files
SHARED_DIR=${BASE_PATH}/shared
mkdir -p "$SHARED_DIR/gentxs"

# List of all node names
NODE_NAMES=("regen-node1" "regen-node2" "regen-node3")

# Retrieve environment variables for validator address and mnemonic
NODE_ENV_NAME=$(echo "${NODE_NAME^^}" | tr '-' '_')

VALIDATOR_MNEMONIC_VAR="${NODE_ENV_NAME}_VALIDATOR_MNEMONIC"
VALIDATOR_ADDRESS_VAR="${NODE_ENV_NAME}_VALIDATOR_ADDRESS"

VALIDATOR_MNEMONIC="${!VALIDATOR_MNEMONIC_VAR}"
VALIDATOR_ADDRESS="${!VALIDATOR_ADDRESS_VAR}"

# Add debugging statements
echo "[$NODE_NAME] NODE_ENV_NAME: $NODE_ENV_NAME"
echo "[$NODE_NAME] VALIDATOR_MNEMONIC_VAR: $VALIDATOR_MNEMONIC_VAR"
echo "[$NODE_NAME] VALIDATOR_ADDRESS_VAR: $VALIDATOR_ADDRESS_VAR"
echo "[$NODE_NAME] VALIDATOR_MNEMONIC: $VALIDATOR_MNEMONIC"
echo "[$NODE_NAME] VALIDATOR_ADDRESS: $VALIDATOR_ADDRESS"

# Ensure environment variables are set
if [ -z "$VALIDATOR_MNEMONIC" ] || [ -z "$VALIDATOR_ADDRESS" ]; then
  echo "[$NODE_NAME] ERROR: Validator mnemonic and address must be set via environment variables."
  exit 1
fi

COMPLETION_FILE="$SHARED_DIR/regen-node1_genesis_init_done"

# **Regen-node1 initializes the genesis.json**
if [ "$NODE_NAME" == "regen-node1" ]; then
  echo "[$NODE_NAME] Initializing the Regen node..."
  regen init "$NODE_NAME" --chain-id regen-devnet --home "$HOME_DIR"

  # **Update the staking bond_denom to uregen**
  sed -i 's/"bond_denom": "stake"/"bond_denom": "uregen"/' "$HOME_DIR/config/genesis.json"

  # **Add all validator accounts to the genesis file**
  echo "[$NODE_NAME] Adding validator accounts to genesis..."
  for NODE in "${NODE_NAMES[@]}"; do
    NODE_LOOP_ENV_NAME=$(echo "${NODE^^}" | tr '-' '_')
    ADDRESS_VAR="${NODE_LOOP_ENV_NAME}_VALIDATOR_ADDRESS"
    ADDRESS="${!ADDRESS_VAR}"
    echo "[$NODE_NAME] Processing $NODE: ADDRESS_VAR=$ADDRESS_VAR, ADDRESS=$ADDRESS"

    if [ -z "$ADDRESS" ]; then
      echo "[$NODE_NAME] ERROR: Address for $NODE is not set."
      exit 1
    fi
    # Use the address directly
    regen add-genesis-account "$ADDRESS" 100000000uregen --home "$HOME_DIR"
  done

  # **Copy the initial genesis.json to the shared directory**
  cp "$HOME_DIR/config/genesis.json" "$SHARED_DIR/genesis.json"
  touch "$COMPLETION_FILE"

else
  # **Other nodes wait for genesis initialization by regen-node1**
  while [ ! -f "$COMPLETION_FILE" ]; do
    echo "[$NODE_NAME] Waiting for regen-node1 to initialize genesis..."
    sleep 2
  done

  echo "[$NODE_NAME] Copying genesis.json from regen-node1..."
  regen init "$NODE_NAME" --chain-id regen-devnet --home "$HOME_DIR"
  cp "$SHARED_DIR/genesis.json" "$HOME_DIR/config/genesis.json"
fi

# **Import the validator key using the mnemonic**
echo "[$NODE_NAME] Importing validator key..."
echo "$VALIDATOR_MNEMONIC" | regen keys add my_validator --recover --keyring-backend test --home "$HOME_DIR"

# **Generate the genesis transaction to stake tokens**
echo "[$NODE_NAME] Generating gentx..."
regen gentx my_validator 50000000uregen --keyring-backend test --chain-id regen-devnet --home "$HOME_DIR"

# **Copy the gentx to the shared directory**
echo "[$NODE_NAME] Copying gentx to shared directory..."
cp "$HOME_DIR/config/gentx/"*.json "$SHARED_DIR/gentxs/${NODE_NAME}_gentx.json"

# **Wait for all gentx files from all nodes**
for NODE in "${NODE_NAMES[@]}"; do
  while [ ! -f "$SHARED_DIR/gentxs/${NODE}_gentx.json" ]; do
    echo "[$NODE_NAME] Waiting for gentx file from $NODE..."
    sleep 2
  done
done

# **Regen-node1 collects all gentx files and creates the final genesis.json**
COMPLETION_FILE_GENTX="$SHARED_DIR/regen-node1_gentx_collect_done"
if [ "$NODE_NAME" == "regen-node1" ]; then
  echo "[$NODE_NAME] Collecting all genesis transactions..."
  regen collect-gentxs --gentx-dir "$SHARED_DIR/gentxs" --home "$HOME_DIR"

  echo "[$NODE_NAME] Validating the final genesis file..."
  regen validate-genesis --home "$HOME_DIR"

  cp "$HOME_DIR/config/genesis.json" "$SHARED_DIR/genesis.json"
  touch "$COMPLETION_FILE_GENTX"
else
  # **Other nodes wait for the final genesis.json**
  while [ ! -f "$COMPLETION_FILE_GENTX" ]; do
    echo "[$NODE_NAME] Waiting for regen-node1 to collect and finalize gentx..."
    sleep 2
  done

  cp "$SHARED_DIR/genesis.json" "$HOME_DIR/config/genesis.json"
  regen validate-genesis --home "$HOME_DIR"
fi

# **Retrieve Node ID and save it to shared directory**
NODE_ID=$(regen tendermint show-node-id --home "$HOME_DIR")
echo "[$NODE_NAME] My Node ID: $NODE_ID"
echo "$NODE_ID" > "$SHARED_DIR/${NODE_NAME}_id"

# **Wait for all Node IDs**
for NODE in "${NODE_NAMES[@]}"; do
  while [ ! -f "$SHARED_DIR/${NODE}_id" ]; do
    echo "[$NODE_NAME] Waiting for Node ID from $NODE..."
    sleep 2
  done
done

# **Configure persistent peers**
PERSISTENT_PEERS=""
for NODE in "${NODE_NAMES[@]}"; do
  if [ "$NODE" != "$NODE_NAME" ]; then
    PEER_ID=$(cat "$SHARED_DIR/${NODE}_id")
    PERSISTENT_PEERS+="$PEER_ID@$NODE:26656,"
  fi
done

PERSISTENT_PEERS=${PERSISTENT_PEERS%,}
sed -i "s/^persistent_peers = .*/persistent_peers = \"$PERSISTENT_PEERS\"/" "$HOME_DIR/config/config.toml"
sed -i 's/^addr_book_strict *=.*/addr_book_strict = false/' "$HOME_DIR/config/config.toml"


# **Start the node**
echo "[$NODE_NAME] Starting the Regen node..."
exec regen start --home "$HOME_DIR" --minimum-gas-prices="0.025uregen"
