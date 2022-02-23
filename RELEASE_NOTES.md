## Regen Ledger v3.0.0

Regen Ledger `v3.0.0` adds basket functionality to the ecocredit module.

### Baskets

Regen Ledger `v3.0.0` updates the ecocredit module to include basket functionality, enabling the aggregation of heterogeneous ecosystem service credits into baskets. Credits from different credit classes and batches that meet a defined criteria can be deposited within a basket in exchange for basket tokens. The basket tokens are fully fungible with other tokens from the same basket, and are tracked in the standard bank module, which means these assets will be made visible in wallets like Keplr, and also able to be transferred via IBC to other chains in the Cosmos ecosystem. Each basket token can later be redeemed for an underlying ecocredit from the given basket, and the ecocredits received may be retired by the account receiving them (for offsetting emissions).

Regen Ledger `v3.0.0` includes a scoped-down, minimum-viable basket implementation with the intention to bring IBC compliant carbon credits to the interchain. For more information about the full specification for basket functionality, see the [basket specification](https://github.com/regen-network/regen-ledger/blob/master/rfcs/002-baskets-specification.md).

The MVP version of baskets proposed in Regen Ledger `v3.0.0` differs from the full specification in the RFC above in that:

- `BasketCritera` is restricted to only allow for:
  - A list of credit classes
  - A recency filter represented either as a fixed minimum batch start date, or a rolling recency window (e.g. batch start date must be within the last 6 months)
- Retrieving ecocredits from a basket can only be done via `Take`, not `Pick`. All calls of `Take` will always retrieve the oldest ecocredits first (by batch start date), ensuring the basket flushes out old credits over time.

## Changelog

For a full list of changes since regen-ledger `v2.1.0`, please see the [CHANGELOG.md](./CHANGELOG.md)
