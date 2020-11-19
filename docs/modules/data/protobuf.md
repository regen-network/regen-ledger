 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/data/v1alpha1/events.proto](#regen/data/v1alpha1/events.proto)
    - [EventAnchorData](#regen.data.v1alpha1.EventAnchorData)
    - [EventSignData](#regen.data.v1alpha1.EventSignData)
    - [EventStoreData](#regen.data.v1alpha1.EventStoreData)
  
- [regen/data/v1alpha1/genesis.proto](#regen/data/v1alpha1/genesis.proto)
    - [GenesisState](#regen.data.v1alpha1.GenesisState)
  
- [regen/data/v1alpha1/query.proto](#regen/data/v1alpha1/query.proto)
    - [QueryByCidRequest](#regen.data.v1alpha1.QueryByCidRequest)
    - [QueryByCidResponse](#regen.data.v1alpha1.QueryByCidResponse)
    - [QueryBySignerRequest](#regen.data.v1alpha1.QueryBySignerRequest)
    - [QueryBySignerResponse](#regen.data.v1alpha1.QueryBySignerResponse)
  
    - [Query](#regen.data.v1alpha1.Query)
  
- [regen/data/v1alpha1/tx.proto](#regen/data/v1alpha1/tx.proto)
    - [MsgAnchorDataRequest](#regen.data.v1alpha1.MsgAnchorDataRequest)
    - [MsgAnchorDataResponse](#regen.data.v1alpha1.MsgAnchorDataResponse)
    - [MsgSignDataRequest](#regen.data.v1alpha1.MsgSignDataRequest)
    - [MsgSignDataResponse](#regen.data.v1alpha1.MsgSignDataResponse)
    - [MsgStoreDataRequest](#regen.data.v1alpha1.MsgStoreDataRequest)
    - [MsgStoreDataResponse](#regen.data.v1alpha1.MsgStoreDataResponse)
  
    - [Msg](#regen.data.v1alpha1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/data/v1alpha1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/data/v1alpha1/events.proto



<a name="regen.data.v1alpha1.EventAnchorData"></a>

### EventAnchorData
EventAnchorData is an event emitted when data is anchored on-chain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cid | [bytes](#bytes) |  | cid is a Content Identifier for the data corresponding to the IPFS CID specification: https://github.com/multiformats/cid. |






<a name="regen.data.v1alpha1.EventSignData"></a>

### EventSignData
EventAnchorData is an event emitted when data is signed on-chain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cid | [bytes](#bytes) |  | cid is a Content Identifier for the data corresponding to the IPFS CID specification: https://github.com/multiformats/cid. |
| signers | [string](#string) | repeated | signers are the addresses of the accounts which have signed the data. |






<a name="regen.data.v1alpha1.EventStoreData"></a>

### EventStoreData
EventAnchorData is an event emitted when data is stored on-chain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cid | [bytes](#bytes) |  | cid is a Content Identifier for the data corresponding to the IPFS CID specification: https://github.com/multiformats/cid. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/data/v1alpha1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/data/v1alpha1/genesis.proto



<a name="regen.data.v1alpha1.GenesisState"></a>

### GenesisState
TODO





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/data/v1alpha1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/data/v1alpha1/query.proto



<a name="regen.data.v1alpha1.QueryByCidRequest"></a>

### QueryByCidRequest
QueryByCidRequest is the Query/ByCid request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cid | [bytes](#bytes) |  | cid is a Content Identifier for the data corresponding to the IPFS CID specification: https://github.com/multiformats/cid. |






<a name="regen.data.v1alpha1.QueryByCidResponse"></a>

### QueryByCidResponse
QueryByCidResponse is the Query/ByCid response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is the timestamp of the block at which the data was anchored. |
| signers | [string](#string) | repeated | signers are the addresses of the accounts which have signed the data. |
| content | [bytes](#bytes) |  | content is the content of the data, if it was stored on-chain. |






<a name="regen.data.v1alpha1.QueryBySignerRequest"></a>

### QueryBySignerRequest
QueryBySignerRequest is the Query/BySigner request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| signer | [string](#string) |  | signer is the address of the signer to query by. |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination is the PageRequest to use for pagination. |






<a name="regen.data.v1alpha1.QueryBySignerResponse"></a>

### QueryBySignerResponse
QueryBySignerResponse is the Query/BySigner response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cids | [bytes](#bytes) | repeated | cids are in the CIDs returned in this page of the query. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination is the pagination PageResponse. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.data.v1alpha1.Query"></a>

### Query
Query is the regen.data.v1alpha1 Query service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ByCid | [QueryByCidRequest](#regen.data.v1alpha1.QueryByCidRequest) | [QueryByCidResponse](#regen.data.v1alpha1.QueryByCidResponse) | ByCid queries data based on its CID. |
| BySigner | [QueryBySignerRequest](#regen.data.v1alpha1.QueryBySignerRequest) | [QueryBySignerResponse](#regen.data.v1alpha1.QueryBySignerResponse) | BySigner queries data based on signers. |

 <!-- end services -->



<a name="regen/data/v1alpha1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/data/v1alpha1/tx.proto



<a name="regen.data.v1alpha1.MsgAnchorDataRequest"></a>

### MsgAnchorDataRequest
MsgAnchorDataRequest is the Msg/AnchorData request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the address of the sender of the transaction. The sender in StoreData is not attesting to the veracity of the underlying data. They can simply be a intermediary providing services. |
| cid | [bytes](#bytes) |  | cid is a Content Identifier for the data corresponding to the IPFS CID specification: https://github.com/multiformats/cid. |






<a name="regen.data.v1alpha1.MsgAnchorDataResponse"></a>

### MsgAnchorDataResponse
MsgAnchorDataRequest is the Msg/AnchorData response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is the timestamp of the block at which the data was anchored. |






<a name="regen.data.v1alpha1.MsgSignDataRequest"></a>

### MsgSignDataRequest
MsgSignDataRequest is the Msg/SignData request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| signers | [string](#string) | repeated | signers are the addresses of the accounts signing the data. By making a SignData request, the signers are attesting to the veracity of the data referenced by the cid. The precise meaning of this may vary depending on the underlying data. |
| cid | [bytes](#bytes) |  | cid is a Content Identifier for the data corresponding to the IPFS CID specification: https://github.com/multiformats/cid. |






<a name="regen.data.v1alpha1.MsgSignDataResponse"></a>

### MsgSignDataResponse
MsgSignDataResponse is the Msg/SignData response type.






<a name="regen.data.v1alpha1.MsgStoreDataRequest"></a>

### MsgStoreDataRequest
MsgStoreDataRequest is the Msg/StoreData request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the address of the sender of the transaction. The sender in StoreData is not attesting to the veracity of the underlying data. They can simply be a intermediary providing services. |
| cid | [bytes](#bytes) |  | cid is a Content Identifier for the data corresponding to the IPFS CID specification: https://github.com/multiformats/cid. |
| content | [bytes](#bytes) |  | content is the content of the data corresponding to the provided CID.

Currently only data for CID's using sha2-256 and blake2b-256 hash algorithms can be stored on-chain. |






<a name="regen.data.v1alpha1.MsgStoreDataResponse"></a>

### MsgStoreDataResponse
MsgStoreDataRequest is the Msg/StoreData response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.data.v1alpha1.Msg"></a>

### Msg
Msg is the regen.data.v1alpha1 Msg service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AnchorData | [MsgAnchorDataRequest](#regen.data.v1alpha1.MsgAnchorDataRequest) | [MsgAnchorDataResponse](#regen.data.v1alpha1.MsgAnchorDataResponse) | AnchorData "anchors" a piece of data to the blockchain based on its secure hash, effectively providing a tamper resistant timestamp.

The sender in AnchorData is not attesting to the veracity of the underlying data. They can simply be a intermediary providing timestamp services. SignData should be used to create a digital signature attesting to the veracity of some piece of data. |
| SignData | [MsgSignDataRequest](#regen.data.v1alpha1.MsgSignDataRequest) | [MsgSignDataResponse](#regen.data.v1alpha1.MsgSignDataResponse) | SignData allows for signing of an arbitrary piece of data on the blockchain. By "signing" data the signers are making a statement about the veracity of the data itself. It is like signing a legal document, meaning that I agree to all conditions and to the best of my knowledge everything is true. When anchoring data, the sender is not attesting to the veracity of the data, they are simply communicating that it exists.

On-chain signatures have the following benefits: - on-chain identities can be managed using different cryptographic keys that change over time through key rotation practices - an on-chain identity may represent an organization and through delegation individual members may sign on behalf of the group - the blockchain transaction envelope provides built-in replay protection and timestamping

SignData implicitly calls AnchorData if the data was not already anchored.

SignData can be called multiple times for the same CID with different signers and those signers will be appended to the list of signers. |
| StoreData | [MsgStoreDataRequest](#regen.data.v1alpha1.MsgStoreDataRequest) | [MsgStoreDataResponse](#regen.data.v1alpha1.MsgStoreDataResponse) | StoreData stores a piece of data corresponding to a CID on the blockchain.

Currently only data for CID's using sha2-256 and blake2b-256 hash algorithms can be stored on-chain.

StoreData implicitly calls AnchorData if the data was not already anchored.

The sender in StoreData is not attesting to the veracity of the underlying data. They can simply be a intermediary providing storage services. SignData should be used to create a digital signature attesting to the veracity of some piece of data. |

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

