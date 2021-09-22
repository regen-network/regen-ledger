package orm

import (
	"fmt"
	"reflect"
)

// buildKeyFromParts encodes and concatenates primary key and index parts.
// They can be []byte, string, and integer types. The function will return
// an error if there is a part of any other type.
// Key parts except the last part follow these rules:
//  - []byte is encoded with a single byte length prefix
//  - strings are null-terminated
//  - integers are encoded using 4 or 8 byte big endian.
func buildKeyFromParts(parts []interface{}) ([]byte, error) {
	bytesSlice := make([][]byte, len(parts))
	totalLen := 0
	var err error
	for i, part := range parts {
		bytesSlice[i], err = keyPartBytes(part, len(parts) > 1 && i == len(parts)-1)
		if err != nil {
			return nil, err
		}
		// bytesSlice[i] = keyPartBytes(part)
		totalLen += len(bytesSlice[i])
	}
	key := make([]byte, 0, totalLen)
	for _, bs := range bytesSlice {
		key = append(key, bs...)
	}
	return key, nil
}

func keyPartBytes(part interface{}, last bool) ([]byte, error) {
	switch v := part.(type) {
	case []byte:
		if last || len(v) == 0 {
			return v, nil
		}
		// if len(v) == 0 {
		// 	return nil, nil
		// 	// return nil, errors.Wrap(ErrArgument, "empty index key")
		// }
		return AddLengthPrefix(v), nil
	case string:
		if last || len(v) == 0 {
			return []byte(v), nil
		}
		// if len(v) == 0 {
		// 	return nil, nil
		// 	// return nil, errors.Wrap(ErrArgument, "empty index key")
		// }
		return NullTerminatedBytes(v), nil
	case uint64:
		return EncodeSequence(v), nil
	default:
		return nil, fmt.Errorf("type %T not allowed as key part", v)
	}
}

// AddLengthPrefix prefixes the byte array with its length as 8 bytes. The function will panic
// if the bytes length is bigger than 255.
func AddLengthPrefix(bytes []byte) []byte {
	byteLen := len(bytes)
	if byteLen > 255 {
		panic("Cannot create primary key with an []byte of length greater than 255 bytes. Try again with a smaller []byte.")
	}

	prefixedBytes := make([]byte, 1+len(bytes))
	copy(prefixedBytes, []byte{uint8(byteLen)})
	copy(prefixedBytes[1:], bytes)
	return prefixedBytes
}

// NullTerminatedBytes converts string to byte array and null terminate it
func NullTerminatedBytes(s string) []byte {
	bytes := make([]byte, len(s)+1)
	copy(bytes, s)
	return bytes
}

func stripRowID(indexKey []byte, indexKeyType reflect.Type) (RowID, error) {
	switch indexKeyType {
	case reflect.TypeOf(([]byte)(nil)):
		searchableKeyLen := indexKey[0]
		return indexKey[1+searchableKeyLen:], nil
	case reflect.TypeOf((string)("")):
		searchableKeyLen := 0
		for i, b := range indexKey {
			if b == 0 {
				searchableKeyLen = i
				break
			}
		}
		return indexKey[1+searchableKeyLen:], nil
	case reflect.TypeOf((uint64)(0)):
		return indexKey[EncodedSeqLength:], nil
	default:
		return nil, fmt.Errorf("type %T not allowed as index key", reflect.New(indexKeyType).Interface())
	}
}
