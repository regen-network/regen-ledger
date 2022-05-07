# Run a Full Node

This document provides instructions for running a full node for a [live network](../../ledger/get-started/live-networks.md) (either Regen Mainnet, Redwood Testnet, or Hambach Testnet).

The following instructions assume that you have already completed the following:

- [Initial Setup](README)
- [Install Regen](install-regen.md)
- [Install Cosmovisor](install-cosmovisor.md)

## Quickstart

If you would like to manually set up a full node, skip to the [next section](#install-regen). Alternatively, you can run the following quickstart script:

```bash
bash <(curl -s https://raw.githubusercontent.com/regen-network/mainnet/blob/main/scripts/mainnet-val-setup.sh)
```

## Initialize Node

Create the configuration files and data directory by initializing the node. In the following command, replace `[moniker]` with a name of your choice. 

*For Regen Mainnet:*

```bash
regen init [moniker] --chain-id regen-1
```

*For Redwood Testnet:*

```bash
regen init [moniker] --chain-id regen-redwood-1
```

*For Hambach Testnet:*

```bash
regen init [moniker] --chain-id regen-hambach-1
```

## Update Genesis

Update the genesis file.

<!-- TODO: update to use dedicated full node operated by RND -->

*For Regen Mainnet:*

```bash
curl http://mainnet.regen.network:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

*For Redwood Testnet:*

```bash
curl http://redwood.regen.network:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

*For Hambach Testnet:*

```bash
curl http://hambach.regen.network:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

## Update Peers

Add a seed node for initial peer discovery.

*For Regen Mainnet:*

```bash
PERSISTENT_PEERS="c4460b52c34ad4f12168d05807e998bb8e8b4812@mainnet.regen.network:26656,aebb8431609cb126a977592446f5de252d8b7fa1@regen.rpc.vitwit.com:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

*For Redwood Testnet:*

```bash
PERSISTENT_PEERS="d5ceac343e48c7522c3a5a8c0cf5cb896d1f8a60@redwood.regen.network:26656,61f53f226a4a71968a87583f58902405e289b4b9@redwood-sentry.vitwit.com:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

*For Hambach Testnet:*

```bash
PERSISTENT_PEERS="4f5c0be7705bf4acb5b99dcaf93190059ac283a1@hambach.regen.network:26656,578b74c81f08a812b5f1a76a53b00a8ad3cfec57@hambach-sentry.vitwit.com:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

## Start Node

At this point, the node is ready. If you do not need to run a dedicated full node in a separate process, you can start the node using the `regen` binary.

Start node:

```bash
regen start
```

## Set Genesis Binary

Create the folder for the genesis binary and copy the `regen` binary:

```bash
mkdir -p $HOME/.regen/cosmovisor/genesis/bin
cp $GOBIN/regen $HOME/.regen/cosmovisor/genesis/bin
```

## Cosmovisor Service

The next step will be to configure `cosmovisor` as a `systemd` service. For more information about the environment variables used to configure `cosmovisor`, see [Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor).

Create the `cosmovisor.service` file:

```bash
echo "[Unit]
Description=Cosmovisor daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=regen"
Environment="DAEMON_HOME=${HOME}/.regen"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
Environment="UNSAFE_SKIP_BACKUP=false"
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

```bash
sudo mv cosmovisor.service /lib/systemd/system/cosmovisor.service
```

Reload systemctl and start `cosmovisor`:

```bash
sudo systemctl daemon-reload
sudo systemctl start cosmovisor
```

Check the status of the `cosmovisor` service:

```bash
sudo systemctl status cosmovisor
```

Enable cosmovisor to start automatically when the machine reboots:

```bash
sudo systemctl enable cosmovisor.service
```
## Using StateSync

[Regen Mainnet](../../ledger/get-started/live-networks.md#regen-mainnet) and [Redwood Testnet](../../ledger/get-started/live-networks.md#redwood-testnet) also support [statesync](https://docs.cosmos.network/v0.44/architecture/adr-040-storage-and-smt-state-commitments.html#snapshots-for-storage-sync-and-state-versioning) which allows node operators to quickly spin up a node without downloading the existing chain data. It should be noted that not many nodes should be spun up on the network using this method as these nodes will be unable to propogate the historical data to other nodes.

*Download and execute the script for Regen Mainnet*:
```bash 
export MONIKER=<your-node-moniker>
curl -s -L https://raw.githubusercontent.com/regen-network/regen-ledger/master/scripts/statesync.bash | bash -s $MONIKER
```

*Download and execute the script for Redwood Testnet*:
```bash 
export MONIKER=<your-node-moniker>
curl -s -L https://raw.githubusercontent.com/regen-network/testnets/master/scripts/testnet-statesync.bash | bash -s $MONIKER
```

## Prepare Upgrade

The next step will be to [create a validator](create-a-validator.md) and prepare your node for the [upgrade process](../migrations/upgrade.md).
