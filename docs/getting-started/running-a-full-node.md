# Running a Full Node

This document provides instructions for running a full node for a live network.

## Prerequisites

In order to install the `cosmovisor` and `regen` binaries, you'll need the following: 

- Git `>=2`
- Make `>=4`
- Go `>=1.15`

In order to run a full node for a live network, we recommend the following:

- 8GB RAM
- 4vCPUs
- 200GB Disk space

For more information on hardware requirements, see [Prerequisites](./prerequisites). 

## Install Cosmovisor

...

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

Download the genesis file for either `regen-1` or `regen-devnet-5`.

For `regen-1`:
```
wget https://raw.githubusercontent.com/regen-network/mainnet/main/regen-1/genesis.json
```

For `regen-devnet-5`:
```
...
```

Move `genesis.json` to node config:
```
mv genesis.json ~/.regen/config/genesis.json
```

## Update Genesis

```sh
curl http://18.220.101.192:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

## Update Peers

Update `~/.regen/config/config.toml`:
```
persistent_peers = "69975e7afdf731a165e40449fcffc75167a084fc@104.131.169.70:26656,d35d652b6cb3bf7d6cb8d4bd7c036ea03e7be2ab@116.203.182.185:26656,ffacd3202ded6945fed12fa4fd715b1874985b8c@3.98.38.91:26656"
```

## Start Node

Start node:
```
regen start
```
