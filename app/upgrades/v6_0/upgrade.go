package v6_0 //nolint:revive,stylecheck
import (
	"context"
	"fmt"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/regen-network/regen-ledger/v7/app/upgrades"
)

const Name = "v6_0"

var Upgrade = upgrades.Upgrade{
	UpgradeName: Name,
	CreateUpgradeHandler: func(manager *module.Manager, configurator module.Configurator) upgradetypes.UpgradeHandler {
		return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			sdkCtx.Logger().Info("Starting module migrations...")
			vmManager, err := manager.RunMigrations(ctx, configurator, fromVM)
			if err != nil {
				return nil, err
			}

			sdkCtx.Logger().Info(fmt.Sprintf("Migration %s completed", Name))

			return vmManager, nil
		}
	},
	StoreUpgrades: storetypes.StoreUpgrades{},
}
