package orm

// IndexKeyCodec defines methods for building index keys from a searchable key
// and the RowID from the original table, and splitting the index key to
// retrieve the RowID.
// type IndexKeyCodec interface {
// 	// PrefixSearchableKey adds an optional prefix to the searchable key and
// 	// should be called before all prefix lookups and stores.
// 	PrefixSearchableKey(searchableKey []byte) []byte
// 	// BuildIndexKey encodes a searchable key and the target RowID.
// 	BuildIndexKey(searchableKey []byte, rowID RowID) ([]byte, error)
// 	// StripRowID returns the RowID from the indexKey. It is the reverse
// 	// operation to BuildIndexKey.
// 	StripRowID(indexKey []byte) RowID
// }

// var _, _ IndexKeyCodec = Max255DynamicLengthIndexKeyCodec{},
// 	FixLengthIndexKeyCodec{}

// // Max255DynamicLengthIndexKeyCodec works with up to 255 byte dynamic size
// // searchable keys. They are encoded as `concat(len(searchableKey),
// // searchableKey, rowID)` and can be used with PrimaryKey or external Key tables
// // for example.
// type Max255DynamicLengthIndexKeyCodec struct{}

// // PrefixSearchableKey adds a length prefix to the searchable key
// func (Max255DynamicLengthIndexKeyCodec) PrefixSearchableKey(searchableKey []byte) []byte {
// 	return AddLengthPrefix(searchableKey)
// }

// // BuildIndexKey builds the index key by adding a length prefix to searchableKey
// // and appending it with rowID. The searchableKey length must not be greater
// // than 255.
// func (Max255DynamicLengthIndexKeyCodec) BuildIndexKey(searchableKey []byte, rowID RowID) ([]byte, error) {
// 	rowIDLen := len(rowID)
// 	if rowIDLen == 0 {
// 		return nil, ErrArgument.Wrap("Empty RowID")
// 	}

// 	searchableKeyLen := len(searchableKey)
// 	res := make([]byte, 1+searchableKeyLen+rowIDLen)
// 	copy(res, AddLengthPrefix(searchableKey))
// 	copy(res[1+searchableKeyLen:], rowID)
// 	return res, nil
// }

// // StripRowID returns the RowID from the indexKey. It is the reverse operation
// // to BuildIndexKey, dropping the searchableKey and its length prefix.
// func (Max255DynamicLengthIndexKeyCodec) StripRowID(indexKey []byte) RowID {
// 	searchableKeyLen := indexKey[0]
// 	return indexKey[1+searchableKeyLen:]
// }

// // FixLengthIndexKeyCodec expects the searchableKey to always have the same
// // length with all entries. They are encoded as `concat(searchableKey, rowID)`
// // and can be used with AutoUint64Tables and length EncodedSeqLength for
// // example.
// type FixLengthIndexKeyCodec struct {
// 	searchableKeyLength int
// 	rowIDLength         int
// }

// // FixLengthIndexKeys is a constructor for FixLengthIndexKeyCodec.
// func FixLengthIndexKeys() *FixLengthIndexKeyCodec {
// 	return &FixLengthIndexKeyCodec{}
// }

// // PrefixSearchableKey adds no prefix
// func (FixLengthIndexKeyCodec) PrefixSearchableKey(searchableKey []byte) []byte {
// 	return searchableKey
// }

// // BuildIndexKey builds the index key by appending searchableKey with rowID.
// // The searchableKey length must not be greater than what is defined by
// // searchableKeyLength in construction.
// func (c FixLengthIndexKeyCodec) BuildIndexKey(searchableKey []byte, rowID RowID) ([]byte, error) {
// 	switch n := len(rowID); {
// 	case n == 0:
// 		return nil, ErrArgument.Wrap("Empty RowID")
// 	case n > c.rowIDLength:
// 		return nil, ErrArgument.Wrap("RowID exceeds max size")
// 	}
// 	if n := len(searchableKey); n != c.searchableKeyLength {
// 		return nil, ErrArgument.Wrapf(
// 			"searchableKey is incorrect length, expected %d, got %d",
// 			c.searchableKeyLength,
// 			n,
// 		)
// 	}
// 	res := make([]byte, c.searchableKeyLength+c.rowIDLength)
// 	copy(res, searchableKey)
// 	copy(res[c.searchableKeyLength:], rowID)
// 	return res, nil
// }

// // StripRowID returns the RowID from the indexKey. It is the reverse operation
// // to BuildIndexKey, dropping the searchableKey.
// func (c FixLengthIndexKeyCodec) StripRowID(indexKey []byte) RowID {
// 	return indexKey[c.searchableKeyLength:]
// }
