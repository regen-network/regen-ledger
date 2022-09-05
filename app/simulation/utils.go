package simulation

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	ica "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts"

	regen "github.com/regen-network/regen-ledger/v4/app"
)

// Get flags every time the simulator is run
func init() {
	simapp.GetSimulatorFlags()
}

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
	return simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		simapp.AppStateFn(app.AppCodec(), app.SimulationManager()),
		simtypes.RandomAccounts,
		regen.SimulationOperations(app, app.AppCodec(), config),
		app.ModuleAccountAddrs(),
		config,
		app.AppCodec(),
	)
}

// removeICAFromSimulation is a utility function that removes from genesis exporting due to a panic bug.
//
// TODO: remove after https://github.com/cosmos/ibc-go/issues/2151 is resolved
func removeICAFromSimulation(app *regen.RegenApp) {
	remove := func(target string, mods []string) []string {
		for i, mod := range mods {
			if mod == target {
				return append(mods[:i], mods[i+1:]...)
			}
		}
		return mods
	}

	icaModName := ica.AppModule{}.Name()

	app.ModuleManager.OrderExportGenesis = remove(icaModName, app.ModuleManager.OrderExportGenesis)
}
