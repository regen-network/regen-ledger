# Running a Validator

This document provides instructions for running a validator node for a [live network](./live-networks.md). With Regen Mainnet, Redwood Testnet, and Hambach Testnet already launched and running, this document will focus on how to become a validator for a network post-genesis.

## Prerequisites

In order to install the `cosmovisor` and `regen` binaries, you'll need the following: 

- Git `>=2`
- Make `>=4`
- Go `>=1.17`

For more information (including hardware recommendations), see [Prerequisites](./prerequisites.md). 

## Install Regen

Clone the `regen-ledger` repository:
```
git clone https://github.com/regen-network/regen-ledger
```

Change to the `regen-ledger` directory:
```
cd regen-ledger
```

Check out the version that the network launched with.

*For Regen Mainnet:*
```
git checkout v1.0.0
```

*For Redwood Testnet:*
```
git checkout v1.0.0
```

*For Hambach Testnet:*
```
git checkout v2.0.0-beta1
```

Install the `regen` binary (the `EXPERIMENTAL` option enables experimental features).

*For Regen Mainnet:*
```
make install
```

*For Redwood Testnet:*
```
make install
```

*For Hambach Testnet:*
```
EXPERIMENTAL=true make install
```

Check to ensure the install was successful:
```
regen version
```

## Initialize Node

Create the configuration files and data directory by initializing the node. In the following command, replace `[moniker]` with a name of your choice. 

*For Regen Mainnet:*
```
regen init [moniker] --chain-id regen-1
```

*For Redwood Testnet:*
```
regen init [moniker] --chain-id regen-redwood-1
```

*For Hambach Testnet:*
```
regen init [moniker] --chain-id regen-hambach-1
```

## Update Genesis

Update the genesis file using a node endpoint.

<!-- TODO: update to use dedicated full node operated by RND -->

*For Regen Mainnet:*
```
curl http://104.131.169.70:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

*For Redwood Testnet:*
```
curl http://redwood.regen.network:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

*For Hambach Testnet:*
```
curl http://hambach.regen.network:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

## Update Peers

Add a seed node for initial peer discovery.

<!-- TODO: update to use dedicated full node operated by RND -->

*For Regen Mainnet:*
```
PERSISTENT_PEERS="69975e7afdf731a165e40449fcffc75167a084fc@104.131.169.70:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

*For Redwood Testnet:*
```
PERSISTENT_PEERS="a5528d8f5fabd3d50e91e8d6a97e355403c5b842@redwood.regen.network:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

*For Hambach Testnet:*
```
PERSISTENT_PEERS="4f5c0be7705bf4acb5b99dcaf93190059ac283a1@hambach.regen.network:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

## Install Cosmovisor

[Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor) is a process manager for running application binaries. Using Cosmovisor is not required but recommended for node operators that would like to automate the upgrade process.

To install `cosmovisor`, run the following command:
```
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0
```

Check to ensure the install was successful:
```
cosmovisor version
```

## Set Genesis Binary

Create the folder for the genesis binary and copy the `regen` binary:
```
mkdir -p $HOME/.regen/cosmovisor/genesis/bin
cp $GOBIN/regen $HOME/.regen/cosmovisor/genesis/bin
```

## Cosmovisor Service

The next step will be to configure `cosmovisor` as a `systemd` service. For more information about the environment variables used to configure `cosmovisor`, see [Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor).

::: warning
You'll want to carefully consider the options you set when configuring cosmovisor. The current version of cosmovisor does not require the checksum parameter to be included in the URL of the downloadable upgrade binary, so the auto-download option should be used with caution.
:::

Create the `cosmovisor.service` file:
```
echo "[Unit]
Description=Cosmovisor daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=regen"
Environment="DAEMON_HOME=${HOME}/.regen"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
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
sudo systemctl daemon-reload
sudo systemctl start cosmovisor
```

Check the status of the `cosmovisor` service:
```
sudo systemctl status cosmovisor
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

## Create Validator

The next step will be to create a validator. You will need to have enough REGEN tokens to stake and to submit the transaction. For more information about the REGEN token, see the [token page](https://www.regen.network/token/). 

::: warning
You'll want to carefully consider the options you set when creating a validator.
:::

Submit a transaction to create a validator:

```
regen tx staking create-validator \
  --amount=<stake_amount> \
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
