package server

import "fmt"

// batchDenomT is used to prevent errors when forming keys as accounts and denoms are
// both represented as strings
type batchDenomT string

func TradableBalanceKey(acc string, denom batchDenomT) []byte {
	key := []byte{TradableBalancePrefix}
	str := fmt.Sprintf("%s|%s", acc, denom)
	return append(key, str...)
}

func TradableSupplyKey(batchDenom batchDenomT) []byte {
	key := []byte{TradableSupplyPrefix}
	return append(key, batchDenom...)
}

func RetiredBalanceKey(acc string, batchDenom batchDenomT) []byte {
	key := []byte{RetiredBalancePrefix}
	str := fmt.Sprintf("%s|%s", acc, batchDenom)
	return append(key, str...)
}

func RetiredSupplyKey(batchDenom batchDenomT) []byte {
	key := []byte{RetiredSupplyPrefix}
	return append(key, batchDenom...)
}
