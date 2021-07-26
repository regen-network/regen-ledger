package server

import (
	"encoding/base64"
)

const (
	AnchorTablePrefix byte = 0x0
	CIDSignerPrefix   byte = 0x1
	SignerCIDPrefix   byte = 0x2
	DataTablePrefix   byte = 0x3
)

func AnchorKey(cid []byte) []byte {
	return append([]byte{AnchorTablePrefix}, cid...)
}

func CIDBase64String(cid []byte) string {
	return base64.StdEncoding.EncodeToString(cid)
}

func CIDSignerKey(cidStr string, signer string) []byte {
	key := CIDSignerIndexPrefix(cidStr)
	key = append(key, signer...)
	return key
}

func CIDSignerIndexPrefix(cidStr string) []byte {
	key := []byte{CIDSignerPrefix}
	key = append(key, cidStr...)
	key = append(key, 0)
	return key
}

func SignerCIDKey(signer string, cid []byte) []byte {
	key := SignerCIDIndexPrefix(signer)
	key = append(key, cid...)
	return key
}

func SignerCIDIndexPrefix(signer string) []byte {
	key := []byte{SignerCIDPrefix}
	key = append(key, signer...)
	key = append(key, 0)
	return key
}

func DataKey(cid []byte) []byte {
	return append([]byte{DataTablePrefix}, cid...)
}
