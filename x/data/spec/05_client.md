# Client

## CLI

A user can query and interact with the `data` module using the CLI.

### Query

The `query` commands allow users to query `data` state.

```bash
regen query data --help
```

#### by-iri

The `by-iri` command allows users to query anchored data based on its content hash (i.e. IRI).

```bash
regen query data by-iri [iri] [flags]
```

Example:

```bash
regen query data by-iri regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
```

Example Output:

```bash
entry:
  hash:
    graph:
      canonicalization_algorithm: GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015
      digest_algorithm: DIGEST_ALGORITHM_BLAKE2B_256
      hash: YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=
      merkle_tree: GRAPH_MERKLE_TREE_NONE_UNSPECIFIED
  iri: regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
  timestamp: "2022-01-01T00:00:00.000000000Z"
```

#### by-signer

The `by-signer` command allows users to query anchored data based on a signer.

```bash
regen query data by-signer [signer] [flags]
```

Example:

```bash
regen query data by-signer regen1..
```

Example Output:

```bash
entries:
- hash:
    graph:
      canonicalization_algorithm: GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015
      digest_algorithm: DIGEST_ALGORITHM_BLAKE2B_256
      hash: YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=
      merkle_tree: GRAPH_MERKLE_TREE_NONE_UNSPECIFIED
  iri: regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
  timestamp: "2022-01-01T00:00:00.000000000Z"
pagination:
  next_key: null
  total: "1"
```

#### signers

The `signers` command allows users to query signers based on a content hash (i.e. IRI).

```bash
regen query data signers [iri] [flags]
```

Example:

```bash
regen query data signers regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
```

Example Output:

```bash
pagination:
  next_key: null
  total: "1"
signers:
- regen1..
```

### Transactions

The `tx` commands allow users to interact with the `data` module.

```bash
regen tx data --help
```

#### anchor

The `anchor` command allows users to anchor data to the blockchain based on its content hash (i.e. IRI).

```bash
regen tx data anchor [iri] [flags]
```

Example:

```bash
regen tx data anchor regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf --from regen1..
```

#### sign

The `sign` command allows users to sign anchored data on the blockchain.

```bash
regen tx data sign [iri] [flags]
```

Example:

```bash
regen tx data sign regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf --from regen1..
```

## gRPC

A user can query the `data` module using gRPC endpoints.

### ByIRI

The `ByIRI` endpoint allows users to query anchored data by the content hash (i.e. IRI).

```bash
regen.data.v1alpha2.Query/ByIRI
```

Example:

```bash
grpcurl -plaintext \
    -d '{"iri":"regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"}' \
    localhost:9090 \
    regen.data.v1alpha2.Query/ByIRI
```

Example Output:

```bash
{
  "entry": {
    "hash": {
      "graph": {
        "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
        "digestAlgorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
        "canonicalizationAlgorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015"
      }
    },
    "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
    "timestamp": "2022-01-01T00:00:00.000000000Z"
  }
}
```

### BySigner

The `BySigner` endpoint allows users to query anchored data by signer.

```bash
regen.data.v1alpha2.Query/BySigner
```

Example:

```bash
grpcurl -plaintext \
    -d '{"signer":"regen1.."}' \
    localhost:9090 \
    regen.data.v1alpha2.Query/BySigner
```

Example Output:

```bash
{
  "entries": [
    {
      "hash": {
        "graph": {
          "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
          "digestAlgorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
          "canonicalizationAlgorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015"
        }
      },
      "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "timestamp": "2022-01-01T00:00:00.000000000Z"
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### Signers

The `Signers` endpoint allows users to query signers by a content hash (i.e. IRI).

```bash
regen.data.v1alpha2.Query/Signers
```

Example:

```bash
grpcurl -plaintext \
    -d '{"iri":"regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"}' \
    localhost:9090 \
    regen.data.v1alpha2.Query/Signers
```

Example Output:

```bash
{
  "signers": [
    "regen16md38uw5z9v4du2dtq4qgake8ewyf36u6qgfza"
  ],
  "pagination": {
    "total": "1"
  }
}
```

## REST

A user can query the `data` module using REST endpoints.

### by-iri

The `content` endpoint allows users to query anchored data based on a content hash (i.e. IRI).

```bash
/regen/data/v1alpha2/by-iri/{iri}
```

Example:

```bash
curl localhost:1317/regen/data/v1alpha2/by-iri/regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
```

Example Output:

```bash
{
  "entry": {
    "hash": {
      "graph": {
        "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
        "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
        "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
        "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
      }
    },
    "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
    "timestamp": "2022-01-01T00:00:00.000000000Z"
  }
}
```

### by-signer

The `by-signer` endpoint allows users to query anchored data based on a signer.

```bash
/regen/data/v1alpha2/by-signer/{signer}
```

Example:

```bash
curl localhost:1317/regen/data/v1alpha2/by-signer/regen1..
```

Example Output:

```bash
{
  "entries": [
    {
      "hash": {
        "graph": {
          "hash": "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXoxMjM0NTY=",
          "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
          "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
          "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
        }
      },
      "iri": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "timestamp": "2022-01-01T00:00:00.000000000Z"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### signers

The `signers` endpoint allows users to query signers based on a content hash (i.e. IRI).

```bash
/regen/data/v1alpha2/signers/{iri}
```

Example:

```bash
curl localhost:1317/regen/data/v1alpha2/signers/regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
```

Example Output:

```bash
{
  "signers": [
    "regen1.."
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```
