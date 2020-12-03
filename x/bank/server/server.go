package server

import (
	"github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/bank"
)

type serverImpl struct {
	denomNamespaceAdmins map[string]module.ModuleID
	key                  server.RootModuleKey
}

func RegisterServices(denomNamespaceAdmins map[string]module.ModuleID, configurator server.Configurator) {
	impl := serverImpl{
		denomNamespaceAdmins: denomNamespaceAdmins,
		key:                  configurator.ModuleKey(),
	}
	bank.RegisterMsgServer(configurator.MsgServer(), impl)
	bank.RegisterQueryServer(configurator.QueryServer(), impl)
}
