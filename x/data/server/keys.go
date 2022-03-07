package server

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	IriIDTablePrefix byte = iota
	AnchorTimestampPrefix
	IDSignerPrefix
	SignerIDPrefix
	ORMStatePrefix
)

func AnchorTimestampKey(id []byte) []byte {
	return append([]byte{AnchorTimestampPrefix}, id...)
}

func IDSignerTimestampKey(id []byte, address sdk.AccAddress) []byte {
	if len(id) > 255 {
		panic(fmt.Errorf("id length must be <= 255, found: %d", len(id)))
	}

	key := make([]byte, 0, len(id)+len(address)+2)
	key = append(key, IDSignerPrefix, byte(len(id)))
	key = append(key, id...)
	key = append(key, address...)
	return key
}

func IDSignerIndexPrefix(id []byte) []byte {
	if len(id) > 255 {
		panic(fmt.Errorf("id length must be <= 255, found: %d", len(id)))
	}

	key := make([]byte, 0, len(id)+2)
	key = append(key, IDSignerPrefix, byte(len(id)))
	key = append(key, id...)
	return key
}

func SignerIDKey(address sdk.AccAddress, id []byte) []byte {
	if len(address) > 255 {
		panic(fmt.Errorf("address length must be <= 255, found: %d", len(address)))
	}

	key := make([]byte, 0, len(id)+len(address)+2)
	key = append(key, SignerIDPrefix, byte(len(address)))
	key = append(key, address...)
	key = append(key, id...)
	return key
}

func SignerIDIndexPrefix(address sdk.AccAddress) []byte {
	if len(address) > 255 {
		panic(fmt.Errorf("address length must be <= 255, found: %d", len(address)))
	}

	key := make([]byte, 0, len(address)+2)
	key = append(key, SignerIDPrefix, byte(len(address)))
	key = append(key, address...)
	return key
}
