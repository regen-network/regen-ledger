# Run a Full Node

This document provides instructions for running a full node for a [live network](../../ledger/get-started/live-networks.md) (either Regen Mainnet, Redwood Testnet, or Hambach Testnet).

The following instructions assume that you have already completed the following:

- [Initial Setup](README)
- [Install Regen](install-regen.md)
- [Install Cosmovisor](install-cosmovisor.md) (optional)

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

## Prepare Upgrade

The next step will be to [create a validator](create-a-validator.md) and prepare your node for the [upgrade process](../migrations/upgrade.md).
