# Core Functionality

## Cosmos SDK Background

Regen Ledger is built ontop of the Cosmos SDK. Much of the language and usage patterns when interacting with Regen Ledger follow directly from usage patterns and architecture laid out in the Cosmos SDK.

For more information on Cosmos's work, and what it means to build an "application specific blockchain", the [Cosmos SDK Docs](https://docs.cosmos.network/master/intro/overview.html#what-are-application-specific-blockchains) is a great place to start.

## Regen Ledger Overview

Regen Ledger is a single application binary that:
- Runs a fully functioning node in the public proof-of-stake Regen Network
- Stores application state locally, using an [IAVL Tree](https://github.com/cosmos/iavl)
- Exposes an API server with both gRPC and REST interfaces for querying blockchain state and sending transactions
- Exposes a command line interface for account creation and key management as well as for querying blockchain state and sending transactions

Regen Ledger has two custom modules in development that support our desired functionality for ecological data, claims and credits.

### Ecocredit Module

The **Ecocredit Module** is our module for managing the issuance, trading, and retiring of credits pertaining to verifiable changes in ecosystem health.

Initially, this module will be used for Regen Network's own [CarbonPlus Grasslands](https://regen-registry.s3.amazonaws.com/GHG+and+Co-Benefits+in+Grazing+Systems+Credit+Class.pdf) credit, but we've designed this module first and foremost to be an open platform for all credit designers - something like an [EIP721](https://eips.ethereum.org/EIPS/eip-721) token standard for ecosystem service credits.

Check out [Ecocredit Module Overview](./modules/ecocredit/README.md) for more information about the ecocredit module.

### Data Module

::: tip COMING SOON
The data module is under active development and will soon be made available using the experimental build of Regen Ledger and on [Hambach Testnet](/getting-started/live-networks.md#hambach-testnet).
:::

High quality and verifiable ecological data is an essential component in any modern ecosystem service marketplace. The **Data Module** is intended to sit alongside the Ecocredit Module, serving as a generic repository for more complex metadata pertaining to a credit batch or an ecosystem service project.

The basic functionality of the data module includes storing, timestamping, and digitally signing data on Regen Ledger.

Check out [Data Module Overview](./modules/data/README.md) for more information about the data module.

## Additional Functionality

Supplemental to the core featureset described, Regen Ledger has out-of-the-box support for the creation of DAOs, multi-sig wallets, and smart contracting capabilities. These more complex features are enabled through an additional set of modules from the Cosmos ecosystem:

- **Groups Module** – allowing nested accounts, or subkeys, with custom voting schemas for message execution
- **CosmWasm Integration** – enabling WASM based smart contracts to live on Regen Ledger

### Groups Module

::: tip COMING SOON
A beta version of the group module is currently available using the experimental build of Regen Ledger and on [Hambach Testnet](/getting-started/live-networks.md#hambach-testnet).
:::

### CosmWasm

::: tip COMING SOON
CosmWasm support will soon be made available using the experimental build of Regen Ledger and on [Hambach Testnet](/getting-started/live-networks.md#hambach-testnet).
:::
