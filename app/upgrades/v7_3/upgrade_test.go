package v7_3_test

import (
	"context"
	"testing"
	"time"

	"cosmossdk.io/core/header"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	abci "github.com/cometbft/cometbft/abci/types"
	cmttypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/regen-network/regen-ledger/v7/app"
	"github.com/regen-network/regen-ledger/v7/app/testsuite"
	"github.com/regen-network/regen-ledger/v7/app/upgrades/v7_3"
)

const upgradeHeight = int64(5)

func setup(t *testing.T) (*app.RegenApp, sdk.Context) {
	t.Helper()
	options := testsuite.DefaultOptions()
	options.AppOpts = simtestutil.NewAppOptionsWithFlagHome(t.TempDir())
	a := testsuite.NewAppWithCustomOptions(t, false, options)
	t.Cleanup(func() { require.NoError(t, a.Close()) })
	ctx := a.NewContextLegacy(false, cmttypes.Header{
		Height: 1, ChainID: "regen-upgrade-test", Time: time.Now().UTC(),
	})
	ctx = atHeight(ctx, 1)
	require.True(t, a.UpgradeKeeper.HasHandler(v7_3.Name))
	require.NoError(t, a.UpgradeKeeper.ScheduleUpgrade(ctx, upgradetypes.Plan{
		Name: v7_3.Name, Height: upgradeHeight,
	}))
	return a, ctx
}

func atHeight(ctx sdk.Context, height int64) sdk.Context {
	return ctx.WithBlockHeight(height).WithHeaderInfo(header.Info{
		Height: height, ChainID: ctx.ChainID(), Time: ctx.BlockTime(),
	})
}

func TestUpgradePreservesV7State(t *testing.T) {
	a, ctx := setup(t)

	// Use non-default permissions so accidentally re-running the v7_0 handler
	// would be detected, even if a production chain has since changed them.
	params := a.WasmKeeper.GetParams(ctx)
	params.CodeUploadAccess = wasmtypes.AccessConfig{
		Permission: wasmtypes.AccessTypeAnyOfAddresses,
		Addresses:  []string{testsuite.CreateRandomAccounts(1)[0].String()},
	}
	params.InstantiateDefaultPermission = wasmtypes.AccessTypeNobody
	require.NoError(t, a.WasmKeeper.SetParams(ctx, params))

	mintParams, err := a.MintKeeper.Params.Get(ctx)
	require.NoError(t, err)
	supply := a.BankKeeper.GetSupply(ctx, mintParams.MintDenom)
	icaParams := a.ICAHostKeeper.GetParams(ctx)
	versions, err := a.UpgradeKeeper.GetModuleVersionMap(ctx)
	require.NoError(t, err)
	require.Equal(t, a.ModuleManager.GetVersionMap(), versions)

	ctx = atHeight(ctx, upgradeHeight)
	_, err = a.PreBlocker(ctx, &abci.RequestFinalizeBlock{Height: upgradeHeight})
	require.NoError(t, err)

	doneHeight, err := a.UpgradeKeeper.GetDoneHeight(ctx, v7_3.Name)
	require.NoError(t, err)
	require.Equal(t, upgradeHeight, doneHeight)
	_, err = a.UpgradeKeeper.GetUpgradePlan(ctx)
	require.ErrorIs(t, err, upgradetypes.ErrNoUpgradePlanFound)
	updatedVersions, err := a.UpgradeKeeper.GetModuleVersionMap(ctx)
	require.NoError(t, err)
	require.Equal(t, versions, updatedVersions)
	require.Equal(t, params, a.WasmKeeper.GetParams(ctx))
	updatedMintParams, err := a.MintKeeper.Params.Get(ctx)
	require.NoError(t, err)
	require.Equal(t, mintParams, updatedMintParams)
	require.Equal(t, supply, a.BankKeeper.GetSupply(ctx, mintParams.MintDenom))
	require.Equal(t, icaParams, a.ICAHostKeeper.GetParams(ctx))

	// Subsequent blocks must not execute this handler a second time.
	a.UpgradeKeeper.SetUpgradeHandler(v7_3.Name, func(context.Context, upgradetypes.Plan, module.VersionMap) (module.VersionMap, error) {
		t.Fatal("upgrade handler executed twice")
		return nil, nil
	})
	_, err = a.PreBlocker(atHeight(ctx, upgradeHeight+1), &abci.RequestFinalizeBlock{Height: upgradeHeight + 1})
	require.NoError(t, err)
}

func TestUpgradeRejectsEarlyActivation(t *testing.T) {
	a, ctx := setup(t)
	ctx = atHeight(ctx, upgradeHeight-1)
	_, err := a.PreBlocker(ctx, &abci.RequestFinalizeBlock{Height: upgradeHeight - 1})
	require.ErrorContains(t, err, "BINARY UPDATED BEFORE TRIGGER")
	doneHeight, err := a.UpgradeKeeper.GetDoneHeight(ctx, v7_3.Name)
	require.NoError(t, err)
	require.Zero(t, doneHeight)
	plan, err := a.UpgradeKeeper.GetUpgradePlan(ctx)
	require.NoError(t, err)
	require.Equal(t, v7_3.Name, plan.Name)
}

func TestUpgradeDoesNotCompleteOnMigrationFailure(t *testing.T) {
	a, ctx := setup(t)
	versions, err := a.UpgradeKeeper.GetModuleVersionMap(ctx)
	require.NoError(t, err)
	// SDK 0.53.6 removed the historical mint 1->2 migration. This unsupported
	// input must fail, proving the handler runs migrations rather than skipping them.
	versions[minttypes.ModuleName] = 1
	require.NoError(t, a.UpgradeKeeper.SetModuleVersionMap(ctx, versions))
	ctx = atHeight(ctx, upgradeHeight)
	_, err = a.PreBlocker(ctx, &abci.RequestFinalizeBlock{Height: upgradeHeight})
	require.ErrorContains(t, err, "migration")
	doneHeight, err := a.UpgradeKeeper.GetDoneHeight(ctx, v7_3.Name)
	require.NoError(t, err)
	require.Zero(t, doneHeight)
	plan, err := a.UpgradeKeeper.GetUpgradePlan(ctx)
	require.NoError(t, err)
	require.Equal(t, v7_3.Name, plan.Name)
}
