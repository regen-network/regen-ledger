# Running a Validator

This document provides instructions for running a validator node for a live network. With both Regen Mainnet and Regen Devnet already launched and running, this document will focus on how to become a validator for a live network that has already launched.

## Prerequisites

In order to install the `cosmovisor` and `regen` binaries, you'll need the following: 

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

Check out the latest stable version:
```
git checkout v1.0.0
```

Install the `regen` binary:
```
make install
```

Check to make sure the install was successful:
```
regen version
```

## Add Validator Key

As a validator who signs blocks, your node must have a public/private key pair. Regen Ledger keys can be managed with the `regen keys` subcommand. A new key pair can be generated using:

```
regen keys add [name]
```

::: warning
If you create a new key, make sure you store the mnemonic phrase in a safe place. You will not be able to recover your new key without the mnemonic phrase.
:::

If you'd like to use an existing key or a custom keyring backend, you can find more information about adding keys and keyring backends in the Cosmos SDK [Keyring](https://docs.cosmos.network/master/run-node/keyring.html) documentation.

## Initialize Node

Create the configuration files and data directory by initializing the node.

For Regen Mainnet (`regen-1`):
```
regen init [moniker] --chain-id regen-1
```

For Regen Devnet (`regen-devnet-5`):
```
regen init [moniker] --chain-id regen-devnet-1
```

## Update Genesis

Update the genesis file using a node endpoint.

<!-- TODO: update to use dedicated full node operated by RND -->

For Regen Mainnet (`regen-1`):
```
curl http://104.131.169.70:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

<!-- TODO: update to use dedicated full node operated by RND -->

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

## Install Cosmovisor

[Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor) is a process manager for running application binaries. Using Cosmovisor is not required but recommended for node operators that would like to automate the upgrade process.

To install `cosmovisor`, run the following command:
```
go get github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor
```

## Cosmovisor Service

The next step will be to configure `cosmovisor` as a `systemd` service. For more information about the environment variables used to configure `cosmovisor`, see [Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor).

::: warning
You'll want to carefully consider the options you set when configuring cosmovisor.
:::

Create the `cosmovisor.service` file:
```
echo "[Unit]
Description=Cosmovisor daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=regen"
Environment="DAEMON_HOME=${HOME}/.${DAEMON}"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
User=${USER}
ExecStart=${GOBIN}/cosmovisor start
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
" >cosmovisor.service
```

Move the file to the systemd directory:
```
sudo mv cosmovisor.service /lib/systemd/system/cosmovisor.service
```

Reload systemctl and start `cosmovisor`:
```
sudo -S systemctl daemon-reload
sudo -S systemctl start cosmovisor
```

## Create Validator

The next step will be to create a validator. You will need to have enough REGEN tokens to stake and to submit the transaction. For more information about the REGEN token, see the [token page](https://www.regen.network/token/). 

::: warning
You'll want to carefully consider the options you set when creating a validator.
:::

Submit a transaction to create a validator:

```
regen tx staking create-validator \
  --amount=9000000uregen \
  --pubkey=$(regen tendermint show-validator) \
  --moniker="<node_moniker>" \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --gas="auto" \
  --from=<key_name>
```