# Msg Service

The ecocredit module provides a message service for interacting with the state of the ecocredit module. The messages are defined in the proto files available on [Buf Schema Registry](https://buf.build/regen/regen-ledger).

For examples on how to interact with state using the `regen` binary, see [regen tx ecocredit](../../commands/regen_tx_ecocredit.md).

## Ecocredit Core

- [Cancel](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.Cancel)
- [CreateBatch](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.CreateBatch)
- [CreateClass](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.CreateClass)
- [CreateProject](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.CreateProject)
- [MintBatchCredits](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.MintBatchCredits)
- [SealBatch](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.SealBatch)
- [Retire](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.Retire)
- [Send](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.Send)

## Basket Submodule

- [Create](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Msg.Create)
- [Put](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Msg.Put)
- [Take](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Msg.Take)

## Marketplace Submodule

- [AllowAskDenom](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.AllowAskDenom)
- [BuyDirect](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.BuyDirect)
- [CancelSellOrder](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.CancelSellOrder)
- [Sell](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.Sell)
- [UpdateSellOrders](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.UpdateSellOrders)
