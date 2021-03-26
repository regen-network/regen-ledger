// +build experimental

package app

import (
	"path/filepath"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/gov"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	moduletypes "github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/types/module/server"
	data "github.com/regen-network/regen-ledger/x/data/module"
	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit/module"
	group "github.com/regen-network/regen-ledger/x/group/module"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{
		gov.NewAppModuleBasic(
			append(wasmclient.ProposalHandlers, paramsclient.ProposalHandler, distrclient.ProposalHandler,
				upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler)...,
		),
		data.Module{},
		ecocredit.Module{},
		group.Module{},
		wasm.AppModuleBasic{},
	}
}

func setCustomKVStoreKeys() []string {
	return []string{wasm.StoreKey}
}

func (app *RegenApp) setCustomKeeprs(bApp *baseapp.BaseApp, keys map[string]*sdk.KVStoreKey, appCodec codec.Marshaler, govRouter govtypes.Router, homePath string) {
	// just re-use the full router - do we want to limit this more?
	var wasmRouter = bApp.Router()
	wasmDir := filepath.Join(homePath, "wasm")

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	supportedFeatures := "staking"
	app.wasmKeeper = wasm.NewKeeper(
		appCodec,
		keys[wasm.StoreKey],
		app.GetSubspace(wasm.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.DistrKeeper,
		wasmRouter,
		wasmDir,
		getWasmConfig(),
		supportedFeatures,
		nil,
		nil,
	)

	// The gov proposal types can be individually enabled
	govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(app.wasmKeeper, wasm.EnableAllProposals))
}

// setCustomModules registers new modules with the server module manager.
func setCustomModules(app *RegenApp, interfaceRegistry types.InterfaceRegistry) *server.Manager {

	/* New Module Wiring START */
	newModuleManager := server.NewManager(app.BaseApp, codec.NewProtoCodec(interfaceRegistry))

	// BEGIN HACK: this is a total, ugly hack until x/auth supports ADR 033 or we have a suitable alternative
	groupModule := group.Module{AccountKeeper: app.AccountKeeper}
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

	return newModuleManager
	/* New Module Wiring END */
}

func (app *RegenApp) registerUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler("Mau", func(ctx sdk.Context, plan upgradetypes.Plan) {
		// no-op handler, does nothing
	})
}

func (app *RegenApp) setCustomModuleManager() []module.AppModule {
	return []module.AppModule{
		wasm.NewAppModule(&app.wasmKeeper),
	}
}

func setCustomOrderInitGenesis() []string {
	return []string{
		wasm.ModuleName,
	}
}

func (app *RegenApp) setCustomSimulationManager() []module.AppModuleSimulation {
	return []module.AppModuleSimulation{
		wasm.NewAppModule(&app.wasmKeeper),
	}
}

func initCustomParamsKeeper(paramsKeeper *paramskeeper.Keeper) {
	paramsKeeper.Subspace(wasm.ModuleName)
}
