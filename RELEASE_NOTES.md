# Regen Ledger v5.0.0

## New Features

The new features made available in Regen Ledger `v5.0.0` are as follows:

### DAO support via Group Accounts

The `group` module started out as a custom module built within the Regen Ledger repository and has since been added to Cosmos SDK as a core module refined and maintained by the Cosmos SDK core developers (which includes Regen Network Development team members). The `group` module enables the creation and management of on-chain multisig accounts and voting for message execution based on configurable decision policies.

What does this mean within the context of Regen Ledger functionality? All entities on chain that are currently managed using a legacy multisig or an individual account can now be managed by a group account. For example, the role of the credit class admin can be reassigned to a group account and the group can create and update decision policies for the execution of messages that are restricted to the role of the credit class admin. Other roles for on-chain entities in the ecocredit module that can be managed by group accounts include the credit class creator, the credit class issuer, the project admin, and the basket curator. The role of the resolver manager in the data module can also be managed by a group account.

The group module also enables any individual or group of individuals separate of the predefined roles for on-chain entities to create and manage a group account. A great example being participants of the enDAOment program... Another example being a group of individuals requesting community funds... Another example being an individual that would like to create a two-factor authentication experience where they are required to sign-off on message execution using multiple accounts (i.e. one account on the phone, one account on the laptop, and requiring both accounts to vote on the execution of a message).

For more information about the group module, check out the [group module documentation](https://docs.cosmos.network/v0.46/modules/group/).

### Message-Based Governance Proposals

Cosmos SDK v0.46 included a new approach to governance proposals...

All parameter change proposals are...

Bridge operations must now specify a target/source that exists in the AllowedBridgeChain table.

For more information about the gov module, check out the [gov module documentation](https://docs.cosmos.network/v0.46/modules/gov/).

### Interchain Accounts

Two new modules have been added to support interchain accounts. Interchain accounts enables cross-chain account management built upon IBC. One of the modules is an application module built and maintained by the IBC team within the `ibc-go` repository (`ica`) and the other is an application module built and maintained by the RND team within the `regen-ledger` repository (`intertx`).

Interchain accounts are accounts controlled programmatically by counterparty chains via IBC packets. Unlike a traditional account, an interchain account does not have a private key and therefore does not sign transactions. The account is registered on a "host chain" via a "controller chain" and the controller chain sends instructions (IBC packets with Cosmos SDK messages) to the host chain that the interchain account then executes.

The `ica` module has two submodules (`host` and `controller`). The `host` submodule enables a Regen Ledger chain (i.e. Regen Mainnet and Regen Redwood) to act as a "host chain" and the `controller` submodule enables a Regen Ledger chain to act as a "controller chain". The `host` and `controller` submodules will not be enabled following the upgrade of an existing Regen Ledger chain and therefore each will require an on-chain governance proposal to enable. Which messages allowed to be executed by interchain accounts will also need to be added to an `allowed_messages` parameter in the `host` submodule via subsequent governance proposals.

The `intertx` module is the implementation of the `controller` functionality allowing for the registering of interchain accounts and submitting transactions...

For more information about interchain accounts, check out the [interchain accounts documentation](https://ibc.cosmos.network/main/apps/interchain-accounts/overview).

### Relayer Incentivization

The `fee` module is a self-contained [middleware](https://ibc.cosmos.network/main/ibc/middleware/develop.html) module that extends the base IBC application module. The fee module was designed as an incentivization mechanism to help cover the operational costs of running a relayer (i.e. running full nodes to query transaction proofs and paying for transaction fees associated with IBC packets).

There are three fees within the fee model, one for receiving the packet, one for acknowledging the packet, and one for timeouts. The fees are held in escrow until the packet is either successful or times out. In the case of a successful packet, the timeout fee will be reimbursed, and in the case of an unsuccessful packet, the receiving and acknowledging fee will be reimbursed.

The first version of the fee module only supports incentivization of new channels and existing channels will need to wait for additional functionality to support upgradeability. Using the fee middleware with IBC transactions is optional and acts more like a "tip". End users can manually incentivize IBC packets using the CLI and client developers can leverage the gRPC endpoints to integrate fees within their application.

For more information about fee middleware, check out the [fee middleware documentation](https://ibc.cosmos.network/main/middleware/ics29-fee/overview).

## Additional Changes

### gRPC Error Codes

A community member reported that the gRPC error codes for queries were not being reported correctly, in some cases the error code was `Unknown` and in other cases the error code did not match the standard gRPC error codes. Not all projects building on Regen Ledger will use the error messages provided by Regen Ledger and may choose to create their own error messages based on the error codes returned. To ensure we are providing a good developer experience, we have fixed and updated gRPC error codes to return the expected gRPC error codes.

### SDK Module Manager

Regen Ledger has historically been using a custom module manager within the application wiring. Regen Ledger v5.0 migrates from the custom module manager to the latest Cosmos SDK module manager.

### Experimental Build

Following Regen Ledger v4.0, and now with Regen Ledger v5.0, all experimental features that were being developed within the Regen Ledger codebase have been stabilized and included in the stable application build. The experimental application build option has therefore been removed. We will be considering a separate release cycle for experimental features if we choose to continue providing experimental features alongside a stable application build.

## Changelog

For a full list of changes since Regen Ledger v4.1, please see [CHANGELOG.md](./CHANGELOG.md).

## Validator Upgrade Guide

An upgrade guide for validators is available at [Upgrade Guide v5.0](https://docs.regen.network/validators/upgrades/v5.0-upgrade.html).

## Developer Migration Guide

A migration guide for application developers is available at [Migration Guide v5.0](https://docs.regen.network/ledger/migrations/v5.0-migration.html).
