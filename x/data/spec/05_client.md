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
regen query data by-iri ...
```

Example Output:

```bash
# not yet implemented
```

#### by-signer

The `by-signer` command allows users to query anchored data based on a signer.

```bash
regen query data by-signer [signer] [flags]
```

Example:

```bash
regen query data by-signer ...
```

Example Output:

```bash
# not yet implemented
```

#### signers

The `signers` command allows users to query signers based on a content hash (i.e. IRI).

```bash
regen query data signers [signers] [flags]
```

Example:

```bash
regen query data signers ...
```

Example Output:

```bash
# not yet implemented
```

### Transactions

The `tx` commands allow users to interact with the `data` module.

```bash
regen tx data --help
```

#### anchor

The `anchor` command allows users to anchor a data to the blockchain based on its content hash (i.e. IRI).

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
```

### Signers

The `Signers` endpoint allows users to query signers by a content hash (i.e. IRI).

```bash
regen.data.v1alpha2.Query/Signers
```

Example:

```bash
grpcurl -plaintext \
    -d '{"signers":"regen1.., regen1.."}' \
    localhost:9090 \
    regen.data.v1alpha2.Query/Signers
```

Example Output:

```bash
```

## REST

A user can query the `data` module using REST endpoints.

### content

The `content` endpoint allows users to query anchored data based on a content hash (i.e. IRI) or a signer.

```bash
/regen/data/v1alpha1/content/{iri_or_signer}
```

Example:

```bash
curl localhost:1317/regen/data/v1alpha1/content/regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
```

Example Output:

```bash
```

### signers

The `signers` endpoint allows users to query signers based on a content hash (i.e. IRI).

```bash
/regen/data/v1alpha1/signers/{iri}
```

Example:

```bash
curl localhost:1317/regen/data/v1alpha1/signers/regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
```

Example Output:

```bash
```
