# Msg Service

The ecocredit module provides a message service for interacting with the state of the ecocredit module. The messages are defined in the proto files available on [Buf Schema Registry](https://buf.build/regen/regen-ledger).

For examples on how to interact with state using the `regen` binary, see [regen tx ecocredit](../../commands/regen_tx_ecocredit.md).

## Ecocredit Core

- [CreateClass](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.v1#CreateClass)
- [CreateProject](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.v1#CreateProject)
- [CreateBatch](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.v1#CreateBatch)
- [Send](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.v1#Send)
- [Retire](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.v1#Retire)
- [Cancel](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.v1#Cancel)

## Basket Submodule

- [Create](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.basket.v1#Create)
- [Put](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.basket.v1#Put)
- [Take](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.basket.v1#Take)

## Marketplace Submodule

- [Sell](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.marketplace.v1#Sell)
- [UpdateSellOrders](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.marketplace.v1#UpdateSellOrders)
- [Buy](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.marketplace.v1#Buy)
- [AllowAskDenom](https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.marketplace.v1#AllowAskDenom)
