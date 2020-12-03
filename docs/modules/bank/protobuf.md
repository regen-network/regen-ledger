 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/bank/v1alpha1/types.proto](#regen/bank/v1alpha1/types.proto)
    - [Coin](#regen.bank.v1alpha1.Coin)
  
- [regen/bank/v1alpha1/query.proto](#regen/bank/v1alpha1/query.proto)
    - [QueryBalanceRequest](#regen.bank.v1alpha1.QueryBalanceRequest)
    - [QueryBalanceResponse](#regen.bank.v1alpha1.QueryBalanceResponse)
    - [QueryPrecisionRequest](#regen.bank.v1alpha1.QueryPrecisionRequest)
    - [QueryPrecisionResponse](#regen.bank.v1alpha1.QueryPrecisionResponse)
    - [QuerySupplyOfRequest](#regen.bank.v1alpha1.QuerySupplyOfRequest)
    - [QuerySupplyOfResponse](#regen.bank.v1alpha1.QuerySupplyOfResponse)
  
    - [Query](#regen.bank.v1alpha1.Query)
  
- [regen/bank/v1alpha1/rules.proto](#regen/bank/v1alpha1/rules.proto)
    - [ACLRule](#regen.bank.v1alpha1.ACLRule)
    - [BooleanRule](#regen.bank.v1alpha1.BooleanRule)
  
- [regen/bank/v1alpha1/tx.proto](#regen/bank/v1alpha1/tx.proto)
    - [MsgBurnRequest](#regen.bank.v1alpha1.MsgBurnRequest)
    - [MsgBurnResponse](#regen.bank.v1alpha1.MsgBurnResponse)
    - [MsgCreateDenomRequest](#regen.bank.v1alpha1.MsgCreateDenomRequest)
    - [MsgCreateDenomResponse](#regen.bank.v1alpha1.MsgCreateDenomResponse)
    - [MsgMintRequest](#regen.bank.v1alpha1.MsgMintRequest)
    - [MsgMintRequest.Issuance](#regen.bank.v1alpha1.MsgMintRequest.Issuance)
    - [MsgMintResponse](#regen.bank.v1alpha1.MsgMintResponse)
    - [MsgMoveRequest](#regen.bank.v1alpha1.MsgMoveRequest)
    - [MsgMoveResponse](#regen.bank.v1alpha1.MsgMoveResponse)
    - [MsgSendRequest](#regen.bank.v1alpha1.MsgSendRequest)
    - [MsgSendResponse](#regen.bank.v1alpha1.MsgSendResponse)
    - [MsgSetPrecisionRequest](#regen.bank.v1alpha1.MsgSetPrecisionRequest)
    - [MsgSetPrecisionResponse](#regen.bank.v1alpha1.MsgSetPrecisionResponse)
  
    - [Msg](#regen.bank.v1alpha1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/bank/v1alpha1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/bank/v1alpha1/types.proto



<a name="regen.bank.v1alpha1.Coin"></a>

### Coin



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [string](#string) |  |  |
| amount | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/bank/v1alpha1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/bank/v1alpha1/query.proto



<a name="regen.bank.v1alpha1.QueryBalanceRequest"></a>

### QueryBalanceRequest
QueryBalanceRequest is the request type for the Query/Balance RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address is the address to query balances for. |
| denom | [string](#string) |  | denom is the coin denom to query balances for. |






<a name="regen.bank.v1alpha1.QueryBalanceResponse"></a>

### QueryBalanceResponse
QueryBalanceResponse is the response type for the Query/Balance RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balance | [Coin](#regen.bank.v1alpha1.Coin) |  | balance is the balance of the coin. |






<a name="regen.bank.v1alpha1.QueryPrecisionRequest"></a>

### QueryPrecisionRequest
QueryPrecisionRequest is the Query/Precision request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [string](#string) |  | denom is the unique ID of credit batch to query. |






<a name="regen.bank.v1alpha1.QueryPrecisionResponse"></a>

### QueryPrecisionResponse
QueryPrecisionResponse is the Query/Precision response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_decimal_places | [uint32](#uint32) |  | max_decimal_places is the maximum number of decimal places that can be used to represent some quantity of coins. It is an experimental feature to concretely explore an idea proposed in https://github.com/cosmos/cosmos-sdk/issues/7113. |






<a name="regen.bank.v1alpha1.QuerySupplyOfRequest"></a>

### QuerySupplyOfRequest
QuerySupplyOfRequest is the request type for the Query/SupplyOf RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [string](#string) |  | denom is the coin denom to query balances for. |






<a name="regen.bank.v1alpha1.QuerySupplyOfResponse"></a>

### QuerySupplyOfResponse
QuerySupplyOfResponse is the response type for the Query/SupplyOf RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [Coin](#regen.bank.v1alpha1.Coin) |  | amount is the supply of the coin. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.bank.v1alpha1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Balance | [QueryBalanceRequest](#regen.bank.v1alpha1.QueryBalanceRequest) | [QueryBalanceResponse](#regen.bank.v1alpha1.QueryBalanceResponse) | Balance queries the balance of a single coin for a single account. |
| SupplyOf | [QuerySupplyOfRequest](#regen.bank.v1alpha1.QuerySupplyOfRequest) | [QuerySupplyOfResponse](#regen.bank.v1alpha1.QuerySupplyOfResponse) | SupplyOf queries the supply of a single coin. |
| Precision | [QueryPrecisionRequest](#regen.bank.v1alpha1.QueryPrecisionRequest) | [QueryPrecisionResponse](#regen.bank.v1alpha1.QueryPrecisionResponse) | Precision queries the number of decimal places that can be used to represent coins. See Tx/SetPrecision for more details. |

 <!-- end services -->



<a name="regen/bank/v1alpha1/rules.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/bank/v1alpha1/rules.proto



<a name="regen.bank.v1alpha1.ACLRule"></a>

### ACLRule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allowed_addresses | [string](#string) | repeated |  |






<a name="regen.bank.v1alpha1.BooleanRule"></a>

### BooleanRule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/bank/v1alpha1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/bank/v1alpha1/tx.proto



<a name="regen.bank.v1alpha1.MsgBurnRequest"></a>

### MsgBurnRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| burner_address | [string](#string) |  |  |
| coins | [Coin](#regen.bank.v1alpha1.Coin) | repeated |  |






<a name="regen.bank.v1alpha1.MsgBurnResponse"></a>

### MsgBurnResponse







<a name="regen.bank.v1alpha1.MsgCreateDenomRequest"></a>

### MsgCreateDenomRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace_admin | [string](#string) |  |  |
| denom_namespace | [string](#string) |  |  |
| denom_name | [string](#string) |  |  |
| denom_admin | [string](#string) |  | denom_admin specifies an address that has administrative access over the denom and can change important parameters and rules. If left empty, the denom's rules can only be changed by governance. |
| mint_rule | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| send_rule | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| move_rule | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| burn_rule | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| max_decimal_places | [uint32](#uint32) |  |  |






<a name="regen.bank.v1alpha1.MsgCreateDenomResponse"></a>

### MsgCreateDenomResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [string](#string) |  |  |






<a name="regen.bank.v1alpha1.MsgMintRequest"></a>

### MsgMintRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter_address | [string](#string) |  |  |
| issuance | [MsgMintRequest.Issuance](#regen.bank.v1alpha1.MsgMintRequest.Issuance) | repeated |  |






<a name="regen.bank.v1alpha1.MsgMintRequest.Issuance"></a>

### MsgMintRequest.Issuance



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recipient | [string](#string) |  |  |
| coins | [Coin](#regen.bank.v1alpha1.Coin) | repeated |  |






<a name="regen.bank.v1alpha1.MsgMintResponse"></a>

### MsgMintResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_decimal_places | [uint32](#uint32) |  |  |






<a name="regen.bank.v1alpha1.MsgMoveRequest"></a>

### MsgMoveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mover_address | [string](#string) |  |  |
| from_address | [string](#string) |  |  |
| to_address | [string](#string) |  |  |
| amount | [Coin](#regen.bank.v1alpha1.Coin) | repeated |  |






<a name="regen.bank.v1alpha1.MsgMoveResponse"></a>

### MsgMoveResponse







<a name="regen.bank.v1alpha1.MsgSendRequest"></a>

### MsgSendRequest
MsgSendRequest represents a message to send coins from one account to another.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| from_address | [string](#string) |  |  |
| to_address | [string](#string) |  |  |
| amount | [Coin](#regen.bank.v1alpha1.Coin) | repeated |  |






<a name="regen.bank.v1alpha1.MsgSendResponse"></a>

### MsgSendResponse
MsgSendResponse defines the Msg/Send response type.






<a name="regen.bank.v1alpha1.MsgSetPrecisionRequest"></a>

### MsgSetPrecisionRequest
MsgRetireRequest is the Msg/SetPrecision request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_admin | [string](#string) |  | denom_admin is the address of the denom admin. |
| denom | [string](#string) |  | denom is the unique ID of the credit batch. |
| max_decimal_places | [uint32](#uint32) |  | max_decimal_places is the new maximum number of decimal places that can be used to represent some quantity of credit units. It is an experimental feature to concretely explore an idea proposed in https://github.com/cosmos/cosmos-sdk/issues/7113. |






<a name="regen.bank.v1alpha1.MsgSetPrecisionResponse"></a>

### MsgSetPrecisionResponse
MsgSetPrecisionResponse is the Msg/SetPrecision response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.bank.v1alpha1.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateDenom | [MsgCreateDenomRequest](#regen.bank.v1alpha1.MsgCreateDenomRequest) | [MsgCreateDenomResponse](#regen.bank.v1alpha1.MsgCreateDenomResponse) |  |
| Mint | [MsgMintRequest](#regen.bank.v1alpha1.MsgMintRequest) | [MsgMintResponse](#regen.bank.v1alpha1.MsgMintResponse) | Mint is a method for minting new coins. It is subject to each denom's send rule. |
| Move | [MsgMoveRequest](#regen.bank.v1alpha1.MsgMoveRequest) | [MsgMoveResponse](#regen.bank.v1alpha1.MsgMoveResponse) | Move is a method for a module, contract, or other trusted third party to move coins from one account to another account. It is subject to each denom's move rule. |
| Send | [MsgSendRequest](#regen.bank.v1alpha1.MsgSendRequest) | [MsgSendResponse](#regen.bank.v1alpha1.MsgSendResponse) | Send is a method for sending coins from one account to another account. It is subject to each denom's send rule. |
| Burn | [MsgBurnRequest](#regen.bank.v1alpha1.MsgBurnRequest) | [MsgBurnResponse](#regen.bank.v1alpha1.MsgBurnResponse) | Burn is method for burning coins. It is subject to each denom's burn rule. |
| SetPrecision | [MsgSetPrecisionRequest](#regen.bank.v1alpha1.MsgSetPrecisionRequest) | [MsgSetPrecisionResponse](#regen.bank.v1alpha1.MsgSetPrecisionResponse) | SetPrecision allows an issuer to increase the decimal precision of a denom. It is an experimental feature to concretely explore an idea proposed in https://github.com/cosmos/cosmos-sdk/issues/7113. The number of decimal places allowed for a credit batch is determined by the original number of decimal places used with calling CreateDenom. SetPrecision allows the number of allowed decimal places to be increased, effectively making the supply more granular without actually changing any balances. It allows asset issuers to be able to issue an asset without needing to think about how many subdivisions are needed upfront. While it may not be relevant for credits which likely have a fairly stable market value, I wanted to experiment a bit and this serves as a proof of concept for a broader bank redesign where say for instance a coin like the ATOM or XRN could be issued in its own units rather than micro or nano-units. Instead an operation like SetPrecision would allow trading in micro, nano or pico in the future based on market demand. Arbitrary, unbounded precision is not desirable because this can lead to spam attacks (like sending 0.000000000000000000000000000001 coins). This is effectively fixed precision so under the hood it is still basically an integer, but the fixed precision can be increased so its more adaptable long term than just an integer. |

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

