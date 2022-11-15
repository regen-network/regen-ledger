#!/usr/bin/env bash

set -eo pipefail

REGEN_CHAIN_ID="regen-sandbox"

regen() {
  $BINARY --home $REGENHOME "$@"
}

# parse flags
POSITIONAL_ARGS=()
while [[ $# -gt 0 ]]; do
  case $1 in
    -o|--overwrite-home-dir)
      OVERWRITE_HOMEDIR=true
      shift # past argument
      ;;
    -*|--*)
      echo "Unknown option $1"
      exit 1
      ;;
    *)
      POSITIONAL_ARGS+=("$1") # save positional arg
      shift # past argument
      ;;
  esac
done

set -- "${POSITIONAL_ARGS[@]}" # restore positional parameters


if [ -d "$REGENHOME" ] && [ "$OVERWRITE_HOMEDIR" != true ]; then
  echo "Regen home ($REGENHOME) already exists, skipping bootstrap..."
else
  rm -rf $REGENHOME
  regen config keyring-backend test
  regen config node http://localhost:26657
  regen config chain-id $REGEN_CHAIN_ID
  regen config broadcast-mode block
  regen config output json

  # initialize .regen home directory and genesis.json
  regen init test_moniker --chain-id $REGEN_CHAIN_ID 2>&1 | jq -Rr '. as $raw | try (fromjson | "Created genesis. chain-id: \(.chain_id), moniker: \(.moniker)") catch $raw'

  # modify genesis file and config files (platform dependent usage of sed)
  if [[ $(uname -s) == 'Darwin' ]]; then
    # change stake denom to uregen
    sed -i "" "s/stake/uregen/g" $REGENHOME/config/genesis.json
    # set min gas price
    sed -i "" "s/minimum-gas-prices = \"\"/minimum-gas-prices = \"0.025uregen\"/g" $REGENHOME/config/app.toml
    # decrease block-time so tests run faster
    sed -i "" "s/timeout_commit = \"5s\"/timeout_commit = \"500ms\"/g" $REGENHOME/config/config.toml
    # bind on all interfaces, enabling ports to be exposed outside docker
    sed -i "" "s/127\.0\.0\.1/0.0.0.0/g" $REGENHOME/config/config.toml
  else
    # change stake denom to uregen
    sed -i "s/stake/uregen/g" $REGENHOME/config/genesis.json
    # set min gas price
    sed -i "s/minimum-gas-prices = \"\"/minimum-gas-prices = \"0.025uregen\"/g" $REGENHOME/config/app.toml
    # decrease block-time so tests run faster
    sed -i "s/timeout_commit = \"5s\"/timeout_commit = \"500ms\"/g" $REGENHOME/config/config.toml
    # bind on all interfaces, enabling ports to be exposed outside docker
    sed -i "s/127\.0\.0\.1/0.0.0.0/g" $REGENHOME/config/config.toml
  fi

  # app specific genesis file modifications
  cat .regen/config/genesis.json | jq '.app_state.ecocredit."regen.ecocredit.v1.AllowedBridgeChain"[0] = {"chain_name": "polygon"}' > genesis.json.tmp && mv genesis.json.tmp .regen/config/genesis.json

  # setup initial wallets, and validator address with gentx
  if ! [ -z "$REGEN_MNEMONIC" ]; then
    echo "Adding key to keyring for account name: addr1"
    echo $REGEN_MNEMONIC | regen keys add addr1 --account 0 --recover > /dev/null
  else
    REGEN_MNEMONIC=$(regen keys add addr1 | jq -r '.mnemonic')
    echo ""
    echo "No \$REGEN_MNEMONIC provided, using generated mnemonic:"
    echo "    $REGEN_MNEMONIC"
    echo ""
    echo "Adding key to keyring for account name: addr1"
  fi

  echo "Adding key to keyring for account name: addr2"
  echo $REGEN_MNEMONIC | regen keys add addr2 --account 1 --recover > /dev/null
  regen add-genesis-account addr2 10000000000uregen --keyring-backend test

  echo "Adding key to keyring for account name: addr3"
  echo $REGEN_MNEMONIC | regen keys add addr3 --account 2 --recover > /dev/null
  regen add-genesis-account addr3 10000000000uregen --keyring-backend test

  echo "Adding key to keyring for account name: addr4"
  echo $REGEN_MNEMONIC | regen keys add addr4 --account 3 --recover > /dev/null
  regen add-genesis-account addr4 10000000000uregen --keyring-backend test

  echo "Adding key to keyring for account name: addr5"
  echo $REGEN_MNEMONIC | regen keys add addr5 --account 4 --recover > /dev/null
  regen add-genesis-account addr5 10000000000uregen --keyring-backend test

  echo "Setting up validator (from addr1)..."
  regen add-genesis-account addr1 50000000000uregen --keyring-backend test
  regen gentx addr1 40000000000uregen --ip 127.0.0.1

  regen collect-gentxs 2>&1 | jq -r '"Collecting gentxs from \(.gentxs_dir) into genesis.json"'
fi

# Make sure we kill the regen process if our script exits while node is running
# in the background
trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM EXIT INT

# Start regen chain and immediately move it to the background
echo "Starting up regen node..."
regen start --log_level warn &
REGEN_PID=$!

# Sleep for regen node to full boot up
sleep 3

# run initialization of suites
SUITE_NAMES=$1
for suite in ${SUITE_NAMES//,/ }
do
  echo ""
  echo "INFO: Initializing state from './setup/$suite.sh'"
  ./setup/$suite.sh
done

# wait again on the regen node process so it can be terminated with ctrl+C
echo "Node started & state inialized!"
wait $REGEN_PID

