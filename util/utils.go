package util

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

func MustDecodeBech32(bech string) (hrp string, data []byte) {
	hrp, data, err := bech32.DecodeAndConvert(bech)
	if err != nil {
		panic(err)
	}
	return hrp, data
}
