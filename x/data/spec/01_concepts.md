# Concepts

### IRI

An Internationalized Resource Identifier (IRI) is used within the data module to identify a piece of data. The IRI for a piece of data is always unique and generated deterministically. The IRI contains the full [content hash](#content-hash) enabling support for stateless conversions between IRI and content hash.

The IRI for a [raw content hash](#raw-content-hash) and a [graph content hash](#graph-content-hash) follow the same format but differ in how the content hash is encoded to accommodate for different properties. Each IRI includes a prefix, base58 encoded data (the encoded content hash), and a file extension.

The format for an IRI:

```
regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf
```

The pattern for a raw content hash:

```
regen:{base58check(concat(byte(0x0), byte(digest_algorithm), hash))}.{media_type extension}
```

The pattern for a graph content hash:

```
regen:{base58check(concat(byte(0x1), byte(canonicalization_algorithm), byte(merkle_tree), byte(digest_algorithm), hash))}.rdf
```

### Content Hash

A content hash is a hash-based content identifier for a piece of data. A content hash can either be of type [raw](#raw-content-hash) or [graph](#graph-content-hash). A content hash defines the hash (the content hash itself) and the digest algorithm used to generate the hash. Each type defines additional properties specific to its type.

#### Raw Content Hash

A raw content hash specifies "raw" data that does not use deterministic, canonical encoding. Users of raw content hashes must maintain a copy of the hashed data that is preserved bit by bit. In addition to defining the hash (the content hash itself) and the digest algorithm, a raw content hash also defines the media type (e.g. TXT, JSON, CSV, XML, PDF, etc.). For a complete list of the supported media types for a raw content hash, see [RawMediaType](https://buf.build/regen/regen-ledger/docs/main:regen.data.v1#regen.data.v1.RawMediaType).

#### Graph Content Hash

A graph content hash specifies "graph" data that conforms to the [RDF data model](https://www.w3.org/TR/rdf11-concepts/) and therefore uses deterministic, canonical encoding allowing implementations to choose from various formats for content hash encoding while maintaining the guarantee that the underlying canonical hash will not change. In addition to defining the hash (the content hash itself) and the digest algorithm, a graph content hash also defines the canonicalization algorithm and the type of merkle tree. In the current implementation, Universal RDF Dataset Canonicalization Algorithm 2015 (URDNA2015) is the only canonicalization algorithm supported and no merkle tree types are supported.

### Anchor

Anchoring data is a way to prove a piece of data was known to exist at a certain point in time. This can also be referred to as "secure timestamping". When data is anchored, the content hash is converted to a unique deterministic identifier (an [IRI](#iri)) that is stored on chain alongside a timestamp representing the time at which the data was anchored (i.e. the block time of the transaction).

When anchoring data, the sender of the transaction is not attesting to the veracity of the data. The sender can simply be an intermediary providing secure timestamping services. [Attest](#attest) should be used to attest to the veracity of the data.

If the data is altered in any way, both the content hash and the resulting identifier would be different and the data would need to be anchored again, creating a new content entry for the updated data and leaving the previous content entry as a record of the data prior to any changes.

### Attest

Attesting to data is a way to validate a piece of data. Attesting to data is comparable to signing a legal document and implies the contents of the data are accepted to be true by the attestor to the best of their knowledge. When data is attested to, an attestor entry is stored on chain including the identifier ([IRI](#iri)) of the anchored data, the address of the attestor, and a timestamp representing the time at which the data was attested to. The attestor entry can be thought of as a digital signature (separate from but dependant upon a transaction signature).

Attesting to data is only supported with data that uses deterministic, canonical encoding (i.e. the attestor must provide a [graph content hash](#graph-content-hash)). If the data has not already been anchored, attesting to data will also anchor the data. If the same attestor attempts to attest to the same piece of data, the attestation will be ignored (i.e. the previous entry will not be updated and a new entry will not be added). Any number of different attestors can attest to the same piece of data. 

### Resolver

A resolver is used to retrieve data that has been stored off chain and anchored on chain. A resolver defines a URL that refers to an HTTP service where the data is served and should be made available with a simple GET request that takes the [IRI](#iri) of the anchored data.

When defining a resolver, the address that defines the resolver becomes the manager of the resolver. The manager has the ability to register data to the resolver. When registering data that has not been anchored, the data will be automatically anchored before being registered to the resolver.
