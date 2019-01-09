package utils

import (
	"github.com/tendermint/tendermint/libs/bech32"
)

func MustEncodeBech32(hrp string, data []byte) string {
	str, err := bech32.ConvertAndEncode(hrp, data)
	if err != nil {
		panic(err)
	}
	return str
}
