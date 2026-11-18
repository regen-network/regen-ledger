# RFC-004: Issuance Policy

- Created: 2023-01-26
- Status: __DRAFT__
- Superseded By: NA
- RFC PR: [#1746](https://github.com/regen-network/regen-ledger/pull/1746)
- Authors: @ryanchristo

### Table of Contents

- [Summary](#summary)
- [Need](#need)
- [Approach](#approach)
- [Rationale](#rationale)
- [References](#references)
- [Changelog](#changelog)

## Summary

<!-- Brief explanation of what the RFC attempts to address. -->

This RFC aims to lay out a high-level architecture and feature set for on-chain functionality that would enable project-based "issuance policies", requirements enforced programmatically when credits are issued from a project.

## Need

<!-- What's the identified need? A need should relate to an important and specific opportunity or use case. -->

Regen Ledger provides a framework that enables individuals or organizations to design and issue credits for ecosystem services in the form of on-chain assets. Ecosystem service credits are issued in batches by a credit class issuer for projects registered within the credit class.

There are currently no issuance requirements enforced programmatically, i.e. what percentages of each issuance are to be issued to which accounts and whether credits are to be issued in a tradable or retired state. Project issuance policies would enable credit class administrators to set requirements on the issuance of credits for specific projects. Project issuance policies would also enable issuance policies to be created and managed programmatically by on-chain functionality such as forward contracts. 

### Off-Chain Forward Contracts

A forward contract requires credits to be delivered either in full or over time to an investor as the credits become available. Issuance policies would enable an off-chain agreement to be enforced on chain; a percentage of each credit issuance would be issued to an account until an amount of credits has been delivered or an end date has been reached. Issuance policies would enable such requirements to be enforced programmatically and therefore properly accounted for by credit class issuers.

### On-Chain Forward Contracts

This proposal is written in parallel with a proposal for [on-chain forward contract functionality][1]. On-chain forward contract functionality enables a project to create a forward contract, the forward contract to be approved by a credit class issuer, and investors to purchase future credits issued by the project. In the current specification, an issuance policy would be created and managed programmatically at the time of investment and therefore guaranteeing a percentage of credits from each issuance to be issued to the investor.

### Royalties and Other Use Cases

Other agreements such as royalties (where credits are used as the unit of payment) can be enabled through issuance policies. Various parties involved in producing, monitoring, verifying, reviewing, or issuing credits may have contractual agreements with the credit class or project that guarantees credits are received as a form of payment. Issuance policies could serve a variety of use cases.

## Approach

<!-- The recommended approach to fulfill the needs presented in the previous section. -->

Project issuance policies would be implemented as a policy submodule within the existing ecocredit module. For reference, [feat(x/ecocredit): issuance policy proof-of-concept][2] was submitted as an initial proof-of-concept to help illustrate this architecture.

### Create Policy

```proto
// MsgCreate is the Msg/Create request type.
message MsgCreate {
  option (cosmos.msg.v1.signer) = "admin";

  // project_id is the unique identifier of the project that the issuance policy
  // applies to. The project issuance policy will require all credits from this
  // project to be issued in compliance with this policy.
  uint64 project_id = 1;

  // admin is the address of the admin of the issuance policy. The admin must be
  // the credit class admin within which the project is registered or the module
  // within which the issuance policy will be managed programmatically.
  string admin = 2;

  // percent is the percentage of each credit issuance that will be issued to
  // the recipient. The value is represented as a decimal string that must be
  // less than 1 with a maximum precision of 6.
  string percent = 3;

  // recipient is the address of the account that will receive the percent of
  // each credit issuance.
  string recipient = 4;

  // auto_retire is a boolean that determines whether the credits will be
  // retired upon issuance or remain in a tradable state.
  bool auto_retire = 5;

  // retirement_jurisdiction is the jurisdiction of the recipient and is only
  // required if auto-retire is enabled.
  string retirement_jurisdiction = 6;
}
```

## Policy State

```proto
// Policy defines a project issuance policy and the table within which the
// policy is stored.
message Policy {
  option (cosmos.orm.v1alpha1.table) = {
    id : 1,
    primary_key : {fields : "id", auto_increment : true}
    index : {id : 1, fields : "project_id"}
    index : {id : 2, fields : "admin"}
    index : {id : 3, fields : "recipient"}
  };

  // id is the unique identifier and table row of the project issuance policy.
  uint64 id = 1;

  // admin is the admin of the issuance policy, which is either the admin of the
  // credit class within which the project is registered or the module managing
  // the issuance policy programmatically.
  bytes admin = 2;

  // project_id is the unique identifier of the project that the issuance policy
  // applies to. The project issuance policy will require all credits from this
  // project to be issued in compliance with this policy.
  uint64 project_id = 3;

  // percent is the percentage of each credit issuance that will be issued to
  // the recipient. The value is represented as a decimal string that must be
  // less than 1 with a maximum precision of 6.
  string percent = 4;

  // recipient is the address of the account that will receive the percent of
  // each credit issuance.
  bytes recipient = 5;

  // auto_retire is a boolean that determines whether the credits will be
  // retired upon issuance or remain in a tradable state.
  bool auto_retire = 6;

  // retirement_jurisdiction is the jurisdiction of the recipient and is only
  // required if auto-retire is enabled.
  string retirement_jurisdiction = 7;
}
```

## Rationale

<!-- Include an overview of what tradeoffs exist when taking this approach, what benefits come from it, and/or what alternatives were considered. -->

## References

1. [docs: add RFC for forward contract functionality][1]
2. [feat(x/ecocredit): issuance policy proof-of-concept][2]

[1]: https://github.com/regen-network/regen-ledger/pull/1474
[2]: https://github.com/regen-network/regen-ledger/pull/1421

## Changelog

<!-- An RFC should include a changelog, providing a record of any significant changes. -->

- [#1746](https://github.com/regen-network/regen-ledger/pull/1746) First draft submitted by @ryanchristo
