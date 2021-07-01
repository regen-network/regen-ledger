# Live Networks

This document provides an overview of live networks currently running the `regen-ledger` blockchain application and how to interact with a live network using the `regen` binary.

## Available Networks

### regen-1

`regen-1` is the chain ID for the main network.

http://104.131.169.70:26657/

http://104.131.169.70:26657/genesis

For more information, see [regen-network/mainnet](https://github.com/regen-network/mainnet).

### regen-devnet-5

`regen-devnet-5` is the chain ID for the development network.

http://18.220.101.192:26657/

http://18.220.101.192:26657/genesis

For more information, see [regen-network/testnets](https://github.com/regen-network/testnets).

## Test Commands

By default the command line client will connect to a [local node](#running-a-node). To connect to mainnet or a devnet, you must know the network's chain ID and the address of a public node.

...

In order to interact with a node from a live network, you will need to provide a `--node` flag with a valid peer address. See [the API documentation](../api.md) for more information.

...

Regardless of whether you're running a local node, or connecting directly to a live network, you can run `regen status` to verify that the CLI is able to connect to the blockchain node.

...

If connecting to a live network, make sure you provide a `--node` flag with the correct address of a live peer. For example, to check the status of `regen-devnet-5`, run the following command:

```
regen status --node http://18.220.101.192:26657
```
