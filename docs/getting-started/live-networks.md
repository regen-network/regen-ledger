# Live Networks

This document provides an overview of live networks currently running the `regen-ledger` blockchain application and how to interact with a live network using the `regen` binary.

By default, the `regen` binary will connect to a local node. In order to connect to a live network, you'll need to know the address of a public node.

## Available Networks

### Regen Mainnet

`regen-1` is the chain ID for Regen Mainnet.

<!-- TODO: add information about genesis binary and upgrade binaries -->

Regen Mainnet launched with the `v1.0.0` release tag of `regen-ledger`.

<!-- TODO: update to use dedicated full node operated by RND -->

The following URL is the node endpoint for one of our lead validators:

[http://104.131.169.70:26657/](http://104.131.169.70:26657/)

For more information, see [regen-network/mainnet](https://github.com/regen-network/mainnet).

### Regen Devnet

`regen-devnet-5` is the chain ID for Regen Devnet.

<!-- TODO: add information about genesis binary and upgrade binaries -->

Regen Devnet launched with the `v1.0.0-rc0` release tag of `regen-ledger`.

<!-- TODO: update to use dedicated full node operated by RND -->

The following URL is the node endpoint for one of our lead validators:

[http://18.220.101.192:26657/](http://18.220.101.192:26657/)

For more information, see [regen-network/testnets](https://github.com/regen-network/testnets).

## Interacting With A Node

First, you'll need to install the `regen` binary. For installation instructions, see [Install Regen](./#install-regen).

In order to interact with a node from a live network, you'll need to provide a `--node` flag with a valid node address to your commands. Before submitting any queries or transactions, you should first check the status of the node using the `status` command.

To check the status of the Regen Mainnet node provided above, run the following command:
```
regen status --node http://104.131.169.70:26657
```

To check the status of the Regen Devnet node provided above, run the following command:
```
regen status --node http://18.220.101.192:26657
```

<!-- TODO: add `regen config node` instructions once updated to v1.1 -->

For more information about what commands are available, see [Command-Line Interface](http://localhost:8080/api.html#command-line-interface).