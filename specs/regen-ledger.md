# Regen Ledger Specification

## Overview

This document is an attempt to give a formal specification to Regen
Ledger and compute infrastructure that interacts with it. It is a
living document and will evolve over time as the system evolves. The
functionality specified may be at various stages of implementation,
this document will attempt to track those statuses but there may be
significant discrepancies until the system is stable. Community input
is definitely welcome. The primary forum for public comments is
[Gitlab Issues](https://github.com/regen-network/regen-ledger/issues).

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

## Regen URI Scheme

Regen Ledger defines its own URI scheme with the `regen` prefix for usage
in various modules. The semantics of this URI scheme will be defined
in each module as they arise. In the future, this information may be
consolidated into a single reference section.

## Data and RDF Indexes

[RDF](https://www.w3.org/TR/rdf11-concepts/) data (serialized as [JSON-LD](https://json-ld.org))
and [Sparql](https://www.w3.org/TR/sparql11-query/)
are at the core of Regen Ledger's data technologies.

### Rationale

The potential usage of RDF technologies in Regen Ledger originated from a
number of aims:

1. We would like to have namespaced identifiers for elements in data
schemas that can be shared between different schemas and optionally
have some semantic meaning defined elsewhere
2. The goal of creating an ecological data commons is a priority for
Regen Network

RDF technology and the movement which created it are closely enough
aligned with these goals that it was a technology considered from early on.

Beyond these initial desires for a data commons, there are a few other
more immediate needs have been identified:

1. We need to have a way to provide off-chain and on-chain compute functions
access to data that is already stored on-chain as well as remote data that
is tracked by hash on the blockchain
2. Ideally we could just query the blockchain for that data in a contract
3. We need some basic functional programming language that can be used to
compute new data to be stored on chain based on data already on chain. For
instance, we might need to convert tons of carbon in some verification claim
to coins to be released from escrow or minted. Or we might want to convert
tons of carbon to a regeneration score.

Sparql seems like a natural fit for 1 and 2, although there are other options
like SQL. Sparql may seem like a strange choice for 3 because it's not marketed
as a programming language. But upon studying our use cases more, it became clear
that most things like generating new ecological claims from existing claims or
computing a reward amount from ecological claims could be achieved by querying
the existing claims and generating some new claim. Sparql `CONSTRUCT` allows
us to query an existing RDF dataset and produce a new RDF graph from that. Sparql has
full expression support, geo-spatial support via GeoSparql, and also the ability
to query remote endpoints via federation which could be used to include off-chain
data tracked by hash on chain in a single query. While an ideal solution
might also include a more robust type checker and some other nice features, it
seems like Sparql checks most of the boxes to achieve these immediate important
needs for the system. In the future, a more custom built programming language,
which we have been calling Ceres, may reach maturity and fulfill a similar
role and even be adapted to work on-chain. For now, it appears existing RDF and
Sparql tools can be adapted to provide sufficient near term functionality.
In addition, there are some potential upsides that could arise out of the usage
of Sparql in terms of schema alignment - for instance Sparql engines can support
RDFS and OWL inferencing for out of the box things like sub-classes and sub-properties.

RDF and Sparql are not without their limitations. In particular, despite
being in existence for more than a decade these technologies have yet to
see widespread adoption which limits the existence of high quality tools
for using them. Numerous other criticism have been directed towards the
semantic web movement as a whole. Nonetheless, we feel like the technologies
as they are specified are a good enough fit and that the existing open source
tooling is mature enough to justify their usage for this application.

#### Potential Alternatives

One potential alternative to using Sparql and a query language in general is to rely
on verification requesters pushing data to protocols. For instance, if there is a
protocol that requires carbon verification, the verification requester could just
push a signed carbon verification claim and be done with it. If the protocol
code simply had access to the blockchain key-value store, it could do simple verification
that the claim exist on the blockchain. While this type of configuration could work for
a lot of cases, it prevents other desirable use cases like checking for the absence
of a valid counter claim in the blockchain registry of claims. Maybe this could be mitigated
by providing raw index access to the contract programming language, but going down along
this path we eventually end up reinventing the whole infrastructure for a database and
query language. If we could do this in a type-safe, formally verified way like we were
planning with Ceres, this might actually be desirable, but it will require significant
time and personnel investment and this needs to be planned along with other priorities.
Another option which could provide similar benefits is finding an appropriate RDF
schema language (possibly a subset of SHACL) which allows us to do sophisticated type
checking of a subset of Sparql queries. Within the RDF space, it's worth mentioning that
there are other logic languages such as RIF, Silk-2, Ergo, etc. that could prove
useful for our use cases. This all is worth exploring when we have more actual usage
experience. Another alternative that was considered was simply using a dialect of SQL
which has expression and query support. The downsides of SQL compared to Sparql are
that SQL is generally tied to a storage engine whereas Sparql can be applied to memory
and persistent datasets and patched together or filtered combinations of these, Sparql
supports late-bound external data sources via `SERVICE`, and the RDF data model is more
flexible and consistent across the whole domain (graphs all the way down) - SQL with JSON
support could work but is a bit less elegant.

### Storing and Tracking Graphs and Dataset

Arbitrary RDF graphs or datasets can be stored or tracked on-chain via
the `data` module. For indexing reasons, the `data` module differentiates between
graphs and datasets although the procedure for storing them is similar.

#### Storing

To store data on-chain, a valid RDF dataset can be submitted in a
supported format (JSON-LD currently). The URI for a graph will be of the
format `regen://<block-height>/graph/<hash>` where `block-height` is the
hex-encoded block height at which the data was committed `<hash>` is the [URDNA2015](https://json-ld.github.io/normalization/spec/)
hash of the data. Likewise, datasets get the URI `regen://<block-height>/dataset/<hash>`.
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
as the named graph `regen://<block-height>/graph/<hash>` in the RDF data set.

#### `SERVICE` endpoints

Each dataset that is stored on or off-chain should be accessible as the
dataset at the [Sparql 1.1. Federation](https://www.w3.org/TR/2013/REC-sparql11-federated-query-20130321/)
`SERVICE` endpoint `regen://<block-height>/dataset/<hash>`. Compliant index databases should
overload the handling of service endpoint queries to either pull the dataset from the chain
if the dataset is stored on-chain or retrieve the dataset from the URL provided
on-chain which is supposed to permanently reference this data. In either case,
the Sparql engine which actually query these endpoints locally rather than deferring
the processing to a remote server. For off-chain data, the Sparql processor
should verify the hash of the dataset before completing the query and
throw a data inconsistency error if hash verification fails. Query authors
can ignore issues with bad or missing off-chain datasets using the
`SERVICE SILENT`  construct.

### Efficient binary serialization of RDF

## Verification Framework

### Goals

The goal of the Regen Ledger verification framework is to provide a way to
come to consensus around world state, in particular the ecological health
of different parts of the world, while either relying minimally on
"trusted third parties" or having more trust and transparency into the
conclusions of trusted third parties. Note that eliminating trusted third
parties altogether is a non-goal of the verification framework as the effort
to do that may have too many unintended side effects. Our goal, rather, is
to increase transparency wherever possibly, while recognizing both the real-world
utility as well as risks of relying on other humans.

### Components

The verification framework as a whole relies on two components:
the [Oracle Function and Consensus Framework](#oracle-function-and-consensus-framework)
and [Agents](#agents). Oracle functions are a way of coming to conclusions
about world state, in particular ecological state just by relying on "well-known"
data. Agents are a way of creating organizations of individuals and groups
that can gain and lose rights to act as trusted third party verifiers at
various steps in verification protocols.

### Dealing with Uncertainty

## Oracle Function and Consensus Framework

## Rationale for Oracle Functions

Oracle functions are functions that are run off-chain
(at least for now) and should give deterministic results - i.e. give
the same output given the same input, while at the same time having
access to "well-known world state" up to a certain point in time. This
"well-known world state" access is one of the key differentiating factors
of the oracle function framework. In functional programming
a pure function usually has no access to external resources like the file
system and HTTP resources. However, because we are building a system
with cryptographic integrity checks - in particular a blockchain which
also gives us a form of consensus around what is known when - we can
make this data available to oracle functions. So the base world state
that is available to all oracle functions is the set of data stored on
Regen Ledger up to a certain height plus all of the remote, off-chain
data which has been stored by hash up to that height. In addition to
this data set, we also make available other data that is "well-known"
enough to be reasonably tamper resistant. This data includes the
satellite imagery collections produced by ESA and NASA, and may be expanded
to include other well-known "public" data sets.

## Achieving Determinism and Consensus

One of the primary goals of oracle functions is to be able to achieve
a deterministic result that all observers can agree is the one correct
result of running the function with the given inputs. There are many
things that could get in the way of this, such as:

1. Floating point indeterminism 
2. Improper function implementation that uses non-deterministic world state
leaked into compute environments such as random number generators, the system
clock, and unintentional access to the file system and/or network
3. Inconsistent access to remote resources tracked on the blockchain (i.e. some
oracle runners may have read access to those resources and others may not)
4. Willful misrepresentation on the part of oracle runners
5. Faulty indexing of blockchain state or faulty function execution
6. Hash pre-image attacks

For each of these cases, let's explore who is at fault (if anyone) and what can be
done:
1. Use of floating point math probably can't be avoided in ecological data science,
so we should take whatever precautions we can to minimize indeterminism, but realistically
we probably can't eliminate it (insert REFERENCES). Ultimately this is nobody's fault
and is something that needs to be dealt with as an explicit part of the consensus algorithm.
i.e. oracles must come to consensus around an agreed upon floating point result and to
be able to achieve this, protocol curators must ensure that functions that use floating
point math provide sufficient tolerance ranges
2. This is ultimately the fault of the protocol curator. There must be some mechanism
for coming to consensus that this is the issue and for dealing with the aftermath
(which probably in the ideal case results in fixing the underlying compute function,
but this may or may not be possible in all cases)
3. This is the fault of the verification requester assuming they have access to the data.
To deal with this, there must be a protocol for oracles to report which remote resources
they were not able to access and a way for verification requesters to re-run functions
which are inconsistent for this reason after they have either fixed remote access
permissions or availability, or instructed oracles to ignore certain inaccessible resources
consistently
4. There must be a protocol for identifying and dealing with malicious activity on
the part of oracle runners which results in them being banned from the system and
probably the seizure of some bond amount
5. This is the fault of the oracle runner, but is not necessarily malicious. It may
result in the slashing of an oracle bond, but does not necessarily result in system
banning unless it is consistently unresolved
6. For now we assume we can avoid this entirely by choice of a sufficiently robust hash
algorithm, but this assumption should be re-examined as quantum computers evolve

### Oracle Consensus Protocol

The oracle consensus protocol is a mechanism for coming to reasonable consensus
about the correct results of off-chain computations. A number of factors unique
to the use cases we are solving for are:

* Some computations in the system may be fairly intensive taking several minutes
  or even hours (geospatial data analysis)
* We do not need instant finality for most of these computational results. Ecological
  change of state occurs on the timescale of months and years, so a few extra minutes,
  hours or even days to achieve consensus around analysis results is usually reasonable
* The results of some calculations need to be public and stored back on the ledger
  and other results need to remain private, but some amount of tracking via hashes
  on the ledger is advantageous and increases trust
  
Outlined below is the proposed protocol for byzantine fault tolerant off-chain computations:

* An oracle pool refers to a set of oracles sharing common contracts around
  cost, bonding, and arbitration - these monetary details will be discussed separately.
  For the sake of the cryptographic algorithm, we assume that all oracles are bonded,
  have agreed to certain per computation payment terms, and have consented to have all disputes
  that can't be resolved computationally resolved by the named arbiter. We also
  assume that the compute function or verification protocol curator is also bonded
  and has agreed to the same named arbiter
* A verification protocol requesting computation will normally specify the minimum
  number of oracles that need to perform a computation. Note that due to the
  challenge process and challenge window, it may often be safe to set this value to 1.
  For this example let us assume that the minimum is 2 and that a third oracle is
  asked to perform the computation as well 10% of the time randomly
* When a computation is requested of an oracle, the first oracle to do the computation
  is chosen at random based on the modulus of a block hash
* Once the first oracle has completed the computation, they choose a random nonce
  which they keep private and then store the hash of the nonce appended to the
  computation result on the blockchain. This makes it hard for another oracle to
  brute force guess the computation result when the possible result set is small (for instance binary or integral)
* The second computation oracle is chosen randomly based on the result of the block hash
  where the first oracle tracks their result. The second oracle performs the computation
  and tracks it on the blockchain using the same nonce + hash method
* A third oracle may or may not be chosen randomly based on the block hash of the second
  oracle's result
* Once the initial oracle set has computed all results, they privately share the results
  of their computation and the nonces they used (either through secure back-channels or by public PGP messages).
  The oracles check both that the hashes stored on the blockchain knowing these nonces
  and also the computation results. Because a function may specify floating point tolerance ranges,
  the oracles may be checking that results are within the tolerance range as opposed to identical.
  Each oracle votes publicly on the blockchain about the correctness of the results of
  the other oracles
* If there is not 100% agreement amongst this initial oracle set, additional oracles will be
  selected until either there are 4 out of 5 or 7 out of 9 oracles that agree upon the result.
  If either 4 out of 5 or 7 out of 9 oracles concur on the result, it will be assumed
  that the oracles that disagreed are byzantine and will have their bonds slashed unless a
  successful appeal to the arbiter determines that indeterminism in the compute function
  was responsible for the discrepancy. If either 100% initial agreement or 4 out of 5 or 7
  out of 9 consensus cannot be achieved, the compute function or verification protocol 
  curator will be held as the responsible party and have their bond slashed as well as
  compute function suspended unless the arbiter determines upon appeal that in fact byzantine
  oracles were present
* The verification protocol depending on this computation will usually set a challenge
  window depending on the computational complexity and stakes. During the challenge window
  a third party observer may challenge the oracle's result and have the arbiter intervene.
  This involves posting a bond which may be broken if the arbiter decides against the challenger.
  
  This challenge window enables protocol authors to sometimes set the minimum compute oracle
  threshold to 1 for very complex computations
* Note that the protocol above is identical for cases where results are eventually made
  public or kept private

## Oracle Function Types

The following types of compute functions are defined:

### SPARQL Functions

SPARQL compute functions specify are get compute results from
from the data that is already stored and tracked on
Regen Ledger.

#### SPARQL CONSTRUCT

[SPARQL CONSTRUCT]() functions are used to generate new RDF graphs from the
data already tracked and stored on Regen Ledger.

#### SPARQL ASK

[SPARQL ASK]() functions compute a boolean true/false value from
the data already stored and tracked on Regen Ledger.

### Docker Image

Docker compute functions are used to do more complex computations that
can depend both on data store or tagged on Regen Ledger as well as other
well known data sources like public satellite imagery.

## Agents

## Ecological State Protocols

## Identity Claims

## Contracts

### Escrow Accounts

### Token Minting

### Ecosystem Health Endowments

