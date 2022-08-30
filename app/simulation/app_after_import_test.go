//go:build !nosimulation

package simulation

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/simapp"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	regen "github.com/regen-network/regen-ledger/v4/app"
)

func TestAppAfterImport(t *testing.T) {
	config, db, dir, logger, skip, err := simapp.SetupSimulation("app-after-import-1", "sim-1")
	if skip {
		t.Skip("skipping app-after-import simulation")
	}
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		db.Close()
		require.NoError(t, os.RemoveAll(dir))
	}()

	app := regen.NewRegenApp(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		regen.DefaultNodeHome,
		simapp.FlagPeriodValue,
		regen.MakeEncodingConfig(),
		simapp.EmptyAppOptions{},
		fauxMerkleModeOpt,
	)
	require.Equal(t, regen.AppName, app.Name())

	// run randomized simulation
	stopEarly, simParams, simErr := simulateFromSeed(t, app, config)

	// export state and simParams before the simulation error is checked
	err = simapp.CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		simapp.PrintStats(db)
	}

	if stopEarly {
		fmt.Println("can't export or import a zero-validator genesis, exiting test...")
		return
	}

	fmt.Printf("exporting genesis...\n")

	exported, err := app.ExportAppStateAndValidators(true, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	_, newDB, newDir, _, _, err := simapp.SetupSimulation("app-after-import-2", "sim-2")
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		newDB.Close()
		require.NoError(t, os.RemoveAll(newDir))
	}()

	newApp := regen.NewRegenApp(
		log.NewNopLogger(),
		newDB,
		nil,
		true,
		map[int64]bool{},
		regen.DefaultNodeHome,
		simapp.FlagPeriodValue,
		regen.MakeEncodingConfig(),
		simapp.EmptyAppOptions{},
		fauxMerkleModeOpt,
	)
	require.Equal(t, regen.AppName, newApp.Name())

	newApp.InitChain(abci.RequestInitChain{
		AppStateBytes: exported.AppState,
	})

	_, _, err = simulation.SimulateFromSeed(
		t,
		os.Stdout,
		newApp.BaseApp,
		simapp.AppStateFn(app.AppCodec(), app.SimulationManager()),
		simtypes.RandomAccounts,
		regen.SimulationOperations(newApp, newApp.AppCodec(), config),
		app.ModuleAccountAddrs(),
		config,
		app.AppCodec(),
	)
	require.NoError(t, err)
}
