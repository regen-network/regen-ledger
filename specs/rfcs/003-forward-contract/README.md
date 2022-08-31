# RFC-003: Forward Contract

- Created: YYYY-MM-DD
- Status: __DRAFT__ | __IN REVIEW__ | __ACCEPTED__ | __REJECTED__ | __SUPERSEDED__ | __ABANDONED__
- Superseded By:
- RFC PR: [#]()
- Authors:

### Table of Contents

- [Summary](#summary)
- [Need](#need)
- [Approach](#approach)
- [Rationale](#rationale)
- [References](#references)
- [Changelog](#changelog)

## Summary

<!-- Brief explanation of what the RFC attempts to address. -->

...

## Need

<!-- What's the identified need? A need should relate to an important and specific opportunity or use case. -->

...

## Approach

<!-- The recommended approach to fulfill the needs presented in the previous section. -->

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
  // information about the contract such as estimated total supply, forward
  // contract supply, and estimated price per credit type unit.
  string metadata = 3;

  // funds_to_collect is the denom and amount the project is collecting.
  cosmos.base.v1beta1.Coin funds_to_collect = 4;

  // volume_percentage is the percent of all credits issued that will be
  // available to purchase in shares.
  string volume_percentage = 5;

  // start_date is the delivery start date.
  google.protobuf.Timestamp start_date = 6;

  // end_date is the delivery end date.
  google.protobuf.Timestamp end_date = 7;
}
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
  // information about the contract such as estimated total supply, forward
  // contract supply, and estimated price per credit type unit.
  string metadata = 3;

  // funds_to_collect is the denom and amount the project is collecting.
  cosmos.base.v1beta1.Coin funds_to_collect = 4;

  // volume_percentage is the percent of all credits issued that will be
  // available to purchase in shares.
  string volume_percentage = 5;

  // start_date is the delivery start date.
  google.protobuf.Timestamp start_date = 6;

  // end_date is the delivery end date.
  google.protobuf.Timestamp end_date = 7;
}
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

### Fund Project

Any account can view available contracts and fund a project. When an account funds a project, the account is buying a share of the available shares defined within the contract. The token denom must match the token denom used in the contract.

The funds are transferred to the project admin upon funding and (either the shares are defined in custom state similar to ecocredits or bank tokens specific to the contract are minted and transferred to the funder - see [Contract Shares](#contract-shares) below).

When an account funds a project and therefore owns a share of future credits issued from the project, an issuance policy is automatically created and managed programmatically (i.e. no account has the authority to update or remove the issuance policy).

### Contract Shares

Tokens may be distributed to shareholders based on the amount of funding they provide but managing the issuance of credits based on token holders may introduce additional challenges...

How to represent a forward contract as an asset that can be traded?

- (1) Selling a share of a forward contract creates a new object in state that represents the shares, deducting the share from the available shares in the contract and creating a new object with the share owner and the share percentage.
- (2) When a forward contract is created, tokens are minted that represent the share of future credits and the tokens are the asset.
  - (A) The amount of tokens minted is defined at the time the contract is created.
  - (B) The amount of tokens minted is calculated based on the share and reflects the percent of future credits.

### Contract Expiration

How does contract expiration work?

- (1) The issuance policy is removed on the end date and therefore no further credits are issued to the funder.
- (2) The issuance policy is never removed. The issuance policy would need to define a range of dates within which the policy is valid.

### Automated Credit Issuance

The credits issued to shareholders should be automated and the functionality to support this will be explored as a separate feature set. If automatic issuance is not delivered with forward contract functionality, it will be the responsibility of the credit issuers to issue credits according to existing contracts.

<!-- (TODO: replace with issuance policy RFC link) -->
See [(x/ecocredit): issuance policy proof-of-concept][3]

### Project Approval Process

In additional to forward contract functionality, the [Forward Contract Bond Module][1] document describes a set of features around a project approval process. This feature set is being explored independently of the functionality described within this document and, although these features might be built in parallel, one is not dependent on the other.

The existing credit class and project functionality can be used to accomplish the project approval process outlined in [Forward Contract Bond Module][1]. In the current implementation, a project can only be created within a credit class by an approved issuer of the credit class. An organization that intends to audit projects would need to administer a credit class and designate issuers to manage project audits. The application process would be handled off chain and the issuer would approve a project by creating a project on chain.

<!-- TODO: replace with independent projects RFC link -->
See [(x/ecocredit): independent projects proof-of-concept][4]

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

- [](#) First draft submitted by <authors>
