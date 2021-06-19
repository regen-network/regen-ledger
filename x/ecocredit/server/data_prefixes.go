package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

// batchDenomT is used to prevent errors when forming keys as accounts and denoms are
// both represented as strings
type batchDenomT string

// - 0x0 <accAddrLen (1 Byte)><accAddr_Bytes><denom_Bytes>
// - 0x1 <denom_Bytes>
// - 0x2 <accAddrLen (1 Byte)><accAddr_Bytes><denom_Bytes>
// - 0x3 <denom_Bytes>
// - 0x7 <denom_Bytes>

func TradableBalanceKey(acc sdk.AccAddress, denom batchDenomT) []byte {
	key := []byte{TradableBalancePrefix}
	key = append(key, address.MustLengthPrefix(acc)...)
	return append(key, denom...)
}

func ParseTradableBalanceKey(key []byte) (sdk.AccAddress, batchDenomT) {
	addrLen := key[1]
	addr := sdk.AccAddress(key[2 : 2+addrLen])
	return addr, batchDenomT(key[2+addrLen:])
}

func TradableSupplyKey(batchDenom batchDenomT) []byte {
	key := []byte{TradableSupplyPrefix}
	return append(key, batchDenom...)
}

func ParseTradableSupplyKey(key []byte) batchDenomT {
	return batchDenomT(key[1:])
}

func RetiredBalanceKey(acc sdk.AccAddress, batchDenom batchDenomT) []byte {
	key := []byte{RetiredBalancePrefix}
	key = append(key, address.MustLengthPrefix(acc)...)
	return append(key, batchDenom...)
}

func ParseRetiredBalanceKey(key []byte) (sdk.AccAddress, batchDenomT) {
	addrLen := key[1]
	addr := sdk.AccAddress(key[2 : 2+addrLen])
	return addr, batchDenomT(key[2+addrLen:])
}

func RetiredSupplyKey(batchDenom batchDenomT) []byte {
	key := []byte{RetiredSupplyPrefix}
	return append(key, batchDenom...)
}

func ParseRetiredSupplyKey(key []byte) batchDenomT {
	return batchDenomT(key[1:])
}

func MaxDecimalPlacesKey(batchDenom batchDenomT) []byte {
	key := []byte{MaxDecimalPlacesPrefix}
	return append(key, batchDenom...)
}

func ParseMaxDecimalPlacesKey(key []byte) batchDenomT {
	return batchDenomT(key[1:])
}
