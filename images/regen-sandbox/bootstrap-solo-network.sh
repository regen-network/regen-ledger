#!/bin/bash

set -eo pipefail

REGEN_CHAIN_ID="regen-sandbox"
REGEN="$BINARY --home $REGENHOME"

if [ -d "$REGENHOME" ]; then
  echo "Regen home ($REGENHOME) already exists, skipping bootstrap..."
else
  $REGEN config keyring-backend test
  $REGEN config node http://localhost:26657
  $REGEN config chain-id $REGEN_CHAIN_ID
  $REGEN config broadcast-mode block
  $REGEN config output json
  
  $REGEN init test_moniker --chain-id $REGEN_CHAIN_ID 2>&1 | jq -Rr '. as $raw | try (fromjson | "Created genesis. chain-id: \(.chain_id), moniker: \(.moniker)") catch $raw'
  
  if [[ $(uname -s) == 'Darwin' ]]; then
    # change stake denom to uregen
    sed -i "" "s/stake/uregen/g" $REGENHOME/config/genesis.json
    # set min gas price
    sed -i "" "s/minimum-gas-prices = \"\"/minimum-gas-prices = \"0.025uregen\"/g" $REGENHOME/config/app.toml
    # decrease block-time so tests run faster
    sed -i "" "s/timeout_commit = \"5s\"/timeout_commit = \"500ms\"/g" $REGENHOME/config/config.toml
  else
    # change stake denom to uregen
    sed -i "s/stake/uregen/g" $REGENHOME/config/genesis.json
    # set min gas price
    sed -i "s/minimum-gas-prices = \"\"/minimum-gas-prices = \"0.025uregen\"/g" $REGENHOME/config/app.toml
    # decrease block-time so tests run faster
    sed -i "s/timeout_commit = \"5s\"/timeout_commit = \"500ms\"/g" $REGENHOME/config/config.toml
  fi
  
  if ! [ -z "$REGEN_MNEMONIC" ]; then
    echo "Adding key to keyring for account name: addr1"
    echo $REGEN_MNEMONIC | $REGEN keys add addr1 --account 0 --recover > /dev/null
    $REGEN add-genesis-account addr1 50000000000uregen --keyring-backend test
    $REGEN gentx addr1 40000000000uregen --ip 127.0.0.1
  
    echo "Adding key to keyring for account name: addr2"
    echo $REGEN_MNEMONIC | $REGEN keys add addr2 --account 1 --recover > /dev/null
    $REGEN add-genesis-account addr2 10000000000uregen --keyring-backend test
  
    echo "Adding key to keyring for account name: addr3"
    echo $REGEN_MNEMONIC | $REGEN keys add addr3 --account 2 --recover > /dev/null
    $REGEN add-genesis-account addr3 10000000000uregen --keyring-backend test
  
    echo "Adding key to keyring for account name: addr4"
    echo $REGEN_MNEMONIC | $REGEN keys add addr4 --account 3 --recover > /dev/null
    $REGEN add-genesis-account addr4 10000000000uregen --keyring-backend test
  
    echo "Adding key to keyring for account name: addr5"
    echo $REGEN_MNEMONIC | $REGEN keys add addr5 --account 4 --recover > /dev/null
    $REGEN add-genesis-account addr5 10000000000uregen --keyring-backend test
  else
    echo "No \$REGEN_MNEMONIC provided."
    exit 1
  fi

  $REGEN collect-gentxs 2>&1 | jq -r '"Collecting gentxs from \(.gentxs_dir)"'
fi

# Make sure we kill the regen process if our script exits while node is running
# in the background
trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM EXIT INT

# Start regen chain and immediately move it to the background
echo "Starting up regen node..."
$REGEN start --log_level warn &
REGEN_PID=$!

# Sleep for regen node to full boot up
sleep 3

# run initialization of suites
SUITE_NAMES=$1
for suite in ${SUITE_NAMES//,/ }
do
  echo "INFO: Initializing state from './setup/$suite.sh'"
  ./setup/$suite.sh
done

# wait again on the regen node process so it can be terminated with ctrl+C
echo "Node started & state inialized!"
wait $REGEN_PID

