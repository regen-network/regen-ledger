//go:build !experimental
// +build !experimental

// DONTCOVER

package app

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	// ecocreditmodule "github.com/regen-network/regen-ledger/x/ecocredit/module"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distrclient.ProposalHandler,
			upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
		),
	}
}

// setCustomModules registers new modules with the server module manager.
// It does nothing here and returns an empty manager since we're not using experimental mode.
func setCustomModules(app *RegenApp, interfaceRegistry types.InterfaceRegistry) *server.Manager {
	return server.NewManager(app.BaseApp, codec.NewProtoCodec(interfaceRegistry))
}
func setCustomKVStoreKeys() []string {
	return []string{}
}

func (app *RegenApp) registerUpgradeHandlers() {

	// This is the upgrade plan name we used in the gov proposal.
	const upgradeName = "v3.0.0"
	app.UpgradeKeeper.SetUpgradeHandler(upgradeName, func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// TODO: once integrating basepp need to register a proper ecocredit consensus version
		// fromVM[ecocredit.ModuleName] = ecocreditmodule.Module{}.ConsensusVersion() - 1

		ecoParams, ok := app.ParamsKeeper.GetSubspace(ecocredit.ModuleName)
		if !ok {
			return nil, fmt.Errorf("unable to upgrade: subspace %s not found", ecocredit.ModuleName)
		}

		// set basket creation fee to 1'0000 REGEN
		ecoParams.Set(ctx, ecocredit.KeyBasketCreationFee, sdk.NewCoins(sdk.NewInt64Coin("uregen", 1e9)))

		return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	})
}

func (app *RegenApp) setCustomModuleManager() []module.AppModule {
	return []module.AppModule{}
}

func (app *RegenApp) setCustomKeeprs(_ *baseapp.BaseApp, keys map[string]*sdk.KVStoreKey, appCodec codec.Codec, _ govtypes.Router, _ string) {
}

func setCustomOrderInitGenesis() []string {
	return []string{}
}

func (app *RegenApp) setCustomSimulationManager() []module.AppModuleSimulation {
	return []module.AppModuleSimulation{}
}

func initCustomParamsKeeper(_ *paramskeeper.Keeper) {}
