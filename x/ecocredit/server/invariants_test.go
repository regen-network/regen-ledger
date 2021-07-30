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
					Type:       ecocredit.Balance_TYPE_TRADABLE,
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210",
					Type:       ecocredit.Balance_TYPE_TRADABLE,
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210",
					Type:       ecocredit.Balance_TYPE_RETIRED,
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
					Type:       ecocredit.Balance_TYPE_TRADABLE,
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210.456",
					Type:       ecocredit.Balance_TYPE_TRADABLE,
				},
				{
					Address:    acc2.String(),
					BatchDenom: "3/4",
					Balance:    "210.456",
					Type:       ecocredit.Balance_TYPE_TRADABLE,
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
					Type:       ecocredit.Balance_TYPE_TRADABLE,
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210.456",
					Type:       ecocredit.Balance_TYPE_TRADABLE,
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
					Type:       ecocredit.Balance_TYPE_TRADABLE,
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210.456",
					Type:       ecocredit.Balance_TYPE_TRADABLE,
				},
				{
					BatchDenom: "3/4",
					Address:    acc2.String(),
					Balance:    "1234",
					Type:       ecocredit.Balance_TYPE_TRADABLE,
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

			initBalances(t, store, tc.balances)

			initSupply(t, store, tc.supply, ecocredit.Balance_TYPE_TRADABLE)

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
					Type:       ecocredit.Balance_TYPE_RETIRED,
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210",
					Type:       ecocredit.Balance_TYPE_RETIRED,
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210",
					Type:       ecocredit.Balance_TYPE_TRADABLE,
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
					Type:       ecocredit.Balance_TYPE_RETIRED,
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210.456",
					Type:       ecocredit.Balance_TYPE_RETIRED,
				},
				{
					Address:    acc2.String(),
					BatchDenom: "3/4",
					Balance:    "210.456",
					Type:       ecocredit.Balance_TYPE_RETIRED,
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
					Type:       ecocredit.Balance_TYPE_RETIRED,
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210.456",
					Type:       ecocredit.Balance_TYPE_RETIRED,
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
					Type:       ecocredit.Balance_TYPE_RETIRED,
				},
				{
					Address:    acc2.String(),
					BatchDenom: "1/2",
					Balance:    "210.456",
					Type:       ecocredit.Balance_TYPE_RETIRED,
				},
				{
					BatchDenom: "3/4",
					Address:    acc2.String(),
					Balance:    "1234",
					Type:       ecocredit.Balance_TYPE_RETIRED,
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
			initBalances(t, store, tc.balances)

			initSupply(t, store, tc.supply, ecocredit.Balance_TYPE_RETIRED)
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
		d, err := math.ParseNonNegativeDecimal(b.Balance)
		require.NoError(t, err)
		addr, err := sdk.AccAddressFromBech32(b.Address)
		require.NoError(t, err)
		key, err := getBalanceKey(b.Type, addr, batchDenomT(b.BatchDenom))
		require.NoError(t, err)
		setDecimal(store, key, d)
	}
}

func initSupply(t *testing.T, store sdk.KVStore, supply []*ecocredit.Supply, balanceType ecocredit.Balance_Type) {
	for _, s := range supply {
		d, err := math.ParseNonNegativeDecimal(s.Supply)
		require.NoError(t, err)
		key, err := getSupplyKey(balanceType, batchDenomT(s.BatchDenom))
		require.NoError(t, err)
		setDecimal(store, key, d)
	}
}
