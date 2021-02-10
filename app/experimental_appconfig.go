// +build !stable

package app

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/group"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{
		data.AppModuleBasic{},
		ecocredit.AppModuleBasic{},
		group.AppModuleBasic{},
	}
}

func setCustomKVStoreKeys() []string {
	return []string{
		data.StoreKey,
		ecocredit.StoreKey,
		group.StoreKey,
	}
}

func (app *RegenApp) setCustomKeepers() {
	// TODO Register regen module keepers here
}

func (app *RegenApp) setCustomAppModules() []module.AppModule {
	return []module.AppModule{
		// TODO register regen app modules here
	}
}

func (app *RegenApp) setCustomEndBlockModules() []string {
	return []string{
		// TODO add endblock modules
	}
}

func (app *RegenApp) setCustomInitGenesisOrder() []string {
	return []string{
		// TODO
	}
}

func (app *RegenApp) setCustomSimModules() []module.AppModuleSimulation {
	return []module.AppModuleSimulation{
		data.NewAppModuleSimulation(
			// ...
		),
	}
}
