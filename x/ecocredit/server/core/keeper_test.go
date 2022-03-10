package core

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	mocks2 "github.com/regen-network/regen-ledger/x/ecocredit/server/core/mocks"
)

type baseSuite struct {
	t            *testing.T
	db           ormdb.ModuleDB
	stateStore   api.StateStore
	ctx          context.Context
	k            Keeper
	ctrl         *gomock.Controller
	addr         sdk.AccAddress
	bankKeeper   *mocks.MockBankKeeper
	paramsKeeper *mocks2.MockParamKeeper
	storeKey     *sdk.KVStoreKey
	sdkCtx       sdk.Context
}

func setupBase(t *testing.T) *baseSuite {
	// prepare database
	s := &baseSuite{t: t}
	var err error
	s.db, err = ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	s.stateStore, err = api.NewStateStore(s.db)
	assert.NilError(t, err)

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	s.storeKey = sdk.NewKVStoreKey("test")
	cms.MountStoreWithDB(s.storeKey, sdk.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// setup test keeper
	s.ctrl = gomock.NewController(t)
	assert.NilError(t, err)
	s.bankKeeper = mocks.NewMockBankKeeper(s.ctrl)
	s.paramsKeeper = mocks2.NewMockParamKeeper(s.ctrl)
	s.k = NewKeeper(s.stateStore, s.bankKeeper, s.paramsKeeper)
	_, _, s.addr = testdata.KeyTestPubAddr()

	return s
}

// setupClassProjectBatch setups a class "C01", a project "PRO", a batch "C01-20200101-20210101-01", and a
// supply/balance of "10.5" for both retired and tradable.
func (s baseSuite) setupClassProjectBatch(t *testing.T) (className, projectName, batchDenom string) {
	className, projectName, batchDenom = "C01", "PRO", "C01-20200101-20210101-01"
	assert.NilError(t, s.stateStore.ClassInfoStore().Insert(s.ctx, &api.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   "",
		CreditType: "C",
	}))
	assert.NilError(t, s.stateStore.ProjectInfoStore().Insert(s.ctx, &api.ProjectInfo{
		Name:            "PRO",
		ClassId:         1,
		ProjectLocation: "US-OR",
		Metadata:        "",
	}))
	assert.NilError(t, s.stateStore.BatchInfoStore().Insert(s.ctx, &api.BatchInfo{
		ProjectId:  1,
		BatchDenom: "C01-20200101-20210101-01",
		Metadata:   "",
		StartDate:  &timestamppb.Timestamp{Seconds: 2},
		EndDate:    &timestamppb.Timestamp{Seconds: 2},
	}))
	assert.NilError(t, s.stateStore.BatchSupplyStore().Insert(s.ctx, &api.BatchSupply{
		BatchId:         1,
		TradableAmount:  "10.5",
		RetiredAmount:   "10.5",
		CancelledAmount: "",
	}))
	assert.NilError(t, s.stateStore.BatchBalanceStore().Insert(s.ctx, &api.BatchBalance{
		Address:  s.addr,
		BatchId:  1,
		Tradable: "10.5",
		Retired:  "10.5",
	}))
	return
}

// this is an example of how we will unit test the basket functionality with mocks
func TestKeeperExample(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	require.NotNil(t, s.k)
}
