 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/data/v1alpha2/types.proto](#regen/data/v1alpha2/types.proto)
    - [Content](#regen.data.v1alpha2.Content)
    - [ContentHash](#regen.data.v1alpha2.ContentHash)
    - [ContentHash.Graph](#regen.data.v1alpha2.ContentHash.Graph)
    - [ContentHash.Raw](#regen.data.v1alpha2.ContentHash.Raw)
    - [SignerEntry](#regen.data.v1alpha2.SignerEntry)
  
    - [DigestAlgorithm](#regen.data.v1alpha2.DigestAlgorithm)
    - [GraphCanonicalizationAlgorithm](#regen.data.v1alpha2.GraphCanonicalizationAlgorithm)
    - [GraphMerkleTree](#regen.data.v1alpha2.GraphMerkleTree)
    - [MediaType](#regen.data.v1alpha2.MediaType)
  
- [regen/data/v1alpha2/events.proto](#regen/data/v1alpha2/events.proto)
    - [EventAnchorData](#regen.data.v1alpha2.EventAnchorData)
    - [EventSignData](#regen.data.v1alpha2.EventSignData)
    - [EventStoreRawData](#regen.data.v1alpha2.EventStoreRawData)
  
- [regen/data/v1alpha2/genesis.proto](#regen/data/v1alpha2/genesis.proto)
    - [GenesisContentEntry](#regen.data.v1alpha2.GenesisContentEntry)
    - [GenesisState](#regen.data.v1alpha2.GenesisState)
  
- [regen/data/v1alpha2/query.proto](#regen/data/v1alpha2/query.proto)
    - [ContentEntry](#regen.data.v1alpha2.ContentEntry)
    - [QueryByHashRequest](#regen.data.v1alpha2.QueryByHashRequest)
    - [QueryByHashResponse](#regen.data.v1alpha2.QueryByHashResponse)
    - [QueryBySignerRequest](#regen.data.v1alpha2.QueryBySignerRequest)
    - [QueryBySignerResponse](#regen.data.v1alpha2.QueryBySignerResponse)
  
    - [Query](#regen.data.v1alpha2.Query)
  
- [regen/data/v1alpha2/tx.proto](#regen/data/v1alpha2/tx.proto)
    - [MsgAnchorDataRequest](#regen.data.v1alpha2.MsgAnchorDataRequest)
    - [MsgAnchorDataResponse](#regen.data.v1alpha2.MsgAnchorDataResponse)
    - [MsgSignDataRequest](#regen.data.v1alpha2.MsgSignDataRequest)
    - [MsgSignDataResponse](#regen.data.v1alpha2.MsgSignDataResponse)
    - [MsgStoreRawDataRequest](#regen.data.v1alpha2.MsgStoreRawDataRequest)
    - [MsgStoreRawDataResponse](#regen.data.v1alpha2.MsgStoreRawDataResponse)
  
    - [Msg](#regen.data.v1alpha2.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/data/v1alpha2/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/data/v1alpha2/types.proto



<a name="regen.data.v1alpha2.Content"></a>

### Content
Content is a wrapper for content stored on-chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| raw_data | [bytes](#bytes) |  | raw_data is the oneof field for raw data |






<a name="regen.data.v1alpha2.ContentHash"></a>

### ContentHash
ContentHash specifies a hash based content identifier for a piece of data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| raw | [ContentHash.Raw](#regen.data.v1alpha2.ContentHash.Raw) |  | Raw specifies "raw" data which does not specify a deterministic, canonical encoding. Users of these hashes MUST maintain a copy of the hashed data which is preserved bit by bit. All other content encodings specify a deterministic, canonical encoding allowing implementations to choose from a variety of alternative formats for transport and encoding while maintaining the guarantee that the canonical hash will not change. The media type for "raw" data is defined by the MediaType enum. |
| graph | [ContentHash.Graph](#regen.data.v1alpha2.ContentHash.Graph) |  | Graph specifies graph data that conforms to the RDF data model. The canonicalization algorithm used for an RDF graph is specified by GraphCanonicalizationAlgorithm. |






<a name="regen.data.v1alpha2.ContentHash.Graph"></a>

### ContentHash.Graph
Graph is the content hash type used for RDF graph data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [bytes](#bytes) |  | hash represents the hash of the data based on the specified digest_algorithm |
| digest_algorithm | [DigestAlgorithm](#regen.data.v1alpha2.DigestAlgorithm) |  | digest_algorithm represents the hash digest algorithm. |
| canonicalization_algorithm | [GraphCanonicalizationAlgorithm](#regen.data.v1alpha2.GraphCanonicalizationAlgorithm) |  | graph_canonicalization_algorithm represents the RDF graph canonicalization algorithm. |
| merkle_tree | [GraphMerkleTree](#regen.data.v1alpha2.GraphMerkleTree) |  | merkle_tree is the merkle tree type used for the graph hash, if any |






<a name="regen.data.v1alpha2.ContentHash.Raw"></a>

### ContentHash.Raw
Raw is the content hash type used for raw data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [bytes](#bytes) |  | hash represents the hash of the data based on the specified digest_algorithm |
| digest_algorithm | [DigestAlgorithm](#regen.data.v1alpha2.DigestAlgorithm) |  | digest_algorithm represents the hash digest algorithm. |
| media_type | [MediaType](#regen.data.v1alpha2.MediaType) |  | media_type represents the MediaType for raw data. |






<a name="regen.data.v1alpha2.SignerEntry"></a>

### SignerEntry
SignerEntry is a signer entry wrapping a signer address and timestamp


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| signer | [string](#string) |  | signer is the address of the signer |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is the time at which the data was signed |





 <!-- end messages -->


<a name="regen.data.v1alpha2.DigestAlgorithm"></a>

### DigestAlgorithm
DigestAlgorithm is the hash digest algorithm

| Name | Number | Description |
| ---- | ------ | ----------- |
| DIGEST_ALGORITHM_UNSPECIFIED | 0 | unspecified and invalid |
| DIGEST_ALGORITHM_BLAKE2B_256 | 1 | BLAKE2b-256 |



<a name="regen.data.v1alpha2.GraphCanonicalizationAlgorithm"></a>

### GraphCanonicalizationAlgorithm
GraphCanonicalizationAlgorithm is the graph canonicalization algorithm

| Name | Number | Description |
| ---- | ------ | ----------- |
| GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED | 0 | unspecified and invalid |
| GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015 | 1 | URDNA2015 graph hashing |



<a name="regen.data.v1alpha2.GraphMerkleTree"></a>

### GraphMerkleTree
GraphMerkleTree is the graph merkle tree type used for hashing, if any

| Name | Number | Description |
| ---- | ------ | ----------- |
| GRAPH_MERKLE_TREE_NONE_UNSPECIFIED | 0 | no merkle tree |



<a name="regen.data.v1alpha2.MediaType"></a>

### MediaType
MediaType defines MIME media types to be used with a ContentHash.Raw hash.

| Name | Number | Description |
| ---- | ------ | ----------- |
| MEDIA_TYPE_UNSPECIFIED | 0 | MEDIA_TYPE_UNSPECIFIED can be used for raw binary data |
| MEDIA_TYPE_TEXT_PLAIN | 1 | plain text |
| MEDIA_TYPE_JSON | 2 | JSON |
| MEDIA_TYPE_CSV | 3 | CSV |
| MEDIA_TYPE_XML | 4 | XML |
| MEDIA_TYPE_PDF | 5 | PDF |
| MEDIA_TYPE_TIFF | 16 | TIIF |
| MEDIA_TYPE_JPG | 17 | JPG |
| MEDIA_TYPE_PNG | 18 | PNG |
| MEDIA_TYPE_SVG | 19 | SVG |
| MEDIA_TYPE_WEBP | 20 | WEBP |
| MEDIA_TYPE_AVIF | 21 | AVIF |
| MEDIA_TYPE_GIF | 22 | GIF |
| MEDIA_TYPE_APNG | 23 | APNG |
| MEDIA_TYPE_MPEG | 32 | MPEG |
| MEDIA_TYPE_MP4 | 33 | MP4 |
| MEDIA_TYPE_WEBM | 34 | WEBM |
| MEDIA_TYPE_OGG | 35 | OGG |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/data/v1alpha2/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/data/v1alpha2/events.proto



<a name="regen.data.v1alpha2.EventAnchorData"></a>

### EventAnchorData
EventAnchorData is an event emitted when data is anchored on-chain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iri | [string](#string) |  | iri is the data IRI |






<a name="regen.data.v1alpha2.EventSignData"></a>

### EventSignData
EventSignData is an event emitted when data is signed on-chain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iri | [string](#string) |  | iri is the data IRI |
| signers | [string](#string) | repeated | signers are the addresses of the accounts which have signed the data. |






<a name="regen.data.v1alpha2.EventStoreRawData"></a>

### EventStoreRawData
EventStoreRawData is an event emitted when data is stored on-chain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iri | [string](#string) |  | iri is the data IRI |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/data/v1alpha2/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/data/v1alpha2/genesis.proto



<a name="regen.data.v1alpha2.GenesisContentEntry"></a>

### GenesisContentEntry
GenesisContentEntry is a genesis content entry


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [ContentHash](#regen.data.v1alpha2.ContentHash) |  | hash is the ContentHash |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is the anchor Timestamp |
| signers | [SignerEntry](#regen.data.v1alpha2.SignerEntry) | repeated | signers are the signers, if any |
| content | [Content](#regen.data.v1alpha2.Content) |  | content is the actual content if stored on-chain |






<a name="regen.data.v1alpha2.GenesisState"></a>

### GenesisState
GenesisState is the genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [GenesisContentEntry](#regen.data.v1alpha2.GenesisContentEntry) | repeated | entries are the content entries |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/data/v1alpha2/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/data/v1alpha2/query.proto



<a name="regen.data.v1alpha2.ContentEntry"></a>

### ContentEntry
ContentEntry describes data referenced and possibly stored on chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [ContentHash](#regen.data.v1alpha2.ContentHash) |  | hash is the content hash |
| iri | [string](#string) |  | iri is the content IRI |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is the anchor Timestamp |
| signers | [SignerEntry](#regen.data.v1alpha2.SignerEntry) | repeated | signers are the signers, if any |
| content | [Content](#regen.data.v1alpha2.Content) |  | content is the actual content if stored on-chain |






<a name="regen.data.v1alpha2.QueryByHashRequest"></a>

### QueryByHashRequest
QueryByContentHashRequest is the Query/ByContentHash request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [ContentHash](#regen.data.v1alpha2.ContentHash) |  | hash is the hash-based identifier for the anchored content. |






<a name="regen.data.v1alpha2.QueryByHashResponse"></a>

### QueryByHashResponse
QueryByContentHashResponse is the Query/ByContentHash response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [ContentEntry](#regen.data.v1alpha2.ContentEntry) |  | entry is the ContentEntry |






<a name="regen.data.v1alpha2.QueryBySignerRequest"></a>

### QueryBySignerRequest
QueryBySignerRequest is the Query/BySigner request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| signer | [string](#string) |  | signer is the address of the signer to query by. |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination is the PageRequest to use for pagination. |






<a name="regen.data.v1alpha2.QueryBySignerResponse"></a>

### QueryBySignerResponse
QueryBySignerResponse is the Query/BySigner response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [ContentEntry](#regen.data.v1alpha2.ContentEntry) | repeated | entries is the ContentEntry's signed by the queried signer |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination is the pagination PageResponse. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.data.v1alpha2.Query"></a>

### Query
Query is the regen.data.v1alpha1 Query service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ByHash | [QueryByHashRequest](#regen.data.v1alpha2.QueryByHashRequest) | [QueryByHashResponse](#regen.data.v1alpha2.QueryByHashResponse) | ByHash queries data based on its ContentHash. |
| BySigner | [QueryBySignerRequest](#regen.data.v1alpha2.QueryBySignerRequest) | [QueryBySignerResponse](#regen.data.v1alpha2.QueryBySignerResponse) | BySigner queries data based on signers. |

 <!-- end services -->



<a name="regen/data/v1alpha2/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/data/v1alpha2/tx.proto



<a name="regen.data.v1alpha2.MsgAnchorDataRequest"></a>

### MsgAnchorDataRequest
MsgAnchorDataRequest is the Msg/AnchorData request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the address of the sender of the transaction. The sender in StoreData is not attesting to the veracity of the underlying data. They can simply be a intermediary providing services. |
| hash | [ContentHash](#regen.data.v1alpha2.ContentHash) |  | hash is the hash-based identifier for the anchored content. |






<a name="regen.data.v1alpha2.MsgAnchorDataResponse"></a>

### MsgAnchorDataResponse
MsgAnchorDataRequest is the Msg/AnchorData response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is the timestamp of the block at which the data was anchored. |






<a name="regen.data.v1alpha2.MsgSignDataRequest"></a>

### MsgSignDataRequest
MsgSignDataRequest is the Msg/SignData request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| signers | [string](#string) | repeated | signers are the addresses of the accounts signing the data. By making a SignData request, the signers are attesting to the veracity of the data referenced by the cid. The precise meaning of this may vary depending on the underlying data. |
| hash | [ContentHash.Graph](#regen.data.v1alpha2.ContentHash.Graph) |  | hash is the hash-based identifier for the anchored content. Only RDF graph data can be signed as its data model is intended to specifically convey semantic meaning. |






<a name="regen.data.v1alpha2.MsgSignDataResponse"></a>

### MsgSignDataResponse
MsgSignDataResponse is the Msg/SignData response type.






<a name="regen.data.v1alpha2.MsgStoreRawDataRequest"></a>

### MsgStoreRawDataRequest
MsgStoreRawDataRequest is the Msg/StoreRawData request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the address of the sender of the transaction. The sender in StoreData is not attesting to the veracity of the underlying data. They can simply be a intermediary providing services. |
| content_hash | [ContentHash.Raw](#regen.data.v1alpha2.ContentHash.Raw) |  | content_hash is the hash-based identifier for the anchored content. |
| content | [bytes](#bytes) |  | content is the content of the raw data corresponding to the provided content hash. |






<a name="regen.data.v1alpha2.MsgStoreRawDataResponse"></a>

### MsgStoreRawDataResponse
MsgStoreRawDataRequest is the Msg/StoreRawData response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.data.v1alpha2.Msg"></a>

### Msg
Msg is the regen.data.v1alpha1 Msg service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AnchorData | [MsgAnchorDataRequest](#regen.data.v1alpha2.MsgAnchorDataRequest) | [MsgAnchorDataResponse](#regen.data.v1alpha2.MsgAnchorDataResponse) | AnchorData "anchors" a piece of data to the blockchain based on its secure hash, effectively providing a tamper resistant timestamp.

The sender in AnchorData is not attesting to the veracity of the underlying data. They can simply be a intermediary providing timestamp services. SignData should be used to create a digital signature attesting to the veracity of some piece of data. |
| SignData | [MsgSignDataRequest](#regen.data.v1alpha2.MsgSignDataRequest) | [MsgSignDataResponse](#regen.data.v1alpha2.MsgSignDataResponse) | SignData allows for signing of an arbitrary piece of data on the blockchain. By "signing" data the signers are making a statement about the veracity of the data itself. It is like signing a legal document, meaning that I agree to all conditions and to the best of my knowledge everything is true. When anchoring data, the sender is not attesting to the veracity of the data, they are simply communicating that it exists.

On-chain signatures have the following benefits: - on-chain identities can be managed using different cryptographic keys that change over time through key rotation practices - an on-chain identity may represent an organization and through delegation individual members may sign on behalf of the group - the blockchain transaction envelope provides built-in replay protection and timestamping

SignData implicitly calls AnchorData if the data was not already anchored.

SignData can be called multiple times for the same content hash with different signers and those signers will be appended to the list of signers. |
| StoreRawData | [MsgStoreRawDataRequest](#regen.data.v1alpha2.MsgStoreRawDataRequest) | [MsgStoreRawDataResponse](#regen.data.v1alpha2.MsgStoreRawDataResponse) | StoreRawData stores a piece of raw data corresponding to an ContentHash.Raw on the blockchain.

StoreRawData implicitly calls AnchorData if the data was not already anchored.

The sender in StoreRawData is not attesting to the veracity of the underlying data. They can simply be a intermediary providing storage services. SignData should be used to create a digital signature attesting to the veracity of some piece of data. |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

