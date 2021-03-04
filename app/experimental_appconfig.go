// +build experimental

package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	moduletypes "github.com/regen-network/regen-ledger/types/module"
	data "github.com/regen-network/regen-ledger/x/data/module"
	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit/module"
	group "github.com/regen-network/regen-ledger/x/group/module"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{
		data.Module{},
		ecocredit.Module{},
		group.Module{},
	}
}

func setCustomModules(app *RegenApp) []moduletypes.Module {
	// BEGIN HACK: this is a total, ugly hack until x/auth supports ADR 033 or we have a suitable alternative
	groupModule := group.Module{AccountKeeper: app.AccountKeeper}
	// use a separate newModules from the global NewModules here because we need to pass state into the group module
	newModules := []moduletypes.Module{
		ecocredit.Module{},
		data.Module{},
		groupModule,
	}
	// END HACK
	return newModules
}

func (app *RegenApp) registerUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler("Mau", func(ctx sdk.Context, plan upgradetypes.Plan) {
		// no-op handler, does nothing
	})
}
