package simulation

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/simapp"

	regen "github.com/regen-network/regen-ledger/v4/app"
)

func TestApp(t *testing.T) {
	cfg, db, dir, logger, skip, err := simapp.SetupSimulation("app-simulation", "simulation")
	if skip {
		t.Skip("skipping application simulation")
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
	_, simParams, simErr := simulateFromSeed(t, app, cfg)

	// export state and simParams before the simulation error is checked
	err = simapp.CheckExportSimulation(app, cfg, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if cfg.Commit {
		simapp.PrintStats(db)
	}
}
