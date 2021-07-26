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

The initial implementation of Regen Ledger has two modules that support 
our desired functionality for ecological data, claims and credits.


### Ecocredit Module

The **Ecocredit Module** is our module for managing the issuance, trading, and retiring
of credits pertaining to verifiable changes in ecosystem health. Initially,
this module will be used for Regen Network's own
[CarbonPlus Grasslands](https://regen-registry.s3.amazonaws.com/GHG+and+Co-Benefits+in+Grazing+Systems+Credit+Class.pdf)
credit, but we've designed this module first and foremost to be an open platform
for all credit designers, something like an ERC20 token standard for ecosystem
service credits.

Checkout the [Ecocredit Module Overview](./modules/ecocredit/) for a deeper
look into how the ecocredit module works in action.

### Data Module

High quality and verifiable ecological data is an essential component in any modern
ecosystem service marketplace. The Data Module is intended to sit alongside the Ecocredit
Module, serving as a generic repository for more complex metadata pertaining a credit batch
or ecosystem service project.

The basic functionality of the data module includes storing, timestamping,
and digitally signing data on Regen Ledger.

Checkout the [Data Module Overview](./modules/data/) for a deeper
look into how the data module works in action.

## Additional Functionality

Supplemental to the core featureset described, Regen Ledger has out-of-the-box support for the creation of DAOs, multi-sig wallets, and smart contracting capabilities. These more complex features are enabled through an additional set of modules from the Cosmos ecosystem:
- **Groups Module** – allowing nested accounts, or subkeys, with custom voting schemas for message execution
- **CosmWasm Integration** – enabling WASM based smart contracts to live on Regen Ledger

### Groups Module

::: tip COMING SOON

We'll be updating this section soon, when we have the groups module live on Regen
Ledger!

:::

### CosmWasm

::: tip COMING SOON

We'll be updating this section soon, when we have CosmWasm live on Regen
Ledger!

:::
