package orm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// IndexerFunc creates one or multiple index keys for the source object.
type IndexerFunc func(value interface{}) ([]interface{}, error)

// IndexerFunc creates exactly one index key for the source object.
type UniqueIndexerFunc func(value interface{}) (interface{}, error)

// Indexer manages the persistence for an Index based on searchable keys and operations.
type Indexer struct {
	indexerFunc IndexerFunc
	addFunc     func(store sdk.KVStore, secondaryIndexKey []byte, rowID RowID) error
}

// NewIndexer returns an indexer that supports multiple reference keys for an entity.
func NewIndexer(indexerFunc IndexerFunc) (*Indexer, error) {
	if indexerFunc == nil {
		return nil, ErrArgument.Wrap("Indexer func must not be nil")
	}
	return &Indexer{
		indexerFunc: pruneEmptyKeys(indexerFunc),
		addFunc:     multiKeyAddFunc,
	}, nil
}

// NewUniqueIndexer returns an indexer that requires exactly one reference keys for an entity.
func NewUniqueIndexer(f UniqueIndexerFunc) (*Indexer, error) {
	if f == nil {
		return nil, ErrArgument.Wrap("Indexer func must not be nil")
	}
	adaptor := func(indexerFunc UniqueIndexerFunc) IndexerFunc {
		return func(v interface{}) ([]RowID, error) {
			k, err := indexerFunc(v)
			return []RowID{k}, err
		}
	}
	idx, err := NewIndexer(adaptor(f))
	if err != nil {
		return nil, err
	}
	idx.addFunc = uniqueKeysAddFunc
	return idx, nil
}

// OnCreate persists the secondary index entries for the new object.
func (i Indexer) OnCreate(store sdk.KVStore, rowID RowID, value interface{}) error {
	secondaryIndexKeys, err := i.indexerFunc(value)
	if err != nil {
		return err
	}

	for _, secondaryIndexKey := range secondaryIndexKeys {
		if err := i.addFunc(store, secondaryIndexKey, rowID); err != nil {
			return err
		}
	}
	return nil
}

// OnDelete removes the secondary index entries for the deleted object.
func (i Indexer) OnDelete(store sdk.KVStore, rowID RowID, value interface{}) error {
	secondaryIndexKeys, err := i.indexerFunc(value)
	if err != nil {
		return err
	}

	for _, secondaryIndexKey := range secondaryIndexKeys {
		indexKey, err := buildKeyFromParts([]interface{}{secondaryIndexKey, rowID})
		if err != nil {
			return err
		}
		store.Delete(indexKey)
	}
	return nil
}

// OnUpdate rebuilds the secondary index entries for the updated object.
func (i Indexer) OnUpdate(store sdk.KVStore, rowID RowID, newValue, oldValue interface{}) error {
	oldSecIdxKeys, err := i.indexerFunc(oldValue)
	if err != nil {
		return err
	}
	newSecIdxKeys, err := i.indexerFunc(newValue)
	if err != nil {
		return err
	}
	for _, oldIdxKey := range difference(oldSecIdxKeys, newSecIdxKeys) {
		indexKey, err := i.indexKeyCodec.BuildIndexKey(oldIdxKey, rowID)
		if err != nil {
			return err
		}
		store.Delete(indexKey)
	}
	for _, newIdxKey := range difference(newSecIdxKeys, oldSecIdxKeys) {
		if err := i.addFunc(store, i.indexKeyCodec, newIdxKey, rowID); err != nil {
			return err
		}
	}
	return nil
}

// uniqueKeysAddFunc enforces keys to be unique
func uniqueKeysAddFunc(store sdk.KVStore, codec IndexKeyCodec, secondaryIndexKey []byte, rowID RowID) error {
	if len(secondaryIndexKey) == 0 {
		return errors.Wrap(ErrArgument, "empty index key")
	}
	it := store.Iterator(PrefixRange(codec.PrefixSearchableKey(secondaryIndexKey)))
	defer it.Close()
	if it.Valid() {
		return ErrUniqueConstraint
	}

	indexKey, err := codec.BuildIndexKey(secondaryIndexKey, rowID)
	if err != nil {
		return err
	}

	store.Set(indexKey, []byte{})
	return nil
}

// multiKeyAddFunc allows multiple entries for a key
func multiKeyAddFunc(store sdk.KVStore, codec IndexKeyCodec, secondaryIndexKey []byte, rowID RowID) error {
	if len(secondaryIndexKey) == 0 {
		return errors.Wrap(ErrArgument, "empty index key")
	}

	indexKey, err := codec.BuildIndexKey(secondaryIndexKey, rowID)
	if err != nil {
		return err
	}

	store.Set(indexKey, []byte{})
	return nil
}

// difference returns the list of elements that are in a but not in b.
func difference(a []RowID, b []RowID) []RowID {
	set := make(map[string]struct{}, len(b))
	for _, v := range b {
		set[string(v)] = struct{}{}
	}
	var result []RowID
	for _, v := range a {
		if _, ok := set[string(v)]; !ok {
			result = append(result, v)
		}
	}
	return result
}

// pruneEmptyKeys drops any empty key from IndexerFunc f returned
func pruneEmptyKeys(f IndexerFunc) IndexerFunc {
	return func(v interface{}) ([]interface{}, error) {
		keys, err := f(v)
		if err != nil || keys == nil {
			return keys, err
		}
		r := make([]interface{}, 0, len(keys))
		for i := range keys {
			key, err := keyPartBytes(keys[i], true)
			if err != nil {
				return nil, err
			}
			if len(key) != 0 {
				r = append(r, keys[i])
			}
		}
		return r, nil
	}
}
