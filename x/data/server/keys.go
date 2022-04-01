package server

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	IriIDTablePrefix byte = iota
	AnchorTimestampPrefix
	IDAttestorPrefix
	AttestorIDPrefix
	ORMStatePrefix
)

func AnchorTimestampKey(id []byte) []byte {
	return append([]byte{AnchorTimestampPrefix}, id...)
}

func IDAttestorTimestampKey(id []byte, address sdk.AccAddress) []byte {
	if len(id) > 255 {
		panic(fmt.Errorf("id length must be <= 255, found: %d", len(id)))
	}

	key := make([]byte, 0, len(id)+len(address)+2)
	key = append(key, IDAttestorPrefix, byte(len(id)))
	key = append(key, id...)
	key = append(key, address...)
	return key
}

func IDAttestorIndexPrefix(id []byte) []byte {
	if len(id) > 255 {
		panic(fmt.Errorf("id length must be <= 255, found: %d", len(id)))
	}

	key := make([]byte, 0, len(id)+2)
	key = append(key, IDAttestorPrefix, byte(len(id)))
	key = append(key, id...)
	return key
}

func AttestorIDKey(address sdk.AccAddress, id []byte) []byte {
	if len(address) > 255 {
		panic(fmt.Errorf("address length must be <= 255, found: %d", len(address)))
	}

	key := make([]byte, 0, len(id)+len(address)+2)
	key = append(key, AttestorIDPrefix, byte(len(address)))
	key = append(key, address...)
	key = append(key, id...)
	return key
}

func AttestorIDIndexPrefix(address sdk.AccAddress) []byte {
	if len(address) > 255 {
		panic(fmt.Errorf("address length must be <= 255, found: %d", len(address)))
	}

	key := make([]byte, 0, len(address)+2)
	key = append(key, AttestorIDPrefix, byte(len(address)))
	key = append(key, address...)
	return key
}
