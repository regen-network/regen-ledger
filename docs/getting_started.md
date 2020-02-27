# Getting Started

## Installation

### Building from source

Prerequisites: The [Go](https://golang.org/doc/install) compiler **version 1.11** (we use
go modules) or later and GNU make. The go bin path (usually `$HOME/go/bin`) should be
on your system `PATH`.

Run `make install` to build the code.`xrnd` is the blockchain daemon. `xrncli`
is the command line blockchain client. Both of them provide help messages when run.

#### Go 1.13

If you are running Go 1.13, there are a few minor changes that will need to be made to `Makefile` and `go.mod` to use properly formatted pseudo-versions.

In `Makefile` (line 80):

```
-       go install -mod=readonly $(BUILD_FLAGS) ./cmd/xrnd
+       go install $(BUILD_FLAGS) ./cmd/xrnd
```

In `go.mod` (lines 18, 48):

```
-    github.com/leanovate/gopter v0.0.0-20190000000000-6e7780f59df75750618bf30eeafcb1a88e457fcb
+    github.com/leanovate/gopter v0.0.0-20190326081808-6e7780f59df7
...
...
-  replace github.com/cosmos/cosmos-sdk => github.com/regen-network/cosmos-sdk v0.0.0-0.20190329021654-281da889de6ca3c7784d5570fd95de78d7d23a59
+  replace github.com/cosmos/cosmos-sdk => github.com/regen-network/cosmos-sdk v0.0.0-20190827142854-281da889de6c
```


## Configuring the command line client

By default the command line client will connect to a [local node](#running-a-local-node).
To connect to a testnet, you must know the testnet's chain ID and the address of the node. More information on the current testnet can be found in the [testnets repository](https://github.com/regen-network/testnets).

`xrncli` can be configured to connect to a testnet node automatically by setting these
parameters in `$HOME/.xrncli/config/config.toml`. This file can be generated
automatically by running:

```sh
xrncli init --chain-id [chain ID] --node [node address]
```

Check current testnet status [here](https://github.com/regen-network/testnets).

Run `xrncli status` to see if you are able to connect to the block-chain node.

### Creating a key

To run any transactions you must have a key. A new one can be generated using:

```sh
xrncli keys add [key-name]
```

Make sure you save the seed mnemonic in a safe place!

## Running a local node

Initialize the node config

```sh
xrnd init
```

Start the node
```sh
xrnd start
```

Local node config will be saved to `~/.xrnd` and the chain ID can be found in `~/.xrnd/config/genesis.json`.
