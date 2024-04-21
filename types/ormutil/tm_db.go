package ormutil

import (
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cosmos/cosmos-sdk/orm/types/kv"
	// tdbm "github.com/tendermint/tm-db"
)

func NewStoreAdapter(db dbm.DB) kv.Store {
	return storeAdapter{db}
}

// storeAdapter adapts cometbft/cometbft-db -> cosmos-sdk/tm-db Iterator for ORM kv.Store
type storeAdapter struct {
	dbm.DB
}

func (ta storeAdapter) Iterator(start, end []byte) (kv.Iterator, error) {
	return ta.DB.Iterator(start, end)
}

func (ta storeAdapter) ReverseIterator(start, end []byte) (kv.Iterator, error) {
	return ta.DB.ReverseIterator(start, end)
}

func (ta storeAdapter) Close() error {
	return ta.DB.Close()
}
