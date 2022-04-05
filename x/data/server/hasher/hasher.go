package hasher

import (
	"encoding/binary"
	"fmt"
	"hash"

	"golang.org/x/crypto/blake2b"
)

// Hasher generates a unique binary identifier for a longer piece of binary data
// using an efficient, non-cryptographic hash function.
type Hasher interface {
	// CreateID is an idempotent method for creating a unique shortened identifier
	// for the provided binary value.
	CreateID(value []byte, collisions int) []byte
}

// NewHasher creates a new hasher instance. Default parameters are currently set to use the first
// 4-bytes of the 64-bit BLAKE2b, non-cryptographic hash. In the case of a collision, more bytes
// of the hash will be used for disambiguation but this happens in a minority of cases except
// for massively large data sets.
func NewHasher() (Hasher, error) {
	return NewHasherWithOptions(HashOptions{})
}

// NewHasherWithOptions creates a Hash with custom options. Most users should just use NewHasher
// with the default values.
func NewHasherWithOptions(options HashOptions) (Hasher, error) {
	minLength := options.MinLength
	if minLength == 0 {
		minLength = 4
	}

	newHash := options.NewHash
	if newHash == nil {
		newHash = func() hash.Hash {
			hash, err := blake2b.New(8, nil)
			if err != nil {
				panic(err) // an error should not occur creating a hash
			}
			return hash
		}
	}

	hashLen := len(newHash().Sum(nil))
	if minLength > hashLen {
		return nil, fmt.Errorf("option MinLength %d is greater than hash length %d", minLength, hashLen)
	}

	bufLen := hashLen + binary.MaxVarintLen64

	return hasher{
		minLen:  minLength,
		bufLen:  bufLen,
		newHash: newHash,
		hashLen: hashLen,
	}, nil
}

// HashOptions is used to specify custom hash options and should only be used by advanced users.
type HashOptions struct {
	// NewHash is a function which returns a new hash.Hash instance.
	NewHash func() hash.Hash

	// MinLength is the minimum number of hash bytes that will be used to create a lookup identifier.
	MinLength int
}

type hasher struct {
	minLen  int
	bufLen  int
	newHash func() hash.Hash
	hashLen int
}

func (t hasher) CreateID(value []byte, collisions int) (id []byte) {
	hasher := t.newHash()
	_, err := hasher.Write(value)
	if err != nil {
		// we panic here because hash.Write returning an error shouldn't happen
		panic(err)
	}
	hashBz := hasher.Sum(nil)

	id = make([]byte, t.minLen, t.bufLen)
	copy(id[:], hashBz[:t.minLen])

	// Deal with collisions by appending the equivalent number of bytes
	// from hashBz. If using this method will exceed hash length, append
	// a disambiguation varint. Such collisions are almost impossible with
	// good settings, but can happen with a suboptimal hash function.
	if t.minLen+collisions < t.hashLen {
		id = append(id, hashBz[collisions])
	} else {
		id = id[:t.bufLen]
		n := binary.PutUvarint(id[t.hashLen:], uint64(collisions))
		id = id[:t.hashLen+n]
	}

	return id
}
