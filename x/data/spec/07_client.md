# Client

## CLI

A user can query and interact with the `data` module using the CLI.

### Query

The `query` commands allow users to query `data` state.

For examples on how to query state using CLI, see the data module [Query commands](https://docs.regen.network/commands/regen_query_data.html) documentation.

### Transactions

The `tx` commands allow users to interact with the `data` module.

For examples on how to submit transactions using CLI, see the data module [Transaction commands](https://docs.regen.network/commands/regen_tx_data.html) documentation.

## gRPC

A user can query the `data` module using gRPC endpoints.

### AnchorByIRI

The `AnchorByIRI` endpoint allows users to query a data anchor by the IRI of the data.

```bash
regen.data.v1.Query/AnchorByIRI
```

Example:

```bash
grpcurl -plaintext \
    -d '{"iri":"regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"}' \
    localhost:9090 \
    regen.data.v1.Query/AnchorByIRI
```

Example Output:

```bash
{
  "anchor": {
    "contentHash": {
      "graph": {
        "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
        "digestAlgorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
        "canonicalizationAlgorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
        "markleTree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
      }
    },
    "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
    "timestamp": "2022-01-01T00:00:00.000000000Z"
  }
}
```

### AnchorByHash

The `AnchorByHash` endpoint allows users to query a data anchor by the ContentHash of the data

```bash
regen.data.v1.Query/AnchorByHash
```

Example:

```bash
grpcurl -plaintext \
    -d '{"contentHash": {
        "graph": {
          "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
          "digestAlgorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
          "canonicalizationAlgorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015"
        }
      }
    }' \
    localhost:9090 \
    regen.data.v1.Query/AnchorByHash
```

Example Output:

```bash
{
  "anchor": {
    "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
    "contentHash": {
      "graph": {
        "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
        "digestAlgorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
        "canonicalizationAlgorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015"
      }
    },
    "timestamp": "2022-07-06T11:54:58.464341467Z"
  }
}
```

### AttestationsByAttestor

The `AttestationsByAttestor` endpoint allows users to query data attestations by an attestor.

```bash
regen.data.v1.Query/AttestationsByAttestor
```

Example:

```bash
grpcurl -plaintext \
    -d '{"attestor": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5"}' \
    localhost:9090 \
    regen.data.v1.Query/AttestationsByAttestor
```

Example Output:

```bash
{
  "attestations": [
    {
      "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "attestor": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5",
      "timestamp": "2022-07-07T05:44:06.119352311Z"
    }
  ]
}
```

### AttestationsByIRI

The `AttestationsByIRI` endpoint allows users to query data attestations by an iri.

```bash
regen.data.v1.Query/AttestationsByIRI
```

Example:

```bash
grpcurl -plaintext \ 
    -d '{"iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"}' \
    localhost:9090 \
    regen.data.v1.Query/AttestationsByIRI
```

Example Output:

```bash
{
  "attestations": [
    {
      "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "attestor": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5",
      "timestamp": "2022-07-07T05:44:06.119352311Z"
    }
  ]
}
```

### AttestationsByHash

The `AttestationsByHash` endpoint allows users to query by the ContentHash of the data.

```bash
regen.data.v1.Query/AttestationsByHash
```

Example:

```bash
grpcurl -plaintext \
    -d '{
        "content_hash": {
          "graph": {
            "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
            "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
            "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY="
          }
        }
    }' \
    localhost:9090 \
    regen.data.v1.Query/AttestationsByHash
```

Example Output:

```bash
{
  "attestations": [
    {
      "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "attestor": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5",
      "timestamp": "2022-07-07T05:44:06.119352311Z"
    }
  ]
}
```

### Resolver

The `Resolver` allows users to query resolver by its unique identifier.

```bash
regen.data.v1.Query/Resolver
```

Example:

```bash
grpcurl -plaintext \
    -d '{"id":1}' \
    localhost:9090 \
    regen.data.v1.Query/Resolver
```

Example Output:

```bash
{
  "resolver": {
    "id": "1",
    "url": "http://foo.bar",
    "manager": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5"
  }
}
```

### ResolversByIRI

The `ResolversByIRI` allows users to query resolvers with registered data by the IRI of the data.

```bash
regen.data.v1.Query/ResolversByIRI
```

Example:

```bash
grpcurl -plaintext \
  -d '{"iri":"regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"}' \
  localhost:9090 \
  regen.data.v1.Query/ResolversByIRI
```

Example Output:

```bash
{
  "resolvers": [
    {
      "id": "1",
      "url": "http://foo.bar",
      "manager": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5"
    }
  ]
}
```

### ResolversByHash

The `ResolversByHash` allows users to query resolvers with registered data by the ContentHash of the data.

```bash
regen.data.v1.Query/ResolversByHash
```

Example:

```bash
grpcurl -plaintext \
    -d '{
      "content_hash":{
        "graph":{
          "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
          "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
          "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015"
        }
      }
    }' \
    localhost:9090 \
    regen.data.v1.Query/ResolversByHash
```

Example Output:

```bash
{
  "resolvers": [
    {
      "id": "1",
      "url": "http://foo.bar",
      "manager": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5"
    }
  ]
}
```

### ResolversByURL

The `ResolversByURL` allows users to query resolvers by URL.

```bash
regen.data.v1.Query/ResolversByURL
```

Example:

```bash
grpcurl -plaintext \
    -d '{"url":"http://foo.bar"}' \
    localhost:9090 \
    regen.data.v1.Query/ResolversByURL
```

Example Output:

```bash
{
  "resolvers": [
    {
      "id": "1",
      "url": "http://foo.bar",
      "manager": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5"
    }
  ]
}
```

### ConvertIRIToHash

The `ConvertIRIToHash` allows users to convert IRI to a ContentHash.

```bash
regen.data.v1.Query/ConvertIRIToHash
```

Example:

```bash
grpcurl -plaintext \
    -d '{"iri":"regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"}' \
    localhost:9090 \
    regen.data.v1.Query/ConvertIRIToHash
```

Example Output:

```bash
{
  "contentHash": {
    "graph": {
      "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
      "digestAlgorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
      "canonicalizationAlgorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015"
    }
  }
}
```

### ConvertHashToIRI

The `ConvertHashToIRI` endpoint allows users to convert ContentHash to an IRI.

```bash
regen.data.v1.ConvertHashToIRI
```

Example:

```bash
grpcurl -plaintext \
    -d '{
      "content_hash": {
        "graph": {
          "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
          "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
          "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015"
        }
      }
    }' \
    localhost:9090 \
    regen.data.v1.Query/ConvertHashToIRI
```

Example Output:

```bash
{
  "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
}
```

## REST

A user can query the `data` module using REST endpoints.

### anchor-by-iri

The `anchor-by-iri` endpoint allows users to query a data anchor by the IRI of the data.

```bash
/regen/data/v1/anchor-by-iri/{iri}
/regen/data/v1/anchors/iri/{iri}
```

Example:

```bash
curl localhost:1317/regen/data/v1/anchors/iri/regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
```

Example Output:

```bash
{
  "anchor": {
    "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
    "content_hash": {
      "raw": null,
      "graph": {
        "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
        "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
        "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
        "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
      }
    },
    "timestamp": "2022-07-06T11:54:58.464341467Z"
  }
}
```

### anchor-by-hash

The `anchor-by-hash` allows users to query a data anchor by the ContentHash of the data.

```bash
/regen/data/v1/anchor-by-hash
/regen/data/v1/anchors/hash
```

Example:

```bash
curl \
    -d '{"content_hash": {"graph": {"canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015","digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256","hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY="}}}' \
    -H 'Content-Type: application/json' \
    localhost:1317/regen/data/v1/anchors/hash
```

Example Output:

```bash
{
    "anchor": {
        "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
        "content_hash": {
            "raw": null,
            "graph": {
                "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
                "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
                "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
                "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
            }
        },
        "timestamp": "2022-07-06T11:54:58.464341467Z"
    }
}
```

### attestations-by-attestor

The `attestations-by-attestor` endpoint allows users to query data attestations by an attestor.

```bash
/regen/data/v1/attestations-by-attestor/{attestor}
/regen/data/v1/attestations/attestor/{attestor}
```

Example:

```bash
curl localhost:1317/regen/data/v1/attestations-by-attestor/regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5
```

Example Output:

```bash
{
    "attestations": [
        {
            "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
            "attestor": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5",
            "timestamp": "2022-07-07T05:44:06.119352311Z"
        }
    ],
    "pagination": null
}
```

### attestations-by-iri

The `attestations-by-iri` allows users to query data attestations by IRI.

```bash
/regen/data/v1/attestations-by-iri/{iri}
/regen/data/v1/attestations/iri/{iri}
```

Example:

```bash
curl localhost:1317/regen/data/v1/attestations-by-iri/regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
```

Example Output:

```bash
{
    "attestations": [
        {
            "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
            "attestor": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5",
            "timestamp": "2022-07-07T05:44:06.119352311Z"
        }
    ],
    "pagination": null
}
```

### attestations-by-hash

The `attestations-by-hash` allows users to query data attestations by the ContentHash of the data.

```bash
/regen/data/v1/attestations-by-hash
/regen/data/v1/attestations/hash
```

Example:

```bash
curl \
    -d '{"content_hash": {"graph": {"canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015","digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256","hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY="}}}' \
    -H 'Content-Type: application/json' \
    localhost:1317/regen/data/v1/attestations/hash
```

Example Output:

```bash
{
    "attestations": [
        {
            "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
            "attestor": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5",
            "timestamp": "2022-07-07T05:44:06.119352311Z"
        }
    ],
    "pagination": null
}
```

### resolver

The `resovler` endpoint allows users to query a resolver by its unique identifier.

```bash
/regen/data/v1/resolver/{id}
/regen/data/v1/resolvers/{id}
```

Example:

```bash
curl localhost:1317/regen/data/v1/resolvers/1 
```

Example Output:

```bash
{
  "resolver": {
    "id": "1",
    "url": "http://foo.bar",
    "manager": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5"
  }
}
```

### resolvers-by-iri

The `resolvers-by-iri` endpoint allows users to query resolvers with registered data by the IRI
of the data.

```bash
/regen/data/v1/resolvers-by-iri/{iri}
/regen/data/v1/resolvers/iri/{iri}
```

Example:

```bash
curl locahost:1317/regen/data/v1/resolvers/iri/regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
```

Example Output:

```bash
{
    "resolvers": [
        {
            "id": "1",
            "url": "http://foo.bar",
            "manager": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5"
        }
    ],
    "pagination": null
}
```

### resolvers-by-hash

The `resolvers-by-hash` endpoint allows users to query resolvers with registered data by the ContentHash 
of the data.

```bash
/regen/data/v1/resolvers-by-hash
/regen/data/v1/resolvers/hash
```

Example:

```bash
curl \
    -d '{"content_hash":{"graph":{"hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=","digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256","canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015"}}}' \
    -H 'Content-Type: application/json' \
    localhost:1317/regen/data/v1/resolvers/hash
```

Example Output:

```bash
{
    "resolvers": [
        {
            "id": "1",
            "url": "http://foo.bar",
            "manager": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5"
        }
    ],
    "pagination": null
}
```

### resolvers-by-url

The `resolvers-by-url` allows users to query resolvers by URL.

```bash
/regen/data/v1/resolvers-by-url
/regen/data/v1/resolvers/url
```

Example:

```bash
curl \
    -d '{"url":"http://foo.bar"}' \
    -H 'Content-Type: application/json' \
    localhost:1317/regen/data/v1/resolvers/url
```

Example Output:

```bash
{
    "resolvers": [
        {
            "id": "1",
            "url": "http://foo.bar",
            "manager": "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5"
        }
    ],
    "pagination": null
}
```

### convertIRI-to-hash

The `convertIRI-to-hash` allows users to convert IRI to a ContentHash.

```bash
/regen/data/v1/convert-iri-to-hash/{iri}
```

Example:

```bash
curl localhost:1317/regen/data/v1/convert-iri-to-hash/regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
```


Example Output:

```bash
{
    "content_hash": {
        "raw": null,
        "graph": {
            "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
            "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
            "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
            "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
        }
    }
}
```

### convert-hash-to-IRI

The `convert-hash-to-IRI` endpoint allows users to convert ContentHash to an IRI.

```bash
/regen/data/v1/convert-hash-to-iri
```

Example:

```bash
curl \
    -d '{"content_hash": {"raw": null,"graph": {"hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=","digest_algorithm":"DIGEST_ALGORITHM_BLAKE2B_256","canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015"}}}' \
    -H 'Content-Type: application/json' \
    localhost:1317/regen/data/v1/convert-hash-to-iri
```

Example Output:

```bash
{
    "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
}
```