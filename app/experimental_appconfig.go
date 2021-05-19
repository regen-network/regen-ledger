// +build experimental

package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	moduletypes "github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/types/module/server"
	data "github.com/regen-network/regen-ledger/x/data/module"
	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit/module"
	group "github.com/regen-network/regen-ledger/x/group/module"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distrclient.ProposalHandler,
			upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
		),
		data.Module{},
		ecocredit.Module{},
		group.Module{},
	}
}

func setCustomKVStoreKeys() []string {
	return []string{}
}

func (app *RegenApp) setCustomKeeprs(bApp *baseapp.BaseApp, keys map[string]*sdk.KVStoreKey, appCodec codec.Codec, govRouter govtypes.Router, homePath string) {
}

// setCustomModules registers new modules with the server module manager.
func setCustomModules(app *RegenApp, interfaceRegistry types.InterfaceRegistry) *server.Manager {

	/* New Module Wiring START */
	newModuleManager := server.NewManager(app.BaseApp, codec.NewProtoCodec(interfaceRegistry))

	// BEGIN HACK: this is a total, ugly hack until x/auth & x/bank supports ADR 033 or we have a suitable alternative
	groupModule := group.Module{AccountKeeper: app.AccountKeeper, BankKeeper: app.BankKeeper}
	// use a separate newModules from the global NewModules here because we need to pass state into the group module
	newModules := []moduletypes.Module{
		ecocredit.Module{},
		data.Module{},
		groupModule,
	}
	err := newModuleManager.RegisterModules(newModules)
	if err != nil {
		panic(err)
	}
	// END HACK

	err = newModuleManager.CompleteInitialization()
	if err != nil {
		panic(err)
	}

	/* New Module Wiring END */
	return newModuleManager
}

func (app *RegenApp) registerUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler("v0.43.0-beta1-upgrade", func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

		// skipping x/auth migrations. It is already patched in regen-ledger v1.0
		fromVM["auth"] = auth.AppModule{}.ConsensusVersion()

		return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	})
}

func (app *RegenApp) setCustomModuleManager() []module.AppModule {
	return []module.AppModule{}
}

func setCustomOrderInitGenesis() []string {
	return []string{}
}

func (app *RegenApp) setCustomSimulationManager() []module.AppModuleSimulation {
	return []module.AppModuleSimulation{
		group.Module{
			Registry:      app.interfaceRegistry,
			BankKeeper:    app.BankKeeper,
			AccountKeeper: app.AccountKeeper,
		},
	}
}

func initCustomParamsKeeper(paramsKeeper *paramskeeper.Keeper) {
}
