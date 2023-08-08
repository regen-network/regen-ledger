# Message-Based Governance Proposals

This tutorial uses the [command-line interface (CLI)](../../ledger/interfaces.md#command-line-interface) to demonstrates how to submit message-based governance proposals. This tutorial uses messages from the `ecocredit` and `authz` modules as examples but any message for any module wired up to Regen Ledger works the same. For example, you could create and manage a group using governance proposals and the `group` module.

Community members are encouraged to submit governance proposals on all networks governed and maintained by the Regen Network community. If you are planning to submit a proposal to Regen Mainnet, make sure you follow [Governance Guidelines](https://github.com/regen-network/governance#guidelines) for a successful outcome.

Also, if you are submitting a proposal to Regen Mainnet, make sure you provide a meaningful title and description. The description should include a rationale as to why this proposal should pass and can be written in either plain text or markdown (see [Proposal #15](https://wallet.keplr.app/chains/regen/proposals/15) for an example).

## Prerequisites

- [Install Regen](../../ledger/get-started/README.md)
- [Manage Keys](../../ledger/get-started/manage-keys.md)
- [Redwood Testnet](../../ledger/get-started/redwood-testnet.md)

## Recommended

- [Regen Ledger v5 Release Notes](https://github.com/regen-network/regen-ledger/releases/tag/v5.0.0)

## Introduction

Governance proposals have historically been used to update a specific set of parameters defined within each application module. These types of proposals are now referred to as "legacy" proposals and are slowly being replaced by message-based governance proposals.

Following the upgrade to Regen Ledger `v5`, the `submit-proposal` command now submits message-based governance proposals, so if you need to submit a proposal to update a module parameter outside the `ecocredit` module, you will mostly likely need to use the [submit-legacy-proposal](../../commands/regen_tx_gov_submit-legacy-proposal.md) command rather than `submit-proposal`.

When the `ecocredit` module was updated in `v5` to use message-based governance proposals, how ecocredit "parameters" are changed was also updated, which means `submit-legacy-proposal` no longer works with the `ecocredit` module.

Although how "parameters" are changed, you can still query them using the following command:

```sh
regen q ecocredit params
```

Each of these are now stored separately in state and are more loosely defined as parameters. A new set of messages were also added to update each independently. The messages can only be executed by the `authority` account, which is currently configured to be the `gov` module account.

#### Gov Account

With message-based governance proposals, the `gov` module account is the signer for each message within a proposal when the proposal is executed. As the signer, we need to know the address of the `gov` module account for when we construct our messages.

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

#### Submit Proposals

The next sections of the tutorial provides governance proposal examples. Each example only includes one message but all messages could be included within a single proposal in which case all messages included would all be executed within the same transaction when the proposal is executed.

## Add Credit Type

...

### Create Proposal

Using the following as a template, create a JSON file that includes information about the proposal and the message to be executed if the proposal is successful:

```json
{
  "title": "Add biodiversity credit type",
  "description": "This proposal adds a new biodiversity credit type.",
  "messages": [
    {
      "@type": "/regen.ecocredit.v1.MsgAddCreditType",
      "authority": "regen10d07y265gmmuvt4z0w9aw880jnsr700j9qceqh",
      "": ""
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

### Create Proposal

Using the following as a template, create a JSON file that includes information about the proposal and the message to be executed if the proposal is successful:

```json
{
  "title": "Add $REGEN to the marketplace allowlist",
  "description": "This proposal adds $REGEN to the marketplace allowlist",
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

In this section we submit a message-based governance proposal to grant an account the authorization to call `MsgAddCreditType` without going through a network-wide governance process. Keep in mind that we could grant an authorization for any message, and doing so enables the account represented as the grantee to act on behalf of the gov module. 

Check out [Authz Overview](https://docs.cosmos.network/v0.46/modules/authz/) for more information about authorizations.

### Create Proposal

Using the following as a template, create a JSON file that includes information about the proposal and the message to be executed if the proposal is successful:

```json
{
  "title": "Grant authorization for adding credit types",
  "description": "This proposal grants an authorization to a group account to call MsgAddCreditType on behalf of the gov module",
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
        "expiration": null
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

## Conclusion

For more messages, check out [Modules](../../modules).