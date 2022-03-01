# Upgrade Overview

This document provides an overview of the upgrade process for software upgrades on [Regen Mainnet](../getting-started/live-networks.md#regen-mainnet) and [Redwood Testnet](../getting-started/live-networks.md#redwood-testnet). Instructions for each upgrade can be found at the following links:

- [Upgrade Guide v2.0](./v2.0-upgrade.md)
- [Upgrade Guide v3.0](./v3.0-upgrade.md)

## Software Upgrade

A software upgrade is required when breaking changes are introduced in a new version of [Regen Ledger](https://github.com/regen-network/regen-ledger). In order to upgrade a live network to a new version of software that introduces breaking changes, there must be an agreed upon block height at which blocks are no longer produced with the old version of the software and only with the new.

The expected version and the upgrade height are defined within a [software upgrade](https://docs.cosmos.network/master/modules/gov/01_concepts.html#software-upgrade) proposal. If the proposal passes, the chain will halt at the proposed upgrade height. At this point, node operators will need to stop running the current binary and start running the upgrade binary.

### In-Place Store Migrations

Regen Ledger leverages a feature introduced in Cosmos SDK v0.44 called ["In-Place Store Migrations"](https://docs.cosmos.network/master/core/upgrade.html). This feature prevents the need to stop the chain, manually perform migrations, and then restart the chain with an updated genesis file. Each upgrade binary defines the migrations that need to take place within an upgrade handler that gets called once the new binary has been started.

### Using Cosmovisor

We recommend node operators use [Cosmovisor](https://docs.cosmos.network/master/run-node/cosmovisor.html), a small process manager for running application binaries for Cosmos SDK-based blockchains. When the chain halts at the proposed upgrade height, Cosmovisor will automatically stop the current binary and start the upgrade binary. Node operators can prepare the upgrade binary ahead of time and then relax at the time of the upgrade. Cosmovisor also includes an option to automatically download the upgrade binary at the time of the upgrade and an option to backup the data before performing migrations.

## Testnet Upgrade

Each upgrade on Regen Mainnet will be preceded by an upgrade on Redwood Testnet. The upgrade on Redwood Testnet provides validators an opportunity to go through the upgrade process on a live network before going through the same process on Regen Mainnet. Each upgrade will be rigorously tested on short-lived test networks prior to performing the upgrade on a live network.

The voting period for Redwood Testnet is currently set to `86400s` (1 day) and the voting period for Regen Mainnet is currently set to `1209600s` (14 days). After the upgrade proposal for Regen Mainnet has been submitted, an upgrade proposal will be submitted on Redwood Testnet.
