package server

import sdk "github.com/cosmos/cosmos-sdk/types"

type EcocreditAccount interface {
	TradableBalanceKey(denom batchDenomT) []byte
	RetiredBalanceKey(denom batchDenomT) []byte
	String() string
}

type EcocreditEOA sdk.AccAddress

func (ea EcocreditEOA) TradableBalanceKey(denom batchDenomT) []byte {
	return TradableBalanceKey(sdk.AccAddress(ea), denom)
}

func (ea EcocreditEOA) RetiredBalanceKey(denom batchDenomT) []byte {
	return RetiredBalanceKey(sdk.AccAddress(ea), denom)
}

func (ea EcocreditEOA) String() string {
	return sdk.AccAddress(ea).String()
}

func (bd basketDenomT) TradableBalanceKey(denom batchDenomT) []byte {
	return BasketBatchKey(bd, denom)
}

func (bd basketDenomT) RetiredBalanceKey(_ batchDenomT) []byte {
	return nil
}

func (bd basketDenomT) String() string {
	return string(bd)
}
