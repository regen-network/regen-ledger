package server

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/modules/incubator/orm"
	_ "github.com/cosmos/modules/incubator/orm"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

const (
	TradeableBalancePrefix = 0x0
	TradeableSupplyPrefix  = 0x1
	RetiredBalancePrefix   = 0x2
	RetiredSupplyPrefix    = 0x3
)

type serverImpl struct {
	denomPrefix string

	storeKey sdk.StoreKey

	bankKeeper bankkeeper.Keeper

	classInfoSeq   orm.Sequence
	classInfoTable orm.NaturalKeyTable

	batchInfoSeq   orm.Sequence
	batchInfoTable orm.NaturalKeyTable
}

type Server interface {
	ecocredit.MsgServer
	ecocredit.QueryServer
}

func NewServer(denomPrefix string, storeKey sdk.StoreKey, bankKeeper bankkeeper.Keeper) Server {
	s := serverImpl{denomPrefix: denomPrefix, storeKey: storeKey, bankKeeper: bankKeeper}

	return s
}

func TradeableBalanceKey(acc string, batchDenom string) []byte {
	key := []byte{TradeableBalancePrefix}
	str := fmt.Sprintf("%s|%s", acc, batchDenom)
	return append(key, str...)
}

func TradeableSupplyKey(batchDenom string) []byte {
	key := []byte{TradeableSupplyPrefix}
	return append(key, batchDenom...)
}

func RetiredBalanceKey(acc string, batchDenom string) []byte {
	key := []byte{RetiredBalancePrefix}
	str := fmt.Sprintf("%s|%s", acc, batchDenom)
	return append(key, str...)
}

func RetiredSupplyKey(batchDenom string) []byte {
	key := []byte{RetiredSupplyPrefix}
	return append(key, batchDenom...)
}
