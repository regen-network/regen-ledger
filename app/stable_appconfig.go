package app

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/group"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ica "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts"
	icacontrollertypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"

	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (app *RegenApp) registerUpgradeHandlers() {
	upgradeName := "v5.0"
	app.UpgradeKeeper.SetUpgradeHandler(upgradeName,
		func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

			// set regen module consensus version
			fromVM[ecocredit.ModuleName] = 2
			fromVM[data.ModuleName] = 1

			// Add Interchain Accounts host module
			// set the ICS27 consensus version so InitGenesis is not run
			// we want to manually initialize the module with our own parameters
			fromVM[icatypes.ModuleName] = app.ModuleManager.Modules[icatypes.ModuleName].ConsensusVersion()

			// create ICS27 Controller submodule params, controller module not enabled.
			controllerParams := icacontrollertypes.Params{}

			// setup interchain accounts host
			// we can either set up predefined params here as part of the upgrade, or
			// we can let governance update via the legacy param proposals.
			// to add a message you add a type url like so:
			// sdk.MsgTypeURL(&core.MsgRetire{})
			hostParams := icahosttypes.Params{
				HostEnabled:   true,
				AllowMessages: []string{},
			}
			app.ICAHostKeeper.SetParams(ctx, hostParams)

			icaModule, ok := app.ModuleManager.Modules[icatypes.ModuleName].(ica.AppModule)
			if !ok {
				panic(fmt.Sprintf("expected interchain account module to be of type %T, got %T", ica.AppModule{}, app.ModuleManager.Modules[icatypes.ModuleName]))
			}
			// InitModule is called as an alternative to InitGenesis. It performs the same actions however InitModule
			// allows us to define and set the parameters directly.
			icaModule.InitModule(ctx, controllerParams, hostParams)

			// transfer module consensus version has been bumped to 2
			return app.ModuleManager.RunMigrations(ctx, app.configurator, fromVM)
		})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == upgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				group.ModuleName,
				icahosttypes.StoreKey,
			},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

func (app *RegenApp) setCustomAnteHandler(cfg client.TxConfig) (sdk.AnteHandler, error) {
	return ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			SignModeHandler: cfg.SignModeHandler(),
			FeegrantKeeper:  app.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)
}
