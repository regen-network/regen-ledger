package server

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
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

	if value.Cmp(x) == -1 {
		return ecocredit.ErrInsufficientFunds
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
		stop, err := cb(string(ecocredit.ParseSupplyKey(iter.Key())), string(iter.Value()))
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
		addr, denom := ecocredit.ParseBalanceKey(iter.Key())
		if cb(addr.String(), string(denom), string(iter.Value())) {
			break
		}
	}
}

func verifyCreditBalance(store storetypes.KVStore, ownerAddr sdk.AccAddress, batchDenom string, quantity string) error {
	bd := ecocredit.BatchDenomT(batchDenom)

	balance, err := getDecimal(store, ecocredit.TradableBalanceKey(ownerAddr, bd))
	if err != nil {
		return err
	}

	q, err := math.NewPositiveDecFromString(quantity)
	if err != nil {
		return err
	}

	if balance.Cmp(q) == -1 {
		return ecocredit.ErrInsufficientFunds
	}

	return nil
}

// getCoinNeeded calculates the amount of coin needed for purchasing [quantity] ecocredits at [bidPrice] price
func getCoinNeeded(quantity math.Dec, bidPrice *sdk.Coin) (coinNeeded sdk.Coin, err error) {

	// amount is the amount (in bid denom) of coins necessary for purchasing
	// the exact quantity requested
	amountDec, err := quantity.Mul(math.NewDecFromInt64(bidPrice.Amount.Int64()))
	if err != nil {
		return coinNeeded, err
	}

	// amountNeeded should always be rounded down to the nearest integer number
	// as a buyer cannot spend fractional sdk.Coin (e.g. fractional uregen).
	// The ecocredit quantity sent to the buyer should always be calculated by
	// multiplying the integral amountNeeded by bidPrice as opposed to the quantity sent
	// in the Buy() request.
	amountNeeded, err := amountDec.QuoInteger(math.NewDecFromInt64(1))
	if err != nil {
		return coinNeeded, err
	}

	amountNeededInt64, err := amountNeeded.Int64()
	if err != nil {
		return coinNeeded, err
	}

	return sdk.NewCoin(bidPrice.Denom, sdk.NewInt(amountNeededInt64)), nil
}
