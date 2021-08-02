# Regen Ledger API & CLI

::: tip

All the commands below work both if you're connecting to a [remote node](./getting-started.md#connecting-to-an-existing-network) (with the `--node` flag) or to [your own local node](./getting-started.md#running-a-localnet).

For simplicity's sake, you can use our public node by appending `--node http://18.220.101.192:26657` to all the commands below.

:::

Regen provides three interfaces for interacting with a node: a command-line interface (CLI), a gRPC endpoint and a REST API.

## Command-line Interface

The most straightforward way to interact with a node is using the CLI. The `regen` binary serves both as node and as a client CLI, and we will see here how to use it for the latter case.

We will interact with Regen's [Data module](./modules/data/) in the example, also called `x/data`. The Data module allows a user to anchor any piece of data on the blockchain, by storing its hash in on-chain state. `x/data` uses the [CID](https://github.com/multiformats/cid) specification for addressing the hash of this data.

### Create a KeyPair

For sending transactions on the blockchain, you will need a keypair. Regen Ledger keys can be managed with the `regen keys` subcommand. A new key pair can be generated using:

```sh
KEY_NAME=my_key # Or choose your own key name.
regen keys add $KEY_NAME

# We will also save the generated address in a variable for later use.
MY_ADDRESS=$(regen keys show $KEY_NAME -a)
```

If you'd like to use a custom keyring backend, you can find more information on the [Cosmos SDK keyring docs](https://docs.cosmos.network/master/run-node/keyring.html). Make sure you save the seed mnemonic in a safe place!

To receive some funds on this address, copy the content of the `$MY_ADDRESS` variable, and paste your address on [our faucet](https://regen.vitwit.com/faucet) to receive some funds.

### Sending a Transaction

To send a transaction for anchoring a CID, run the following command:

```sh
EXAMPLE_CID=zb2rhe5P4gXftAwvA4eXQ5HJwsER2owDyS9sKaQRRVQPn93bA
regen tx data anchor $KEY_NAME $EXAMPLE_CID
```

A summary of the transaction will be displayed, and you can confirm sending it by pressing the `y` key. After a couple of seconds, the confirmation of the transaction will be displayed, in JSON format. Make sure that the `code` field contains `0`, as this means the transaction was successful. If not, then look at the `logs` field for potential errors.

Let's sign the anchor we just created. Signing an anchor in the Data module is attesting of the veracity of the data the CID represents.

```sh
regen tx data sign $KEY_NAME $EXAMPLE_CID
```

Similarily to above, you should see a transaction confirmation after pressing `y`.

### Querying the State

Once the transaction(s) are sent, we can query the state to retrieve on-chain data. This is done with the following command:

```sh
regen q data $EXAMPLE_CID
```

This command will fetch the given CID on chain, as well as all its (potential) signers. Since one user signed the CID (yourself), you should see an output similar to:

```sh
content: null
signers:
- regen1qx9tsl46vkgf9tf8pl3chm7mkmqa5s30ydjdzx
timestamp: "2020-11-14T03:45:50.924123Z"
```

## gRPC

[gRPC](https://grpc.io/docs/what-is-grpc/introduction/) is a modern RPC framework that leverages
[protocol buffers](https://developers.google.com/protocol-buffers) for encoding requests
and responses between a client and service. In the case of Regen Ledger, we use gRPC mostly for querying blockchain 
state (credit or token balances, data signature records, etc). As a client developer, this means that you
can query Regen Ledger state directly by using a gRPC library for your language, in combination with Regen Ledger's
protobuf definitions defined [here](https://github.com/regen-network/regen-ledger/tree/master/proto/regen).

In some languages, you may be able to leverage a pre-existing client library to take care of most of the heavy lifting,
including compiling protobuf messages. For javascript/typescript developers, [CosmJS](https://github.com/cosmos/cosmjs)
is a great place to start.

::: tip

While CosmJS has basic support for all Cosmos SDK based blockchains, you will still need to compile the protobuf messages
for Regen Ledger's own modules (e.g. data module, ecocredit module) if you intend to interact with credits or on-chain
ecological data.

And be sure to use [cosmjs/stargate](https://cosmos.github.io/cosmjs/latest/stargate/index.html)
 client!
:::

For the purposes of this guide, we'll use [gRPCurl](https://github.com/fullstorydev/grpcurl), a command-line tool
which acts as a `curl` replacement for gRPC services. Follow the instructions on their [github](https://github.com/fullstorydev/grpcurl)
to download the binary.

Now let's get to querying!

Assuming you have a local node running (either a localnet, or connected to our devnet), you should be able
to run the following to list the query services available:

```sh
grpcurl -plaintext localhost:9090 list
```

You should see a list of gRPC services like `cosmos.bank.v1beta1.Query`, `regen.data.v1alpha1.Query`. Each of these
represents a different API endpoint which you can query for some relevant state from the blockchain.

Let's see if we can query the same CID that we anchored and signed from the previous section. To find out
which methods are exposed via the `regen.data.v1alpha1.Query` service, we need to provide `grpcurl` with
the .proto files corresponding to this service:

```sh
$ grpcurl -proto ../proto/regen/data/v1alpha1/query.proto describe regen.data.v1alpha1.Query
# regen.data.v1alpha1.Query is a service:
# // Query is the regen.data.v1alpha1 Query service
# service Query {
#   // ByCid queries data based on its CID.
#   rpc ByCid ( .regen.data.v1alpha1.QueryByCidRequest ) returns ( .regen.data.v1alpha1.QueryByCidResponse );
# }
```

Here we see that there is one method, `ByCid`, which takes a `QueryByCidRequest`, and returns a `QueryByCidResponse`.
The structure of these request & response messages can similarly be quiered with `grpcurl`, or you can find
details on them in our [protobuf documentation](./modules/data/protobuf.md#regen-data-v1alpha1-query-proto).

Now that we know the structure of our `QueryByCidRequest`, we can query the `ByCid` method directly using a JSON
encoding of the `QueryByCidRequest` message.

_Note: Since gRPCurl requires bytes be encoded as base64 strings, we have to do some gymnastics to decode our CID
and re-encode the raw bytes using base64. The base64 CID string below is the correct one for the `$EXAMPLE_CID` in
the CLI tutorial above. For more details on how CIDs work see the [CID spec](https://github.com/multiformats/cid)._

```sh
grpcurl -proto ../proto/regen/data/v1alpha1/query.proto \
    -d '{"cid": "AVUSIG5v95UKNhh6gBYTQm6Fjc5obNfX48D8Qu4DMActJFyV"}' \
    -plaintext localhost:9090 regen.data.v1alpha1.Query/ByCid
```

The result should look something like this:

```json
{
  "timestamp": "2020-11-14T03:45:50.924123Z",
  "signers": [
    "regen1qx9tsl46vkgf9tf8pl3chm7mkmqa5s30ydjdzx"
  ]
}
```


## REST API

::: tip COMING SOON

All gRPC services and methods on Regen Ledger will soon be made available for more convenient
REST based queries through [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway).
:::

Currently gRPC-gateway endpoints have yet to be added for Regen Ledger's own modules, but the basic
Cosmos REST API does exist, so you can stil use the REST API for queries to general modules like `x/bank`,
`x/staking`, etc.

If you're eager to play around with what's available so far while we're still working on
full REST API support for Regen Ledger, make sure you have API server and (optionally)
Swagger UI enabled in your `~/.regen/config/app.toml` file, and go to
`http://localhost:1317/swagger/` to read through the OpenAPI documentation.


