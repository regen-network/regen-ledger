package module

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/regen-network/regen-ledger/types/module"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/bank"
	"github.com/regen-network/regen-ledger/x/bank/server"
)

type Module struct {
	DenomNamespaceAdmins map[string]module.ModuleID
}

func (m Module) Name() string {
	panic("implement me")
}

func (m Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	bank.RegisterTypes(registry)
}

func (m Module) RegisterServices(configurator servermodule.Configurator) {
	server.RegisterServices(m.DenomNamespaceAdmins, configurator)
}

var _ servermodule.Module = Module{}
