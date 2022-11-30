//go:build !nosimulation

package simulation

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/simapp"

	regen "github.com/regen-network/regen-ledger/v5/app"
)

func TestApp(t *testing.T) {
	config, db, dir, logger, skip, err := simapp.SetupSimulation("app", "simulation")
	if skip {
		t.Skip("skipping app simulation")
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
	_, simParams, simErr := simulateFromSeed(t, app, config)

	// export state and simParams before the simulation error is checked
	err = simapp.CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		simapp.PrintStats(db)
	}
}
