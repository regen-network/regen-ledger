#!/bin/bash

# Set environment variables
export coins="10000000000stake,100000000000samoleans"
export coinsV="5000000000stake"
export REGEN_KEYRING_BACKEND=test
export CHAINID=regen-local
export TMCFG=~/.regen/config/config.toml
export APPCFG=~/.regen/config/app.toml

# Remove existing regen configuration
rm -rf ~/.regen

# Add a new key for the validator
./build/regen keys add validator --keyring-backend test

# Initialize the chain
./build/regen init dev-val --chain-id $CHAINID

# Add genesis account
./build/regen genesis add-genesis-account validator --keyring-backend test $coins

# Generate a genesis transaction
./build/regen genesis gentx validator $coinsV --chain-id $CHAINID --keyring-backend test

# Collect genesis transactions
./build/regen genesis collect-gentxs

# Validate genesis file
./build/regen genesis validate-genesis

# Modify configuration files
perl -i -pe 's|timeout_commit = ".*?"|timeout_commit = "2s"|g' $TMCFG
perl -i -pe 's|minimum-gas-prices = ""|minimum-gas-prices = "0stake"|g' $APPCFG

# Start the node
./build/regen start --api.enable true --grpc.address="0.0.0.0:9090"
