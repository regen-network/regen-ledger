package server

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/modules/incubator/orm"
	_ "github.com/cosmos/modules/incubator/orm"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

const (
	RetiredBalancePrefix = 0x4
	RetiredSupplyPrefix  = 0x5
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

func RetiredBalanceKey(acc string, batchDenom string) []byte {
	key := []byte{RetiredBalancePrefix}
	str := fmt.Sprintf("%s/%s", acc, batchDenom)
	return append(key, str...)
}

func RetiredSupplyKey(batchDenom string) []byte {
	key := []byte{RetiredBalancePrefix}
	return append(key, batchDenom...)
}
