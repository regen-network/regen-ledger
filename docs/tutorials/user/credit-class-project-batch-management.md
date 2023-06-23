# Credit Class, Project, and Batch Management

This tutorial covers the creation and management of credit classes, projects, and batches using the [command-line interface (CLI)](../../ledger/infrastructure/interfaces.html#command-line-interface). This tutorial will demonstrate with data standards and practices used by Regen Network Development for [Regen Registry Program](https://library.regen.network/).

For information about creating and managing credit classes, projects, and credit batches using the [Regen Marketplace](https://app.regen.network/) application, see [Regen Network Guidebook](https://guides.regen.network/guides/regen-marketplace).

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

To create classes, projects, and batches on chain using the `ecocredit` module, we first need to know what data we are appending to each each class, project, and batch.

Credit classes, projects, and batches are stored as objects in on-chain application state. Each object has a `metadata` field that can be any arbitrary `string` with a maximum length of `256`.

The maximum length on the `metadata` field limits the amount of data stored in on-chain state, which can become expensive and slow down the network. It was also designed for content hashes, which enables data privacy using custom hashing algorithms or data resolvers with access control.

Regen Network Development uses an [Internationalized Resource Identifier (IRI)](../../modules/data/01_concepts#iri) for all credit classes, projects, and batches managed by Regen Registry. This is the recommended way for anyone wanting to leverage Regen Ledger for their own credit origination process and it ensures the supporting data is compatible with applications built and maintained by Regen Network Development.

### JSON-LD

The [Regen Registry Standards](https://github.com/regen-network/regen-registry-standards) repository includes the data schemas currently being used by Regen Network Development. Following this approach, we start with building a [JSON-LD](https://json-ld.org/) object for each credit class, project, and batch we intend to create.

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

Once we have the supporting data for a new credit class, project, and batch, we can then generate an IRI for each using the following command:

```sh
curl -X POST -d '<json-ld>' https://api.registry.regen.network/iri-gen
```

At this point we have our supporting data for a credit class, project, and batch, we also have an IRI for each piece of data, but we have not stored the data anywhere.

## Data Resolvers

If you are managing your own credit origination process, you need to host your own data. If you are not ready to figure out a solution for data storage but you are ready to create a credit class, project, and batch, you can skip to the [next section](#credit-class).

To make your data available in Regen Network Development applications, you also need to use the same IRI specification discussed in the previous section and to create a data resolver using the `data` module that will point applications to an endpoint where the data can be queried by IRI.

Regen Network Development applications leverage data resolvers to look up data hosted off chain using the IRI stored in the `metadata` field of each credit class, project, and batch.

### Define Resolver

To create a data resolver, run the following command:

```sh
regen tx data define-resolver <url>
```
With this command, we are creating a data resolver with a url of `<url>`. This is the url at which you would be hosting data. When provided an IRI (e.g. `<url>/<iri>`, the assumption is that an application can fetch the data in either a complete or partial form depending on how you manage privacy.

For more information about the command, run the following:

```sh
regen tx data define-resolver --help
```

Now that you have a data resolver, you can query it by id:

```sh
regen q data resolver <id>
```

The data for each credit class, project, and batch can now be registered to the resolver. The account that created the resolver is the only account that can register data to the resolver.

### Register Resolver

To register data to a resolver, run the following command:

```sh
regen tx data register-resolver ...
```

For more information about the command, run the following:

```sh
regen tx data register-resolver --help
```

Now that you have register data to a resolver, you can lookup the resolver using the IRI of the data.

```sh
regen q data resolver <id>
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

## Regen Mainnet

Everything you've done here can also be done using [Regen Mainnet](../../ledger/get-started/live-networks.md#regen-mainnet). All you need to do is update the configuration for the `regen` binary to use a different chain ID and node endpoint (you'll also need to own official REGEN tokens). See [Live Networks](../../ledger/get-started/live-networks.md) for configuration instructions.

## Regen Marketplace

You can now view your new credit class, project, and batch using a version of [Regen Marketplace](https://dev.app.regen.network/) connected to Redwood Testnet. You also might notice there is not much information on the pages but you have some new capabilities. Check out [Regen Network Guidebook](https://guides.regen.network/guides/regen-marketplace) to learn about managing credit classes, projects, and batches using the Regen Marketplace application.
