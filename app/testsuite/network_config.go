package testsuite

import (
	"fmt"
	"time"

	dbm "github.com/cometbft/cometbft-db"

	tmrand "github.com/cometbft/cometbft/libs/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	pruningtypes "github.com/cosmos/cosmos-sdk/store/pruning/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/regen-network/regen-ledger/types/v2/testutil/network"
	"github.com/regen-network/regen-ledger/v5/app"
	"github.com/regen-network/regen-ledger/v5/app/params"
)

func NewRegenAppConstructor(encCfg params.EncodingConfig) network.AppConstructor {
	return func(val network.Validator) servertypes.Application {
		return app.NewRegenApp(
			val.Ctx.Logger, dbm.NewMemDB(), nil, true, make(map[int64]bool), val.Ctx.Config.RootDir, 0,
			encCfg,
			EmptyAppOptions{},
			baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
		)
	}
}

// DefaultConfig returns a sane default configuration suitable for nearly all
// testing requirements.
func DefaultConfig() network.Config {
	encCfg := app.MakeEncodingConfig()

	return network.Config{
		Codec:             encCfg.Codec,
		TxConfig:          encCfg.TxConfig,
		LegacyAmino:       encCfg.Amino,
		InterfaceRegistry: encCfg.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor:    NewRegenAppConstructor(encCfg),
		GenesisState:      app.ModuleBasics.DefaultGenesis(encCfg.Codec),
		TimeoutCommit:     2 * time.Second,
		ChainID:           "chain-" + tmrand.NewRand().Str(6),
		NumValidators:     2,
		BondDenom:         sdk.DefaultBondDenom,
		MinGasPrices:      fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
		AccountTokens:     sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
		StakingTokens:     sdk.TokensFromConsensusPower(1e7, sdk.DefaultPowerReduction),
		BondedTokens:      sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
		PruningStrategy:   pruningtypes.PruningOptionNothing,
		CleanupDir:        true,
		SigningAlgo:       string(hd.Secp256k1Type),
		KeyringOptions:    []keyring.Option{},
		PrintMnemonic:     false,
	}
}
