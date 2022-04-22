//go:build experimental
// +build experimental

// DONTCOVER

package app

import (
	"fmt"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"

	moduletypes "github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/client/core"
	group "github.com/regen-network/regen-ledger/x/group/module"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{
		gov.NewAppModuleBasic(
			append(
				wasmclient.ProposalHandlers,
				paramsclient.ProposalHandler, distrclient.ProposalHandler,
				upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
				core.CreditTypeProposalHandler,
			)...,
		),
		wasm.AppModuleBasic{},
		group.Module{},
	}
}

func setCustomKVStoreKeys() []string {
	return []string{wasm.StoreKey}
}

func setCustomMaccPerms() map[string][]string {
	return map[string][]string{
		wasm.ModuleName: {authtypes.Burner},
	}
}

func (app *RegenApp) setCustomKeepers(bApp *baseapp.BaseApp, keys map[string]*sdk.KVStoreKey, appCodec codec.Codec,
	govRouter govtypes.Router, homePath string, appOpts servertypes.AppOptions,
	wasmOpts []wasm.Option) {
	wasmDir := filepath.Join(homePath, "wasm")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	app.wasmCfg = wasmConfig

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	supportedFeatures := "iterator,staking,stargate"
	app.wasmKeeper = wasm.NewKeeper(
		appCodec,
		keys[wasm.StoreKey],
		app.GetSubspace(wasm.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.DistrKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.scopedWasmKeeper,
		app.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		supportedFeatures,
		wasmOpts...,
	)
}

// setCustomModules registers new modules with the server module manager.
func setCustomModules(app *RegenApp, interfaceRegistry types.InterfaceRegistry) *server.Manager {

	/* New Module Wiring START */
	newModuleManager := server.NewManager(app.BaseApp, codec.NewProtoCodec(interfaceRegistry))

	// BEGIN HACK: this is a total, ugly hack until x/auth & x/bank supports ADR 033 or we have a suitable alternative

	groupModule := group.Module{AccountKeeper: app.AccountKeeper, BankKeeper: app.BankKeeper}
	// use a separate newModules from the global NewModules here because we need to pass state into the group module
	newModules := []moduletypes.Module{
		groupModule,
	}
	err := newModuleManager.RegisterModules(newModules)
	if err != nil {
		panic(err)
	}
	// END HACK

	/* New Module Wiring END */
	return newModuleManager
}

func (app *RegenApp) registerUpgradeHandlers() {}

func (app *RegenApp) setCustomModuleManager() []module.AppModule {
	return []module.AppModule{
		wasm.NewAppModule(app.appCodec, &app.wasmKeeper, app.StakingKeeper),
	}
}

func setCustomOrderInitGenesis() []string {
	return []string{
		// wasm after ibc transfer
		wasm.ModuleName,
	}
}

func setCustomOrderBeginBlocker() []string {
	return []string{
		wasm.ModuleName,
	}
}

func setCustomOrderEndBlocker() []string {
	return []string{
		wasm.ModuleName,
	}
}

func (app *RegenApp) setCustomAnteHandler(encCfg simappparams.EncodingConfig,
	wasmKey *sdk.KVStoreKey, wasmCfg *wasm.Config) (sdk.AnteHandler, error) {
	return NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				FeegrantKeeper:  app.FeeGrantKeeper,
				SignModeHandler: encCfg.TxConfig.SignModeHandler(),
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			IBCChannelKeeper:  app.IBCKeeper.ChannelKeeper,
			WasmConfig:        wasmCfg,
			TXCounterStoreKey: wasmKey,
		},
	)

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
	paramsKeeper.Subspace(wasm.ModuleName)
}

func (app *RegenApp) initializeCustomScopedKeepers() {
	app.scopedWasmKeeper = app.CapabilityKeeper.ScopeToModule(wasm.ModuleName)
}
