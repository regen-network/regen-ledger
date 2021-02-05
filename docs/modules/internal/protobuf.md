 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/data/internal/v1alpha2/types.proto](#regen/data/internal/v1alpha2/types.proto)
    - [CompactDataset](#regen.data.internal.v1alpha2.CompactDataset)
    - [GraphID](#regen.data.internal.v1alpha2.GraphID)
    - [Node](#regen.data.internal.v1alpha2.Node)
    - [ObjectGraph](#regen.data.internal.v1alpha2.ObjectGraph)
    - [Properties](#regen.data.internal.v1alpha2.Properties)
  
    - [WellknownDatatype](#regen.data.internal.v1alpha2.WellknownDatatype)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/data/internal/v1alpha2/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/data/internal/v1alpha2/types.proto



<a name="regen.data.internal.v1alpha2.CompactDataset"></a>

### CompactDataset



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [Node](#regen.data.internal.v1alpha2.Node) | repeated |  |
| new_iris | [string](#string) | repeated |  |






<a name="regen.data.internal.v1alpha2.GraphID"></a>

### GraphID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| internal_id | [bytes](#bytes) |  |  |
| local_ref | [sint32](#sint32) |  |  |






<a name="regen.data.internal.v1alpha2.Node"></a>

### Node



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| internal_id | [bytes](#bytes) |  |  |
| local_ref | [sint32](#sint32) |  |  |
| properties | [Properties](#regen.data.internal.v1alpha2.Properties) | repeated |  |






<a name="regen.data.internal.v1alpha2.ObjectGraph"></a>

### ObjectGraph



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_internal_id | [bytes](#bytes) |  |  |
| object_local_ref | [sint32](#sint32) |  |  |
| well_known_datatype | [WellknownDatatype](#regen.data.internal.v1alpha2.WellknownDatatype) |  |  |
| data_type_internal_id | [bytes](#bytes) |  |  |
| data_type_local_ref | [sint32](#sint32) |  |  |
| lang_tag | [string](#string) |  |  |
| graphs | [GraphID](#regen.data.internal.v1alpha2.GraphID) | repeated |  |
| str_value | [string](#string) |  |  |






<a name="regen.data.internal.v1alpha2.Properties"></a>

### Properties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| internal_id | [bytes](#bytes) |  |  |
| local_ref | [sint32](#sint32) |  |  |
| objects | [ObjectGraph](#regen.data.internal.v1alpha2.ObjectGraph) | repeated |  |





 <!-- end messages -->


<a name="regen.data.internal.v1alpha2.WellknownDatatype"></a>

### WellknownDatatype


| Name | Number | Description |
| ---- | ------ | ----------- |
| DATATYPE_UNSPECIFIED | 0 |  |
| DATATYPE_BOOL_FALSE | 1 |  |
| DATATYPE_BOOL_TRUE | 2 |  |
| DATATYPE_DECIMAL | 3 |  |
| DATATYPE_INTEGER | 4 |  |
| DATATYPE_STRING | 5 |  |
| DATATYPE_ANY_URI | 6 |  |
| DATATYPE_DATE | 7 |  |
| DATATYPE_TIME | 8 |  |
| DATATYPE_DATE_TIME | 9 |  |
| DATATYPE_BASE64_STRING | 10 |  |
| DATATYPE_WKT_LITERAL | 11 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

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

