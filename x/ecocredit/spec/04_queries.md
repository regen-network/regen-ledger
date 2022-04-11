# Query Service

The ecocredit module provides a query service for querying the state of the ecocredit module. The queries are defined in the proto files available on [Buf Schema Registry](https://buf.build/regen/regen-ledger).

For examples on how to query state using gRPC and REST, see the [ecocredit client](06_client.md) documentation, and for examples using the `regen` binary, see [regen query ecocredit](../../commands/regen_query_ecocredit.html).

## Ecocredit Core

- [Classes](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#Classes)
- [ClassInfo](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#ClassInfo)
- [ClassIssuers](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#ClassIssuers)
- [Projects](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#Projects)
- [ProjectInfo](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#ProjectInfo)
- [Batches](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#Batches)
- [BatchesByClass](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#BatchesByClass)
- [BatchInfo](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#BatchInfo)
- [Balance](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#Balance)
- [Supply](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#Supply)
- [CreditTypes](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#CreditTypes)
- [Params](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#Params)

## Basket Submodule

- [Basket](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#Basket)
- [Baskets](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#Baskets)
- [BasketBalance](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#BasketBalance)
- [BasketBalances](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#BasketBalances)

## Marketplace Submodule

- [SellOrder](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#SellOrder)
- [SellOrders](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#SellOrders)
- [SellOrdersByBatchDenom](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#SellOrdersByBatchDenom)
- [SellOrdersByAddress](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#SellOrdersByAddress)
- [BuyOrder](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#BuyOrder)
- [BuyOrders](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#BuyOrders)
- [BuyOrdersByAddress](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#BuyOrdersByAddress)
- [AllowedDenoms](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#AllowedDenoms)
