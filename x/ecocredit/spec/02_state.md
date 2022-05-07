# State

The `ecocredit` module uses the `orm` package as an abstraction layer over the `KVStore` that enables the creation of database tables with primary and secondary keys. The tables used within the ecocredit module are defined in the proto files available on [Buf Schema Registry](https://buf.build/regen/regen-ledger).

## Ecocredit Core

<!-- listed alphabetically -->

- [Batch](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Batch)
- [BatchBalance](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.BatchBalance)
- [BatchOrigTx](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.BatchOrigTx)
- [BatchSequence](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.BatchSequence)
- [BatchSupply](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.BatchSupply)
- [Class](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Class)
- [ClassIssuer](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.ClassIssuer)
- [ClassSequence](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.ClassSequence)
- [CreditType](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.CreditType)
- [Project](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Project)
- [ProjectSequence](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.ProjectSequence)

## Basket Submodule

<!-- listed alphabetically -->

- [Basket](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Basket)
- [BasketBalance](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.BasketBalance)
- [BasketClass](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.BasketClass)

## Marketplace Submodule

<!-- listed alphabetically -->

- [AllowedDenom](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.AllowedDenom)
- [Market](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Market)
- [SellOrder](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.SellOrder)
