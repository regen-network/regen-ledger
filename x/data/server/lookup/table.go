package lookup

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash"
	"hash/fnv"
)

type Table interface {
	GetOrCreateID(store KVStore, value []byte) []byte
	GetValue(store KVStore, id []byte) []byte
}

func NewTable(prefix []byte) (Table, error) {
	return NewTableWithOptions(TableOptions{
		Prefix: prefix,
	})
}

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

type TableOptions struct {
	NewHash   func() hash.Hash
	MinLength int
	Prefix    []byte
}

type table struct {
	minLen    int
	bufLen    int
	newHash   func() hash.Hash
	prefix    []byte
	prefixLen int
	hashLen   int
}

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

	for i := t.minLen; i < t.hashLen; i++ {
		id = append(id[t.prefixLen:], hashBz[:i]...)
		if tryId(store, id, value) {
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
		if tryId(store, id, value) {
			return id, t.hashLen + n - t.minLen
		}

		i++
	}
}

func tryId(store KVStore, id []byte, value []byte) bool {
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
