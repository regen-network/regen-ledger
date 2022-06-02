# Data Module

::: tip Coming Soon
The first version of the data module will be available in Regen Ledger v4.0.
:::

The following documentation provides a technical overview of the data module and is designed for developers building tools and services that interact with Regen Ledger.

For more information about how data is being stored for specific credit classes, projects, and credit batches through the Regen Registry Program, please see [Regen Registry Library](https://library.regen.network/).

## Overview

The data module enables content hashes for different types of data to be anchored with a timestamp, attested to by verifiers, and registered with a resolver. The primary use case for the data module is to provide verifiable supporting data for credit classes, projects, and credit batches that are created and managed using the ecocredit module.

Anchoring data is done using a secure content hash for either raw data (non-canonical) or graph data (conforms to the RDF data model) and provides a tamper-proof timestamp, effectively saying that the data exists at a certain point in time (also known as "secure timestamping"). When anchoring data, the sender of the transaction is not attesting to the veracity of the data.

Once data is anchored, verifiers can attest to the veracity of the data. Attesting to data is like signing a legal document, meaning that the attestor agrees to all conditions and to the best of their knowledge everything is true. Data can be attested to by multiple attestors and each attestation is secured with a tamper-proof timestamp indicating the date and time when the data was verified.

The data module also enables the creation and management of data resolvers. When a data resolver is created, the creator of the resolver becomes the admin. The admin can then register anchored data to the resolver, providing a list of content hashes that the data resolver claims to serve.

## Contents

1. **[Concepts](01_concepts.md)**
1. **[State](02_state.md)**
1. **[Msg Service](03_messages.md)**
1. **[Events](05_events.md)**
1. **[Types](06_types.md)**
1. **[Client](07_client.md)**

## Protobuf

For a complete list of the Protobuf definitions, see the following documentation:

- [regen.data.v1](https://buf.build/regen/regen-ledger/docs/main:regen.data.v1)
