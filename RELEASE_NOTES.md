# Regen Ledger v5.0.0

## New Features

The new features made available in Regen Ledger `v5.0.0` are as follows:

### Group Module

The `group` module enables the creation and management of on-chain multisig accounts and enables voting for message execution based on configurable decision policies...

The `group` module started out as a custom module built within the Regen Ledger repository and has since been added to Cosmos SDK as a core module refined and maintained by the Cosmos SDK core developers (which includes Regen Network Development team members).

### Interchain Accounts and Fees

The `host` module...

The `controller` module...

The `intertx` module...

The `fee` module...

### Message-Based Governance Proposals

...

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
