//go:build !experimental
// +build !experimental

package app

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/CosmWasm/wasmd/app"
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
		app.DefaultNodeHome, 0, encCfg, simapp.EmptyAppOptions{}, nil)
	bz := NewDefaultGenesisState(encCfg.Marshaler)
	stateBytes, err := json.MarshalIndent(bz, "", " ")
	require.NoError(t, err)

	regenApp.InitChain(
		abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: app.DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)

	ctx := regenApp.BaseApp.NewContext(false, tmproto.Header{})

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
		Base:    "eco.uC.rNCT",
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
		Base:    "eco.uC.NCT",
		Display: "eco.C.NCT",
		Name:    "NCT",
		Symbol:  "NCT",
	})

	err = migrateDenomUnits(ctx, regenApp.BankKeeper)
	require.NoError(t, err)

	md, _ := regenApp.BankKeeper.GetDenomMetaData(ctx, "eco.uC.rNCT")
	fmt.Println(md.DenomUnits)
	require.True(t, false)

}
