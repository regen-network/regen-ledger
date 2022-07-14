# Regen Ledger v4.0.0

## New Features

The new features made available in Regen Ledger `v4.0.0` are as follows:

### On-Chain Projects

We've updated the core functionality of the ecocredit module to support on-chain projects. This means on-the-ground project developers providing ecosystem services will now be represented as on-chain entities and information about each project will be stored separate from the credit batch and each credit batch will store the associated project ID.

In this initial implementation, projects exist within credit classes and only an approved issuer of a given credit class can create a project associated with that credit class. The issuer of the credit class becomes the project admin when the project is created and can then reassign the admin role to any address.

For more information about projects, check out the [ecocredit module documentation](https://docs.regen.network/modules/ecocredit/).

### Marketplace Functionality

The ecocredit module includes a new marketplace submodule that supports a simple storefront model for creating sell orders and purchasing credits directly from those sell orders. When a sell order is created, the credits being sold are held in escrow. The default behavior is to have credits auto-retired upon sale but the seller has the option to disable auto-retirement and the buyer has the option to disable auto-retirement if the seller has disabled auto-retirement.

Credit owners can only list credits for sale with a token denom that is on an "allowed denom" list specific to the marketplace and controlled through on-chain governance. The allowed denom list will be empty at the time of the upgrade and the community will be able to submit proposals to add allowed denoms following the upgrade. See this discussion for more information.

For more information about marketplace functionality, check out the [ecocredit module documentation](https://docs.regen.network/modules/ecocredit/).

### Ecological Data Services

The first version of the data module has arrived and it supports the ability to anchor data on chain, attest to the veracity of anchored data, and to define a data resolver and register anchored data to that resolver. Anchoring data (also known as "secure timestamping") does not store the data on chain but rather a content hash of the data alongside a timestamp that represents the time at which the data was anchored. If the data is altered in any way, the content hash will be different and the data will need to be anchored again as a separate entry.

The initial use case for the data module will be to anchor data specific to each credit class, project, and credit batch, including but not limited to methodologies for credit classes, baseline monitoring reports for projects, and monitoring reports for credit batches. Anchoring data generates a unique deterministic identifier (an IRI) that will then be stored in the metadata field for each credit class, project, and credit batch. The data can optionally be registered to a resolver for convenient public (or private/verified) lookups and attested to as a means of verification.

The intention of this design is to allow for those anchoring datasets to have control over the privacy of their data. Credit issuers and project admins can leverage Regen Ledger for data anchoring and attestation, while keeping the raw datasets associated with those IRIs private if they choose. In a future release, we intend to support merklized hash formats, which would enable individual elements of datasets to be selectively disclosed to the public or to a specific buyer.

For more information about the data module, check out the [data module documentation](https://docs.regen.network/modules/data/).

### Cross-Chain Credits

Over the past few months, we have been working alongside the Toucan team to develop a bridge service that will enable bridging ecosystem service credits to/from the Polygon blockchain. The initial use case of the bridge service will be to bridge Toucan's TCO2 tokens to Regen Ledger for use in our upcoming NCT basket.

In support of these efforts, we have added functionality in Regen Ledger v4.0 to support dynamic batch minting that enables bridged assets from the same vintage to be minted to a pre-existing credit batch. Each credit batch will be "sealed" by default so that credit batches with credits issued natively on Regen Ledger can remain immutable.

When credits are bridged from Regen to Polygon, the credits will be cancelled, indicating that the credits have moved to another chain or registry. The functionality to support bridging assets will be included in Regen Ledger v4.0 but the bridge service itself will be launched separately.

## Improvements

### Improved Storage

Regen Ledger v4.0 makes use of an ORM storage model implemented within the orm package within Cosmos SDK that acts as an abstraction layer over the existing KV store. The orm package enables the creation of database tables with primary and secondary keys. This abstraction layer provides support for efficient lookups and will improve the velocity of feature development.

### Improved API Naming

Regen Ledger v4.0 includes a significant number of minor API changes intended to provide more consistent naming throughout the project and to provide an overall better user experience. The API is defined in proto files that are now available on Buf Schema Registry.

## Additional Changes

### Credit Batch Denoms

Adding support for on-chain projects required updating the format of the credit batch denom to include the project ID. The credit batch denom was previously formatted to include the credit type abbreviation, the credit class ID, the start and end dates for the monitoring period, and the credit batch sequence number scoped to the credit class. The credit batch denom is now formatted to include the project ID and the credit batch sequence number is now scoped to the project.

An example of a credit batch in Regen Ledger v3.0:

```
C01-20200101-20210101-001
```

*The first credit batch from the first credit class of the "C" credit type with a start date of January 1st 2020 and an end date of January 1st 2021 (credit class id "C01" and batch sequence "001").*

An example of a credit batch in Regen Ledger v4.0:

```
C01-001-20200101-20210101-001
```

*The first credit batch from the first project of the first credit class of the "C" credit type with a start date of January 1st 2020 and an end date of January 1st 2021 (credit class id "C01", project id "C01-001", and batch sequence "001").*

## Changelog

[link]

## Validator Upgrade Guide

[link]

## Developer Migration Guide

[link]
