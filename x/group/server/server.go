package server

import (
	"github.com/regen-network/regen-ledger/x/group/types"

	"github.com/cosmos/cosmos-sdk/types/module"
)

type serverImpl struct {
	*Keeper
}

func newServer(keeper *Keeper) serverImpl {
	return serverImpl{Keeper: keeper}
}

func RegisterServices(keeper *Keeper, configurator module.Configurator) {
	types.RegisterMsgServer(configurator.MsgServer(), newServer(keeper))
}
