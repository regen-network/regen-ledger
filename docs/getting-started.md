# Getting Started

## Installation

### Building from source

Prerequisites: The [Go](https://golang.org/doc/install) compiler **version 1.15** (we use
go modules) or later and GNU make. The go bin path (usually `$HOME/go/bin`) should be
in your system `PATH`.

Run `make install` to build the code. The `regen` binary serves as both the blockchain daemon, as well as
the command line blockchain client.

## Configuring the command line client

By default the command line client will connect to a [local node](#running-a-local-node).
To connect to a testnet or devnet, you must know the network's chain ID and the address of the node. More information on the current testnets and devnet can be found in the [testnets repository](https://github.com/regen-network/testnets).

`regen` can be configured to connect to a testnet node automatically by setting these
parameters in `$HOME/.regen/config/config.toml`. This file can be generated
automatically by running:

```sh
regen init --chain-id [chain ID] --node [moniker]
```
If you're wanting to connect to our devnet that launched during the [Open Climate Collabathon](https://www.collabathon.openclimate.earth/), the chain-id is `regen-devnet-1`.


If not connecting to a live testnet or devnet, initialize the node config for local development:

```sh
regen init [moniker]
```


## Running a node

Start the node:
```sh
regen start
```
::: tip

Starting a node is not necessary if you're only wanting to interact with a live blockchain. In that case, you can use the `regen` binary purely as a CLI for interacting with a live network. Just make sure you always provide a `--node` flag with a valid peer address.

:::

Regardless of whether you're running a local node (as above), or connecting directly to a live network, you can run `regen status` to verify that the CLI is able to connect to the blockchain node. If connecting to a live network, make sure you provide a `--node` flag with the correct address of a live peer. For connecting to `regen-devnet-1`, use the following:

```sh
regen status --node http://18.220.101.192:26657
```


Local node config will be saved to `~/.regen` and the chain ID can be found in `~/.regen/config/genesis.json`.

## Creating a key pair

To execute any transactions you must have a key pair. Regen Ledger keys can be managed with the `regen keys` subcommand. A new key pair can be generated using:

```sh
regen keys add [key-name]
```
If you'd like to use a custom keyring backend, you can find more information on the [Cosmos SDK keyring docs](https://docs.cosmos.network/master/run-node/keyring.html).


Make sure you save the seed mnemonic in a safe place!
