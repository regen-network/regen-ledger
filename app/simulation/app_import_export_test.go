//go:build !nosimulation

package simulation

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/x/feegrant"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	storetypes "cosmossdk.io/store/types"
	sims "github.com/cosmos/cosmos-sdk/testutil/simsx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	regen "github.com/regen-network/regen-ledger/v7/app"
)

var emptyWasmOption []wasmkeeper.Option

// NewAppFactory creates a proper app factory function for sims.Run and sims.NewSimulationAppInstance
func NewAppFactory() func(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *regen.RegenApp {
	return func(
		logger log.Logger,
		db dbm.DB,
		traceStore io.Writer,
		loadLatest bool,
		appOpts servertypes.AppOptions,
		baseAppOptions ...func(*baseapp.BaseApp),
	) *regen.RegenApp {
		// Get the invCheckPeriod from appOpts, with a default value if not found
		var invCheckPeriod uint = 1
		if val := appOpts.Get(server.FlagInvCheckPeriod); val != nil {
			if period, ok := val.(uint); ok {
				invCheckPeriod = period
			}
		}

		return regen.NewRegenApp(
			logger,
			db,
			traceStore,
			loadLatest,
			invCheckPeriod,
			appOpts,
			emptyWasmOption,
			baseAppOptions...,
		)
	}
}

// Alternative factory function in case sims uses different types
func NewRegenAppFactoryForSims() func(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *regen.RegenApp {
	return func(
		logger log.Logger,
		db dbm.DB,
		traceStore io.Writer,
		loadLatest bool,
		appOpts servertypes.AppOptions,
		baseAppOptions ...func(*baseapp.BaseApp),
	) *regen.RegenApp {
		// Get the invCheckPeriod from appOpts, with a default value if not found
		var invCheckPeriod uint = 1
		if val := appOpts.Get(server.FlagInvCheckPeriod); val != nil {
			if period, ok := val.(uint); ok {
				invCheckPeriod = period
			}
		}
		// Convert sims.AppOptionsMap to servertypes.AppOptions
		return regen.NewRegenApp(
			logger,
			db,
			traceStore,
			loadLatest,
			invCheckPeriod,
			appOpts, // This might work if AppOptionsMap implements AppOptions interface
			emptyWasmOption,
			baseAppOptions...,
		)
	}
}

// NewRegenAppForSims creates a wrapper function that matches the expected signature for sims.NewSimulationAppInstance
func NewRegenAppForSims(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *regen.RegenApp {
	// Get the invCheckPeriod from appOpts, with a default value if not found
	var invCheckPeriod uint = 1
	invCheckPeriod = cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod))

	return regen.NewRegenApp(
		logger,
		db,
		nil, // traceStore (regen.NewRegenApp expects io.Writer, but we pass nil)
		loadLatest,
		invCheckPeriod,
		appOpts,
		emptyWasmOption,
		baseAppOptions...,
	)
}

func setupStateFactory(app *regen.RegenApp) sims.SimStateFactory {
	return sims.SimStateFactory{
		Codec:         app.AppCodec(),
		AppStateFn:    simtestutil.AppStateFn(app.AppCodec(), app.SimulationManager(), app.DefaultGenesis()),
		BlockedAddr:   app.BlockAddresses(),
		AccountSource: app.AccountKeeper,
		BalanceSource: app.BankKeeper,
	}
}

var (
	exportAllModules       = []string{}
	exportWithValidatorSet = []string{}
)

func TestAppImportExport(t *testing.T) {
	sims.Run(t, NewAppFactory(), setupStateFactory, func(tb testing.TB, ti sims.TestInstance[*regen.RegenApp], accs []simtypes.Account) {
		tb.Helper()
		app := ti.App
		tb.Log("exporting genesis...")
		exported, err := app.ExportAppStateAndValidators(false, exportWithValidatorSet, exportAllModules)
		require.NoError(tb, err)

		tb.Log("importing genesis...")
		newTestInstance := sims.NewSimulationAppInstance(tb, ti.Cfg, NewRegenAppForSims)
		newApp := newTestInstance.App
		var genesisState regen.GenesisState
		require.NoError(tb, json.Unmarshal(exported.AppState, &genesisState))
		ctxB := newApp.NewContextLegacy(true, cmtproto.Header{Height: app.LastBlockHeight()})
		_, err = newApp.ModuleManager.InitGenesis(ctxB, newApp.AppCodec(), genesisState)
		if IsEmptyValidatorSetErr(err) {
			tb.Skip("Skipping simulation as all validators have been unbonded")
			return
		}
		require.NoError(tb, err)
		err = newApp.StoreConsensusParams(ctxB, exported.ConsensusParams)
		require.NoError(tb, err)

		tb.Log("comparing stores...")
		// skip certain prefixes
		skipPrefixes := map[string][][]byte{
			stakingtypes.StoreKey: {
				stakingtypes.UnbondingQueueKey, stakingtypes.RedelegationQueueKey, stakingtypes.ValidatorQueueKey,
				stakingtypes.HistoricalInfoKey, stakingtypes.UnbondingIDKey, stakingtypes.UnbondingIndexKey,
				stakingtypes.UnbondingTypeKey,
				stakingtypes.ValidatorUpdatesKey,
			},
			authzkeeper.StoreKey:   {authzkeeper.GrantQueuePrefix},
			feegrant.StoreKey:      {feegrant.FeeAllowanceQueueKeyPrefix},
			slashingtypes.StoreKey: {slashingtypes.ValidatorMissedBlockBitmapKeyPrefix},
		}
		AssertEqualStores(tb, app, newApp, app.SimulationManager().StoreDecoders, skipPrefixes)
	})
}

func IsEmptyValidatorSetErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), "validator set is empty after InitGenesis")
}

