package simulation

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"testing"

	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/testutil"
	"github.com/cosmos/cosmos-sdk/runtime"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"

	regen "github.com/regen-network/regen-ledger/v7/app"

	dbm "github.com/cosmos/cosmos-db"
)

var FlagEnableStreamingValue bool

// Get flags every time the simulator is run
func init() {
	simcli.GetSimulatorFlags()
	flag.BoolVar(&FlagEnableStreamingValue, "EnableStreaming", false, "Enable streaming service")
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
	t.Helper()
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

// CheckExportSimulation exports the app state and simulation parameters to JSON
// if the export paths are defined.
func CheckExportSimulation(app runtime.AppI, config simtypes.Config, params simtypes.Params) error {
	if config.ExportStatePath != "" {
		fmt.Println("exporting app state...")
		exported, err := app.ExportAppStateAndValidators(false, nil, nil)
		if err != nil {
			return err
		}

		if err := os.WriteFile(config.ExportStatePath, []byte(exported.AppState), 0o600); err != nil {
			return err
		}
	}

	if config.ExportParamsPath != "" {
		fmt.Println("exporting simulation params...")
		paramsBz, err := json.MarshalIndent(params, "", " ")
		if err != nil {
			return err
		}

		if err := os.WriteFile(config.ExportParamsPath, paramsBz, 0o600); err != nil {
			return err
		}
	}
	return nil
}

// PrintStats prints the corresponding statistics from the app DB.
func PrintStats(db dbm.DB) {
	fmt.Println("\nLevelDB Stats")
	fmt.Println(db.Stats()["leveldb.stats"])
	fmt.Println("LevelDB cached block size", db.Stats()["leveldb.cachedblock"])
}

func MakeTestTxConfig() client.TxConfig {
	interfaceRegistry := testutil.CodecOptions{
		AccAddressPrefix: "regen",
		ValAddressPrefix: "regenvaloper",
	}.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	return tx.NewTxConfig(cdc, tx.DefaultSignModes)
}

// Operations retrieves the simulation params from the provided file path
// and returns all the modules weighted operations
//
// nolint:revive
func SimulationOperations(app runtime.AppI, cdc codec.JSONCodec, config simtypes.Config) []simtypes.WeightedOperation {
	return simtestutil.BuildSimulationOperations(app, cdc, config, MakeTestTxConfig())
}
