 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/data/v1alpha2/types.proto](#regen/data/v1alpha2/types.proto)
    - [ID](#regen.data.v1alpha2.ID)
  
    - [DigestAlgorithm](#regen.data.v1alpha2.DigestAlgorithm)
    - [GraphCanonicalizationAlgorithm](#regen.data.v1alpha2.GraphCanonicalizationAlgorithm)
    - [IDType](#regen.data.v1alpha2.IDType)
    - [MediaType](#regen.data.v1alpha2.MediaType)
  
- [regen/data/v1alpha2/events.proto](#regen/data/v1alpha2/events.proto)
    - [EventAnchorData](#regen.data.v1alpha2.EventAnchorData)
    - [EventSignData](#regen.data.v1alpha2.EventSignData)
    - [EventStoreData](#regen.data.v1alpha2.EventStoreData)
  
- [regen/data/v1alpha2/genesis.proto](#regen/data/v1alpha2/genesis.proto)
    - [GenesisState](#regen.data.v1alpha2.GenesisState)
  
- [regen/data/v1alpha2/query.proto](#regen/data/v1alpha2/query.proto)
    - [QueryByIdRequest](#regen.data.v1alpha2.QueryByIdRequest)
    - [QueryByIdResponse](#regen.data.v1alpha2.QueryByIdResponse)
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



<a name="regen.data.v1alpha2.ID"></a>

### ID
ID specifies a hash based content identifier for a piece of data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [IDType](#regen.data.v1alpha2.IDType) |  | type specifies the IDType for the ID |
| hash | [bytes](#bytes) |  | hash represents the hash of the data based on the specified digest_algorithm and canonicalization algorithm for types other than ID_TYPE_RAW_UNSPECIFIED. |
| digest_algorithm | [DigestAlgorithm](#regen.data.v1alpha2.DigestAlgorithm) |  | digest_algorithm represents the hash digest algorithm. |
| media_type | [MediaType](#regen.data.v1alpha2.MediaType) |  | media_type represents the MediaType for data with ID_TYPE_RAW_UNSPECIFIED. It should be left unset if type is not ID_TYPE_RAW_UNSPECIFIED. |
| graph_canonicalization_algorithm | [GraphCanonicalizationAlgorithm](#regen.data.v1alpha2.GraphCanonicalizationAlgorithm) |  | graph_canonicalization_algorithm represents the RDF graph canonicalization algorithm. It should be left unset if type is not ID_TYPE_GRAPH. |





 <!-- end messages -->


<a name="regen.data.v1alpha2.DigestAlgorithm"></a>

### DigestAlgorithm


| Name | Number | Description |
| ---- | ------ | ----------- |
| DIGEST_ALGORITHM_SHA256_UNSPECIFIED | 0 |  |
| DIGEST_ALGORITHM_BLAKE2B_256 | 1 |  |



<a name="regen.data.v1alpha2.GraphCanonicalizationAlgorithm"></a>

### GraphCanonicalizationAlgorithm


| Name | Number | Description |
| ---- | ------ | ----------- |
| GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015_UNSPECIFIED | 0 |  |



<a name="regen.data.v1alpha2.IDType"></a>

### IDType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ID_TYPE_RAW_UNSPECIFIED | 0 | ID_TYPE_RAW_UNSPECIFIED specifies "raw" data which does not specify a deterministic, canonical encoding. Users of these hashes MUST maintain a copy of the hashed data which is preserved bit by bit. All other encodings in IDType specify a deterministic, canonical encoding allowing implementations to choose from a variety of alternative formats for transport and encoding while maintaining the guarantee that the canonical hash will not change. The media type for "raw" data used with ID_TYPE_RAW_UNSPECIFIED is defined by the MediaType enum. |
| ID_TYPE_GRAPH | 1 | ID_TYPE_GRAPH specifies graph data that conforms to the RDF data model. The canonicalization algorithm used for an RDF graph is specified by GraphCanonicalizationAlgorithm. |
| ID_TYPE_GEOGRAPHY | 2 | ID_TYPE_GEOGRAPHY specifies geographic data conforming to this canonicalization algorithm:

We define the canonical lexical representation of a geography as its [WKT](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry) serialization with the following additional restrictions:

## **General Rules**

- only the WGS 84 reference coordinate system is supported and all coordinates should be assumed to be in WGS 84 - only 2D and 3D coordinates are supported, coordinates with linear referencing systems or an “M” component are unsupported

## **`POLYGON` Rules**

- exterior linear rings must be defined in counterclockwise direction - interior linear rings must be defined in clockwise direction - where there is more than one ring, the first MUST be the exterior ring and other rings MUST be linear rings - this means that the area of a polygon MUST be continuous and unbroken. To represent multiple non-contiguous shapes, a `MULTIPOLYGON` should be used instead. Ex: in the case where inside of a hole in a polygon, there is an island that we want to include. Because the island is disjoint, this shape must be represented as a `MULTIPOLYGON` instead of a `POLYGON` with a hole and an island in the hole - the first and last point in a linear ring must be identical - *canonical coordinate ordering rule*: the first (and last) point in a linear ring must be the point with the minimum X value in the ring. In the event where there is more than one point in a ring with the minimum X value, the first point must be the one with the minimum Y value. For 3D geometries, if there is is more than one point with the minimum X and Y values, the one with the minimum Z value should be the first. - if there is more than one interior linear ring, those rings should be ordered using the above *canonical coordinate ordering rule*, i.e. the ring with the minimum X coordinate comes first, etc.

## `**MULTIPOLYGON` Rules**

- each polygon should conform to the above `POLYGON` rules - if there is more than one polygon, those polygons should be ordered using the above *canonical coordinate ordering rule*, i.e. the polygon with the minimum X coordinate comes first, etc. |



<a name="regen.data.v1alpha2.MediaType"></a>

### MediaType
MediaType defines MIME media types to be used with ID_TYPE_RAW_UNSPECIFIED.

| Name | Number | Description |
| ---- | ------ | ----------- |
| MEDIA_TYPE_UNSPECIFIED | 0 |  |
| MEDIA_TYPE_TEXT_PLAIN | 1 | basic formats |
| MEDIA_TYPE_JSON | 2 |  |
| MEDIA_TYPE_CSV | 3 |  |
| MEDIA_TYPE_XML | 4 |  |
| MEDIA_TYPE_PROTOBUF_ANY | 5 |  |
| MEDIA_TYPE_PDF | 6 |  |
| MEDIA_TYPE_TIFF | 16 | images |
| MEDIA_TYPE_JPG | 17 |  |
| MEDIA_TYPE_PNG | 18 |  |
| MEDIA_TYPE_SVG | 19 |  |
| MEDIA_TYPE_WEBP | 20 |  |
| MEDIA_TYPE_AVIF | 21 |  |
| MEDIA_TYPE_GIF | 22 |  |
| MEDIA_TYPE_APNG | 23 |  |
| MEDIA_TYPE_MPEG | 32 | audio-visual media containers |
| MEDIA_TYPE_MP4 | 33 |  |
| MEDIA_TYPE_WEBM | 34 |  |
| MEDIA_TYPE_OGG | 35 |  |


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
| id | [ID](#regen.data.v1alpha2.ID) |  | id is the hash-based identifier for the anchored content. |






<a name="regen.data.v1alpha2.EventSignData"></a>

### EventSignData
EventAnchorData is an event emitted when data is signed on-chain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [ID](#regen.data.v1alpha2.ID) |  | id is the hash-based identifier for the anchored content. |
| signers | [string](#string) | repeated | signers are the addresses of the accounts which have signed the data. |






<a name="regen.data.v1alpha2.EventStoreData"></a>

### EventStoreData
EventStoreData is an event emitted when data is stored on-chain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [ID](#regen.data.v1alpha2.ID) |  | id is the hash-based identifier for the anchored content. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/data/v1alpha2/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/data/v1alpha2/genesis.proto



<a name="regen.data.v1alpha2.GenesisState"></a>

### GenesisState
TODO





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/data/v1alpha2/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/data/v1alpha2/query.proto



<a name="regen.data.v1alpha2.QueryByIdRequest"></a>

### QueryByIdRequest
QueryByIdRequest is the Query/ById request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [ID](#regen.data.v1alpha2.ID) |  | id is the hash-based identifier for the anchored content. |






<a name="regen.data.v1alpha2.QueryByIdResponse"></a>

### QueryByIdResponse
QueryByIdResponse is the Query/ById response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is the timestamp of the block at which the data was anchored. |
| signers | [string](#string) | repeated | signers are the addresses of the accounts which have signed the data. |
| content | [bytes](#bytes) |  | content is the content of the data, if it was stored on-chain. |






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
| ids | [ID](#regen.data.v1alpha2.ID) | repeated | ids are in the IDs returned in this page of the query. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination is the pagination PageResponse. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.data.v1alpha2.Query"></a>

### Query
Query is the regen.data.v1alpha1 Query service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ById | [QueryByIdRequest](#regen.data.v1alpha2.QueryByIdRequest) | [QueryByIdResponse](#regen.data.v1alpha2.QueryByIdResponse) | ById queries data based on its ID. |
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
| id | [ID](#regen.data.v1alpha2.ID) |  | id is the hash-based identifier for the anchored content. |






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
| id | [ID](#regen.data.v1alpha2.ID) |  | id is the hash-based identifier for the anchored content. |






<a name="regen.data.v1alpha2.MsgSignDataResponse"></a>

### MsgSignDataResponse
MsgSignDataResponse is the Msg/SignData response type.






<a name="regen.data.v1alpha2.MsgStoreRawDataRequest"></a>

### MsgStoreRawDataRequest
MsgStoreRawDataRequest is the Msg/StoreRawData request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the address of the sender of the transaction. The sender in StoreData is not attesting to the veracity of the underlying data. They can simply be a intermediary providing services. |
| id | [ID](#regen.data.v1alpha2.ID) |  | id is the hash-based identifier for the anchored content. The id's type must equal ID_TYPE_RAW_UNSPECIFIED. |
| content | [bytes](#bytes) |  | content is the content of the raw data corresponding to the provided ID. |






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

SignData can be called multiple times for the same ID with different signers and those signers will be appended to the list of signers. |
| StoreRawData | [MsgStoreRawDataRequest](#regen.data.v1alpha2.MsgStoreRawDataRequest) | [MsgStoreRawDataResponse](#regen.data.v1alpha2.MsgStoreRawDataResponse) | StoreRawData stores a piece of raw data corresponding to an ID on the blockchain.

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

