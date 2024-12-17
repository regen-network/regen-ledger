package app

import (
	"github.com/cometbft/cometbft/libs/log"

	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	"github.com/cosmos/cosmos-sdk/baseapp"

	// wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func (app *RegenApp) registerUpgrades() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	app.registerUpgrade6_0(upgradeInfo)
}

func (app *RegenApp) registerUpgrade6_0(upgradeInfo upgradetypes.Plan) {
	planName := "v6.0"

	// Set param key table for params module migration
	for _, subspace := range app.ParamsKeeper.GetSubspaces() {
		subspace := subspace
		found := true
		var keyTable paramstypes.KeyTable
		switch subspace.Name() {
		case authtypes.ModuleName:
			keyTable = authtypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case banktypes.ModuleName:
			keyTable = banktypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case stakingtypes.ModuleName:
			keyTable = stakingtypes.ParamKeyTable()
		case minttypes.ModuleName:
			keyTable = minttypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case distrtypes.ModuleName:
			keyTable = distrtypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case slashingtypes.ModuleName:
			keyTable = slashingtypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case govtypes.ModuleName:
			keyTable = govv1.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case crisistypes.ModuleName:
			keyTable = crisistypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		// case wasmtypes.ModuleName:
		// 	keyTable = wasmtypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		default:
			// subspace not handled
			found = false
		}
		if found && !subspace.HasKeyTable() {
			subspace.WithKeyTable(keyTable)
		}
	}
	baseAppLegacySS := app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())

	app.UpgradeKeeper.SetUpgradeHandler(planName,
		func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			printPlanName(planName, ctx.Logger())

			// Migrate CometBFT consensus parameters from x/params module to a dedicated x/consensus module.
			baseapp.MigrateParams(ctx, baseAppLegacySS, &app.ConsensusParamsKeeper)

			// explicitly update the IBC 02-client params, adding the localhost client type
			params := app.IBCKeeper.ClientKeeper.GetParams(ctx)
			params.AllowedClients = append(params.AllowedClients, ibcexported.Localhost)
			app.IBCKeeper.ClientKeeper.SetParams(ctx, params)

			fromVM, err := app.ModuleManager.RunMigrations(ctx, app.configurator, fromVM)
			if err != nil {
				return fromVM, err
			}
			// Cosmos SDK v0.47 introduced new gov param: MinInitialDepositRatio
			govParams := app.GovKeeper.GetParams(ctx)
			govParams.MinInitialDepositRatio = sdk.NewDecWithPrec(1, 1).String()
			err = app.GovKeeper.SetParams(ctx, govParams)
			return fromVM, err
		},
	)

	app.storeUpgrade(planName, upgradeInfo, storetypes.StoreUpgrades{
		Added: []string{
			consensustypes.ModuleName,
			crisistypes.ModuleName,
		},
	})
}

// helper function to check if the store loader should be upgraded
// configure store loader that checks if version == upgradeHeight and applies store upgrades
func (app *RegenApp) storeUpgrade(planName string, ui upgradetypes.Plan, stores storetypes.StoreUpgrades) {
	if ui.Name == planName && !app.UpgradeKeeper.IsSkipHeight(ui.Height) {
		app.SetStoreLoader(
			upgradetypes.UpgradeStoreLoader(ui.Height, &stores))
	}
}

func printPlanName(planName string, logger log.Logger) {
	logger.Info("-----------------------------\n-----------------------------")
	logger.Info("Upgrade handler execution", "name", planName)
}
