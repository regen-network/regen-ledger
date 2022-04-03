# State

The `ecocredit` module uses the `orm` package as an abstraction layer over the `KVStore` that enables the creation of database tables with primary and secondary keys. The tables used within the ecocredit module are defined in the proto files available on [Buf Schema Registry](https://buf.build/regen/regen-ledger).

## Ecocredit Core

- [BatchBalance](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.BatchBalance)
- [BatchInfo](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.BatchInfo)
- [BatchSequence](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.BatchSequence)
- [BatchSupply](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.BatchSupply)
- [ClassInfo](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.ClassInfo)
- [ClassIssuer](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.ClassIssuer)
- [ClassSequence](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.ClassSequence)
- [CreditType](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.CreditType)

## Basket Submodule

- [Basket](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Basket)
- [BasketBalance](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.BasketBalance)
- [BasketCredit](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.BasketCredit)
- [DateCriteria](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.DateCriteria)

## Marketplace Submodule

- [AllowedDenom](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.AllowedDenom)
- [BatchSelector](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.BatchSelector)
- [BuyOrder](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.BuyOrder)
- [BuyOrder.Selection](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.BuyOrder.Selection)
- [ClassSelector](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.ClassSelector)
