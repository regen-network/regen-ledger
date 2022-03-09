package server

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

func setupStore(t *testing.T) (sdk.Context, *sdk.KVStoreKey) {
	interfaceRegistry := types.NewInterfaceRegistry()
	ecocredit.RegisterTypes(interfaceRegistry)
	key := sdk.NewKVStoreKey(ecocredit.ModuleName)
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	err := cms.LoadLatestVersion()
	require.NoError(t, err)
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	return ctx, key
}

func TestTradableSupplyInvariants(t *testing.T) {
	acc1 := sdk.AccAddress([]byte("account1"))
	acc2 := sdk.AccAddress([]byte("account2"))

	testCases := []struct {
		msg           string
		balances      []*ecocredit.Balance
		supply        []*ecocredit.Supply
		basketBalance map[string]math.Dec
		expBroken     bool
	}{
		{
			"valid test case",
			[]*ecocredit.Balance{
				{
					Address:         acc1.String(),
					BatchDenom:      "1/2",
					TradableBalance: "100",
				},
				{
					Address:         acc2.String(),
					BatchDenom:      "1/2",
					TradableBalance: "210",
				},
				{
					Address:        acc2.String(),
					BatchDenom:     "1/2",
					RetiredBalance: "210",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom:     "1/2",
					TradableSupply: "320",
					RetiredSupply:  "210",
				},
			},
			map[string]math.Dec{"1/2": math.NewDecFromInt64(10)},
			false,
		},
		{
			"valid test case multiple denom",
			[]*ecocredit.Balance{
				{
					Address:         acc1.String(),
					BatchDenom:      "1/2",
					TradableBalance: "100.123",
				},
				{
					Address:         acc2.String(),
					BatchDenom:      "1/2",
					TradableBalance: "210.456",
				},
				{
					Address:         acc2.String(),
					BatchDenom:      "3/4",
					TradableBalance: "210.456",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom:     "1/2",
					TradableSupply: "320.579",
					RetiredSupply:  "0",
				},
				{
					BatchDenom:     "3/4",
					TradableSupply: "220.456",
					RetiredSupply:  "0",
				},
			},
			map[string]math.Dec{"1/2": math.NewDecFromInt64(10), "3/4": math.NewDecFromInt64(10)},
			false,
		},
		{
			"fail with error tradable balance not found",
			[]*ecocredit.Balance{
				{
					Address:         acc1.String(),
					BatchDenom:      "1/2",
					TradableBalance: "100.123",
				},
				{
					Address:         acc2.String(),
					BatchDenom:      "1/2",
					TradableBalance: "210.456",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom:     "1/2",
					TradableSupply: "310.579",
					RetiredSupply:  "0",
				},
				{
					BatchDenom:     "3/4",
					TradableSupply: "1234",
					RetiredSupply:  "0",
				},
			},
			map[string]math.Dec{},
			true,
		},
		{
			"fail with error supply does not match",
			[]*ecocredit.Balance{
				{
					Address:         acc1.String(),
					BatchDenom:      "1/2",
					TradableBalance: "100.123",
				},
				{
					Address:         acc2.String(),
					BatchDenom:      "1/2",
					TradableBalance: "210.456",
				},
				{
					BatchDenom:      "3/4",
					Address:         acc2.String(),
					TradableBalance: "1234",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom:     "1/2",
					TradableSupply: "325.57",
					RetiredSupply:  "0",
				},
				{
					BatchDenom:     "3/4",
					TradableSupply: "1234",
					RetiredSupply:  "0",
				},
			},
			map[string]math.Dec{},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		ctx, storeKey := setupStore(t)
		store := ctx.KVStore(storeKey)
		t.Run(tc.msg, func(t *testing.T) {
			initBalances(t, store, tc.balances)
			initSupply(t, store, tc.supply)

			msg, broken := tradableSupplyInvariant(store, tc.basketBalance)
			if tc.expBroken {
				require.True(t, broken, msg)
			} else {
				require.False(t, broken, msg)
			}
		})
	}
}

func TestRetiredSupplyInvariants(t *testing.T) {
	acc1 := sdk.AccAddress([]byte("account1"))
	acc2 := sdk.AccAddress([]byte("account2"))

	testCases := []struct {
		msg       string
		balances  []*ecocredit.Balance
		supply    []*ecocredit.Supply
		expBroken bool
	}{
		{
			"valid test case",
			[]*ecocredit.Balance{
				{
					Address:        acc1.String(),
					BatchDenom:     "1/2",
					RetiredBalance: "100",
				},
				{
					Address:        acc2.String(),
					BatchDenom:     "1/2",
					RetiredBalance: "210",
				},
				{
					Address:         acc2.String(),
					BatchDenom:      "1/2",
					TradableBalance: "210",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom:     "1/2",
					RetiredSupply:  "310",
					TradableSupply: "210",
				},
			},
			false,
		},
		{
			"valid test case multiple denom",
			[]*ecocredit.Balance{
				{
					Address:        acc1.String(),
					BatchDenom:     "1/2",
					RetiredBalance: "100.123",
				},
				{
					Address:        acc2.String(),
					BatchDenom:     "1/2",
					RetiredBalance: "210.456",
				},
				{
					Address:        acc2.String(),
					BatchDenom:     "3/4",
					RetiredBalance: "210.456",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom:     "1/2",
					RetiredSupply:  "310.579",
					TradableSupply: "0",
				},
				{
					BatchDenom:     "3/4",
					RetiredSupply:  "210.456",
					TradableSupply: "0",
				},
			},
			false,
		},
		{
			"fail with error retired balance not found",
			[]*ecocredit.Balance{
				{
					Address:        acc1.String(),
					BatchDenom:     "1/2",
					RetiredBalance: "100.123",
				},
				{
					Address:        acc2.String(),
					BatchDenom:     "1/2",
					RetiredBalance: "210.456",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom:     "1/2",
					RetiredSupply:  "310.579",
					TradableSupply: "0",
				},
				{
					BatchDenom:     "3/4",
					RetiredSupply:  "1234",
					TradableSupply: "0",
				},
			},
			true,
		},
		{
			"fail with error retired supply does not match",
			[]*ecocredit.Balance{
				{
					Address:        acc1.String(),
					BatchDenom:     "1/2",
					RetiredBalance: "100.123",
				},
				{
					Address:        acc2.String(),
					BatchDenom:     "1/2",
					RetiredBalance: "210.456",
				},
				{
					BatchDenom:     "3/4",
					Address:        acc2.String(),
					RetiredBalance: "1234",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom:     "1/2",
					RetiredSupply:  "310.57",
					TradableSupply: "0",
				},
				{
					BatchDenom:     "3/4",
					RetiredSupply:  "1234",
					TradableSupply: "0",
				},
			},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		ctx, storeKey := setupStore(t)
		store := ctx.KVStore(storeKey)
		t.Run(tc.msg, func(t *testing.T) {
			initBalances(t, store, tc.balances)
			initSupply(t, store, tc.supply)

			msg, broken := retiredSupplyInvariant(store)
			if tc.expBroken {
				require.True(t, broken, msg)
			} else {
				require.False(t, broken, msg)
			}
		})
	}
}

func initBalances(t *testing.T, store sdk.KVStore, balances []*ecocredit.Balance) {
	for _, b := range balances {
		denomT := ecocredit.BatchDenomT(b.BatchDenom)
		addr, err := sdk.AccAddressFromBech32(b.Address)
		require.NoError(t, err)
		if b.TradableBalance != "" {
			d, err := math.NewNonNegativeDecFromString(b.TradableBalance)
			require.NoError(t, err)
			key := ecocredit.TradableBalanceKey(addr, denomT)
			ecocredit.SetDecimal(store, key, d)
		}
		if b.RetiredBalance != "" {
			d, err := math.NewNonNegativeDecFromString(b.RetiredBalance)
			require.NoError(t, err)
			key := ecocredit.RetiredBalanceKey(addr, denomT)
			ecocredit.SetDecimal(store, key, d)
		}
	}
}

func initSupply(t *testing.T, store sdk.KVStore, supply []*ecocredit.Supply) {
	for _, s := range supply {
		denomT := ecocredit.BatchDenomT(s.BatchDenom)
		d, err := math.NewNonNegativeDecFromString(s.TradableSupply)
		require.NoError(t, err)
		key := ecocredit.TradableSupplyKey(denomT)
		ecocredit.AddAndSetDecimal(store, key, d)
		d, err = math.NewNonNegativeDecFromString(s.RetiredSupply)
		require.NoError(t, err)
		key = ecocredit.RetiredSupplyKey(denomT)
		ecocredit.AddAndSetDecimal(store, key, d)
	}
}
