 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/ecocredit/basket/v1/types.proto](#regen/ecocredit/basket/v1/types.proto)
    - [Basket](#regen.ecocredit.basket.v1.Basket)
    - [BasketBalance](#regen.ecocredit.basket.v1.BasketBalance)
    - [BasketCredit](#regen.ecocredit.basket.v1.BasketCredit)
    - [DateCriteria](#regen.ecocredit.basket.v1.DateCriteria)
  
- [regen/ecocredit/basket/v1/events.proto](#regen/ecocredit/basket/v1/events.proto)
    - [EventCreate](#regen.ecocredit.basket.v1.EventCreate)
    - [EventPut](#regen.ecocredit.basket.v1.EventPut)
    - [EventTake](#regen.ecocredit.basket.v1.EventTake)
  
- [regen/ecocredit/basket/v1/query.proto](#regen/ecocredit/basket/v1/query.proto)
    - [QueryBasketBalanceRequest](#regen.ecocredit.basket.v1.QueryBasketBalanceRequest)
    - [QueryBasketBalanceResponse](#regen.ecocredit.basket.v1.QueryBasketBalanceResponse)
    - [QueryBasketBalancesRequest](#regen.ecocredit.basket.v1.QueryBasketBalancesRequest)
    - [QueryBasketBalancesResponse](#regen.ecocredit.basket.v1.QueryBasketBalancesResponse)
    - [QueryBasketRequest](#regen.ecocredit.basket.v1.QueryBasketRequest)
    - [QueryBasketResponse](#regen.ecocredit.basket.v1.QueryBasketResponse)
    - [QueryBasketsRequest](#regen.ecocredit.basket.v1.QueryBasketsRequest)
    - [QueryBasketsResponse](#regen.ecocredit.basket.v1.QueryBasketsResponse)
  
    - [Query](#regen.ecocredit.basket.v1.Query)
  
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



<a name="regen.ecocredit.basket.v1.Basket"></a>

### Basket
Basket represents a basket in state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id is the uint64 ID of the basket. It is used internally for reducing storage space. |
| basket_denom | [string](#string) |  | basket_denom is the basket bank denom. |
| name | [string](#string) |  | name is the unique name of the basket specified in MsgCreate. Basket names must be unique across all credit types and choices of exponent above and beyond the uniqueness constraint on basket_denom. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire indicates whether or not the credits will be retired upon withdraw from the basket. |
| credit_type_abbrev | [string](#string) |  | credit_type_abbrev is the abbreviation of the credit type this basket is able to hold. |
| date_criteria | [DateCriteria](#regen.ecocredit.basket.v1.DateCriteria) |  | date_criteria is the date criteria for batches admitted to the basket. |
| exponent | [uint32](#uint32) |  | exponent is the exponent for converting credits to/from basket tokens. |






<a name="regen.ecocredit.basket.v1.BasketBalance"></a>

### BasketBalance
BasketBalance stores the amount of credits from a batch in a basket


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basket_id | [uint64](#uint64) |  | basket_id is the ID of the basket |
| batch_denom | [string](#string) |  | batch_denom is the denom of the credit batch |
| balance | [string](#string) |  | balance is the amount of ecocredits held in the basket |
| batch_start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | batch_start_date is the start date of the batch. This field is used to create an index which is used to remove the oldest credits first. |






<a name="regen.ecocredit.basket.v1.BasketCredit"></a>

### BasketCredit
BasketCredit represents the information for a credit batch inside a basket.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| amount | [string](#string) |  | amount is the number of credits being put into or taken out of the basket. Decimal values are acceptable within the precision of the corresponding credit type for this batch. |






<a name="regen.ecocredit.basket.v1.DateCriteria"></a>

### DateCriteria
DateCriteria represents the information for credit acceptance in a basket.
At most, only one of the values should be set.
NOTE: gogo proto `oneof` is not compatible with Amino signing, hence we directly define
both `start_date_window` and `min_start_date`. In the future, with pulsar, this should change
and we should use `oneof`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min_start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | min_start_date (optional) is the earliest start date for batches of credits allowed into the basket. At most only one of `start_date_window` and `min_start_date` can be set. |
| start_date_window | [google.protobuf.Duration](#google.protobuf.Duration) |  | start_date_window (optional) is a duration of time measured into the past which sets a cutoff for batch start dates when adding new credits to the basket. Based on the current block timestamp, credits whose start date is before `block_timestamp - batch_date_window` will not be allowed into the basket. At most only one of `start_date_window` and `min_start_date` can be set. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/basket/v1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/basket/v1/events.proto



<a name="regen.ecocredit.basket.v1.EventCreate"></a>

### EventCreate
EventCreate is an event emitted when a basket is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basket_denom | [string](#string) |  | basket_denom is the basket bank denom. |
| curator | [string](#string) |  | curator is the address of the basket curator who is able to change certain basket settings. |






<a name="regen.ecocredit.basket.v1.EventPut"></a>

### EventPut
EventPut is an event emitted when credits are put into a basket in return
for basket tokens.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the owner of the credits put into the basket. |
| basket_denom | [string](#string) |  | basket_denom is the basket bank denom that the credits were added to. |
| credits | [BasketCredit](#regen.ecocredit.basket.v1.BasketCredit) | repeated | credits are the credits that were added to the basket. |
| amount | [string](#string) |  | amount is the integer number of basket tokens converted from credits. |






<a name="regen.ecocredit.basket.v1.EventTake"></a>

### EventTake
EventTake is an event emitted when credits are taken from a basket starting
from the oldest credits first.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the owner of the credits taken from the basket. |
| basket_denom | [string](#string) |  | basket_denom is the basket bank denom that credits were taken from. |
| credits | [BasketCredit](#regen.ecocredit.basket.v1.BasketCredit) | repeated | credits are the credits that were taken from the basket. |
| amount | [string](#string) |  | amount is the integer number of basket tokens converted to credits. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/basket/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/basket/v1/query.proto



<a name="regen.ecocredit.basket.v1.QueryBasketBalanceRequest"></a>

### QueryBasketBalanceRequest
QueryBasketBalanceRequest is the Query/BasketBalance request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basket_denom | [string](#string) |  | basket_denom is the denom of the basket. |
| batch_denom | [string](#string) |  | batch_denom is the denom of the credit batch. |






<a name="regen.ecocredit.basket.v1.QueryBasketBalanceResponse"></a>

### QueryBasketBalanceResponse
QueryBasketBalanceResponse is the Query/BasketBalance response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balance | [string](#string) |  | balance is the amount of the queried credit batch in the basket. |






<a name="regen.ecocredit.basket.v1.QueryBasketBalancesRequest"></a>

### QueryBasketBalancesRequest
QueryBasketBalancesRequest is the Query/BasketBalances request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basket_denom | [string](#string) |  | basket_denom is the denom of the basket. |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.basket.v1.QueryBasketBalancesResponse"></a>

### QueryBasketBalancesResponse
QueryBasketBalancesResponse is the Query/BasketBalances response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balances | [BasketBalance](#regen.ecocredit.basket.v1.BasketBalance) | repeated | balances is a list of credit balances in the basket. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.basket.v1.QueryBasketRequest"></a>

### QueryBasketRequest
QueryBasketRequest is the Query/Basket request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basket_denom | [string](#string) |  | basket_denom represents the denom of the basket to query. |






<a name="regen.ecocredit.basket.v1.QueryBasketResponse"></a>

### QueryBasketResponse
QueryBasketResponse is the Query/Basket response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basket | [Basket](#regen.ecocredit.basket.v1.Basket) |  | basket is the queried basket. |
| classes | [string](#string) | repeated | classes are the credit classes that can be deposited in the basket. |






<a name="regen.ecocredit.basket.v1.QueryBasketsRequest"></a>

### QueryBasketsRequest
QueryBasketsRequest is the Query/Baskets request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.basket.v1.QueryBasketsResponse"></a>

### QueryBasketsResponse
QueryBasketsResponse is the Query/Baskets response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baskets | [Basket](#regen.ecocredit.basket.v1.Basket) | repeated | baskets are the fetched baskets. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.ecocredit.basket.v1.Query"></a>

### Query
Msg is the regen.ecocredit.basket.v1beta1 Query service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Basket | [QueryBasketRequest](#regen.ecocredit.basket.v1.QueryBasketRequest) | [QueryBasketResponse](#regen.ecocredit.basket.v1.QueryBasketResponse) | Basket queries one basket by denom. |
| Baskets | [QueryBasketsRequest](#regen.ecocredit.basket.v1.QueryBasketsRequest) | [QueryBasketsResponse](#regen.ecocredit.basket.v1.QueryBasketsResponse) | Baskets lists all baskets in the ecocredit module. |
| BasketBalances | [QueryBasketBalancesRequest](#regen.ecocredit.basket.v1.QueryBasketBalancesRequest) | [QueryBasketBalancesResponse](#regen.ecocredit.basket.v1.QueryBasketBalancesResponse) | BasketBalances lists the balance of each credit batch in the basket. |
| BasketBalance | [QueryBasketBalanceRequest](#regen.ecocredit.basket.v1.QueryBasketBalanceRequest) | [QueryBasketBalanceResponse](#regen.ecocredit.basket.v1.QueryBasketBalanceResponse) | BasketBalance queries the balance of a specific credit batch in the basket. |

 <!-- end services -->



<a name="regen/ecocredit/basket/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/basket/v1/tx.proto



<a name="regen.ecocredit.basket.v1.MsgCreate"></a>

### MsgCreate
MsgCreate is the Msg/Create request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| curator | [string](#string) |  | curator is the address of the basket curator who is able to change certain basket settings. |
| name | [string](#string) |  | name will be used to together with prefix to create a bank denom for this basket token. It can be between 3-8 alphanumeric characters, with the first character being alphabetic.

The bank denom will be formed from name, credit type and exponent and be of the form `eco.<prefix><credit_type_abbrev>.<name>` where prefix is derived from exponent. |
| description | [string](#string) |  | description is a human-readable description of the basket denom that should be at most 256 characters. |
| exponent | [uint32](#uint32) |  | exponent is the exponent that will be used for converting credits to basket tokens and for bank denom metadata. It also limits the precision of credit amounts when putting credits into a basket. An exponent of 6 will mean that 10^6 units of a basket token will be issued for 1.0 credits and that this should be displayed as one unit in user interfaces. It also means that the maximum precision of credit amounts is 6 decimal places so that the need to round is eliminated. The exponent must be >= the precision of the credit type at the time the basket is created and be of one of the following values 0, 1, 2, 3, 6, 9, 12, 15, 18, 21, or 24 which correspond to the exponents which have an official SI prefix.

The exponent will be used to form the prefix part of the bank denom and will be mapped as follows: 0 - no prefix 1 - d (deci) 2 - c (centi) 3 - m (milli) 6 - u (micro) 9 - n (nano) 12 - p (pico) 15 - f (femto) 18 - a (atto) 21 - z (zepto) 24 - y (yocto) |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire allows auto-retirement to be disabled. The credits will be auto-retired if disable_auto_retire is false unless the credits were previously put into the basket by the address picking them from the basket, in which case they will remain tradable. |
| credit_type_abbrev | [string](#string) |  | credit_type_abbrev is the abbreviation of the credit type this basket is able to hold. |
| allowed_classes | [string](#string) | repeated | allowed_classes are the credit classes allowed to be put in the basket |
| date_criteria | [DateCriteria](#regen.ecocredit.basket.v1.DateCriteria) |  | date_criteria is the date criteria for batches admitted to the basket. At most, only one of the fields in the date_criteria should be set. |
| fee | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | fee is the fee that the curator will pay to create the basket. It must be >= the required Params.basket_creation_fee. We include the fee explicitly here so that the curator explicitly acknowledges paying this fee and is not surprised to learn that the paid a big fee and didn't know beforehand. |






<a name="regen.ecocredit.basket.v1.MsgCreateResponse"></a>

### MsgCreateResponse
MsgCreateResponse is the Msg/Create response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basket_denom | [string](#string) |  | basket_denom is the unique denomination ID of the newly created basket. |






<a name="regen.ecocredit.basket.v1.MsgPut"></a>

### MsgPut
MsgPut is the Msg/Put request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the owner of credits being put into the basket. |
| basket_denom | [string](#string) |  | basket_denom is the basket denom to add credits to. |
| credits | [BasketCredit](#regen.ecocredit.basket.v1.BasketCredit) | repeated | credits are credits to add to the basket. If they do not match the basket's admission criteria the operation will fail. If there are any "dust" credits left over when converting credits to basket tokens, these credits will not be converted to basket tokens and instead remain with the owner. |






<a name="regen.ecocredit.basket.v1.MsgPutResponse"></a>

### MsgPutResponse
MsgPutResponse is the Msg/Put response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount_received | [string](#string) |  | amount_received is the integer amount of basket tokens received. |






<a name="regen.ecocredit.basket.v1.MsgTake"></a>

### MsgTake
MsgTake is the Msg/Take request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the owner of the basket tokens. |
| basket_denom | [string](#string) |  | basket_denom is the basket bank denom to take credits from. |
| amount | [string](#string) |  | amount is the integer number of basket tokens to convert into credits. |
| retirement_location | [string](#string) |  | retirement_location is the optional retirement location for the credits which will be used only if retire_on_take is true for this basket. |
| retire_on_take | [bool](#bool) |  | retire_on_take is a boolean that dictates whether the ecocredits received in exchange for the basket tokens will be received as retired or tradable credits. |






<a name="regen.ecocredit.basket.v1.MsgTakeResponse"></a>

### MsgTakeResponse
MsgTakeResponse is the Msg/Take response type.


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

