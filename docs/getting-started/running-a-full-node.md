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
regen init node
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

Update `~/.regen/config/config.toml` to include persistent peers.

For Regen Mainnet (`regen-1`):
```
persistent_peers = "69975e7afdf731a165e40449fcffc75167a084fc@104.131.169.70:26656,d35d652b6cb3bf7d6cb8d4bd7c036ea03e7be2ab@116.203.182.185:26656,ffacd3202ded6945fed12fa4fd715b1874985b8c@3.98.38.91:26656"
```

For Regen Devnet (`regen-devnet-5`):
```
```

## Start Node

Start node:
```
regen start
```
