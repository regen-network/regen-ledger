package server_test

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	coreserver "github.com/regen-network/regen-ledger/x/ecocredit/server/core"
)

func setupStore(t *testing.T) (sdk.Context, api.StateStore) {
	// interfaceRegistry := types.NewInterfaceRegistry()
	// ecocredit.RegisterTypes(interfaceRegistry)
	// key := sdk.NewKVStoreKey(ecocredit.ModuleName)
	// db := dbm.NewMemDB()
	// cms := store.NewCommitMultiStore(db)
	// cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	// err := cms.LoadLatestVersion()
	// require.NoError(t, err)
	// ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())

	ecocreditKey := sdk.NewKVStoreKey("ecocredit")
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(ecocreditKey, sdk.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	sdkCtx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)
	ss, err := api.NewStateStore(ormdb)
	require.Nil(t, err)
	return sdkCtx, ss
}

func TestTradableSupplyInvariants(t *testing.T) {
	acc1 := sdk.AccAddress([]byte("account1"))
	acc2 := sdk.AccAddress([]byte("account2"))

	testCases := []struct {
		msg           string
		balances      []*core.BatchBalance
		supply        []*core.BatchSupply
		basketBalance map[uint64]math.Dec
		expBroken     bool
	}{
		{
			"valid test case",
			[]*core.BatchBalance{
				{
					Address:  acc1.Bytes(),
					BatchId:  1,
					Tradable: "100",
				},
				{
					Address:  acc2.Bytes(),
					BatchId:  1,
					Tradable: "210",
				},
				{
					Address: acc2.Bytes(),
					BatchId: 1,
					Retired: "210",
				},
			},
			[]*core.BatchSupply{
				{
					BatchId:        1,
					TradableAmount: "320",
					RetiredAmount:  "210",
				},
			},
			map[uint64]math.Dec{1: math.NewDecFromInt64(10)},
			false,
		},
		{
			"valid test case multiple denom",
			[]*core.BatchBalance{
				{
					Address:  acc1.Bytes(),
					BatchId:  1,
					Tradable: "100.123",
				},
				{
					Address:  acc2.Bytes(),
					BatchId:  1,
					Tradable: "210.456",
				},
				{
					Address:  acc2.Bytes(),
					BatchId:  2,
					Tradable: "210.456",
				},
			},
			[]*core.BatchSupply{
				{
					BatchId:        1,
					TradableAmount: "320.579",
					RetiredAmount:  "0",
				},
				{
					BatchId:        2,
					TradableAmount: "220.456",
					RetiredAmount:  "0",
				},
			},
			map[uint64]math.Dec{1: math.NewDecFromInt64(10), 2: math.NewDecFromInt64(10)},
			false,
		},
		// {
		// 	"fail with error tradable balance not found",
		// 	[]*ecocredit.Balance{
		// 		{
		// 			Address:         acc1.String(),
		// 			BatchDenom:      "1/2",
		// 			TradableBalance: "100.123",
		// 		},
		// 		{
		// 			Address:         acc2.String(),
		// 			BatchDenom:      "1/2",
		// 			TradableBalance: "210.456",
		// 		},
		// 	},
		// 	[]*ecocredit.Supply{
		// 		{
		// 			BatchDenom:     "1/2",
		// 			TradableSupply: "310.579",
		// 			RetiredSupply:  "0",
		// 		},
		// 		{
		// 			BatchDenom:     "3/4",
		// 			TradableSupply: "1234",
		// 			RetiredSupply:  "0",
		// 		},
		// 	},
		// 	map[string]math.Dec{},
		// 	true,
		// },
		// {
		// 	"fail with error supply does not match",
		// 	[]*ecocredit.Balance{
		// 		{
		// 			Address:         acc1.String(),
		// 			BatchDenom:      "1/2",
		// 			TradableBalance: "100.123",
		// 		},
		// 		{
		// 			Address:         acc2.String(),
		// 			BatchDenom:      "1/2",
		// 			TradableBalance: "210.456",
		// 		},
		// 		{
		// 			BatchDenom:      "3/4",
		// 			Address:         acc2.String(),
		// 			TradableBalance: "1234",
		// 		},
		// 	},
		// 	[]*ecocredit.Supply{
		// 		{
		// 			BatchDenom:     "1/2",
		// 			TradableSupply: "325.57",
		// 			RetiredSupply:  "0",
		// 		},
		// 		{
		// 			BatchDenom:     "3/4",
		// 			TradableSupply: "1234",
		// 			RetiredSupply:  "0",
		// 		},
		// 	},
		// 	map[string]math.Dec{},
		// 	true,
		// },
	}

	for _, tc := range testCases {
		tc := tc
		sdkCtx, ss := setupStore(t)
		ctx := sdk.WrapSDKContext(sdkCtx)
		t.Run(tc.msg, func(t *testing.T) {
			initBalances(t, ctx, ss, tc.balances)
			initSupply(t, ctx, ss, tc.supply)

			msg, broken := server.BatchSupplyInvariant(ctx, coreserver.NewKeeper(ss), tc.basketBalance)
			if tc.expBroken {
				require.True(t, broken, msg)
			} else {
				require.False(t, broken, msg)
			}
		})
	}
}

// func TestRetiredSupplyInvariants(t *testing.T) {
// 	acc1 := sdk.AccAddress([]byte("account1"))
// 	acc2 := sdk.AccAddress([]byte("account2"))

