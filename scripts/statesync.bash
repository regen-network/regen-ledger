#!/bin/bash
# microtick and bitcanna contributed significantly here.
# This is the Regen Network State sync file, which is based on the Gaia State sync file, which is based on scripts written bgy Bitcanna and Microtick.  The key difference is that this approach uses environment variables to override convifugration files.

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
export REGEN_STATESYNC_RPC_SERVERS="http://public-rpc.regen.vitwit.com:26657,http://public-rpc.regen.vitwit.com:26657"
export REGEN_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export REGEN_STATESYNC_TRUST_HASH=$TRUST_HASH
export REGEN_P2P_PERSISTENT_PEERS="69975e7afdf731a165e40449fcffc75167a084fc@104.131.169.70:26656,d35d652b6cb3bf7d6cb8d4bd7c036ea03e7be2ab@116.203.182.185:26656,ffacd3202ded6945fed12fa4fd715b1874985b8c@3.98.38.91:26656,d153ddc748f85f490ae7f6195137d6c39d53c89b@50.116.44.210:26656,81ae5a804e7f29d8dbad49c8734bfca47810569a@34.150.234.59:26656,95d4be08b7705efa59fd657efe65d06a75b232b5@161.97.93.76:26656,6791609444e95982ee2013199a523172d25054df@80.64.211.92:26656,f8d3cb29e4550ddf5ee633c4ec6bc232a487cb71@45.35.34.70:26656,92543347634034d6473075e3cee041cb58afa81d@162.55.132.230:2161"
export REGEN_P2P_SEEDS="aebb8431609cb126a977592446f5de252d8b7fa1@104.236.201.138:26656"

regen start --x-crisis-skip-assert-invariants
