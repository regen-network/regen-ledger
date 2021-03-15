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
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
)

func setCustomModuleBasics() []module.AppModuleBasic {
	return []module.AppModuleBasic{
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distrclient.ProposalHandler,
				upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
		),
	}
}

func setCustomKVStoreKeys () []string {
	return []string{}
}

func setCustomModules(app *RegenApp, interfaceRegistry types.InterfaceRegistry) {}

func (app *RegenApp) registerUpgradeHandlers() {}

func (app *RegenApp) setCustomModuleManager() []module.AppModule {
	return []module.AppModule{}
}

func (app *RegenApp) setCustomKeeprs(bApp *baseapp.BaseApp, keys map[string]*sdk.KVStoreKey, appCodec codec.Marshaler, govRouter govtypes.Router) {

}
