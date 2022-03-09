 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/ecocredit/orderbook/v1alpha1/memory.proto](#regen/ecocredit/orderbook/v1alpha1/memory.proto)
    - [BuyOrderBatchSelector](#regen.ecocredit.orderbook.v1alpha1.BuyOrderBatchSelector)
    - [BuyOrderClassSelector](#regen.ecocredit.orderbook.v1alpha1.BuyOrderClassSelector)
    - [BuyOrderProjectSelector](#regen.ecocredit.orderbook.v1alpha1.BuyOrderProjectSelector)
    - [BuyOrderSellOrderMatch](#regen.ecocredit.orderbook.v1alpha1.BuyOrderSellOrderMatch)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/ecocredit/orderbook/v1alpha1/memory.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/orderbook/v1alpha1/memory.proto



<a name="regen.ecocredit.orderbook.v1alpha1.BuyOrderBatchSelector"></a>

### BuyOrderBatchSelector
BuyOrderBatchSelector indexes a buy order with batch selector.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order_id | [uint64](#uint64) |  | buy_order_id is the buy order ID. |
| batch_id | [uint64](#uint64) |  | batch_id is the batch ID. |






<a name="regen.ecocredit.orderbook.v1alpha1.BuyOrderClassSelector"></a>

### BuyOrderClassSelector
BuyOrderClassSelector indexes a buy order with class selector.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order_id | [uint64](#uint64) |  | buy_order_id is the buy order ID. |
| class_id | [uint64](#uint64) |  | class_id is the class ID. |
| project_location | [string](#string) |  | project_location is the project location in the selector's criteria. |
| min_start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | min_start_date is the minimum start date in the selector's criteria. |
| max_end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | max_end_date is the maximum end date in the selector's criteria. |






<a name="regen.ecocredit.orderbook.v1alpha1.BuyOrderProjectSelector"></a>

### BuyOrderProjectSelector
BuyOrderProjectSelector indexes a buy order with project selector.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order_id | [uint64](#uint64) |  | buy_order_id is the buy order ID. |
| project_id | [uint64](#uint64) |  | project_id is the project ID. |
| min_start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | min_start_date is the minimum start date in the selector's criteria. |
| max_end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | max_end_date is the maximum end date in the selector's criteria. |






<a name="regen.ecocredit.orderbook.v1alpha1.BuyOrderSellOrderMatch"></a>

### BuyOrderSellOrderMatch
BuyOrderSellOrderMatch defines the data the FIFO/price-time-priority matching
algorithm used to actually match buy and sell orders.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| market_id | [uint64](#uint64) |  | market_id defines the market within which this match exists. |
| buy_order_id | [uint64](#uint64) |  | buy_order_id is the buy order ID. |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the sell order ID. |
| bid_price_complement | [fixed32](#fixed32) |  | bid_price_complement is the the complement (^ operator) of the bid price encoded as a uint32 (which should have sufficient precision) - effectively ~price * 10^exponent (usually 10^6). The complement is used so that bids can be sorted high to low. |
| ask_price | [fixed32](#fixed32) |  | ask_price is the ask price encoded to a uint32. Ask prices are sorted low to high. |





 <!-- end messages -->

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

