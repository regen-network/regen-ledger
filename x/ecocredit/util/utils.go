package util

import (
	"encoding/binary"
	"fmt"
	"strings"
	"unicode"

	"github.com/btcsuite/btcutil/base58"
	"github.com/enigmampc/btcutil/bech32"
)

// ConvertAndEncode converts from a base64 encoded byte string to base32 encoded byte string and then to bech32.
func ConvertAndEncode(hrp string, data []byte) (string, error) {
	converted, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("encoding bech32 failed: %w", err)
	}

	return bech32.Encode(hrp, converted)
}

// DecodeAndConvert decodes a bech32 encoded string and converts to base64 encoded bytes.
func DecodeAndConvert(bech string) (string, []byte, error) {
	hrp, data, err := bech32.Decode(bech, 1023)
	if err != nil {
		return "", nil, fmt.Errorf("decoding bech32 failed: %w", err)
	}

	converted, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, fmt.Errorf("decoding bech32 failed: %w", err)
	}

	return hrp, converted, nil
}

func Uint64ToBase58Check(x uint64) string {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, x)
	return base58.CheckEncode(buf[:n], 0)
}

func FastRemoveWhitespace(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
