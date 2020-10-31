package server

import (
	"fmt"
	"github.com/cockroachdb/apd/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (s serverImpl) getDec(store sdk.KVStore, key []byte) (*apd.Decimal, error) {
	bz := store.Get(key)

	value, _, err := apd.NewFromString(string(bz))
	if err != nil {
		return nil, sdkerrors.Wrap(err, "can't unmarshal decimal")
	}

	return value, nil
}

func (s serverImpl) setDec(store sdk.KVStore, key []byte, value *apd.Decimal) {
	store.Set(key, []byte(value.String()))
}

// IEEE 754-2008 decimal128
var dec128Context = apd.Context{
	Precision:   34,
	MaxExponent: 6144,
	MinExponent: -6143,
	Traps:       apd.DefaultTraps,
	Rounding:    apd.RoundHalfEven,
}

func (s serverImpl) addDec(store sdk.KVStore, key []byte, x *apd.Decimal) error {
	value, err := s.getDec(store, key)
	if err != nil {
		return err
	}

	err = add(value, value, x)
	if err != nil {
		return err
	}

	s.setDec(store, key, value)
	return nil
}

func (s serverImpl) safeSubDec(store sdk.KVStore, key []byte, x *apd.Decimal) error {
	value, err := s.getDec(store, key)
	if err != nil {
		return err
	}

	_, err = dec128Context.Sub(value, value, x)
	if err != nil {
		return err
	}

	if isNegative(x) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "result is negative")
	}

	s.setDec(store, key, value)
	return nil
}

func add(res, x, y *apd.Decimal) error {
	_, err := dec128Context.Add(res, x, y)
	return err
}

func isPositive(x *apd.Decimal) bool {
	return x.Sign() > 0 && !x.IsZero()
}

func isNegative(x *apd.Decimal) bool {
	return x.Sign() < 0 && !x.IsZero()
}

func requirePositive(x *apd.Decimal) error {
	if !isPositive(x) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s is not positive", x))
	}
	return nil
}
