# Regen Ledger API & CLI

Regen Ledger provides three interfaces for interacting with a node:

- command-line interface
- gRPC interface
- REST interface

Each interface can be used to query the state. The command-line interface can be used to generate, sign, and broadcast transactions, whereas the gRPC and REST interfaces can only be used to broadcast transactions (a transaction needs to be generated and signed either programmatically or using the command-line interface before it can be broadcasted using the gRPC or REST interfaces).

## Command-Line Interface

The most straightforward way to interact with a node is using the command-line interface.

The `regen` binary serves as the node client and the application client, meaning the `regen` binary can be used to both run a node and interact with it. In [Get Started](../get-started/README.md), we started a local node using the `regen` binary and then interacted with that node by submitting queries and transactions. For more examples of interacting with a node using the command-line interface, see [Tutorials](../../tutorials/README.md).

To learn more about the available commands, [install regen](../get-started/README.md#install-regen) and run the following:

```bash
regen --help
```

For transaction commands:

```bash
regen tx --help
```

For query commands:

```bash
regen query --help
```

For a full list of the available commands, see [Commands](../../commands.md).

## gRPC Interface

[gRPC](https://grpc.io/docs/what-is-grpc/introduction/) is a modern RPC framework that leverages [protocol buffers](https://developers.google.com/protocol-buffers) for encoding requests and responses between a client and service. Regen Ledger uses gRPC primarily for querying blockchain state (credit or token balances, data signature records, etc). As a client developer, this means you can query Regen Ledger state directly by using a gRPC library in your programming language of choice, in combination with Regen Ledger's protobuf definitions defined [here](https://buf.build/regen/regen-ledger).

In addition to using a gRPC library, you can also use [grpcurl](https://github.com/fullstorydev/grpcurl) - a command-line tool that lets you interact with gRPC servers. If you have a local node running, you can list the protobuf services available using the following command:

```bash
grpcurl -plaintext localhost:9090 list
```

To execute a call, you can use the following format:

```bash
grpcurl \
    -plaintext \
    -d '{"address":"<address>"}' \
    localhost:9090 \
    cosmos.bank.v1beta1.Query/AllBalances
```

In some programming languages, you may be able to leverage a pre-existing client library to take care of most of the heavy lifting, including compiling protobuf messages. For javascript/typescript developers, [CosmJS](https://github.com/cosmos/cosmjs) is a great place to start.

::: tip
While CosmJS provides basic support for all Cosmos SDK based blockchains, you will still need to compile the protobuf messages for Regen Ledger's own modules (e.g. ecocredit module, data module) if you intend to interact with ecosytem service credits or on-chain ecological data.
:::

For more information about the gRPC interface, check out the [Cosmos SDK Documentation](https://docs.cosmos.network/main/run-node/interact-node.html).

## REST Interface

All gRPC services and methods on Regen Ledger are available for more convenient REST based queries through [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway).

For example, you can query the balance of an address using the following `curl` command:

```bash
curl \
    -X GET \
    -H "Content-Type: application/json" \
    http://localhost:1317/cosmos/bank/v1beta1/balances/<address>
```

In order to interact with the REST interface, make sure you have API server and (optionally) Swagger UI enabled in your `~/.regen/config/app.toml` file. With Swagger UI enabled, you can go to `http://localhost:1317/swagger/` to read through the OpenAPI documentation.

For more information about the REST interface, check out the [Cosmos SDK Documentation](https://docs.cosmos.network/main/run-node/interact-node.html).
