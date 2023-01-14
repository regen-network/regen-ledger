# Query Service

The ecocredit module provides a query service for querying the state of the ecocredit module. The queries are defined in the proto files available on [Buf Schema Registry](https://buf.build/regen/regen-ledger).

For examples on how to query state using gRPC and REST, see the [ecocredit client](07_client.md) documentation, and for examples using the `regen` binary, see [regen query ecocredit](../../commands/regen_query_ecocredit.html).

## Ecocredit Core

<!-- listed alphabetically -->

- [AllBalances](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.AllBalances)
- [AllowedBridgeChains](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.AllowedBridgeChains)
- [AllowedClassCreators](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.AllowedClassCreators)
- [Balance](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Balance)
- [Balances](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Balances)
- [BalancesByBatch](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.BalancesByBatch)
- [Batch](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Batch)
- [Batches](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Batches)
- [BatchesByClass](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.BatchesByClass)
- [BatchesByIssuer](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.BatchesByIssuer)
- [BatchesByProject](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.BatchesByProject)
- [Class](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Class)
- [ClassCreatorAllowlist](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.ClassCreatorAllowlist)
- [Classes](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Classes)
- [ClassesByAdmin](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.ClassesByAdmin)
- [ClassFee](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.ClassFee)
- [ClassIssuers](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.ClassIssuers)
- [CreditType](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.CreditType)
- [CreditTypes](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.CreditTypes)
- [Params](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Params)
- [Project](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Project)
- [Projects](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Projects)
- [ProjectsByAdmin](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.ProjectsByAdmin)
- [ProjectsByClass](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.ProjectsByClass)
- [ProjectsByReferenceId](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.ProjectsByReferenceId)
- [Supply](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Query.Supply)

## Basket Submodule

<!-- listed alphabetically -->

- [Basket](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Query.Basket)
- [BasketBalance](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Query.BasketBalance)
- [BasketBalances](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Query.BasketBalances)
- [BasketFee](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Query.BasketFee)
- [Baskets](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Query.Baskets)

## Marketplace Submodule

<!-- listed alphabetically -->

- [AllowedDenoms](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Query.AllowedDenoms)
- [SellOrder](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Query.SellOrder)
- [SellOrders](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Query.SellOrders)
- [SellOrdersByBatch](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Query.SellOrdersByBatch)
- [SellOrdersBySeller](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Query.SellOrdersBySeller)
