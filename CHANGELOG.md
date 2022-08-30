# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### General

#### Changed

- [#1244](https://github.com/regen-network/regen-ledger/pull/1244) Update all modules to Cosmos SDK v0.46

### app

#### Added

- [#1340](https://github.com/regen-network/regen-ledger/pull/1340) Add Cosmos SDK group module to app configuration

#### Changed

- [#1350](https://github.com/regen-network/regen-ledger/pull/1350) Move application entry point to root directory
- [#1357](https://github.com/regen-network/regen-ledger/pull/1350) Migrate from custom module manager to sdk module manager

#### Removed

- [#1258](https://github.com/regen-network/regen-ledger/pull/1258) Remove group module from experimental config
- [#1350](https://github.com/regen-network/regen-ledger/pull/1350) Remove experimental app configuration
- [#1357](https://github.com/regen-network/regen-ledger/pull/1357) Remove unused RegenApp functions for setting custom keepers

### types

#### Removed

- [#1357](https://github.com/regen-network/regen-ledger/pull/1357) Remove custom context, module manager, module keys, cli, rest, base module types

#### Changed

- [#1357](https://github.com/regen-network/regen-ledger/pull/1357) Refactor fixture factory to use baseApp routing and Cosmos SDK module manager
- [#1357](https://github.com/regen-network/regen-ledger/pull/1357) Rename FixtureFactory to Factory to prevent package stuttering


### x/data

#### Added

- [#1351](https://github.com/regen-network/regen-ledger/pull/1351) Add `signer` option to transaction messages
- [#1395](https://github.com/regen-network/regen-ledger/pull/1395) Add `DataId` state validation checks
- [#1395](https://github.com/regen-network/regen-ledger/pull/1395) Add `DataAnchor` state validation checks
- [#1395](https://github.com/regen-network/regen-ledger/pull/1395) Add `DataAttestor` state validation checks
- [#1395](https://github.com/regen-network/regen-ledger/pull/1395) Add `Resolver` state validation checks
- [#1395](https://github.com/regen-network/regen-ledger/pull/1395) Add `DataResolver` state validation checks

### x/ecocredit

#### API Breaking Changes

- [#1337](https://github.com/regen-network/regen-ledger/pull/1337) The `NewKeeper` method in `ecocredit/core` requires an `authority` address.
- [#1337](https://github.com/regen-network/regen-ledger/pull/1337) The `AddCreditType` proposal handler has been removed.
- [#1342](https://github.com/regen-network/regen-ledger/pull/1342) The `NewKeeper` method in `ecocredit/marketplace` requires an `authority` address.
- [#1342](https://github.com/regen-network/regen-ledger/pull/1342) The `AllowedDenom` proposal handler has been removed.


#### Added

- [#1337](https://github.com/regen-network/regen-ledger/pull/1342) Add `AddAllowedDenom` msg-based gov proposal
- [#1337](https://github.com/regen-network/regen-ledger/pull/1337) Add `AddCreditType` msg-based gov proposal
- [#1346](https://github.com/regen-network/regen-ledger/pull/1346) Add `RemoveAllowedDenom` msg-based gov proposal
- [#1349](https://github.com/regen-network/regen-ledger/pull/1349) Add `UpdateBasketFees` msg-based gov proposal
- [#1351](https://github.com/regen-network/regen-ledger/pull/1351) Add `signer` option to transaction messages
- [#1362](https://github.com/regen-network/regen-ledger/pull/1362) Add `BatchBalance` state validation checks
- [#1362](https://github.com/regen-network/regen-ledger/pull/1362) Add `BatchContract` state validation checks
- [#1362](https://github.com/regen-network/regen-ledger/pull/1362) Add `BatchSequence` state validation checks
- [#1362](https://github.com/regen-network/regen-ledger/pull/1362) Add `BatchSupply` state validation checks
- [#1362](https://github.com/regen-network/regen-ledger/pull/1362) Add `ClassSequence` state validation checks
- [#1362](https://github.com/regen-network/regen-ledger/pull/1362) Add `OriginTxIndex` state validation checks
- [#1362](https://github.com/regen-network/regen-ledger/pull/1362) Add `ProjectSequence` state validation checks
- [#1354](https://github.com/regen-network/regen-ledger/pull/1354) Add `AddClassCreator` msg-based gov proposal
- [#1354](https://github.com/regen-network/regen-ledger/pull/1354) Add `RemoveClassCreator` msg-based gov proposal
- [#1354](https://github.com/regen-network/regen-ledger/pull/1354) Add `ToggleCreditClassAllowlist` msg-based gov proposal
- [#1354](https://github.com/regen-network/regen-ledger/pull/1354) Add `UpdateClassFees` msg-based gov proposal
- [#1391](https://github.com/regen-network/regen-ledger/pull/1391) Add `BasketFees` params query
- [#1412](https://github.com/regen-network/regen-ledger/pull/1412) Add `EventRemoveAllowedDenom`
- [#1416](https://github.com/regen-network/regen-ledger/pull/1416) Add `QueryAllBalances` query
- [#1417](https://github.com/regen-network/regen-ledger/pull/1417) Add ecocredit params state migration

#### Changed

- [#1362](https://github.com/regen-network/regen-ledger/pull/1362) Update `Batch` state validation checks
- [#1362](https://github.com/regen-network/regen-ledger/pull/1362) Update `Class` state validation checks
- [#1362](https://github.com/regen-network/regen-ledger/pull/1362) Update `ClassIssuer` state validation checks
- [#1362](https://github.com/regen-network/regen-ledger/pull/1362) Update `CreditType` state validation checks
- [#1362](https://github.com/regen-network/regen-ledger/pull/1362) Update `Project` state validation checks
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `DateCriteria.ToApi` to `DateCriteria.ToAPI`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `QueryProjectsByReferenceIdCmd` to `QueryProjectsByReferenceIDCmd`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `FlagReferenceId` to `FlagReferenceID`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `MaxReferenceIdLength` to `MaxReferenceIDLength`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `RegexClassId` to `RegexClassID`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `RegexProjectId` to `RegexProjectID`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `FormatClassId` to `FormatClassID`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `FormatProjectId` to `FormatProjectID`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `ValidateClassId` to `ValidateClassID`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `ValidateProjectId` to `ValidateProjectID`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `GetClassIDFromProjectId` to `GetClassIDFromProjectID`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `GetClassIdFromBatchDenom` to `GetClassIDFromBatchDenom`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `GetProjectIdFromBatchDenom` to `GetProjectIDFromBatchDenom`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `GetCreditTypeAbbrevFromClassId` to `GetCreditTypeAbbrevFromClassID`
- [#1384](https://github.com/regen-network/regen-ledger/pull/1384) Renamed `BasketSupplyInvariant` to `SupplyInvariant`

### Fixed

- [#1411](https://github.com/regen-network/regen-ledger/pull/1411/) `EventCreateBatch`, `EventMintBatchCredits`, and and `EventSend` now show "0" instead of "" for 0 value decimals.

#### Removed

- [#1337](https://github.com/regen-network/regen-ledger/pull/1337) Remove `AddCreditType` proposal handler
- [#1342](https://github.com/regen-network/regen-ledger/pull/1342) Remove `AllowedDenom` proposal handler
- [#1354](https://github.com/regen-network/regen-ledger/pull/1354) Removed `paramsKeeper` parameter from `core/Keeper`.

### x/group

#### Removed

- [#1258](https://github.com/regen-network/regen-ledger/pull/1258) Remove group module in favor of Cosmos SDK group module

## [v4.0.1](https://github.com/regen-network/regen-ledger/releases/tag/v4.0.1) - TBD

### x/ecocredit

#### Fixed

- [#1360](https://github.com/regen-network/regen-ledger/pull/1360) Register ecocredit v1alpha1 messages to allow for historical queries

## [v4.0.0](https://github.com/regen-network/regen-ledger/releases/tag/v4.0.0) - 2022-07-26

### General

#### Added

- [#1255](https://github.com/regen-network/regen-ledger/pull/1255) Add arm and windows package builds to release build

#### Changed

- [#1088](https://github.com/regen-network/regen-ledger/pull/1088) Update all modules to Go `1.18`
- [#1131](https://github.com/regen-network/regen-ledger/pull/1131) Update app module version to `v4`
- [#1097](https://github.com/regen-network/regen-ledger/pull/1097) Update `make` commands and build options

#### Fixed

- [#685](https://github.com/regen-network/regen-ledger/pull/685) Fix swagger-gen to include ibc-go swagger docs

### types

#### Added

- [#720](https://github.com/regen-network/regen-ledger/pull/720) Add begin and end blocker support to module server
- [#783](https://github.com/regen-network/regen-ledger/pull/783) Add `BigInt` conversion method to `Dec` interface 

#### Fixed

- [#1210](https://github.com/regen-network/regen-ledger/pull/1210) Fix `PositiveDecimalFromString` to error if not finite
- [#1292](https://github.com/regen-network/regen-ledger/pull/1292) Fix nil timestamp and duration conversion

### x/data

#### Added

- [#708](https://github.com/regen-network/regen-ledger/pull/708) Add resolver message and query service
- [#866](https://github.com/regen-network/regen-ledger/pull/866) Add resolver commands for message and query service
- [#887](https://github.com/regen-network/regen-ledger/pull/887) Add data module to stable app configuration
- [#887](https://github.com/regen-network/regen-ledger/pull/887) Add `ByHash` and IRI/ContentHash conversion queries
- [#953](https://github.com/regen-network/regen-ledger/pull/953) Add legacy Amino signing support
- [#1051](https://github.com/regen-network/regen-ledger/pull/1051) Add `EventDefineResolver` and `EventRegisterResolver`
- [#1107](https://github.com/regen-network/regen-ledger/pull/1107) Add commands and endpoints for IRI/ContentHash conversion
- [#1132](https://github.com/regen-network/regen-ledger/pull/1132) Add `Query/ResolversByURL`
- [#1156](https://github.com/regen-network/regen-ledger/pull/1156) Add post endpoints to all `ByHash` queries  
- [#1219](https://github.com/regen-network/regen-ledger/pull/1219) Add unique constraints on resolver url and manager

#### Changed

- [#848](https://github.com/regen-network/regen-ledger/pull/848) Update protobuf package version to `v1`
- [#927](https://github.com/regen-network/regen-ledger/pull/927) Update `Msg/Sign` to `Msg/Attest`
- [#956](https://github.com/regen-network/regen-ledger/pull/956) Update data module to use `cosmos-sdk/orm`
- [#970](https://github.com/regen-network/regen-ledger/pull/970) Update data module to use `cosmos-sdk/orm`
- [#969](https://github.com/regen-network/regen-ledger/pull/971) Update `Msg/Attest` to single attestor and multiple pieces of data
- [#1014](https://github.com/regen-network/regen-ledger/pull/1014) Update `hash` field names to `content_hash` 
- [#1156](https://github.com/regen-network/regen-ledger/pull/1156) Update `Query/ByIRI` to `Query/AnchorByIRI` 
- [#1156](https://github.com/regen-network/regen-ledger/pull/1156) Update `Query/ByHash` to `Query/AnchorByHash` 
- [#1156](https://github.com/regen-network/regen-ledger/pull/1156) Update `Query/AttestorsByIRI` to `Query/AttestationsByIRI` 
- [#1156](https://github.com/regen-network/regen-ledger/pull/1156) Update `Query/AttestorsByHash` to `Query/AttestationsByHash` 
- [#1156](https://github.com/regen-network/regen-ledger/pull/1156) Update `ContentEntry` to `AnchorInfo` 
- [#1156](https://github.com/regen-network/regen-ledger/pull/1156) Update `AttestorEntry` to `AttestorInfo` 
- [#1156](https://github.com/regen-network/regen-ledger/pull/1156) Update `Query/IRIByHash` to `Query/ConvertHashToIRI` 
- [#1156](https://github.com/regen-network/regen-ledger/pull/1156) Update `Query/HashByIRI` to `Query/ConvertIRIToHash` 

#### Fixed

- [#939](https://github.com/regen-network/regen-ledger/pull/934) Fix gas consumption to consume for each iteration
- [#946](https://github.com/regen-network/regen-ledger/pull/946) Remove use of `oneof` to support legacy Amino signing
- [#1128](https://github.com/regen-network/regen-ledger/pull/1128) Remove unique constrains on resolver URL

### x/ecocredit

#### Added

- [#823](https://github.com/regen-network/regen-ledger/pull/823) Add `Msg/CreateProject`
- [#840](https://github.com/regen-network/regen-ledger/pull/840) Add `Query/Projects`
- [#841](https://github.com/regen-network/regen-ledger/pull/841) Add `Query/ProjectInfo`
- [#870](https://github.com/regen-network/regen-ledger/pull/870) Add `curator` to `Basket`
- [#873](https://github.com/regen-network/regen-ledger/pull/847) Add marketplace `Msg/Sell`
- [#874](https://github.com/regen-network/regen-ledger/pull/874) Add `years_in_the_past` basket `DataCriteria`
- [#878](https://github.com/regen-network/regen-ledger/pull/878) Add `escrowed` to `BatchBalance`
- [#878](https://github.com/regen-network/regen-ledger/pull/878) Add escrow functionality to sell orders
- [#885](https://github.com/regen-network/regen-ledger/pull/885) Add `Query/BatchesByClass`
- [#888](https://github.com/regen-network/regen-ledger/pull/888) Add `admin` to `Project`
- [#855](https://github.com/regen-network/regen-ledger/pull/855) Add `EventUpdateClassAdmin`
- [#855](https://github.com/regen-network/regen-ledger/pull/855) Add `EventUpdateClassIssuers`
- [#855](https://github.com/regen-network/regen-ledger/pull/855) Add `EventUpdateClassMetadata`
- [#891](https://github.com/regen-network/regen-ledger/pull/891) Add marketplace `Msg/CancelSellOrder`
- [#891](https://github.com/regen-network/regen-ledger/pull/891) Add marketplace `Msg/UpdateSellOrders`
- [#906](https://github.com/regen-network/regen-ledger/pull/906) Add marketplace `Query/SellOrder`
- [#912](https://github.com/regen-network/regen-ledger/pull/912) Add `fee` to `Msg/CreateClass`
- [#917](https://github.com/regen-network/regen-ledger/pull/917) Add marketplace `Query/SellOrders`
- [#936](https://github.com/regen-network/regen-ledger/pull/936) Add `issuance_date` to `Batch`
- [#937](https://github.com/regen-network/regen-ledger/pull/937) Add `Msg/MintBatchCredits` api
- [#937](https://github.com/regen-network/regen-ledger/pull/937) Add `Msg/SealBatch` api
- [#942](https://github.com/regen-network/regen-ledger/pull/942) Add `Query/Balances`
- [#899](https://github.com/regen-network/regen-ledger/pull/899) Add marketplace `Msg/BuyDirect`
- [#954](https://github.com/regen-network/regen-ledger/pull/954) Add `Query/ClassesByAdmin`
- [#955](https://github.com/regen-network/regen-ledger/pull/955) Add `Query/ClassesByIssuer`
- [#957](https://github.com/regen-network/regen-ledger/pull/957) Add state migrations for `v4.0`
- [#969](https://github.com/regen-network/regen-ledger/pull/969) Add marketplace `AllowAskDenom`
- [#991](https://github.com/regen-network/regen-ledger/pull/991) Add `Msg/MintBatchCredits` message validation
- [#991](https://github.com/regen-network/regen-ledger/pull/991) Add `Msg/SealBatch` message validation
- [#1009](https://github.com/regen-network/regen-ledger/pull/1009) Add `Msg/UpdateProjectAdmin` message validation
- [#1009](https://github.com/regen-network/regen-ledger/pull/1009) Add `Msg/UpdateProjectMetadata` message validation
- [#1010](https://github.com/regen-network/regen-ledger/pull/1010) Add `Msg/UpdateProjectAdmin` server implementation
- [#1010](https://github.com/regen-network/regen-ledger/pull/1010) Add `Msg/UpdateProjectMetadata` server implementation
- [#1015](https://github.com/regen-network/regen-ledger/pull/1015) Add `Msg/CreditTypeProposal`
- [#1037](https://github.com/regen-network/regen-ledger/pull/1037) Add `EventUpdateProjectAdmin`
- [#1037](https://github.com/regen-network/regen-ledger/pull/1037) Add `EventUpdateProjectMetadata`
- [#1008](https://github.com/regen-network/regen-ledger/pull/1008) Add `buy-direct` and `buy-direct-batch` CLI commands
- [#1033](https://github.com/regen-network/regen-ledger/pull/1033) Add `class-issuers` CLI command
- [#1054](https://github.com/regen-network/regen-ledger/pull/1054) Add `basket_info` to `Query/BasketResponse`
- [#1059](https://github.com/regen-network/regen-ledger/pull/1059) Add `Msg/MintBatchCredits` server implementation
- [#1060](https://github.com/regen-network/regen-ledger/pull/1060) Add `Msg/SealBatch` server implementation
- [#1061](https://github.com/regen-network/regen-ledger/pull/1061) Add `CreditType` state migration
- [#1072](https://github.com/regen-network/regen-ledger/pull/1072) Add `Msg/AllowDenomProposal`
- [#1064](https://github.com/regen-network/regen-ledger/pull/1064) Add `Get` method for a single parameter
- [#1094](https://github.com/regen-network/regen-ledger/pull/1094) Add `reference_id` to `Project`
- [#1095](https://github.com/regen-network/regen-ledger/pull/1095) Add batch denom migration to state migrations
- [#1099](https://github.com/regen-network/regen-ledger/pull/1099) Add marketplace `Query/AllowedDenoms`
- [#1101](https://github.com/regen-network/regen-ledger/pull/1101) Add `Msg/Bridge`
- [#1112](https://github.com/regen-network/regen-ledger/pull/1112) Add duplicate `BatchOriginTx` check
- [#1116](https://github.com/regen-network/regen-ledger/pull/1116) Add `Query/BatchesByProject`
- [#1116](https://github.com/regen-network/regen-ledger/pull/1116) Add commands for querying batches
- [#1121](https://github.com/regen-network/regen-ledger/pull/1121) Add `v4.0.0` upgrade handler
- [#1125](https://github.com/regen-network/regen-ledger/pull/1125) Add `EventMint`
- [#1148](https://github.com/regen-network/regen-ledger/pull/1148) Add marketplace `EventBuyDirect`
- [#1148](https://github.com/regen-network/regen-ledger/pull/1148) Add marketplace `EventCancel`
- [#1141](https://github.com/regen-network/regen-ledger/pull/1141) Add `Query/ProjectsByAdmin`
- [#1160](https://github.com/regen-network/regen-ledger/pull/1160) Add reusable `Credits` type.
- [#1168](https://github.com/regen-network/regen-ledger/pull/1168) Add additional bindings for query endpoints
- [#1178](https://github.com/regen-network/regen-ledger/pull/1178) Add `Query/Projects`
- [#1174](https://github.com/regen-network/regen-ledger/pull/1174) Add `Msg/BridgeReceive`
- [#1197](https://github.com/regen-network/regen-ledger/pull/1197) Add alternative bindings for marketplace queries
- [#1198](https://github.com/regen-network/regen-ledger/pull/1198) Add alternative bindings for basket queries
- [#1205](https://github.com/regen-network/regen-ledger/pull/1205) Add basket denom units migration
- [#1209](https://github.com/regen-network/regen-ledger/pull/1209) Add regen metadata units migration
- [#1213](https://github.com/regen-network/regen-ledger/pull/1213) Add simpler version of `send` command
- [#1218](https://github.com/regen-network/regen-ledger/pull/1218) Add restriction on buyer being seller
- [#1220](https://github.com/regen-network/regen-ledger/pull/1220) Add project reference id to `Query/Project` response
- [#1224](https://github.com/regen-network/regen-ledger/pull/1224) Add `EventBridgeReceive` and `EventBridge`
- [#1225](https://github.com/regen-network/regen-ledger/pull/1225) Add `note` and `contract` to `OriginTx`
- [#1226](https://github.com/regen-network/regen-ledger/pull/1226) Add `BatchContract` to map batch to contract
- [#1229](https://github.com/regen-network/regen-ledger/pull/1229) Add `class_key` to `OriginTxIndex`
- [#1274](https://github.com/regen-network/regen-ledger/pull/1274) Add `cancel-sell-order` command

#### Changed

- [#804](https://github.com/regen-network/regen-ledger/pull/804) Upgrade module to use `cosmos-sdk/orm`
- [#816](https://github.com/regen-network/regen-ledger/pull/823) Update `Msg/CreateClass` to use `cosmos-sdk/orm`
- [#824](https://github.com/regen-network/regen-ledger/pull/823) Update `Msg/CreateBatch` to use `cosmos-sdk/orm`
- [#825](https://github.com/regen-network/regen-ledger/pull/825) Update `Msg/Send` to use `cosmos-sdk/orm`
- [#826](https://github.com/regen-network/regen-ledger/pull/826) Update `Msg/Retire` to use `cosmos-sdk/orm`
- [#828](https://github.com/regen-network/regen-ledger/pull/828) Update `Msg/Cancel` to use `cosmos-sdk/orm`
- [#830](https://github.com/regen-network/regen-ledger/pull/830) Update `Msg/UpdateClassAdmin` to use `cosmos-sdk/orm`
- [#830](https://github.com/regen-network/regen-ledger/pull/830) Update `Msg/UpdateClassIssuers` to use `cosmos-sdk/orm`
- [#830](https://github.com/regen-network/regen-ledger/pull/830) Update `Msg/UpdateClassMetadata` to use `cosmos-sdk/orm`
- [#837](https://github.com/regen-network/regen-ledger/pull/837) Update `Query/Classes` to use `cosmos-sdk/orm`
- [#838](https://github.com/regen-network/regen-ledger/pull/838) Update `Query/ClassInfo` to use `cosmos-sdk/orm`
- [#839](https://github.com/regen-network/regen-ledger/pull/839) Update `Query/ClassIssuers` to use `cosmos-sdk/orm`
- [#842](https://github.com/regen-network/regen-ledger/pull/842) Update `Query/Batches` to use `cosmos-sdk/orm`
- [#843](https://github.com/regen-network/regen-ledger/pull/843) Update `Query/BatchInfo` to use `cosmos-sdk/orm`
- [#844](https://github.com/regen-network/regen-ledger/pull/844) Update `Query/Balance` to use `cosmos-sdk/orm`
- [#845](https://github.com/regen-network/regen-ledger/pull/845) Update `Query/Supply` to use `cosmos-sdk/orm`
- [#846](https://github.com/regen-network/regen-ledger/pull/846) Update `Query/Params` to use `cosmos-sdk/orm`
- [#847](https://github.com/regen-network/regen-ledger/pull/847) Update `Query/CreditTypes` to use `cosmos-sdk/orm`
- [#848](https://github.com/regen-network/regen-ledger/pull/848) Update protobuf package versions to `v1`
- [#863](https://github.com/regen-network/regen-ledger/pull/863) Update `metadata` fields from bytes to strings
- [#901](https://github.com/regen-network/regen-ledger/pull/901) Update client commands to use `v1` api
- [#913](https://github.com/regen-network/regen-ledger/pull/913) Update module wiring to use submodules
- [#935](https://github.com/regen-network/regen-ledger/pull/935) Update module wiring to use core submodule
- [#977](https://github.com/regen-network/regen-ledger/pull/977) Update genesis import/export to use `v1` api
- [#1017](https://github.com/regen-network/regen-ledger/pull/1017) Update invariants to use `v1` api
- [#1020](https://github.com/regen-network/regen-ledger/pull/1020) Update project and retirement `location` to `jurisdiction`
- [#1021](https://github.com/regen-network/regen-ledger/pull/1021) Update `id` to `key` and `<object>_id` to `id`
- [#1021](https://github.com/regen-network/regen-ledger/pull/1021) Update `credit_type` to `credit_type_abbrev`
- [#1022](https://github.com/regen-network/regen-ledger/pull/1022) Update `Query/ClassResponse` fields
- [#1022](https://github.com/regen-network/regen-ledger/pull/1022) Update `Query/BatchResponse` fields
- [#1022](https://github.com/regen-network/regen-ledger/pull/1022) Update `Query/ProjectResponse` fields
- [#1039](https://github.com/regen-network/regen-ledger/pull/1039) Update `BatchDenom` to `Denom`
- [#1040](https://github.com/regen-network/regen-ledger/pull/1040) Update `ClassInfo` to `Class`
- [#1040](https://github.com/regen-network/regen-ledger/pull/1040) Update `ProjectInfo` to `Project`
- [#1040](https://github.com/regen-network/regen-ledger/pull/1040) Update `BatchInfo` to `Batch`
- [#1052](https://github.com/regen-network/regen-ledger/pull/1052) Update `Project` fields to exclude `project_` prefix
- [#1054](https://github.com/regen-network/regen-ledger/pull/1054) Update `Query/SellOrdersResponse` fields
- [#1054](https://github.com/regen-network/regen-ledger/pull/1054) Update `Query/AllowedDenomsResponse` fields
- [#1046](https://github.com/regen-network/regen-ledger/pull/1046) Update project ID and batch denom format
- [#1090](https://github.com/regen-network/regen-ledger/pull/1090) Update query names to exclude `Info`
- [#1123](https://github.com/regen-network/regen-ledger/pull/1123) Rename `AskPrice` to `AskAmount`
- [#1123](https://github.com/regen-network/regen-ledger/pull/1123) Rename `BatchId` to `BatchKey`
- [#1123](https://github.com/regen-network/regen-ledger/pull/1123) Rename `CreditType` to `CreditTypeAbbrev`
- [#1123](https://github.com/regen-network/regen-ledger/pull/1123) Rename `Tradable` to `TradableAmount`
- [#1123](https://github.com/regen-network/regen-ledger/pull/1123) Rename `Retired` to `RetiredAmount`
- [#1123](https://github.com/regen-network/regen-ledger/pull/1123) Rename `Escrowed` to `EscrowedAmount`
- [#1123](https://github.com/regen-network/regen-ledger/pull/1123) Rename `AskPrice` to `AskAmount`
- [#1153](https://github.com/regen-network/regen-ledger/pull/1153) Rename marketplace `owner` to `seller`
- [#1153](https://github.com/regen-network/regen-ledger/pull/1153) Rename `account` to `address`
- [#1153](https://github.com/regen-network/regen-ledger/pull/1153) Rename `holder` to `owner`
- [#1167](https://github.com/regen-network/regen-ledger/pull/1167) Update basket fee to be burned upon basket creation
- [#1168](https://github.com/regen-network/regen-ledger/pull/1168) Update `Query/Projects` to `Query/ProjectsByClass`
- [#1160](https://github.com/regen-network/regen-ledger/pull/1160) Update `Msg/Bridge` to use reusable `Credits` type
- [#1160](https://github.com/regen-network/regen-ledger/pull/1160) Update `Msg/Cancel` to use reusable `Credits` type
- [#1161](https://github.com/regen-network/regen-ledger/pull/1161) Update `Msg/Retire` to use reusable `Credits` type
- [#1161](https://github.com/regen-network/regen-ledger/pull/1161) Rename `issuer` to `admin` in `Msg/CreateProject`
- [#1161](https://github.com/regen-network/regen-ledger/pull/1161) Rename `metadata` to `new_metadata` in `Msg/UpdateClassMetadata`
- [#1192](https://github.com/regen-network/regen-ledger/pull/1192) Rename `tradable_supply` to `tradable_amount`
- [#1192](https://github.com/regen-network/regen-ledger/pull/1192) Rename `retired_supply` to `retired_amount`
- [#1197](https://github.com/regen-network/regen-ledger/pull/1197) Rename `SellOrdersByBatchDenom` to `SellOrdersByBatch`
- [#1199](https://github.com/regen-network/regen-ledger/pull/1199) Update ecocredit commands to consistently use json
- [#1200](https://github.com/regen-network/regen-ledger/pull/1200) Update basic validation for basket denom
- [#1213](https://github.com/regen-network/regen-ledger/pull/1213) Rename `send` command to `send-bulk`
- [#1288](https://github.com/regen-network/regen-ledger/pull/1288) Rename `types` query command to `credit-types`
- [#1288](https://github.com/regen-network/regen-ledger/pull/1288) Rename `balance` query command to `batch-balance`
- [#1288](https://github.com/regen-network/regen-ledger/pull/1288) Rename `supply` query command to `batch-supply`

#### Removed

- [#1043](https://github.com/regen-network/regen-ledger/pull/1043) Remove `credit_types` from `Params`
- [#1044](https://github.com/regen-network/regen-ledger/pull/1044) Remove unnecessary fields in events
- [#1086](https://github.com/regen-network/regen-ledger/pull/1086) Remove unnecessary fields in events

#### Fixed

- [#875](https://github.com/regen-network/regen-ledger/pull/875) Fix `put-in-basket` flags and error messages
- [#939](https://github.com/regen-network/regen-ledger/pull/934) Fix gas consumption for each iteration
- [#1135](https://github.com/regen-network/regen-ledger/pull/1135) Deprecate basket exponent and use credit type precision
- [#1144](https://github.com/regen-network/regen-ledger/pull/1144) Fix `MsgCreateBatch` issuance with same recipient
- [#1177](https://github.com/regen-network/regen-ledger/pull/1177) Remove exponent from `create-basket` command
- [#1180](https://github.com/regen-network/regen-ledger/pull/1180) Fix batch dates and project IDs in `regen-1` state migrations
- [#1194](https://github.com/regen-network/regen-ledger/pull/1194) Fix start and end date in batch denom to always use UTC
- [#1195](https://github.com/regen-network/regen-ledger/pull/1195) Fix genesis validation to include basket balances
- [#1202](https://github.com/regen-network/regen-ledger/pull/1202) Fix order of storing basket token denominations
- [#1234](https://github.com/regen-network/regen-ledger/pull/1234) Fix `CreateClassFee` implementation if not required
- [#1290](https://github.com/regen-network/regen-ledger/pull/1290) Fix error message when basket not found
- [#1290](https://github.com/regen-network/regen-ledger/pull/1290) Return empty balance when basket balance not found

### orm

- [#1076](https://github.com/regen-network/regen-ledger/pull/1076) Remove `orm` module. The final version of the `orm` module released from regen-ledger is [orm/v1.0.0-beta1](https://github.com/regen-network/regen-ledger/releases/tag/orm%2Fv1.0.0-beta1). The data module and ecocredit module have migrated to [cosmos-sdk/orm@v1.0.0-alpha.11](https://github.com/cosmos/cosmos-sdk/tree/orm/v1.0.0-alpha.11/orm) while the alpha version of the group module remains unchanged since the last release and continues to use [orm/v1.0.0-beta1](https://github.com/regen-network/regen-ledger/releases/tag/orm%2Fv1.0.0-beta1).

## [v3.0.0](https://github.com/regen-network/regen-ledger/releases/tag/v3.0.0) - 2022-02-25

### x/ecocredit

#### Added

- [#737](https://github.com/regen-network/regen-ledger/pull/737) Add basket proto definitions.
- [#735](https://github.com/regen-network/regen-ledger/pull/735) Add minimal baskets ORM and keeper setup.
- [#747](https://github.com/regen-network/regen-ledger/pull/747) Add sdk.Msg implementation for MsgPut. 
- [#748](https://github.com/regen-network/regen-ledger/pull/748) Add sdk.Msg implementation for MsgTake. 
- [#745](https://github.com/regen-network/regen-ledger/pull/745) Add sdk.Msg implementation for MsgCreate. 
- [#751](https://github.com/regen-network/regen-ledger/pull/751) Add BasketBalance query server method. 
- [#757](https://github.com/regen-network/regen-ledger/pull/757) Add start date window for date criteria. 
- [#760](https://github.com/regen-network/regen-ledger/pull/760) Add BasketBalance query CLI commands. 
- [#735](https://github.com/regen-network/regen-ledger/pull/735) Add Basket query server method. 
- [#758](https://github.com/regen-network/regen-ledger/pull/758) Add Put message server method. 
- [#746](https://github.com/regen-network/regen-ledger/pull/746) Add Take message server method. 
- [#763](https://github.com/regen-network/regen-ledger/pull/763) Add BasketBalances query server method. 
- [#766](https://github.com/regen-network/regen-ledger/pull/766) Add Basket query CLI command.
- [#749](https://github.com/regen-network/regen-ledger/pull/749) Add Take transaction CLI command.
- [#766](https://github.com/regen-network/regen-ledger/pull/766) Add Baskets query CLI command.
- [#754](https://github.com/regen-network/regen-ledger/pull/754) Add Create transaction CLI command.
- [#765](https://github.com/regen-network/regen-ledger/pull/765) Add codec and server registration.
- [#762](https://github.com/regen-network/regen-ledger/pull/762) Add Create message server method.
- [#764](https://github.com/regen-network/regen-ledger/pull/764) Add basket genesis initialization.
- [#772](https://github.com/regen-network/regen-ledger/pull/772) Add basket event proto definitions.
- [#771](https://github.com/regen-network/regen-ledger/pull/771) Add basket integration tests.
- [#776](https://github.com/regen-network/regen-ledger/pull/776) Add basket name and prefix updates.
- [#787](https://github.com/regen-network/regen-ledger/pull/787) Add basket supply invariant.
- [#769](https://github.com/regen-network/regen-ledger/pull/769) Add basket simulation tests.
- [#803](https://github.com/regen-network/regen-ledger/pull/803) Add classes to basket query response.

#### Changed

- [#764](https://github.com/regen-network/regen-ledger/pull/764) Update genesis to support ORM and non-ORM genesis.
- [#789](https://github.com/regen-network/regen-ledger/pull/789) Update consensus version of ecocredit module and service registration.

#### Fixed

- [#807](https://github.com/regen-network/regen-ledger/pull/807) Fix attributes on ecocredit receive events

## [v2.1.0](https://github.com/regen-network/regen-ledger/releases/tag/v2.1.0) - 2021-11-23

### General

#### Fixed

- [#654](https://github.com/regen-network/regen-ledger/pull/654) Add patch for IBC connection parameter

#### Changed

- [#657](https://github.com/regen-network/regen-ledger/pull/657) Update go.mod & imports to adhere to golang semver guidelines for regen-ledger/v2
- [#658](https://github.com/regen-network/regen-ledger/pull/658) Upgrade `ibc-go` to v2.0.0

## [v2.0.0](https://github.com/regen-network/regen-ledger/releases/tag/v2.0.0) - 2021-10-29

### General

#### Added

- [#388](https://github.com/regen-network/regen-ledger/pull/388) Add support for rosetta
- [#482](https://github.com/regen-network/regen-ledger/pull/482)
    Add support for on-chain creation of Permanent Locked Accounts
    ([regen-network/cosmos-sdk#42](http://github.com/regen-network/cosmos-sdk/pull/42))
- [#349](https://github.com/regen-network/regen-ledger/pull/349) Add x/feegrant & x/authz from Cosmos SDK v0.43
- [#538](https://github.com/regen-network/regen-ledger/pull/538) Add script for starting a local test node

#### Changed

- [#422](https://github.com/regen-network/regen-ledger/pull/422) remove `Request` suffix in Msgs
- [#322](https://github.com/regen-network/regen-ledger/pull/322) Split regen ledger into multiple go modules
- [#580](https://github.com/regen-network/regen-ledger/pull/580) Update SDK fork to v0.44.2-regen-1
- [#587](https://github.com/regen-network/regen-ledger/pull/587) Update Go to v1.17.

#### Fixed

- [#386](https://github.com/regen-network/regen-ledger/pull/386) fix IBC proposal registration

### x/ecocredit

#### Added

- (genesis) [#389](https://github.com/regen-network/regen-ledger/pull/389) add genesis import and export
- [#385](https://github.com/regen-network/regen-ledger/pull/385) add support for credit cancelling
- [#425](https://github.com/regen-network/regen-ledger/pull/425) add params for an allowlist of permissioned credit designers
- [#451](https://github.com/regen-network/regen-ledger/pull/451) add queries to list classes and batches with a class
- [#183](https://github.com/regen-network/regen-ledger/pull/183) add grpc-gateway support for query routes
- [#539](https://github.com/regen-network/regen-ledger/pull/539) Add methods for updating a credit class
- [#555](https://github.com/regen-network/regen-ledger/pull/555) Add ecocredit params query

#### Changed

- [#375](https://github.com/regen-network/regen-ledger/pull/375) add fixed fee for creating new credit class
- [#392](https://github.com/regen-network/regen-ledger/pull/392) update class ID and batch denomination formats
- [#328](https://github.com/regen-network/regen-ledger/pull/328) record retirement locations of ecocredit
- [#393](https://github.com/regen-network/regen-ledger/pull/393) add dates as top level fields in credit batches
- [#394](https://github.com/regen-network/regen-ledger/pull/394) add project location as field in credit batches
- [#435](https://github.com/regen-network/regen-ledger/pull/435) use dec wrapper for decimal operations
- [#424](https://github.com/regen-network/regen-ledger/pull/424) add credit types to credit class
- [#500](https://github.com/regen-network/regen-ledger/pull/500) Rename credit class designer to admin
- [#540](https://github.com/regen-network/regen-ledger/pull/540) Add max-metadata check for credit class and credit batch
- [#526](https://github.com/regen-network/regen-ledger/pull/526) Add gas per-loop-iteration in ecocredit messages
- [#554](https://github.com/regen-network/regen-ledger/pull/554) Add ValidateDenom for MsgSend, MsgRetire and MsgCancel

#### Fixed

- [#591](https://github.com/regen-network/regen-ledger/pull/591) Set credit class fee in upgrade handler
- [#592](https://github.com/regen-network/regen-ledger/pull/592) Fix `undefined` error message when creating a credit class

### x/group

#### Added

- [#330](https://github.com/regen-network/regen-ledger/pull/330) add invariant checks for groups' vote sums
- [#333](https://github.com/regen-network/regen-ledger/pull/333) try to execute group proposal on submission or on new vote
- [#183](https://github.com/regen-network/regen-ledger/pull/183) add grpc-gateway support for query routes

### orm

#### Fixed

- [#518](https://github.com/regen-network/regen-ledger/pull/518) Fix bytes key field to have a max length
- [#525](https://github.com/regen-network/regen-ledger/pull/525) Fix IndexKeyCodec prefixing issue.

## [v1.0.0](https://github.com/regen-network/regen-ledger/releases/tag/v1.0.0) - 2021-04-13

This release is the version of regen-ledger that will be used for the mainnet launch of Regen Network's blockchain (chain-id: `regen-1`).

It enables configurable builds for regen ledger (by building with an `EXPERIMENTAL=true/false` build flag). With this new configuration, we've made the following delineation.

- Stable build (EXPERIMENTAL=false) is intended for Regen Network's mainnet, and any testing networks aiming to replicate the mainnet configuration.
  - Includes all standard modules from the Cosmos SDK (bank/staking/gov/etc.), as well as IBC
- Experimental builds, are intended to have more experimental features which have not gone through a full internal audit and are intended for devnets and 3rd party developers who want to work on integrating with future features of regen ledger.
  - In addition to stable build modules, experimental build includes:
    - Regen specific modules (x/ecocredit, x/data)
    - CosmWasm
    - x/group

It is not guaranteed that APIs of features in the experimental build will remain consistent until they are migrated to the stable configuration.

### Added

- make configurable builds (#256)
- add remaining group events
- add group module documentation (#314)

### Changed

- upgrade to Cosmos SDK v0.42.4
- update group tx commands
- remove colon from regen addresses

## [v0.6.0](https://github.com/regen-network/regen-ledger/releases/tag/v0.6.0) - 2021-02-04

This release contains first iterations of the `x/ecocredit` and `x/data` modules which were launched in a Devnet as part of the Open Climate Collabathon in Nov 2020.

It is more or less a full rewrite of regen-ledger to upgrade it to Stargate (Cosmos SDK v0.40)

It also includes an initial draft of the `x/group` module for on-chain multisig and DAO use cases.

### Added

- Data Module Proof of Consept (#118)
- Eco-Credit Module Proof of Concept (#119)
- Addition of vuepress docs site: docs.regen.network (#158)
- Add CosmWasm module to regen ledger (#148)
- Add group module (#154)

### Changed

- Custom protobuf service codegen (#207)
- Update to SDK v0.40.0 (#219)
- Remove usage/naming of `gaia` / `XrnApp` / `simd`

## [v0.5.0](https://github.com/regen-network/regen-ledger/releases/tag/v0.5.0) - 2019-09-21

This release provides the amazonas test upgrade the regen-test-1001 testnet. Specifically this release packages the following changes to the upgrade module:

when an upgrade is planned, the new binary which contains code for the planned upgrade will panic if it is started too early
upgrade scripts are disabled because they were glitchy to setup and not recommended

## [v0.4.0](https://github.com/regen-network/regen-ledger/releases/tag/v0.4.0) - 2019-06-04

### Changed

- [\#166185199](https://www.pivotaltracker.com/story/show/166185199) Temporarily disable all custom modules beside `geo` because they need to be integrated with the new app module setup and this can be a good test case for a coordinated tesnet upgrade
- [\#163156528](https://www.pivotaltracker.com/story/show/163156528) Use stored geo shape for ESP results
- [\#164056249](https://www.pivotaltracker.com/story/show/164056249) Rename `agent` -> `group` module, align structure of groups with specification document
- [\#16](https://github.com/regen-network/regen-ledger/issues/16) The on-chain store data command now only works with graphs defined by the graph package
- [\#15](https://github.com/regen-network/regen-ledger/issues/15) Test and debug upgrade module in Cosmos PR [\#4233](https://github.com/cosmos/cosmos-sdk/pull/4233) against an
internal testnet

### Added

- [\#163101853](https://www.pivotaltracker.com/story/show/163101853) Proposal index tags support
- [\#163101852](https://www.pivotaltracker.com/story/show/163101852) ESP version proposal index tags
- [\#163107520](https://www.pivotaltracker.com/story/show/163107520) Bech32 agent ID's
- [\#163101848](https://www.pivotaltracker.com/story/show/163101848) ESP version index tags
- [\#163101847](https://www.pivotaltracker.com/story/show/163101847) ESP result index tags
- [\#163101851](https://www.pivotaltracker.com/story/show/163101851) ESP result proposal index tags
- [\#163102182](https://www.pivotaltracker.com/story/show/163102182) Cosmos cli support for printing Tags and string Data after submitting tx's
- [\#163098168](https://www.pivotaltracker.com/story/show/163098168) Return agent ID after creating from CLI
- [\#163098166](https://www.pivotaltracker.com/story/show/163098166) Get ESP result ID from CLI after submitting
- [\#162943157](https://www.pivotaltracker.com/story/show/162943157) Store geo shape support
- [\#163098169](https://www.pivotaltracker.com/story/show/163098169) Proposal CLI get query support
- [\#163098032](https://www.pivotaltracker.com/story/show/163098032) Agent CLI query support
- [\#163987749](https://www.pivotaltracker.com/story/show/163987749) [\#163831108](https://www.pivotaltracker.com/story/show/163831108) Configured Gitlab CI for build and tests
- [\#163963339](https://www.pivotaltracker.com/story/show/163963339) Upgrade to latest Cosmos SDK
- [\#164334136](https://www.pivotaltracker.com/story/show/164334136) Implement Postgreql indexer package
- [\#164380651](https://www.pivotaltracker.com/story/show/164380651) Implement Postgres indexing for geo module
- [\#17](https://github.com/regen-network/regen-ledger/issues/17) Add define property schema support
- [\#18](https://github.com/regen-network/regen-ledger/issues/18) Graph package and binary serialization format
- [\#27](https://github.com/regen-network/regen-ledger/issues/27) Create claim module
- [\#166185199](https://www.pivotaltracker.com/story/show/166185199) Integrate Cosmos staking modules

## [v0.3.0](https://github.com/regen-network/regen-ledger/releases/tag/v0.3.0) - 2019-01-09

- Updated `xrn-test-2` testnet at height `1000`

### Fixed

- Storage of ESP results bug

## [v0.2.0](https://github.com/regen-network/regen-ledger/releases/tag/v0.2.0) - 2019-01-09

- Launched `xrn-test-2` testnet

### Added

- [\#163071770](https://www.pivotaltracker.com/story/show/163071770) Agent genesis.json support
- [\#162944049](https://www.pivotaltracker.com/story/show/162944049) Planned upgrade module
- [\#162944050](https://www.pivotaltracker.com/story/show/162944050) Consortium upgrade action support

### Fixed

- [\#163040931](https://www.pivotaltracker.com/story/show/163040931) Signature verification failed bug

## [v0.1.0](https://github.com/regen-network/regen-ledger/releases/tag/v0.1.0) - 2018-12-19

- Launched `xrn-1` testnet

### Added

- [\#162640229](https://www.pivotaltracker.com/story/show/162640229) Create ESP cosmos msg
- [\#162640226](https://www.pivotaltracker.com/story/show/162640226) Create simple agent cosmos msg
- [\#162640231](https://www.pivotaltracker.com/story/show/162640231) Report ESP result cosmos msg
- [\#162640230](https://www.pivotaltracker.com/story/show/162640230) Register ESP version cosmos msg
- Created `data` module

[Unreleased]: https://github.com/regen-network/regen-ledger/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/regen-network/regen-ledger/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/regen-network/regen-ledger/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/regen-network/regen-ledger/compare/fcc6887b...v0.1.0
