package server

import (
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/bank"
)

type serverImpl struct {
	denomNamespaceAdmins map[string]types.ModuleID
	key                  server.RootModuleKey
}

func RegisterServices(denomNamespaceAdmins map[string]types.ModuleID, configurator server.Configurator) {
	impl := serverImpl{
		denomNamespaceAdmins: denomNamespaceAdmins,
		key:                  configurator.ModuleKey(),
	}
	bank.RegisterMsgServer(configurator.MsgServer(), impl)
	bank.RegisterQueryServer(configurator.QueryServer(), impl)
}
