// +build !experimental

package app

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{}
}

func setCustomModules(app *RegenApp, interfaceRegistry types.InterfaceRegistry) {}

func (app *RegenApp) registerUpgradeHandlers() {}
