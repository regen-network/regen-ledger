---
name: Module Readiness Checklist
about: Pre-flight checklist that modules must pass before ready for being enabled in stable builds of regen ledger
labels: 'module-readiness-checklist'
---

- [ ] Good unit and integration test coverage for all state machine logic (@assignee)
- [ ] Internal API audit (at least 1 person) (@assignee)
  - [ ] Are Msg and Query methods and types well-named and organized?
  - [ ] Is everything well documented (godoc as well as README.md in module directory)
- [ ] Internal state machine audit (at least 2 people) (@assignee1, @assignee2)
  - [ ] Read through MsgServer code and verify correctness upon visual inspection
  - [ ] Ensure all state machine code which could be confusing is properly commented
  - [ ] Make sure state machine logic matches Msg method documentation
  - [ ] Ensure that all state machine edge cases are covered with tests and that test coverage is sufficient
  - [ ] Assess potential threats for each method including spam attacks and ensure that threats have been addressed sufficiently. This should be done by writing up threat assessment for each method
  - [ ] Assess potential risks of any new third party dependencies and decide whether a dependency audit is needed
- [ ] Internal completeness audit, fully impleted with tests (at least 1 person) (@assignee)
  - [ ] Genesis import and export of all state
  - [ ] Query services
  - [ ] CLI methods
- [ ] Testnet / devnet testing after above internal audits
  - [ ] All Msg methods have been tested especially in light of any potential threats identified
  - [ ] Genesis import and export has been tested
- [ ] Nice to have (and needed in some cases if threats could be high): official 3rd party audit
