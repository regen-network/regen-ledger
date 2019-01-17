# Regen Ledger Specification

## Overview

This document is an attempt to give a formal specification to Regen
Ledger and compute infrastructure that interacts with it. It is a
living document and will evolve over time as the system evolves. The
functionality specified may be at various stages of implementation,
this document will attempt to track those statuses but there may be
significant discrepancies until the system is stable. Community input
is definitely welcome. The primary forum for public comments is
[Gitlab Issues](https://gitlab.com/regen-network/regen-ledger/issues).

### Philosophy

This specification aims to balance pragmatism with idealism and find a
middle ground that will enable a system that
can be developed in a relatively short amount of time with today's
technology and that also satisfies all of the most important aims 
of the problem domain of ecological verification and contracts in a robust
and future resistant manner. That is we aim to build a system that
will still be relevant decades down the road (with of course graceful
upgrades along the way) but that will not take us years to build.

### Prerequisites

While this document attempts to give appropriate context where necessary,
the following resources may be useful for understanding the technology
that the system depends on in a general sense, and are those recommended
prerequisite reading:

* [Tendermint Docs](https://tendermint.com/docs/)
* [Cosmos SDK Docs](https://cosmos.network/docs/)
* [RDF 1.1 Concepts and Abstract Syntax](https://www.w3.org/TR/rdf11-concepts/)
* [Sparql 1.1 Query Language](https://www.w3.org/TR/sparql11-query/)

## XRN URI Scheme

Regen Ledger defines its own URI scheme with the `xrn` prefix for usage
in various modules. The semantics of this URI scheme will be defined
in each module as they arise. In the future, this information may be
consolidated into a single reference section.

## Data and RDF Indexes

[RDF](https://www.w3.org/TR/rdf11-concepts/) data is at the core of Regen
Ledger's data technologies. The rationale for this choice is documented
in a separate section [Why RDF?](#why-rdf).

This section documents what types of RDF data can be stored or tracked
on-chain via the `data` module and how this data should be indexed in an
RDF database and available through its Sparql query engine.

### Storing and Tracking Graphs and Dataset

Arbitrary RDF graphs or datasets can be stored or tracked on-chain. For
indexing reasons, the `data` module differentiates from graphs and
datasets although the procedure for storing them is similar.

#### Storing

To store data on-chain, a valid RDF dataset can be submitted in a
supported format (JSON-LD currently). The URI for a graph will be of the
format `xrn://<block-height>graph/<hash>` where `block-height` is the
hex-encoded block height at which the data was committed `<hash>` is the [URDNA2015](https://json-ld.github.io/normalization/spec/)
hash of the data. Likewise, datasets get the URI `xrn:dataset/<hash>`.
Note that all simple graphs will be accessible from the dataset URI's
as well as they are also valid datasets.

#### Tracking

Data stored off-chain can be stored on-chain by providing the
[URDNA2015](https://json-ld.github.io/normalization/spec/) hash
of the dataset and a URL from which the data should be permanently
accessible (TODO: define a mechanism to update this URL if necessary).
The data itself may require some valid form of authentication in order
to access. Access to private data in verification algorithms is discussed
in the verification section. Off-chain data can always be converted to
on-chain data at a later date by actually transacting the dataset
matching this hash onto the chain.

### Indexing and Sparql usage

For usage in [compute functions](#compute-functions) which are a core
part of Regen Ledger's verification methodology, RDF data is indexed
in an RDF quad-store database. This quad-store database must have
support for [Sparql 1.1 Query](https://www.w3.org/TR/sparql11-query/),
configurable [Sparql 1.1 Federation](https://www.w3.org/TR/2013/REC-sparql11-federated-query-20130321/)
and [GeoSparql](https://www.opengeospatial.org/standards/geosparql).

#### The default graph

The default graph of the Regen Ledger RDF data set is where on-chain
information regarding verifications lives and will be documented in a
later section.

#### Named graphs

Each graph that is stored on-chain should be available
as the named graph `xrn://<block-height>/graph/<hash>` in the RDF data set.

#### `SERVICE` endpoints

Each dataset that is stored on or off-chain should be accessible as the
dataset at the [Sparql 1.1. Federation](https://www.w3.org/TR/2013/REC-sparql11-federated-query-20130321/)
`SERVICE` endpoint `xrn://<block-height>/dataset/<hash>`. Compliant index databases should
overload the handling of service endpoint queries to either pull the dataset from the chain
if the dataset is stored on-chain or retrieve the dataset from the URL provided
on-chain which is supposed to permanently reference this data. In either case,
the Sparql engine which actually query these endpoints locally rather than deferring
the processing to a remote server. For off-chain data, the Sparql processor
should verify the hash of the dataset before completing the query and
throw a data inconsistency error if hash verification fails. Query authors
can ignore issues with bad or missing off-chain datasets using the
`SERVICE SILENT`  construct.

## Agents

## Verification Compute Functions

## Rationale

Verification compute functions are functions that are run off-chain
(at least for now) and should give deterministic results - i.e. give
the same output given the same input, while at the same time having
access to "well known world state" up to a certain point in time. This
"well known world state" access is one of the key differentiating factors
of the verification compute function framework. In functional programming
a pure function usually has no access to external resources like the file
system and HTTP resources. However, because we are building a system
with cryptographic integrity checks - in particular a blockchain which
also gives us a form of consensus around what is known when - we can
make this data available to compute functions. 

The following types of compute functions are defined:

### SPARQL Functions

SPARQL compute functions specify are get compute results from
from the data that is already stored or tracked on
Regen Ledger.

#### SPARQL CONSTRUCT

[SPARQL CONSTRUCT]() compute functions specify are used to generate a
new RDF from the RDF dataset index of data already stored or tracked on
Regen Ledger.

#### SPARQL ASK

[SPARQL ASK]() compute functions compute a boolean true/false value from
the index of data already stored or tracked on Regen Ledger.

### Docker Image

Docker compute functions are used to do more complex computations that
can depend both on data store or tagged on Regen Ledger as well as other
well known data sources like public satellite imagery.

## Verification

## Ecological State Protocols

## Identity Claims

## Rationale

### Why RDF?

