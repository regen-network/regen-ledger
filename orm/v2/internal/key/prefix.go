package key

import "encoding/binary"

func MakePrefix(namespacePrefix []byte, id uint32) []byte {
	nsPrefixLen := len(namespacePrefix)
	prefix := make([]byte, nsPrefixLen+binary.MaxVarintLen32)
	copy(prefix, namespacePrefix)
	n := binary.PutUvarint(prefix[nsPrefixLen:], uint64(id))
	return prefix[:nsPrefixLen+n]
}
