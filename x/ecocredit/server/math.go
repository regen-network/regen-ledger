package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/cockroachdb/apd/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func storeGetDec(store sdk.KVStore, key []byte) (*apd.Decimal, error) {
	bz := store.Get(key)
	if bz == nil {
		return apd.New(0, 0), nil
	}

	value, _, err := apd.NewFromString(string(bz))
	if err != nil {
		return nil, sdkerrors.Wrap(err, fmt.Sprintf("can't unmarshal %s as decimal", bz))
	}

	return value, nil
}

func storeSetDec(store sdk.KVStore, key []byte, value *apd.Decimal) {
	// always remove all trailing zeros for canonical representation
	value, _ = value.Reduce(value)
	// use floating notation here always for canonical representation
	str := value.Text('f')
	store.Set(key, []byte(str))
}

func storeAddDec(store sdk.KVStore, key []byte, x *apd.Decimal) error {
	value, err := storeGetDec(store, key)
	if err != nil {
		return err
	}

	err = add(value, value, x)
	if err != nil {
		return err
	}

	storeSetDec(store, key, value)
	return nil
}

func storeSafeSubDec(store sdk.KVStore, key []byte, x *apd.Decimal) error {
	value, err := storeGetDec(store, key)
	if err != nil {
		return err
	}

	err = safeSub(value, value, x)
	if err != nil {
		return err
	}

	storeSetDec(store, key, value)
	return nil
}

var exactContext = apd.Context{
	Precision:   0,
	MaxExponent: apd.MaxExponent,
	MinExponent: apd.MinExponent,
	Traps:       apd.DefaultTraps | apd.Inexact | apd.Rounded,
}

func add(res, x, y *apd.Decimal) error {
	_, err := exactContext.Add(res, x, y)
	if err != nil {
		return sdkerrors.Wrap(err, "decimal addition error")
	}
	return nil
}

func safeSub(res, x, y *apd.Decimal) error {
	_, err := exactContext.Sub(res, x, y)
	if err != nil {
		return sdkerrors.Wrap(err, "decimal subtraction error")
	}

	if res.Sign() < 0 {
		return sdkerrors.ErrInsufficientFunds
	}

	return nil
}

func storeSetUInt32(store sdk.KVStore, key []byte, value uint32) error {
	bz := make([]byte, 0, 4)
	buf := bytes.NewBuffer(bz)
	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return err
	}

	store.Set(key, buf.Bytes())
	return nil
}

func storeGetUInt32(store sdk.KVStore, key []byte) (uint32, error) {
	bz := store.Get(key)
	if bz == nil {
		return 0, nil
	}

	buf := bytes.NewReader(bz)
	var res uint32
	err := binary.Read(buf, binary.LittleEndian, &res)
	if err != nil {
		return 0, err
	}

	return res, nil
}
