// +build !experimental

package app

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{}
}

// setCustomModules registers new modules with the server module manager.
// It does nothing here and returns an empty manager since we're not using experimental mode.
func setCustomModules(app *RegenApp, interfaceRegistry types.InterfaceRegistry) *servermodule.Manager {
	return &servermodule.Manager{}
}

func (app *RegenApp) registerUpgradeHandlers() {}
