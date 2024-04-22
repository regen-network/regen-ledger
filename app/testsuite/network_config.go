package testsuite

import (
	"fmt"
	"os"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	pruningtypes "github.com/cosmos/cosmos-sdk/store/pruning/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"

	"github.com/regen-network/regen-ledger/v5/app"
)

// NewTestNetworkFixture returns a new simapp AppConstructor for network simulation tests
func NewTestNetworkFixture() network.TestFixture {
	dir, err := os.MkdirTemp("", "regen")
	if err != nil {
		panic(fmt.Sprintf("failed creating temporary directory: %v", err))
	}
	defer os.RemoveAll(dir)

	a := app.NewRegenApp(log.NewNopLogger(), dbm.NewMemDB(), nil, true, map[int64]bool{}, 0,
		simtestutil.NewAppOptionsWithFlagHome(dir))

	appCtr := func(val network.ValidatorI) servertypes.Application {
		cfg := val.GetAppConfig()
		skipUpgrades := map[int64]bool{}

		return app.NewRegenApp(
			val.GetCtx().Logger, dbm.NewMemDB(), nil, true, skipUpgrades, 0,
			simtestutil.NewAppOptionsWithFlagHome(val.GetCtx().Config.RootDir),
			baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(cfg.Pruning)),
			baseapp.SetMinGasPrices(cfg.MinGasPrices),
			baseapp.SetChainID(val.GetCtx().Viper.GetString(flags.FlagChainID)),
		)
	}

	return network.TestFixture{
		AppConstructor: appCtr,
		GenesisState:   a.DefaultGenesis(),
		EncodingConfig: testutil.TestEncodingConfig{
			InterfaceRegistry: a.InterfaceRegistry(),
			Codec:             a.AppCodec(),
			TxConfig:          a.TxConfig(),
			Amino:             a.LegacyAmino(),
		},
	}
}
