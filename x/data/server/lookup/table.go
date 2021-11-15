package lookup

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash"
	"hash/fnv"
)

// Table is the interface for interacting with a lookup table.
type Table interface {
	// GetOrCreateID is an idempotent method for creating or retrieving an unique
	// shortened identifier for the provided binary value.
	GetOrCreateID(store KVStore, value []byte) []byte

	// GetID returns the shortened identifier for the provided binary value if
	// it exists or nil.
	GetID(store KVStore, value []byte) []byte

	// GetValue returns the binary data (if any) corresponding to the provided shortened identifier.
	GetValue(store KVStore, id []byte) []byte
}

// NewTable creates a new lookup table for the provided, optional KVStore prefix using default parameters.
// Default parameters are currently set to use the first 4-bytes of the FNV-1a 64-bit, non-cryptographic hash.
// In the case of a collision, more bytes of the hash will be used for disambiguation but this happens
// in a minority of cases except for massively large data sets.
func NewTable(prefix []byte) (Table, error) {
	return NewTableWithOptions(TableOptions{
		Prefix: prefix,
	})
}

// NewTableWithOptions creates a Table with custom options. Most users should just use NewTable
// with the default values.
func NewTableWithOptions(options TableOptions) (Table, error) {
	prefixLen := len(options.Prefix)
	minLength := options.MinLength
	if minLength == 0 {
		minLength = 4
	}

	newHash := options.NewHash
	if newHash == nil {
		newHash = func() hash.Hash {
			return fnv.New64a()
		}
	}

	hashLen := len(newHash().Sum(nil))
	if minLength > hashLen {
		return nil, fmt.Errorf("option MinLength %d is greater than hash length %d", minLength, hashLen)
	}

	bufLen := prefixLen + hashLen + binary.MaxVarintLen64
	initLen := prefixLen + minLength

	return table{
		minLen:    minLength,
		bufLen:    bufLen,
		initLen:   initLen,
		newHash:   newHash,
		prefix:    options.Prefix,
		prefixLen: prefixLen,
		hashLen:   hashLen,
	}, nil
}

// TableOptions is used to specify custom table options and should only be used by advanced users.
type TableOptions struct {
	// NewHash is a function which returns a new hash.Hash instance.
	NewHash func() hash.Hash

	// MinLength is the minimum number of hash bytes that will be used to create a lookup identifier.
	MinLength int

	// Prefix is an optional prefix to be pre-pended to all KVStore keys.
	Prefix []byte
}

type table struct {
	minLen    int
	bufLen    int
	newHash   func() hash.Hash
	prefix    []byte
	prefixLen int
	hashLen   int
	initLen   int
}

// KVSTore is the interface for key-value stores that Tables operate on.
type KVStore interface {
	// Get returns nil iff key doesn't exist. Panics on nil key.
	Get(key []byte) []byte

	// Set sets the key. Panics on nil key or value.
	Set(key, value []byte)
}

func (t table) GetValue(store KVStore, id []byte) []byte {
	buf := make([]byte, t.prefixLen+len(id))
	copy(buf, t.prefix)
	copy(buf[t.prefixLen:], id)
	return store.Get(id)
}

func (t table) GetOrCreateID(store KVStore, value []byte) []byte {
	id, _ := t.getOrCreateID(store, value, true)
	return id
}

func (t table) GetID(store KVStore, value []byte) []byte {
	id, _ := t.getOrCreateID(store, value, false)
	return id
}

func (t table) getOrCreateID(store KVStore, value []byte, create bool) (id []byte, numCollisions int) {
	hasher := t.newHash()
	_, err := hasher.Write(value)
	if err != nil {
		// we panic here because hash.Write returning an error shouldn't happen
		panic(err)
	}
	hashBz := hasher.Sum(nil)

	id = make([]byte, t.initLen, t.bufLen)
	copy(id, t.prefix)
	copy(id[len(t.prefix):], hashBz[:t.minLen])
	// take the first i bytes of hashBz starting with t.minLen and increasing
	// in cases where there are collisions

	for i := t.minLen; ; i++ {
		found, doesntExist := tryGetOrSetIDIfNotFound(store, id, value, create)

		if found {
			return id, i - t.minLen
		}

		if doesntExist {
			return nil, 0
		}

		if i >= t.hashLen {
			break
		}
		id = append(id, hashBz[i])
	}

	// Deal with collisions by appending a varint disambiguation value.
	// Such collisions are almost impossible with good settings, but can
	// happen with a sub-optimal hash function.
	preLen := t.prefixLen + t.hashLen
	for i := uint64(0); ; i++ {
		id = id[:t.bufLen]
		n := binary.PutUvarint(id[preLen:], i)
		id = id[:preLen+n]
		found, doesntExist := tryGetOrSetIDIfNotFound(store, id, value, create)

		if found {
			return id, t.hashLen + int(i) - t.minLen
		}

		if doesntExist {
			return nil, 0
		}
	}
}

func tryGetOrSetIDIfNotFound(store KVStore, id, value []byte, create bool) (found, doesntExist bool) {
	bz := store.Get(id)

	// id doesn't exist yet
	if bz == nil {
		if create {
			store.Set(id, value)
			return true, false
		} else {
			// doesntExist is true in the case when we're just trying to get an
			// ID and not create it - this means there is no ID registered for
			// this value
			return false, true
		}
	}
	return bytes.Equal(value, bz), false
}
