package v5

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/group"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ica "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts"
	icacontrollertypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v5/modules/apps/29-fee/types"

	"github.com/regen-network/regen-ledger/v4/app/upgrades"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

const Name = "v5"

var Upgrade = upgrades.Upgrade{
	UpgradeName: Name,
	CreateUpgradeHandler: func(mm *module.Manager, cfg module.Configurator) upgradetypes.UpgradeHandler {
		return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			// set regen module consensus version
			fromVM[ecocredit.ModuleName] = 2
			fromVM[data.ModuleName] = 1

			// save oldIcaVersion, so we can skip icahost.InitModule in longer term tests.
			oldIcaVersion := fromVM[icatypes.ModuleName]

			// Add Interchain Accounts host module
			// set the ICS27 consensus version so InitGenesis is not run
			fromVM[icatypes.ModuleName] = mm.Modules[icatypes.ModuleName].ConsensusVersion()

			// create ICS27 Controller submodule params, controller module not enabled.
			controllerParams := icacontrollertypes.Params{ControllerEnabled: false}

			// create ICS27 Host submodule params, host module not enabled.
			hostParams := icahosttypes.Params{
				HostEnabled:   false,
				AllowMessages: []string{},
			}

			mod, found := mm.Modules[icatypes.ModuleName]
			if !found {
				panic(fmt.Sprintf("module %s is not in the module manager", icatypes.ModuleName))
			}

			icaMod, ok := mod.(ica.AppModule)
			if !ok {
				panic(fmt.Sprintf("expected module %s to be type %T, got %T", icatypes.ModuleName, ica.AppModule{}, mod))
			}

			// skip InitModule in upgrade tests after the upgrade has gone through.
			if oldIcaVersion != fromVM[icatypes.ModuleName] {
				icaMod.InitModule(ctx, controllerParams, hostParams)
			}

			// transfer module consensus version has been bumped to 2
			return mm.RunMigrations(ctx, cfg, fromVM)
		}
	},
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			group.ModuleName,
			icahosttypes.StoreKey,
			ibcfeetypes.StoreKey,
			icacontrollertypes.StoreKey,
		},
	},
}
