package v7_3 //nolint:revive,stylecheck

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/regen-network/regen-ledger/v7/app/upgrades"
)

const Name = "v7.3.0"

var Upgrade = upgrades.Upgrade{
	UpgradeName: Name,
	CreateUpgradeHandler: func(manager *module.Manager, _ *wasmkeeper.Keeper, configurator module.Configurator) upgradetypes.UpgradeHandler {
		return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			// Run only migrations required by the stored module versions. Existing
			// v7 state keeps its stores and parameters, including wasm permissions.
			return manager.RunMigrations(ctx, configurator, fromVM)
		}
	},
	StoreUpgrades: storetypes.StoreUpgrades{},
}
