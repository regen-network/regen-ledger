# Core Functionality

## Cosmos SDK Background

As Regen Ledger is built ontop of the Cosmos SDK, much of the language and usage patterns when interacting with Regen Ledger follow directly from usage patterns and architecture laid out in the Cosmos SDK.

For more information on Cosmos's work, and what it means to build an "application specific blockchain", the [Cosmos SDK Docs](https://docs.cosmos.network/master/intro/overview.html#what-are-application-specific-blockchains) is a great place to start.

## Regen Ledger Overview

Regen Ledger is a single application binary that:
- Runs a fully functioning node in the public proof-of-stake Regen Network
- Stores application state locally, using an [IAVL Tree](https://github.com/cosmos/iavl)
- Exposes an API server with both gRPC and REST interfaces for querying blockchain state and sending transactions
- Exposes a command line interface for account creation and key management as well as for querying blockchain state and sending transactions

The initial implementation of Regen Ledger has two modules that support our desired functionality for ecological data, claims and credits:
- An **Ecocredit Module** for managing the issuance, trading, and retiring of credits pertaining to verifiable changes in ecosystem health
- A **Data Module** for storing, timestamping, and digitally signing data on Regen Ledger

### Ecocredit Module

::: warning TODO

This section not yet written.

:::

### Data Module

::: warning TODO

This section not yet written.

:::

## Additional Functionality

Supplemental to the core featureset described, Regen Ledger has out-of-the-box support for the creation of DAOs, multi-sig wallets, and smart contracting capabilities. These more complex features are enabled through an additional set of modules from the Cosmos ecosystem:
- **Groups Module** – allowing nested accounts, or subkeys, with custom voting schemas for message execution
- **CosmWasm Integration** – enabling WASM based smart contracts to live on Regen Ledger

### Groups Module

::: warning TODO

This section not yet written.

:::

### CosmWasm

::: warning TODO

This section not yet written.

:::
