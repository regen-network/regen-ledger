package hasher

import (
	"encoding/binary"
	"fmt"
	"hash"
	"hash/fnv"
)

// Hasher generates a unique binary identifier for a longer piece of binary data
// using an efficient, non-cryptographic hash function.
type Hasher interface {
	// CreateID is an idempotent method for creating a unique shortened identifier
	// for the provided binary value.
	CreateID(value []byte, collisions int) []byte
}

// NewHasher creates a new hasher instance. Default parameters are currently set to use the first
// 4-bytes of the FNV-1a 64-bit, non-cryptographic hash. In the case of a collision, more bytes
// of the hash will be used for disambiguation but this happens in a minority of cases except
// for massively large data sets.
func NewHasher(prefix []byte) (Hasher, error) {
	return NewHasherWithOptions(HashOptions{
		Prefix: prefix,
	})
}

// NewHasherWithOptions creates a Hash with custom options. Most users should just use NewHasher
// with the default values.
func NewHasherWithOptions(options HashOptions) (Hasher, error) {
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

	return hasher{
		minLen:    minLength,
		bufLen:    bufLen,
		initLen:   initLen,
		newHash:   newHash,
		prefix:    options.Prefix,
		prefixLen: prefixLen,
		hashLen:   hashLen,
	}, nil
}

// HashOptions is used to specify custom hash options and should only be used by advanced users.
type HashOptions struct {
	// NewHash is a function which returns a new hash.Hash instance.
	NewHash func() hash.Hash

	// MinLength is the minimum number of hash bytes that will be used to create a lookup identifier.
	MinLength int

	// Prefix is an optional prefix to be pre-pended to all keys.
	Prefix []byte
}

type hasher struct {
	minLen    int
	bufLen    int
	newHash   func() hash.Hash
	prefix    []byte
	prefixLen int
	hashLen   int
	initLen   int
}

func (t hasher) CreateID(value []byte, collisions int) []byte {
	id := t.createID(value, collisions)
	return id
}

func (t hasher) createID(value []byte, collisions int) (id []byte) {
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

	// Deal with collisions by appending the equivalent number of bytes
	// from hashBz. If using this method will exceed hash length, append
	// a disambiguation varint. Such collisions are almost impossible with
	// good settings, but can happen with a suboptimal hash function.
	if t.minLen+collisions < t.hashLen {
		id = append(id, hashBz[collisions])
	} else {
		preLen := t.prefixLen + t.hashLen
		id = id[:t.bufLen]
		n := binary.PutUvarint(id[preLen:], uint64(collisions))
		id = id[:preLen+n]
	}

	return id
}
