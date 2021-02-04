package types

import "crypto/sha256"

func AddressHash(prefix string, contents []byte) []byte {
	prefixHash := sha256.Sum256([]byte(prefix))
	preImage := make([]byte, 32+len(contents))
	preImage = append(preImage, prefixHash[:]...)
	preImage = append(preImage, contents...)
	sum := sha256.Sum256(preImage)
	return sum[:32]
}
