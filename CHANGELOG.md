# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### General

#### Fixed

* [#591](https://github.com/regen-network/regen-ledger/pull/591) Set credit class fee in upgrade handler

## [v2.0.0-beta1](https://github.com/regen-network/regen-ledger/releases/tag/v2.0.0-beta1) - 2021-08-19

### General

#### Added

* [#388](https://github.com/regen-network/regen-ledger/pull/388) Add support for rosetta
* [#482](https://github.com/regen-network/regen-ledger/pull/482)
    Add support for on-chain creation of Permanent Locked Accounts
    ([regen-network/cosmos-sdk#42](http://github.com/regen-network/cosmos-sdk/pull/42))
* [#349](https://github.com/regen-network/regen-ledger/pull/349) Add x/feegrant & x/authz from Cosmos SDK v0.43

#### Changed

* [#422](https://github.com/regen-network/regen-ledger/pull/422) remove `Request` suffix in Msgs
* [#322](https://github.com/regen-network/regen-ledger/pull/322) Split regen ledger into multiple go modules
* [#482](https://github.com/regen-network/regen-ledger/pull/482) Upgrade Cosmos SDK to v0.43.0

#### Fixed

* [#386](https://github.com/regen-network/regen-ledger/pull/386) fix IBC proposal registration

### `x/ecocredit`

#### Added

* (genesis) [#389](https://github.com/regen-network/regen-ledger/pull/389) add genesis import and export
* [#385](https://github.com/regen-network/regen-ledger/pull/385) add support for credit cancelling
* [#425](https://github.com/regen-network/regen-ledger/pull/425) add params for an allowlist of permissioned credit designers
* [#451](https://github.com/regen-network/regen-ledger/pull/451) add queries to list classes and batches with a class
* [#183](https://github.com/regen-network/regen-ledger/pull/183) add grpc-gateway support for query routes


#### Changed

* [#375](https://github.com/regen-network/regen-ledger/pull/375) add fixed fee for creating new credit class
* [#392](https://github.com/regen-network/regen-ledger/pull/392) update class ID and batch denomination formats
* [#328](https://github.com/regen-network/regen-ledger/pull/328) record retirement locations of ecocredit
* [#393](https://github.com/regen-network/regen-ledger/pull/393) add dates as top level fields in credit batches
* [#394](https://github.com/regen-network/regen-ledger/pull/394) add project location as field in credit batches
* [#435](https://github.com/regen-network/regen-ledger/pull/435) use dec wrapper for decimal operations
* [#424](https://github.com/regen-network/regen-ledger/pull/424) add credit types to credit class

### `x/group`

#### Added

* [#330](https://github.com/regen-network/regen-ledger/pull/330) add invariant checks for groups' vote sums
* [#333](https://github.com/regen-network/regen-ledger/pull/333) try to execute group proposal on submission or on new vote
* [#183](https://github.com/regen-network/regen-ledger/pull/183) add grpc-gateway support for query routes

## [1.0.0] - 2021-04-13

This release is the version of regen-ledger that will be used for the mainnet launch of Regen Network's blockchain (chain-id: `regen-1`).

It enables configurable builds for regen ledger (by building with an `EXPERIMENTAL=true/false` build flag). With this new configuration, we've made the following delineation.

* Stable build (EXPERIMENTAL=false) is intended for Regen Network's mainnet, and any testing networks aiming to replicate the mainnet configuration.
  * Includes all standard modules from the Cosmos SDK (bank/staking/gov/etc.), as well as IBC
* Experimental builds, are intended to have more experimental features which have not gone through a full internal audit and are intended for devnets and 3rd party developers who want to work on integrating with future features of regen ledger.
  * In addition to stable build modules, experimental build includes:
    * Regen specific modules (x/ecocredit, x/data)
    * CosmWasm
    * x/group

It is not guaranteed that APIs of features in the experimental build will remain consistent until they are migrated to the stable configuration.

### Added
* make configurable builds (#256)
* add remaining group events
* add group module documentation (#314)

### Changed
* upgrade to Cosmos SDK v0.42.4
* update group tx commands
* remove colon from regen addresses

## [0.6.0] - 2021-02-04

This release contains first iterations of the `x/ecocredit` and `x/data` modules which were launched in a Devnet as part of the Open Climate Collabathon in Nov 2020.

It is more or less a full rewrite of regen-ledger to upgrade it to Stargate (Cosmos SDK v0.40)

It also includes an initial draft of the `x/group` module for on-chain multisig and DAO use cases.

### Added

* Data Module Proof of Consept (#118)
* Eco-Credit Module Proof of Concept (#119)
* Addition of vuepress docs site: docs.regen.network (#158)
* Add CosmWasm module to regen ledger (#148)
* Add group module (#154)


### Changed

* Custom protobuf service codegen (#207)
* Update to SDK v0.40.0 (#219)
* Remove usage/naming of `gaia` / `XrnApp` / `simd`

## [0.5.0] - 2019-09-21

This release provides the amazonas test upgrade the regen-test-1001 testnet. Specifically this release packages the following changes to the upgrade module:

when an upgrade is planned, the new binary which contains code for the planned upgrade will panic if it is started too early
upgrade scripts are disabled because they were glitchy to setup and not recommended

## [0.4.0] - 2019-06-04

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

## [0.3.0] - 2019-01-09

- Updated `xrn-test-2` testnet at height `1000`

### Fixed
- Storage of ESP results bug

## [0.2.0] - 2019-01-09

- Launched `xrn-test-2` testnet

### Added

- [\#163071770](https://www.pivotaltracker.com/story/show/163071770) Agent genesis.json support
- [\#162944049](https://www.pivotaltracker.com/story/show/162944049) Planned upgrade module
- [\#162944050](https://www.pivotaltracker.com/story/show/162944050) Consortium upgrade action support

### Fixed
- [\#163040931](https://www.pivotaltracker.com/story/show/163040931) Signature verification failed bug

## [0.1.0] - 2018-12-19

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
