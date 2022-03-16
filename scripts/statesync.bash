#!/bin/bash
# This is the Regen Network State sync file, which is based on the Gaia State sync file, which is based on scripts written by Bitcanna and Microtick. The key difference is that this approach uses environment variables to override configuration files.
# http://public-rpc.regen.vitwit.com:26657
# https://regen.stakesystems.io:2053
# https://rpc.regen.forbole.com:443



set -uxe

# set environment variables
export GOPATH=~/go
export PATH=$PATH:~/go/bin
export HOME_DIR=~/.regen

MONIKER=$1
if [ -z $MONIKER ]
then
    MONIKER=test-sync
fi

# MAKE HOME FOLDER AND GET GENESIS
regen init $MONIKER --home $HOME_DIR
wget https://raw.githubusercontent.com/regen-network/mainnet/main/regen-1/genesis.json -O $HOME_DIR/config/genesis.json 

INTERVAL=10

# GET TRUST HASH AND TRUST HEIGHT

LATEST_HEIGHT=$(curl -s http://public-rpc.regen.vitwit.com:26657/block | jq -r .result.block.header.height);
BLOCK_HEIGHT=$(($LATEST_HEIGHT-$INTERVAL))
TRUST_HASH=$(curl -s "http://public-rpc.regen.vitwit.com:26657/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)


# TELL USER WHAT WE ARE DOING
echo "TRUST HEIGHT: $BLOCK_HEIGHT"
echo "TRUST HASH: $TRUST_HASH"


# expor state sync vars
export REGEN_STATESYNC_ENABLE=true
export REGEN_P2P_MAX_NUM_OUTBOUND_PEERS=200
export REGEN_P2P_MAX_NUM_INBOUND_PEERS=200
export REGEN_STATESYNC_RPC_SERVERS="http://public-rpc.regen.vitwit.com:26657,https://rpc.regen.forbole.com:443,https://regen.stakesystems.io:2053"
export REGEN_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export REGEN_STATESYNC_TRUST_HASH=$TRUST_HASH
export REGEN_P2P_SEEDS=$(curl -s https://raw.githubusercontent.com/regen-network/mainnet/main/regen-1/seed-nodes.txt | paste -sd,)
export REGEN_P2P_PERSISTENT_PEERS=$(curl -s https://raw.githubusercontent.com/regen-network/mainnet/main/regen-1/peer-nodes.txt | paste -sd,)

sed -i '/persistent_peers =/c\persistent_peers = "'"$REGEN_P2P_PERSISTENT_PEERS"'"' $HOME_DIR/config/config.toml
sed -i '/seeds =/c\seeds = "'"$REGEN_P2P_SEEDS"'"' $HOME_DIR/config/config.toml
sed -i '/max_num_outbound_peers =/c\max_num_outbound_peers = '$REGEN_P2P_MAX_NUM_OUTBOUND_PEERS'' $HOME_DIR/config/config.toml
sed -i '/max_num_inbound_peers =/c\max_num_inbound_peers = '$REGEN_P2P_MAX_NUM_INBOUND_PEERS'' $HOME_DIR/config/config.toml
sed -i '/enable =/c\enable = true' $HOME_DIR/config/config.toml
sed -i '/rpc_servers =/c\rpc_servers = "'"$REGEN_STATESYNC_RPC_SERVERS"'"' $HOME_DIR/config/config.toml
sed -i '/trust_height =/c\trust_height = '$REGEN_STATESYNC_TRUST_HEIGHT'' $HOME_DIR/config/config.toml
sed -i '/trust_hash =/c\trust_hash = "'"$REGEN_STATESYNC_TRUST_HASH"'"' $HOME_DIR/config/config.toml

regen start --x-crisis-skip-assert-invariants