// 	testCases := []struct {
// 		msg       string
// 		balances  []*ecocredit.Balance
// 		supply    []*ecocredit.Supply
// 		expBroken bool
// 	}{
// 		{
// 			"valid test case",
// 			[]*ecocredit.Balance{
// 				{
// 					Address:        acc1.String(),
// 					BatchDenom:     "1/2",
// 					RetiredBalance: "100",
// 				},
// 				{
// 					Address:        acc2.String(),
// 					BatchDenom:     "1/2",
// 					RetiredBalance: "210",
// 				},
// 				{
// 					Address:         acc2.String(),
// 					BatchDenom:      "1/2",
// 					TradableBalance: "210",
// 				},
// 			},
// 			[]*ecocredit.Supply{
// 				{
// 					BatchDenom:     "1/2",
// 					RetiredSupply:  "310",
// 					TradableSupply: "210",
// 				},
// 			},
// 			false,
// 		},
// 		{
// 			"valid test case multiple denom",
// 			[]*ecocredit.Balance{
// 				{
// 					Address:        acc1.String(),
// 					BatchDenom:     "1/2",
// 					RetiredBalance: "100.123",
// 				},
// 				{
// 					Address:        acc2.String(),
// 					BatchDenom:     "1/2",
// 					RetiredBalance: "210.456",
// 				},
// 				{
// 					Address:        acc2.String(),
// 					BatchDenom:     "3/4",
// 					RetiredBalance: "210.456",
// 				},
// 			},
// 			[]*ecocredit.Supply{
// 				{
// 					BatchDenom:     "1/2",
// 					RetiredSupply:  "310.579",
// 					TradableSupply: "0",
// 				},
// 				{
// 					BatchDenom:     "3/4",
// 					RetiredSupply:  "210.456",
// 					TradableSupply: "0",
// 				},
// 			},
// 			false,
// 		},
// 		{
// 			"fail with error retired balance not found",
// 			[]*ecocredit.Balance{
// 				{
// 					Address:        acc1.String(),
// 					BatchDenom:     "1/2",
// 					RetiredBalance: "100.123",
// 				},
// 				{
// 					Address:        acc2.String(),
// 					BatchDenom:     "1/2",
// 					RetiredBalance: "210.456",
// 				},
// 			},
// 			[]*ecocredit.Supply{
// 				{
// 					BatchDenom:     "1/2",
// 					RetiredSupply:  "310.579",
// 					TradableSupply: "0",
// 				},
// 				{
// 					BatchDenom:     "3/4",
// 					RetiredSupply:  "1234",
// 					TradableSupply: "0",
// 				},
// 			},
// 			true,
// 		},
// 		{
// 			"fail with error retired supply does not match",
// 			[]*ecocredit.Balance{
// 				{
// 					Address:        acc1.String(),
// 					BatchDenom:     "1/2",
// 					RetiredBalance: "100.123",
// 				},
// 				{
// 					Address:        acc2.String(),
// 					BatchDenom:     "1/2",
// 					RetiredBalance: "210.456",
// 				},
// 				{
// 					BatchDenom:     "3/4",
// 					Address:        acc2.String(),
// 					RetiredBalance: "1234",
// 				},
// 			},
// 			[]*ecocredit.Supply{
// 				{
// 					BatchDenom:     "1/2",
// 					RetiredSupply:  "310.57",
// 					TradableSupply: "0",
// 				},
// 				{
// 					BatchDenom:     "3/4",
// 					RetiredSupply:  "1234",
// 					TradableSupply: "0",
// 				},
// 			},
// 			true,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		tc := tc
// 		ctx, storeKey := setupStore(t)
// 		store := ctx.KVStore(storeKey)
// 		t.Run(tc.msg, func(t *testing.T) {
// 			initBalances(t, store, tc.balances)
// 			initSupply(t, store, tc.supply)

// 			msg, broken := retiredSupplyInvariant(store)
// 			if tc.expBroken {
// 				require.True(t, broken, msg)
// 			} else {
// 				require.False(t, broken, msg)
// 			}
// 		})
// 	}
// }

func initBalances(t *testing.T, ctx context.Context, ss api.StateStore, balances []*core.BatchBalance) {
	for _, b := range balances {
		addr := sdk.AccAddress(b.Address)
		if b.Tradable != "" {
			_, err := math.NewNonNegativeDecFromString(b.Tradable)
			require.NoError(t, err)
			require.NoError(t, ss.BatchBalanceTable().Save(ctx, &api.BatchBalance{
				Address:  addr,
				BatchId:  b.BatchId,
				Tradable: b.Tradable,
			}))
		}
		if b.Retired != "" {
			_, err := math.NewNonNegativeDecFromString(b.Retired)
			require.NoError(t, err)
			require.NoError(t, ss.BatchBalanceTable().Save(ctx, &api.BatchBalance{
				Address: addr,
				BatchId: b.BatchId,
				Retired: b.Retired,
			}))
		}
	}
}

func initSupply(t *testing.T, ctx context.Context, ss api.StateStore, supply []*core.BatchSupply) {
	for _, s := range supply {
		ss.BatchSupplyTable().Save(ctx, &api.BatchSupply{
			BatchId:         s.BatchId,
			TradableAmount:  s.TradableAmount,
			RetiredAmount:   s.RetiredAmount,
			CancelledAmount: s.CancelledAmount,
		})
	}
}
