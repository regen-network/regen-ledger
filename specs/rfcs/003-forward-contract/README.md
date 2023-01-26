# RFC-003: Forward Contract

- Created: 2023-01-26
- Status: __DRAFT__
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

Regen Ledger provides a framework that enables individuals or organizations to design and issue credits for ecosystem services in the form of on-chain assets. Ecosystem service credits represent positive ecological outcomes and are issued after those outcomes have been measured and quantified. Projects providing ecosystem services usually receive payment after individuals or organizations purchase credits leaving projects responsible for paying upfront costs or seeking out financial support to start, continue, or expand their operations.

Forward contract functionality would enable projects to offer a percentage of future credits issued and receive funding before those credits have been measured and quantified. A Project would work alongside a credit class issuer to estimate the volume and price of credits and then submit a forward contract to be approved by the credit class issuer. Once the contract has been approved, individuals and organizations would then be able to provide upfront funding for projects in exchange for a percentage of future credits issued.

### Fast Forward Pilot

The "Fast Forward Pilot" described within this document is potentially one of many pilots that are being developed as a part of the Fast Forward working group. The pilot described here involves UNDO Carbon (the supplier), Spirals Protocol (the buyer), and Regen Network Development (the infrastructure provider).

The following expectations were discussed and the "✓" indicates what we aligned on for an initial implementation:

**Initial payment**

1. Buyer pays full upfront cost for future credits ✓
2. Buyer pays partial upfront cost for future credits
3. Buyer payment is held in escrow for future credits
4. Buyer pays no upfront cost for future credits

**Payment amount**

1. Discount from current market price based on years out
2. Discount from current market price based on risk
3. Discount from current market price based on combination ✓

**Representation**

1. Buyer receives no asset and contract is baked into the protocol ✓
2. Buyer receives a liquid asset representing the percent of future credits
    1. The asset is tradable
    2. The asset is not tradable
3. Buyer receives vouchers for credits that will mature (i.e. ex-ante)

**Credit delivery (time)**

1. Buyer receives credits as they are issued (percent of each issuance) ✓
2. Buyer receives credits at the end of the contract (total percent/credits)
3. Buyer receives vouchers at the time of purchase (i.e. ex-ante)

**Credit delivery (quantity)**

1. Buyer receives percentage in the contract (no insurance)
    1. If more/less credits are issued, buyer receives more/less
    2. Once expected credit amount issued, buyer receives no more ✓
2. Buyer receives percentage in the contract (with insurance)
    1. Previously issued credits are held in a reserve pool ✓
    2. Future credits issued if credits not delivered within time-frame
3. Buyer receives exact credits that mature (i.e. ex-ante)

...

### Earthbanc Use Case

See [Forward Contract Bond Module][1].

...

## Approach

<!-- The recommended approach to fulfill the needs presented in the previous section. -->

