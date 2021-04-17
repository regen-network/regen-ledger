package app_test

import (
	"fmt"
	"time"
	"testing"



    "github.com/stretchr/testify/suite"
	"github.com/regen-network/regen-ledger/types/testutil/network" //could not import error
	//data "github.com/regen-network/regen-ledger/x/data/client/testsuite"
	group "github.com/regen-network/regen-ledger/x/data/group/client/testsuite"
	testutil "github.com/regen-network/regen-ledger/types/testutil/network_test" //could not import error
	"github.com/regen-network/regen-ledger/app"	
	"github.com/cosmos/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/tm-db"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	tmrand "github.com/tendermint/tendermint/libs/rand"

)
// NewSimApp is not used in repository, why we need it?
func NewSimApp(val network.Validator) servertypes.Application {
	return app.NewRegenApp(
		val.Ctx.Logger, dbm.NewMemDB(), nil, true, make(map[int64]bool), val.Ctx.Config.RootDir, 0,
		app.MakeEncodingConfig(),
		simapp.EmptyAppOptions{},
		baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
		baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
	)
}

// DefaultConfig returns a sane default configuration suitable for nearly all
// testing requirements.
func DefaultConfig() network.Config {
	encCfg := app.MakeEncodingConfig()

	return network.Config{
		Codec:             encCfg.Marshaler,
		TxConfig:          encCfg.TxConfig,
		LegacyAmino:       encCfg.Amino,
		InterfaceRegistry: encCfg.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor:    NewSimApp,// do we need arguments to NewSimApp?
		GenesisState:      app.ModuleBasics.DefaultGenesis(encCfg.Marshaler),
		TimeoutCommit:     2 * time.Second,
		ChainID:           "chain-" + tmrand.NewRand().Str(6),
		NumValidators:     4,
		BondDenom:         sdk.DefaultBondDenom,
		MinGasPrices:      fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
		AccountTokens:     sdk.TokensFromConsensusPower(1000),
		StakingTokens:     sdk.TokensFromConsensusPower(500),
		BondedTokens:      sdk.TokensFromConsensusPower(100),
		PruningStrategy:   storetypes.PruningOptionNothing,
		CleanupDir:        true,
		SigningAlgo:       string(hd.Secp256k1Type),
		KeyringOptions:    []keyring.Option{},
	}
}

func TestModules(t *testing.T) {
	t.Parallel()
	cfg := DefaultConfig() // (requires app, or the NewRegenApp(), this will not sbe in types/testutil)

	//suite.Run(t, data.NewIntegrationTestSuite(cfg))
	suite.Run(t, group.NewIntegrationTestSuite(cfg))
	suite.Run(t, testutil.NewIntegrationTestSuite(cfg))

}
