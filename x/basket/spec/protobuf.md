 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/ecocredit/basket/v1/types.proto](#regen/ecocredit/basket/v1/types.proto)
    - [BasketCredit](#regen.ecocredit.basket.v1.BasketCredit)
  
- [regen/ecocredit/basket/v1/tx.proto](#regen/ecocredit/basket/v1/tx.proto)
    - [MsgCreate](#regen.ecocredit.basket.v1.MsgCreate)
    - [MsgCreateResponse](#regen.ecocredit.basket.v1.MsgCreateResponse)
    - [MsgPut](#regen.ecocredit.basket.v1.MsgPut)
    - [MsgPutResponse](#regen.ecocredit.basket.v1.MsgPutResponse)
    - [MsgTake](#regen.ecocredit.basket.v1.MsgTake)
    - [MsgTakeResponse](#regen.ecocredit.basket.v1.MsgTakeResponse)
  
    - [Msg](#regen.ecocredit.basket.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/ecocredit/basket/v1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/basket/v1/types.proto



<a name="regen.ecocredit.basket.v1.BasketCredit"></a>

### BasketCredit
BasketCredit represents the information for a credit batch inside a basket.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| amount | [string](#string) |  | amount is the number of credits being put into or taken out of the basket. Decimal values are acceptable within the precision of the corresponding credit type for this batch. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/basket/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/basket/v1/tx.proto



<a name="regen.ecocredit.basket.v1.MsgCreate"></a>

### MsgCreate
MsgCreateBasket is the Msg/CreateBasket request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| curator | [string](#string) |  | curator is the address of the basket curator who is able to change certain basket settings. |
| name | [string](#string) |  | name will be used to create a bank denom for this basket token. |
| display_name | [string](#string) |  | display_name will be used to create a bank Metadata display name for this basket token. |
| exponent | [uint32](#uint32) |  | exponent is the exponent that will be used for converting credits to basket tokens and for bank denom metadata. An exponent of 6 will mean that 10^6 units of a basket token will be issued for 1.0 credits and that this should be displayed as one unit in user interfaces. The exponent must be >= the precision of the credit type to minimize the need for rounding (rounding may still be needed if the precision changes to be great than the exponent). |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire allows auto-retirement to be disabled. The credits will be auto-retired if disable_auto_retire is false unless the credits were previously put into the basket by the address picking them from the basket, in which case they will remain tradable. |
| credit_type_name | [string](#string) |  | credit_type_name filters against credits from this credit type name. |
| allowed_classes | [string](#string) | repeated | allowed_classes are the credit classes allowed to be put in the basket |
| min_start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | min_start_date is the earliest start date for batches of credits allowed into the basket. |
| fee | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | fee is the fee that the curator will pay to create the basket. It must be >= the required Params.basket_creation_fee. We include the fee explicitly here so that the curator explicitly acknowledges paying this fee and is not surprised to learn that the paid a big fee and didn't know beforehand. |






<a name="regen.ecocredit.basket.v1.MsgCreateResponse"></a>

### MsgCreateResponse
MsgCreateBasketResponse is the Msg/CreateBasket response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basket_denom | [string](#string) |  | basket_denom is the unique denomination ID of the newly created basket. |






<a name="regen.ecocredit.basket.v1.MsgPut"></a>

### MsgPut
MsgAddToBasket is the Msg/AddToBasket request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the owner of credits being put into the basket. |
| basket_denom | [string](#string) |  | basket_denom is the basket denom to add credits to. |
| credits | [BasketCredit](#regen.ecocredit.basket.v1.BasketCredit) | repeated | credits are credits to add to the basket. If they do not match the basket's admission criteria the operation will fail. If there are any "dust" credits left over when converting credits to basket tokens, these credits will not be converted to basket tokens and instead remain with the owner. |






<a name="regen.ecocredit.basket.v1.MsgPutResponse"></a>

### MsgPutResponse
MsgAddToBasketResponse is the Msg/AddToBasket response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount_received | [string](#string) |  | amount_received is the integer amount of basket tokens received. |






<a name="regen.ecocredit.basket.v1.MsgTake"></a>

### MsgTake
MsgTakeFromBasket is the Msg/TakeFromBasket request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the owner of the basket tokens. |
| basket_denom | [string](#string) |  | basket_denom is the basket denom to take credits from. |
| amount | [string](#string) |  | amount is the integer number of basket tokens to convert into credits. |
| retirement_location | [string](#string) |  | retirement_location is the optional retirement location for the credits which will be used only if retire_on_take is true for this basket. |
| retire_on_take | [bool](#bool) |  | retire_on_take is a boolean that dictates whether the ecocredits received in exchange for the basket tokens will be received as retired or tradable credits. |






<a name="regen.ecocredit.basket.v1.MsgTakeResponse"></a>

### MsgTakeResponse
MsgTakeFromBasketResponse is the Msg/TakeFromBasket response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credits | [BasketCredit](#regen.ecocredit.basket.v1.BasketCredit) | repeated | credits are the credits taken out of the basket. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.ecocredit.basket.v1.Msg"></a>

### Msg
Msg is the regen.ecocredit.basket.v1beta1 Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [MsgCreate](#regen.ecocredit.basket.v1.MsgCreate) | [MsgCreateResponse](#regen.ecocredit.basket.v1.MsgCreateResponse) | Create creates a bank denom which wraps credits. |
| Put | [MsgPut](#regen.ecocredit.basket.v1.MsgPut) | [MsgPutResponse](#regen.ecocredit.basket.v1.MsgPutResponse) | Put puts credits into a basket in return for basket tokens. |
| Take | [MsgTake](#regen.ecocredit.basket.v1.MsgTake) | [MsgTakeResponse](#regen.ecocredit.basket.v1.MsgTakeResponse) | Take takes credits from a basket starting from the oldest credits first. |

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

