 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/ecocredit/v1alpha2/events.proto](#regen/ecocredit/v1alpha2/events.proto)
    - [EventAllowAskDenom](#regen.ecocredit.v1alpha2.EventAllowAskDenom)
    - [EventBuyOrderCreated](#regen.ecocredit.v1alpha2.EventBuyOrderCreated)
    - [EventBuyOrderFilled](#regen.ecocredit.v1alpha2.EventBuyOrderFilled)
    - [EventCancel](#regen.ecocredit.v1alpha2.EventCancel)
    - [EventCreateBatch](#regen.ecocredit.v1alpha2.EventCreateBatch)
    - [EventCreateClass](#regen.ecocredit.v1alpha2.EventCreateClass)
    - [EventCreateProject](#regen.ecocredit.v1alpha2.EventCreateProject)
    - [EventReceive](#regen.ecocredit.v1alpha2.EventReceive)
    - [EventRetire](#regen.ecocredit.v1alpha2.EventRetire)
    - [EventSell](#regen.ecocredit.v1alpha2.EventSell)
    - [EventUpdateSellOrder](#regen.ecocredit.v1alpha2.EventUpdateSellOrder)
  
- [regen/ecocredit/v1alpha2/types.proto](#regen/ecocredit/v1alpha2/types.proto)
    - [AskDenom](#regen.ecocredit.v1alpha2.AskDenom)
    - [BasketCredit](#regen.ecocredit.v1alpha2.BasketCredit)
    - [BasketCriteria](#regen.ecocredit.v1alpha2.BasketCriteria)
    - [BatchInfo](#regen.ecocredit.v1alpha2.BatchInfo)
    - [BuyOrder](#regen.ecocredit.v1alpha2.BuyOrder)
    - [BuyOrder.Selection](#regen.ecocredit.v1alpha2.BuyOrder.Selection)
    - [ClassInfo](#regen.ecocredit.v1alpha2.ClassInfo)
    - [CreditType](#regen.ecocredit.v1alpha2.CreditType)
    - [CreditTypeSeq](#regen.ecocredit.v1alpha2.CreditTypeSeq)
    - [Filter](#regen.ecocredit.v1alpha2.Filter)
    - [Filter.And](#regen.ecocredit.v1alpha2.Filter.And)
    - [Filter.DateRange](#regen.ecocredit.v1alpha2.Filter.DateRange)
    - [Filter.Or](#regen.ecocredit.v1alpha2.Filter.Or)
    - [Params](#regen.ecocredit.v1alpha2.Params)
    - [ProjectInfo](#regen.ecocredit.v1alpha2.ProjectInfo)
    - [SellOrder](#regen.ecocredit.v1alpha2.SellOrder)
  
- [regen/ecocredit/v1alpha2/genesis.proto](#regen/ecocredit/v1alpha2/genesis.proto)
    - [Balance](#regen.ecocredit.v1alpha2.Balance)
    - [GenesisState](#regen.ecocredit.v1alpha2.GenesisState)
    - [Supply](#regen.ecocredit.v1alpha2.Supply)
  
- [regen/ecocredit/v1alpha2/query.proto](#regen/ecocredit/v1alpha2/query.proto)
    - [Basket](#regen.ecocredit.v1alpha2.Basket)
    - [QueryAllowedAskDenomsRequest](#regen.ecocredit.v1alpha2.QueryAllowedAskDenomsRequest)
    - [QueryAllowedAskDenomsResponse](#regen.ecocredit.v1alpha2.QueryAllowedAskDenomsResponse)
    - [QueryBalanceRequest](#regen.ecocredit.v1alpha2.QueryBalanceRequest)
    - [QueryBalanceResponse](#regen.ecocredit.v1alpha2.QueryBalanceResponse)
    - [QueryBasketCreditsRequest](#regen.ecocredit.v1alpha2.QueryBasketCreditsRequest)
    - [QueryBasketCreditsResponse](#regen.ecocredit.v1alpha2.QueryBasketCreditsResponse)
    - [QueryBasketRequest](#regen.ecocredit.v1alpha2.QueryBasketRequest)
    - [QueryBasketResponse](#regen.ecocredit.v1alpha2.QueryBasketResponse)
    - [QueryBasketsRequest](#regen.ecocredit.v1alpha2.QueryBasketsRequest)
    - [QueryBasketsResponse](#regen.ecocredit.v1alpha2.QueryBasketsResponse)
    - [QueryBatchInfoRequest](#regen.ecocredit.v1alpha2.QueryBatchInfoRequest)
    - [QueryBatchInfoResponse](#regen.ecocredit.v1alpha2.QueryBatchInfoResponse)
    - [QueryBatchesRequest](#regen.ecocredit.v1alpha2.QueryBatchesRequest)
    - [QueryBatchesResponse](#regen.ecocredit.v1alpha2.QueryBatchesResponse)
    - [QueryBuyOrderRequest](#regen.ecocredit.v1alpha2.QueryBuyOrderRequest)
    - [QueryBuyOrderResponse](#regen.ecocredit.v1alpha2.QueryBuyOrderResponse)
    - [QueryBuyOrdersByAddressRequest](#regen.ecocredit.v1alpha2.QueryBuyOrdersByAddressRequest)
    - [QueryBuyOrdersByAddressResponse](#regen.ecocredit.v1alpha2.QueryBuyOrdersByAddressResponse)
    - [QueryBuyOrdersRequest](#regen.ecocredit.v1alpha2.QueryBuyOrdersRequest)
    - [QueryBuyOrdersResponse](#regen.ecocredit.v1alpha2.QueryBuyOrdersResponse)
    - [QueryClassInfoRequest](#regen.ecocredit.v1alpha2.QueryClassInfoRequest)
    - [QueryClassInfoResponse](#regen.ecocredit.v1alpha2.QueryClassInfoResponse)
    - [QueryClassesRequest](#regen.ecocredit.v1alpha2.QueryClassesRequest)
    - [QueryClassesResponse](#regen.ecocredit.v1alpha2.QueryClassesResponse)
    - [QueryCreditTypesRequest](#regen.ecocredit.v1alpha2.QueryCreditTypesRequest)
    - [QueryCreditTypesResponse](#regen.ecocredit.v1alpha2.QueryCreditTypesResponse)
    - [QueryParamsRequest](#regen.ecocredit.v1alpha2.QueryParamsRequest)
    - [QueryParamsResponse](#regen.ecocredit.v1alpha2.QueryParamsResponse)
    - [QueryProjectInfoRequest](#regen.ecocredit.v1alpha2.QueryProjectInfoRequest)
    - [QueryProjectInfoResponse](#regen.ecocredit.v1alpha2.QueryProjectInfoResponse)
    - [QueryProjectsRequest](#regen.ecocredit.v1alpha2.QueryProjectsRequest)
    - [QueryProjectsResponse](#regen.ecocredit.v1alpha2.QueryProjectsResponse)
    - [QuerySellOrderRequest](#regen.ecocredit.v1alpha2.QuerySellOrderRequest)
    - [QuerySellOrderResponse](#regen.ecocredit.v1alpha2.QuerySellOrderResponse)
    - [QuerySellOrdersByAddressRequest](#regen.ecocredit.v1alpha2.QuerySellOrdersByAddressRequest)
    - [QuerySellOrdersByAddressResponse](#regen.ecocredit.v1alpha2.QuerySellOrdersByAddressResponse)
    - [QuerySellOrdersByBatchDenomRequest](#regen.ecocredit.v1alpha2.QuerySellOrdersByBatchDenomRequest)
    - [QuerySellOrdersByBatchDenomResponse](#regen.ecocredit.v1alpha2.QuerySellOrdersByBatchDenomResponse)
    - [QuerySellOrdersRequest](#regen.ecocredit.v1alpha2.QuerySellOrdersRequest)
    - [QuerySellOrdersResponse](#regen.ecocredit.v1alpha2.QuerySellOrdersResponse)
    - [QuerySupplyRequest](#regen.ecocredit.v1alpha2.QuerySupplyRequest)
    - [QuerySupplyResponse](#regen.ecocredit.v1alpha2.QuerySupplyResponse)
  
    - [Query](#regen.ecocredit.v1alpha2.Query)
  
- [regen/ecocredit/v1alpha2/tx.proto](#regen/ecocredit/v1alpha2/tx.proto)
    - [MsgAddToBasket](#regen.ecocredit.v1alpha2.MsgAddToBasket)
    - [MsgAddToBasketResponse](#regen.ecocredit.v1alpha2.MsgAddToBasketResponse)
    - [MsgAllowAskDenom](#regen.ecocredit.v1alpha2.MsgAllowAskDenom)
    - [MsgAllowAskDenomResponse](#regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse)
    - [MsgBuy](#regen.ecocredit.v1alpha2.MsgBuy)
    - [MsgBuy.Order](#regen.ecocredit.v1alpha2.MsgBuy.Order)
    - [MsgBuy.Order.Selection](#regen.ecocredit.v1alpha2.MsgBuy.Order.Selection)
    - [MsgBuyResponse](#regen.ecocredit.v1alpha2.MsgBuyResponse)
    - [MsgCancel](#regen.ecocredit.v1alpha2.MsgCancel)
    - [MsgCancel.CancelCredits](#regen.ecocredit.v1alpha2.MsgCancel.CancelCredits)
    - [MsgCancelResponse](#regen.ecocredit.v1alpha2.MsgCancelResponse)
    - [MsgCreateBasket](#regen.ecocredit.v1alpha2.MsgCreateBasket)
    - [MsgCreateBasketResponse](#regen.ecocredit.v1alpha2.MsgCreateBasketResponse)
    - [MsgCreateBatch](#regen.ecocredit.v1alpha2.MsgCreateBatch)
    - [MsgCreateBatch.BatchIssuance](#regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance)
    - [MsgCreateBatchResponse](#regen.ecocredit.v1alpha2.MsgCreateBatchResponse)
    - [MsgCreateClass](#regen.ecocredit.v1alpha2.MsgCreateClass)
    - [MsgCreateClassResponse](#regen.ecocredit.v1alpha2.MsgCreateClassResponse)
    - [MsgCreateProject](#regen.ecocredit.v1alpha2.MsgCreateProject)
    - [MsgCreateProjectResponse](#regen.ecocredit.v1alpha2.MsgCreateProjectResponse)
    - [MsgPickFromBasket](#regen.ecocredit.v1alpha2.MsgPickFromBasket)
    - [MsgPickFromBasketResponse](#regen.ecocredit.v1alpha2.MsgPickFromBasketResponse)
    - [MsgRetire](#regen.ecocredit.v1alpha2.MsgRetire)
    - [MsgRetire.RetireCredits](#regen.ecocredit.v1alpha2.MsgRetire.RetireCredits)
    - [MsgRetireResponse](#regen.ecocredit.v1alpha2.MsgRetireResponse)
    - [MsgSell](#regen.ecocredit.v1alpha2.MsgSell)
    - [MsgSell.Order](#regen.ecocredit.v1alpha2.MsgSell.Order)
    - [MsgSellResponse](#regen.ecocredit.v1alpha2.MsgSellResponse)
    - [MsgSend](#regen.ecocredit.v1alpha2.MsgSend)
    - [MsgSend.SendCredits](#regen.ecocredit.v1alpha2.MsgSend.SendCredits)
    - [MsgSendResponse](#regen.ecocredit.v1alpha2.MsgSendResponse)
    - [MsgTakeFromBasket](#regen.ecocredit.v1alpha2.MsgTakeFromBasket)
    - [MsgTakeFromBasketResponse](#regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse)
    - [MsgUpdateClassAdmin](#regen.ecocredit.v1alpha2.MsgUpdateClassAdmin)
    - [MsgUpdateClassAdminResponse](#regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse)
    - [MsgUpdateClassIssuers](#regen.ecocredit.v1alpha2.MsgUpdateClassIssuers)
    - [MsgUpdateClassIssuersResponse](#regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse)
    - [MsgUpdateClassMetadata](#regen.ecocredit.v1alpha2.MsgUpdateClassMetadata)
    - [MsgUpdateClassMetadataResponse](#regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse)
    - [MsgUpdateSellOrders](#regen.ecocredit.v1alpha2.MsgUpdateSellOrders)
    - [MsgUpdateSellOrders.Update](#regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update)
    - [MsgUpdateSellOrdersResponse](#regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse)
  
    - [Msg](#regen.ecocredit.v1alpha2.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/ecocredit/v1alpha2/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1alpha2/events.proto



<a name="regen.ecocredit.v1alpha2.EventAllowAskDenom"></a>

### EventAllowAskDenom
EventAllowAskDenom is an event emitted when an ask denom is added.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [string](#string) |  | denom is the denom to allow (ex. ibc/GLKHDSG423SGS) |
| display_denom | [string](#string) |  | display_denom is the denom to display to the user and is informational |
| exponent | [uint32](#uint32) |  | exponent is the exponent that relates the denom to the display_denom and is informational |






<a name="regen.ecocredit.v1alpha2.EventBuyOrderCreated"></a>

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






<a name="regen.ecocredit.v1alpha2.EventBuyOrderFilled"></a>

### EventBuyOrderFilled
EventBuyOrderFilled is an event emitted when a buy order is filled.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order_id | [uint64](#uint64) |  | buy_order_id is the unique ID of the buy order. |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the unique ID of the sell order. |
| batch_denom | [string](#string) |  | batch_denom is the credit batch ID of the purchased credits. |
| quantity | [string](#string) |  | quantity is the quantity of the purchased credits. |
| total_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | total_price is the total price for the purchased credits. |






<a name="regen.ecocredit.v1alpha2.EventCancel"></a>

### EventCancel
EventCancel is an event emitted when credits are cancelled. When credits are
cancelled from multiple batches in the same transaction, a separate event is
emitted for each batch_denom. This allows for easier indexing.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| canceller | [string](#string) |  | canceller is the account which has cancelled the credits, which should be the holder of the credits. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch. |
| amount | [string](#string) |  | amount is the decimal number of credits that have been cancelled. |






<a name="regen.ecocredit.v1alpha2.EventCreateBatch"></a>

### EventCreateBatch
EventCreateBatch is an event emitted when a credit batch is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch. |
| issuer | [string](#string) |  | issuer is the account address of the issuer of the credit batch. |
| total_amount | [string](#string) |  | total_amount is the total number of credits in the credit batch. |
| start_date | [string](#string) |  | start_date is the beginning of the period during which this credit batch was quantified and verified. |
| end_date | [string](#string) |  | end_date is the end of the period during which this credit batch was quantified and verified. |
| project_location | [string](#string) |  | project_location is the location of the project. Full documentation can be found in MsgCreateProject.project_location. |
| project_id | [string](#string) |  | project_id is the unique ID of the project this batch belongs to. |






<a name="regen.ecocredit.v1alpha2.EventCreateClass"></a>

### EventCreateClass
EventCreateClass is an event emitted when a credit class is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class. |
| admin | [string](#string) |  | admin is the admin of the credit class. |






<a name="regen.ecocredit.v1alpha2.EventCreateProject"></a>

### EventCreateProject
EventCreateProject is an event emitted when a project is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [string](#string) |  | project_id is the unique ID of the project. |
| class_id | [string](#string) |  | class_id is the unique ID of credit class for this project. |
| issuer | [string](#string) |  | issuer is the issuer of the credit batches for this project. |
| project_location | [string](#string) |  | project_location is the location of the project. Full documentation can be found in MsgCreateProject.project_location. |






<a name="regen.ecocredit.v1alpha2.EventReceive"></a>

### EventReceive
EventReceive is an event emitted when credits are received either upon
creation of a new batch or upon transfer. Each batch_denom created or
transferred will result in a separate EventReceive for easy indexing.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the sender of the credits in the case that this event is the result of a transfer. It will not be set when credits are received at initial issuance. |
| recipient | [string](#string) |  | recipient is the recipient of the credits |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch. |
| tradable_amount | [string](#string) |  | tradable_amount is the decimal number of tradable credits received. |
| retired_amount | [string](#string) |  | retired_amount is the decimal number of retired credits received. |






<a name="regen.ecocredit.v1alpha2.EventRetire"></a>

### EventRetire
EventRetire is an event emitted when credits are retired. When credits are
retired from multiple batches in the same transaction, a separate event is
emitted for each batch_denom. This allows for easier indexing.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| retirer | [string](#string) |  | retirer is the account which has done the "retiring". This will be the account receiving credits in the case that credits were retired upon issuance using Msg/CreateBatch or retired upon transfer using Msg/Send. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch. |
| amount | [string](#string) |  | amount is the decimal number of credits that have been retired. |
| location | [string](#string) |  | location is the location of the beneficiary or buyer of the retired credits. It is a string of the form <country-code>[-<sub-national-code>[ <postal-code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters. |






<a name="regen.ecocredit.v1alpha2.EventSell"></a>

### EventSell
EventSell is an event emitted when a sell order is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [uint64](#uint64) |  | order_id is the unique ID of sell order. |
| batch_denom | [string](#string) |  | batch_denom is the credit batch being sold. |
| quantity | [string](#string) |  | quantity is the quantity of credits being sold. |
| ask_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | ask_price is the price the seller is asking for each unit of the batch_denom. Each credit unit of the batch will be sold for at least the ask_price or more. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire disables auto-retirement of credits which allows a buyer to disable auto-retirement in their buy order enabling them to resell the credits to another buyer. |






<a name="regen.ecocredit.v1alpha2.EventUpdateSellOrder"></a>

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





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/v1alpha2/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1alpha2/types.proto



<a name="regen.ecocredit.v1alpha2.AskDenom"></a>

### AskDenom
AskDenom represents the information for an ask denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [string](#string) |  | denom is the denom to allow (ex. ibc/GLKHDSG423SGS) |
| display_denom | [string](#string) |  | display_denom is the denom to display to the user and is informational |
| exponent | [uint32](#uint32) |  | exponent is the exponent that relates the denom to the display_denom and is informational |






<a name="regen.ecocredit.v1alpha2.BasketCredit"></a>

### BasketCredit
BasketCredit represents the information for a credit batch inside a basket.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| tradable_amount | [string](#string) |  | tradable_amount is the number of credits in this transfer that can be traded by the recipient. Decimal values are acceptable within the precision returned by Query/Precision. |






<a name="regen.ecocredit.v1alpha2.BasketCriteria"></a>

### BasketCriteria
BasketCriteria defines a criteria by which credits can be added to a basket.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filter | [Filter](#regen.ecocredit.v1alpha2.Filter) |  | filter defines condition(s) that credits should satisfy in order to be added to the basket. |
| multiplier | [string](#string) |  | multiplier is an integer number which is applied to credit units when converting to basket units. For example if the multiplier is 2000, then 1.1 credits will result in 2200 basket tokens. If there are any fractional amounts left over in this calculation when adding credits to a basket, those fractional amounts will not get added to the basket. |






<a name="regen.ecocredit.v1alpha2.BatchInfo"></a>

### BatchInfo
BatchInfo represents the high-level on-chain information for a credit batch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [string](#string) |  | project_id is the unique ID of the project this batch belongs to. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch. |
| total_amount | [string](#string) |  | total_amount is the total number of active credits in the credit batch. Some of the issued credits may be cancelled and will be removed from total_amount and tracked in amount_cancelled. total_amount and amount_cancelled will always sum to the original amount of credits that were issued. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata attached to the credit batch. |
| amount_cancelled | [string](#string) |  | amount_cancelled is the number of credits in the batch that have been cancelled, effectively undoing there issuance. The sum of total_amount and amount_cancelled will always sum to the original amount of credits that were issued. |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_date is the beginning of the period during which this credit batch was quantified and verified. |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | end_date is the end of the period during which this credit batch was quantified and verified. |






<a name="regen.ecocredit.v1alpha2.BuyOrder"></a>

### BuyOrder
BuyOrder represents the information for a buy order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order_id | [uint64](#uint64) |  | buy_order_id is the unique ID of buy order. |
| buyer | [string](#string) |  | buyer is the address that created the buy order |
| selection | [BuyOrder.Selection](#regen.ecocredit.v1alpha2.BuyOrder.Selection) |  | selection is the buy order selection. |
| quantity | [string](#string) |  | quantity is the quantity of credits to buy. If the quantity of credits available is less than this amount the order will be partially filled unless disable_partial_fill is true. |
| bid_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | bid price is the bid price for this buy order. A credit unit will be settled at a purchase price that is no more than the bid price. The buy order will fail if the buyer does not have enough funds available to complete the purchase. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire allows auto-retirement to be disabled. If it is set to true the credits will not auto-retire and can be resold assuming that the corresponding sell order has auto-retirement disabled. If the sell order hasn't disabled auto-retirement and the buy order tries to disable it, that buy order will fail. |
| disable_partial_fill | [bool](#bool) |  | disable_partial_fill disables the default behavior of partially filling buy orders if the requested quantity is not available. |






<a name="regen.ecocredit.v1alpha2.BuyOrder.Selection"></a>

### BuyOrder.Selection
Selection defines a buy order selection.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the sell order ID against which the buyer is trying to buy. When sell_order_id is set, this is known as a direct buy order because it is placed directly against a specific sell order. |






<a name="regen.ecocredit.v1alpha2.ClassInfo"></a>

### ClassInfo
ClassInfo represents the high-level on-chain information for a credit class.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class. |
| admin | [string](#string) |  | admin is the admin of the credit class. |
| issuers | [string](#string) | repeated | issuers are the approved issuers of the credit class. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata to attached to the credit class. |
| credit_type | [CreditType](#regen.ecocredit.v1alpha2.CreditType) |  | credit_type describes the type of credit (e.g. carbon, biodiversity), as well as unit and precision. |
| num_batches | [uint64](#uint64) |  | The number of batches issued in this credit class. |






<a name="regen.ecocredit.v1alpha2.CreditType"></a>

### CreditType
CreditType defines the measurement unit/precision of a certain credit type
(e.g. carbon, biodiversity...)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | the type of credit (e.g. carbon, biodiversity, etc) |
| abbreviation | [string](#string) |  | abbreviation is a 1-3 character uppercase abbreviation of the CreditType name, used in batch denominations within the CreditType. It must be unique. |
| unit | [string](#string) |  | the measurement unit (e.g. kg, ton, etc) |
| precision | [uint32](#uint32) |  | the decimal precision |






<a name="regen.ecocredit.v1alpha2.CreditTypeSeq"></a>

### CreditTypeSeq
CreditTypeSeq associates a sequence number with a credit type abbreviation.
This represents the number of credit classes created with that credit type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| abbreviation | [string](#string) |  | The credit type abbreviation |
| seq_number | [uint64](#uint64) |  | The sequence number of classes of the credit type |






<a name="regen.ecocredit.v1alpha2.Filter"></a>

### Filter
Filter defines condition(s) that credits should satisfy in order to be added
to the basket. It can handled nested conditions linked with and/or operators.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| and | [Filter.And](#regen.ecocredit.v1alpha2.Filter.And) |  | and specifies a list of filters where all conditions should be satisfied. |
| or | [Filter.Or](#regen.ecocredit.v1alpha2.Filter.Or) |  | or specifies a list of filters where at least one of the conditions should be satisfied. |
| credit_type_name | [string](#string) |  | credit_type_name filters against credits from this credit type name. |
| class_id | [string](#string) |  | class_id filters against credits from this credit class id. |
| project_id | [string](#string) |  | project_id filters against credits from this project. |
| batch_denom | [string](#string) |  | batch_denom filters against credits from this batch. |
| class_admin | [string](#string) |  | class_admin filters against credits issued by this class admin. |
| issuer | [string](#string) |  | issuer filters against credits issued by this issuer address. |
| owner | [string](#string) |  | owner filters against credits currently owned by this address. |
| project_location | [string](#string) |  | project_location can be specified in three levels of granularity: country, sub-national-code, or postal code. If just country is given, for instance "US" then any credits in the "US" will be matched even their project location is more specific, ex. "US-NY 12345". If a country, sub-national-code and postal code are all provided then only projects in that postal code will match. |
| date_range | [Filter.DateRange](#regen.ecocredit.v1alpha2.Filter.DateRange) |  | date_range filters against credit batch start_date and/or end_date. |
| tag | [string](#string) |  | tag specifies a curation tag to match against. |






<a name="regen.ecocredit.v1alpha2.Filter.And"></a>

### Filter.And
And specifies an "and" condition between the list of filters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [Filter](#regen.ecocredit.v1alpha2.Filter) | repeated | filters is a list of filters where all conditions should be satisfied. |






<a name="regen.ecocredit.v1alpha2.Filter.DateRange"></a>

### Filter.DateRange
DateRange defines a period for credit batches in a basket.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_date is the beginning of the period during which this credit batch was quantified and verified. If it is empty then there is no start date limit. |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | end_date is the end of the period during which this credit batch was quantified and verified. If it is empty then there is no end date limit. |






<a name="regen.ecocredit.v1alpha2.Filter.Or"></a>

### Filter.Or
And specifies an "or" condition between the list of filters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [Filter](#regen.ecocredit.v1alpha2.Filter) | repeated | filters is a list of filters where at least one of the conditions should be satisfied. |






<a name="regen.ecocredit.v1alpha2.Params"></a>

### Params
Params defines the updatable global parameters of the ecocredit module for
use with the x/params module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credit_class_fee | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | credit_class_fee is the fixed fee charged on creation of a new credit class |
| allowed_class_creators | [string](#string) | repeated | allowed_class_creators is an allowlist defining the addresses with the required permissions to create credit classes |
| allowlist_enabled | [bool](#bool) |  | allowlist_enabled is a param that enables/disables the allowlist for credit creation |
| credit_types | [CreditType](#regen.ecocredit.v1alpha2.CreditType) | repeated | credit_types is a list of definitions for credit types |






<a name="regen.ecocredit.v1alpha2.ProjectInfo"></a>

### ProjectInfo
ProjectInfo represents the high-level on-chain information for a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [string](#string) |  | project_id is the unique ID of the project. |
| class_id | [string](#string) |  | class_id is the unique ID of credit class for this project. |
| issuer | [string](#string) |  | issuer is the issuer of the credit batches for this project. |
| project_location | [string](#string) |  | project_location is the location of the project. Full documentation can be found in MsgCreateProject.project_location. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata attached to the project. |






<a name="regen.ecocredit.v1alpha2.SellOrder"></a>

### SellOrder
SellOrder represents the information for a sell order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [uint64](#uint64) |  | order_id is the unique ID of sell order. |
| owner | [string](#string) |  | owner is the address of the owner of the credits being sold. |
| batch_denom | [string](#string) |  | batch_denom is the credit batch being sold. |
| quantity | [string](#string) |  | quantity is the quantity of credits being sold. |
| ask_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | ask_price is the price the seller is asking for each unit of the batch_denom. Each credit unit of the batch will be sold for at least the ask_price or more. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire disables auto-retirement of credits which allows a buyer to disable auto-retirement in their buy order enabling them to resell the credits to another buyer. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/v1alpha2/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1alpha2/genesis.proto



<a name="regen.ecocredit.v1alpha2.Balance"></a>

### Balance
Balance represents tradable or retired units of a credit batch with an
account address, batch_denom, and balance.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address is the account address of the account holding credits. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| tradable_balance | [string](#string) |  | tradable_balance is the tradable balance of the credit batch. |
| retired_balance | [string](#string) |  | retired_balance is the retired balance of the credit batch. |






<a name="regen.ecocredit.v1alpha2.GenesisState"></a>

### GenesisState
GenesisState defines ecocredit module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#regen.ecocredit.v1alpha2.Params) |  | Params contains the updateable global parameters for use with the x/params module |
| class_info | [ClassInfo](#regen.ecocredit.v1alpha2.ClassInfo) | repeated | class_info is the list of credit class info. |
| batch_info | [BatchInfo](#regen.ecocredit.v1alpha2.BatchInfo) | repeated | batch_info is the list of credit batch info. |
| sequences | [CreditTypeSeq](#regen.ecocredit.v1alpha2.CreditTypeSeq) | repeated | sequences is the list of credit type sequence. |
| balances | [Balance](#regen.ecocredit.v1alpha2.Balance) | repeated | balances is the list of credit batch tradable/retired units. |
| supplies | [Supply](#regen.ecocredit.v1alpha2.Supply) | repeated | supplies is the list of credit batch tradable/retired supply. |
| project_info | [ProjectInfo](#regen.ecocredit.v1alpha2.ProjectInfo) | repeated | project_info is the list of projects. |
| project_seq_num | [uint64](#uint64) |  | project_seq_num is the project table orm.Sequence, it is used to generate the next project id. |






<a name="regen.ecocredit.v1alpha2.Supply"></a>

### Supply
Supply represents a tradable or retired supply of a credit batch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| tradable_supply | [string](#string) |  | tradable_supply is the tradable supply of the credit batch. |
| retired_supply | [string](#string) |  | retired_supply is the retired supply of the credit batch. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/v1alpha2/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1alpha2/query.proto



<a name="regen.ecocredit.v1alpha2.Basket"></a>

### Basket
Basket defines a credit basket.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| curator | [string](#string) |  | curator is the address of the basket curator who is able to change certain basket settings. |
| name | [string](#string) |  | name will be used to create a bank denom for this basket token of the form ecocredit:{curator}:{name}. |
| display_name | [string](#string) |  | display_name will be used to create a bank Metadata display name for this basket token of the form ecocredit:{curator}:{display_name}. |
| exponent | [uint32](#uint32) |  | exponent is the exponent that will be used for denom metadata. An exponent of 6 will mean that 10^6 units of a basket token should be displayed as one unit in user interfaces. |
| basket_criteria | [BasketCriteria](#regen.ecocredit.v1alpha2.BasketCriteria) | repeated | basket_criteria is the criteria by which credits can be added to the basket. Basket criteria will be applied in order and the first criteria which applies to a credit will determine its multiplier in the basket. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire allows auto-retirement to be disabled. The credits will be auto-retired if disable_auto_retire is false unless the credits were previously put into the basket by the address picking them from the basket, in which case they will remain tradable. |
| allow_picking | [bool](#bool) |  | allow_picking specifies whether an address which didn't deposit the credits in the basket can pick those credits or not. |






<a name="regen.ecocredit.v1alpha2.QueryAllowedAskDenomsRequest"></a>

### QueryAllowedAskDenomsRequest
QueryAllowedAskDenomsRequest is the Query/AllowedAskDenoms request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1alpha2.QueryAllowedAskDenomsResponse"></a>

### QueryAllowedAskDenomsResponse
QueryAllowedAskDenomsResponse is the Query/AllowedAskDenoms response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ask_denoms | [AskDenom](#regen.ecocredit.v1alpha2.AskDenom) | repeated | ask_denoms is a list of coin denoms allowed to use in the ask price of sell orders. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an optional pagination for the response. |






<a name="regen.ecocredit.v1alpha2.QueryBalanceRequest"></a>

### QueryBalanceRequest
QueryBalanceRequest is the Query/Balance request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  | account is the address of the account whose balance is being queried. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch balance to query. |






<a name="regen.ecocredit.v1alpha2.QueryBalanceResponse"></a>

### QueryBalanceResponse
QueryBalanceResponse is the Query/Balance response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tradable_amount | [string](#string) |  | tradable_amount is the decimal number of tradable credits. |
| retired_amount | [string](#string) |  | retired_amount is the decimal number of retired credits. |






<a name="regen.ecocredit.v1alpha2.QueryBasketCreditsRequest"></a>

### QueryBasketCreditsRequest
QueryBasketCreditsRequest is the Query/BasketCredits request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basket_denom | [string](#string) |  | basket_denom is the basket denom to query credits for. |






<a name="regen.ecocredit.v1alpha2.QueryBasketCreditsResponse"></a>

### QueryBasketCreditsResponse
QueryBasketCreditsResponse is the Query/BasketCredits response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credits | [BasketCredit](#regen.ecocredit.v1alpha2.BasketCredit) | repeated | credits are the credits inside the basket. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.v1alpha2.QueryBasketRequest"></a>

### QueryBasketRequest
QueryBasketRequest is the Query/Basket request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basket_denom | [string](#string) |  | basket_denom represents the denom of the basket to query. |






<a name="regen.ecocredit.v1alpha2.QueryBasketResponse"></a>

### QueryBasketResponse
QueryBasketResponse is the Query/Basket response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basket | [Basket](#regen.ecocredit.v1alpha2.Basket) |  | basket is the queried basket. |






<a name="regen.ecocredit.v1alpha2.QueryBasketsRequest"></a>

### QueryBasketsRequest
QueryBasketsRequest is the Query/Baskets request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1alpha2.QueryBasketsResponse"></a>

### QueryBasketsResponse
QueryBasketsResponse is the Query/Baskets response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baskets | [Basket](#regen.ecocredit.v1alpha2.Basket) | repeated | baskets are the fetched baskets. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.v1alpha2.QueryBatchInfoRequest"></a>

### QueryBatchInfoRequest
QueryBatchInfoRequest is the Query/BatchInfo request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch to query. |






<a name="regen.ecocredit.v1alpha2.QueryBatchInfoResponse"></a>

### QueryBatchInfoResponse
QueryBatchInfoResponse is the Query/BatchInfo response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [BatchInfo](#regen.ecocredit.v1alpha2.BatchInfo) |  | info is the BatchInfo for the credit batch. |






<a name="regen.ecocredit.v1alpha2.QueryBatchesRequest"></a>

### QueryBatchesRequest
QueryBatchesRequest is the Query/Batches request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [string](#string) |  | project_id is the unique ID of the project to query. |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1alpha2.QueryBatchesResponse"></a>

### QueryBatchesResponse
QueryBatchesResponse is the Query/Batches response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batches | [BatchInfo](#regen.ecocredit.v1alpha2.BatchInfo) | repeated | batches are the fetched credit batches within the project. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.v1alpha2.QueryBuyOrderRequest"></a>

### QueryBuyOrderRequest
QueryBuyOrderRequest is the Query/BuyOrder request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order_id | [uint64](#uint64) |  | buy_order_id is the id of the buy order. |






<a name="regen.ecocredit.v1alpha2.QueryBuyOrderResponse"></a>

### QueryBuyOrderResponse
QueryBuyOrderResponse is the Query/BuyOrder response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order | [BuyOrder](#regen.ecocredit.v1alpha2.BuyOrder) |  | buy_order contains all information related to a buy order. |






<a name="regen.ecocredit.v1alpha2.QueryBuyOrdersByAddressRequest"></a>

### QueryBuyOrdersByAddressRequest
QueryBuyOrdersByAddressRequest is the Query/BuyOrdersByAddress request type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address of the buy order creator |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1alpha2.QueryBuyOrdersByAddressResponse"></a>

### QueryBuyOrdersByAddressResponse
QueryBuyOrdersByAddressResponse is the Query/BuyOrdersByAddress response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_orders | [BuyOrder](#regen.ecocredit.v1alpha2.BuyOrder) | repeated | buy_orders is a list of buy orders. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an optional pagination for the response. |






<a name="regen.ecocredit.v1alpha2.QueryBuyOrdersRequest"></a>

### QueryBuyOrdersRequest
QueryBuyOrdersRequest is the Query/BuyOrders request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1alpha2.QueryBuyOrdersResponse"></a>

### QueryBuyOrdersResponse
QueryBuyOrdersResponse is the Query/BuyOrders response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_orders | [BuyOrder](#regen.ecocredit.v1alpha2.BuyOrder) | repeated | buy_orders is a list of buy orders. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an optional pagination for the response. |






<a name="regen.ecocredit.v1alpha2.QueryClassInfoRequest"></a>

### QueryClassInfoRequest
QueryClassInfoRequest is the Query/ClassInfo request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class to query. |






<a name="regen.ecocredit.v1alpha2.QueryClassInfoResponse"></a>

### QueryClassInfoResponse
QueryClassInfoResponse is the Query/ClassInfo request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [ClassInfo](#regen.ecocredit.v1alpha2.ClassInfo) |  | info is the ClassInfo for the credit class. |






<a name="regen.ecocredit.v1alpha2.QueryClassesRequest"></a>

### QueryClassesRequest
QueryClassesRequest is the Query/Classes request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1alpha2.QueryClassesResponse"></a>

### QueryClassesResponse
QueryClassesResponse is the Query/Classes response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| classes | [ClassInfo](#regen.ecocredit.v1alpha2.ClassInfo) | repeated | classes are the fetched credit classes. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.v1alpha2.QueryCreditTypesRequest"></a>

### QueryCreditTypesRequest
QueryCreditTypesRequest is the Query/Credit_Types request type






<a name="regen.ecocredit.v1alpha2.QueryCreditTypesResponse"></a>

### QueryCreditTypesResponse
QueryCreditTypesRequest is the Query/Credit_Types response type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credit_types | [CreditType](#regen.ecocredit.v1alpha2.CreditType) | repeated | list of credit types |






<a name="regen.ecocredit.v1alpha2.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the Query/Params request type.






<a name="regen.ecocredit.v1alpha2.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the Query/Params response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#regen.ecocredit.v1alpha2.Params) |  | params defines the parameters of the ecocredit module. |






<a name="regen.ecocredit.v1alpha2.QueryProjectInfoRequest"></a>

### QueryProjectInfoRequest
QueryProjectInfoRequest is the Query/Project request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [string](#string) |  | project_id is the unique ID of the project to query. |






<a name="regen.ecocredit.v1alpha2.QueryProjectInfoResponse"></a>

### QueryProjectInfoResponse
QueryProjectInfoResponse is the Query/Project response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [ProjectInfo](#regen.ecocredit.v1alpha2.ProjectInfo) |  | info is the ProjectInfo for the project. |






<a name="regen.ecocredit.v1alpha2.QueryProjectsRequest"></a>

### QueryProjectsRequest
QueryProjectsRequest is the Query/Projects request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class to query. |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1alpha2.QueryProjectsResponse"></a>

### QueryProjectsResponse
QueryProjectsResponse is the Query/Projects response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| projects | [ProjectInfo](#regen.ecocredit.v1alpha2.ProjectInfo) | repeated | projects are the fetched projects. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.v1alpha2.QuerySellOrderRequest"></a>

### QuerySellOrderRequest
QuerySellOrderRequest is the Query/SellOrder request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the id of the requested sell order. |






<a name="regen.ecocredit.v1alpha2.QuerySellOrderResponse"></a>

### QuerySellOrderResponse
QuerySellOrderResponse is the Query/SellOrder response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_order | [SellOrder](#regen.ecocredit.v1alpha2.SellOrder) |  | sell_order contains all information related to a sell order. |






<a name="regen.ecocredit.v1alpha2.QuerySellOrdersByAddressRequest"></a>

### QuerySellOrdersByAddressRequest
QuerySellOrdersByAddressRequest is the Query/SellOrdersByAddress request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address is the creator of the sell order |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1alpha2.QuerySellOrdersByAddressResponse"></a>

### QuerySellOrdersByAddressResponse
QuerySellOrdersByAddressResponse is the Query/SellOrdersByAddressResponse response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_orders | [SellOrder](#regen.ecocredit.v1alpha2.SellOrder) | repeated | sell_orders is a list of sell orders. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an optional pagination for the response. |






<a name="regen.ecocredit.v1alpha2.QuerySellOrdersByBatchDenomRequest"></a>

### QuerySellOrdersByBatchDenomRequest
QuerySellOrdersByDenomRequest is the Query/SellOrdersByDenom request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is an ecocredit denom |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1alpha2.QuerySellOrdersByBatchDenomResponse"></a>

### QuerySellOrdersByBatchDenomResponse
QuerySellOrdersByDenomResponse is the Query/SellOrdersByDenom response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_orders | [SellOrder](#regen.ecocredit.v1alpha2.SellOrder) | repeated | sell_orders is a list of sell orders. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an optional pagination for the response. |






<a name="regen.ecocredit.v1alpha2.QuerySellOrdersRequest"></a>

### QuerySellOrdersRequest
QuerySellOrdersRequest is the Query/SellOrders request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1alpha2.QuerySellOrdersResponse"></a>

### QuerySellOrdersResponse
QuerySellOrdersResponse is the Query/SellOrders response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_orders | [SellOrder](#regen.ecocredit.v1alpha2.SellOrder) | repeated | sell_orders is a list of sell orders. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.v1alpha2.QuerySupplyRequest"></a>

### QuerySupplyRequest
QuerySupplyRequest is the Query/Supply request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch to query. |






<a name="regen.ecocredit.v1alpha2.QuerySupplyResponse"></a>

### QuerySupplyResponse
QuerySupplyResponse is the Query/Supply response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tradable_supply | [string](#string) |  | tradable_supply is the decimal number of tradable credits in the batch supply. |
| retired_supply | [string](#string) |  | retired_supply is the decimal number of retired credits in the batch supply. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.ecocredit.v1alpha2.Query"></a>

### Query
Msg is the regen.ecocredit.v1alpha2 Query service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Classes | [QueryClassesRequest](#regen.ecocredit.v1alpha2.QueryClassesRequest) | [QueryClassesResponse](#regen.ecocredit.v1alpha2.QueryClassesResponse) | Classes queries for all credit classes with pagination. |
| ClassInfo | [QueryClassInfoRequest](#regen.ecocredit.v1alpha2.QueryClassInfoRequest) | [QueryClassInfoResponse](#regen.ecocredit.v1alpha2.QueryClassInfoResponse) | ClassInfo queries for information on a credit class. |
| Projects | [QueryProjectsRequest](#regen.ecocredit.v1alpha2.QueryProjectsRequest) | [QueryProjectsResponse](#regen.ecocredit.v1alpha2.QueryProjectsResponse) | Projects queries for all projects within a class with pagination. |
| ProjectInfo | [QueryProjectInfoRequest](#regen.ecocredit.v1alpha2.QueryProjectInfoRequest) | [QueryProjectInfoResponse](#regen.ecocredit.v1alpha2.QueryProjectInfoResponse) | ClassInfo queries for information on a project. |
| Batches | [QueryBatchesRequest](#regen.ecocredit.v1alpha2.QueryBatchesRequest) | [QueryBatchesResponse](#regen.ecocredit.v1alpha2.QueryBatchesResponse) | Batches queries for all batches in the given project with pagination. |
| BatchInfo | [QueryBatchInfoRequest](#regen.ecocredit.v1alpha2.QueryBatchInfoRequest) | [QueryBatchInfoResponse](#regen.ecocredit.v1alpha2.QueryBatchInfoResponse) | BatchInfo queries for information on a credit batch. |
| Balance | [QueryBalanceRequest](#regen.ecocredit.v1alpha2.QueryBalanceRequest) | [QueryBalanceResponse](#regen.ecocredit.v1alpha2.QueryBalanceResponse) | Balance queries the balance (both tradable and retired) of a given credit batch for a given account. |
| Supply | [QuerySupplyRequest](#regen.ecocredit.v1alpha2.QuerySupplyRequest) | [QuerySupplyResponse](#regen.ecocredit.v1alpha2.QuerySupplyResponse) | Supply queries the tradable and retired supply of a credit batch. |
| CreditTypes | [QueryCreditTypesRequest](#regen.ecocredit.v1alpha2.QueryCreditTypesRequest) | [QueryCreditTypesResponse](#regen.ecocredit.v1alpha2.QueryCreditTypesResponse) | CreditTypes returns the list of allowed types that credit classes can have. See Types/CreditType for more details. |
| Params | [QueryParamsRequest](#regen.ecocredit.v1alpha2.QueryParamsRequest) | [QueryParamsResponse](#regen.ecocredit.v1alpha2.QueryParamsResponse) | Params queries the ecocredit module parameters. |
| SellOrder | [QuerySellOrderRequest](#regen.ecocredit.v1alpha2.QuerySellOrderRequest) | [QuerySellOrderResponse](#regen.ecocredit.v1alpha2.QuerySellOrderResponse) | SellOrder queries a sell order by its ID |
| SellOrders | [QuerySellOrdersRequest](#regen.ecocredit.v1alpha2.QuerySellOrdersRequest) | [QuerySellOrdersResponse](#regen.ecocredit.v1alpha2.QuerySellOrdersResponse) | SellOrders queries a paginated list of all sell orders |
| SellOrdersByBatchDenom | [QuerySellOrdersByBatchDenomRequest](#regen.ecocredit.v1alpha2.QuerySellOrdersByBatchDenomRequest) | [QuerySellOrdersByBatchDenomResponse](#regen.ecocredit.v1alpha2.QuerySellOrdersByBatchDenomResponse) | SellOrdersByDenom queries a paginated list of all sell orders of a specific ecocredit denom |
| SellOrdersByAddress | [QuerySellOrdersByAddressRequest](#regen.ecocredit.v1alpha2.QuerySellOrdersByAddressRequest) | [QuerySellOrdersByAddressResponse](#regen.ecocredit.v1alpha2.QuerySellOrdersByAddressResponse) | SellOrdersByAddress queries a paginated list of all sell orders from a specific address |
| BuyOrder | [QueryBuyOrderRequest](#regen.ecocredit.v1alpha2.QueryBuyOrderRequest) | [QueryBuyOrderResponse](#regen.ecocredit.v1alpha2.QueryBuyOrderResponse) | BuyOrder queries a buy order by its id |
| BuyOrders | [QueryBuyOrdersRequest](#regen.ecocredit.v1alpha2.QueryBuyOrdersRequest) | [QueryBuyOrdersResponse](#regen.ecocredit.v1alpha2.QueryBuyOrdersResponse) | BuyOrders queries a paginated list of all buy orders |
| BuyOrdersByAddress | [QueryBuyOrdersByAddressRequest](#regen.ecocredit.v1alpha2.QueryBuyOrdersByAddressRequest) | [QueryBuyOrdersByAddressResponse](#regen.ecocredit.v1alpha2.QueryBuyOrdersByAddressResponse) | BuyOrdersByAddress queries a paginated list of buy orders by creator address |
| AllowedAskDenoms | [QueryAllowedAskDenomsRequest](#regen.ecocredit.v1alpha2.QueryAllowedAskDenomsRequest) | [QueryAllowedAskDenomsResponse](#regen.ecocredit.v1alpha2.QueryAllowedAskDenomsResponse) | AllowedAskDenoms queries all denoms allowed to be set in the AskPrice of a sell order |
| Basket | [QueryBasketRequest](#regen.ecocredit.v1alpha2.QueryBasketRequest) | [QueryBasketResponse](#regen.ecocredit.v1alpha2.QueryBasketResponse) | Basket queries one basket by denom. |
| Baskets | [QueryBasketsRequest](#regen.ecocredit.v1alpha2.QueryBasketsRequest) | [QueryBasketsResponse](#regen.ecocredit.v1alpha2.QueryBasketsResponse) | Baskets lists all baskets in the ecocredit module. |
| BasketCredits | [QueryBasketCreditsRequest](#regen.ecocredit.v1alpha2.QueryBasketCreditsRequest) | [QueryBasketCreditsResponse](#regen.ecocredit.v1alpha2.QueryBasketCreditsResponse) | BasketCredits lists all ecocredits inside a given basket. |

 <!-- end services -->



<a name="regen/ecocredit/v1alpha2/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1alpha2/tx.proto



<a name="regen.ecocredit.v1alpha2.MsgAddToBasket"></a>

### MsgAddToBasket
MsgAddToBasket is the Msg/AddToBasket request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the owner of credits being added to the basket. |
| basket_denom | [string](#string) |  | basket_denom is the basket denom to add credits to. |
| credits | [BasketCredit](#regen.ecocredit.v1alpha2.BasketCredit) | repeated | credits are credits to add to the basket. If they do not match the basket's admission criteria the operation will fail. |






<a name="regen.ecocredit.v1alpha2.MsgAddToBasketResponse"></a>

### MsgAddToBasketResponse
MsgAddToBasketResponse is the Msg/AddToBasket response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount_received | [string](#string) |  | amount_received is the amount of basket tokens received. |






<a name="regen.ecocredit.v1alpha2.MsgAllowAskDenom"></a>

### MsgAllowAskDenom
MsgAllowAskDenom is the Msg/AllowAskDenom request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root_address | [string](#string) |  | root_address is the address of the governance account which can authorize ask denoms |
| denom | [string](#string) |  | denom is the denom to allow (ex. ibc/GLKHDSG423SGS) |
| display_denom | [string](#string) |  | display_denom is the denom to display to the user and is informational |
| exponent | [uint32](#uint32) |  | exponent is the exponent that relates the denom to the display_denom and is informational |






<a name="regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse"></a>

### MsgAllowAskDenomResponse
MsgAllowAskDenomResponse is the Msg/AllowAskDenom response type.






<a name="regen.ecocredit.v1alpha2.MsgBuy"></a>

### MsgBuy
MsgBuy is the Msg/Buy request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buyer | [string](#string) |  | buyer is the address of the credit buyer. |
| orders | [MsgBuy.Order](#regen.ecocredit.v1alpha2.MsgBuy.Order) | repeated | orders are the new buy orders. |






<a name="regen.ecocredit.v1alpha2.MsgBuy.Order"></a>

### MsgBuy.Order
Order is a buy order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selection | [MsgBuy.Order.Selection](#regen.ecocredit.v1alpha2.MsgBuy.Order.Selection) |  | selection is the buy order selection. |
| quantity | [string](#string) |  | quantity is the quantity of credits to buy. If the quantity of credits available is less than this amount the order will be partially filled unless disable_partial_fill is true. |
| bid_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | bid price is the bid price for this buy order. A credit unit will be settled at a purchase price that is no more than the bid price. The buy order will fail if the buyer does not have enough funds available to complete the purchase. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire allows auto-retirement to be disabled. If it is set to true the credits will not auto-retire and can be resold assuming that the corresponding sell order has auto-retirement disabled. If the sell order hasn't disabled auto-retirement and the buy order tries to disable it, that buy order will fail. |
| disable_partial_fill | [bool](#bool) |  | disable_partial_fill disables the default behavior of partially filling buy orders if the requested quantity is not available. |






<a name="regen.ecocredit.v1alpha2.MsgBuy.Order.Selection"></a>

### MsgBuy.Order.Selection
Selection defines a buy order selection.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the sell order ID against which the buyer is trying to buy. When sell_order_id is set, this is known as a direct buy order because it is placed directly against a specific sell order. |






<a name="regen.ecocredit.v1alpha2.MsgBuyResponse"></a>

### MsgBuyResponse
MsgBuyResponse is the Msg/Buy response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buy_order_ids | [uint64](#uint64) | repeated | buy_order_ids are the buy order IDs of the newly created buy orders. Buy orders may not settle instantaneously, but rather in batches at specified batch epoch times. |






<a name="regen.ecocredit.v1alpha2.MsgCancel"></a>

### MsgCancel
MsgCancel is the Msg/Cancel request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| holder | [string](#string) |  | holder is the credit holder address. |
| credits | [MsgCancel.CancelCredits](#regen.ecocredit.v1alpha2.MsgCancel.CancelCredits) | repeated | credits are the credits being cancelled. |






<a name="regen.ecocredit.v1alpha2.MsgCancel.CancelCredits"></a>

### MsgCancel.CancelCredits
CancelCredits specifies a batch and the number of credits being cancelled.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| amount | [string](#string) |  | amount is the number of credits being cancelled. Decimal values are acceptable within the precision returned by Query/Precision. |






<a name="regen.ecocredit.v1alpha2.MsgCancelResponse"></a>

### MsgCancelResponse
MsgCancelResponse is the Msg/Cancel response type.






<a name="regen.ecocredit.v1alpha2.MsgCreateBasket"></a>

### MsgCreateBasket
MsgCreateBasket is the Msg/CreateBasket request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| curator | [string](#string) |  | curator is the address of the basket curator who is able to change certain basket settings. |
| name | [string](#string) |  | name will be used to create a bank denom for this basket token of the form ecocredit:{curator}:{name}. |
| display_name | [string](#string) |  | display_name will be used to create a bank Metadata display name for this basket token of the form ecocredit:{curator}:{display_name}. |
| exponent | [uint32](#uint32) |  | exponent is the exponent that will be used for denom metadata. An exponent of 6 will mean that 10^6 units of a basket token should be displayed as one unit in user interfaces. |
| basket_criteria | [BasketCriteria](#regen.ecocredit.v1alpha2.BasketCriteria) | repeated | basket_criteria is the criteria by which credits can be added to the basket. Basket criteria will be applied in order and the first criteria which applies to a credit will determine its multiplier in the basket. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire allows auto-retirement to be disabled. The credits will be auto-retired if disable_auto_retire is false unless the credits were previously put into the basket by the address picking them from the basket, in which case they will remain tradable. |
| allow_picking | [bool](#bool) |  | allow_picking specifies whether an address which didn't deposit the credits in the basket can pick those credits or not. |






<a name="regen.ecocredit.v1alpha2.MsgCreateBasketResponse"></a>

### MsgCreateBasketResponse
MsgCreateBasketResponse is the Msg/CreateBasket response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basket_denom | [string](#string) |  | basket_denom is the unique denomination ID of the newly created basket. |






<a name="regen.ecocredit.v1alpha2.MsgCreateBatch"></a>

### MsgCreateBatch
MsgCreateBatch is the Msg/CreateBatch request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| issuer | [string](#string) |  | issuer is the address of the batch issuer. |
| project_id | [string](#string) |  | project_id is the unique ID of the project this batch belongs to. |
| issuance | [MsgCreateBatch.BatchIssuance](#regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance) | repeated | issuance are the credits issued in the batch. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata attached to the credit batch. |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_date is the beginning of the period during which this credit batch was quantified and verified. |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | end_date is the end of the period during which this credit batch was quantified and verified. |






<a name="regen.ecocredit.v1alpha2.MsgCreateBatch.BatchIssuance"></a>

### MsgCreateBatch.BatchIssuance
BatchIssuance represents the issuance of some credits in a batch to a
single recipient.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recipient | [string](#string) |  | recipient is the account of the recipient. |
| tradable_amount | [string](#string) |  | tradable_amount is the number of credits in this issuance that can be traded by this recipient. Decimal values are acceptable. |
| retired_amount | [string](#string) |  | retired_amount is the number of credits in this issuance that are effectively retired by the issuer on receipt. Decimal values are acceptable. |
| retirement_location | [string](#string) |  | retirement_location is the location of the beneficiary or buyer of the retired credits. This must be provided if retired_amount is positive. It is a string of the form <country-code>[-<sub-national-code>[ <postal-code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters. |






<a name="regen.ecocredit.v1alpha2.MsgCreateBatchResponse"></a>

### MsgCreateBatchResponse
MsgCreateBatchResponse is the Msg/CreateBatch response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique denomination ID of the newly created batch. |






<a name="regen.ecocredit.v1alpha2.MsgCreateClass"></a>

### MsgCreateClass
MsgCreateClass is the Msg/CreateClass request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | admin is the address of the account that created the credit class. |
| issuers | [string](#string) | repeated | issuers are the account addresses of the approved issuers. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata to attached to the credit class. |
| credit_type_name | [string](#string) |  | credit_type_name describes the type of credit (e.g. "carbon", "biodiversity"). |






<a name="regen.ecocredit.v1alpha2.MsgCreateClassResponse"></a>

### MsgCreateClassResponse
MsgCreateClassResponse is the Msg/CreateClass response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of the newly created credit class. |






<a name="regen.ecocredit.v1alpha2.MsgCreateProject"></a>

### MsgCreateProject
MsgCreateProjectResponse is the Msg/CreateProject request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| issuer | [string](#string) |  | issuer is the address of an approved issuer for the credit class through which batches will be issued. It is not required, however, that this same issuer issue all batches for a project. |
| class_id | [string](#string) |  | class_id is the unique ID of the class within which the project is created. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata attached to the project. |
| project_location | [string](#string) |  | project_location is the location of the project backing the credits in this batch. It is a string of the form <country-code>[-<sub-national-code>[ <postal-code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters. country-code is required, while sub-national-code and postal-code can be added for increasing precision. |
| project_id | [string](#string) |  | project_id is an optional user-specified project ID which can be used instead of an auto-generated ID. If project_id is provided, it must be unique within the credit class and match the regex [A-Za-z0-9]{2,16} or else the operation will fail. If project_id is omitted an ID will automatically be generated. |






<a name="regen.ecocredit.v1alpha2.MsgCreateProjectResponse"></a>

### MsgCreateProjectResponse
MsgCreateProjectResponse is the Msg/CreateProject response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [string](#string) |  | project_id is the ID of the newly created project. |






<a name="regen.ecocredit.v1alpha2.MsgPickFromBasket"></a>

### MsgPickFromBasket
MsgPickFromBasket is the Msg/PickFromBasket request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the owner of the basket tokens. |
| basket_denom | [string](#string) |  | basket_denom is the basket denom to pick credits from. |
| credits | [BasketCredit](#regen.ecocredit.v1alpha2.BasketCredit) | repeated | credits are the units of credits being picked from the basket |
| retirement_location | [string](#string) |  | retirement_location is the optional retirement location for the credits which will be used only if retire_on_take is true for this basket. |






<a name="regen.ecocredit.v1alpha2.MsgPickFromBasketResponse"></a>

### MsgPickFromBasketResponse
MsgPickFromBasketResponse is the Msg/PickFromBasket response type.






<a name="regen.ecocredit.v1alpha2.MsgRetire"></a>

### MsgRetire
MsgRetire is the Msg/Retire request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| holder | [string](#string) |  | holder is the credit holder address. |
| credits | [MsgRetire.RetireCredits](#regen.ecocredit.v1alpha2.MsgRetire.RetireCredits) | repeated | credits are the credits being retired. |
| location | [string](#string) |  | location is the location of the beneficiary or buyer of the retired credits. It is a string of the form <country-code>[-<sub-national-code>[ <postal-code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters. |






<a name="regen.ecocredit.v1alpha2.MsgRetire.RetireCredits"></a>

### MsgRetire.RetireCredits
RetireCredits specifies a batch and the number of credits being retired.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| amount | [string](#string) |  | amount is the number of credits being retired. Decimal values are acceptable within the precision returned by Query/Precision. |






<a name="regen.ecocredit.v1alpha2.MsgRetireResponse"></a>

### MsgRetireResponse
MsgRetire is the Msg/Retire response type.






<a name="regen.ecocredit.v1alpha2.MsgSell"></a>

### MsgSell
MsgSell is the Msg/Sell request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the address of the owner of the credits being sold. |
| orders | [MsgSell.Order](#regen.ecocredit.v1alpha2.MsgSell.Order) | repeated | orders are the sell orders being created. |






<a name="regen.ecocredit.v1alpha2.MsgSell.Order"></a>

### MsgSell.Order
Order is the content of a new sell order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the credit batch being sold. |
| quantity | [string](#string) |  | quantity is the quantity of credits being sold from this batch. If it is less then the balance of credits the owner has available at the time this sell order is matched, the quantity will be adjusted downwards to the owner's balance. However, if the balance of credits is less than this quantity at the time the sell order is created, the operation will fail. |
| ask_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | ask_price is the price the seller is asking for each unit of the batch_denom. Each credit unit of the batch will be sold for at least the ask_price or more. |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire disables auto-retirement of credits which allows a buyer to disable auto-retirement in their buy order enabling them to resell the credits to another buyer. |






<a name="regen.ecocredit.v1alpha2.MsgSellResponse"></a>

### MsgSellResponse
MsgSellResponse is the Msg/Sell response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_order_ids | [uint64](#uint64) | repeated | sell_order_ids are the sell order IDs of the newly created sell orders. |






<a name="regen.ecocredit.v1alpha2.MsgSend"></a>

### MsgSend
MsgSend is the Msg/Send request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the address of the account sending credits. |
| recipient | [string](#string) |  | sender is the address of the account receiving credits. |
| credits | [MsgSend.SendCredits](#regen.ecocredit.v1alpha2.MsgSend.SendCredits) | repeated | credits are the credits being sent. |






<a name="regen.ecocredit.v1alpha2.MsgSend.SendCredits"></a>

### MsgSend.SendCredits
SendCredits specifies a batch and the number of credits being transferred.
This is split into tradable credits, which will remain tradable on receipt,
and retired credits, which will be retired on receipt.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| tradable_amount | [string](#string) |  | tradable_amount is the number of credits in this transfer that can be traded by the recipient. Decimal values are acceptable within the precision returned by Query/Precision. |
| retired_amount | [string](#string) |  | retired_amount is the number of credits in this transfer that are effectively retired by the issuer on receipt. Decimal values are acceptable within the precision returned by Query/Precision. |
| retirement_location | [string](#string) |  | retirement_location is the location of the beneficiary or buyer of the retired credits. This must be provided if retired_amount is positive. It is a string of the form <country-code>[-<sub-national-code>[ <postal-code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters. |






<a name="regen.ecocredit.v1alpha2.MsgSendResponse"></a>

### MsgSendResponse
MsgSendResponse is the Msg/Send response type.






<a name="regen.ecocredit.v1alpha2.MsgTakeFromBasket"></a>

### MsgTakeFromBasket
MsgTakeFromBasket is the Msg/TakeFromBasket request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the owner of the basket tokens. |
| basket_denom | [string](#string) |  | basket_denom is the basket denom to take credits from. |
| amount | [string](#string) |  | amount is the number of credits to take from the basket. |
| retirement_location | [string](#string) |  | retirement_location is the optional retirement location for the credits which will be used only if retire_on_take is true for this basket. |






<a name="regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse"></a>

### MsgTakeFromBasketResponse
MsgTakeFromBasketResponse is the Msg/TakeFromBasket response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credits | [BasketCredit](#regen.ecocredit.v1alpha2.BasketCredit) | repeated | credits are the credits taken out of the basket. |






<a name="regen.ecocredit.v1alpha2.MsgUpdateClassAdmin"></a>

### MsgUpdateClassAdmin
MsgUpdateClassAdmin is the Msg/UpdateClassAdmin request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | admin is the address of the account that is the admin of the credit class. |
| class_id | [string](#string) |  | class_id is the unique ID of the credit class. |
| new_admin | [string](#string) |  | new_admin is the address of the new admin of the credit class. |






<a name="regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse"></a>

### MsgUpdateClassAdminResponse
MsgUpdateClassAdminResponse is the MsgUpdateClassAdmin response type.






<a name="regen.ecocredit.v1alpha2.MsgUpdateClassIssuers"></a>

### MsgUpdateClassIssuers
MsgUpdateClassIssuers is the Msg/UpdateClassIssuers request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | admin is the address of the account that is the admin of the credit class. |
| class_id | [string](#string) |  | class_id is the unique ID of the credit class. |
| issuers | [string](#string) | repeated | issuers are the updated account addresses of the approved issuers. |






<a name="regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse"></a>

### MsgUpdateClassIssuersResponse
MsgUpdateClassIssuersResponse is the MsgUpdateClassIssuers response type.






<a name="regen.ecocredit.v1alpha2.MsgUpdateClassMetadata"></a>

### MsgUpdateClassMetadata
MsgUpdateClassMetadata is the Msg/UpdateClassMetadata request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | admin is the address of the account that is the admin of the credit class. |
| class_id | [string](#string) |  | class_id is the unique ID of the credit class. |
| metadata | [bytes](#bytes) |  | metadata is the updated arbitrary metadata to be attached to the credit class. |






<a name="regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse"></a>

### MsgUpdateClassMetadataResponse
MsgUpdateClassMetadataResponse is the MsgUpdateClassMetadata response type.






<a name="regen.ecocredit.v1alpha2.MsgUpdateSellOrders"></a>

### MsgUpdateSellOrders
MsgUpdateSellOrders is the Msg/UpdateSellOrders request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | owner is the owner of the sell orders. |
| updates | [MsgUpdateSellOrders.Update](#regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update) | repeated | updates are updates to existing sell orders. |






<a name="regen.ecocredit.v1alpha2.MsgUpdateSellOrders.Update"></a>

### MsgUpdateSellOrders.Update
Update is an update to an existing sell order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sell_order_id | [uint64](#uint64) |  | sell_order_id is the ID of an existing sell order. |
| new_quantity | [string](#string) |  | new_quantity is the updated quantity of credits available to sell, if it is set to zero then the order is cancelled. |
| new_ask_price | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | new_ask_price is the new ask price for this sell order |
| disable_auto_retire | [bool](#bool) |  | disable_auto_retire updates the disable_auto_retire field in the sell order. |






<a name="regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse"></a>

### MsgUpdateSellOrdersResponse
MsgUpdateSellOrdersResponse is the Msg/UpdateSellOrders response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.ecocredit.v1alpha2.Msg"></a>

### Msg
Msg is the regen.ecocredit.v1alpha1 Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateClass | [MsgCreateClass](#regen.ecocredit.v1alpha2.MsgCreateClass) | [MsgCreateClassResponse](#regen.ecocredit.v1alpha2.MsgCreateClassResponse) | CreateClass creates a new credit class with an approved list of issuers and optional metadata. |
| CreateProject | [MsgCreateProject](#regen.ecocredit.v1alpha2.MsgCreateProject) | [MsgCreateProjectResponse](#regen.ecocredit.v1alpha2.MsgCreateProjectResponse) | CreateProject creates a new project within a credit class. |
| CreateBatch | [MsgCreateBatch](#regen.ecocredit.v1alpha2.MsgCreateBatch) | [MsgCreateBatchResponse](#regen.ecocredit.v1alpha2.MsgCreateBatchResponse) | CreateBatch creates a new batch of credits for an existing project. This will create a new batch denom with a fixed supply. Issued credits can be distributed to recipients in either tradable or retired form. |
| Send | [MsgSend](#regen.ecocredit.v1alpha2.MsgSend) | [MsgSendResponse](#regen.ecocredit.v1alpha2.MsgSendResponse) | Send sends tradable credits from one account to another account. Sent credits can either be tradable or retired on receipt. |
| Retire | [MsgRetire](#regen.ecocredit.v1alpha2.MsgRetire) | [MsgRetireResponse](#regen.ecocredit.v1alpha2.MsgRetireResponse) | Retire retires a specified number of credits in the holder's account. |
| Cancel | [MsgCancel](#regen.ecocredit.v1alpha2.MsgCancel) | [MsgCancelResponse](#regen.ecocredit.v1alpha2.MsgCancelResponse) | Cancel removes a number of credits from the holder's account and also deducts them from the tradable supply, effectively cancelling their issuance on Regen Ledger |
| UpdateClassAdmin | [MsgUpdateClassAdmin](#regen.ecocredit.v1alpha2.MsgUpdateClassAdmin) | [MsgUpdateClassAdminResponse](#regen.ecocredit.v1alpha2.MsgUpdateClassAdminResponse) | UpdateClassAdmin updates the credit class admin |
| UpdateClassIssuers | [MsgUpdateClassIssuers](#regen.ecocredit.v1alpha2.MsgUpdateClassIssuers) | [MsgUpdateClassIssuersResponse](#regen.ecocredit.v1alpha2.MsgUpdateClassIssuersResponse) | UpdateClassIssuers updates the credit class issuer list |
| UpdateClassMetadata | [MsgUpdateClassMetadata](#regen.ecocredit.v1alpha2.MsgUpdateClassMetadata) | [MsgUpdateClassMetadataResponse](#regen.ecocredit.v1alpha2.MsgUpdateClassMetadataResponse) | UpdateClassMetadata updates the credit class metadata |
| Sell | [MsgSell](#regen.ecocredit.v1alpha2.MsgSell) | [MsgSellResponse](#regen.ecocredit.v1alpha2.MsgSellResponse) | Sell creates new sell orders. |
| UpdateSellOrders | [MsgUpdateSellOrders](#regen.ecocredit.v1alpha2.MsgUpdateSellOrders) | [MsgUpdateSellOrdersResponse](#regen.ecocredit.v1alpha2.MsgUpdateSellOrdersResponse) | UpdateSellOrders updates existing sell orders. |
| Buy | [MsgBuy](#regen.ecocredit.v1alpha2.MsgBuy) | [MsgBuyResponse](#regen.ecocredit.v1alpha2.MsgBuyResponse) | Buy creates credit buy orders. |
| AllowAskDenom | [MsgAllowAskDenom](#regen.ecocredit.v1alpha2.MsgAllowAskDenom) | [MsgAllowAskDenomResponse](#regen.ecocredit.v1alpha2.MsgAllowAskDenomResponse) | AllowAskDenom is a governance operation which authorizes a new ask denom to be used in sell orders |
| CreateBasket | [MsgCreateBasket](#regen.ecocredit.v1alpha2.MsgCreateBasket) | [MsgCreateBasketResponse](#regen.ecocredit.v1alpha2.MsgCreateBasketResponse) | CreateBasket creates a bank denom which wraps credits. |
| AddToBasket | [MsgAddToBasket](#regen.ecocredit.v1alpha2.MsgAddToBasket) | [MsgAddToBasketResponse](#regen.ecocredit.v1alpha2.MsgAddToBasketResponse) | AddToBasket adds credits to a basket in return for basket tokens. |
| TakeFromBasket | [MsgTakeFromBasket](#regen.ecocredit.v1alpha2.MsgTakeFromBasket) | [MsgTakeFromBasketResponse](#regen.ecocredit.v1alpha2.MsgTakeFromBasketResponse) | TakeFromBasket takes credits from a basket without regard for which credits they are. The credits will be auto-retired if disable_auto_retire is false. Credits will be chosen randomly using the previous block hash as a consensus source of randomness. More concretely, the implementation is as follows: - take the previous block hash and convert it into an uint64, - given the total number of different credits within the basket `n`, the first credits that will get picked correspond to: hash modulo n (in terms of order), - then if we need to take more credits, we get some from the next one and so on. |
| PickFromBasket | [MsgPickFromBasket](#regen.ecocredit.v1alpha2.MsgPickFromBasket) | [MsgPickFromBasketResponse](#regen.ecocredit.v1alpha2.MsgPickFromBasketResponse) | PickFromBasket picks specific credits from a basket. If allow_picking is set to false, then only an address which deposited credits in the basket can pick those credits. All other addresses will be blocked from picking those credits. The credits will be auto-retired if disable_auto_retire is false unless the credits were previously put into the basket by the address picking them from the basket, in which case they will remain tradable. This functionality allows the owner of a credit to have more control over the credits they are putting in baskets then ordinary users to deal with the scenario where basket tokens end up being worth significantly less than the credits on their own. |

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

