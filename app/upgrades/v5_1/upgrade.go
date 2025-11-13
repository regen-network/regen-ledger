package v5_1 //nolint:revive,stylecheck

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/regen-network/regen-ledger/v7/app/upgrades"
)

const Name = "v5.1"

var Upgrade = upgrades.Upgrade{
	UpgradeName: Name,
	CreateUpgradeHandler: func(mm *module.Manager, _ *wasmkeeper.Keeper, cfg module.Configurator) upgradetypes.UpgradeHandler {
		return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			// run in-place store migrations for ecocredit module
			return mm.RunMigrations(ctx, cfg, fromVM)
		}
	},
	StoreUpgrades: storetypes.StoreUpgrades{},
}
