// DONTCOVER

package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/group"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

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
			app.UpgradeKeeper.SetModuleVersionMap(ctx, fromVM)

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
