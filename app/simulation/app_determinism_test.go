//go:build !nosimulation

package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/libs/log"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"

	regen "github.com/regen-network/regen-ledger/v5/app"
	"github.com/regen-network/regen-ledger/v5/app/testsuite"
)

func TestAppDeterminism(t *testing.T) {
	if !simcli.FlagEnabledValue {
		t.Skip("skipping app-determinism simulation")
	}

	config := simcli.NewConfigFromFlags()
	config.InitialBlockHeight = 1
	config.ExportParamsPath = ""
	config.OnOperation = false
	config.AllInvariants = false
	config.ChainID = SimAppChainID

	numSeeds := 3
	numTimesToRunPerSeed := 3
	appHashList := make([]json.RawMessage, numTimesToRunPerSeed)

	for i := 0; i < numSeeds; i++ {
		config.Seed = rand.Int63()

		for j := 0; j < numTimesToRunPerSeed; j++ {
			var logger log.Logger
			if simcli.FlagVerboseValue {
				logger = log.TestingLogger()
			} else {
				logger = log.NewNopLogger()
			}

			db := dbm.NewMemDB()
			app := regen.NewRegenApp(
				logger,
				db,
				nil,
				true,
				testsuite.EmptyAppOptions{},
			)
			require.Equal(t, regen.AppName, app.Name())

			fmt.Printf(
				"running app-determinism simulation; seed %d: %d/%d, attempt: %d/%d\n",
				config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
			)

			_, _, err := simulateFromSeed(t, app, config)

			require.NoError(t, err)

			if config.Commit {
				simtestutil.PrintStats(db)
			}

			appHash := app.LastCommitID().Hash
			appHashList[j] = appHash

			if j != 0 {
				require.Equal(
					t, string(appHashList[0]), string(appHashList[j]),
					"non-determinism in seed %d: %d/%d, attempt: %d/%d\n",
					config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
				)
			}
		}
	}
}
