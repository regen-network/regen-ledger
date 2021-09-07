 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/ecocredit/v1alpha1/events.proto](#regen/ecocredit/v1alpha1/events.proto)
    - [EventCancel](#regen.ecocredit.v1alpha1.EventCancel)
    - [EventCreateBatch](#regen.ecocredit.v1alpha1.EventCreateBatch)
    - [EventCreateClass](#regen.ecocredit.v1alpha1.EventCreateClass)
    - [EventReceive](#regen.ecocredit.v1alpha1.EventReceive)
    - [EventRetire](#regen.ecocredit.v1alpha1.EventRetire)
  
- [regen/ecocredit/v1alpha1/types.proto](#regen/ecocredit/v1alpha1/types.proto)
    - [BatchInfo](#regen.ecocredit.v1alpha1.BatchInfo)
    - [ClassInfo](#regen.ecocredit.v1alpha1.ClassInfo)
    - [CreditType](#regen.ecocredit.v1alpha1.CreditType)
    - [CreditTypeSeq](#regen.ecocredit.v1alpha1.CreditTypeSeq)
    - [Params](#regen.ecocredit.v1alpha1.Params)
  
- [regen/ecocredit/v1alpha1/genesis.proto](#regen/ecocredit/v1alpha1/genesis.proto)
    - [Balance](#regen.ecocredit.v1alpha1.Balance)
    - [GenesisState](#regen.ecocredit.v1alpha1.GenesisState)
    - [Supply](#regen.ecocredit.v1alpha1.Supply)
  
- [regen/ecocredit/v1alpha1/query.proto](#regen/ecocredit/v1alpha1/query.proto)
    - [QueryBalanceRequest](#regen.ecocredit.v1alpha1.QueryBalanceRequest)
    - [QueryBalanceResponse](#regen.ecocredit.v1alpha1.QueryBalanceResponse)
    - [QueryBatchInfoRequest](#regen.ecocredit.v1alpha1.QueryBatchInfoRequest)
    - [QueryBatchInfoResponse](#regen.ecocredit.v1alpha1.QueryBatchInfoResponse)
    - [QueryBatchesRequest](#regen.ecocredit.v1alpha1.QueryBatchesRequest)
    - [QueryBatchesResponse](#regen.ecocredit.v1alpha1.QueryBatchesResponse)
    - [QueryClassInfoRequest](#regen.ecocredit.v1alpha1.QueryClassInfoRequest)
    - [QueryClassInfoResponse](#regen.ecocredit.v1alpha1.QueryClassInfoResponse)
    - [QueryClassesRequest](#regen.ecocredit.v1alpha1.QueryClassesRequest)
    - [QueryClassesResponse](#regen.ecocredit.v1alpha1.QueryClassesResponse)
    - [QueryCreditTypesRequest](#regen.ecocredit.v1alpha1.QueryCreditTypesRequest)
    - [QueryCreditTypesResponse](#regen.ecocredit.v1alpha1.QueryCreditTypesResponse)
    - [QuerySupplyRequest](#regen.ecocredit.v1alpha1.QuerySupplyRequest)
    - [QuerySupplyResponse](#regen.ecocredit.v1alpha1.QuerySupplyResponse)
  
    - [Query](#regen.ecocredit.v1alpha1.Query)
  
- [regen/ecocredit/v1alpha1/tx.proto](#regen/ecocredit/v1alpha1/tx.proto)
    - [MsgCancel](#regen.ecocredit.v1alpha1.MsgCancel)
    - [MsgCancel.CancelCredits](#regen.ecocredit.v1alpha1.MsgCancel.CancelCredits)
    - [MsgCancelResponse](#regen.ecocredit.v1alpha1.MsgCancelResponse)
    - [MsgCreateBatch](#regen.ecocredit.v1alpha1.MsgCreateBatch)
    - [MsgCreateBatch.BatchIssuance](#regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance)
    - [MsgCreateBatchResponse](#regen.ecocredit.v1alpha1.MsgCreateBatchResponse)
    - [MsgCreateClass](#regen.ecocredit.v1alpha1.MsgCreateClass)
    - [MsgCreateClassResponse](#regen.ecocredit.v1alpha1.MsgCreateClassResponse)
    - [MsgRetire](#regen.ecocredit.v1alpha1.MsgRetire)
    - [MsgRetire.RetireCredits](#regen.ecocredit.v1alpha1.MsgRetire.RetireCredits)
    - [MsgRetireResponse](#regen.ecocredit.v1alpha1.MsgRetireResponse)
    - [MsgSend](#regen.ecocredit.v1alpha1.MsgSend)
    - [MsgSend.SendCredits](#regen.ecocredit.v1alpha1.MsgSend.SendCredits)
    - [MsgSendResponse](#regen.ecocredit.v1alpha1.MsgSendResponse)
  
    - [Msg](#regen.ecocredit.v1alpha1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/ecocredit/v1alpha1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1alpha1/events.proto



<a name="regen.ecocredit.v1alpha1.EventCancel"></a>

### EventCancel
EventCancel is an event emitted when credits are cancelled. When credits are
cancelled from multiple batches in the same transaction, a separate event is
emitted for each batch_denom. This allows for easier indexing.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| canceller | [string](#string) |  | canceller is the account which has cancelled the credits, which should be the holder of the credits. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch. |
| amount | [string](#string) |  | amount is the decimal number of credits that have been cancelled. |






<a name="regen.ecocredit.v1alpha1.EventCreateBatch"></a>

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
| project_location | [string](#string) |  | project_location is the location of the project backing the credits in this batch. Full documentation can be found in MsgCreateBatch.project_location. |






<a name="regen.ecocredit.v1alpha1.EventCreateClass"></a>

### EventCreateClass
EventCreateClass is an event emitted when a credit class is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class. |
| admin | [string](#string) |  | admin is the admin of the credit class. |






<a name="regen.ecocredit.v1alpha1.EventReceive"></a>

### EventReceive
EventReceive is an event emitted when credits are received either upon
creation of a new batch or upon transfer. Each batch_denom created or
transferred will result in a separate EventReceive for easy indexing.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the sender of the credits in the case that this event is the result of a transfer. It will not be set when credits are received at initial issuance. |
| recipient | [string](#string) |  | recipient is the recipient of the credits |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch. |
| amount | [string](#string) |  | amount is the decimal number of both tradable and retired credits received. |






<a name="regen.ecocredit.v1alpha1.EventRetire"></a>

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





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/v1alpha1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1alpha1/types.proto



<a name="regen.ecocredit.v1alpha1.BatchInfo"></a>

### BatchInfo
BatchInfo represents the high-level on-chain information for a credit batch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch. |
| issuer | [string](#string) |  | issuer is the issuer of the credit batch. |
| total_amount | [string](#string) |  | total_amount is the total number of active credits in the credit batch. Some of the issued credits may be cancelled and will be removed from total_amount and tracked in amount_cancelled. total_amount and amount_cancelled will always sum to the original amount of credits that were issued. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata to attached to the credit batch. |
| amount_cancelled | [string](#string) |  | amount_cancelled is the number of credits in the batch that have been cancelled, effectively undoing there issuance. The sum of total_amount and amount_cancelled will always sum to the original amount of credits that were issued. |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_date is the beginning of the period during which this credit batch was quantified and verified. |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | end_date is the end of the period during which this credit batch was quantified and verified. |
| project_location | [string](#string) |  | project_location is the location of the project backing the credits in this batch. Full documentation can be found in MsgCreateBatch.project_location. |






<a name="regen.ecocredit.v1alpha1.ClassInfo"></a>

### ClassInfo
ClassInfo represents the high-level on-chain information for a credit class.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class. |
| admin | [string](#string) |  | admin is the admin of the credit class. |
| issuers | [string](#string) | repeated | issuers are the approved issuers of the credit class. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata to attached to the credit class. |
| credit_type | [CreditType](#regen.ecocredit.v1alpha1.CreditType) |  | credit_type describes the type of credit (e.g. carbon, biodiversity), as well as unit and precision. |
| num_batches | [uint64](#uint64) |  | The number of batches issued in this credit class. |






<a name="regen.ecocredit.v1alpha1.CreditType"></a>

### CreditType
CreditType defines the measurement unit/precision of a certain credit type
(e.g. carbon, biodiversity...)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | the type of credit (e.g. carbon, biodiversity, etc) |
| abbreviation | [string](#string) |  | abbreviation is a 1-3 character uppercase abbreviation of the CreditType name, used in batch denominations within the CreditType. It must be unique. |
| unit | [string](#string) |  | the measurement unit (e.g. kg, ton, etc) |
| precision | [uint32](#uint32) |  | the decimal precision |






<a name="regen.ecocredit.v1alpha1.CreditTypeSeq"></a>

### CreditTypeSeq
CreditTypeSeq associates a sequence number with a credit type abbreviation.
This represents the number of credit classes created with that credit type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| abbreviation | [string](#string) |  | The credit type abbreviation |
| seq_number | [uint64](#uint64) |  | The sequence number of classes of the credit type |






<a name="regen.ecocredit.v1alpha1.Params"></a>

### Params
Params defines the updatable global parameters of the ecocredit module for
use with the x/params module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credit_class_fee | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | credit_class_fee is the fixed fee charged on creation of a new credit class |
| allowed_class_creators | [string](#string) | repeated | allowed_class_creators is an allowlist defining the addresses with the required permissions to create credit classes |
| allowlist_enabled | [bool](#bool) |  | allowlist_enabled is a param that enables/disables the allowlist for credit creation |
| credit_types | [CreditType](#regen.ecocredit.v1alpha1.CreditType) | repeated | credit_types is a list of definitions for credit types |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/v1alpha1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1alpha1/genesis.proto



<a name="regen.ecocredit.v1alpha1.Balance"></a>

### Balance
Balance represents tradable or retired units of a credit batch with an
account address, batch_denom, and balance.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address is the account address of the account holding credits. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| tradable_balance | [string](#string) |  | tradable_balance is the tradable balance of the credit batch. |
| retired_balance | [string](#string) |  | retired_balance is the retired balance of the credit batch. |






<a name="regen.ecocredit.v1alpha1.GenesisState"></a>

### GenesisState
GenesisState defines ecocredit module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#regen.ecocredit.v1alpha1.Params) |  | Params contains the updateable global parameters for use with the x/params module |
| class_info | [ClassInfo](#regen.ecocredit.v1alpha1.ClassInfo) | repeated | class_info is the list of credit class info. |
| batch_info | [BatchInfo](#regen.ecocredit.v1alpha1.BatchInfo) | repeated | batch_info is the list of credit batch info. |
| sequences | [CreditTypeSeq](#regen.ecocredit.v1alpha1.CreditTypeSeq) | repeated | sequences is the list of credit type sequence. |
| balances | [Balance](#regen.ecocredit.v1alpha1.Balance) | repeated | balances is the list of credit batch tradable/retired units. |
| supplies | [Supply](#regen.ecocredit.v1alpha1.Supply) | repeated | supplies is the list of credit batch tradable/retired supply. |






<a name="regen.ecocredit.v1alpha1.Supply"></a>

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



<a name="regen/ecocredit/v1alpha1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1alpha1/query.proto



<a name="regen.ecocredit.v1alpha1.QueryBalanceRequest"></a>

### QueryBalanceRequest
QueryBalanceRequest is the Query/Balance request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  | account is the address of the account whose balance is being queried. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch balance to query. |






<a name="regen.ecocredit.v1alpha1.QueryBalanceResponse"></a>

### QueryBalanceResponse
QueryBalanceResponse is the Query/Balance response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tradable_amount | [string](#string) |  | tradable_amount is the decimal number of tradable credits. |
| retired_amount | [string](#string) |  | retired_amount is the decimal number of retired credits. |






<a name="regen.ecocredit.v1alpha1.QueryBatchInfoRequest"></a>

### QueryBatchInfoRequest
QueryBatchInfoRequest is the Query/BatchInfo request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch to query. |






<a name="regen.ecocredit.v1alpha1.QueryBatchInfoResponse"></a>

### QueryBatchInfoResponse
QueryBatchInfoResponse is the Query/BatchInfo response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [BatchInfo](#regen.ecocredit.v1alpha1.BatchInfo) |  | info is the BatchInfo for the credit batch. |






<a name="regen.ecocredit.v1alpha1.QueryBatchesRequest"></a>

### QueryBatchesRequest
QueryBatchesRequest is the Query/Batches request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of the credit class to query. |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1alpha1.QueryBatchesResponse"></a>

### QueryBatchesResponse
QueryBatchesResponse is the Query/Batches response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batches | [BatchInfo](#regen.ecocredit.v1alpha1.BatchInfo) | repeated | batches are the fetched credit batches within the class. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.v1alpha1.QueryClassInfoRequest"></a>

### QueryClassInfoRequest
QueryClassInfoRequest is the Query/ClassInfo request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class to query. |






<a name="regen.ecocredit.v1alpha1.QueryClassInfoResponse"></a>

### QueryClassInfoResponse
QueryClassInfoResponse is the Query/ClassInfo request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [ClassInfo](#regen.ecocredit.v1alpha1.ClassInfo) |  | info is the ClassInfo for the credit class. |






<a name="regen.ecocredit.v1alpha1.QueryClassesRequest"></a>

### QueryClassesRequest
QueryClassesRequest is the Query/Classes request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1alpha1.QueryClassesResponse"></a>

### QueryClassesResponse
QueryClassesResponse is the Query/Classes response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| classes | [ClassInfo](#regen.ecocredit.v1alpha1.ClassInfo) | repeated | classes are the fetched credit classes. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.v1alpha1.QueryCreditTypesRequest"></a>

### QueryCreditTypesRequest
QueryCreditTypesRequest is the Query/Credit_Types request type






<a name="regen.ecocredit.v1alpha1.QueryCreditTypesResponse"></a>

### QueryCreditTypesResponse
QueryCreditTypesRequest is the Query/Credit_Types response type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credit_types | [CreditType](#regen.ecocredit.v1alpha1.CreditType) | repeated | list of credit types |






<a name="regen.ecocredit.v1alpha1.QuerySupplyRequest"></a>

### QuerySupplyRequest
QuerySupplyRequest is the Query/Supply request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch to query. |






<a name="regen.ecocredit.v1alpha1.QuerySupplyResponse"></a>

### QuerySupplyResponse
QuerySupplyResponse is the Query/Supply response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tradable_supply | [string](#string) |  | tradable_supply is the decimal number of tradable credits in the batch supply. |
| retired_supply | [string](#string) |  | retired_supply is the decimal number of retired credits in the batch supply. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.ecocredit.v1alpha1.Query"></a>

### Query
Msg is the regen.ecocredit.v1alpha1 Query service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Classes | [QueryClassesRequest](#regen.ecocredit.v1alpha1.QueryClassesRequest) | [QueryClassesResponse](#regen.ecocredit.v1alpha1.QueryClassesResponse) | Classes queries for all credit classes with pagination. |
| ClassInfo | [QueryClassInfoRequest](#regen.ecocredit.v1alpha1.QueryClassInfoRequest) | [QueryClassInfoResponse](#regen.ecocredit.v1alpha1.QueryClassInfoResponse) | ClassInfo queries for information on a credit class. |
| Batches | [QueryBatchesRequest](#regen.ecocredit.v1alpha1.QueryBatchesRequest) | [QueryBatchesResponse](#regen.ecocredit.v1alpha1.QueryBatchesResponse) | Batches queries for all batches in the given credit class with pagination. |
| BatchInfo | [QueryBatchInfoRequest](#regen.ecocredit.v1alpha1.QueryBatchInfoRequest) | [QueryBatchInfoResponse](#regen.ecocredit.v1alpha1.QueryBatchInfoResponse) | BatchInfo queries for information on a credit batch. |
| Balance | [QueryBalanceRequest](#regen.ecocredit.v1alpha1.QueryBalanceRequest) | [QueryBalanceResponse](#regen.ecocredit.v1alpha1.QueryBalanceResponse) | Balance queries the balance (both tradable and retired) of a given credit batch for a given account. |
| Supply | [QuerySupplyRequest](#regen.ecocredit.v1alpha1.QuerySupplyRequest) | [QuerySupplyResponse](#regen.ecocredit.v1alpha1.QuerySupplyResponse) | Supply queries the tradable and retired supply of a credit batch. |
| CreditTypes | [QueryCreditTypesRequest](#regen.ecocredit.v1alpha1.QueryCreditTypesRequest) | [QueryCreditTypesResponse](#regen.ecocredit.v1alpha1.QueryCreditTypesResponse) | CreditTypes returns the list of allowed types that credit classes can have. See Types/CreditType for more details. |

 <!-- end services -->



<a name="regen/ecocredit/v1alpha1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1alpha1/tx.proto



<a name="regen.ecocredit.v1alpha1.MsgCancel"></a>

### MsgCancel
MsgCancel is the Msg/Cancel request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| holder | [string](#string) |  | holder is the credit holder address. |
| credits | [MsgCancel.CancelCredits](#regen.ecocredit.v1alpha1.MsgCancel.CancelCredits) | repeated | credits are the credits being cancelled. |






<a name="regen.ecocredit.v1alpha1.MsgCancel.CancelCredits"></a>

### MsgCancel.CancelCredits
CancelCredits specifies a batch and the number of credits being cancelled.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| amount | [string](#string) |  | amount is the number of credits being cancelled. Decimal values are acceptable within the precision returned by Query/Precision. |






<a name="regen.ecocredit.v1alpha1.MsgCancelResponse"></a>

### MsgCancelResponse
MsgCancelResponse is the Msg/Cancel response type.






<a name="regen.ecocredit.v1alpha1.MsgCreateBatch"></a>

### MsgCreateBatch
MsgCreateBatch is the Msg/CreateBatch request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| issuer | [string](#string) |  | issuer is the address of the batch issuer. |
| class_id | [string](#string) |  | class_id is the unique ID of the class. |
| issuance | [MsgCreateBatch.BatchIssuance](#regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance) | repeated | issuance are the credits issued in the batch. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata to attached to the credit batch. |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_date is the beginning of the period during which this credit batch was quantified and verified. |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | end_date is the end of the period during which this credit batch was quantified and verified. |
| project_location | [string](#string) |  | project_location is the location of the project backing the credits in this batch. It is a string of the form <country-code>[-<sub-national-code>[ <postal-code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters. country-code is required, while sub-national-code and postal-code can be added for increasing precision. |






<a name="regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance"></a>

### MsgCreateBatch.BatchIssuance
BatchIssuance represents the issuance of some credits in a batch to a
single recipient.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recipient | [string](#string) |  | recipient is the account of the recipient. |
| tradable_amount | [string](#string) |  | tradable_amount is the number of credits in this issuance that can be traded by this recipient. Decimal values are acceptable. |
| retired_amount | [string](#string) |  | retired_amount is the number of credits in this issuance that are effectively retired by the issuer on receipt. Decimal values are acceptable. |
| retirement_location | [string](#string) |  | retirement_location is the location of the beneficiary or buyer of the retired credits. This must be provided if retired_amount is positive. It is a string of the form <country-code>[-<sub-national-code>[ <postal-code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters. |






<a name="regen.ecocredit.v1alpha1.MsgCreateBatchResponse"></a>

### MsgCreateBatchResponse
MsgCreateBatchResponse is the Msg/CreateBatch response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique denomination ID of the newly created batch. |






<a name="regen.ecocredit.v1alpha1.MsgCreateClass"></a>

### MsgCreateClass
MsgCreateClass is the Msg/CreateClass request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | admin is the address of the account which can manage the credit class. The admin has permissions to change the list of issuers and perform other administrative operations. |
| issuers | [string](#string) | repeated | issuers are the account addresses of the approved issuers. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata to attached to the credit class. |
| credit_type_name | [string](#string) |  | credit_type describes the type of credit (e.g. "carbon", "biodiversity"). |






<a name="regen.ecocredit.v1alpha1.MsgCreateClassResponse"></a>

### MsgCreateClassResponse
MsgCreateClassResponse is the Msg/CreateClass response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of the newly created credit class. |






<a name="regen.ecocredit.v1alpha1.MsgRetire"></a>

### MsgRetire
MsgRetire is the Msg/Retire request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| holder | [string](#string) |  | holder is the credit holder address. |
| credits | [MsgRetire.RetireCredits](#regen.ecocredit.v1alpha1.MsgRetire.RetireCredits) | repeated | credits are the credits being retired. |
| location | [string](#string) |  | location is the location of the beneficiary or buyer of the retired credits. It is a string of the form <country-code>[-<sub-national-code>[ <postal-code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters. |






<a name="regen.ecocredit.v1alpha1.MsgRetire.RetireCredits"></a>

### MsgRetire.RetireCredits
RetireCredits specifies a batch and the number of credits being retired.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| amount | [string](#string) |  | amount is the number of credits being retired. Decimal values are acceptable within the precision returned by Query/Precision. |






<a name="regen.ecocredit.v1alpha1.MsgRetireResponse"></a>

### MsgRetireResponse
MsgRetire is the Msg/Retire response type.






<a name="regen.ecocredit.v1alpha1.MsgSend"></a>

### MsgSend
MsgSend is the Msg/Send request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the address of the account sending credits. |
| recipient | [string](#string) |  | sender is the address of the account receiving credits. |
| credits | [MsgSend.SendCredits](#regen.ecocredit.v1alpha1.MsgSend.SendCredits) | repeated | credits are the credits being sent. |






<a name="regen.ecocredit.v1alpha1.MsgSend.SendCredits"></a>

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






<a name="regen.ecocredit.v1alpha1.MsgSendResponse"></a>

### MsgSendResponse
MsgSendResponse is the Msg/Send response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.ecocredit.v1alpha1.Msg"></a>

### Msg
Msg is the regen.ecocredit.v1alpha1 Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateClass | [MsgCreateClass](#regen.ecocredit.v1alpha1.MsgCreateClass) | [MsgCreateClassResponse](#regen.ecocredit.v1alpha1.MsgCreateClassResponse) | CreateClass creates a new credit class with an approved list of issuers and optional metadata. |
| CreateBatch | [MsgCreateBatch](#regen.ecocredit.v1alpha1.MsgCreateBatch) | [MsgCreateBatchResponse](#regen.ecocredit.v1alpha1.MsgCreateBatchResponse) | CreateBatch creates a new batch of credits for an existing credit class. This will create a new batch denom with a fixed supply. Issued credits can be distributed to recipients in either tradable or retired form. |
| Send | [MsgSend](#regen.ecocredit.v1alpha1.MsgSend) | [MsgSendResponse](#regen.ecocredit.v1alpha1.MsgSendResponse) | Send sends tradable credits from one account to another account. Sent credits can either be tradable or retired on receipt. |
| Retire | [MsgRetire](#regen.ecocredit.v1alpha1.MsgRetire) | [MsgRetireResponse](#regen.ecocredit.v1alpha1.MsgRetireResponse) | Retire retires a specified number of credits in the holder's account. |
| Cancel | [MsgCancel](#regen.ecocredit.v1alpha1.MsgCancel) | [MsgCancelResponse](#regen.ecocredit.v1alpha1.MsgCancelResponse) | Cancel removes a number of credits from the holder's account and also deducts them from the tradable supply, effectively cancelling their issuance on Regen Ledger |

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

