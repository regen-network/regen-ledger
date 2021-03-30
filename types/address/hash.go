package address

import (
	"crypto/sha256"
)

// Len is the length of base addresses
const Len = sha256.Size

type Addressable interface {
	Address() []byte
}

// Hash creates a new address from address type and key
// Deprecated: use SDK function instead.
func Hash(typ string, key []byte) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(typ))
	th := hasher.Sum(nil)

	hasher.Reset()
	// the is error always nil, it's here only to satisfy the io.Writer interface
	if _, err := hasher.Write(th); err != nil {
		panic(err)
	}
	if _, err := hasher.Write(key); err != nil {
		panic(err)
	}

	return hasher.Sum(nil)
}

// Module is a specialized version of a composed address for modules. Each module account
// is constructed from a module name and module account key.
func Module(moduleName string, key []byte) []byte {
	mKey := append([]byte(moduleName), 0)
	return Hash("module", append(mKey, key...))
}
