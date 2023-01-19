# RFC-003: Issuance Policy

- Created: 2023-01-18
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

Regen Ledger enables entities (individuals or organizations) to design and issue credits for ecosystem services in the form of on-chain assets. Ecosystem service credits are issued in credit batches by an issuer of the credit class within which the project is registered. At the time of issuance, there are currently no restrictions enforced programmatically as to how those credits are issued, i.e. what amounts or percentages of each issuance are to be issued to which addresses and whether those credits are to be issued in a tradable or retired state.

A credit class issuer is an account entrusted by the credit class admin to issue credits for projects registered within the credit class. The requirements enforced by issuance policies is not designed as a security precaution for dishonest credit class issuers but rather as an aid to credit class issuers managing the issuance of credits for multiple projects and having to track and account for contractual agreements with various parties receiving those credits. Issuance policies also provide added assurance to those contracted to receive credits will in fact receive those credits based on an agreement held with the credit class and/or project.

### Off-Chain Forward Contracts

A forward contract requires credits to be delivered either in full or over time to an investor as the credits become available. Issuance policies would enable an off-chain agreement to be enforced on chain; a percentage of each credit issuance could be issued to an account until an amount of credits has been delivered or an end date has been reached. An issuance policy would enable this requirement to be enforced programmatically on chain and therefore properly accounted for by credit class issuers.

### On-Chain Forward Contracts

This proposal is written in parallel with a proposal for [on-chain forward contract functionality][1]. On-chain forward contract functionality would enable a project to create a forward contract, the forward contract would then be approved by a credit class issuer, and investors would then be able to purchase shares of future credits issued by the project. In the initial implementation, an issuance policy would be created at the time of investment and therefore each credit issuance from the project within the timeframe of the contract would require a percentage of the credits to be issued to the investor.

### Royalties and Other Agreements

Other agreements such as royalties (where credits are used as the unit of payment) can be enforced through issuance policies. Various parties involved in producing, monitoring, verifying, reviewing, or issuing credits may have contractual agreements with the credit class or project that guarantees credits are received in the form of royalties.

## Approach

<!-- The recommended approach to fulfill the needs presented in the previous section. -->

## Rationale

<!-- Include an overview of what tradeoffs exist when taking this approach, what benefits come from it, and/or what alternatives were considered. -->

## References

1. [docs: add RFC for forward contract functionality][1]

[1]: https://github.com/regen-network/regen-ledger/pull/1474

## Changelog

<!-- An RFC should include a changelog, providing a record of any significant changes. -->

- [#1746](https://github.com/regen-network/regen-ledger/pull/1746) First draft submitted by @ryanchristo
