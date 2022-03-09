 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/ecocredit/v1/events.proto](#regen/ecocredit/v1/events.proto)
    - [EventCancel](#regen.ecocredit.v1.EventCancel)
    - [EventCreateBatch](#regen.ecocredit.v1.EventCreateBatch)
    - [EventCreateClass](#regen.ecocredit.v1.EventCreateClass)
    - [EventCreateProject](#regen.ecocredit.v1.EventCreateProject)
    - [EventReceive](#regen.ecocredit.v1.EventReceive)
    - [EventRetire](#regen.ecocredit.v1.EventRetire)
  
- [regen/ecocredit/v1/state.proto](#regen/ecocredit/v1/state.proto)
    - [BatchBalance](#regen.ecocredit.v1.BatchBalance)
    - [BatchInfo](#regen.ecocredit.v1.BatchInfo)
    - [BatchSequence](#regen.ecocredit.v1.BatchSequence)
    - [BatchSupply](#regen.ecocredit.v1.BatchSupply)
    - [ClassInfo](#regen.ecocredit.v1.ClassInfo)
    - [ClassIssuer](#regen.ecocredit.v1.ClassIssuer)
    - [ClassSequence](#regen.ecocredit.v1.ClassSequence)
    - [CreditType](#regen.ecocredit.v1.CreditType)
    - [ProjectInfo](#regen.ecocredit.v1.ProjectInfo)
    - [ProjectSequence](#regen.ecocredit.v1.ProjectSequence)
  
- [regen/ecocredit/v1/types.proto](#regen/ecocredit/v1/types.proto)
    - [Params](#regen.ecocredit.v1.Params)
  
- [regen/ecocredit/v1/query.proto](#regen/ecocredit/v1/query.proto)
    - [QueryBalanceRequest](#regen.ecocredit.v1.QueryBalanceRequest)
    - [QueryBalanceResponse](#regen.ecocredit.v1.QueryBalanceResponse)
    - [QueryBatchInfoRequest](#regen.ecocredit.v1.QueryBatchInfoRequest)
    - [QueryBatchInfoResponse](#regen.ecocredit.v1.QueryBatchInfoResponse)
    - [QueryBatchesRequest](#regen.ecocredit.v1.QueryBatchesRequest)
    - [QueryBatchesResponse](#regen.ecocredit.v1.QueryBatchesResponse)
    - [QueryClassInfoRequest](#regen.ecocredit.v1.QueryClassInfoRequest)
    - [QueryClassInfoResponse](#regen.ecocredit.v1.QueryClassInfoResponse)
    - [QueryClassIssuersRequest](#regen.ecocredit.v1.QueryClassIssuersRequest)
    - [QueryClassIssuersResponse](#regen.ecocredit.v1.QueryClassIssuersResponse)
    - [QueryClassesRequest](#regen.ecocredit.v1.QueryClassesRequest)
    - [QueryClassesResponse](#regen.ecocredit.v1.QueryClassesResponse)
    - [QueryCreditTypesRequest](#regen.ecocredit.v1.QueryCreditTypesRequest)
    - [QueryCreditTypesResponse](#regen.ecocredit.v1.QueryCreditTypesResponse)
    - [QueryParamsRequest](#regen.ecocredit.v1.QueryParamsRequest)
    - [QueryParamsResponse](#regen.ecocredit.v1.QueryParamsResponse)
    - [QueryProjectInfoRequest](#regen.ecocredit.v1.QueryProjectInfoRequest)
    - [QueryProjectInfoResponse](#regen.ecocredit.v1.QueryProjectInfoResponse)
    - [QueryProjectsRequest](#regen.ecocredit.v1.QueryProjectsRequest)
    - [QueryProjectsResponse](#regen.ecocredit.v1.QueryProjectsResponse)
    - [QuerySupplyRequest](#regen.ecocredit.v1.QuerySupplyRequest)
    - [QuerySupplyResponse](#regen.ecocredit.v1.QuerySupplyResponse)
  
    - [Query](#regen.ecocredit.v1.Query)
  
- [regen/ecocredit/v1/tx.proto](#regen/ecocredit/v1/tx.proto)
    - [MsgCancel](#regen.ecocredit.v1.MsgCancel)
    - [MsgCancel.CancelCredits](#regen.ecocredit.v1.MsgCancel.CancelCredits)
    - [MsgCancelResponse](#regen.ecocredit.v1.MsgCancelResponse)
    - [MsgCreateBatch](#regen.ecocredit.v1.MsgCreateBatch)
    - [MsgCreateBatch.BatchIssuance](#regen.ecocredit.v1.MsgCreateBatch.BatchIssuance)
    - [MsgCreateBatchResponse](#regen.ecocredit.v1.MsgCreateBatchResponse)
    - [MsgCreateClass](#regen.ecocredit.v1.MsgCreateClass)
    - [MsgCreateClassResponse](#regen.ecocredit.v1.MsgCreateClassResponse)
    - [MsgCreateProject](#regen.ecocredit.v1.MsgCreateProject)
    - [MsgCreateProjectResponse](#regen.ecocredit.v1.MsgCreateProjectResponse)
    - [MsgRetire](#regen.ecocredit.v1.MsgRetire)
    - [MsgRetire.RetireCredits](#regen.ecocredit.v1.MsgRetire.RetireCredits)
    - [MsgRetireResponse](#regen.ecocredit.v1.MsgRetireResponse)
    - [MsgSend](#regen.ecocredit.v1.MsgSend)
    - [MsgSend.SendCredits](#regen.ecocredit.v1.MsgSend.SendCredits)
    - [MsgSendResponse](#regen.ecocredit.v1.MsgSendResponse)
    - [MsgUpdateClassAdmin](#regen.ecocredit.v1.MsgUpdateClassAdmin)
    - [MsgUpdateClassAdminResponse](#regen.ecocredit.v1.MsgUpdateClassAdminResponse)
    - [MsgUpdateClassIssuers](#regen.ecocredit.v1.MsgUpdateClassIssuers)
    - [MsgUpdateClassIssuersResponse](#regen.ecocredit.v1.MsgUpdateClassIssuersResponse)
    - [MsgUpdateClassMetadata](#regen.ecocredit.v1.MsgUpdateClassMetadata)
    - [MsgUpdateClassMetadataResponse](#regen.ecocredit.v1.MsgUpdateClassMetadataResponse)
  
    - [Msg](#regen.ecocredit.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/ecocredit/v1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1/events.proto



<a name="regen.ecocredit.v1.EventCancel"></a>

### EventCancel
EventCancel is an event emitted when credits are cancelled. When credits are
cancelled from multiple batches in the same transaction, a separate event is
emitted for each batch_denom. This allows for easier indexing.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| canceller | [string](#string) |  | canceller is the account which has cancelled the credits, which should be the holder of the credits. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch. |
| amount | [string](#string) |  | amount is the decimal number of credits that have been cancelled. |






<a name="regen.ecocredit.v1.EventCreateBatch"></a>

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
| project_id | [string](#string) |  | project_id is the unique ID of the project this batch belongs to. |






<a name="regen.ecocredit.v1.EventCreateClass"></a>

### EventCreateClass
EventCreateClass is an event emitted when a credit class is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class. |
| admin | [string](#string) |  | admin is the admin of the credit class. |






<a name="regen.ecocredit.v1.EventCreateProject"></a>

### EventCreateProject
EventCreateProject is an event emitted when a project is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [string](#string) |  | project_id is the unique ID of the project. |
| class_id | [string](#string) |  | class_id is the unique ID of credit class for this project. |
| issuer | [string](#string) |  | issuer is the issuer of the credit batches for this project. |
| project_location | [string](#string) |  | project_location is the location of the project. Full documentation can be found in MsgCreateProject.project_location. |






<a name="regen.ecocredit.v1.EventReceive"></a>

### EventReceive
EventReceive is an event emitted when credits are received either via
creation of a new batch, transfer of credits, or taking credits from a
basket. Each batch_denom created, transferred or taken from a basket will
result in a separate EventReceive for easy indexing.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the sender of the credits in the case that this event is the result of a transfer. It will not be set when credits are received at initial issuance or taken from a basket. |
| recipient | [string](#string) |  | recipient is the recipient of the credits. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch. |
| tradable_amount | [string](#string) |  | tradable_amount is the decimal number of tradable credits received. |
| retired_amount | [string](#string) |  | retired_amount is the decimal number of retired credits received. |
| basket_denom | [string](#string) |  | basket_denom is the denom of the basket. When the basket_denom field is set, it indicates that this event was triggered by the transfer of credits from a basket. It will not be set if the credits were transferred or received at initial issuance. |






<a name="regen.ecocredit.v1.EventRetire"></a>

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



<a name="regen/ecocredit/v1/state.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1/state.proto



<a name="regen.ecocredit.v1.BatchBalance"></a>

### BatchBalance
BatchBalance stores each users credit balance


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [bytes](#bytes) |  | address is the address of the credit holder |
| batch_id | [uint64](#uint64) |  | batch_id is the id of the credit batch |
| tradable | [string](#string) |  | tradable is the tradable amount of credits |
| retired | [string](#string) |  | retired is the retired amount of credits |






<a name="regen.ecocredit.v1.BatchInfo"></a>

### BatchInfo
BatchInfo represents the high-level on-chain information for a credit batch.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id is an auto-incrementing integer to succinctly identify the batch |
| project_id | [uint64](#uint64) |  | project_id is the unique ID of the project this batch belongs to. |
| batch_denom | [string](#string) |  | batch_denom is the unique string identifier of the credit batch formed from the project name, batch sequence number and dates. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata attached to the credit batch. |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_date is the beginning of the period during which this credit batch was quantified and verified. |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | end_date is the end of the period during which this credit batch was quantified and verified. |






<a name="regen.ecocredit.v1.BatchSequence"></a>

### BatchSequence
BatchSequence tracks the sequence number for batches within a project


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [string](#string) |  | project_id is the id of the project for a batch |
| next_batch_id | [uint64](#uint64) |  | next_batch_id is a sequence number incrementing on each issued batch |






<a name="regen.ecocredit.v1.BatchSupply"></a>

### BatchSupply
BatchSupply tracks the supply of a credit batch


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_id | [uint64](#uint64) |  | batch_id is the id of the batch |
| tradable_amount | [string](#string) |  | tradable_amount is the total number of tradable credits in the credit batch. Some of the issued credits may be cancelled and will be removed from tradable_amount and tracked in amount_cancelled. tradable_amount + retired_amount + amount_cancelled will always sum to the original credit issuance amount. |
| retired_amount | [string](#string) |  | retired_amount is the total amount of credits that have been retired from the credit batch. |
| cancelled_amount | [string](#string) |  | cancelled_amount is the number of credits in the batch that have been cancelled, effectively undoing the issuance. The sum of total_amount and amount_cancelled will always equal the original credit issuance amount. |






<a name="regen.ecocredit.v1.ClassInfo"></a>

### ClassInfo
ClassInfo represents the high-level on-chain information for a credit class.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id is the unique ID of credit class. |
| name | [string](#string) |  | abbrev is the unique string name for this credit class formed from its credit type and an auto-generated integer. |
| admin | [bytes](#bytes) |  | admin is the admin of the credit class. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata to attached to the credit class. |
| credit_type | [string](#string) |  | credit_type is the abbreviation of the credit type. |






<a name="regen.ecocredit.v1.ClassIssuer"></a>

### ClassIssuer
ClassIssuers is a JOIN table for Class Info that stores the credit class issuers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [uint64](#uint64) |  | class_id is the row ID of a credit class. |
| issuer | [bytes](#bytes) |  | issuer is the approved issuer of the credit class. |






<a name="regen.ecocredit.v1.ClassSequence"></a>

### ClassSequence
ClassSequence is a sequence number for creating credit class identifiers for each
credit type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credit_type | [string](#string) |  | credit_type is the credit type abbreviation |
| next_class_id | [uint64](#uint64) |  | next_class_id is the next class ID for this credit type |






<a name="regen.ecocredit.v1.CreditType"></a>

### CreditType
CreditType defines the measurement unit/precision of a certain credit type
(e.g. carbon, biodiversity...)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| abbreviation | [string](#string) |  | abbreviation is a 1-3 character uppercase abbreviation of the CreditType name, used in batch denominations within the CreditType. It must be unique. |
| name | [string](#string) |  | the type of credit (e.g. carbon, biodiversity, etc) |
| unit | [string](#string) |  | the measurement unit (e.g. kg, ton, etc) |
| precision | [uint32](#uint32) |  | the decimal precision |






<a name="regen.ecocredit.v1.ProjectInfo"></a>

### ProjectInfo
ProjectInfo represents the high-level on-chain information for a project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id is the unique ID of the project |
| name | [string](#string) |  | name is an unique name of this project formed from its credit class name and either an auto-generated number or user provided string. |
| class_id | [uint64](#uint64) |  | class_id is the ID of credit class for this project. |
| project_location | [string](#string) |  | project_location is the location of the project. Full documentation can be found in MsgCreateProject.project_location. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata attached to the project. |






<a name="regen.ecocredit.v1.ProjectSequence"></a>

### ProjectSequence
ProjectSequence stores and increments the sequence number for projects
within a given credit class.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [uint64](#uint64) |  | class_id is the id of the credit class |
| next_project_id | [uint64](#uint64) |  | next_project_id is the sequence number for the project |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/v1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1/types.proto



<a name="regen.ecocredit.v1.Params"></a>

### Params
Params defines the updatable global parameters of the ecocredit module for
use with the x/params module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credit_class_fee | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | credit_class_fee is the fixed fee charged on creation of a new credit class |
| allowed_class_creators | [string](#string) | repeated | allowed_class_creators is an allowlist defining the addresses with the required permissions to create credit classes |
| allowlist_enabled | [bool](#bool) |  | allowlist_enabled is a param that enables/disables the allowlist for credit creation |
| credit_types | [CreditType](#regen.ecocredit.v1.CreditType) | repeated | credit_types is a list of definitions for credit types |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/ecocredit/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1/query.proto



<a name="regen.ecocredit.v1.QueryBalanceRequest"></a>

### QueryBalanceRequest
QueryBalanceRequest is the Query/Balance request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  | account is the address of the account whose balance is being queried. |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch balance to query. |






<a name="regen.ecocredit.v1.QueryBalanceResponse"></a>

### QueryBalanceResponse
QueryBalanceResponse is the Query/Balance response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tradable_amount | [string](#string) |  | tradable_amount is the decimal number of tradable credits. |
| retired_amount | [string](#string) |  | retired_amount is the decimal number of retired credits. |






<a name="regen.ecocredit.v1.QueryBatchInfoRequest"></a>

### QueryBatchInfoRequest
QueryBatchInfoRequest is the Query/BatchInfo request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch to query. |






<a name="regen.ecocredit.v1.QueryBatchInfoResponse"></a>

### QueryBatchInfoResponse
QueryBatchInfoResponse is the Query/BatchInfo response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [BatchInfo](#regen.ecocredit.v1.BatchInfo) |  | info is the BatchInfo for the credit batch. |






<a name="regen.ecocredit.v1.QueryBatchesRequest"></a>

### QueryBatchesRequest
QueryBatchesRequest is the Query/Batches request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [string](#string) |  | project_id is the unique ID of the project to query. |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1.QueryBatchesResponse"></a>

### QueryBatchesResponse
QueryBatchesResponse is the Query/Batches response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batches | [BatchInfo](#regen.ecocredit.v1.BatchInfo) | repeated | batches are the fetched credit batches within the project. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.v1.QueryClassInfoRequest"></a>

### QueryClassInfoRequest
QueryClassInfoRequest is the Query/ClassInfo request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class to query. |






<a name="regen.ecocredit.v1.QueryClassInfoResponse"></a>

### QueryClassInfoResponse
QueryClassInfoResponse is the Query/ClassInfo request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [ClassInfo](#regen.ecocredit.v1.ClassInfo) |  | info is the ClassInfo for the credit class. |






<a name="regen.ecocredit.v1.QueryClassIssuersRequest"></a>

### QueryClassIssuersRequest
QueryClassIssuersRequest is the Query/ClassIssuers request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class to query. |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1.QueryClassIssuersResponse"></a>

### QueryClassIssuersResponse
QueryClassIssuersRequest is the Query/ClassIssuers response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| issuers | [string](#string) | repeated | issuers is a list of issuers for the credit class |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.v1.QueryClassesRequest"></a>

### QueryClassesRequest
QueryClassesRequest is the Query/Classes request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1.QueryClassesResponse"></a>

### QueryClassesResponse
QueryClassesResponse is the Query/Classes response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| classes | [ClassInfo](#regen.ecocredit.v1.ClassInfo) | repeated | classes are the fetched credit classes. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.v1.QueryCreditTypesRequest"></a>

### QueryCreditTypesRequest
QueryCreditTypesRequest is the Query/Credit_Types request type






<a name="regen.ecocredit.v1.QueryCreditTypesResponse"></a>

### QueryCreditTypesResponse
QueryCreditTypesRequest is the Query/Credit_Types response type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credit_types | [CreditType](#regen.ecocredit.v1.CreditType) | repeated | list of credit types |






<a name="regen.ecocredit.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the Query/Params request type.






<a name="regen.ecocredit.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the Query/Params response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#regen.ecocredit.v1.Params) |  | params defines the parameters of the ecocredit module. |






<a name="regen.ecocredit.v1.QueryProjectInfoRequest"></a>

### QueryProjectInfoRequest
QueryProjectInfoRequest is the Query/Project request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [string](#string) |  | project_id is the unique ID of the project to query. |






<a name="regen.ecocredit.v1.QueryProjectInfoResponse"></a>

### QueryProjectInfoResponse
QueryProjectInfoResponse is the Query/Project response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [ProjectInfo](#regen.ecocredit.v1.ProjectInfo) |  | info is the ProjectInfo for the project. |






<a name="regen.ecocredit.v1.QueryProjectsRequest"></a>

### QueryProjectsRequest
QueryProjectsRequest is the Query/Projects request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of credit class to query. |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.ecocredit.v1.QueryProjectsResponse"></a>

### QueryProjectsResponse
QueryProjectsResponse is the Query/Projects response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| projects | [ProjectInfo](#regen.ecocredit.v1.ProjectInfo) | repeated | projects are the fetched projects. |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.ecocredit.v1.QuerySupplyRequest"></a>

### QuerySupplyRequest
QuerySupplyRequest is the Query/Supply request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of credit batch to query. |






<a name="regen.ecocredit.v1.QuerySupplyResponse"></a>

### QuerySupplyResponse
QuerySupplyResponse is the Query/Supply response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tradable_supply | [string](#string) |  | tradable_supply is the decimal number of tradable credits in the batch supply. |
| retired_supply | [string](#string) |  | retired_supply is the decimal number of retired credits in the batch supply. |
| cancelled_amount | [string](#string) |  | cancelled_amount is the decimal number of cancelled credits in the batch supply. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.ecocredit.v1.Query"></a>

### Query
Msg is the regen.ecocredit.v1 Query service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Classes | [QueryClassesRequest](#regen.ecocredit.v1.QueryClassesRequest) | [QueryClassesResponse](#regen.ecocredit.v1.QueryClassesResponse) | Classes queries for all credit classes with pagination. |
| ClassInfo | [QueryClassInfoRequest](#regen.ecocredit.v1.QueryClassInfoRequest) | [QueryClassInfoResponse](#regen.ecocredit.v1.QueryClassInfoResponse) | ClassInfo queries for information on a credit class. |
| ClassIssuers | [QueryClassIssuersRequest](#regen.ecocredit.v1.QueryClassIssuersRequest) | [QueryClassIssuersResponse](#regen.ecocredit.v1.QueryClassIssuersResponse) | ClassIssuers queries for the addresses of the issuers for a credit class. |
| Projects | [QueryProjectsRequest](#regen.ecocredit.v1.QueryProjectsRequest) | [QueryProjectsResponse](#regen.ecocredit.v1.QueryProjectsResponse) | Projects queries for all projects within a class with pagination. |
| ProjectInfo | [QueryProjectInfoRequest](#regen.ecocredit.v1.QueryProjectInfoRequest) | [QueryProjectInfoResponse](#regen.ecocredit.v1.QueryProjectInfoResponse) | ClassInfo queries for information on a project. |
| Batches | [QueryBatchesRequest](#regen.ecocredit.v1.QueryBatchesRequest) | [QueryBatchesResponse](#regen.ecocredit.v1.QueryBatchesResponse) | Batches queries for all batches in the given project with pagination. |
| BatchInfo | [QueryBatchInfoRequest](#regen.ecocredit.v1.QueryBatchInfoRequest) | [QueryBatchInfoResponse](#regen.ecocredit.v1.QueryBatchInfoResponse) | BatchInfo queries for information on a credit batch. |
| Balance | [QueryBalanceRequest](#regen.ecocredit.v1.QueryBalanceRequest) | [QueryBalanceResponse](#regen.ecocredit.v1.QueryBalanceResponse) | Balance queries the balance (both tradable and retired) of a given credit batch for a given account. |
| Supply | [QuerySupplyRequest](#regen.ecocredit.v1.QuerySupplyRequest) | [QuerySupplyResponse](#regen.ecocredit.v1.QuerySupplyResponse) | Supply queries the tradable and retired supply of a credit batch. |
| CreditTypes | [QueryCreditTypesRequest](#regen.ecocredit.v1.QueryCreditTypesRequest) | [QueryCreditTypesResponse](#regen.ecocredit.v1.QueryCreditTypesResponse) | CreditTypes returns the list of allowed types that credit classes can have. See Types/CreditType for more details. |
| Params | [QueryParamsRequest](#regen.ecocredit.v1.QueryParamsRequest) | [QueryParamsResponse](#regen.ecocredit.v1.QueryParamsResponse) | Params queries the ecocredit module parameters. |

 <!-- end services -->



<a name="regen/ecocredit/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/ecocredit/v1/tx.proto



<a name="regen.ecocredit.v1.MsgCancel"></a>

### MsgCancel
MsgCancel is the Msg/Cancel request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| holder | [string](#string) |  | holder is the credit holder address. |
| credits | [MsgCancel.CancelCredits](#regen.ecocredit.v1.MsgCancel.CancelCredits) | repeated | credits are the credits being cancelled. |






<a name="regen.ecocredit.v1.MsgCancel.CancelCredits"></a>

### MsgCancel.CancelCredits
CancelCredits specifies a batch and the number of credits being cancelled.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| amount | [string](#string) |  | amount is the number of credits being cancelled. Decimal values are acceptable within the precision returned by Query/Precision. |






<a name="regen.ecocredit.v1.MsgCancelResponse"></a>

### MsgCancelResponse
MsgCancelResponse is the Msg/Cancel response type.






<a name="regen.ecocredit.v1.MsgCreateBatch"></a>

### MsgCreateBatch
MsgCreateBatch is the Msg/CreateBatch request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| issuer | [string](#string) |  | issuer is the address of the batch issuer. |
| project_id | [string](#string) |  | project_id is the unique ID of the project this batch belongs to. |
| issuance | [MsgCreateBatch.BatchIssuance](#regen.ecocredit.v1.MsgCreateBatch.BatchIssuance) | repeated | issuance are the credits issued in the batch. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata attached to the credit batch. |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_date is the beginning of the period during which this credit batch was quantified and verified. |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | end_date is the end of the period during which this credit batch was quantified and verified. |






<a name="regen.ecocredit.v1.MsgCreateBatch.BatchIssuance"></a>

### MsgCreateBatch.BatchIssuance
BatchIssuance represents the issuance of some credits in a batch to a
single recipient.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recipient | [string](#string) |  | recipient is the account of the recipient. |
| tradable_amount | [string](#string) |  | tradable_amount is the number of credits in this issuance that can be traded by this recipient. Decimal values are acceptable. |
| retired_amount | [string](#string) |  | retired_amount is the number of credits in this issuance that are effectively retired by the issuer on receipt. Decimal values are acceptable. |
| retirement_location | [string](#string) |  | retirement_location is the location of the beneficiary or buyer of the retired credits. This must be provided if retired_amount is positive. It is a string of the form <country-code>[-<sub-national-code>[ <postal-code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters. |






<a name="regen.ecocredit.v1.MsgCreateBatchResponse"></a>

### MsgCreateBatchResponse
MsgCreateBatchResponse is the Msg/CreateBatch response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique denomination ID of the newly created batch. |






<a name="regen.ecocredit.v1.MsgCreateClass"></a>

### MsgCreateClass
MsgCreateClass is the Msg/CreateClass request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | admin is the address of the account that created the credit class. |
| issuers | [string](#string) | repeated | issuers are the account addresses of the approved issuers. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata to attached to the credit class. |
| credit_type_abbrev | [string](#string) |  | credit_type_abbrev describes the abbreviation of a credit type (e.g. "C", "BIO"). |






<a name="regen.ecocredit.v1.MsgCreateClassResponse"></a>

### MsgCreateClassResponse
MsgCreateClassResponse is the Msg/CreateClass response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [string](#string) |  | class_id is the unique ID of the newly created credit class. |






<a name="regen.ecocredit.v1.MsgCreateProject"></a>

### MsgCreateProject
MsgCreateProjectResponse is the Msg/CreateProject request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| issuer | [string](#string) |  | issuer is the address of an approved issuer for the credit class through which batches will be issued. It is not required, however, that this same issuer issue all batches for a project. |
| class_id | [string](#string) |  | class_id is the unique ID of the class within which the project is created. |
| metadata | [bytes](#bytes) |  | metadata is any arbitrary metadata attached to the project. |
| project_location | [string](#string) |  | project_location is the location of the project backing the credits in this batch. It is a string of the form <country-code>[-<sub-national-code>[ <postal-code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters. country-code is required, while sub-national-code and postal-code can be added for increasing precision. |
| project_id | [string](#string) |  | project_id is an optional user-specified project ID which can be used instead of an auto-generated ID. If project_id is provided, it must be unique within the credit class and match the regex [A-Za-z0-9]{2,16} or else the operation will fail. If project_id is omitted an ID will automatically be generated. |






<a name="regen.ecocredit.v1.MsgCreateProjectResponse"></a>

### MsgCreateProjectResponse
MsgCreateProjectResponse is the Msg/CreateProject response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [string](#string) |  | project_id is the ID of the newly created project. |






<a name="regen.ecocredit.v1.MsgRetire"></a>

### MsgRetire
MsgRetire is the Msg/Retire request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| holder | [string](#string) |  | holder is the credit holder address. |
| credits | [MsgRetire.RetireCredits](#regen.ecocredit.v1.MsgRetire.RetireCredits) | repeated | credits are the credits being retired. |
| location | [string](#string) |  | location is the location of the beneficiary or buyer of the retired credits. It is a string of the form <country-code>[-<sub-national-code>[ <postal-code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters. |






<a name="regen.ecocredit.v1.MsgRetire.RetireCredits"></a>

### MsgRetire.RetireCredits
RetireCredits specifies a batch and the number of credits being retired.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch_denom | [string](#string) |  | batch_denom is the unique ID of the credit batch. |
| amount | [string](#string) |  | amount is the number of credits being retired. Decimal values are acceptable within the precision returned by Query/Precision. |






<a name="regen.ecocredit.v1.MsgRetireResponse"></a>

### MsgRetireResponse
MsgRetire is the Msg/Retire response type.






<a name="regen.ecocredit.v1.MsgSend"></a>

### MsgSend
MsgSend is the Msg/Send request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the address of the account sending credits. |
| recipient | [string](#string) |  | sender is the address of the account receiving credits. |
| credits | [MsgSend.SendCredits](#regen.ecocredit.v1.MsgSend.SendCredits) | repeated | credits are the credits being sent. |






<a name="regen.ecocredit.v1.MsgSend.SendCredits"></a>

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






<a name="regen.ecocredit.v1.MsgSendResponse"></a>

### MsgSendResponse
MsgSendResponse is the Msg/Send response type.






<a name="regen.ecocredit.v1.MsgUpdateClassAdmin"></a>

### MsgUpdateClassAdmin
MsgUpdateClassAdmin is the Msg/UpdateClassAdmin request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | admin is the address of the account that is the admin of the credit class. |
| class_id | [string](#string) |  | class_id is the unique ID of the credit class. |
| new_admin | [string](#string) |  | new_admin is the address of the new admin of the credit class. |






<a name="regen.ecocredit.v1.MsgUpdateClassAdminResponse"></a>

### MsgUpdateClassAdminResponse
MsgUpdateClassAdminResponse is the MsgUpdateClassAdmin response type.






<a name="regen.ecocredit.v1.MsgUpdateClassIssuers"></a>

### MsgUpdateClassIssuers
MsgUpdateClassIssuers is the Msg/UpdateClassIssuers request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | admin is the address of the account that is the admin of the credit class. |
| class_id | [string](#string) |  | class_id is the unique ID of the credit class. |
| add_issuers | [string](#string) | repeated | add_issuers are the issuers to add to the class issuers list. |
| remove_issuers | [string](#string) | repeated | remove_issuers are the issuers to remove from the class issuers list. |






<a name="regen.ecocredit.v1.MsgUpdateClassIssuersResponse"></a>

### MsgUpdateClassIssuersResponse
MsgUpdateClassIssuersResponse is the MsgUpdateClassIssuers response type.






<a name="regen.ecocredit.v1.MsgUpdateClassMetadata"></a>

### MsgUpdateClassMetadata
MsgUpdateClassMetadata is the Msg/UpdateClassMetadata request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | admin is the address of the account that is the admin of the credit class. |
| class_id | [string](#string) |  | class_id is the unique ID of the credit class. |
| metadata | [bytes](#bytes) |  | metadata is the updated arbitrary metadata to be attached to the credit class. |






<a name="regen.ecocredit.v1.MsgUpdateClassMetadataResponse"></a>

### MsgUpdateClassMetadataResponse
MsgUpdateClassMetadataResponse is the MsgUpdateClassMetadata response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.ecocredit.v1.Msg"></a>

### Msg
Msg is the regen.ecocredit.v1 Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateClass | [MsgCreateClass](#regen.ecocredit.v1.MsgCreateClass) | [MsgCreateClassResponse](#regen.ecocredit.v1.MsgCreateClassResponse) | CreateClass creates a new credit class with an approved list of issuers and optional metadata. |
| CreateProject | [MsgCreateProject](#regen.ecocredit.v1.MsgCreateProject) | [MsgCreateProjectResponse](#regen.ecocredit.v1.MsgCreateProjectResponse) | CreateProject creates a new project within a credit class. |
| CreateBatch | [MsgCreateBatch](#regen.ecocredit.v1.MsgCreateBatch) | [MsgCreateBatchResponse](#regen.ecocredit.v1.MsgCreateBatchResponse) | CreateBatch creates a new batch of credits for an existing project. This will create a new batch denom with a fixed supply. Issued credits can be distributed to recipients in either tradable or retired form. |
| Send | [MsgSend](#regen.ecocredit.v1.MsgSend) | [MsgSendResponse](#regen.ecocredit.v1.MsgSendResponse) | Send sends tradable credits from one account to another account. Sent credits can either be tradable or retired on receipt. |
| Retire | [MsgRetire](#regen.ecocredit.v1.MsgRetire) | [MsgRetireResponse](#regen.ecocredit.v1.MsgRetireResponse) | Retire retires a specified number of credits in the holder's account. |
| Cancel | [MsgCancel](#regen.ecocredit.v1.MsgCancel) | [MsgCancelResponse](#regen.ecocredit.v1.MsgCancelResponse) | Cancel removes a number of credits from the holder's account and also deducts them from the tradable supply, effectively cancelling their issuance on Regen Ledger |
| UpdateClassAdmin | [MsgUpdateClassAdmin](#regen.ecocredit.v1.MsgUpdateClassAdmin) | [MsgUpdateClassAdminResponse](#regen.ecocredit.v1.MsgUpdateClassAdminResponse) | UpdateClassAdmin updates the credit class admin |
| UpdateClassIssuers | [MsgUpdateClassIssuers](#regen.ecocredit.v1.MsgUpdateClassIssuers) | [MsgUpdateClassIssuersResponse](#regen.ecocredit.v1.MsgUpdateClassIssuersResponse) | UpdateClassIssuers updates the credit class issuer list |
| UpdateClassMetadata | [MsgUpdateClassMetadata](#regen.ecocredit.v1.MsgUpdateClassMetadata) | [MsgUpdateClassMetadataResponse](#regen.ecocredit.v1.MsgUpdateClassMetadataResponse) | UpdateClassMetadata updates the credit class metadata |

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

