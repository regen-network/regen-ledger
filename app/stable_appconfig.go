// +build stable

package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{}
}

func setCustomKVStoreKeys() []string {
	return []string{}
}

func (app *RegenApp) setCustomKeepers() {
}

func (app *RegenApp) setCustomAppModules() []module.AppModule {
	return []module.AppModule{}
}

func (app *RegenApp) setCustomEndBlockModules() []string {
	return []string{}
}

func (app *RegenApp) setCustomInitGenesisOrder() []string {
	return []string{}
}

func (app *RegenApp) setCustomSimModules() []module.AppModuleSimulation {
	return []module.AppModuleSimulation{}
}
