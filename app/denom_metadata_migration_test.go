//go:build !experimental
// +build !experimental

package app

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

func TestMigrateDenomMetadata(t *testing.T) {
	encCfg := MakeEncodingConfig()
	db := dbm.NewMemDB()

	regenApp := NewRegenApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{},
		DefaultNodeHome, 0, encCfg, simapp.EmptyAppOptions{}, nil)
	bz := NewDefaultGenesisState(encCfg.Marshaler)
	stateBytes, err := json.MarshalIndent(bz, "", " ")
	require.NoError(t, err)

	regenApp.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)

	ctx := regenApp.BaseApp.NewContext(false, tmproto.Header{})
	regenApp.BankKeeper.SetDenomMetaData(ctx, types.Metadata{
		Description: "The native staking token of Regen Network",
		Base:        "uregen",
		Display:     "regen",
		DenomUnits: []*types.DenomUnit{
			{
				Denom:    "uregen",
				Exponent: 0,
				Aliases:  []string{"microregen"},
			},
			{
				Denom:    "mregen",
				Exponent: 3,
				Aliases:  []string{"milliregen"},
			},
			{
				Denom:    "regen",
				Exponent: 6,
			},
		},
	})

	regenApp.BankKeeper.SetDenomMetaData(ctx, types.Metadata{
		Description: "The native staking token of Cosmos Hub",
		Base:        "uatom",
		Display:     "atom",
		DenomUnits: []*types.DenomUnit{
			{
				Denom:    "uatom",
				Exponent: 0,
				Aliases:  []string{"microatom"},
			},
			{
				Denom:    "matom",
				Exponent: 3,
				Aliases:  []string{"milliatom"},
			},
			{
				Denom:    "atom",
				Exponent: 6,
			},
		},
		Name:   "Cosmos Atom",
		Symbol: "ATOM",
	})

	metadata, found := regenApp.BankKeeper.GetDenomMetaData(ctx, "uregen")
	require.True(t, found)
	require.Equal(t, metadata.Name, "")
	require.Equal(t, metadata.Symbol, "")

	err = migrateDenomMetadata(ctx, regenApp.BankKeeper)
	require.NoError(t, err)

	metadata, found = regenApp.BankKeeper.GetDenomMetaData(ctx, "uregen")
	require.True(t, found)
	require.Equal(t, metadata.Name, "Regen")
	require.Equal(t, metadata.Symbol, "REGEN")

	metadata, found = regenApp.BankKeeper.GetDenomMetaData(ctx, "uatom")
	require.True(t, found)
	require.Equal(t, metadata.Name, "Cosmos Atom")
	require.Equal(t, metadata.Symbol, "ATOM")
}