This proposal separates the approach into two stages. The initial stage (i.e. [Stage 1](#stage-1)) is designed to serve the requirements of the [Fast Forward Pilot](#fast-forward-pilot) and lay the foundation for the [Earthbanc Use Case](#earthbanc-use-case). The second stage (i.e. [Stage 2](#stage-2)) is designed to serve the remaining requirements of the [Earthbanc Use Case](#earthbanc-use-case) and is left open-ended for further specification in subsequent version or separate proposal.

### Stage 1

The first stage includes the implementation of direct credit issuance and support for forward contracts that are specific to a single project whereby the project is vetted by a credit class issuer and the risk of the project under-delivering is either shared by the credit class (a risk in reserve credits and reputation) and the investor(s) (a risk in investment) or held solely by the credit class (the credit class provides reserve credits that fully back the future credits issued).

In the initial implementation, there would only be one option for receiving credits, which is the direct issuance of credits to the account that purchases credits via the forward contract (i.e. the investor). The credits are delivered over time as they are issued; the investor receives a percentage of each credit issuance from the project where the monitoring period meets the date criteria of the contract. The percentage of credits delivered with each credit issuance is based on the percentage the investor purchased and enforced by on-chain functionality.

Each forward contract is specific to a single project. The project should be properly vetted by the credit class issuer(s), and the investor(s) will need to trust the credit class and/or vet the project themselves. The admin and issuer(s) of a credit class are responsible for defining their own vetting process for projects and the issuer that approves the contract will be responsible for following that process and assessing the risk of the project and/or work with a third party or other issuers of the credits class to do so.

The price and volume estimates for future credits issued from the project within the timeframe of the contract will be set by the project admin and approved by the credit class issuer through an on-chain contract creation and review process. How the credit price and volume estimates are calculated are outside the scope of on-chain functionality but additional information about how the calculations were made should be stored on chain within the forward contract as verifiable supporting data.

To mitigate the risk of investment, and to therefore improve the likelihood of a project receiving funding, each contract will have a reserve pool specifically for the contract where previously issued credits from the project or an equivalent project can be deposited. The reserve pool would only accept credits once the contract has been approved and the credits would be held in the reserve pool until the end of the contract or until the reserve pool has more credits than what remains in the contract (in which case the difference could be withdrawn).

The reserve pool would only accept credits from the same credit class and from the credit class issuer that approved the contract. Whether the reserve pool backs a partial amount or the total amount of credits being sold is up to the credit class issuer; the more credits held in the reserve pool providing less risk for the investor(s) and more likelihood of the project receiving funds.

The accepted form of funds (i.e. the accepted token denomination) is decided by the project admin and approved by the credit class issuer. There would be no restrictions on what token denomination the project chooses but the token denomination and amount would be unalterable after the contract has been approved, therefore a stable coin would be the most probable choice.

### Contract Submodule

Forward contract functionality would be implemented as a `contract` submodule within the existing `ecocredit` module. For reference, [(x/ecocredit): forward contract proof-of-concept][2] was submitted as an initial proof-of-concept to help illustrate this architecture.

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

  // volume_percentage is the percent of future credits issued that will be
  // available to purchase.
  string volume_percentage = 5;

  // start_date is the contract start date.
  google.protobuf.Timestamp start_date = 6;

  // end_date is the contract end date.
  google.protobuf.Timestamp end_date = 7;
}
```

```feature
Rule: Only the project admin can create a contract
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

  // volume_percentage is the percent of future credits issued that will be
  // available to purchase.
  string volume_percentage = 5;

  // start_date is the contract start date.
  google.protobuf.Timestamp start_date = 6;

  // end_date is the contract end date.
  google.protobuf.Timestamp end_date = 7;
}
```

```feature
Rule: Only the project admin can update the contract
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
Rule: Only a project admin can cancel the contract
```

### Approve Contract

A credit class issuer can approve a contract. Once the contract is "approved", the contract cannot be updated.

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

### Reserve Credits

The credit class issuer can reserve credits in a reserve pool specifically for the forward contract to help mitigate the risk of investment by providing credits that will be distributed to the investor if the project under-delivers.

```protobuf
// MsgReserve is the Msg/Reserve request type.
message MsgReserve {
  option (cosmos.msg.v1.signer) = "issuer";

  // contract_id is the unique identifier of the contract.
  uint64 contract_id = 1;

  // issuer is the address of the issuer that approved the contract. Only the
  // issuer in the contract can send credits to the contract reserve.
  string issuer = 2;

  // batch_denom is the batch denom of the credits being sent.
  string batch_denom = 3;

  // tradable_amount is the amount of tradable credits being sent.
  string tradable_amount = 4;
}
```

```feature
Rule: Only the credit class issuer in the contract can send credits to the reserve
```

```feature
Rule: The reserve can only receive credits from the same credit class as the project
```

```feature
Rule: The reserve cannot receive more credits than the total volume of contracted credits
```

### Invest in Contract

Any account can view available contracts and fund a project. When an account funds a project, the account has a claim to future credits issued from the project. The funds are transferred directly to the project admin.

When an account funds a project and therefore owns a claim on future credits issued from the project, an issuance policy is automatically created and managed programmatically (i.e. no account has the authority to update the issuance policy and the issuance policy would only expire when the contract has ended).

```protobuf
// MsgInvest is the Msg/Invest request type.
message MsgInvest {

  // id is the unique identifier of the contract.
  uint64 id = 1;

  // funder is the address of the account funding the project and receiving a
  // share of future credits issued from the project.
  string funder = 2;

  // volume_percentage is the percent of all credits issued that the funder will
  // receive.
  string volume_percentage = 3;

  // funds is the token denom and amount the funder is providing in return for
  // the specified volume percentage. The required amount is determined based on
  // the volume percentage provided and only the required amount is sent.
  cosmos.base.v1beta1.Coin funds = 4;
  
  // auto_retire determines whether the credits will be automatically retied upon
  // issuance (i.e. the issuance policy will be set to auto-retire).
  bool auto_retire = 5;
  
  // retirement_jurisdiction is the jurisdiction of the funder. A jurisdiction is
  // only required if auto-retire is enabled.
  string retirement_jurisdiction = 6;
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
Rule: The funds are deducted from the token balance of the funder
```

```feature
Rule: The funds are added to the token balance of the project admin
```

### Contract State

```protobuf
// Contract defines a forward contract and the forward contract table.
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
  string metadata = 4;

  // volume_percentage is the remaining percent of all credits issued that
  // are be available to claim.
  string volume_percentage = 5;

  // funds_to_collect is the denom and remaining amount of funds the project
  // is collecting.
  cosmos.base.v1beta1.Coin funds_to_collect = 6;

  // start_date is the contract start date.
  google.protobuf.Timestamp start_date = 7;

  // end_date is the contract end date.
  google.protobuf.Timestamp end_date = 8;
  
  // buffer_window is the duration after the end date in which credits may
  // still be issued with a monitoring period that falls within the start
  // and end date of the contract. The credits held in the reserve pool are
  // not distributed or returned until end date + buffer window.
  google.protobuf.Duration buffer_window = 9;
}
```

### Contract Reserve State

```protobuf
// ContractReserve defines a forward contract reserve (aka "buffer pool")
// and the table within which forward contract reserves are stored.
message ContractReserve {
  option (cosmos.orm.v1alpha1.table) = {
    id : 2,
    primary_key : {fields : "contract_id"}
  };

  // contract_id is the table row identifier of the contract.
  uint64 contract_id = 1;

  // balances is the list of credit batch balances held in the reserve.
  repeated Balance balances = 2;

  // Balance defines a balance of credits held in the reserve.
  message Balance {

    // batch_denom is the denom of the credits in the reserve.
    string batch_denom = 1;

    // tradable_amount is the amount of the credits in the reserve.
    string tradable_amount = 2;
  }
}
```

### Contract Expiration

The contract will automatically expire once the contracted credit volume has been delivered or when the contract end date (with an optional buffer window) has been reached. If the end date (plus the optional buffer window) is reached before the contracted credit volume has been delivered, credits from the reserve pool will be delivered in place of the contracted credits up to the amount available in the reserve pool. 

### Automated Credit Issuance

Direct credit issuance to the investor(s) will be enforced with on-chain functionality. At the time of funding the project, an investor has the option of choosing to receive credits in a retired or tradable state upon issuance. This functionality will be handled by "issuance policies" and explored as a separate feature set built alongside the initial implementation of forward contract functionality.

<!-- (TODO: replace with issuance policy RFC link) -->
See [(x/ecocredit): issuance policy proof-of-concept][3] for more information.

### Stage 2

Following the initial implementation, additional functionality could be added to support the liquidity of claims on future credits issued from a project enabling investors to receive tradable assets representing such claims.

Investors would have the option of receiving tokens immediately instead of receiving credits directly over time. The tokens would then be used to redeem credits from a claim account (i.e. an account managed programmatically by the `contract` submodule). Receiving tokens would enable investors to transfer their claim and the owner(s) of those tokens would then be able to redeem credits at a time of their choosing.

The amount of tokens sent to the investor upon funding a project would be calculated based on the percent of the claim. The tokens would be specific to the project and the denomination would include information about the most recent credits redeemed. Rather than an issuance policy being created using the investor account as the recipient, an issuance policy would be created using the claim account as the recipient and the credits would be held by the claim account until redeemed by the token owner.

When a token owner redeems credits from the claim account, the owner would exchange the tokens for the credits issued from the project and receive new tokens equal to the amount of tokens used to redeem credits. The denomination of the tokens would include the batch sequence number of the most recent batch from which the credits were redeemed and redeeming credits would burn the tokens sent to the claim account and return tokens with an updated denomination.

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
