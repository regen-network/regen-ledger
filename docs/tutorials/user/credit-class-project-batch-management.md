# Credit Class, Project, and Batch Management

This tutorial covers the creation and management of credit classes, projects, and batches using the [command-line interface (CLI)](../../ledger/infrastructure/interfaces.html#command-line-interface). This tutorial will demonstrate with data standards and practices used by Regen Network Development for [Regen Registry Program](https://library.regen.network/).

For information about creating and managing credit classes, projects, and credit batches using the [Regen Marketplace](https://app.regen.network/) application, see [Regen Network Guidebook](https://guides.regen.network/guides/regen-marketplace).

## Prerequisites

- [Install Regen](../../ledger/get-started)
- [Redwood Testnet](../../ledger/get-started/live-networks.md#redwood-testnet)

## Recommended

- [Data Module Overview](../../modules/data)
- [Data Module Concepts](../../modules/data/01_concepts.md)
- [Ecocredit Module Overview](../../modules/ecocredit)
- [Ecocredit Module Concepts](../../modules/ecocredit/01_concepts.md)

## Metadata

To create classes, projects, and batches on chain using the `ecocredit` module, we first need to know what our supporting data looks like for each class, project, and batch we intend to create.

Credit classes, projects, and batches are stored as objects in on-chain application state. Each object has a `metadata` field that can be any arbitrary `string` with a maximum length of `256`. There are no additional restrictions on `metadata` but it was designed for content hashes.

Regen Network Development uses a custom [Internationalized Resource Identifier (IRI)](../../modules/data/01_concepts#iri) as the value of `metadata` for credit classes, projects, and batches created and managed by Regen Registry. If you are managing your own credit origination process, we recommend doing the same. If you use the same IRI generation method, your data will be readable by Regen Network Development applications.

The IRI contains a content hash with embedded information about how the content hash was created and how the data was hashed. To generate an IRI for the `metadata` field, we will first need to construct "graph" data for our credit classes, projects, and batches.

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

Once we have the supporting data for a credit class, project, and batch, we can then generate an IRI for each using the following command:

```sh
curl -X POST -d '<json-ld>' https://api.registry.regen.network/iri-gen
```

This is the IRI that we will use for `metadata` when creating our credit class, project, and batch.

## Data Resolvers

At this point in the tutorial, we have our supporting data for a credit class, project, and batch, and we have an IRI for each piece of data, but the data is only stored locally on our computers.

If you are managing your own credit origination process, you need to host your own data. If you are not ready to figure out a solution for hosting data but you are ready to create a credit class, project, and batch, then feel free to skip to the [next section](#credit-class) and come back later.

To make your data available in Regen Network Development applications, you need to use the same IRI generation method mentioned in the previous section and create a data resolver using the `data` module that points applications to the hosted data when provided the IRI of the data.

### Define Resolver

The following command will create a data resolver with a url of `[url]`. This is the url at which you are hosting the data. When provided an IRI (e.g. `[url] + [iri]`), the assumption is that an application will be able to fetch the data in complete or partial form depending on how you manage privacy.

To create a data resolver, run the following command:

```sh
regen tx data define-resolver [url]
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_data_define-resolver.md).

Now that you created a data resolver, you can query it by id:

```sh
regen q data resolver <id>
```

### Content Hash

Before we can anchor and register the data we constructed in the previous section, we must convert the IRIs we generated into a JSON object representing the content hash.

To convert an IRI to a content hash, run the following command:

```sh
regen q data convert-iri-to-hash <iri>
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_query_data_convert-iri-to-hash.md).

### Register Resolver

The data for each credit class, project, and batch can now be registered to the resolver. The account that created the resolver is the only account that can register data to the resolver.

With the next command you can register all your data at once. You can register credit class, project, and batch data using the content hashes generated in the last step.

To register data to a resolver, run the following command:

```sh
regen tx data register-resolver [resolver_id] [content_hashes_json]
```
For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_data_register-resolver.md).

Now that you registered data to the resolver, you can look up resolvers by IRI:

```sh
regen q data resolvers-by-iri [iri]
```

## Credit Class

Now that we have supporting data, we have our IRIs for each, and we may or may not be hosting and resolving our own data, we can move on to creating and managing a credit class.

A credit class represents a collection of projects and issuers whereby the projects are following the same standards and practices and the issuers are issuing credits on behalf of the projects.

For more information about credit classes, see [Ecocredit Concepts](http://localhost:8080/modules/ecocredit/01_concepts.html).

### Create Credit Class

To create a credit class, run the following command:

```sh
regen tx ecocredit create-class ...
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_create-class.md).

### Update Credit Class

To update a credit class admin, run the following command:

```sh
regen tx ecocredit update-class-admin ...
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_update-class-admin.md).

To update credit class issuers, run the following command:

```sh
regen tx ecocredit update-class-issuers ...
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_update-class-issuers.md).

To update credit class metadata, run the following command:

```sh
regen tx ecocredit update-class-metadata ...
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_update-class-metadata.md).

## Project

Creating and updating a project...

### Create Project

To create a project, run the following command:

```sh
regen tx ecocredit create-project ...
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_create-project.md).

### Update Project

To update a project admin, run the following command:

```sh
regen tx ecocredit update-project-admin ...
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_update-project-admin.md).

To update project metadata, run the following command:

```sh
regen tx ecocredit update-project-metadata ...
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_update-project-metadata.md).

## Credit Batch

Creating and updating a credit batch...


### Create Batch

To create a batch, run the following command:

```sh
regen tx ecocredit create-batch ...
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_create-batch.md).

### Update Batch

A credit batch can only be updated if the batch is "open". This is not common or recommended unless you are bridging credits to and from another chain or registry.

To update batch metadata, run the following command:

```sh
regen tx ecocredit update-batch-metadata ...
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_update-batch-metadata.md).

## Regen Mainnet

Everything you've done here can also be done using [Regen Mainnet](../../ledger/get-started/live-networks.md#regen-mainnet). All you need to do is update the configuration for the `regen` binary to use a different chain ID and node endpoint (you'll also need to own official REGEN tokens). See [Live Networks](../../ledger/get-started/live-networks.md) for configuration instructions.

## Regen Marketplace

You can now view your new credit class, project, and batch using a version of [Regen Marketplace](https://dev.app.regen.network/) connected to Redwood Testnet. You also might notice there is not much information on the pages but you have some new capabilities. Check out [Regen Network Guidebook](https://guides.regen.network/guides/regen-marketplace) to learn about managing credit classes, projects, and batches using the Regen Marketplace application.
