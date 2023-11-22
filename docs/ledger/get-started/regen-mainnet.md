# Regen Mainnet

This document provides information about available node endpoints for Regen Mainnet and how to interact with a node using the `regen` binary. For general information about Regen Mainnet such as supported block explorers and wallets, see [Regen Ledger Overview](../README.md).

## Prerequisites

- [Install Regen](./README.md)
- [Manage Keys](./manage-keys.md)

## Node Endpoints

The following RPC endpoints are for full nodes operated by RND and VitWit:

- [http://mainnet.regen.network:26657/](http://mainnet.regen.network:26657/)
- [http://regen.rpc.vitwit.com:26657/](http://regen.rpc.vitwit.com:26657/)

The following RPC endpoint is for an archive node operated by RND:

- [http://archive.regen.network:26657/](http://archive.regen.network:26657/)

## Configuration

To interact with a node on Regen Mainnet, we first need to make sure we have `chain-id` and `node` properly configured for the `regen` binary.

To view your current configuration, run the following command:

```bash
regen config
```

The above command displays the configuration in the `~/.regen/config/client.toml` file. This file can be updated using the same `config` command along with key-value pairs (see below).

#### Chain ID

The chain ID for Regen Mainnet is `regen-1`. When using the `regen` binary to communicate with a node on Regen Mainnet, you need to update the `chain-id` in your configuration.

To configure the chain ID for all commands, run the following:

```sh
regen config chain-id regen-1
```

#### Node Endpoint

When interacting with a live network, you need to connect to a remote node or have a node running locally that is in sync with the network (in this case, Regen Mainnet).

To configure the node endpoint for all commands, run the following:

```sh
regen config node http://mainnet.regen.network:26657/
```
