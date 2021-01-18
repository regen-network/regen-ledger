package server

import (
	"encoding/base64"
	"encoding/binary"
)

const (
	IriIdTablePrefix   byte = 0x0
	IdIriTablePrefix   byte = 0x1
	AnchorTablePrefix  byte = 0x2
	IDSignerPrefix     byte = 0x3
	SignerIDPrefix     byte = 0x4
	RawDataTablePrefix byte = 0x5
)

func IriIdKey(iri string) []byte {
	key := make([]byte, len(iri)+1)
	key[0] = IriIdTablePrefix
	copy(key[1:], iri)
	return key
}

func IdIriKey(id uint64) []byte {
	key := make([]byte, 9)
	key[0] = IriIdTablePrefix
	binary.LittleEndian.PutUint64(key[1:], id)
	return key
}

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
