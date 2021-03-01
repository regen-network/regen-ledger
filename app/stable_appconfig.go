// +build experimental
// TODO: build flags not working for simulations

package app

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{}
}

func setCustomModules(app *RegenApp, interfaceRegistry types.InterfaceRegistry) *servermodule.Manager {
	return &servermodule.Manager{}
}

func (app *RegenApp) registerUpgradeHandlers() {}
