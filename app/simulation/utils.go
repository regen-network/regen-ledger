package simulation

import (
	"encoding/json"
	"flag"
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
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"

	regen "github.com/regen-network/regen-ledger/v7/app"
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
		SimulationOperations(app, cdc, config),
		app.BlockAddresses(),
		config,
		app.AppCodec(),
	)
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
//nolint:revive
func SimulationOperations(app runtime.AppI, cdc codec.JSONCodec, config simtypes.Config) []simtypes.WeightedOperation {
	return BuildSimulationOperations(app, cdc, config, MakeTestTxConfig())
}

func BuildSimulationOperations(app runtime.AppI, cdc codec.JSONCodec, config simtypes.Config, txConfig client.TxConfig) []simtypes.WeightedOperation {
	simState := module.SimulationState{
		AppParams: make(simtypes.AppParams),
		Cdc:       cdc,
		TxConfig:  txConfig,
		BondDenom: "regen",
	}

	if config.ParamsFile != "" {
		bz, err := os.ReadFile(config.ParamsFile)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bz, &simState.AppParams)
		if err != nil {
			panic(err)
		}
	}

	simState.LegacyProposalContents = app.SimulationManager().GetProposalContents(simState) //nolint:staticcheck // we're testing the old way here
	simState.ProposalMsgs = app.SimulationManager().GetProposalMsgs(simState)
	return app.SimulationManager().WeightedOperations(simState)
}
