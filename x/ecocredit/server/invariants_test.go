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
		msg       string
		balances  []*ecocredit.Balance
		supply    []*ecocredit.Supply
		expBroken bool
	}{
		{
			"valid test case",
			[]*ecocredit.Balance{
				{
					Address:    acc1.String(),
					BatchDenom: "1/2",
					Balance:    "100",
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom: "1/2",
					Supply:     "310",
				},
			},
			false,
		},
		{
			"valid test case multiple denom",
			[]*ecocredit.Balance{
				{
					Address:    acc1.String(),
					BatchDenom: "1/2",
					Balance:    "100.123",
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210.456",
				},
				{
					Address:    acc2.String(),
					BatchDenom: "3/4",
					Balance:    "210.456",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom: "1/2",
					Supply:     "310.579",
				},
				{
					BatchDenom: "3/4",
					Supply:     "210.456",
				},
			},
			false,
		},
		{
			"fail with error tradable balance not found",
			[]*ecocredit.Balance{
				{
					Address:    acc1.String(),
					BatchDenom: "1/2",
					Balance:    "100.123",
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210.456",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom: "1/2",
					Supply:     "310.579",
				},
				{
					BatchDenom: "3/4",
					Supply:     "1234",
				},
			},
			true,
		},
		{
			"fail with error supply does not match",
			[]*ecocredit.Balance{
				{
					Address:    acc1.String(),
					BatchDenom: "1/2",
					Balance:    "100.123",
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210.456",
				},
				{
					BatchDenom: "3/4",
					Address:    acc2.String(),
					Balance:    "1234",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom: "1/2",
					Supply:     "310.57",
				},
				{
					BatchDenom: "3/4",
					Supply:     "1234",
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

			initBalances(t, store, tc.balances, func(addr sdk.AccAddress, denom batchDenomT) []byte {
				return TradableBalanceKey(addr, denom)
			})

			initSupply(t, store, tc.supply, func(denom batchDenomT) []byte {
				return TradableSupplyKey(denom)
			})

			msg, broken := tradableSupplyInvariant(store)
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
					Address:    acc1.String(),
					BatchDenom: "1/2",
					Balance:    "100",
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom: "1/2",
					Supply:     "310",
				},
			},
			false,
		},
		{
			"valid test case multiple denom",
			[]*ecocredit.Balance{
				{
					Address:    acc1.String(),
					BatchDenom: "1/2",
					Balance:    "100.123",
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210.456",
				},
				{
					Address:    acc2.String(),
					BatchDenom: "3/4",
					Balance:    "210.456",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom: "1/2",
					Supply:     "310.579",
				},
				{
					BatchDenom: "3/4",
					Supply:     "210.456",
				},
			},
			false,
		},
		{
			"fail with error retired balance not found",
			[]*ecocredit.Balance{
				{
					Address:    acc1.String(),
					BatchDenom: "1/2",
					Balance:    "100.123",
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210.456",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom: "1/2",
					Supply:     "310.579",
				},
				{
					BatchDenom: "3/4",
					Supply:     "1234",
				},
			},
			true,
		},
		{
			"fail with error retired supply does not match",
			[]*ecocredit.Balance{
				{
					Address:    acc1.String(),
					BatchDenom: "1/2",
					Balance:    "100.123",
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210.456",
				},
				{
					BatchDenom: "3/4",
					Address:    acc2.String(),
					Balance:    "1234",
				},
			},
			[]*ecocredit.Supply{
				{
					BatchDenom: "1/2",
					Supply:     "310.57",
				},
				{
					BatchDenom: "3/4",
					Supply:     "1234",
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
			initBalances(t, store, tc.balances, func(addr sdk.AccAddress, denom batchDenomT) []byte {
				return RetiredBalanceKey(addr, denom)
			})

			initSupply(t, store, tc.supply, func(denom batchDenomT) []byte {
				return RetiredSupplyKey(denom)
			})
			msg, broken := retiredSupplyInvariant(store)
			if tc.expBroken {
				require.True(t, broken, msg)
			} else {
				require.False(t, broken, msg)
			}
		})
	}
}

func initBalances(t *testing.T, store sdk.KVStore, balances []*ecocredit.Balance, keyFunc balanceKey) {
	for _, b := range balances {
		d, err := math.ParseNonNegativeDecimal(b.Balance)
		require.NoError(t, err)
		addr, err := sdk.AccAddressFromBech32(b.Address)
		require.NoError(t, err)
		setDecimal(store, keyFunc(addr, batchDenomT(b.BatchDenom)), d)
	}
}

func initSupply(t *testing.T, store sdk.KVStore, supply []*ecocredit.Supply, keyFunc supplyKey) {
	for _, s := range supply {
		d, err := math.ParseNonNegativeDecimal(s.Supply)
		require.NoError(t, err)
		setDecimal(store, keyFunc(batchDenomT(s.BatchDenom)), d)
	}
}
