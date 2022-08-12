# Core Functionality

## Built on Cosmos SDK

Regen Ledger is a proof-of-stake blockchain application built on top of [Cosmos SDK](https://github.com/cosmos/cosmos-sdk).

Much of the language and usage patterns when interacting with Regen Ledger follow directly from the usage patterns and architecture laid out in Cosmos SDK.

For more information about Cosmos SDK and what it means to build an "application specific blockchain", the [Cosmos SDK Documentation](https://docs.cosmos.network/master/intro/overview.html#what-are-application-specific-blockchains) is a great place to start.

## Regen Ledger Overview

Regen Ledger is a single application binary that:
- Runs a fully functioning node in a public proof-of-stake network
- Exposes an API server with both gRPC and REST interfaces for querying blockchain state and sending transactions
- Exposes a command line interface for account creation and key management as well as for querying blockchain state and sending transactions

Regen Ledger has two custom modules that support our desired functionality for ecosystem service credits and ecological data claims.

### Ecocredit Module

The **Ecocredit Module** is our module for managing the issuance, trading, and retiring of credits pertaining to verifiable changes in ecosystem health.

Check out [Ecocredit Module Overview](../../modules/ecocredit/README.md) for more information about the ecocredit module.

### Data Module

The **Data Module** is used alongside the Ecocredit Module, enabling the anchoring of ecological data, attesting to the veracity of anchored data, and registering anchored data to a resolver.

Check out [Data Module Overview](../../modules/data/README.md) for more information about the data module.
