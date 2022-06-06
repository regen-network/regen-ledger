//go:build !experimental
// +build !experimental

package app

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/CosmWasm/wasmd/app"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

func TestLostFunds(t *testing.T) {
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
}
