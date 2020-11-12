# Regen Ledger API & CLI

::: tip

All the commands below work both if you're connecting to a [remote node](./getting-started#connecting-to-an-existing-network) (with the `--node` flag) or to [your own local node](./getting-started#running-a-localnet).

For simplicity's sake, you can use our public node by appending `--node http://18.220.101.192:26657` to all the commands below.

:::

Regen provides three interfaces for interacting with a node: a command-line interface (CLI), a gRPC endpoint and a REST API.

## Command-line Interface

The most straightforward way to interact with a node is using the CLI. The `regen` binary serves both as node and as a client CLI, and we will see here how to use it for the latter case.

We will interact with Regen's [Data module](./modules/data/) in the example, also called `x/data`. The Data module allows a user to anchor any piece of data on the blockchain, by storing its hash in the state. `x/data` uses the [CID](https://github.com/multiformats/cid) specification for addressing the hash of this data.

### Create a KeyPair

For sending transactions on the blockchain, you will need a keypair. Regen Ledger keys can be managed with the `regen keys` subcommand. A new key pair can be generated using:

```sh
KEY_NAME=my_key # Or choose your own key name.
regen keys add $KEY_NAME

# We will also save the generated address in a variable for later use.
MY_ADDRESS=$(regen keys show $KEY_NAME -a)
```

If you'd like to use a custom keyring backend, you can find more information on the [Cosmos SDK keyring docs](https://docs.cosmos.network/master/run-node/keyring.html). Make sure you save the seed mnemonic in a safe place!

To receive some funds on this address, copy the content of the `$MY_ADDRESS` variable, and paste your address on [our faucet](https://faucet.devnet.regen.vitwit.com) to receive some funds.

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
TODO
```

## gRPC

TODO Which gRPC client should be use? grpcurl? CosmJS?

## REST API
