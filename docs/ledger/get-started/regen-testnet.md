# Regen Testnet

This document provides information about available node endpoints for Regen Testnet and how to interact with a node using the `regen` binary. For general information about Regen Testnet such as supported block explorers and wallets, see [Regen Ledger Overview](../README.md).

## Prerequisites

- [Install Regen](./README.md)
- [Manage Keys](./manage-keys.md)

## Node Endpoints

The following RPC endpoint is for the full node operated by Vitwit:

- [https://rpc-regen-upgrade.vitwit.com/](https://rpc-regen-upgrade.vitwit.com/)

The following API endpoint is also available:

- [https://api-regen-upgrade.vitwit.com/](https://api-regen-upgrade.vitwit.com/)

## Testnet Tokens

The following token faucets currently supports Regen Testnet:

- [Vitwit](https://faucet-regen-upgrade.vitwit.com/)
- [Regen](https://app.regen.network/faucet)

## Configuration

To interact with a node on Regen Testnet, we first need to make sure we have `chain-id` and `node` properly configured for the `regen` binary.

To view your current configuration, run the following command:

```bash
regen config
```

The above command displays the configuration in the `~/.regen/config/client.toml` file. This file can be updated using the same `config` command along with key-value pairs (see below).

#### Chain ID

The chain ID for Regen Testnet is `regen-upgrade`. When using the `regen` binary to communicate with a node on Regen Testnet, you need to update the `chain-id` in your configuration.

To configure the chain ID for all commands, run the following:

```sh
regen config chain-id regen-upgrade
```

#### Node Endpoint

When interacting with a live network, you need to connect to a remote node or have a node running locally that is in sync with the network (in this case, Regen Testnet).

To configure the node endpoint for all commands, run the following:

```sh
regen config node https://rpc-regen-upgrade.vitwit.com:443
```
