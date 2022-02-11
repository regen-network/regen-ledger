package ecocredit

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types/math"
)

const (
	// ModuleName is the module name constant used in many places
	ModuleName = "ecocredit"

	DefaultParamspace = ModuleName

	TradableBalancePrefix    byte = 0x0
	TradableSupplyPrefix     byte = 0x1
	RetiredBalancePrefix     byte = 0x2
	RetiredSupplyPrefix      byte = 0x3
	CreditTypeSeqTablePrefix byte = 0x4
	ClassInfoTablePrefix     byte = 0x5
	BatchInfoTablePrefix     byte = 0x6
	ORMPrefix                byte = 0x7
)

// BatchDenomT is used to prevent errors when forming keys as accounts and denoms are
// both represented as strings
type BatchDenomT string

// - 0x0 <accAddrLen (1 Byte)><accAddr_Bytes><denom_Bytes>: TradableBalance
// - 0x1 <denom_Bytes>: TradableSupply
// - 0x2 <accAddrLen (1 Byte)><accAddr_Bytes><denom_Bytes>: RetiredBalance
// - 0x3 <denom_Bytes>: RetiredSupply

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

// GetDecimal retrieves a decimal by `key` from the given `store`
func GetDecimal(store sdk.KVStore, key []byte) (math.Dec, error) {
	bz := store.Get(key)
	if bz == nil {
		return math.NewDecFromInt64(0), nil
	}

	value, err := math.NewDecFromString(string(bz))
	if err != nil {
		return math.Dec{}, sdkerrors.Wrap(err, fmt.Sprintf("can't unmarshal %s as decimal", bz))
	}

	return value, nil
}

// SetDecimal stores a decimal by `key` in the given `store`
func SetDecimal(store sdk.KVStore, key []byte, value math.Dec) {
	// always remove all trailing zeros for canonical representation
	value, _ = value.Reduce()

	if value.IsZero() {
		store.Delete(key)
	} else {
		// use floating notation here always for canonical representation
		store.Set(key, []byte(value.String()))
	}
}

// AddAndSetDecimal retrieves a decimal from the given key, adds it to x, and saves it.
func AddAndSetDecimal(store sdk.KVStore, key []byte, x math.Dec) error {
	value, err := GetDecimal(store, key)
	if err != nil {
		return err
	}

	value, err = value.Add(x)
	if err != nil {
		return err
	}

	SetDecimal(store, key, value)
	return nil
}

// SubAndSetDecimal retrieves a decimal from the given key, subtracts x from it, and saves it.
func SubAndSetDecimal(store sdk.KVStore, key []byte, x math.Dec) error {
	value, err := GetDecimal(store, key)
	if err != nil {
		return err
	}

	if value.Cmp(x) == -1 {
		return ErrInsufficientFunds
	}

	value, err = math.SafeSubBalance(value, x)
	if err != nil {
		return err
	}

	SetDecimal(store, key, value)
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

// IterateBalances iterates over balances and calls the specified callback function `cb`
func IterateBalances(store sdk.KVStore, storeKey byte, cb func(address, denom, balance string) bool) {
	iter := sdk.KVStorePrefixIterator(store, []byte{storeKey})
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		addr, denom := ParseBalanceKey(iter.Key())
		if cb(addr.String(), string(denom), string(iter.Value())) {
			break
		}
	}
}
