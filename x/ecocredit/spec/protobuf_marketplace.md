 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/ecocredit/marketplace/v1/events.proto](#regen/ecocredit/marketplace/v1/events.proto)
    - [EventAllowAskDenom](#regen.ecocredit.marketplace.v1.EventAllowAskDenom)
    - [EventBuyOrderCreated](#regen.ecocredit.marketplace.v1.EventBuyOrderCreated)
    - [EventBuyOrderFilled](#regen.ecocredit.marketplace.v1.EventBuyOrderFilled)
    - [EventSell](#regen.ecocredit.marketplace.v1.EventSell)
    - [EventUpdateSellOrder](#regen.ecocredit.marketplace.v1.EventUpdateSellOrder)
  
- [regen/ecocredit/marketplace/v1/types.proto](#regen/ecocredit/marketplace/v1/types.proto)
    - [BatchSelector](#regen.ecocredit.marketplace.v1.BatchSelector)
    - [ClassSelector](#regen.ecocredit.marketplace.v1.ClassSelector)
    - [Filter](#regen.ecocredit.marketplace.v1.Filter)
    - [Filter.Criteria](#regen.ecocredit.marketplace.v1.Filter.Criteria)
    - [ProjectSelector](#regen.ecocredit.marketplace.v1.ProjectSelector)
  
- [regen/ecocredit/marketplace/v1/state.proto](#regen/ecocredit/marketplace/v1/state.proto)
    - [AllowedDenom](#regen.ecocredit.marketplace.v1.AllowedDenom)
    - [BuyOrder](#regen.ecocredit.marketplace.v1.BuyOrder)
    - [BuyOrder.Selection](#regen.ecocredit.marketplace.v1.BuyOrder.Selection)
    - [Market](#regen.ecocredit.marketplace.v1.Market)
    - [SellOrder](#regen.ecocredit.marketplace.v1.SellOrder)
  
- [regen/ecocredit/marketplace/v1/query.proto](#regen/ecocredit/marketplace/v1/query.proto)
    - [QueryAllowedDenomsRequest](#regen.ecocredit.marketplace.v1.QueryAllowedDenomsRequest)
    - [QueryAllowedDenomsResponse](#regen.ecocredit.marketplace.v1.QueryAllowedDenomsResponse)
    - [QueryBuyOrderRequest](#regen.ecocredit.marketplace.v1.QueryBuyOrderRequest)
    - [QueryBuyOrderResponse](#regen.ecocredit.marketplace.v1.QueryBuyOrderResponse)
    - [QueryBuyOrdersByAddressRequest](#regen.ecocredit.marketplace.v1.QueryBuyOrdersByAddressRequest)
    - [QueryBuyOrdersByAddressResponse](#regen.ecocredit.marketplace.v1.QueryBuyOrdersByAddressResponse)
    - [QueryBuyOrdersRequest](#regen.ecocredit.marketplace.v1.QueryBuyOrdersRequest)
    - [QueryBuyOrdersResponse](#regen.ecocredit.marketplace.v1.QueryBuyOrdersResponse)
    - [QuerySellOrderRequest](#regen.ecocredit.marketplace.v1.QuerySellOrderRequest)
    - [QuerySellOrderResponse](#regen.ecocredit.marketplace.v1.QuerySellOrderResponse)
    - [QuerySellOrdersByAddressRequest](#regen.ecocredit.marketplace.v1.QuerySellOrdersByAddressRequest)
    - [QuerySellOrdersByAddressResponse](#regen.ecocredit.marketplace.v1.QuerySellOrdersByAddressResponse)
    - [QuerySellOrdersByBatchDenomRequest](#regen.ecocredit.marketplace.v1.QuerySellOrdersByBatchDenomRequest)
    - [QuerySellOrdersByBatchDenomResponse](#regen.ecocredit.marketplace.v1.QuerySellOrdersByBatchDenomResponse)
    - [QuerySellOrdersRequest](#regen.ecocredit.marketplace.v1.QuerySellOrdersRequest)
    - [QuerySellOrdersResponse](#regen.ecocredit.marketplace.v1.QuerySellOrdersResponse)
  
    - [Query](#regen.ecocredit.marketplace.v1.Query)
  
- [regen/ecocredit/marketplace/v1/tx.proto](#regen/ecocredit/marketplace/v1/tx.proto)
    - [MsgAllowAskDenom](#regen.ecocredit.marketplace.v1.MsgAllowAskDenom)
    - [MsgAllowAskDenomResponse](#regen.ecocredit.marketplace.v1.MsgAllowAskDenomResponse)
    - [MsgBuy](#regen.ecocredit.marketplace.v1.MsgBuy)
    - [MsgBuy.Order](#regen.ecocredit.marketplace.v1.MsgBuy.Order)
    - [MsgBuy.Order.Selection](#regen.ecocredit.marketplace.v1.MsgBuy.Order.Selection)
    - [MsgBuyResponse](#regen.ecocredit.marketplace.v1.MsgBuyResponse)
    - [MsgSell](#regen.ecocredit.marketplace.v1.MsgSell)
    - [MsgSell.Order](#regen.ecocredit.marketplace.v1.MsgSell.Order)
    - [MsgSellResponse](#regen.ecocredit.marketplace.v1.MsgSellResponse)
    - [MsgUpdateSellOrders](#regen.ecocredit.marketplace.v1.MsgUpdateSellOrders)
    - [MsgUpdateSellOrders.Update](#regen.ecocredit.marketplace.v1.MsgUpdateSellOrders.Update)
    - [MsgUpdateSellOrdersResponse](#regen.ecocredit.marketplace.v1.MsgUpdateSellOrdersResponse)
  
    - [Msg](#regen.ecocredit.marketplace.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/ecocredit/marketplace/v1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/marketplace/v1/events.proto



<a name="regen.ecocredit.marketplace.v1.EventAllowAskDenom"></a>

### EventAllowAskDenom
EventAllowAskDenom is an event emitted when an ask denom is added.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [string](#string) |  | denom is the denom to allow (ex. ibc/GLKHDSG423SGS) |
| display_denom | [string](#string) |  | display_denom is the denom to display to the user and is informational |
| exponent | [uint32](#uint32) |  | exponent is the exponent that relates the denom to the display_denom and is informational |






<a name="regen.ecocredit.marketplace.v1.EventBuyOrderCreated"></a>

### EventBuyOrderCreated
EventBuyOrderCreated is an event emitted when a buy order is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order_id | [uint64](#uint64) |  | buy_order_id is the unique ID of buy order. |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the sell order ID against which the buyer is trying to buy. |
| quantity | [string](#string) |  | quantity is the quantity of credits to buy. If the quantity of credits available is less than this amount the order will be partially filled unless disable_partial_fill is true. |
| bid_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | bid price is the bid price for this buy order. A credit unit will be settled at a purchase price that is no more than the bid price. The buy order will fail if the buyer does not have enough funds available to complete the purchase. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire allows auto-retirement to be disabled. If it is set to true the credits will not auto-retire and can be resold assuming that the corresponding sell order has auto-retirement disabled. If the sell order hasn't disabled auto-retirement and the buy order tries to disable it, that buy order will fail. |
| disable_partial_fill | [bool](#bool) |  | disable_partial_fill disables the default behavior of partially filling buy orders if the requested quantity is not available. |
| retirement_location | [string](#string) |  | retirement_location is the optional retirement location for the credits which will be used only if disable_auto_retire is false. |
| expiration | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | expiration is the optional timestamp when the buy order expires. When the expiration time is reached, the buy order is removed from state. |






<a name="regen.ecocredit.marketplace.v1.EventBuyOrderFilled"></a>

### EventBuyOrderFilled
EventBuyOrderFilled is an event emitted when a buy order is filled.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order_id | [uint64](#uint64) |  | buy_order_id is the unique ID of the buy order. |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the unique ID of the sell order. |
| batch_denom | [string](#string) |  | batch_denom is the credit batch ID of the purchased credits. |
| quantity | [string](#string) |  | quantity is the quantity of the purchased credits. |
| total_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total_price is the total price for the purchased credits. |






<a name="regen.ecocredit.marketplace.v1.EventSell"></a>

### EventSell
EventSell is an event emitted when a sell order is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [uint64](#uint64) |  | order_id is the unique ID of sell order. |
| batch_denom | [string](#string) |  | batch_denom is the credit batch being sold. |
| quantity | [string](#string) |  | quantity is the quantity of credits being sold. |
| ask_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | ask_price is the price the seller is asking for each unit of the batch_denom. Each credit unit of the batch will be sold for at least the ask_price or more. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire disables auto-retirement of credits which allows a buyer to disable auto-retirement in their buy order enabling them to resell the credits to another buyer. |
| expiration | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | expiration is an optional timestamp when the sell order expires. When the expiration time is reached, the sell order is removed from state. |






<a name="regen.ecocredit.marketplace.v1.EventUpdateSellOrder"></a>

### EventUpdateSellOrder
EventUpdateSellOrder is an event emitted when a sell order is updated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the owner of the sell orders. |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the ID of an existing sell order. |
| batch_denom | [string](#string) |  | batch_denom is the credit batch being sold. |
| new_quantity | [string](#string) |  | new_quantity is the updated quantity of credits available to sell, if it is set to zero then the order is cancelled. |
| new_ask_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | new_ask_price is the new ask price for this sell order |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire updates the disable_auto_retire field in the sell order. |
| new_expiration | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | new_expiration is an optional timestamp when the sell order expires. When the expiration time is reached, the sell order is removed from state. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/marketplace/v1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/marketplace/v1/types.proto



<a name="regen.ecocredit.marketplace.v1.BatchSelector"></a>

### BatchSelector
BatchSelector is a selector for a credit batch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_id | [uint64](#uint64) |  | batch_id is the credit batch ID. |






<a name="regen.ecocredit.marketplace.v1.ClassSelector"></a>

### ClassSelector
ClassSelector is a selector for a credit class.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [uint64](#uint64) |  | class_id is the credit class ID. |
| project_location | [string](#string) |  | project_location can be specified in three levels of granularity: country, sub-national-code, or postal code. If just country is given, for instance "US" then any credits in the "US" will be matched even their project location is more specific, ex. "US-NY 12345". If a country, sub-national-code and postal code are all provided then only projects in that postal code will match. |
| min_start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_date is the beginning of the period during which a credit batch was quantified and verified. If it is empty then there is no start date limit. |
| max_end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | max_end_date is the end of the period during which a credit batch was quantified and verified. If it is empty then there is no end date limit. |






<a name="regen.ecocredit.marketplace.v1.Filter"></a>

### Filter
Filter is used to create filtered buy orders which match credit batch
sell orders based on selection criteria rather than matching individual
sell orders


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| or | [Filter.Criteria](#regen.ecocredit.marketplace.v1.Filter.Criteria) | repeated | or is a list of criteria for matching credit batches. A credit which matches this filter must match at least one of these criteria. |






<a name="regen.ecocredit.marketplace.v1.Filter.Criteria"></a>

### Filter.Criteria
Criteria is a simple filter criteria for matching a credit batch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_selector | [ClassSelector](#regen.ecocredit.marketplace.v1.ClassSelector) |  | class_selector is a credit class selector. |
| project_selector | [ProjectSelector](#regen.ecocredit.marketplace.v1.ProjectSelector) |  | project_selector is a project selector. |
| batch_selector | [BatchSelector](#regen.ecocredit.marketplace.v1.BatchSelector) |  | batch_selector is a credit batch selector. |






<a name="regen.ecocredit.marketplace.v1.ProjectSelector"></a>

### ProjectSelector
ProjectSelector is a selector for a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [uint64](#uint64) |  | project_id is the project ID. |
| min_start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_date is the beginning of the period during which a credit batch was quantified and verified. If it is empty then there is no start date limit. |
| max_end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | max_end_date is the end of the period during which a credit batch was quantified and verified. If it is empty then there is no end date limit. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/marketplace/v1/state.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/marketplace/v1/state.proto



<a name="regen.ecocredit.marketplace.v1.AllowedDenom"></a>

### AllowedDenom
AllowedDenom represents the information for an allowed ask/bid denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bank_denom | [string](#string) |  | denom is the bank denom to allow (ex. ibc/GLKHDSG423SGS) |
| display_denom | [string](#string) |  | display_denom is the denom to display to the user and is informational. Because the denom is likely an IBC denom, this should be chosen by governance to represent the consensus trusted name of the denom. |
| exponent | [uint32](#uint32) |  | exponent is the exponent that relates the denom to the display_denom and is informational |






<a name="regen.ecocredit.marketplace.v1.BuyOrder"></a>

### BuyOrder
BuyOrder represents the information for a buy order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id is the unique ID of buy order. |
| buyer | [bytes](#bytes) |  | buyer is the bytes address of the account that created the buy order |
| selection | [BuyOrder.Selection](#regen.ecocredit.marketplace.v1.BuyOrder.Selection) |  | selection is the buy order selection. |
| quantity | [string](#string) |  | quantity is the decimal quantity of credits to buy. If the quantity of credits available is less than this amount the order will be partially filled unless disable_partial_fill is true. |
| market_id | [uint64](#uint64) |  | market_id is the market in which this sell order exists and specifies the bank_denom that ask_price corresponds to. |
| bid_price | [string](#string) |  | bid price is the integer bid price for this buy order. A credit unit will be settled at a purchase price that is no more than the bid price. The buy order will fail if the buyer does not have enough funds available to complete the purchase. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire allows auto-retirement to be disabled. If it is set to true the credits will not auto-retire and can be resold assuming that the corresponding sell order has auto-retirement disabled. If the sell order hasn't disabled auto-retirement and the buy order tries to disable it, that buy order will fail. |
| disable_partial_fill | [bool](#bool) |  | disable_partial_fill disables the default behavior of partially filling buy orders if the requested quantity is not available. |
| expiration | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | expiration is the optional timestamp when the buy order expires. When the expiration time is reached, the buy order is removed from state. |
| maker | [bool](#bool) |  | maker indicates that this is a maker order, meaning that when it hit the order book, there were no matching sell orders. |






<a name="regen.ecocredit.marketplace.v1.BuyOrder.Selection"></a>

### BuyOrder.Selection
Selection defines a buy order selection.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the sell order ID against which the buyer is trying to buy. When sell_order_id is set, this is known as a direct buy order because it is placed directly against a specific sell order. |
| filter | [Filter](#regen.ecocredit.marketplace.v1.Filter) |  | filter selects credits to buy based upon the specified filter criteria. |






<a name="regen.ecocredit.marketplace.v1.Market"></a>

### Market
Market describes a distinctly processed market between a credit type and
allowed bank denom. Each market has its own precision in the order book
and is processed independently of other markets. Governance must enable
markets one by one. Every additional enabled market potentially adds more
processing overhead to the blockchain and potentially weakens liquidity in
competing markets. For instance, enabling side by side USD/Carbon and
EUR/Carbon markets may have the end result that each market individually has
less liquidity and longer settlement times. Such decisions should be taken
with care.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id is the unique ID of the market. |
| credit_type | [string](#string) |  | credit_type is the abbreviation of the credit type. |
| bank_denom | [string](#string) |  | bank_denom is an allowed bank denom. |
| precision_modifier | [uint32](#uint32) |  | precision_modifier is an optional modifier used to convert arbitrary precision integer bank amounts to uint32 values used for sorting in the order book. Given an arbitrary precision integer x, its uint32 conversion will be x / 10^precision_modifier using round half away from zero rounding.

uint32 values range from 0 to 4,294,967,295. This allows for a full 9 digits of precision. In most real world markets this amount of precision is sufficient and most common downside - that some orders with very miniscule price differences may be ordered equivalently (because of rounding) - is acceptable. Note that this rounding will not affect settlement price which will always be done exactly.

Given a USD stable coin with 6 decimal digits, a precision_modifier of 0 is probably acceptable as long as credits are always less than $4,294/unit. With precision down to $0.001 (a precision_modifier of 3 in this case), prices can rise up to $4,294,000/unit. Either scenario is probably quite acceptable given that carbon prices are unlikely to rise above $1000/ton any time in the near future.

If credit prices, exceed the maximum range of uint32 with this precision_modifier, orders with high prices will fail and governance will need to adjust precision_modifier to allow for higher prices in exchange for less precision at the lower end. |






<a name="regen.ecocredit.marketplace.v1.SellOrder"></a>

### SellOrder
SellOrder represents the information for a sell order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id is the unique ID of sell order. |
| seller | [bytes](#bytes) |  | seller is the bytes address of the owner of the credits being sold. |
| batch_id | [uint64](#uint64) |  | batch_id is ID of the credit batch being sold. |
| quantity | [string](#string) |  | quantity is the decimal quantity of credits being sold. |
| market_id | [uint64](#uint64) |  | market_id is the market in which this sell order exists and specifies the bank_denom that ask_price corresponds to. |
| ask_price | [string](#string) |  | ask_price is the integer price (encoded as a string) the seller is asking for each unit of the batch_denom. Each credit unit of the batch will be sold for at least the ask_price or more. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire disables auto-retirement of credits which allows a buyer to disable auto-retirement in their buy order enabling them to resell the credits to another buyer. |
| expiration | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | expiration is an optional timestamp when the sell order expires. When the expiration time is reached, the sell order is removed from state. |
| maker | [bool](#bool) |  | maker indicates that this is a maker order, meaning that when it hit the order book, there were no matching buy orders. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/marketplace/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/marketplace/v1/query.proto



<a name="regen.ecocredit.marketplace.v1.QueryAllowedDenomsRequest"></a>

### QueryAllowedDenomsRequest
QueryAllowedDenomsRequest is the Query/AllowedDenoms request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.marketplace.v1.QueryAllowedDenomsResponse"></a>

### QueryAllowedDenomsResponse
QueryAllowedDenomsResponse is the Query/AllowedDenoms response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allowed_denoms | [AllowedDenom](#regen.ecocredit.marketplace.v1.AllowedDenom) | repeated | allowed_denoms is a list of coin denoms allowed to use in the ask price of sell orders. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an optional pagination for the response. |






<a name="regen.ecocredit.marketplace.v1.QueryBuyOrderRequest"></a>

### QueryBuyOrderRequest
QueryBuyOrderRequest is the Query/BuyOrder request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order_id | [uint64](#uint64) |  | buy_order_id is the id of the buy order. |






<a name="regen.ecocredit.marketplace.v1.QueryBuyOrderResponse"></a>

### QueryBuyOrderResponse
QueryBuyOrderResponse is the Query/BuyOrder response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order | [BuyOrder](#regen.ecocredit.marketplace.v1.BuyOrder) |  | buy_order contains all information related to a buy order. |






<a name="regen.ecocredit.marketplace.v1.QueryBuyOrdersByAddressRequest"></a>

### QueryBuyOrdersByAddressRequest
QueryBuyOrdersByAddressRequest is the Query/BuyOrdersByAddress request type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address of the buy order creator |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.marketplace.v1.QueryBuyOrdersByAddressResponse"></a>

### QueryBuyOrdersByAddressResponse
QueryBuyOrdersByAddressResponse is the Query/BuyOrdersByAddress response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_orders | [BuyOrder](#regen.ecocredit.marketplace.v1.BuyOrder) | repeated | buy_orders is a list of buy orders. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an optional pagination for the response. |






<a name="regen.ecocredit.marketplace.v1.QueryBuyOrdersRequest"></a>

### QueryBuyOrdersRequest
QueryBuyOrdersRequest is the Query/BuyOrders request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.marketplace.v1.QueryBuyOrdersResponse"></a>

### QueryBuyOrdersResponse
QueryBuyOrdersResponse is the Query/BuyOrders response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_orders | [BuyOrder](#regen.ecocredit.marketplace.v1.BuyOrder) | repeated | buy_orders is a list of buy orders. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an optional pagination for the response. |






<a name="regen.ecocredit.marketplace.v1.QuerySellOrderRequest"></a>

### QuerySellOrderRequest
QuerySellOrderRequest is the Query/SellOrder request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the id of the requested sell order. |






<a name="regen.ecocredit.marketplace.v1.QuerySellOrderResponse"></a>

### QuerySellOrderResponse
QuerySellOrderResponse is the Query/SellOrder response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_order | [SellOrder](#regen.ecocredit.marketplace.v1.SellOrder) |  | sell_order contains all information related to a sell order. |






<a name="regen.ecocredit.marketplace.v1.QuerySellOrdersByAddressRequest"></a>

### QuerySellOrdersByAddressRequest
QuerySellOrdersByAddressRequest is the Query/SellOrdersByAddress request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address is the creator of the sell order |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.marketplace.v1.QuerySellOrdersByAddressResponse"></a>

### QuerySellOrdersByAddressResponse
QuerySellOrdersByAddressResponse is the Query/SellOrdersByAddressResponse response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_orders | [SellOrder](#regen.ecocredit.marketplace.v1.SellOrder) | repeated | sell_orders is a list of sell orders. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an optional pagination for the response. |






<a name="regen.ecocredit.marketplace.v1.QuerySellOrdersByBatchDenomRequest"></a>

### QuerySellOrdersByBatchDenomRequest
QuerySellOrdersByDenomRequest is the Query/SellOrdersByDenom request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is an ecocredit denom |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.marketplace.v1.QuerySellOrdersByBatchDenomResponse"></a>

### QuerySellOrdersByBatchDenomResponse
QuerySellOrdersByDenomResponse is the Query/SellOrdersByDenom response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_orders | [SellOrder](#regen.ecocredit.marketplace.v1.SellOrder) | repeated | sell_orders is a list of sell orders. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an optional pagination for the response. |






<a name="regen.ecocredit.marketplace.v1.QuerySellOrdersRequest"></a>

### QuerySellOrdersRequest
QuerySellOrdersRequest is the Query/SellOrders request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.marketplace.v1.QuerySellOrdersResponse"></a>

### QuerySellOrdersResponse
QuerySellOrdersResponse is the Query/SellOrders response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_orders | [SellOrder](#regen.ecocredit.marketplace.v1.SellOrder) | repeated | sell_orders is a list of sell orders. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.ecocredit.marketplace.v1.Query"></a>

### Query
Msg is the regen.ecocredit.marketplace.v1 Query service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SellOrder | [QuerySellOrderRequest](#regen.ecocredit.marketplace.v1.QuerySellOrderRequest) | [QuerySellOrderResponse](#regen.ecocredit.marketplace.v1.QuerySellOrderResponse) | SellOrder queries a sell order by its ID |
| SellOrders | [QuerySellOrdersRequest](#regen.ecocredit.marketplace.v1.QuerySellOrdersRequest) | [QuerySellOrdersResponse](#regen.ecocredit.marketplace.v1.QuerySellOrdersResponse) | SellOrders queries a paginated list of all sell orders |
| SellOrdersByBatchDenom | [QuerySellOrdersByBatchDenomRequest](#regen.ecocredit.marketplace.v1.QuerySellOrdersByBatchDenomRequest) | [QuerySellOrdersByBatchDenomResponse](#regen.ecocredit.marketplace.v1.QuerySellOrdersByBatchDenomResponse) | SellOrdersByDenom queries a paginated list of all sell orders of a specific ecocredit denom |
| SellOrdersByAddress | [QuerySellOrdersByAddressRequest](#regen.ecocredit.marketplace.v1.QuerySellOrdersByAddressRequest) | [QuerySellOrdersByAddressResponse](#regen.ecocredit.marketplace.v1.QuerySellOrdersByAddressResponse) | SellOrdersByAddress queries a paginated list of all sell orders from a specific address |
| BuyOrder | [QueryBuyOrderRequest](#regen.ecocredit.marketplace.v1.QueryBuyOrderRequest) | [QueryBuyOrderResponse](#regen.ecocredit.marketplace.v1.QueryBuyOrderResponse) | BuyOrder queries a buy order by its id |
| BuyOrders | [QueryBuyOrdersRequest](#regen.ecocredit.marketplace.v1.QueryBuyOrdersRequest) | [QueryBuyOrdersResponse](#regen.ecocredit.marketplace.v1.QueryBuyOrdersResponse) | BuyOrders queries a paginated list of all buy orders |
| BuyOrdersByAddress | [QueryBuyOrdersByAddressRequest](#regen.ecocredit.marketplace.v1.QueryBuyOrdersByAddressRequest) | [QueryBuyOrdersByAddressResponse](#regen.ecocredit.marketplace.v1.QueryBuyOrdersByAddressResponse) | BuyOrdersByAddress queries a paginated list of buy orders by creator address |
| AllowedDenoms | [QueryAllowedDenomsRequest](#regen.ecocredit.marketplace.v1.QueryAllowedDenomsRequest) | [QueryAllowedDenomsResponse](#regen.ecocredit.marketplace.v1.QueryAllowedDenomsResponse) | AllowedDenoms queries all denoms allowed to be set in the AskPrice of a sell order |

 <!-- end services -->



<a name="regen/ecocredit/marketplace/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/marketplace/v1/tx.proto



<a name="regen.ecocredit.marketplace.v1.MsgAllowAskDenom"></a>

### MsgAllowAskDenom
MsgAllowAskDenom is the Msg/AllowAskDenom request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root_address | [string](#string) |  | root_address is the address of the governance account which can authorize ask denoms |
| denom | [string](#string) |  | denom is the denom to allow (ex. ibc/GLKHDSG423SGS) |
| display_denom | [string](#string) |  | display_denom is the denom to display to the user and is informational |
| exponent | [uint32](#uint32) |  | exponent is the exponent that relates the denom to the display_denom and is informational |






<a name="regen.ecocredit.marketplace.v1.MsgAllowAskDenomResponse"></a>

### MsgAllowAskDenomResponse
MsgAllowAskDenomResponse is the Msg/AllowAskDenom response type.






<a name="regen.ecocredit.marketplace.v1.MsgBuy"></a>

### MsgBuy
MsgBuy is the Msg/Buy request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buyer | [string](#string) |  | buyer is the address of the credit buyer. |
| orders | [MsgBuy.Order](#regen.ecocredit.marketplace.v1.MsgBuy.Order) | repeated | orders are the new buy orders. |






<a name="regen.ecocredit.marketplace.v1.MsgBuy.Order"></a>

### MsgBuy.Order
Order is a buy order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selection | [MsgBuy.Order.Selection](#regen.ecocredit.marketplace.v1.MsgBuy.Order.Selection) |  | selection is the buy order selection. |
| quantity | [string](#string) |  | quantity is the quantity of credits to buy. If the quantity of credits available is less than this amount the order will be partially filled unless disable_partial_fill is true. |
| bid_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | bid price is the bid price for this buy order. A credit unit will be settled at a purchase price that is no more than the bid price. The buy order will fail if the buyer does not have enough funds available to complete the purchase. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire allows auto-retirement to be disabled. If it is set to true the credits will not auto-retire and can be resold assuming that the corresponding sell order has auto-retirement disabled. If the sell order hasn't disabled auto-retirement and the buy order tries to disable it, that buy order will fail. |
| disable_partial_fill | [bool](#bool) |  | disable_partial_fill disables the default behavior of partially filling buy orders if the requested quantity is not available. |
| retirement_location | [string](#string) |  | retirement_location is the optional retirement location for the credits which will be used only if disable_auto_retire is false. |
| expiration | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | expiration is the optional timestamp when the buy order expires. When the expiration time is reached, the buy order is removed from state. |






<a name="regen.ecocredit.marketplace.v1.MsgBuy.Order.Selection"></a>

### MsgBuy.Order.Selection
Selection defines a buy order selection.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the sell order ID against which the buyer is trying to buy. When sell_order_id is set, this is known as a direct buy order because it is placed directly against a specific sell order. |
| filter | [Filter](#regen.ecocredit.marketplace.v1.Filter) |  | filter selects credits to buy based upon the specified filter criteria. |






<a name="regen.ecocredit.marketplace.v1.MsgBuyResponse"></a>

### MsgBuyResponse
MsgBuyResponse is the Msg/Buy response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order_ids | [uint64](#uint64) | repeated | buy_order_ids are the buy order IDs of the newly created buy orders. Buy orders may not settle instantaneously, but rather in batches at specified batch epoch times. |






<a name="regen.ecocredit.marketplace.v1.MsgSell"></a>

### MsgSell
MsgSell is the Msg/Sell request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the address of the owner of the credits being sold. |
| orders | [MsgSell.Order](#regen.ecocredit.marketplace.v1.MsgSell.Order) | repeated | orders are the sell orders being created. |






<a name="regen.ecocredit.marketplace.v1.MsgSell.Order"></a>

### MsgSell.Order
Order is the content of a new sell order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the credit batch being sold. |
| quantity | [string](#string) |  | quantity is the quantity of credits being sold from this batch. If it is less then the balance of credits the owner has available at the time this sell order is matched, the quantity will be adjusted downwards to the owner's balance. However, if the balance of credits is less than this quantity at the time the sell order is created, the operation will fail. |
| ask_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | ask_price is the price the seller is asking for each unit of the batch_denom. Each credit unit of the batch will be sold for at least the ask_price or more. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire disables auto-retirement of credits which allows a buyer to disable auto-retirement in their buy order enabling them to resell the credits to another buyer. |
| expiration | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | expiration is an optional timestamp when the sell order expires. When the expiration time is reached, the sell order is removed from state. |






<a name="regen.ecocredit.marketplace.v1.MsgSellResponse"></a>

### MsgSellResponse
MsgSellResponse is the Msg/Sell response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_order_ids | [uint64](#uint64) | repeated | sell_order_ids are the sell order IDs of the newly created sell orders. |






<a name="regen.ecocredit.marketplace.v1.MsgUpdateSellOrders"></a>

### MsgUpdateSellOrders
MsgUpdateSellOrders is the Msg/UpdateSellOrders request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the owner of the sell orders. |
| updates | [MsgUpdateSellOrders.Update](#regen.ecocredit.marketplace.v1.MsgUpdateSellOrders.Update) | repeated | updates are updates to existing sell orders. |






<a name="regen.ecocredit.marketplace.v1.MsgUpdateSellOrders.Update"></a>

### MsgUpdateSellOrders.Update
Update is an update to an existing sell order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the ID of an existing sell order. |
| new_quantity | [string](#string) |  | new_quantity is the updated quantity of credits available to sell, if it is set to zero then the order is cancelled. |
| new_ask_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | new_ask_price is the new ask price for this sell order |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire updates the disable_auto_retire field in the sell order. |
| new_expiration | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | new_expiration is an optional timestamp when the sell order expires. When the expiration time is reached, the sell order is removed from state. |






<a name="regen.ecocredit.marketplace.v1.MsgUpdateSellOrdersResponse"></a>

### MsgUpdateSellOrdersResponse
MsgUpdateSellOrdersResponse is the Msg/UpdateSellOrders response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.ecocredit.marketplace.v1.Msg"></a>

### Msg
Msg is the regen.ecocredit.marketplace.v1 Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Sell | [MsgSell](#regen.ecocredit.marketplace.v1.MsgSell) | [MsgSellResponse](#regen.ecocredit.marketplace.v1.MsgSellResponse) | Sell creates new sell orders. |
| UpdateSellOrders | [MsgUpdateSellOrders](#regen.ecocredit.marketplace.v1.MsgUpdateSellOrders) | [MsgUpdateSellOrdersResponse](#regen.ecocredit.marketplace.v1.MsgUpdateSellOrdersResponse) | UpdateSellOrders updates existing sell orders. |
| Buy | [MsgBuy](#regen.ecocredit.marketplace.v1.MsgBuy) | [MsgBuyResponse](#regen.ecocredit.marketplace.v1.MsgBuyResponse) | Buy creates credit buy orders. |
| AllowAskDenom | [MsgAllowAskDenom](#regen.ecocredit.marketplace.v1.MsgAllowAskDenom) | [MsgAllowAskDenomResponse](#regen.ecocredit.marketplace.v1.MsgAllowAskDenomResponse) | AllowAskDenom is a governance operation which authorizes a new ask denom to be used in sell orders |

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

