# State

The `ecocredit` module uses the Cosmos SDK's `KVStore` directly for credit batch balances and supplies and the `orm` package (an abstraction over the `KVStore`) for table storage of credit type sequences, credit classes, and credit batches.

## Tradable Balance

`TradableBalance` is the tradable balance of a credit batch.

`TradableBalance` is stored directly in the `KVStore`:

`0x0 | byte(address length) | []byte(address) | []byte(denom) --> []byte(amount)`

## Tradable Supply

`TradableSupply` is the tradable supply of a credit batch.

`TradableSupply` is stored directly in the `KVStore`:

`0x1 | []byte(denom) --> []byte(amount)`

## Retired Balance

`RetiredBalance` is the retired balance of a credit batch.

`RetiredBalance` is stored directly in the `KVStore`:

`0x2 | byte(address length) | []byte(address) | []byte(denom) --> []byte(amount)`

## Retired Supply

`RetiredSupply` is the retired supply of a credit batch.

`RetiredSupply` is stored directly in the `KVStore`:

`0x3 | []byte(denom) --> []byte(amount)`

## Credit Type Sequence Table

`CreditTypeSeq` associates a sequence number with a credit type abbreviation.

The sequence number is incremented on a per credit type basis, and it exists for the purpose of providing a sequence number for the credit class ID. A credit class ID is the combination of a credit type abbreviation and a sequence number (e.g. `C01` is the ID for the first "carbon" credit class).

The `creditTypeSeqTable` stores `CreditTypeSeq`:

`0x4 | []byte(Abbreviation) -> ProtocolBuffer(CreditTypeSeq)`

+++ https://github.com/regen-network/regen-ledger/blob/50eaceda5eabc5970effe491f0d58194852718c9/proto/regen/ecocredit/v1alpha1/types.proto#L112-L122

## Class Info Table

The `classInfoTable` stores `ClassInfo`:

`0x5 | []byte(ClassId) -> ProtocolBuffer(ClassInfo)`

+++ https://github.com/regen-network/regen-ledger/blob/50eaceda5eabc5970effe491f0d58194852718c9/proto/regen/ecocredit/v1alpha1/types.proto#L11-L31

## Batch Info Table

The `batchInfoTable` stores `BatchInfo`:

`0x6 | []byte(BatchId) -> ProtocolBuffer(BatchInfo)`

+++ https://github.com/regen-network/regen-ledger/blob/50eaceda5eabc5970effe491f0d58194852718c9/proto/regen/ecocredit/v1alpha1/types.proto#L33-L72

## Project Info Table

The `projectInfoTable` stores `ProjectInfo`:

`0x10 | []byte(ProjectId) -> ProtocolBuffer(ProjectInfo)`

+++ https://github.com/regen-network/regen-ledger/blob/50eaceda5eabc5970effe491f0d58194852718c9/proto/regen/ecocredit/v1alpha1/types.proto#L33-L72

## Sell Order Table

The `sellOrderTable` stores `SellOrder`:

`0x20 | []byte(OrderId) -> ProtocolBuffer(SellOrder)`

+++ https://github.com/regen-network/regen-ledger/blob/081ae071b159b397b4c10837804b69137295e3af/proto/regen/ecocredit/v1alpha1/types.proto#L122-L146

#### Sell Order Sequence Table

The `sellOrderTable` uses a persistent unique key generator called `Sequence`:

`OrderId`: `0x21 | 0x1 -> BigEndian`

The `0x1` is a fixed key to read/write data to the storage layer.

## Buy Order Table

The `buyOrderTable` stores `BuyOrder`:

`0x25 | []byte(BuyOrderId) -> ProtocolBuffer(BuyOrder)`

+++ https://github.com/regen-network/regen-ledger/blob/081ae071b159b397b4c10837804b69137295e3af/proto/regen/ecocredit/v1alpha1/types.proto#L148-L196

#### Buy Order Sequence Table

The `buyOrderTable` uses a persistent unique key generator called `Sequence`:

`BuyOrderId`: `0x26 | 0x1 -> BigEndian`

The `0x1` is a fixed key to read/write data to the storage layer.

## Ask Denom Table

The `askDenomTable` stores `AskDenom`:

`0x30 | []byte(Denom) -> ProtocolBuffer(AskDenom)`

+++ https://github.com/regen-network/regen-ledger/blob/081ae071b159b397b4c10837804b69137295e3af/proto/regen/ecocredit/v1alpha1/types.proto#L198-L210
