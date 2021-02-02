package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/data"
)

type serverImpl struct {
	storeKey sdk.StoreKey
}

func newServer(storeKey sdk.StoreKey) serverImpl {
	return serverImpl{storeKey: storeKey}
}

func RegisterServices(configurator servermodule.Configurator) {
	impl := newServer(configurator.ModuleKey())
	data.RegisterMsgServer(configurator.MsgServer(), impl)
	data.RegisterQueryServer(configurator.QueryServer(), impl)
}
