package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

// BatchDenomT is used to prevent errors when forming keys as accounts and denoms are
// both represented as strings
type BatchDenomT string

// - 0x0 <accAddrLen (1 Byte)><accAddr_Bytes><denom_Bytes>: TradableBalance
// - 0x1 <denom_Bytes>: TradableSupply
// - 0x2 <accAddrLen (1 Byte)><accAddr_Bytes><denom_Bytes>: RetiredBalance
// - 0x3 <denom_Bytes>: RetiredSupply

// TradableBalanceKey creates the index key for recipient address and batch-denom
func TradableBalanceKey(acc sdk.AccAddress, denom BatchDenomT) []byte {
	key := []byte{TradableBalancePrefix}
	key = append(key, address.MustLengthPrefix(acc)...)
	return append(key, denom...)
}

// ParseBalanceKey parses the recipient address and batch-denom from tradable or retired balance key.
func ParseBalanceKey(key []byte) (sdk.AccAddress, BatchDenomT) {
	addrLen := key[1]
	addr := sdk.AccAddress(key[2 : 2+addrLen])
	return addr, BatchDenomT(key[2+addrLen:])
}

// TradableSupplyKey creates the tradable supply key for a given batch-denom
func TradableSupplyKey(batchDenom BatchDenomT) []byte {
	key := []byte{TradableSupplyPrefix}
	return append(key, batchDenom...)
}

// ParseSupplyKey parses the batch-denom from tradable or retired supply key
func ParseSupplyKey(key []byte) BatchDenomT {
	return BatchDenomT(key[1:])
}

// RetiredBalanceKey creates the index key for recipient address and batch-denom
func RetiredBalanceKey(acc sdk.AccAddress, batchDenom BatchDenomT) []byte {
	key := []byte{RetiredBalancePrefix}
	key = append(key, address.MustLengthPrefix(acc)...)
	return append(key, batchDenom...)
}

// RetiredSupplyKey creates the retired supply key for a given batch-denom
func RetiredSupplyKey(batchDenom BatchDenomT) []byte {
	key := []byte{RetiredSupplyPrefix}
	return append(key, batchDenom...)
}
