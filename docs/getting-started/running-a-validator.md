# Running a Validator

This document provides instructions for running a validator node for a live network.

## Prerequisites

In order to install the `cosmovisor` and `regen` binaries, you'll need the following: 

- Git `>=2`
- Make `>=4`
- Go `>=1.15`

In order to run a validator node for a live network, we recommend the following:

- 8GB RAM
- 4vCPUs
- 200GB Disk space

For more information, see [Prerequisites](./prerequisites). 

## Install Cosmovisor

...

## Build Regen Binary

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

## Creating Keys

As a validator who signs blocks, your node must have a public/private keypair. Regen Ledger keys can be managed with the `regen keys` subcommand. A new key pair can be generated using:

```sh
KEY_NAME=my_validator # Or choose your own key name.
regen keys add $KEY_NAME

# We will also save the generated address in a variable for later use.
MY_VALIDATOR_ADDRESS=$(regen keys show $KEY_NAME -a)
```

If you'd like to use a custom keyring backend, you can find more information on the [Cosmos SDK keyring docs](https://docs.cosmos.network/master/run-node/keyring.html).

**Make sure you save the seed mnemonic in a safe place!**

## Update Genesis

```sh
curl http://18.220.101.192:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

## Update Peers

Be sure to also add well-known seed nodes for your local node's initial peer discovery:

```sh
PERSISTENT_PEERS="a621e6bf1f5981b3e72e059f86cbfc9dc5577fcb@18.220.101.192:26656"
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' ~/.regen/config/config.toml
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

## Cosmovisor Service

In the `cosmovisor` directory create the systemd files. Copy the entire code block below and paste in your shell and hit enter

```sh
echo "[Unit]
Description=Cosmovisor daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=regen"
Environment="DAEMON_HOME=${HOME}/.${DAEMON}"
Environment="DAEMON_RESTART_AFTER_UPGRADE=on"
User=${USER}
ExecStart=${GOBIN}/cosmovisor start
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
" >cosmovisor.service
```

Move the newly create file to the systemd directory reload systemctl and start `cosmovisor`
```sh
sudo mv cosmovisor.service /lib/systemd/system/cosmovisor.service
```
```sh
sudo -S systemctl daemon-reload
```
```sh
sudo -S systemctl start cosmovisor
```
