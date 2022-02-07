package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func setupStore(t *testing.T) ormdb.ModuleDB {
	interfaceRegistry := types.NewInterfaceRegistry()
	ecocredit.RegisterTypes(interfaceRegistry)
	db := dbm.NewMemDB()

	var moduleSchema = ormdb.ModuleSchema{
		FileDescriptors: map[uint32]protoreflect.FileDescriptor{
			1: ecocreditv1beta1.File_regen_ecocredit_v1beta1_state_proto,
		},
	}
	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      nil,
		Hooks:           nil,
	})
	modDB, err := ormdb.NewModuleDB(moduleSchema, ormdb.ModuleDBOptions{
		TypeResolver:  nil,
		FileResolver:  nil,
		JSONValidator: nil,
		GetBackend: func(ctx context.Context) (ormtable.Backend, error) {
			return backend, nil
		},
		GetReadBackend: func(ctx context.Context) (ormtable.ReadBackend, error) {
			return backend, nil
		},
	})
	require.NoError(t, err)
	return modDB
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
					TradableSupply: "310",
					RetiredSupply:  "210",
				},
			},
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
					TradableSupply: "310.579",
					RetiredSupply:  "0",
				},
				{
					BatchDenom:     "3/4",
					TradableSupply: "210.456",
					RetiredSupply:  "0",
				},
			},
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
					TradableSupply: "310.57",
					RetiredSupply:  "0",
				},
				{
					BatchDenom:     "3/4",
					TradableSupply: "1234",
					RetiredSupply:  "0",
				},
			},
			true,
		},
	}

	db := setupStore(t)

	ctx := sdk.NewContext(nil, tmproto.Header{}, false, log.NewNopLogger())
	supplyStore, err := ecocreditv1beta1.NewBatchSupplyStore(db)
	require.NoError(t, err)
	balanceStore, err := ecocreditv1beta1.NewBatchBalanceStore(db)
	require.NoError(t, err)
	batchInfoStore, err := ecocreditv1beta1.NewBatchInfoStore(db)
	require.NoError(t, err)

	initBatchDenom(t, ctx.Context(), batchInfoStore, testCases)

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.msg, func(t *testing.T) {

			initBalances(t, ctx, balanceStore, batchInfoStore, tc.balances)
			initSupply(t, ctx, supplyStore, batchInfoStore, tc.supply)

			msg, broken := tradableSupplyInvariant(supplyStore, balanceStore, batchInfoStore, ctx)
			if tc.expBroken {
				require.True(t, broken, msg)
			} else {
				require.False(t, broken, msg)
			}
		})
	}
}

func initBatchDenom(t *testing.T, ctx context.Context, store ecocreditv1beta1.BatchInfoStore, testCases []struct {
	msg       string
	balances  []*ecocredit.Balance
	supply    []*ecocredit.Supply
	expBroken bool
}) {
	for _, tc := range testCases {
		for _, b := range tc.balances {
			has, err := store.HasByBatchDenom(ctx, b.BatchDenom)
			require.NoError(t, err)
			if !has {
				err := store.Insert(ctx, &ecocreditv1beta1.BatchInfo{
					ProjectId:  0,
					BatchDenom: b.BatchDenom,
					Metadata:   nil,
					StartDate:  timestamppb.Now(),
					EndDate:    timestamppb.Now(),
				})
				require.NoError(t, err)
			}
		}
		for _, s := range tc.supply {
			has, err := store.HasByBatchDenom(ctx, s.BatchDenom)
			require.NoError(t, err)
			if !has {
				err := store.Insert(ctx, &ecocreditv1beta1.BatchInfo{
					ProjectId:  0,
					BatchDenom: s.BatchDenom,
					Metadata:   nil,
					StartDate:  timestamppb.Now(),
					EndDate:    timestamppb.Now(),
				})
				require.NoError(t, err)
			}
		}
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
	ctx := sdk.NewContext(nil, tmproto.Header{}, false, log.NewNopLogger())
	db := setupStore(t)
	supplyStore, err := ecocreditv1beta1.NewBatchSupplyStore(db)
	require.NoError(t, err)
	balanceStore, err := ecocreditv1beta1.NewBatchBalanceStore(db)
	require.NoError(t, err)
	batchInfoStore, err := ecocreditv1beta1.NewBatchInfoStore(db)
	require.NoError(t, err)
	initBatchDenom(t, ctx.Context(), batchInfoStore, testCases)
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			initBalances(t, ctx, balanceStore, batchInfoStore, tc.balances)
			initSupply(t, ctx, supplyStore, batchInfoStore, tc.supply)
			msg, broken := retiredSupplyInvariant(supplyStore, balanceStore, batchInfoStore, ctx)
			if tc.expBroken {
				require.True(t, broken, msg)
			} else {
				require.False(t, broken, msg)
			}
		})
	}
}

