//go:build !nosimulation

package simulation

import (
	"fmt"
	"os"
	"testing"

	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/stretchr/testify/require"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"

	regen "github.com/regen-network/regen-ledger/v7/app"
)

func TestAppAfterImport(t *testing.T) {
	config := simcli.NewConfigFromFlags()
	config.ChainID = SimAppChainID

	//nolint:staticcheck // will be removed in next upgrade
	db, dir, logger, skip, err := simtestutil.SetupSimulation(config, "leveldb-app-sim", "Simulation", simcli.FlagVerboseValue, simcli.FlagEnabledValue)
	if skip {
		t.Skip("skipping application import/export simulation")
	}
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, db.Close())
		require.NoError(t, os.RemoveAll(dir))
	}()

	appOptions := make(simtestutil.AppOptionsMap, 0)
	//nolint:staticcheck // will be removed in next upgrade
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue
	appOptions[flags.FlagHome] = t.TempDir()

	//nolint:staticcheck // will be removed in next upgrade
	app := regen.NewRegenApp(logger, db, nil, true, simcli.FlagPeriodValue, appOptions, emptyWasmOption, fauxMerkleModeOpt, baseapp.SetChainID(SimAppChainID))
	require.Equal(t, "regen", app.Name())

	// Run randomized simulation
	stopEarly, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		simtestutil.AppStateFn(app.AppCodec(), app.SimulationManager(), app.DefaultGenesis()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		SimulationOperations(app, app.AppCodec(), config),
		app.BlockAddresses(),
		config,
		app.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	err = simtestutil.CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		simtestutil.PrintStats(db)
	}

	if stopEarly {
		fmt.Println("can't export or import a zero-validator genesis, exiting test...")
		return
	}

	fmt.Printf("exporting genesis...\n")

	exported, err := app.ExportAppStateAndValidators(true, []string{}, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	//nolint: staticcheck // will be removed in next upgrade
	newDB, newDir, _, _, err := simtestutil.SetupSimulation(config, "leveldb-app-sim-2", "Simulation-2", simcli.FlagVerboseValue, simcli.FlagEnabledValue)
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, newDB.Close())
		require.NoError(t, os.RemoveAll(newDir))
	}()

	newAppOptions := make(simtestutil.AppOptionsMap, 0)
	newAppOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue
	newAppOptions[flags.FlagHome] = t.TempDir()

	//nolint: staticcheck // will be removed in next upgrade
	newApp := regen.NewRegenApp(log.NewNopLogger(), newDB, nil, true, simcli.FlagPeriodValue, newAppOptions, emptyWasmOption, fauxMerkleModeOpt, baseapp.SetChainID(SimAppChainID))
	require.Equal(t, "regen", newApp.Name())

	_, err = newApp.InitChain(&abci.RequestInitChain{
		ChainId:       SimAppChainID,
		AppStateBytes: exported.AppState,
	})
	require.NoError(t, err)

	_, _, err = simulation.SimulateFromSeed(
		t,
		os.Stdout,
		newApp.BaseApp,
		simtestutil.AppStateFn(app.AppCodec(), app.SimulationManager(), app.DefaultGenesis()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		SimulationOperations(newApp, newApp.AppCodec(), config),
		app.BlockAddresses(),
		config,
		app.AppCodec(),
	)
	require.NoError(t, err)
}
