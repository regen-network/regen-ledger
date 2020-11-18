package server

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/data"
)

const DefaultModuleName = "data"

type Module struct{}

var _ servermodule.Module = Module{}

type serverImpl struct {
	storeKey sdk.StoreKey
}

func newServer(storeKey sdk.StoreKey) serverImpl {
	return serverImpl{storeKey: storeKey}
}

func (a Module) RegisterTypes(registry codectypes.InterfaceRegistry) {
	data.RegisterTypes(registry)
}

func (a Module) RegisterServices(configurator servermodule.Configurator) {
	impl := newServer(configurator.ModuleKey())
	data.RegisterMsgServer(configurator.MsgServer(), impl)
	data.RegisterQueryServer(configurator.QueryServer(), impl)
}