func initBalances(t *testing.T, ctx sdk.Context, store ecocreditv1beta1.BatchBalanceStore, batchInfoStore ecocreditv1beta1.BatchInfoStore, balances []*ecocredit.Balance) {
	for _, b := range balances {
		addr, err := sdk.AccAddressFromBech32(b.Address)
		require.NoError(t, err)
		info, err := batchInfoStore.GetByBatchDenom(ctx.Context(), b.BatchDenom)
		require.NoError(t, err)
		id := info.Id
		if b.TradableBalance == "" {
			b.TradableBalance = "0"
		}
		if b.RetiredBalance == "" {
			b.RetiredBalance = "0"
		}

		trad, err := math.NewNonNegativeDecFromString(b.TradableBalance)
		require.NoError(t, err)
		ret, err := math.NewNonNegativeDecFromString(b.RetiredBalance)
		require.NoError(t, err)

		bal, err := store.Get(ctx.Context(), addr, id)
		if err == nil && bal != nil {
			balTrad, err := math.NewNonNegativeDecFromString(bal.Tradable)
			require.NoError(t, err)
			balRet, err := math.NewNonNegativeDecFromString(bal.Retired)
			require.NoError(t, err)
			trad, err = balTrad.Add(trad)
			require.NoError(t, err)
			ret, err = balRet.Add(ret)
			require.NoError(t, err)
			err = store.Update(ctx.Context(), &ecocreditv1beta1.BatchBalance{
				Address:  addr,
				BatchId:  id,
				Tradable: trad.String(),
				Retired:  ret.String(),
			})
			require.NoError(t, err)
		} else {
			err = store.Insert(ctx.Context(), &ecocreditv1beta1.BatchBalance{
				Address:  addr,
				BatchId:  id,
				Tradable: trad.String(),
				Retired:  ret.String(),
			})
			require.NoError(t, err)
		}
	}
}

func initSupply(t *testing.T, ctx sdk.Context, store ecocreditv1beta1.BatchSupplyStore, batchStore ecocreditv1beta1.BatchInfoStore, supply []*ecocredit.Supply) {
	for _, s := range supply {
		info, err := batchStore.GetByBatchDenom(ctx.Context(), s.BatchDenom)
		require.NoError(t, err)
		id := info.Id
		trad, err := math.NewNonNegativeDecFromString(s.TradableSupply)
		require.NoError(t, err)
		ret, err := math.NewNonNegativeDecFromString(s.RetiredSupply)
		require.NoError(t, err)
		sup, err := store.Get(ctx.Context(), id)
		if err == nil && sup != nil {
			supplyTradable, err := math.NewNonNegativeDecFromString(sup.TradableAmount)
			require.NoError(t, err)
			supplyRetired, err := math.NewNonNegativeDecFromString(sup.RetiredAmount)
			require.NoError(t, err)
			trad, err = supplyTradable.Add(trad)
			require.NoError(t, err)
			ret, err = supplyRetired.Add(ret)
			require.NoError(t, err)
			err = store.Update(ctx.Context(), &ecocreditv1beta1.BatchSupply{
				BatchId:         sup.BatchId,
				TradableAmount:  trad.String(),
				RetiredAmount:   ret.String(),
				CancelledAmount: "",
			})
			require.NoError(t, err)
		} else {
			err = store.Insert(ctx.Context(), &ecocreditv1beta1.BatchSupply{
				BatchId:         id,
				TradableAmount:  trad.String(),
				RetiredAmount:   ret.String(),
				CancelledAmount: "",
			})
			require.NoError(t, err)
		}
	}
}
