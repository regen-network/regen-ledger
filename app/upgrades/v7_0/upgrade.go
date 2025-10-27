package v7_0 //nolint:revive,stylecheck

import (
	"context"
	"fmt"

	storetypes "cosmossdk.io/store/types"
	circuittypes "cosmossdk.io/x/circuit/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	protocolpooltypes "github.com/cosmos/cosmos-sdk/x/protocolpool/types"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/CosmWasm/wasmd/x/wasm/types"
	ibcwasmtypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v10/types"

	"github.com/regen-network/regen-ledger/v7/app/upgrades"
)

const Name = "v7_0"

var Upgrade = upgrades.Upgrade{
	UpgradeName: Name,
	CreateUpgradeHandler: func(manager *module.Manager, wasmKeeper *wasmkeeper.Keeper, configurator module.Configurator) upgradetypes.UpgradeHandler {
		return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			sdkCtx.Logger().Info("Starting module migrations...")
			vmManager, err := manager.RunMigrations(ctx, configurator, fromVM)
			if err != nil {
				return nil, err
			}

			govModuleAddrBytes := authtypes.NewModuleAddress(govtypes.ModuleName)

			params := wasmKeeper.GetParams(sdkCtx)
			params.InstantiateDefaultPermission = types.AccessTypeAnyOfAddresses
			params.CodeUploadAccess = types.AccessConfig{
				Permission: types.AccessTypeAnyOfAddresses,
				Addresses:  []string{govModuleAddrBytes.String()},
			}

			if err := wasmKeeper.SetParams(ctx, params); err != nil {
				return nil, err
			}

			sdkCtx.Logger().Info(fmt.Sprintf("Migration %s completed", Name))

			return vmManager, nil
		}
	},
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			circuittypes.ModuleName,
			protocolpooltypes.ModuleName,
			ibcwasmtypes.ModuleName,
		},
	},
}
