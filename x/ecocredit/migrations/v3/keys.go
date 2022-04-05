package v3

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	TradableBalancePrefix    byte = 0x0
	TradableSupplyPrefix     byte = 0x1
	RetiredBalancePrefix     byte = 0x2
	RetiredSupplyPrefix      byte = 0x3
	CreditTypeSeqTablePrefix byte = 0x4
	ClassInfoTablePrefix     byte = 0x5
	BatchInfoTablePrefix     byte = 0x6
)

// TradableBalanceKey creates the index key for recipient address and batch-denom
func TradableBalanceKey(acc sdk.AccAddress, denom BatchDenomT) []byte {
	key := []byte{TradableBalancePrefix}
	key = append(key, address.MustLengthPrefix(acc)...)
	return append(key, denom...)
}

// ParseBalanceKey parses the recipient address and batch-denom from tradable or retired balance key.
// Balance keys take the following form: <storage prefix 1 byte><addr length 1 byte><addr><batchDenom>
func ParseBalanceKey(key []byte) (sdk.AccAddress, BatchDenomT) {
	addrLen := key[1]
	addr := sdk.AccAddress(key[2 : 2+addrLen])
	return addr, BatchDenomT(key[2+addrLen:])
}

// TradableSupplyKey creates the tradable supply key for a given batch-denom
func TradableSupplyKey(batchDenom BatchDenomT) []byte {
	key := []byte{TradableSupplyPrefix}
	return append(key, batchDenom...)
}

// ParseSupplyKey parses the batch-denom from tradable or retired supply key
func ParseSupplyKey(key []byte) BatchDenomT {
	return BatchDenomT(key[1:])
}

// RetiredBalanceKey creates the index key for recipient address and batch-denom
func RetiredBalanceKey(acc sdk.AccAddress, batchDenom BatchDenomT) []byte {
	key := []byte{RetiredBalancePrefix}
	key = append(key, address.MustLengthPrefix(acc)...)
	return append(key, batchDenom...)
}

// RetiredSupplyKey creates the retired supply key for a given batch-denom
func RetiredSupplyKey(batchDenom BatchDenomT) []byte {
	key := []byte{RetiredSupplyPrefix}
	return append(key, batchDenom...)
}

// IterateBalances iterates over balances and calls the specified callback function `cb`
func IterateBalances(store sdk.KVStore, storeKey byte, cb func(address, denom, balance string) (bool, error)) error {
	iter := sdk.KVStorePrefixIterator(store, []byte{storeKey})
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		addr, denom := ParseBalanceKey(iter.Key())
		stop, err := cb(addr.String(), string(denom), string(iter.Value()))
		if err != nil {
			return err
		}

		if stop {
			break
		}
	}

	return nil
}

// IterateSupplies iterates over supplies and calls the specified callback function `cb`
func IterateSupplies(store sdk.KVStore, storeKey byte, cb func(denom, supply string) (bool, error)) error {
	iter := sdk.KVStorePrefixIterator(store, []byte{storeKey})
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		stop, err := cb(string(ParseSupplyKey(iter.Key())), string(iter.Value()))
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}

	return nil
}

// Calculate the ID to use for a new project, based on the class id and
// the project sequence number.
//
// The initial version has format:
// <class id><project seq no>
func FormatProjectID(classID string, projectSeqNo uint64) string {
	return fmt.Sprintf("%s%02d", classID, projectSeqNo)
}
