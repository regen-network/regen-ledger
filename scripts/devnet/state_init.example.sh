#!/bin/bash
set -e

# Constants
CHAIN_ID="regen-devnet"
HOME_DIR="./.regen"  # Use the current directory for state initialization

# Colors and Emojis
GREEN='\033[0;32m'
NC='\033[0m'
INFO="${GREEN}ℹ️${NC}"
SUCCESS="${GREEN}✅${NC}"
WAIT="${GREEN}⏳${NC}"

# Helper function to log messages
log() { echo -e "${1} ${2}"; }

# Load validator addresses and keys from .env
log "$INFO" "Loading environment variables..."
source .env

# Ensure required addresses are available
VALIDATOR_1=${REGEN_NODE1_VALIDATOR_ADDRESS}
VALIDATOR_2=${REGEN_NODE2_VALIDATOR_ADDRESS}
MNEMONIC_1=${REGEN_NODE1_VALIDATOR_MNEMONIC}
MNEMONIC_2=${REGEN_NODE2_VALIDATOR_MNEMONIC}

# Function to import validator accounts if they don't exist
import_key_if_not_exists() {
  local name=$1
  local mnemonic=$2

  if regen keys show "$name" --keyring-backend test --home "$HOME_DIR" &> /dev/null; then
    log "$SUCCESS" "Key '$name' already exists in the keyring. Skipping import."
  else
    log "$INFO" "Importing key for '$name' into the keyring..."
    echo "$mnemonic" | regen keys add "$name" --recover --keyring-backend test --home "$HOME_DIR" --output json
    log "$SUCCESS" "Key '$name' imported successfully."
  fi
}

# Import validator accounts
log "$INFO" "Importing validator accounts into the keyring..."
import_key_if_not_exists "validator1" "$MNEMONIC_1"
import_key_if_not_exists "validator2" "$MNEMONIC_2"

# Check if RPC is available
RPC_URL="http://$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' regen-node1):26657"

log "$INFO" "Checking if RPC is available at $RPC_URL..."

check_rpc_available() {
  local retries=5
  while [ $retries -gt 0 ]; do
    if curl -s "$RPC_URL" > /dev/null; then
      log "$SUCCESS" "RPC endpoint is available."
      return 0
    else
      log "$WAIT" "Waiting for RPC to be available... ($retries retries left)"
      sleep 5
      retries=$((retries - 1))
    fi
  done

  log "$WAIT" "RPC is still unavailable after multiple attempts."
  exit 1
}

check_rpc_available

# 1. Submit a governance proposal
log "$INFO" "Submitting a governance proposal..."
regen tx gov submit-proposal text "Upgrade Proposal" \
  --description "Proposal to upgrade the testnet" \
  --deposit 10000000uregen --from validator1 \
  --chain-id "$CHAIN_ID" --home "$HOME_DIR" --keyring-backend test --yes

# 2. Vote on the proposal
log "$INFO" "Validator 2 voting YES on the proposal..."
regen tx gov vote 1 yes --from validator2 \
  --chain-id "$CHAIN_ID" --home "$HOME_DIR" --keyring-backend test --yes

log "$SUCCESS" "State initialization complete."
