//go:build !nosimulation

package simulation

import (
	"fmt"
	"os"
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/stretchr/testify/require"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"

	regen "github.com/regen-network/regen-ledger/v5/app"
)

func TestAppAfterImport(t *testing.T) {
	config := simcli.NewConfigFromFlags()
	db, dir, logger, skip, err := simtestutil.SetupSimulation(config, "app-after-import-1", "sim-1", false, true)
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
		simcli.FlagPeriodValue,
		simtestutil.EmptyAppOptions{},
		fauxMerkleModeOpt,
	)
	require.Equal(t, regen.AppName, app.Name())

	// run randomized simulation
	stopEarly, simParams, simErr := simulateFromSeed(t, app, config)

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

	exported, err := app.ExportAppStateAndValidators(true, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	newDB, newDir, _, _, err := simtestutil.SetupSimulation(config, "app-after-import-2", "sim-2", false, true)
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
		simcli.FlagPeriodValue,
		simtestutil.EmptyAppOptions{},
		fauxMerkleModeOpt,
	)
	require.Equal(t, regen.AppName, newApp.Name())

	newApp.InitChain(abci.RequestInitChain{
		AppStateBytes: exported.AppState,
	})

	_, _, err = simulateFromSeed(t, newApp, config)
	require.NoError(t, err)
}
