package server

import (
	"encoding/base64"
	"fmt"
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
	return append([]byte{CIDSignerPrefix}, []byte(fmt.Sprintf("%s/%s", cidStr, signer))...)
}

func CIDSignerIndexPrefix(cidStr string) []byte {
	return append([]byte{CIDSignerPrefix}, []byte(fmt.Sprintf("%s/", cidStr))...)
}

func SignerCIDKey(signer string, cidStr string) []byte {
	return append([]byte{SignerCIDPrefix}, []byte(fmt.Sprintf("%s/%s", signer, cidStr))...)
}

func SignerCIDIndexPrefix(signer string) []byte {
	return append([]byte{SignerCIDPrefix}, []byte(fmt.Sprintf("%s/", signer))...)
}

func DataKey(cid []byte) []byte {
	return append([]byte{DataTablePrefix}, cid...)
}
