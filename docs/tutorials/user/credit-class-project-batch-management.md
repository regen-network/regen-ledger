# Credit Class, Project, and Batch Management

This tutorial covers the creation and management of credit classes, projects, and batches using the [command-line interface (CLI)](../../ledger/infrastructure/interfaces.html#command-line-interface). This tutorial will demonstrate with the data model and practices used by Regen Network Development for [Regen Registry Program](https://library.regen.network/).

For more information about managing credit classes, projects, and credit batches using the [Regen Marketplace](https://app.regen.network/) application, see [Regen Network Guidebook](https://guides.regen.network/guides/regen-marketplace).

## Prerequisites

- [Install Regen](../../ledger/get-started)
- Configure Regen to use [Redwood Testnet](../../ledger/get-started/live-networks.md#redwood-testnet)
- Redeem testnet tokens from [Redwood Testnet](../../ledger/get-started/live-networks.md#redwood-testnet)

## Recommended

- [Data Module Overview](../../modules/data)
- [Data Module Concepts](../../modules/data/01_concepts.md)
- [Ecocredit Module Overview](../../modules/ecocredit)
- [Ecocredit Module Concepts](../../modules/ecocredit/01_concepts.md)

## Metadata

In order to create classes, projects, and batches on chain using the `ecocredit` module, we first need to have the supporting data for each class, project, and batch we intend to create.

Credit classes, projects, and batches are stored as objects in on-chain application state. Each object has a `metadata` field that can be any arbitrary `string` with a maximum length of `256`.

The maximum length on the `metadata` field limits the amount of data stored in on-chain state, which can become expensive and slow down the network. It was designed for content hashes, which enables data privacy using custom hashing algorithms or data resolvers with access control.

Regen Network Development uses an [Internationalized Resource Identifier (IRI)](../../modules/data/01_concepts#iri) for all credit classes, projects, and batches managed by Regen Registry. This is the recommended way for anyone wanting to leverage Regen Ledger for their own credit origination process and it enables more information to be displayed in applications built and maintained by Regen Network Development.

### JSON-LD

The [Regen Registry Standards](https://github.com/regen-network/regen-registry-standards) repository includes the data schemas currently being used by Regen Network Development. Using the Regen Network Development approach to metadata, we start with building a [JSON-LD](https://json-ld.org/) object for each credit class, project, and batch.

#### Classes

The following includes the expected fields for a credit class:

```jsonld
{
  "@context": ""
}
```

#### Projects

The following includes the expected fields for a project:

```jsonld
{
  "@context": ""
}
```

#### Batches

The following includes the expected fields for a credit batch:

```jsonld
{
  "@context": ""
}
```

### IRI Generation

Once we have our data for each, we can generate an IRI for the data using the following command:

```sh
curl -X POST -d '<json-ld>' https://api.registry.regen.network/iri-gen
```

### Data Hosting

If you are managing your own credit origination process, you'll need to host your own data. This can be as simple as a single get request and post request on a server connected to a database that stores JSON-LD data. To make your data available in applications developed by the community and Regen Network Development, you'll need to use the same IRI specification and create a data resolver using the `data` module that will point applications to the location of your data.

### Data Resolvers

Regen Network Development applications leverage data resolvers to look up data hosted off chain using the IRI stored in the `metadata` field of each credit class, project, and batch.

To create a data resolver, run the following command:

```sh
regen tx data define-resolver ...
```

### Anchoring Data

To anchor data, run the following command:

```sh
regen tx data anchor ...
```

### Registering Data

To register data to a resolver, run the following command:

```sh
regen tx data register-resolver ...
```

## Credit Class

Creating and updating a credit class...

### Create Credit Class

To create a credit class, run the following command:

```sh
regen tx ecocredit create-class ...
```

### Update Credit Class

To update a credit class, run the following command:

```sh
regen tx ecocredit update-class ...
```

## Project

Creating and updating a project...

### Create Project

To create a project, run the following command:

```sh
regen tx ecocredit create-project ...
```

### Update Project

To update a project, run the following command:

```sh
regen tx ecocredit update-project ...
```

## Credit Batch

Creating and updating a credit batch...


### Create Batch

To create a batch, run the following command:

```sh
regen tx ecocredit create-batch ...
```

### Update Batch

To update a batch, run the following command:

```sh
regen tx ecocredit update-batch ...
```

## Wrapping Up

You can view your credit class, project, and batch in a version of the Regen Marketplace application connected to Redwood Testnet. You also might notice there is not much information their. Check out the [Regen Network Guidebook](https://guides.regen.network/guides/regen-marketplace) for more information about managing credit classes, projects, and batches in the Regen Marketplace application.