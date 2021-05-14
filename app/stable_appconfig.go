// +build !experimental

package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
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
		feegrantmodule.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
	}
}

// setCustomModules registers new modules with the server module manager.
// It does nothing here and returns an empty manager since we're not using experimental mode.
func setCustomModules(_ *RegenApp, _ types.InterfaceRegistry) *server.Manager {
	return &server.Manager{}
}
func setCustomKVStoreKeys() []string {
	return []string{
		feegrant.StoreKey,
		authzkeeper.StoreKey,
	}
}

func (app *RegenApp) registerUpgradeHandlers() {}

func (app *RegenApp) setCustomModuleManager() []module.AppModule {
	return []module.AppModule{
		feegrantmodule.NewAppModule(app.appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(app.appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
	}
}

func (app *RegenApp) setCustomKeeprs(_ *baseapp.BaseApp, keys map[string]*sdk.KVStoreKey, appCodec codec.Codec, _ govtypes.Router, _ string) {
	feegrantKeeper := feegrantkeeper.NewKeeper(
		appCodec, keys[feegrant.StoreKey], &app.AccountKeeper,
	)
	app.FeeGrantKeeper = feegrantKeeper

	authzKeeper := authzkeeper.NewKeeper(
		keys[authzkeeper.StoreKey], appCodec, app.MsgServiceRouter(),
	)
	app.AuthzKeeper = authzKeeper
}

func setCustomOrderInitGenesis() []string {
	return []string{
		feegrant.ModuleName,
		authz.ModuleName,
	}
}

func (app *RegenApp) setCustomSimulationManager() []module.AppModuleSimulation {
	return []module.AppModuleSimulation{
		feegrantmodule.NewAppModule(app.appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(app.appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
	}
}

func initCustomParamsKeeper(_ *paramskeeper.Keeper) {}
