# Initialize Node

The following instructions assume that you have already completed the following:

- [Initial Setup](README)
- [Install Regen](install-regen.md)

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

## Start Node

The node is now ready to connect to the network:

```bash
regen start
```

## Create a Validator

The next step will be to [create a validator](create-a-validator.md).

## Using State Sync

Also, syncing from genesis will be a slow process. You may want to consider [using state sync](using-state-sync.md).
