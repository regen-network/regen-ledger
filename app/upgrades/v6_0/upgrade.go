package v6_0 //nolint:revive,stylecheck
import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/regen-network/regen-ledger/v7/app/upgrades"
)

const Name = "v6_0"

var Upgrade = upgrades.Upgrade{
	UpgradeName: Name,
	CreateUpgradeHandler: func(manager *module.Manager, configurator module.Configurator) upgradetypes.UpgradeHandler {
		return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			ctx.Logger().Info("Starting module migrations...")
			vmManager, err := manager.RunMigrations(ctx, configurator, fromVM)
			if err != nil {
				return nil, err
			}

			ctx.Logger().Info(fmt.Sprintf("Migration %s completed", Name))

			return vmManager, nil

		}
	},
	StoreUpgrades: storetypes.StoreUpgrades{},
}
