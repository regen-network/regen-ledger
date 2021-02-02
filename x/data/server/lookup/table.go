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

	return table{
		minLen:    minLength,
		bufLen:    bufLen,
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
	id, _ := t.getOrCreateID(store, value)
	return id
}

func (t table) getOrCreateID(store KVStore, value []byte) (id []byte, numCollisions int) {
	hasher := t.newHash()
	_, err := hasher.Write(value)
	if err != nil {
		panic(err)
	}
	hashBz := hasher.Sum(nil)

	id = make([]byte, 0, t.bufLen)
	id = append(id, t.prefix...)

	for i := t.minLen; i <= t.hashLen; i++ {
		id = append(id[t.prefixLen:], hashBz[:i]...)
		if tryID(store, id, value) {
			return id, i - t.minLen
		}
		id = id[:t.prefixLen]
	}

	// deal with collisions which are almost impossible with good settings, but can happen with a sub-optimal hash function
	var i uint64 = 0
	preLen := t.prefixLen + t.hashLen
	for {
		id = id[:t.bufLen]
		n := binary.PutUvarint(id[preLen:], i)
		id = id[:preLen+n]
		if tryID(store, id, value) {
			return id, t.hashLen + int(i) - t.minLen
		}

		i++
	}
}

func tryID(store KVStore, id []byte, value []byte) bool {
	bz := store.Get(id)

	// id doesn't exist yet
	if len(bz) == 0 {
		store.Set(id, value)
		return true
	}

	// id exists, check if equal
	if bytes.Equal(value, bz) {
		return true
	}

	return false
}
