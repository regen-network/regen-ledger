//go:build !nosimulation

package simulation

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client/flags"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"

	regen "github.com/regen-network/regen-ledger/v7/app"
)

func TestApp(t *testing.T) {
	config := simcli.NewConfigFromFlags()
	config.ChainID = SimAppChainID

	// nolint:staticcheck
	db, dir, logger, skip, err := simtestutil.SetupSimulation(config, "leveldb-app-sim", "Simulation", simcli.FlagVerboseValue, simcli.FlagEnabledValue)
	if skip {
		t.Skip("skipping app simulation")
	}
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		db.Close()
		require.NoError(t, os.RemoveAll(dir))
	}()

	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = regen.DefaultNodeHome

	app := regen.NewRegenApp(
		logger,
		db,
		nil,
		true,
		// nolint:staticcheck // deprecated but required for upgrade
		simcli.FlagPeriodValue,
		appOptions,
		emptyWasmOption,
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
}
