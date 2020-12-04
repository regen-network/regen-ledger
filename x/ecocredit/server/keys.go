package server

import "fmt"

const (
	RetiredBalancePrefix byte = 0x2
	RetiredSupplyPrefix  byte = 0x3
	IDSeqPrefix          byte = 0x4
	ClassInfoTablePrefix byte = 0x5
	BatchInfoTablePrefix byte = 0x6
)

// batchDenomT is used to prevent errors when forming keys as accounts and denoms are
// both represented as strings
type batchDenomT string

func RetiredBalanceKey(acc string, batchDenom batchDenomT) []byte {
	key := []byte{RetiredBalancePrefix}
	str := fmt.Sprintf("%s|%s", acc, batchDenom)
	return append(key, str...)
}

func RetiredSupplyKey(batchDenom batchDenomT) []byte {
	key := []byte{RetiredSupplyPrefix}
	return append(key, batchDenom...)
}
