# Data Module (CLI)

We will interact with Regen's [Data module](./modules/data/), also called `x/data`, in the example. The Data module allows a user to anchor any piece of data on the blockchain, by storing its hash on-chain. `x/data` uses the [CID](https://github.com/multiformats/cid) specification for addressing the hash of this data.

::: tip
All the commands below work if you're connecting to a [remote node](../getting-started/live-networks) (with the `--node` flag) or to a [local node](../getting-started).
:::

### Create a KeyPair

For sending transactions to the blockchain, you will need a keypair. Regen Ledger keys can be managed with the `regen keys` subcommand. A new key pair can be generated using:

```sh
KEY_NAME=my_key # Or choose your own key name.
regen keys add $KEY_NAME

# We will also save the generated address in a variable for later use.
MY_ADDRESS=$(regen keys show $KEY_NAME -a)
```

If you'd like to use a custom keyring backend, see the [Cosmos SDK keyring docs](https://docs.cosmos.network/main/run-node/keyring.html). Make sure you save the seed mnemonic in a safe place!

To receive some funds on this address, copy the content of the `$MY_ADDRESS` variable, and paste your address on [our faucet](https://regen.vitwit.com/faucet) to receive some funds.

### Sending a Transaction

To send a transaction for anchoring a CID, run the following command:

```sh
EXAMPLE_CID=zb2rhe5P4gXftAwvA4eXQ5HJwsER2owDyS9sKaQRRVQPn93bA
regen tx data anchor $KEY_NAME $EXAMPLE_CID
```

A summary of the transaction will be displayed, with a confirmation message. Pressing `y` will confirm, and `n` will reject and cancel the transaction. After a couple of seconds, the confirmation of the transaction will be displayed, in JSON format. Make sure that the `code` field contains `0`, as this means the transaction was successful. If not, inspect the `logs` field for potential errors.

Let's sign the anchor we just created. Signing an anchor in the Data module is attesting to the veracity of the data the CID represents.

```sh
regen tx data sign $KEY_NAME $EXAMPLE_CID
```

Similarily to above, you should see a transaction confirmation shortly after pressing `y`.
