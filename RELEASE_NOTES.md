# Regen Ledger v5.0

## New Features

The new features made available in Regen Ledger `v5.0` are as follows:

### DAO support via Group Accounts

Regen Ledger now includes the `group` module made available in Cosmos SDK `v0.46`. The `group` module started out as a custom module built within the Regen Ledger repository and has since been added to Cosmos SDK as a core module refined and maintained by the Cosmos SDK core developers. The `group` module enables the creation and management of on-chain multisig accounts and voting for message execution based on configurable decision policies.

What does this mean within the context of Regen Ledger functionality? All entities on a Regen Ledger chain can now be managed by a group account. For example, the role of the credit class admin can be assigned to a group account and the group can create and update decision policies for the execution of messages that are restricted to the role of the credit class admin. This example can be reapplied to all on-chain entities. In the ecocredit module, a group account could be assigned the role of credit class creator, credit class issuer, project admin, and/or basket curator, and in the data module, a group account could be assigned the role of resolver manager.

The `group` module also enables any individual or group of individuals to create and manage a group account independent of the predefined roles for on-chain entities. A prime example being community staking DAOs, which can now be managed by group accounts, therefore enabling the creation and management of decision policies around the execution of messages on behalf of a community staking DAO and the updating of members within the community staking DAO. Another use case of the `group` module is two-factor authentication whereby an individual uses a group account as their primary account that requires them to sign-off on the execution of messages using multiple devices (an account on each device).

For more information about the group module, check out the [group module documentation](https://docs.cosmos.network/v0.46/modules/group/).

### Message-Based Governance Proposals

Regen Ledger now includes the latest version of the `gov` module (`v1`) made available in Cosmos SDK `v0.46`. The previous version of the `gov` module (`v1beta1`) is still wired up in the application for backwards compatibility and to support application modules that have not yet been updated to utilize the latest version. The main feature in the latest version is message-based governance proposals. In combination with the `authz` module and the `group` module, message-based governance proposals unlock new opportunities for governance.

Governance proposals have historically been used for updating a specific set of governance parameters defined within each application module. Message-based governance proposals are similar in that they update governance parameters but those parameters are now more loosely defined, i.e. governance parameters are simply state and no longer need to be defined specifically as a governance parameter. Messages that update such state are implemented with restrictions on which account can call the message (similar to how a credit class admin is the only account with permission to update a credit class) but the account in this case is the `gov` module account.

With message-based governance proposals, any message can be submitted within the proposal to be executed on behalf of the `gov` module. Using the `authz` module alongside message-based governance proposals, it's now possible for a governance proposal to be submitted that would authorize another account to execute a specific message on behalf of the `gov` module account. The other account could be a group account representing a group of individuals that have expertise related to the state being managed. For example, a community staking DAO made up of a group of scientists could be granted authorization to add credit types and credit types could then be added via the voting process of the group.

All governance parameters within the `ecocredit` module have been updated to support message-based governance proposals. The `data` module does not include any governance parameters. All other application modules that are imported and that include governance parameters can be updated with what are now considered "legacy" proposals.

For more information about the gov module, check out the [gov module documentation](https://docs.cosmos.network/v0.46/modules/gov/).

### Interchain Accounts

Two new modules have been added to support interchain accounts. Interchain accounts enables cross-chain account management built on IBC. One of the modules is an application module built and maintained by the IBC team within the `ibc-go` repository (the `ica` module) and the other is an application module built and maintained by the RND team within the `regen-ledger` repository (the `intertx` module).

Interchain accounts are accounts controlled programmatically by counterparty chains via IBC packets. Unlike a traditional account, an interchain account does not have a private key and therefore does not sign transactions. The account is registered on a "host chain" via a "controller chain" and the controller chain sends instructions (IBC packets with Cosmos SDK messages) to the host chain that the interchain account then executes.

The `ica` module has two submodules (`host` and `controller`). The `host` submodule enables a Regen Ledger chain to act as a "host chain" and the `controller` submodule enables a Regen Ledger chain to act as a "controller chain". The `host` and `controller` submodules will not be enabled following the upgrade of an existing Regen Ledger chain and therefore each will require an on-chain governance proposal to enable. Which messages allowed to be executed by interchain accounts will also need to be added to an `allowed_messages` parameter in the `host` submodule with subsequent governance proposals.

The `intertx` module includes functionality to support the `controller` submodule, enabling the registration of interchain accounts and submitting transactions to be executed on a host chain.

For more information about interchain accounts, check out the [interchain accounts documentation](https://ibc.cosmos.network/main/apps/interchain-accounts/overview).

### Relayer Incentivization

The `fee` module is a self-contained [middleware](https://ibc.cosmos.network/main/ibc/middleware/develop.html) module that extends the base IBC application module. The fee module was designed as an incentivization mechanism to help cover the operational costs of running a relayer (i.e. running full nodes to query transaction proofs and paying for transaction fees associated with IBC packets).

There are three fees within the fee model, one for receiving the packet, one for acknowledging the packet, and one for timeouts. The fees are held in escrow until the packet is either successful or times out. In the case of a successful packet, the timeout fee will be reimbursed, and in the case of an unsuccessful packet, the receiving and acknowledging fee will be reimbursed.

The first version of the fee module only supports incentivization of new channels and existing channels will need to wait for additional functionality to support upgradeability. Using the fee middleware with IBC transactions is optional and acts more like a "tip". End users can manually incentivize IBC packets using the CLI and client developers can leverage the gRPC endpoints to integrate relayer fees within their application.

For more information about fee middleware, check out the [fee middleware documentation](https://ibc.cosmos.network/main/middleware/ics29-fee/overview).

## Additional Changes

### SDK Module Manager

Regen Ledger has historically been using a custom module manager within the application wiring. Regen Ledger `v5.0` migrates from the custom module manager to the latest version of the Cosmos SDK module manager and updates the `ecocredit` module and `data` module accordingly.

### gRPC Error Codes

A community member reported that the gRPC error codes for queries were not being reported correctly, in some cases the error code was `Unknown` and in other cases the error code did not match the standard gRPC error codes. Not all projects building on Regen Ledger will use the error messages provided by Regen Ledger and may choose to create their own error messages based on the error codes returned. To ensure we are providing a good developer experience, we have fixed and updated gRPC error codes to return the expected gRPC error codes.

### Experimental Build

Following Regen Ledger `v4.0`, and now with Regen Ledger `v5.0`, all experimental features that were being developed within the Regen Ledger codebase have been stablilized and included in the stable application build. The experimental application build option has therefore been removed. We will consider a separate release that includes CosmWasm that will be used to reboot Hambach Testnet if developers are wanting to experiment with the latest features alongside CosmWasm contracts, otherwise Hambach Testnet will continue running with support for CosmWasm contracts on the experimental build of Regen Ledger v4.0.

## Changelog

For a full list of changes since Regen Ledger `v4.1`, please see [CHANGELOG.md](./CHANGELOG.md).

## Validator Upgrade Guide

An upgrade guide for validators is available at [Upgrade Guide v5.0](https://docs.regen.network/validators/upgrades/v5.0-upgrade.html).

## Developer Migration Guide

A migration guide for application developers is available at [Migration Guide v5.0](https://docs.regen.network/ledger/migrations/v5.0-migration.html).
