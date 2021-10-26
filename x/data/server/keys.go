package server

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	IriIDTablePrefix byte = iota
	AnchorTimestampPrefix
	IDSignerPrefix
	SignerIDPrefix
)

func AnchorTimestampKey(id []byte) []byte {
	return append([]byte{AnchorTimestampPrefix}, id...)
}

func IDSignerTimestampKey(id []byte, address sdk.AccAddress) []byte {
	key := make([]byte, 0, len(id)+len(address)+2)
	key = append(key, IDSignerPrefix)
	key = append(key, byte(len(id)))
	key = append(key, id...)
	key = append(key, address...)
	return key
}

func IDSignerIndexPrefix(id []byte) []byte {
	key := make([]byte, 0, len(id)+2)
	key = append(key, IDSignerPrefix)
	key = append(key, byte(len(id)))
	key = append(key, id...)
	return key
}

func SignerIDKey(address sdk.AccAddress, id []byte) []byte {
	key := make([]byte, 0, len(id)+len(address)+2)
	key = append(key, SignerIDPrefix)
	key = append(key, byte(len(address)))
	key = append(key, address...)
	key = append(key, id...)
	return key
}

func SignerIDIndexPrefix(address sdk.AccAddress) []byte {
	key := make([]byte, 0, len(address)+2)
	key = append(key, SignerIDPrefix)
	key = append(key, byte(len(address)))
	key = append(key, address...)
	return key
}
