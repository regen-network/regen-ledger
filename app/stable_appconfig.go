// +build !experimental

package app

import (
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
	"github.com/regen-network/regen-ledger/types/module/server"
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
func setCustomModules(_ *RegenApp, _ types.InterfaceRegistry) *server.Manager {
	return &server.Manager{}
}
func setCustomKVStoreKeys() []string {
	return []string{}
}

func (app *RegenApp) registerUpgradeHandlers() {}

func (app *RegenApp) setCustomModuleManager() []module.AppModule {
	return []module.AppModule{}
}

func (app *RegenApp) setCustomKeeprs(_ *baseapp.BaseApp, _ map[string]*sdk.KVStoreKey, _ codec.Marshaler, _ govtypes.Router, _ string) {
}

func setCustomOrderInitGenesis() []string {
	return []string{}
}

func (app *RegenApp) setCustomSimulationManager() []module.AppModuleSimulation {
	return []module.AppModuleSimulation{}
}

func initCustomParamsKeeper(_ *paramskeeper.Keeper) {}
