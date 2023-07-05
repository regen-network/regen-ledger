# Redwood Testnet

This document provides information about available node endpoints for Redwood Testnet and how to interact with a node using the `regen` binary. For general information about Redwood Testnet such as supported block explorers and wallets, see [Regen Ledger Overview](../README.md).

## Prerequisites

- [Install Regen](./README.md)
- [Manage Keys](./manage-keys.md)

## Node Endpoints

The following RPC endpoints are for full nodes operated by RND and VitWit:

- [http://redwood.regen.network:26657/](http://redwood.regen.network:26657/)
- [http://redwood-sentry.vitwit.com:26657/](http://redwood-sentry.vitwit.com:26657/)

## Testnet Tokens

To interact with Redwood Testnet, you'll need some testnet tokens. You can redeem unofficial REGEN tokens using the following command:

```sh
curl -X POST -d '{"address": "YOUR_REGEN_ADDRESS"}' http://redwood.regen.network:8000
```

## Configuration

To interact with a node on Redwood Testnet, we first need to make sure we have `chain-id` and `node` properly configured for the `regen` binary.

To view your current configuration, run the following command:

```bash
regen config
```

The above command displays the configuration in the `~/.regen/config/client.toml` file. This file can be updated using the same `config` command along with key-value pairs (see below).

#### Chain ID

The chain ID for Redwood Testnet is `regen-redwood-1`. When using the `regen` binary to communicate with a node on Redwood Testnet, you need to update the `chain-id` in your configuration.

To configure the chain ID for all commands, run the following:

```sh
regen config chain-id regen-redwood-1
```

#### Node Endpoint

When interacting with a live network, you need to connect to a remote node or have a node running locally that is in sync with the network (in this case, Redwood Testnet).

To configure the node endpoint for all commands, run the following:

```sh
regen config node http://redwood.regen.network:26657/
```
