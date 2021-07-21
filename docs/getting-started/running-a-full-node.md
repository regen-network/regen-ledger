# Running a Full Node

This document provides instructions for running a full node for a [live network](./live-networks.html) (either Regen Mainnet or Regen Devnet).

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

<!-- TODO: add information about genesis binary and upgrade binaries -->

Check out the version that the network launched with.

*For Regen Mainnet:*
```
git checkout v1.0.0
```

*For Regen Devnet:*
```
git checkout v1.0.0-rc0
```

Install the `regen` binary (the `EXPERIMENTAL` option enables experimental features).

*For Regen Mainnet:*
```
make install
```

*For Regen Devnet:*
```
EXPERIMENTAL=true make install
```

Check to make sure the install was successful:
```
regen version
```

## Initialize Node

Create the configuration files and data directory by initializing the node. In the following command, replace `[moniker]` with a name of your choice. 

*For Regen Mainnet:*
```
regen init [moniker] --chain-id regen-1
```

*For Regen Devnet:*
```
regen init [moniker] --chain-id regen-devnet-5
```

## Update Genesis

Update the genesis file for either Regen Mainnet or Regen Devnet.

*For Regen Mainnet:*
```
curl http://104.131.169.70:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

*For Regen Devnet:*
```
curl http://18.220.101.192:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

## Update Peers

Add a seed node for initial peer discovery.

<!-- TODO: update to use dedicated full node operated by RND -->

*For Regen Mainnet:*
```
PERSISTENT_PEERS="69975e7afdf731a165e40449fcffc75167a084fc@104.131.169.70:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

<!-- TODO: update to use dedicated full node operated by RND -->

*For Regen Devnet:*
```
PERSISTENT_PEERS="b2679a74d6bd9f89a3c294c447d6930293255e6b@18.220.101.192:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

## Start Node

Start node:
```
regen start
```
