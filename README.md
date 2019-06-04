# Regen Ledger
![banner](docs/regen-network-image.jpg)

[![pipeline status](https://gitlab.com/regen-network/regen-ledger/badges/master/pipeline.svg)](https://gitlab.com/regen-network/regen-ledger/commits/master)
![GitHub issues](https://img.shields.io/github/issues/regen-network/regen-ledger.svg)
[![GitHub issues by-label](https://img.shields.io/github/issues/regen-network/regen-ledger/good%20first%20issue.svg)](https://github.com/regen-network/regen-ledger/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)
[![codecov](https://codecov.io/gh/regen-network/regen-ledger/branch/master/graph/badge.svg)](https://codecov.io/gh/regen-network/regen-ledger)
[![GoDoc](https://godoc.org/github.com/regen-network/regen-ledger?status.svg)](http://godoc.org/github.com/regen-network/regen-ledger)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/regen-network/regen-ledger)](https://goreportcard.com/report/github.com/regen-network/regen-ledger)
[![Gitter chat](https://badges.gitter.im/regen-network/regen-ledger.png)](https://gitter.im/regen-network/regen-ledger "Gitter chat")


A distributed ledger for ecology built on top of the 
[cosmos-sdk](http://github.com/cosmos/cosmos-sdk).

## Getting Started

See [Getting Started](docs/getting_started.md).

## Core Features

Regen Ledger aims to provides the following core features:
* a database of ecological state and change of state claims that spans both
on and off-chain data sources ([Ecological State Database](#ecological-state-database))
* mechanisms for automating the assessment of ecological state, making payments,
and the issuance of tokens and credits ([Compute Functions and Ecological Contracts](#compute-functions-and-ecological-contracts))
* infrastructure for issuing ecologically-backed asset tokens and credits ([Tokens and Credits](#tokens-and-credits))

This project is under heavy development and as result the above features are
implemented to varying degrees of completeness.

### Ecological State Database

One of the core functionalities of Regen Ledger is providing a structured
database of claims regarding ecological state and change of state. A claim is
made up of a few very basic pieces of information:
- the geo-polygon of the portion of the Earth being referred to,
- what is being claimed about this geographical region,
- who is making the claim, and
- any supporting evidence the claimant would like to associate with their claim
 
The actual data for claims can be stored on or off the Regen Ledger blockchain. 
In order to make claim data publicly available to the whole world, it can
be stored directly on the blockchain. In order to keep some or all of the
data private, it can be stored off-chain but "tracked" on-chain by
providing a cryptographic hash and URL, as well as possibly some metadata about
the claim. 

The facilities for storing data on-chain and tracking data off-chain
are managed by Regen Leder's [data](https://godoc.org/github.com/regen-network/regen-ledger/x/data)
and [geo](https://godoc.org/github.com/regen-network/regen-ledger/x/geo) modules.
In order to make it easy to write software that can automatically reason about
claim data, the schemas for all such data must be registered with Regen
Ledger's [schema](https://godoc.org/github.com/regen-network/regen-ledger/x/schema)
module and all submitted data must conform to these schemas. The actually
signing of claims is managed by the [claim](https://godoc.org/github.com/regen-network/regen-ledger/x/claim)
module.

Regen Ledger aims to provide built-in support for indexing claim data in
both the [PostgreSQL](https://www.postgresql.org)/[PostGIS](https://postgis.net)
database and the [Apache Jena](https://jena.apache.org)
RDF data store so that this data can be queried easily and used in compute
functions and contracts.

### Compute Functions and Ecological Contracts

Regen Ledger aims to provide a framework for executing compute functions that
take as input Regen Ledger's ecological state database as well as other "well-known"
public data sources, such as satellite imagery from NASA and ESA. This framework
will define:
- how compute functions can uniformly access private, off-chain data
given appropriate permissions
- how compute functions should be written and executed to ensure that results
are reproducible
- how computers that are executing compute functions (called oracles) should 
interact with Regen Ledger in order to have results stored back into the
ecological state database

This functionality will be managed by the oracle module and described in more detail there.

Ecological contracts in Regen Ledger are modelled as state machines that effectively "observe"
the ecological state database for certain conditions and which execute certain
actions when those conditions are met. For instance, a contract could
be written to make a payment to a farmer at the end of the year if the ecological
state database included claims from a reputable source that the farmer had
used certain practices like cover cropping. Or a contract could be setup as
effectively a "land trust" for a forest that accumulates credits while it
remains forested but has them slashed whenever a deforestation event is tracked
in the ecological state database.

The functionality for ecological contracts will be managed by the contract module.

### Tokens and Credits

In addition to allowing for payments using existing tokens, Regen Ledger will
allow for the creation of custom ecosystem tokens and credits whose issuance
can be controlled directly by ecological contracts.

## Testnet Status

See https://github.com/regen-network/testnets.
<br />
<br />
<br />
And of course there is this, from Mary Oliver, who recently passed away. <br />
We love you! <br />
<br />
## Sleeping in the Forest<br />
<br />
I thought the earth remembered me,<br />
she took me back so tenderly,<br />
arranging her dark skirts, her pockets<br />
full of lichens and seeds.<br />
I slept as never before, a stone on the river bed,<br />
nothing between me and the white fire of the stars<br />
but my thoughts, and they floated light as moths<br />
among the branches of the perfect trees.<br />
All night I heard the small kingdoms<br />
breathing around me, the insects,<br />
and the birds who do their work in the darkness.<br />
All night I rose and fell, as if in water,<br />
grappling with a luminous doom. By morning<br />
I had vanished at least a dozen times<br />
into something better.<br />
<br />
from Sleeping In The Forest by Mary Oliver<br />
© Mary Oliver<br />
