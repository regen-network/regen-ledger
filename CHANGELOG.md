# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
