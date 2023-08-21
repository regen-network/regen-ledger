# Message-Based Governance Proposals

This tutorial uses the [command-line interface (CLI)](../../ledger/interfaces.md#command-line-interface) to demonstrate how to submit message-based governance proposals. This tutorial uses messages from the `ecocredit` and `authz` modules as examples but any message for any module wired up to Regen Ledger works the same. For example, you could create and manage a group using governance proposals and the `group` module.

Community members are encouraged to submit thoughtful governance proposals on all networks governed and maintained by the Regen Network community. If you are planning to submit a proposal to Regen Mainnet, make sure you follow [Governance Guidelines](https://github.com/regen-network/governance#guidelines) for a successful outcome.

## Prerequisites

- [Install Regen](../../ledger/get-started/README.md)
- [Manage Keys](../../ledger/get-started/manage-keys.md)
- [Redwood Testnet](../../ledger/get-started/redwood-testnet.md)

## Recommended

- [Regen Ledger v5 Release Notes](https://github.com/regen-network/regen-ledger/releases/tag/v5.0.0)

## Introduction

For Regen Ledger and other Cosmos SDK applications, governance proposals known as "parameter-change proposals" have historically been used to update a specific set of parameters defined within each application module. As of Cosmos SDK `v0.46`, parameter-change proposals became known as "legacy" proposals in favor of message-based governance proposals.

Following the upgrade to Regen Ledger `v5`, the [submit-proposal](../../commands/regen_tx_gov_submit-proposal.md) command submits message-based governance proposals and the [submit-legacy-proposal](../../commands/regen_tx_gov_submit-legacy-proposal.md) command submits "legacy" proposals including "parameter-change proposals". To submit a parameter-change proposal for any module other than the `ecocredit` module, you need to use the `submit-legacy-proposal` command.

With message-based governance proposals, the messages in each proposal are signed by the `gov` module account if and when the proposal is executed, or in other words, messages within a proposal are called on behalf of the `gov` module account. This makes it possible for the `gov` module account to create and manage entities such as credit classes and groups.

#### Gov Account

With message-based governance proposals, the `gov` module account is the signer for each message in a proposal if and when the proposal is executed. As the signer, we need to know the address of the `gov` module account for when we construct messages.

To look up the `gov` module account, run the following command:

```sh
regen q auth module-account gov
```

If you are connected to Redwood Testnet, you should see the following:

```sh
account:
  '@type': /cosmos.auth.v1beta1.ModuleAccount
  base_account:
    account_number: "7"
    address: regen10d07y265gmmuvt4z0w9aw880jnsr700j9qceqh
    pub_key: null
    sequence: "0"
  name: gov
  permissions:
  - burner
```

#### Ecocredit Module

The `ecocredit` module is currently the only module wired up in Regen Ledger that supports message-based governance proposals specifically for updating "parameters" but any message from any module can be used in message-based governance proposals (as [demonstrated below](#grant-authorization) with `MsgGrant`).

With the migration to message-based governance proposals, a new set of messages were added to the `ecocredit` module that only an `authority` account can execute. In the current configuration of Regen Ledger, the `gov` module account is the `authority` account, and for the purpose of this tutorial we can think of the `authority` account as synonymous with the `gov` module account.

The following is the complete list of `ecocredit` messages added in Regen Ledger `v5` to support the changing of "parameters" using message-based governance proposals:

- [regen.ecocredit.v1.MsgAddAllowedBridgeChain](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.AddAllowedBridgeChain)
- [regen.ecocredit.v1.MsgAddClassCreator](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.AddClassCreator)
- [regen.ecocredit.v1.MsgAddCreditType](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.AddCreditType)
- [regen.ecocredit.v1.MsgRemoveAllowedBridgeChain](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.RemoveAllowedBridgeChain)
- [regen.ecocredit.v1.MsgRemoveClassCreator](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.RemoveClassCreator)
- [regen.ecocredit.v1.MsgSetClassCreatorAllowlist](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.SetClassCreatorAllowlist)
- [regen.ecocredit.v1.MsgUpdateClassFee](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.UpdateClassFee)
- [regen.ecocredit.basket.v1.MsgUpdateDateCriteria](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Msg.UpdateDateCriteria)
- [regen.ecocredit.basket.v1.MsgUpdateBasketFee](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Msg.UpdateBasketFee)
- [regen.ecocredit.marketplace.v1.MsgAddAllowedDenom](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.AddAllowedDenom)
- [regen.ecocredit.marketplace.v1.MsgRemoveAllowedDenom](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.RemoveAllowedDenom)

You can still query `ecocredit` "parameters" using the following command:

```sh
regen q ecocredit params
```

Each of these "parameters" are now stored separately in state. You can also query them individually. To see all available queries for the `ecocredit` module, run the following command:

```sh
regen q ecocredit --help
```

#### Submit Proposals

In the following sections, examples of message-based governance proposals using `ecocredit` and `authz` messages are provided. Each example only includes one message but all messages could be included within a single message-based governance proposal in which case all messages would be executed in the order provided when the proposal is executed.

## Add Credit Type

In this section we submit a message-based governance proposal to add a [credit type](../../modules/ecocredit/01_concepts.md#credit-type) to an allowlist of network approved credit types. A credit type represents the primary unit of measurement used when measuring ecological outcomes. Each credit class is assigned a single credit type upon creation.

For more information about credit types, see [Ecocredit Overview](../../modules/ecocredit) and [Ecocredit Concepts](../../modules/ecocredit/01_concepts.md).

### Create Metadata

Using the following as a template, create a JSON file that includes the proposal metadata.

```json
{
  "title": "Add biodiversity credit type",
  "summary": "This proposal adds a new biodiversity credit type."
}
```

Upload this file to IPFS, and then use the IPFS hash in the following section as "metadata".

### Create Proposal

Using the following as a template, create a JSON file that includes information about the proposal and the message to be executed if the proposal is successful:

```json
{
  "metadata": "<ipfs-hash>",
  "messages": [
    {
      "@type": "/regen.ecocredit.v1.MsgAddCreditType",
      "authority": "regen10d07y265gmmuvt4z0w9aw880jnsr700j9qceqh",
      "credit_type": {
        "abbreviation": "BIO",
        "name": "biodiversity",
        "unit": "species per square kilometer",
        "precision": 6
      }
    }
  ]
}
```

For more information about the message fields, check out the Protobuf API documentation:

- [regen.ecocredit.v1.Msg.AddCreditType](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.AddCreditType)

### Submit Proposal

Once the json file has been created, you can use the following command to submit the proposal:

```bash
regen tx gov submit-proposal [proposal-json]
```

For more information about the command, add `--help` or see [the docs](../../commands/regen_tx_gov_submit-proposal.md).

## Add Allowed Denom

In this section we submit a message-based governance proposal to add a token denom to the list of allowed denoms in the `ecocredit` module. The list of allowed denoms determines which tokens can be used to create sell orders in the marketplace.

Keep in mind that non-native tokens represented by an IBC denom are different with each source. For example, `atom` transferred from one chain is different from `atom` transferred from another chain, i.e. each represents `atom` but they have separate IBC denoms.

Check out [Understanding IBC Denoms](https://tutorials.cosmos.network/tutorials/understanding-ibc-denoms/) for more information.

### Create Metadata

Using the following as a template, create a JSON file that includes the proposal metadata.

```json
{
  "title": "Add $REGEN to the marketplace allowlist",
  "summary": "This proposal adds $REGEN to the marketplace allowlist."
}
```

Upload this file to IPFS, and then use the IPFS hash in the following section as "metadata".

### Create Proposal

Using the following as a template, create a JSON file that includes information about the proposal and the message to be executed if the proposal is successful:

```json
{
  "metadata": "<ipfs-hash>",
  "messages": [
    {
      "@type": "/regen.ecocredit.marketplace.v1.MsgAddAllowedDenom",
      "authority": "regen10d07y265gmmuvt4z0w9aw880jnsr700j9qceqh",
      "bank_denom": "uregen",
      "display_denom": "regen",
      "exponent": 6
    }
  ]
}
```

For more information about the message fields, check out the Protobuf API documentation:

- [regen.ecocredit.marketplace.v1.Msg.AddAllowedDenom](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.AddAllowedDenom)

### Submit Proposal

Once the json file has been created, you can use the following command to submit the proposal:

```bash
regen tx gov submit-proposal [proposal-json]
```

For more information about the command, add `--help` or see [the docs](../../commands/regen_tx_gov_submit-proposal.md).

## Grant Authorization

In this section we submit a message-based governance proposal to grant an account an authorization to call `MsgAddCreditType` without going through a network-wide governance process. Keep in mind that we could grant an authorization for any message, and doing so enables the account represented as the grantee to act on behalf of the gov module. 

Check out [Authz Overview](https://docs.cosmos.network/v0.46/modules/authz/) for more information about authorizations.

### Create Metadata

Using the following as a template, create a JSON file that includes the proposal metadata.

```json
{
  "title": "Grant authorization for adding credit types",
  "summary": "This proposal grants an authorization for adding credit types."
}
```

Upload this file to IPFS, and then use the IPFS hash in the following section as "metadata".

### Create Proposal

Using the following as a template, create a JSON file that includes information about the proposal and the message to be executed if the proposal is successful:

```json
{
  "metadata": "<ipfs-hash>",
  "messages": [
    {
      "@type": "/cosmos.authz.v1beta1.MsgGrant",
      "granter": "regen10d07y265gmmuvt4z0w9aw880jnsr700j9qceqh",
      "grantee": "regen1afk9zr2hn2jsac63h4hm60vl9z3e5u69gndzf7c99cqge3vzwjzs475lmr",
      "grant": {
        "authorization": {
          "@type": "/cosmos.authz.v1beta1.GenericAuthorization",
          "msg": "/regen.ecocredit.v1.MsgAddCreditType"
        },
        "expiration": "2024-01-01T00:00:00Z"
      }
    }
  ]
}
```

For more information about the message fields, check out the Protobuf API documentation:

- [regen.ecocredit.marketplace.v1.Msg.AddAllowedDenom](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.AddAllowedDenom)

### Submit Proposal

Once the json file has been created, you can use the following command to submit the proposal:

```bash
regen tx gov submit-proposal [proposal-json]
```

For more information about the command, add `--help` or see [the docs](../../commands/regen_tx_gov_submit-proposal.md).

## View Proposal

Once you've submitted a proposal, you can view the proposal using a query command.

You can query all proposals using the following command:

```sh
regen q gov proposals
```

You can also query an individual proposal using the following command:

```sh
regen q gov proposal [proposal-id]
```

## Conclusion

Congratulations! You have now successfully submitted a message-based governance proposal.

If you submit a proposal on Redwood Testnet that you would like to see succeed, be sure to reach out on Discord, otherwise the proposal may go unnoticed.

As mentioned above, proposals submitted to Regen Mainnet should follow [Governance Guidelines](https://github.com/regen-network/governance#guidelines) for best results. We also recommend submitting a proposal on Redwood Testnet before submitting on Regen Mainnet if you are new to message-based governance proposals or proposals in general.