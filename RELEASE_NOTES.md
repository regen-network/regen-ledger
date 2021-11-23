## Regen Ledger v2.1.0

### IBC Patch Upgrade

Regen Ledger `v2.1.0` includes an important fix for Regen Mainnet which recently upgraded to Regen Ledger `v2.0.0`.  In the `v2.0.0` upgrade, a bug was introduced that made all new IBC transactions fail to be processed.

This release (`v2.1.0`) hard codes an emergency height-based upgrade which introduces a consensus-breaking change at height `3126912` that resolves the issue. All production validators and full-nodes must update to `v2.1.0` prior to the upgrade height (estimated for Friday Nov 26, 17:00 UTC).

## Changelog

For a full list of changes since regen-ledger v2.0.0, please see the [CHANGELOG.md](./CHANGELOG.md)
