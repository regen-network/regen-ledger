# Interfaces

Regen Ledger provides three interfaces for interacting with a node:

- [Command-Line Interface](#command-line-interface)
- [gRPC Interface](#grpc-interface)
- [REST Interface](#rest-interface)

## Command-Line Interface

When referring to the command-line interface for the Regen Ledger application, we are referring to the `regen` binary. The `regen` binary serves as the node client and the application client, meaning the `regen` binary can be used to both run a node and interact with it.

The `regen` binary is the all-in-one program for Regen Ledger. It can manage keys, run nodes, query state, and generate, sign, and broadcast transactions. If you are comfortable using the command-line, playing around with the `regen` binary is a great way to learn about Regen Ledger.

To learn more about the available commands, [install regen](get-started/README.md#install-regen) and run the following:

```bash
regen --help
```

For transaction commands, run the following:

```bash
regen tx --help
```

For query commands, run the following:

```bash
regen query --help
```

For documentation on all available commands, see [Commands](../commands/README.md).

## gRPC Interface

[gRPC](https://grpc.io/docs/what-is-grpc/introduction/) is a modern RPC framework that leverages [protocol buffers](https://developers.google.com/protocol-buffers) for encoding requests and responses between a client and service. Regen Ledger uses gRPC primarily for querying blockchain state (credit or token balances, data signature records, etc).

As a client developer, you can query state using a gRPC library in the programming language of your choice by importing and generating code from the [Regen Ledger Protobuf API](https://buf.build/regen/regen-ledger).

In some programming languages, you may be able to leverage a pre-existing library to take care of most of the heavy lifting, including the compiling of protobuf messages. If you develop in Go, you can use the [regen-ledger](https://github.com/regen-network/regen-ledger) modules. If you develop in Typescript, check out [regen-js](https://github.com/regen-network/regen-js).

Unlike the command-line interface, the gRPC interface can only be used to broadcast transactions. To generate and sign transactions, a developer will need to use another library such as [keplr-wallet](https://github.com/chainapsis/keplr-wallet) (see [regen-js](https://github.com/regen-network/regen-js) for an example using Typescript and Keplr).

In addition to using a gRPC library programmatically, you can also use [grpcurl](https://github.com/fullstorydev/grpcurl) - a command-line tool that lets you interact with gRPC servers from the command-line.

If you have a local node running, you can list the available services using the following command:

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

If you are running a local node and attempting to execute the above command but you are receiving an error, make sure the gRPC service is enabled in `~/.regen/config/app.toml`.

For more information about the gRPC interface, check out the [Cosmos SDK Documentation](https://docs.cosmos.network/main/run-node/interact-node.html).

## REST Interface

Regen Ledger leverages [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway), which makes all gRPC services and methods available via REST endpoints. As a client developer, you can query state using an HTTP library in the programming language of your choice and without generating code from [Regen Ledger Protobuf API](https://buf.build/regen/regen-ledger).

Like the gRPC interface, the REST interface can only be used to broadcast transactions. To generate and sign transactions, a developer will need to use another library such as [keplr-wallet](https://github.com/chainapsis/keplr-wallet) (see [regen-js](https://github.com/regen-network/regen-js) for an example using Typescript and Keplr).

In addition to using an HTTP library programmatically, you can also use the `curl` command.

To execute a call, you can use the following format:

```bash
curl \
    -X GET \
    -H "Content-Type: application/json" \
    http://localhost:1317/cosmos/bank/v1beta1/balances/<address>
```

You can also use your browser: [http://localhost:1317/cosmos/bank/v1beta1/balances/](http://localhost:1317/cosmos/bank/v1beta1/balances/)

If you are running a local node and the REST server is not working, make sure the REST service is enabled in `~/.regen/config/app.toml`. If you enable Swagger, the OpenAPI documentation will be available at [http://localhost:1317/swagger/](http://localhost:1317/swagger/).

For more information about the REST interface, check out the [Cosmos SDK Documentation](https://docs.cosmos.network/main/run-node/interact-node.html).
