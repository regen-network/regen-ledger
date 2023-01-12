# RFC-003: Forward Contract

- Created: 2023-01-12
- Status: __IN REVIEW__
- Superseded By: NA
- RFC PR: [#1474](https://github.com/regen-network/regen-ledger/pull/1474)
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

This RFC aims to lay out a high-level architecture and feature set for forward contract functionality.

## Need

<!-- What's the identified need? A need should relate to an important and specific opportunity or use case. -->

Regen Ledger enables entities (individuals or organizations) to design and issue credits for ecosystem services in the form of on-chain assets. Ecosystem service credits represent positive ecological outcomes and are issued after those outcomes have been measured and quantified. Project developers receive payment for their services only after buyers purchase their credits and projects are therefore responsible for covering upfront costs or seeking out financial support to start or continue providing ecosystem services.

Forward contract functionality would enable project developers to offer a percentage of future credits issued and receive funding before those credits have been measured and quantified. Project developers would work alongside credit class issuers to submit a forward contract that would offer a percentage of future credits issued in exchange for upfront funding. After the contract has been approved, investors would be able to provide upfront funding for the project in exchange for a percentage of future credits issued from the project.

### Fast Forward Pilot

...

### Earthbanc Use Case

...

## Approach

<!-- The recommended approach to fulfill the needs presented in the previous section. -->

Forward contract functionality would be implemented as a `contract` submodule within the existing `ecocredit` module. For reference, [(x/ecocredit): forward contract proof-of-concept][2] was submitted as an initial proof-of-concept to help illustrate this architecture.

This proposal separates the approach into multiple stages of implementation. The initial stage (i.e. [Stage 1](#stage-1) is designed to serve the needs of the [Fast Forward Pilot](#fast-forward-pilot) and lay the foundation for the [Earthbanc Use Case](#earthbanc-use-case). The second stage (i.e. [Stage 2](#stage-2) is designed to serve the remaining needs of the [Earthbanc Use Case](#earthbanc-use-case).

### Stage 1

The first stage includes the implementation of direct credit issuance and support for forward contracts that are specific to a single project whereby the project is properly vetted by a credit class issuer and the risk is shared by the credit class (a risk in reserve and reputation) and the investor(s) (a risk in investment).

In this initial implementation, there is only one option for receiving future credits issued, which is the direct issuance of credits to the account that has a claim on the forward contract. The credits are delivered over-time as they are issued from the project; the account that has a claim to future credits (i.e. the investor) receives a percentage of each credit batch issuance that has a monitoring period within the timeframe of the contract. The percentage of credits delivered with each credit batch issuance is based on the percentage of the claim to future credits and enforced by on-chain functionality.

Each forward contract is specific to a single project. The project should be properly vetted by the credit class issuer(s) and the investor(s) will need to trust the credit class and/or vet the project themselves. The admin and issuer(s) of a credit class are responsible for defining their own vetting process for projects and the issuer (either an individual or group account) that approves the contract will be responsible for following that process and assessing the risk of the project.

The price and volume estimates for the credits issued from the project within the timeframe of the contract will be set by the project admin and approved by the credit class issuer through an on-chain contract creation and review process. How the price and volume estimates are calculated are outside the scope of the on-chain functionality outlined within this proposal but additional information about how the calculations were made can and should be stored within the forward contract as metadata.

To mitigate the risk of the project under-delivering and therefore the risk in investment, each contract will have a reserve pool specifically for the contract where previously issued credits from the project or an equivalent project can be transferred by the credit class issuer. The reserve pool can only receive credits once the contract has been approved and the credits in the reserve pool would be held escrow until the end of the contract. The reserve pool can only receive credits from the same credit class and up to the volume of future credits being sold via the forward contract and the credits can only be transferred by the credit class issuer or the project admin defined in the contract. Whether the reserve pool backs the total volume of future credits being sold is up to the credit class issuer with more credits providing less risk for the investor(s) and more likelihood of the project receiving pre-financing.

The accepted form of funds (the token denomination) is decided by the project admin and approved by the credit class issuer. There are no restrictions on which token denomination the project chooses to receive funding but a stable coin would be the most probably choice.

### Create Contract

A project admin can create a contract. The contract is "proposed" until "approved" by a credit class issuer.

```protobuf
// MsgCreate is the Msg/Create request type.
message MsgCreate {

  // project_id is the unique identifier of the project.
  string project_id = 1;

  // project_admin is the admin of the project.
  string project_admin = 2;

  // metadata is any arbitrary string that includes or references additional
  // information about the contract including the initial amount of funds to
  // collect, the initial volume percentage offered, estimated total supply,
  // forward contract supply, and estimated price per credit type unit.
  string metadata = 3;

  // funds_to_collect is the denom and amount the project is collecting.
  cosmos.base.v1beta1.Coin funds_to_collect = 4;

  // volume_percentage is the percent of all credits issued that will be
  // available to purchase.
  string volume_percentage = 5;

  // start_date is the delivery start date.
  google.protobuf.Timestamp start_date = 6;

  // end_date is the delivery end date.
  google.protobuf.Timestamp end_date = 7;
}
```

```feature
Rule: Only a project admin can create a contract for their project
```

### Update Contract

A project admin can update a contract with status "proposed". Once the contract status is "approved", the contract cannot be updated.

```protobuf
// MsgUpdate is the Msg/Update request type.
message MsgUpdate {

  // id is the unique identifier of the contract.
  uint64 id = 1;

  // project_admin is the admin of the project.
  string project_admin = 2;

  // metadata is any arbitrary string that includes or references additional
  // information about the contract including the initial amount of funds to
  // collect, the initial volume percentage offered, estimated total supply,
  // forward contract supply, and estimated price per credit type unit.
  string metadata = 3;

  // funds_to_collect is the denom and amount the project is collecting.
  cosmos.base.v1beta1.Coin funds_to_collect = 4;

  // volume_percentage is the percent of all credits issued that will be
  // available to purchase.
  string volume_percentage = 5;

  // start_date is the delivery start date.
  google.protobuf.Timestamp start_date = 6;

  // end_date is the delivery end date.
  google.protobuf.Timestamp end_date = 7;
}
```

```feature
Rule: Only a project admin can update a contract for their project
```

### Cancel Contract

A project admin can cancel a contract with status "proposed". Once the contract status is "approved", the contract cannot be cancelled.

```protobuf
// MsgCancel is the Msg/Cancel request type.
message MsgCancel {

  // id is the unique identifier of the contract.
  uint64 id = 1;

  // project_admin is the admin of the project.
  string project_admin = 2;
}
```

```feature
Rule: Only a project admin can cancel a contract for their project
```

### Approve Contract

A credit class issuer can approve a contract. Once the contract is "approved", the contract cannot be updated and any account can fund the project.

```protobuf
// MsgApprove is the Msg/Approve request type.
message MsgApprove {

  // id is the unique identifier of the contract.
  uint64 id = 1;

  // class_issuer is the address of the credit class issuer that is approving
  // the contract on behalf of the credit class.
  string class_issuer = 2;
}
```

```feature
Rule: The volume percentage cannot exceed the sum percentage of existing issuance policies
```

```feature
Rule: An issuance policy is created for tradable credits and the submodule address is the recipient
```

### Fund Project

Any account can view available contracts and fund a project. When an account funds a project, the account has a claim to credits issued from the project. The funds are transferred directly to the project admin.

When an account funds a project and therefore owns a claim on future credits issued from the project, an issuance policy is automatically created and managed programmatically (i.e. no account has the authority to update the issuance policy and issuance policies would require a property that distinguishes policies managed programmatically from those managed by accounts with authority).

```protobuf
// MsgFundProject is the Msg/FundProject request type.
message MsgFundProject {

  // id is the unique identifier of the contract.
  uint64 id = 1;

  // funder is the address of the account funding the project and receiving a
  // share of future credits issued from the project.
  string funder = 2;

  // volume_percentage is the percent of all credits issued that the funder will
  // receive.
  string volume_percentage = 3;

  // funds is the token denom and amount the funder is providing in return for
  // the specified volume percentage.
  cosmos.base.v1beta1.Coin funds = 4;
  
  // direct_issuance determines whether the credits will be automatically issued
  // to the funder (i.e. the funder account will be assigned as the recipient of
  // the issuance policy) or whether the funder will receive claim tokens (i.e.
  // the contract submodule will be assigned as the recipient of the issuance
  // policy and the credits can be claimed at a later point in time).
  bool direct_issuance = 5;
  
  // auto_retire determines whether the credits will be automatically retied upon
  // issuance (i.e. the issuance policy will be set to auto-retire). This option
  // only applies if direct issuance is enabled.
  bool auto_retire = 6;
  
  // retirement_jurisdiction is the jurisdiction of the funder. A jurisdiction is
  // only required if auto-reture is enabled.
  string retirement_jurisdiction = 7;
}
```

```feature
Rule: The contract must be approved
```

```feature
Rule: The funds denom must match the denom defined in the contract
```

```feature
Rule: The funds amount must be greater than or equal to the calculated cost
```

```feature
Rule: The funds are deducted from the funder
```

```feature
Rule: The funds are added to the project admin
```

### Contract State

```protobuf
// Contract defines a forward contract and the table within which the forward
// contract is stored.
message Contract {
  option (cosmos.orm.v1alpha1.table) = {
    id : 1,
    primary_key : {fields : "id", auto_increment : true}
  };

  // id is the table row identifier of the contract.
  uint64 id = 1;

  // project_id is the unique identifier of the project.
  string project_id = 2;

  // status is the status of the contract (e.g. "proposed", "approved").
  ContractStatus status = 3;

  // metadata is any arbitrary string that includes or references additional
  // information about the contract including the initial amount of funds to
  // collect, the initial volume percentage offered, estimated total supply,
  // forward contract supply, and estimated price per credit type unit.
  google.protobuf.Timestamp metadata = 4;

  // volume_percentage is the remaining percent of all credits issued that
  // are be available to claim.
  string volume_percentage = 5;

  // funds_to_collect is the denom and remaining amount of funds the project
  // is collecting.
  cosmos.base.v1beta1.Coin funds_to_collect = 6;

  // start_date is the delivery start date.
  google.protobuf.Timestamp start_date = 7;

  // end_date is the delivery end date.
  google.protobuf.Timestamp end_date = 8;
}
```

### Contract Expiration

The contract will automatically expire once the contracted credit volume has been delivered or when the contract end date is reached. If the end date is reached before the volume has been delivered, credits from the reserve pool will be delivered in place of the contracted credits up to the amount available in the reserve pool. 

### Automated Credit Issuance

Direct credit issuance to investor will be enforced with on-chain functionality. At the time of funding the project, investors would have the option of choosing to receive credits in a retired or tradable state upon issuance. This functionality will be handled by "issuance policies" and explored as a separate feature set built alongside the initial implementation of forward contract functionality.

<!-- (TODO: replace with issuance policy RFC link) -->
See [(x/ecocredit): issuance policy proof-of-concept][3] for more information.

### Project Approval Process

In additional to forward contract functionality, the [Forward Contract Bond Module][1] document describes a set of features around a project approval process. The contract approval process only partially fulfills the outlined approval process. Additional functionality is being explored independently of the functionality described within this document and, although these features might be built in parallel, one is not dependent on the other.

The existing credit class and project functionality can be used to accomplish the project approval process outlined in [Forward Contract Bond Module][1]. In the current implementation, a project can only be created within a credit class by an approved issuer of the credit class. An organization that intends to audit projects would need to administer a credit class and designate issuers to manage project audits. The application process would be handled off chain and the issuer would approve a project by creating a project on chain.

<!-- TODO: replace with independent projects RFC link -->
See [(x/ecocredit): independent projects proof-of-concept][4] for more information.

### Stage 2

Investors would have the option of receiving tokens that would represent their claim on the contract and later be used to redeem credits from a claim account (i.e. an account managed by the submodule). Receiving tokens would enable investors to transfer their claim and the owner(s) of those tokens would then be able to redeem credits at a later point in time.

When receiving tokens that represent a claim, the amount of tokens received would be calculated based on the percent of credits that the investor has a claim to and the tokens would be specific to the project and the most recent batch claimed. Rather than an issuance policy being created using the investor account as the recipient, an issuance policy would be created (or updated) using a claim account as the recipient. Investor would then be free to transfer the tokens and the owner(s) of the tokens would then be able to claim the credits at a time of their choosing.

When credits are claimed, the owner of the tokens would exchange the tokens for the credits issued from the project and receive new tokens equal to the amount of tokens used to redeem credits. The token denom for claim tokens would include the batch sequence number of the most recent batch claimed and redeeming credits would burn the tokens used to claim the credits and return new tokens with an updated denom.

## Rationale

<!-- Include an overview of what tradeoffs exist when taking this approach, what benefits come from it, and/or what alternatives were considered. -->

...

## References

1. [Forward Contract Bond Module][1]
2. [(x/ecocredit): forward contract proof-of-concept][2]
3. [(x/ecocredit): issuance policy proof-of-concept][3]
4. [(x/ecocredit): independent projects proof-of-concept][4]

[1]: https://docs.google.com/document/d/1_BMb7dUVdYEiL5n1LNCjHtd6WYPXgxRY_gwflIwCXgE
[2]: https://github.com/regen-network/regen-ledger/pull/1420
[3]: https://github.com/regen-network/regen-ledger/pull/1421
[4]: https://github.com/regen-network/regen-ledger/pull/1422

## Changelog

<!-- An RFC should include a changelog, providing a record of any significant changes. -->

- [#1474](https://github.com/regen-network/regen-ledger/pull/1474) First draft submitted by @ryanchristo
