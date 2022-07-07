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

	moduletypes "github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/client/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/client/marketplace"
	group "github.com/regen-network/regen-ledger/x/group/module"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{
		gov.NewAppModuleBasic(
			append(
				paramsclient.ProposalHandler, distrclient.ProposalHandler,
				upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
				core.CreditTypeProposalHandler, marketplace.AllowDenomProposalHandler,
			)...,
		),
		group.Module{},
	}
}

func setCustomKVStoreKeys() []string {
	return []string{}
}

func setCustomMaccPerms() map[string][]string {
	return map[string][]string{}
}

func (app *RegenApp) setCustomKeepers(bApp *baseapp.BaseApp, keys map[string]*sdk.KVStoreKey, appCodec codec.Codec,
	govRouter govtypes.Router, homePath string, appOpts servertypes.AppOptions) {
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
	return []module.AppModule{}
}

func setCustomOrderInitGenesis() []string {
	return []string{}
}

func setCustomOrderBeginBlocker() []string {
	return []string{}
}

func setCustomOrderEndBlocker() []string {
	return []string{}
}

func (app *RegenApp) setCustomAnteHandler(encCfg simappparams.EncodingConfig) (sdk.AnteHandler, error) {
	return nil
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

func (app *RegenApp) initializeCustomScopedKeepers() {
}
