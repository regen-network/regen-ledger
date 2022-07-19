//go:build !experimental
// +build !experimental

package app

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
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

func TestMigrateDenomUnits(t *testing.T) {
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

func TestRecoverFunds(t *testing.T) {
	// address with funds inaccessible
	lostAddr, err := sdk.AccAddressFromBech32("regen1c3lpjaq0ytdtsrnjqzmtj3hceavl8fe2vtkj7f")
	require.NoError(t, err)

	// address that the community member has access to
	newAddr, err := sdk.AccAddressFromBech32("regen14tpuqrwf95evu3ejm9z7dn20ttcyzqy3jjpfv4")
	require.NoError(t, err)

	vestingPeriods := vestingtypes.Periods{
		{
			Length: 0,
			Amount: sdk.NewCoins(sdk.NewInt64Coin("uregen", 1000000)),
		},
		{
			Length: 28630800,
			Amount: sdk.NewCoins(sdk.NewInt64Coin("uregen", 406041682)),
		},
		{
			Length: 2629746,
			Amount: sdk.NewCoins(sdk.NewInt64Coin("uregen", 406041666)),
		},
		{
			Length: 2629746,
			Amount: sdk.NewCoins(sdk.NewInt64Coin("uregen", 406041666)),
		},
	}

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

	regenApp.AccountKeeper.SetAccount(ctx, authtypes.NewBaseAccount(newAddr, nil, 9068, 4))
	oldAccountBalance := sdk.NewCoins(sdk.NewCoin("uregen", sdk.NewInt(1219125014)))
	regenApp.AccountKeeper.SetAccount(ctx, vestingtypes.NewPeriodicVestingAccount(authtypes.NewBaseAccount(lostAddr, nil, 146, 0),
		oldAccountBalance, 1618498800, vestingPeriods))

	regenApp.BankKeeper.MintCoins(ctx, minttypes.ModuleName, oldAccountBalance)
	regenApp.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, lostAddr, oldAccountBalance)

	newAccountBalance := sdk.NewCoins(sdk.NewInt64Coin("uregen", 10747843))
	regenApp.BankKeeper.MintCoins(ctx, minttypes.ModuleName, newAccountBalance)
	regenApp.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, newAddr, newAccountBalance)

	ctx = ctx.WithBlockTime(time.Now())
	err = recoverFunds(ctx, regenApp.AccountKeeper, regenApp.BankKeeper)
	require.NoError(t, err)

	require.True(t, regenApp.BankKeeper.GetBalance(ctx, lostAddr, "uregen").IsZero())
	newbalance := regenApp.BankKeeper.GetAllBalances(ctx, newAddr)
	// verify new account pre-existing balance also included
	require.Equal(t, newbalance, newAccountBalance.Add(oldAccountBalance...))

	acc := regenApp.AccountKeeper.GetAccount(ctx, newAddr)
	pva, ok := acc.(*vestingtypes.PeriodicVestingAccount)
	require.True(t, ok)
	require.Equal(t, pva.GetVestingPeriods().String(), vestingPeriods.String())

	// verify old account is deleted
	acc = regenApp.AccountKeeper.GetAccount(ctx, lostAddr)
	require.Nil(t, acc)
}
