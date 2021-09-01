# State

The `ecocrdit` module uses the `KVStore` directly for credit batch balances and supplies and the `orm` package for table storage of credit type sequences, credit classes, and credit batches.

## Tradable Balance

`TradableBalance` is the tradable balance of a credit batch.

`TradableBalance` is stored directly in the `KVStore`:

`0x0 | byte(address, denom) --> byte(amount)`

## Tradable Supply

`TradableSupply` is the tradable supply of a credit batch.

`TradableSupply` is stored directly in the `KVStore`:

`0x1 | byte(denom) --> byte(amount)`

## Retired Balance

`RetiredBalance` is the retired balance of a credit batch.

`RetiredBalance` is stored directly in the `KVStore`:

`0x2 | byte(address, denom) --> byte(amount)`

## Retired Supply

`RetiredSupply` is the retired supply of a credit batch.

`RetiredSupply` is stored directly in the `KVStore`:

`0x3 | byte(denom) --> byte(amount)`

## Credit Type Sequence Table

`CreditTypeSeq` associates a sequence number with a credit type abbreviation.

The sequence number is incremented on a per credit type basis, and it exists for the purpose of providing a sequence number for the credit class ID. A credit class ID is the combination of a credit type abbreviation and a sequence number (e.g. `C1` is the ID for the first "carbon" credit class).

The `creditTypeSeqTable` stores `CreditTypeSeq`:

`0x4 | []byte(Abbreviation) -> ProtocolBuffer(CreditTypeSeq)`

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/types.proto#L114-122

## Class Info Table

The `classInfoTable` stores `ClassInfo`:

`0x5 | []byte(ClassId) -> ProtocolBuffer(ClassInfo)`

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/types.proto#L11-31

## Batch Info Table

The `batchInfoTable` stores `BatchInfo`:

`0x6 | []byte(BatchId) -> ProtocolBuffer(BatchInfo)`

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/types.proto#L33-74
