package ormstore

import storetypes "github.com/cosmos/cosmos-sdk/types"
import "github.com/cosmos/cosmos-sdk/orm/model/kv"

type storeWrapper struct {
	store storetypes.KVStore
}

func (k storeWrapper) Set(key, value []byte) error {
	k.store.Set(key, value)
	return nil
}

func (k storeWrapper) Delete(key []byte) error {
	k.store.Delete(key)
	return nil
}

func (k storeWrapper) Get(key []byte) ([]byte, error) {
	x := k.store.Get(key)
	return x, nil
}

func (k storeWrapper) Has(key []byte) (bool, error) {
	x := k.store.Has(key)
	return x, nil
}

func (k storeWrapper) Iterator(start, end []byte) (kv.Iterator, error) {
	x := k.store.Iterator(start, end)
	return x, nil
}

func (k storeWrapper) ReverseIterator(start, end []byte) (kv.Iterator, error) {
	x := k.store.ReverseIterator(start, end)
	return x, nil
}

var _ kv.Store = &storeWrapper{}
