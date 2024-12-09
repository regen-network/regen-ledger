//go:build !nosimulation

package simulation

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/server"
	"os"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	regen "github.com/regen-network/regen-ledger/v5/app"
	"github.com/regen-network/regen-ledger/x/data/v3"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
)

func TestAppImportExport(t *testing.T) {
	config := simcli.NewConfigFromFlags()
	//config.Commit = true
	db, dir, logger, skip, err := simtestutil.SetupSimulation(config, "app-import-export", "sim-1", false, true)
	if skip {
		t.Skip("skipping app-import-export simulation")
	}
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		db.Close()
		require.NoError(t, os.RemoveAll(dir))
	}()
	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue

	app := regen.NewRegenApp(
		logger,
		db,
		nil,
		true,
		simcli.FlagPeriodValue,
		appOptions,
		fauxMerkleModeOpt,
	)
	require.Equal(t, regen.AppName, app.Name())

	// run randomized simulation
	_, simParams, simErr := simulateFromSeed(t, app, config)

	// export state and simParams before the simulation error is checked
	err = simtestutil.CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		simtestutil.PrintStats(db)
	}

	fmt.Printf("exporting genesis...\n")

	exported, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	newDB, newDir, _, _, err := simtestutil.SetupSimulation(config, "app-import-export-2", "sim-2", false, true)
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, newDB.Close())
		require.NoError(t, os.RemoveAll(newDir))
	}()

	newApp := regen.NewRegenApp(
		logger,
		db,
		nil,
		true,
		simcli.FlagPeriodValue,
		appOptions,
		fauxMerkleModeOpt,
	)
	require.Equal(t, regen.AppName, newApp.Name())

	var genesisState regen.GenesisState
	err = json.Unmarshal(exported.AppState, &genesisState)
	require.NoError(t, err)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("%v", r)
			if !strings.Contains(err, "validator set is empty after InitGenesis") {
				panic(r)
			}
			logger.Info("Skipping simulation as all validators have been unbonded")
			logger.Info("err", err, "stacktrace", string(debug.Stack()))
		}
	}()

	ctxA := app.NewContext(true, tmproto.Header{Height: app.LastBlockHeight()})
	ctxB := newApp.NewContext(true, tmproto.Header{Height: app.LastBlockHeight()})

	newApp.ModuleManager.InitGenesis(ctxB, app.AppCodec(), genesisState)
	newApp.StoreConsensusParams(ctxB, exported.ConsensusParams)

	fmt.Printf("comparing stores...\n")

	storeKeysPrefixes := []StoreKeysPrefixes{
		{app.GetKey(authtypes.StoreKey), newApp.GetKey(authtypes.StoreKey), [][]byte{}},
		{app.GetKey(banktypes.StoreKey), newApp.GetKey(banktypes.StoreKey), [][]byte{
			banktypes.BalancesPrefix,
		}},
		{app.GetKey(capabilitytypes.StoreKey), newApp.GetKey(capabilitytypes.StoreKey), [][]byte{}},
		{app.GetKey(distrtypes.StoreKey), newApp.GetKey(distrtypes.StoreKey), [][]byte{}},
		{app.GetKey(evidencetypes.StoreKey), newApp.GetKey(evidencetypes.StoreKey), [][]byte{}},
		{app.GetKey(govtypes.StoreKey), newApp.GetKey(govtypes.StoreKey), [][]byte{}},
		{app.GetKey(group.ModuleName), newApp.GetKey(group.ModuleName), [][]byte{}},
		{app.GetKey(minttypes.StoreKey), newApp.GetKey(minttypes.StoreKey), [][]byte{}},
		// {app.GetKey(paramtypes.StoreKey), newApp.GetKey(paramtypes.StoreKey), [][]byte{}}, // FIXME
		{app.GetKey(slashingtypes.StoreKey), newApp.GetKey(slashingtypes.StoreKey), [][]byte{}},
		{app.GetKey(stakingtypes.StoreKey), newApp.GetKey(stakingtypes.StoreKey), [][]byte{
			stakingtypes.UnbondingQueueKey,
			stakingtypes.RedelegationQueueKey,
			stakingtypes.ValidatorQueueKey,
			stakingtypes.HistoricalInfoKey,
			stakingtypes.UnbondingIDKey,
			stakingtypes.UnbondingIndexKey,
			stakingtypes.UnbondingTypeKey,
			stakingtypes.ValidatorUpdatesKey,
		}},

		// ibc modules
		{app.GetKey(ibcexported.StoreKey), newApp.GetKey(ibcexported.StoreKey), [][]byte{}},
		{app.GetKey(ibctransfertypes.StoreKey), newApp.GetKey(ibctransfertypes.StoreKey), [][]byte{}},

		// regen modules
		{app.GetKey(data.ModuleName), newApp.GetKey(data.ModuleName), [][]byte{}},
		{app.GetKey(ecocredit.ModuleName), newApp.GetKey(ecocredit.ModuleName), [][]byte{}},
	}

	for i, skp := range storeKeysPrefixes {
		storeA := ctxA.KVStore(skp.A)
		storeB := ctxB.KVStore(skp.B)
		fmt.Println(i)

		failedKVAs, failedKVBs := sdk.DiffKVStores(storeA, storeB, skp.Prefixes)
		require.Equal(t, len(failedKVAs), len(failedKVBs), "unequal sets of key-values to compare")

		fmt.Printf(
			"compared %d different key/value pairs between %s and %s\n",
			len(failedKVAs), skp.A, skp.B,
		)

		simLog := simtestutil.GetSimulationLog(
			skp.A.Name(),
			app.SimulationManager().StoreDecoders,
			failedKVAs,
			failedKVBs,
		)
		require.Equal(t, len(failedKVAs), 0, simLog)
	}
}
