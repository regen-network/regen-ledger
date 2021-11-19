package store

type KVStore interface {
	// Get returns nil iff key doesn't exist. Panics on nil key.
	Get(key []byte) []byte

	// Has checks if a key exists. Panics on nil key.
	Has(key []byte) bool

	// Set sets the key. Panics on nil key or value.
	Set(key, value []byte)

	// Delete deletes the key. Panics on nil key.
	Delete(key []byte)

	// KVStoreIterator over a domain of keys in ascending order. End is exclusive.
	// Start must be less than end, or the KVStoreIterator is invalid.
	// KVStoreIterator must be closed by caller.
	// To iterate over entire domain, use store.KVStoreIterator(nil, nil)
	// CONTRACT: No writes may happen within a domain while an iterator exists over it.
	// Exceptionally allowed for cachekv.Store, safe to write in the modules.
	Iterator(start, end []byte) KVStoreIterator

	// KVStoreIterator over a domain of keys in descending order. End is exclusive.
	// Start must be less than end, or the KVStoreIterator is invalid.
	// KVStoreIterator must be closed by caller.
	// CONTRACT: No writes may happen within a domain while an iterator exists over it.
	// Exceptionally allowed for cachekv.Store, safe to write in the modules.
	ReverseIterator(start, end []byte) KVStoreIterator
}

type KVStoreIterator interface {
	// Domain returns the start (inclusive) and end (exclusive) limits of the iterator.
	// CONTRACT: start, end readonly []byte
	Domain() (start []byte, end []byte)

	// Valid returns whether the current iterator is valid. Once invalid, the KVStoreIterator remains
	// invalid forever.
	Valid() bool

	// Next moves the iterator to the next key in the database, as defined by order of iteration.
	// If Valid returns false, this method will panic.
	Next()

	// Key returns the key at the current position. Panics if the iterator is invalid.
	// CONTRACT: key readonly []byte
	Key() (key []byte)

	// Value returns the value at the current position. Panics if the iterator is invalid.
	// CONTRACT: value readonly []byte
	Value() (value []byte)

	// Error returns the last error encountered by the iterator, if any.
	Error() error

	// Close closes the iterator, relasing any allocated resources.
	Close() error
}
