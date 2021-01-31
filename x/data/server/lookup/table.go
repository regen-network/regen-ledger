package lookup

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
)

type KVStore interface {
	// Get returns nil iff key doesn't exist. Panics on nil key.
	Get(key []byte) []byte

	// Set sets the key. Panics on nil key or value.
	Set(key, value []byte)
}

func GetOrCreateIDForValue(store KVStore, value []byte) []byte {
	id, _ := getOrCreateIDForValue(store, value, 4, 8)
	return id
}

func getOrCreateIDForValue(store KVStore, value []byte, lo, hi int) (id []byte, numCollisions uint64) {
	hasher := fnv.New64a()
	_, err := hasher.Write(value)
	if err != nil {
		panic(err)
	}
	hash := hasher.Sum(nil)

	// try 32 bit hash
	id = hash[:lo]
	if tryId(store, id, value) {
		return id, 0
	}

	// try 64 bit hash
	id = hash[:hi]
	if tryId(store, id, value) {
		return id, 1
	}

	// deal with collisions
	idBz := make([]byte, 8+binary.MaxVarintLen64)
	copy(idBz[:8], hash)
	var i uint64 = 0
	for {
		n := binary.PutUvarint(idBz[8:], i)
		id = idBz[:8+n]
		if tryId(store, id, value) {
			return id, i + 2
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
