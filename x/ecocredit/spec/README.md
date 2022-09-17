# Ecocredit Module

::: tip Now Available
The first version of the ecocredit module was made available in Regen Ledger v2.0.
:::

The following documentation provides a technical overview of the ecocredit module and is designed for developers building tools and services that interact with Regen Ledger.

For more information about specific ecosystem service credits and methodologies being issued and developed through the Regen Registry Program, please see [Regen Registry Library](https://library.regen.network/).

## Overview

The ecocredit module is divided into three separate components, including "base functionality" and two submodules, the "basket submodule" and the "marketplace submodule".

### Base Functionality

::: tip Now Available
The base functionality was made available in Regen Ledger v2.0. Added support for projects was made available in Regen Ledger v4.0.
:::

The base functionality of the ecocredit module enables the creation and management of credit types, credit classes, and projects, as well as the issuance, transferring, and retirement of ecosystem service credits (i.e. carbon credits, biodiversity credits, soil health credits, etc.).

Credit classes are designed to support a variety of ecosystem service credits, and can be created and managed by individuals, groups, or organizations. Each credit class is associated with a credit type that represents the primary unit of measurement for the methodologies defined within the credit class (i.e. carbon, biodiversity, etc.). Credit types can only be added through on-chain governance.

Ecosystem service credits are issued through a batch issuance process where only approved issuers for a given credit class can issue credits from that credit class. Each credit batch is associated with a project representing the on-the-ground project implementing the methodologies defined within the credit class. Credit batches are unique in that each credit batch has a start and end date representing the period in which the positive ecological outcomes were measured.

Credits issued within a credit batch are only fungible with credits from the same batch. Credits can be issued in a tradable or retired state. Credits issued in a tradable state can be transferred or retired by the owner. Retiring a credit is permanent and implies the credit is being consumed as an offset.

### Basket Submodule

::: tip Now Available
The basket submodule was made available in Regen Ledger v3.0.
:::

The basket submodule enables the creation and management of baskets. Ecosystem service credits that meet a specific criteria defined by a basket can be put into the basket in return for tokens that are fully fungible with other tokens from the same basket. Basket tokens are minted using the bank module from Cosmos SDK and are therefore compatible with IBC, allowing ecosystem service credits to be represented as tokens that are transferable across blockchains. Ecosystem service credits can then be taken out of the basket upon returning the tokens to the basket.

### Marketplace Submodule

::: tip Now Available
The marketplace submodule was made available in Regen Ledger v4.0.
:::

The marketplace submodule enables the creation and management of sell orders for ecosystem service credits and the purchasing of those credits through direct buy orders. Credits can only be listed for approved token denoms that are managed through an on-chain governance process. The credits for each sell order are held in escrow until the sell order either expires, is cancelled by the owner, or purchased via a direct buy order. The owner of a sell order can update or cancel the sell order at any time.

## Contents

1. **[Concepts](01_concepts.md)**
1. **[State](02_state.md)**
1. **[Msg Service](03_messages.md)**
1. **[Query Service](04_queries.md)**
1. **[Events](05_events.md)**
1. **[Types](06_types.md)**
1. **[Client](07_client.md)**

## RFCs

- [RFC-001: Ecocredit Module](https://github.com/regen-network/regen-ledger/blob/main/specs/rfcs/001-ecocredit-module)
- [RFC-002: Basket Functionality](https://github.com/regen-network/regen-ledger/blob/main/specs/rfcs/002-basket-functionality)

## Protobuf

For a complete list of the Protobuf definitions, see the following documentation:

- [regen.ecocredit.v1](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1)
- [regen.ecocredit.basket.v1](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1)
- [regen.ecocredit.marketplace.v1](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1)
