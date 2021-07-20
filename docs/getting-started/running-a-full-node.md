# Running a Full Node

This document provides instructions for running a full node for a live network.

## Prerequisites

In order to install the `regen` binary, you'll need the following: 

- Git `>=2`
- Make `>=4`
- Go `>=1.15`

For more information (including hardware recommendations), see [Prerequisites](./prerequisites). 

## Install Regen

Clone the `regen-ledger` repository:
```
git clone https://github.com/regen-network/regen-ledger
```

Change to the `regen-ledger` directory:
```
cd regen-ledger
```

Check out `v1.0.0`:
```
git checkout v1.0.0
```

Install `regen` binary:
```
make install
```

Initialize node:
```
regen init [moniker]
```

## Update Genesis

Update the genesis file for either Regen Mainnet or Regen Devnet.

For Regen Mainnet (`regen-1`):
```
curl http://104.131.169.70:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

For Regen Devnet (`regen-devnet-5`):
```
curl http://18.220.101.192:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

## Update Peers

Add a seed node for initial peer discovery.

<!-- TODO: update to use dedicated full node operated by RND -->

For Regen Mainnet (`regen-1`):
```
PERSISTENT_PEERS="69975e7afdf731a165e40449fcffc75167a084fc@104.131.169.70:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

<!-- TODO: update to use dedicated full node operated by RND -->

For Regen Devnet (`regen-devnet-5`):
```
PERSISTENT_PEERS="a621e6bf1f5981b3e72e059f86cbfc9dc5577fcb@18.220.101.192:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

## Start Node

Start node:
```
regen start
```
