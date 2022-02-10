package basket

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types/math"
)

const (
	TradableBalancePrefix byte = 0x0
)

type BatchDenomT string

// TradableBalanceKey creates the index key for recipient address and batch-denom
func TradableBalanceKey(acc sdk.AccAddress, denom BatchDenomT) []byte {
	key := []byte{TradableBalancePrefix}
	key = append(key, address.MustLengthPrefix(acc)...)
	return append(key, denom...)
}

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
