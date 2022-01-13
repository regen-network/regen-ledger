package server

import sdk "github.com/cosmos/cosmos-sdk/types"

type EcocreditAccount interface {
	GetTradableBalanceKey(denom batchDenomT) []byte
	GetRetiredBalanceKey(denom batchDenomT) []byte
	String() string
}

type EcocreditAcc sdk.AccAddress

func (ea EcocreditAcc) GetTradableBalanceKey(denom batchDenomT) []byte {
	return TradableBalanceKey(sdk.AccAddress(ea), denom)
}

func (ea EcocreditAcc) GetRetiredBalanceKey(denom batchDenomT) []byte {
	return RetiredBalanceKey(sdk.AccAddress(ea), denom)
}

func (ea EcocreditAcc) String() string {
	return sdk.AccAddress(ea).String()
}

func (bd basketDenomT) GetTradableBalanceKey(denom batchDenomT) []byte {
	return BasketBatchKey(bd, denom)
}

func (bd basketDenomT) GetRetiredBalanceKey(_ batchDenomT) []byte {
	return nil
}

func (bd basketDenomT) String() string {
	return string(bd)
}
