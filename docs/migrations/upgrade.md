# Upgrade Overview

This document provides a general overview of the upgrade process for software upgrades on [Regen Mainnet](../getting-started/live-networks.md#regen-mainnet) and [Redwood Testnet](../getting-started/live-networks.md#redwood-testnet).

Each upgrade on Regen Mainnet will be preceded by an upgrade on Redwood Testnet. The upgrade on Redwood Testnet provides node operators an opportunity to go through the upgrade process on a live network before going through the same process on Regen Mainnet.

*Note: Redwood Testnet uses the same configuration as Regen Mainnet aside from the voting period; the voting period for Redwood Testnet is currently set to `86400s` (1 day) whereas the voting period for Regen Mainnet is currently set to `1209600s` (14 days). Each upgrade will be rigorously tested on temporary short-lived test networks prior to performing the upgrades on Redwood Testnet and Regen Mainnet.*

## Upgrade Overview

A software upgrade is required when breaking changes are introduced in a new version of [Regen Ledger](https://github.com/regen-network/regen-ledger). In order to upgrade a live network to a new version of software that introduces breaking changes, there must be an agreed upon block height at which blocks are no longer produced with the old version of the software and only with the new.

The expected version and the upgrade height are defined within a [software upgrade](https://docs.cosmos.network/master/modules/gov/01_concepts.html#software-upgrade) proposal. If the proposal passes, the chain will halt at the proposed upgrade height. At this point, node operators will need to stop running the current binary and start running the upgrade binary.

We recommend node operators use [Cosmovisor](https://docs.cosmos.network/master/run-node/cosmovisor.html), which is a process manager for running a blockchain's application binary. When the chain halts at the proposed upgrade height, Cosmovisor will automatically stop the current binary and start the upgrade binary. Using Cosmovisor, node operators can prepare the upgrade binary ahead of time and then relax at the time of the upgrade.

### In-Place Store Migrations

For software upgrades, Regen Ledger leverages a feature introduced in Cosmos SDK v0.44 called ["In-Place Store Migrations"](https://docs.cosmos.network/master/core/upgrade.html). This feature prevents the need for node operators to restart the chain using an updated genesis file. Regen Ledger includes a `registerUpgradeHandlers()` method that registers handlers for performing an upgrade's necessary state migrations in-place.

