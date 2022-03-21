# Live Networks

This document provides an overview of live networks currently running the `regen-ledger` blockchain application and how to interact with a live network using the `regen` binary.

By default, the `regen` binary will connect to a local node. In order to connect to a live network, you'll need to know the address of a public node.

## Available Networks

### Regen Mainnet

`regen-1` is the chain ID for Regen Mainnet.

Regen Mainnet launched with the `v1.0.0` release tag of `regen-ledger`.

When starting a full node or a validator node from genesis, you will need to start the node with the `v1.0.0` version (the "genesis binary"). For more information about preparing your node to migrate to the latest version, see [Upgrade Overview](../migrations/upgrade.md).

The following URL is the node address for a full node operated by VitWit:

- [http://mainnet.regen.network:26657/](http://mainnet.regen.network:26657/)
- [http://regen.rpc.vitwit.com:26657/](http://regen.rpc.vitwit.com:26657/)

For more information, see [regen-network/mainnet](https://github.com/regen-network/mainnet).

### Redwood Testnet

`regen-redwood-1` is the chain ID for Redwood Testnet.

Redwood Testnet launched with the `v1.0.0` release tag of `regen-ledger`.

When starting a full node or a validator node from genesis, you will need to start the node with the `v1.0.0` version (the "genesis binary"). For more information about preparing your node to migrate to the latest version, see [Upgrade Overview](../migrations/upgrade.md).

The following URLs are node addresses for full nodes operated by RND and VitWit:

- [http://redwood.regen.network:26657/](http://redwood.regen.network:26657/)
- [http://redwood-sentry.vitwit.com:26657/](http://redwood-sentry.vitwit.com:26657/)

For more information, see [regen-network/testnets](https://github.com/regen-network/testnets).

### Hambach Testnet

`regen-hambach-1` is the chain ID for Hambach Testnet.

Hambach Testnet launched with the `v2.0.0-beta1` release tag of `regen-ledger` using the experimental build (`EXPERIMENTAL=true`).

When the time comes to introduce new experimental features, Hambach Testnet will be restarted with an incremented chain ID rather than going through the upgrade process due to the inability to support migrations for experimental modules. It's important to keep this in mind when developing with Hambach Testnet and we recommend creating scripts that can be used to reseed the network following a restart.

The following URLs are node addresses for full nodes operated by RND and VitWit:

- [http://hambach.regen.network:26657/](http://hambach.regen.network:26657/)
- [http://hambach-sentry.vitwit.com:26657/](http://hambach-sentry.vitwit.com:26657/)

For more information, see [regen-network/testnets](https://github.com/regen-network/testnets).

## Interacting With A Node

First, you'll need to install the `regen` binary. For installation instructions, see [Install Regen](./README.md#install-regen).

In order to interact with a node from a live network, you'll need to provide a `--node` flag with a valid node address to your commands. Before submitting any queries or transactions, you should first check the status of the node using the `status` command.

To check the status of the Regen Mainnet node, run the following command:

```bash
regen status --node http://mainnet.regen.network:26657
```

To check the status of the Redwood Testnet node, run the following command:

```bash
regen status --node http://redwood.regen.network:26657
```

To check the status of the Hambach Testnet node, run the following command:

```bash
regen status --node http://hambach.regen.network:26657
```

If you are using `v2.0.0` or later, you can set the node address once using the `config` command rather than adding the `--node` flag to each command.

```bash
regen config node http://mainnet.regen.network:26657
```

You can check the configuration by running the following command:

```bash
regen config
```

For more information about using the CLI, see [Command-Line Interface](../regen-ledger/interfaces.md#command-line-interface).

## Testnet Tokens

In order to interact with the test networks, you'll need some tokens. You can redeem tokens for each test network by executing the following `curl` commands.

*For Redwood Testnet:*

```bash
curl http://redwood-sentry.vitwit.com:8000/faucet/<account_address>
```

*For Hambach Testnet:*

```bash
curl http://hambach-sentry.vitwit.com:8000/faucet/<account_address>
```