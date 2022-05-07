# Query Service

The ecocredit module provides a query service for querying the state of the ecocredit module. The queries are defined in the proto files available on [Buf Schema Registry](https://buf.build/regen/regen-ledger).

For examples on how to query state using gRPC and REST, see the [ecocredit client](07_client.md) documentation, and for examples using the `regen` binary, see [regen query ecocredit](../../commands/regen_query_ecocredit.html).

## Ecocredit Core

<!-- listed alphabetically -->

- [Balance](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Balance)
- [Balances](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Balances)
- [Batch](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Batch)
- [Batches](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Batches)
- [BatchesByClass](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.BatchesByClass)
- [BatchesByIssuer](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.BatchesByIssuer)
- [Class](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Class)
- [Classes](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Classes)
- [ClassesByAdmin](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.ClassesByAdmin)
- [ClassIssuers](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.ClassIssuers)
- [CreditTypes](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.CreditTypes)
- [Params](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Params)
- [Project](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Project)
- [Projects](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Projects)
- [Supply](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Supply)

## Basket Submodule

<!-- listed alphabetically -->

- [Basket](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Query.Basket)
- [BasketBalance](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Query.BasketBalance)
- [BasketBalances](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Query.BasketBalances)
- [Baskets](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Query.Baskets)

## Marketplace Submodule

<!-- listed alphabetically -->

- [AllowedDenoms](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Query.AllowedDenoms)
- [SellOrder](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Query.SellOrder)
- [SellOrders](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Query.SellOrders)
- [SellOrdersByAddress](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Query.SellOrdersByAddress)
- [SellOrdersByBatchDenom](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Query.SellOrdersByBatchDenom)
