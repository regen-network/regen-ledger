package server_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"

	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
	coreserver "github.com/regen-network/regen-ledger/x/ecocredit/server/core"
)

type baseSuite struct {
	t            *testing.T
	db           ormdb.ModuleDB
	stateStore   api.StateStore
	ctx          context.Context
	k            coreserver.Keeper
	ctrl         *gomock.Controller
	bankKeeper   *mocks.MockBankKeeper
	paramsKeeper *mocks.MockParamKeeper
	storeKey     *storetypes.KVStoreKey
	sdkCtx       sdk.Context
}

func setupBase(t *testing.T) *baseSuite {
	// prepare database
	s := &baseSuite{t: t}
	var err error
	s.db, err = ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	s.stateStore, err = api.NewStateStore(s.db)
	assert.NilError(t, err)

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	s.storeKey = sdk.NewKVStoreKey("test")
	cms.MountStoreWithDB(s.storeKey, storetypes.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// setup test keeper
	s.ctrl = gomock.NewController(t)
	assert.NilError(t, err)
	s.bankKeeper = mocks.NewMockBankKeeper(s.ctrl)
	s.paramsKeeper = mocks.NewMockParamKeeper(s.ctrl)
	_, _, moduleAddress := testdata.KeyTestPubAddr()
	_, _, authorityAddress := testdata.KeyTestPubAddr()
	s.k = coreserver.NewKeeper(s.stateStore, s.bankKeeper, s.paramsKeeper, moduleAddress, authorityAddress)

	return s
}

func TestBatchSupplyInvariant(t *testing.T) {
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
					Address:        acc1,
					BatchKey:       1,
					TradableAmount: "210",
					RetiredAmount:  "110",
				},
			},
			[]*core.BatchSupply{
				{
					BatchKey:       1,
					TradableAmount: "220",
					RetiredAmount:  "110",
				},
			},
			map[uint64]math.Dec{1: math.NewDecFromInt64(10)},
			false,
		},
		{
			"valid test case multiple denom",
			[]*core.BatchBalance{
				{
					Address:        acc1,
					BatchKey:       1,
					TradableAmount: "310.579",
					RetiredAmount:  "0",
				},
				{
					Address:        acc2,
					BatchKey:       2,
					TradableAmount: "210.456",
					RetiredAmount:  "100.1234",
				},
			},
			[]*core.BatchSupply{
				{
					BatchKey:       1,
					TradableAmount: "320.579",
					RetiredAmount:  "0",
				},
				{
					BatchKey:       2,
					TradableAmount: "220.456",
					RetiredAmount:  "100.1234",
				},
			},
			map[uint64]math.Dec{1: math.NewDecFromInt64(10), 2: math.NewDecFromInt64(10)},
			false,
		},
		{
			"fail with error tradable balance not found",
			[]*core.BatchBalance{
				{
					Address:        acc1,
					BatchKey:       1,
					TradableAmount: "100.123",
				},
				{
					Address:        acc2,
					BatchKey:       1,
					TradableAmount: "210.456",
				},
			},
			[]*core.BatchSupply{
				{
					BatchKey:       1,
					TradableAmount: "310.579",
					RetiredAmount:  "0",
				},
				{
					BatchKey:       3,
					TradableAmount: "1234",
					RetiredAmount:  "0",
				},
			},
			map[uint64]math.Dec{},
			true,
		},
		{
			"fail with error supply does not match",
			[]*core.BatchBalance{
				{
					Address:        acc1,
					BatchKey:       1,
					TradableAmount: "310.579",
				},
				{
					BatchKey:       2,
					Address:        acc2,
					TradableAmount: "1234",
				},
			},
			[]*core.BatchSupply{
				{
					BatchKey:       1,
					TradableAmount: "310.579",
					RetiredAmount:  "123",
				},
				{
					BatchKey:       2,
					TradableAmount: "12345",
					RetiredAmount:  "0",
				},
			},
			map[uint64]math.Dec{},
			true,
		},
		{
			"valid case escrowed balance",
			[]*core.BatchBalance{
				{
					Address:        acc1,
					BatchKey:       1,
					TradableAmount: "100",
					EscrowedAmount: "10",
					RetiredAmount:  "1",
				},
				{
					BatchKey:       2,
					Address:        acc2,
					TradableAmount: "1234",
					RetiredAmount:  "123",
					EscrowedAmount: "766",
				},
			},
			[]*core.BatchSupply{
				{
					BatchKey:       1,
					TradableAmount: "110",
					RetiredAmount:  "1",
				},
				{
					BatchKey:       2,
					TradableAmount: "2000",
					RetiredAmount:  "123",
				},
			},
			map[uint64]math.Dec{},
			false,
		},
		{
			"valid case multiple account",
			[]*core.BatchBalance{
				{
					Address:        acc1,
					BatchKey:       1,
					TradableAmount: "100",
					EscrowedAmount: "10",
					RetiredAmount:  "1",
				},
				{
					BatchKey:       1,
					Address:        acc2,
					TradableAmount: "1234",
					RetiredAmount:  "123",
					EscrowedAmount: "766",
				},
				{
					BatchKey:       2,
					Address:        acc2,
					TradableAmount: "1234",
					RetiredAmount:  "123",
					EscrowedAmount: "766",
				},
			},
			[]*core.BatchSupply{
				{
					BatchKey:       1,
					TradableAmount: "2110",
					RetiredAmount:  "124",
				},
				{
					BatchKey:       2,
					TradableAmount: "2000",
					RetiredAmount:  "123",
				},
			},
			map[uint64]math.Dec{},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite := setupBase(t)
		t.Run(tc.msg, func(t *testing.T) {
			initBalances(t, suite.ctx, suite.stateStore, tc.balances)
			initSupply(t, suite.ctx, suite.stateStore, tc.supply)

			msg, broken := coreserver.BatchSupplyInvariant(suite.ctx, suite.k, tc.basketBalance)
			if tc.expBroken {
				require.True(t, broken, msg)
			} else {
				require.False(t, broken, msg)
			}
		})
	}
}

func initBalances(t *testing.T, ctx context.Context, ss api.StateStore, balances []*core.BatchBalance) {
	for _, b := range balances {
		_, err := math.NewNonNegativeDecFromString(b.TradableAmount)
		require.NoError(t, err)

		require.NoError(t, ss.BatchBalanceTable().Insert(ctx, &api.BatchBalance{
			Address:        b.Address,
			BatchKey:       b.BatchKey,
			TradableAmount: b.TradableAmount,
			RetiredAmount:  b.RetiredAmount,
			EscrowedAmount: b.EscrowedAmount,
		}))
	}
}

func initSupply(t *testing.T, ctx context.Context, ss api.StateStore, supply []*core.BatchSupply) {
	for _, s := range supply {
		err := ss.BatchSupplyTable().Insert(ctx, &api.BatchSupply{
			BatchKey:        s.BatchKey,
			TradableAmount:  s.TradableAmount,
			RetiredAmount:   s.RetiredAmount,
			CancelledAmount: s.CancelledAmount,
		})
		require.NoError(t, err)
	}
}