func AssertEqualStores(
	tb testing.TB,
	app, newApp ComparableStoreApp,
	storeDecoders simtypes.StoreDecoderRegistry,
	skipPrefixes map[string][][]byte,
) {
	tb.Helper()
	ctxA := app.NewContextLegacy(true, cmtproto.Header{Height: app.LastBlockHeight()})
	ctxB := newApp.NewContextLegacy(true, cmtproto.Header{Height: app.LastBlockHeight()})

	storeKeys := app.GetStoreKeys()
	require.NotEmpty(tb, storeKeys)

	for _, appKeyA := range storeKeys {
		// only compare kvstores
		if _, ok := appKeyA.(*storetypes.KVStoreKey); !ok {
			continue
		}

		keyName := appKeyA.Name()
		appKeyB := newApp.GetKey(keyName)

		storeA := ctxA.KVStore(appKeyA)
		storeB := ctxB.KVStore(appKeyB)

		failedKVAs, failedKVBs := simtestutil.DiffKVStores(storeA, storeB, skipPrefixes[keyName])
		require.Equal(tb, len(failedKVAs), len(failedKVBs), "unequal sets of key-values to compare %s, key stores %s and %s", keyName, appKeyA, appKeyB)

		tb.Logf("compared %d different key/value pairs between %s and %s\n", len(failedKVAs), appKeyA, appKeyB)
		if !assert.Equal(tb, 0, len(failedKVAs), simtestutil.GetSimulationLog(keyName, storeDecoders, failedKVAs, failedKVBs)) {
			for _, v := range failedKVAs {
				tb.Logf("store mismatch: %q\n", v)
			}
			tb.FailNow()
		}
	}
}

type ComparableStoreApp interface {
	LastBlockHeight() int64
	NewContextLegacy(isCheckTx bool, header cmtproto.Header) sdk.Context
	GetKey(storeKey string) *storetypes.KVStoreKey
	GetStoreKeys() []storetypes.StoreKey
}
