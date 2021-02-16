# Getting Started

## Installation

### Building from source

Prerequisites: The [Go](https://golang.org/doc/install) compiler **version 1.15** (we use
go modules) or later and GNU make. The go bin path (usually `$HOME/go/bin`) should be
in your system `$PATH`.

Run `make install` to build the code. The `regen` binary serves as both the blockchain daemon, as well as
the command line blockchain client.

## Configuring the command line client

By default the command line client will connect to a [local node](#running-a-node).
To connect to a testnet or devnet, you must know the network's chain ID and the address of a public node. More information on the current testnets and devnet can be found in the [testnets repository](https://github.com/regen-network/testnets).

`regen` can be configured to connect to a testnet node automatically by setting these
parameters in `$HOME/.regen/config/config.toml`, under the `rpc.laddr` field. If this file doesn't exist yet, it can be generated automatically by running:

```sh
regen init [moniker] --chain-id [chain ID]
```

If you're wanting to connect to our devnet that launched during the [Open Climate Collabathon](https://www.collabathon.openclimate.earth/), the chain-id is `regen-devnet-1`.

If not connecting to a live testnet or devnet, initialize the node config for local development:

```sh
regen init [moniker]
```

## Running a Node

Running a node consists of first generating a valid genesis file, and then running the `regen start` command to start the node.

::: tip

Starting a node is not necessary if you're only wanting to interact with a live blockchain. In that case, you can use the `regen` binary purely as a CLI for interacting with a live network. Just make sure you always provide a `--node` flag with a valid peer address. See [the API documentation](./api.md) for more info.

:::

### Connecting to an Existing Network

If you're wanting to connect to our devnet that launched during the [Open Climate Collabathon](https://www.collabathon.openclimate.earth/), you can fetch the genesis file here: http://18.220.101.192:26657/genesis. Save it in the `$HOME/.regen/config/` folder:

```sh
curl http://18.220.101.192:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

Be sure to also add well-known seed nodes for your local node's initial peer discovery:

```sh
PERSISTENT_PEERS="a621e6bf1f5981b3e72e059f86cbfc9dc5577fcb@18.220.101.192:26656"
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' ~/.regen/config/config.toml
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

Finally, run the node with

```sh
regen start
```

and you should start syncing the chain.

### Running a Localnet

::: tip

Running a localnet with a new genesis file will effectively create a new chain.

:::

Once you [configured your CLI client](#configuring-the-command-line-client), you need to populate the genesis file with an initial validator set, which, in our case, will be a singleton with your own validator node.

As a validator who signs blocks, your node must have a public/private keypair. Regen Ledger keys can be managed with the `regen keys` subcommand. A new key pair can be generated using:

```sh
KEY_NAME=my_validator # Or choose your own key name.
regen keys add $KEY_NAME

# We will also save the generated address in a variable for later use.
MY_VALIDATOR_ADDRESS=$(regen keys show $KEY_NAME -a)
```

If you'd like to use a custom keyring backend, you can find more information on the [Cosmos SDK keyring docs](https://docs.cosmos.network/master/run-node/keyring.html).

Make sure you save the seed mnemonic in a safe place!

Afterwards, run the following commands to add this newly generated validator to the
genesis file:

```sh
# Populate the validator account with some funds.
# The default bonding token in local networks is called "stake",
# but if you're connecting one of our devnets or testnet, you may
# need to use the denom "tree" instead.
regen add-genesis-account $MY_VALIDATOR_ADDRESS 100000000stake

# Create a gentx.
regen gentx $KEY_NAME --amount 100000stake --chain-id [chain-id]

# Add the gentx to the genesis file.
regen collect-gentxs
```

If you wish to learn more about each individual command in the snippet above, the [Cosmos SDK documentation](https://docs.cosmos.network/master/run-node/run-node.html) is a good place to start.

Finally, start the node with

```sh
regen start
```

and your node should start producing blocks.

### Check the Node's Status

Regardless of whether you're running a local node (as above), or connecting directly to a live network, you can run `regen status` to verify that the CLI is able to connect to the blockchain node. If connecting to a live network, make sure you provide a `--node` flag with the correct address of a live peer. For connecting to `regen-devnet-1`, use the following:

```sh
regen status --node http://18.220.101.192:26657
```

Local node config will be saved to `~/.regen` and the chain ID can be found in `~/.regen/config/genesis.json`.
