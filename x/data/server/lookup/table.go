package lookup

type KVStore interface {
	// Get returns nil iff key doesn't exist. Panics on nil key.
	Get(key []byte) []byte

	// Has checks if a key exists. Panics on nil key.
	Has(key []byte) bool

	// Set sets the key. Panics on nil key or value.
	Set(key, value []byte)
}

//func GetIDForValue(store KVStore, value []byte) []byte {
//	hasher := fnv.New64a()
//	hash := hasher.Sum(value)
//
//	// try 32 bit hash
//	id := hash[:4]
//	bz := store.Get(id)
//	if bytes.Equal(value, bz) {
//		return id
//	}
//
//	// try 64 bit hash
//	id = hash
//	bz = store.Get(id)
//	if bytes.Equal(value, bz) {
//		return id
//	}
//
//	// deal with collisions
//	bz = make([]byte, 8+binary.MaxVarintLen64)
//	var i uint64 = 0
//	for {
//		binary.PutUvarint(bz[8:], i)
//		i++
//	}
//}
//
//func GetValueForID(store KVStore, id []byte) []byte {
//
//}
