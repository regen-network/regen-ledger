# Getting Started

## Installation

### Building from source

Prerequisites: The [Go](https://golang.org/doc/install) compiler **version 1.14** (we use
go modules) or later and GNU make. The go bin path (usually `$HOME/go/bin`) should be
on your system `PATH`.

Run `make install` to build the code. The `regen` binary serves as both the blockchain daemon, as well as
the command line blockchain client.

## Configuring the command line client

By default the command line client will connect to a [local node](#running-a-local-node).
To connect to a testnet, you must know the testnet's chain ID and the address of the node. More information on the current testnet can be found in the [testnets repository](https://github.com/regen-network/testnets).

`regen` can be configured to connect to a testnet node automatically by setting these
parameters in `$HOME/.regen/config/config.toml`. This file can be generated
automatically by running:

```sh
regen init --chain-id [chain ID] --node [node address]
```

Check current testnet status [here](https://github.com/regen-network/testnets).

Run `regen status` to see if you are able to connect to the block-chain node.

### Creating a key

To run any transactions you must have a key. A new one can be generated using:

```sh
regen keys add [key-name]
```

Make sure you save the seed mnemonic in a safe place!

## Running a local node

Initialize the node config

```sh
regen init
```

Start the node
```sh
regen start
```

Local node config will be saved to `~/.regen` and the chain ID can be found in `~/.regen/config/genesis.json`.