# Msg Service

The ecocredit module provides a message service for interacting with the state of the ecocredit module. The messages are defined in the proto files available on [Buf Schema Registry](https://buf.build/regen/regen-ledger).

For examples on how to interact with state using the `regen` binary, see [regen tx ecocredit](../../commands/regen_tx_ecocredit.md).

## Ecocredit Core

<!-- listed alphabetically -->

- [Cancel](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.Cancel)
- [CreateBatch](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.CreateBatch)
- [CreateClass](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.CreateClass)
- [CreateProject](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.CreateProject)
- [MintBatchCredits](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.MintBatchCredits)
- [Retire](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.Retire)
- [SealBatch](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.SealBatch)
- [Send](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.Send)
- [UpdateClassAdmin](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.UpdateClassAdmin)
- [UpdateClassIssuers](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.UpdateClassIssuers)
- [UpdateClassMetadata](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.UpdateClassMetadata)
- [UpdateProjectAdmin](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.UpdateProjectAdmin)
- [UpdateProjectMetadata](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.v1#regen.ecocredit.v1.Msg.UpdateProjectMetadata)

## Basket Submodule

<!-- listed alphabetically -->

- [Create](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Msg.Create)
- [Put](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Msg.Put)
- [Take](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.basket.v1#regen.ecocredit.basket.v1.Msg.Take)

## Marketplace Submodule

<!-- listed alphabetically -->

- [BuyDirect](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.BuyDirect)
- [CancelSellOrder](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.CancelSellOrder)
- [Sell](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.Sell)
- [UpdateSellOrders](https://buf.build/regen/regen-ledger/docs/main:regen.ecocredit.marketplace.v1#regen.ecocredit.marketplace.v1.Msg.UpdateSellOrders)
