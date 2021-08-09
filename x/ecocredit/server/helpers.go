package server

import (
	"bytes"
	"encoding/binary"
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

func getAddAndSetDecimal(store sdk.KVStore, key []byte, x math.Dec) error {
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

func getSubAndSetDecimal(store sdk.KVStore, key []byte, x math.Dec) error {
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

func setUInt32(store sdk.KVStore, key []byte, value uint32) error {
	bz := make([]byte, 0, 4)
	buf := bytes.NewBuffer(bz)
	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return err
	}

	store.Set(key, buf.Bytes())
	return nil
}

func getUint32(store sdk.KVStore, key []byte) (uint32, error) {
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
