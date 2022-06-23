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

func TestDenomUnitsMigration(t *testing.T) {
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

	rNCT := "eco.uc.rNCT"
	nCT := "eco.uc.NCT"
	regenApp.BankKeeper.SetDenomMetaData(ctx, types.Metadata{
		Description: "this is a test rNCT basket",
		DenomUnits: []*types.DenomUnit{
			{
				Denom:    "eco.C.rNCT",
				Exponent: 6,
				Aliases:  []string{},
			},
			{
				Denom:    "eco.uC.rNCT",
				Exponent: 0,
				Aliases:  []string{},
			},
			{
				Denom:    "eco.nC.rNCT",
				Exponent: 9,
				Aliases:  []string{},
			},
			{
				Denom:    "eco.mC.rNCT",
				Exponent: 3,
				Aliases:  []string{},
			},
		},
		Base:    rNCT,
		Display: "eco.C.rNCT",
		Name:    "rNCT",
		Symbol:  "rNCT",
	})

	regenApp.BankKeeper.SetDenomMetaData(ctx, types.Metadata{
		Description: "This is a test NCT basket",
		DenomUnits: []*types.DenomUnit{
			{
				Denom:    "eco.C.NCT",
				Exponent: 6,
				Aliases:  []string{},
			},
			{
				Denom:    "eco.uC.NCT",
				Exponent: 0,
				Aliases:  []string{},
			},
		},
		Base:    nCT,
		Display: "eco.C.NCT",
		Name:    "NCT",
		Symbol:  "NCT",
	})

	migrateDenomUnits(ctx, regenApp.BankKeeper)

	denomMetadata, found := regenApp.BankKeeper.GetDenomMetaData(ctx, rNCT)
	require.True(t, found)
	require.Len(t, denomMetadata.DenomUnits, 4)
	require.Equal(t, denomMetadata.DenomUnits[0].Exponent, uint32(0))
	require.Equal(t, denomMetadata.DenomUnits[1].Exponent, uint32(3))
	require.Equal(t, denomMetadata.DenomUnits[2].Exponent, uint32(6))
	require.Equal(t, denomMetadata.DenomUnits[3].Exponent, uint32(9))

	denomMetadata, found = regenApp.BankKeeper.GetDenomMetaData(ctx, nCT)
	require.True(t, found)
	require.Len(t, denomMetadata.DenomUnits, 2)
	require.Equal(t, denomMetadata.DenomUnits[0].Exponent, uint32(0))
	require.Equal(t, denomMetadata.DenomUnits[1].Exponent, uint32(6))
}
