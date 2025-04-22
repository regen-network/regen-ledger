package simulation

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"

	regen "github.com/regen-network/regen-ledger/v6/app"
)

// Get flags every time the simulator is run
func init() {
	simcli.GetSimulatorFlags()
}

const SimAppChainID = "simulation-app"

type StoreKeysPrefixes struct {
	A        storetypes.StoreKey
	B        storetypes.StoreKey
	Prefixes [][]byte
}

// fauxMerkleModeOpt returns a BaseApp option to use a dbStoreAdapter instead of
// an IAVLStore for faster simulation speed.
func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
}

// interBlockCacheOpt returns a BaseApp option function that sets the persistent
// inter-block write-through cache.
func interBlockCacheOpt() func(*baseapp.BaseApp) {
	return baseapp.SetInterBlockCache(store.NewCommitKVStoreCacheManager())
}

func simulateFromSeed(t *testing.T, app *regen.RegenApp, config simtypes.Config) (bool, simulation.Params, error) {
	cdc := app.AppCodec()
	genesis := regen.NewDefaultGenesisState(cdc)
	return simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		simtestutil.AppStateFn(cdc, app.SimulationManager(), genesis),
		simtypes.RandomAccounts,
		simtestutil.SimulationOperations(app, cdc, config),
		app.BlockAddresses(),
		config,
		app.AppCodec(),
	)
}
