# Running a Full Node

This document provides instructions for running a full node for a [live network](./live-networks.html) (either Regen Mainnet or Regen Devnet).

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

<!-- TODO: update to use dedicated full node operated by RND -->

*For Regen Mainnet:*
```
curl http://104.131.169.70:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

<!-- TODO: update to use dedicated full node operated by RND -->

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

At this point, the node is ready. If you do not need to run a dedicated full node in a separate process, you can start the node using the `regen` binary.

Start node:
```
regen start
```

## Install Cosmovisor

[Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor) is a process manager for running application binaries. Using Cosmovisor is not required but recommended for node operators that would like to automate the upgrade process.

To install `cosmovisor`, run the following command:
```
go get github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor
```

## Set Genesis Binary

Create the folder for the genesis binary and copy the `regen` binary:
```
mkdir -p $HOME/.regen/cosmovisor/genesis/bin
cp ./build/regen $HOME/.regen/cosmovisor/genesis/bin
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
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=true"
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
