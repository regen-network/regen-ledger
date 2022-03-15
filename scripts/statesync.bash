#!/bin/bash
# This is the Regen Network State sync file, which is based on the Gaia State sync file, which is based on scripts written by Bitcanna and Microtick. The key difference is that this approach uses environment variables to override configuration files.
# http://public-rpc.regen.vitwit.com:26657
# https://regen.stakesystems.io:2053
# http://rpc.regen.forbole.com:80



set -uxe

# set environment variables
export GOPATH=~/go
export PATH=$PATH:~/go/bin


# MAKE HOME FOLDER AND GET GENESIS
regen init statesync
wget -O ~/.regen/config/genesis.json https://cloudflare-ipfs.com/ipfs/QmdwHTcBcrowCpiFhuSnf5C3jxRMH8jBRqp7jimbHWLZFs

INTERVAL=1000

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
export REGEN_STATESYNC_RPC_SERVERS="http://public-rpc.regen.vitwit.com:26657,http://rpc.regen.forbole.com:80,https://regen.stakesystems.io:2053"
export REGEN_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export REGEN_STATESYNC_TRUST_HASH=$TRUST_HASH
export REGEN_P2P_SEEDS="aebb8431609cb126a977592446f5de252d8b7fa1@104.236.201.138:26656"
export REGEN_P2P_PERSISTENT_PEERS="69975e7afdf731a165e40449fcffc75167a084fc@104.131.169.70:26656,d35d652b6cb3bf7d6cb8d4bd7c036ea03e7be2ab@116.203.182.185:26656,ffacd3202ded6945fed12fa4fd715b1874985b8c@3.98.38.91:26656"

regen start --x-crisis-skip-assert-invariants
