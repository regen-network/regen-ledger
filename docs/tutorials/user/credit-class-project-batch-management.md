# Credit Class, Project, and Batch Management

This tutorial covers the creation and management of credit classes, projects, and batches using the [command-line interface (CLI)](../../ledger/infrastructure/interfaces.html#command-line-interface). This tutorial will demonstrate with data standards and practices used by Regen Network Development for [Regen Registry Program](https://library.regen.network/).

For information about working with, and having your credit class approved by, Regen Registry, see [Regen Registry Credit Classes](https://library.regen.network/v/regen-registry-program-guide/credit-classes). For information about managing credit classes, projects, and batches using the [Regen Marketplace](https://app.regen.network/) application, see [Regen Network Guidebook](https://guides.regen.network/guides/regen-marketplace).

## Prerequisites

- [Install Regen](../../ledger/get-started)
- [Manage Keys](../../ledger/get-started/manage-keys.md)
- [Redwood Testnet](../../ledger/get-started/redwood-testnet.md)

## Recommended

- [Data Module Overview](../../modules/data)
- [Data Module Concepts](../../modules/data/01_concepts.md)
- [Ecocredit Module Overview](../../modules/ecocredit)
- [Ecocredit Module Concepts](../../modules/ecocredit/01_concepts.md)

## Metadata

To create classes, projects, and batches on chain using the `ecocredit` module, we first need to know what our supporting data looks like for each class, project, and batch we intend to create.

Credit classes, projects, and batches are stored as objects in on-chain application state. Each object has a `metadata` field that can be any arbitrary `string` with a maximum length of `256`. There are no additional restrictions on `metadata` but it was designed for content hashes.

Regen Network Development uses a custom [Internationalized Resource Identifier (IRI)](../../modules/data/01_concepts#iri) as the value of `metadata` for credit classes, projects, and batches created and managed by Regen Registry. If you are managing your own credit origination process, we recommend using the same IRI format, which will make your data readable by applications using Regen Network Development standards.

The IRI contains a content hash with embedded information about how the content hash was created and how the data was hashed. To generate IRIs for the `metadata` fields of a credit class, project, and batch, we first need to construct "graph" data using JSON-LD format. When we say "graph" data here, we mean data that conforms to the [Resource Description Framework (RDF)](https://www.w3.org/TR/rdf11-concepts/) data model. For more information about the relationship between RDF and JSON-LD, see [Relationship to RDF](https://www.w3.org/TR/json-ld/#relationship-to-rdf). 

### JSON-LD

The [Regen Registry Standards](https://github.com/regen-network/regen-registry-standards) repository includes the data schemas currently being used by Regen Network Development. Following this approach, we start with building a [JSON-LD](https://json-ld.org/) object for each credit class, project, and batch we intend to create.

The following templates can be used as a starting point. Feel free to add your own fields in addition to the ones provided. New fields added from other vocabularies should include references in the JSON-LD `@context`. For more information about `@context`, see [The Context](https://www.w3.org/TR/json-ld11/#the-context).

#### Classes

The following includes a template for a credit class:

```jsonld
{
  "@context": {
    "schema": "http://schema.org/",
    "regen": "https://schema.regen.network#",
    "xsd": "http://www.w3.org/2001/XMLSchema#",
    "schema:itemListElement": {
      "@container": "@list"
    },
    "regen:sectoralScope": {
      "@container": "@list"
    },
    "regen:offsetGenerationMethod": {
      "@container": "@list"
    },
    "regen:projectActivities": {
      "@container": "@list"
    },
    "schema:url": {
      "@type": "schema:URL"
    },
    "schema:image": {
      "@type": "schema:URL"
    }
  },
  "@type": "regen:CreditClass",
  "schema:name": "",
  "schema:description": "",
  "schema:url": "",
  "schema:image": "",
  "regen:sourceRegistry": {
    "schema:name": "",
    "schema:url": ""
  },
  "regen:sectoralScope": [
    ""
  ],
  "regen:offsetGenerationMethod": [
    ""
  ],
  "regen:approvedMethodologies": {
    "@type": "schema:ItemList",
    "schema:itemListElement": [
      {
        "schema:name": "",
        "schema:url": "",
        "schema:identifier": "",
        "schema:version": ""
      },
      {
        "schema:name": "",
        "schema:url": "",
        "schema:identifier": "",
        "schema:version": ""
      }
    ],
    "schema:url": ""
  },
  "regen:ecosystemType": [
    ""
  ],
  "regen:verificationMethod": "",
  "regen:projectActivities": [
    ""
  ]
}
```

#### Projects

The following includes a template for a project:

```jsonld
{
  "@context": {
    "schema": "http://schema.org/",
    "regen": "https://schema.regen.network#",
    "qudt": "http://qudt.org/schema/qudt/",
    "unit": "http://qudt.org/vocab/unit/",
    "xsd": "http://www.w3.org/2001/XMLSchema#",
    "regen:projectDesignDocument": {
      "@type": "schema:URL"
    },
    "schema:url": {
      "@type": "schema:URL"
    },
    "qudt:unit": {
      "@type": "qudt:Unit"
    },
    "qudt:numericValue": {
      "@type": "xsd:double"
    },
    "regen:projectStartDate": {
      "@type": "xsd:date"
    },
    "regen:projectEndDate": {
      "@type": "xsd:date"
    },
    "regen:offsetGenerationMethod": {
      "@container": "@list"
    }
  },
  "@type": "regen:Project",
  "schema:name": "",
  "schema:description": "",
  "regen:projectDesignDocument": "",
  "regen:projectDeveloper": {
    "schema:name": "",
    "schema:url": ""
  },
  "regen:projectType": "",
  "regen:projectActivity": {
    "schema:name": "",
    "schema:url": ""
  },
  "regen:offsetGenerationMethod": [
    ""
  ],
  "regen:projectSize": {
    "qudt:unit": "",
    "qudt:numericValue": 0
  },
  "regen:projectStartDate": "",
  "regen:projectEndDate": "",
  "schema:location": {
    "@context": {
      "type": "@type",
      "@vocab": "https://purl.org/geojson/vocab#",
      "coordinates": {
        "@container": "@list"
      },
      "bbox": {
        "@container": "@list"
      }
    },
    "type": "Feature",
    "place_name": "",
    "geometry": {
      "type": "Point",
      "coordinates": [
        0,
        0
      ]
    }
  },
  "regen:offsetProtocol": {
    "schema:name": "",
    "schema:url": "",
    "schema:version": ""
  }
}
```

#### Batches

The following includes a template for a credit batch:

```jsonld
{
  "@context": {
    "schema": "http://schema.org/",
    "regen": "https://schema.regen.network#",
    "regen:additionalCertifications": {
      "@container": "@list"
    },
    "regen:verificationReports": {
      "@container": "@list"
    },
    "schema:url": {
      "@type": "schema:URL"
    }
  },
  "@type": "regen:CreditBatch",
  "regen:verificationReports": [
    {
      "schema:url": ""
    }
  ],
  "regen:projectVerifier": {
    "schema:name": "",
    "schema:url": ""
  },
  "regen:additionalCertifications": [
    {
      "schema:name": "",
      "schema:url": ""
    },
    {
      "schema:name": "",
      "schema:url": ""
    }
  ]
}
```

### IRI Generation

Once we have the supporting data for a credit class, project, and batch, we can then generate an IRI for each using the following command:

```sh
curl -X GET -d '<json-ld>' -H "Content-Type: application/json" https://api.registry.regen.network/iri-gen
```

We now have the IRIs for the `metadata` field of each credit class, project, and batch.

## Data Resolvers

At this point in the tutorial, we have our supporting data for a credit class, project, and batch, and we have an IRI for each piece of data, but the data is only stored locally on our computers.

If you are managing your own credit origination process, you need to host your own data. If you are not ready to figure out a solution for hosting data but you are ready to create a credit class, project, and batch, then feel free to skip to the [next section](#credit-class) and come back later.

To make your data available in applications using Regen Network Development standards, you need to use the same IRI format we introduced in the previous section and create a data resolver using the `data` module that points applications to the hosted data when provided the IRI of the data.

For more information about the `data` module, see [Data Concepts](../../modules/data/01_concepts.md).

### Define Resolver

The following command will create a data resolver with a url of `[url]`. This is the URL at which you are hosting the data. You are claiming that given an IRI, an application can fetch the data at a specific host and path (e.g. `[url] + [iri]`) in complete or partial form.

For example, if data with IRI `regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf` is registered to resolver with URL `https://data.example.com/`, an application should be able to fetch the data at `https://data.example.com/regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf`.

To create a data resolver, run the following command:

```sh
regen tx data define-resolver [url]
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_data_define-resolver.md).

Now that you created a data resolver, you can query it by id:

```sh
regen q data resolver [id]
```

### Content Hash

Before we can anchor and register the data we created in the previous section, we need to convert our IRIs into JSON objects representing content hashes on chain.

To convert an IRI to a content hash, run the following command:

```sh
regen q data convert-iri-to-hash [iri]
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_query_data_convert-iri-to-hash.md).

### Register Resolver

The data for each credit class, project, and batch can now be registered to the resolver. The account that created the resolver is the only account that can register data to the resolver.

With the next command, you can register all your data at once. You can register credit class, project, and batch data using the content hashes we converted from IRIs in the last step.

To register data to a resolver, run the following command:

```sh
regen tx data register-resolver [resolver-id] [content-hashes-json]
```
For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_data_register-resolver.md).

Now that you registered data to the resolver, you can look up resolvers by IRI:

```sh
regen q data resolvers-by-iri [iri]
```

If the data is in fact hosted at the location the data resolver is pointing to, the data can now be fetched within applications using the same standards as Regen Network Development.

## Credit Class

Now that we have data for our class, project, and batch, we have IRIs for each, and we may or may not be hosting the data, we are ready to create a credit class.

A credit class represents a collection of projects and issuers whereby the projects within a credit class are following the same standards and the issuers are issuing credits on their behalf.

For more information about credit classes, see [Ecocredit Concepts](../../modules/ecocredit/01_concepts.md).

### Allowlist and Fee

Before we create a credit class, we need to check the [credit class creator allowlist](../../modules/ecocredit/01_concepts.md#credit-class-creator-allowlist) and that we have enough tokens for the credit class creation fee.

To check whether the credit class creator allowlist is enabled, run the following command:

```sh
regen q ecocredit class-creator-allowlist
```

To check allowed credit class creators (only if allowlist enabled), run the following command:

```sh
regen q ecocredit allowed-class-creators
```

To check the credit class creation fee, run the following command:

```sh
regen q ecocredit class-fee
```

You should see that the allowlist is disabled and therefore any account can create a credit class. You should also see the token amount and denomination needed to create a credit class, which we add to the next command using the `--class-fee` flag.

### Create Credit Class

We are now ready to create a credit class and to use the IRI we generated earlier in this tutorial. The IRI for the credit class can be used in the next command as the value of `[metadata]`.

To create a credit class, run the following command:

```sh
regen tx ecocredit create-class [issuers] [credit-type-abbrev] [metadata]
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_create-class.md).

Now that you created a credit class, you can look up the credit class by id:

```sh
regen q ecocredit class [class-id]
```

### Update Credit Class

Only the admin of a credit class can update the credit class. The account that created the credit class is assigned as the admin upon creation. After the credit class has been created, the admin can update the credit class, including the account assigned as the admin.

To update a credit class admin, run the following command:

```sh
regen tx ecocredit update-class-admin [class-id] [new-admin]
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_update-class-admin.md).

To update credit class issuers, run the following command:

```sh
regen tx ecocredit update-class-issuers [class-id]
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_update-class-issuers.md).

To update credit class metadata, run the following command:

```sh
regen tx ecocredit update-class-metadata [class-id] [new-metadata]
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_update-class-metadata.md).

## Project

Now that we have a credit class, we can create a project within the credit class. Only an account listed as a credit class issuer can create a project within the credit class.

A project represents a group or individual providing ecosystem services. The ecological outcomes of a project are measured and verified and then credits are issued by a credit class issuer.

For more information about projects, see [Ecocredit Concepts](../../modules/ecocredit/01_concepts.md).

### Create Project

We are now ready to create a project and to use the IRI we generated earlier in this tutorial. The IRI for the project can be used in the next command as the value of `[metadata]`.

To create a project, run the following command:

```sh
regen tx ecocredit create-project [class-id] [jurisdiction] [metadata]
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_create-project.md).

Now that you created a project, you can look up the project by id:

```sh
regen q ecocredit project [project-id]
```

### Update Project

Only the admin of a project can update the project. The account that created the project (i.e. the credit class issuer) is assigned as the admin upon creation. After the project has been created, the admin can update the project, including the account assigned as the admin.

To update a project admin, run the following command:

```sh
regen tx ecocredit update-project-admin [project-id] [new-admin]
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_update-project-admin.md).

To update project metadata, run the following command:

```sh
regen tx ecocredit update-project-metadata [project-id] [new-metadata]
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_update-project-metadata.md).

## Credit Batch

Now that we have a project, we can create a credit batch and issue credits. Only an account listed as a credit class issuer can create a credit batch for projects within the credit class.

A credit batch is a vintage of credits for a given project and, at the time of creation, credits are issued to a list of recipients, which can include accounts representing buffer pools.

For more information about credit batches, see [Ecocredit Concepts](../../modules/ecocredit/01_concepts.md).

### Create Batch

We are now ready to create a credit batch and to use the IRI we generated earlier in this tutorial. The IRI for the credit batch can be used in the next command within `[batch-json]`.

To create a batch, run the following command:

```sh
regen tx ecocredit create-batch [batch-json]
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_create-batch.md).

Now that you created a credit batch, you can look up the credit batch by denom:

```sh
regen q ecocredit batch [batch-denom]
```

You can also look up all credit balances for a credit batch by denom:

```sh
regen q ecocredit balances-by-batch [batch-denom]
```

### Update Batch

A credit batch can only be updated if the batch is "open" and only by the issuer. It is not common or recommended to create "open" credit batches
unless you are operating a bridge.

To update batch metadata, run the following command:

```sh
regen tx ecocredit update-batch-metadata [batch-denom] [new-metadata]
```

For more information about the command, add `--help` or check out [the docs](../../commands/regen_tx_ecocredit_update-batch-metadata.md).

## Conclusion

Congratulations! You have now created and updated a credit class, project, and batch with supporting data using the same standards and practices as Regen Network Development.

### Regen Mainnet

Everything you've done here can also be done using [Regen Mainnet](../../ledger/get-started/live-networks.md#regen-mainnet). All you need to do is update the configuration for the `regen` binary to use a different chain ID and node endpoint (you'll also need to own official REGEN tokens). See [Live Networks](../../ledger/get-started/live-networks.md) for configuration instructions.

### Regen Marketplace

You can now view your new credit class, project, and batch using a version of [Regen Marketplace](https://dev.app.regen.network/) connected to Redwood Testnet. You also might notice the pages are a bit empty, but you now have the ability to update them when logged in with Keplr. Check out [Regen Network Guidebook](https://guides.regen.network/guides/regen-marketplace) to learn about managing credit classes, projects, and batches using the Regen Marketplace application.
