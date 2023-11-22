# Create a Validator

This document provides instructions for creating a validator.

## Prerequisites

This assumes that you already have a full node [initialized and running](initialize-node.md).

## Add Validator Key

As a validator who signs blocks, your node must have a public/private key pair. Regen Ledger keys can be managed with the `regen keys` subcommand. A new key pair can be generated using:

```bash
regen keys add [name]
```

::: warning
If you create a new key, make sure you store the mnemonic phrase in a safe place. You will not be able to recover your new key without the mnemonic phrase.
:::

If you'd like to use an existing key or a custom keyring backend, you can find more information about adding keys and keyring backends in the Cosmos SDK [Keyring](https://docs.cosmos.network/main/run-node/keyring.html) documentation.

## Create Validator

The next step will be to create a validator. You will need to have enough REGEN tokens to stake and to submit the transaction. For more information about the REGEN token, see the [token page](https://www.regen.network/token/). 

::: warning
You'll want to carefully consider the options you set when creating a validator.
:::

Submit a transaction to create a validator:

```bash
regen tx staking create-validator \
  --amount=<stake_amount> \
  --pubkey=$(regen tendermint show-validator) \
  --moniker="<node_moniker>" \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --gas="auto" \
  --from=<key_name>
```

## Prepare Upgrade

The next step will be to prepare your node for [software upgrades](../upgrades/README.md).