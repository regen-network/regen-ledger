package server

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/math"
)

func getDecimal(store sdk.KVStore, key []byte) (math.Dec, error) {
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

func setDecimal(store sdk.KVStore, key []byte, value math.Dec) {
	// always remove all trailing zeros for canonical representation
	value, _ = value.Reduce()

	if value.IsZero() {
		store.Delete(key)
	} else {
		// use floating notation here always for canonical representation
		store.Set(key, []byte(value.String()))
	}
}

func addAndSetDecimal(store sdk.KVStore, key []byte, x math.Dec) error {
	value, err := getDecimal(store, key)
	if err != nil {
		return err
	}

	value, err = value.Add(x)
	if err != nil {
		return err
	}

	setDecimal(store, key, value)
	return nil
}

func subAndSetDecimal(store sdk.KVStore, key []byte, x math.Dec) error {
	value, err := getDecimal(store, key)
	if err != nil {
		return err
	}

	value, err = math.SafeSubBalance(value, x)
	if err != nil {
		return err
	}

	setDecimal(store, key, value)
	return nil
}

func iterateSupplies(store sdk.KVStore, storeKey byte, cb func(denom, supply string) (bool, error)) error {
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

func iterateBalances(store sdk.KVStore, storeKey byte, cb func(address, denom, balance string) bool) {
	iter := sdk.KVStorePrefixIterator(store, []byte{storeKey})
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		addr, denom := ParseBalanceKey(iter.Key())
		if cb(addr.String(), string(denom), string(iter.Value())) {
			break
		}
	}
}
