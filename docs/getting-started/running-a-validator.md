# Running a Validator

This document provides instructions for running a validator node for a live network. With both Regen Mainnet and Regen Devnet already launched and running, this document will focus on how to become a validator post-genesis (after a chain has launched).

## Prerequisites

In order to install the `cosmovisor` and `regen` binaries, you'll need the following: 

- Git `>=2`
- Make `>=4`
- Go `>=1.15`

For more information (including hardware recommendations), see [Prerequisites](./prerequisites). 

## Install Cosmovisor

Cosmovisor is a process manager for running application binaries. Using Cosmovisor is not required but recommended for node operators that would like to perform automatic upgrades.

To install `cosmovisor`, run the following command:

```
go get github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor
```

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

As a validator who signs blocks, your node must have a public/private keypair. Regen Ledger keys can be managed with the `regen keys` subcommand. A new key pair can be generated using:

```
regen keys add [name]
```

If you create a new key, make sure you store the mnemonic phrase in a safe place.

If you'd like to use an existing key or a custom keyring backend, you can find more information in the Cosmos SDK [Keyring](https://docs.cosmos.network/master/run-node/keyring.html) documentation.

<!-- TODO: buying tokens and sending to validator account -->

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

Next, you will need to add seed nodes for initial peer discovery:

For Regen Mainnet (`regen-1`):
```
PERSISTENT_PEERS="a621e6bf1f5981b3e72e059f86cbfc9dc5577fcb@18.220.101.192:26656"
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' ~/.regen/config/config.toml
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

For Regen Devnet (`regen-devnet-5`):
```
```

## Cosmovisor Service

In the `cosmovisor` directory, create the systemd files.

```
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

Move the newly created file to the systemd directory, reload systemctl, and start `cosmovisor`:
```
sudo mv cosmovisor.service /lib/systemd/system/cosmovisor.service
sudo -S systemctl daemon-reload
sudo -S systemctl start cosmovisor
```

## Create Validator

<!-- TODO: buying tokens and sending to validator account -->

```
regen tx staking create-validator \
  --amount=9000000uregen \
  --pubkey=$(regen tendermint show-validator) \
  --moniker="<your_moniker>" \
  --chain-id=regen-1 \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --gas="auto" \
  --from=<your_wallet_name>
```