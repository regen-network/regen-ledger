// +build !experimental

package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	moduletypes "github.com/regen-network/regen-ledger/types/module"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{}
}

func setCustomModules(app *RegenApp) []moduletypes.Module {
	return nil
}

func (app *RegenApp) registerUpgradeHandlers() {}
