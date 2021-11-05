# Live Networks

This document provides an overview of live networks currently running the `regen-ledger` blockchain application and how to interact with a live network using the `regen` binary.

By default, the `regen` binary will connect to a local node. In order to connect to a live network, you'll need to know the address of a public node.

## Available Networks

### Regen Mainnet

`regen-1` is the chain ID for Regen Mainnet.

<!-- TODO: add information about genesis binary and upgrade binaries -->

Regen Mainnet launched with the `v1.0.0` release tag of `regen-ledger`.

<!-- TODO: update to use dedicated full node operated by RND -->

The following URL is the node address for a full node operated by VitWit:

[http://104.131.169.70:26657/](http://104.131.169.70:26657/)

For more information, see [regen-network/mainnet](https://github.com/regen-network/mainnet).

### Redwood Testnet

`regen-redwood-1` is the chain ID for Redwood Testnet.

<!-- TODO: add information about genesis binary and upgrade binaries -->

Redwood Testnet launched with the `v1.0.0` release tag of `regen-ledger`.

The following URL is the node address for a full node operated by RND:

[http://redwood.regen.network:26657/](http://redwood.regen.network:26657/)

For more information, see [regen-network/testnets](https://github.com/regen-network/testnets).

### Hambach Testnet

`regen-hambach-1` is the chain ID for Hambach Testnet.

<!-- TODO: add information about genesis binary and upgrade binaries -->

Hambach Testnet launched with the `v2.0.0-beta1` release tag of `regen-ledger` using the experimental build.

The following URL is the node address for a full node operated by RND:

[http://hambach.regen.network:26657/](http://hambach.regen.network:26657/)

For more information, see [regen-network/testnets](https://github.com/regen-network/testnets).

## Interacting With A Node

First, you'll need to install the `regen` binary. For installation instructions, see [Install Regen](./README.md#install-regen).

In order to interact with a node from a live network, you'll need to provide a `--node` flag with a valid node address to your commands. Before submitting any queries or transactions, you should first check the status of the node using the `status` command.

To check the status of the Regen Mainnet node provided above, run the following command:

```bash
regen status --node http://104.131.169.70:26657
```

To check the status of the Redwood Testnet node provided above, run the following command:

```bash
regen status --node http://redwood.regen.network:26657
```

To check the status of the Hambach Testnet node provided above, run the following command:

```bash
regen status --node http://hambach.regen.network:26657
```

<!-- TODO: add `regen config node` instructions once updated to v2.0 -->

For more information, see [Command-Line Interface](../regen-ledger/interfaces.md#command-line-interface).