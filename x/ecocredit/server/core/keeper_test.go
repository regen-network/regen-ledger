package core

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
)

type baseSuite struct {
	t            gocuke.TestingT
	db           ormdb.ModuleDB
	stateStore   api.StateStore
	ctx          context.Context
	k            Keeper
	ctrl         *gomock.Controller
	addr         sdk.AccAddress
	addr2        sdk.AccAddress
	bankKeeper   *mocks.MockBankKeeper
	paramsKeeper *mocks.MockParamKeeper
	storeKey     *storetypes.KVStoreKey
	sdkCtx       sdk.Context
}

func setupBase(t gocuke.TestingT) *baseSuite {
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
	s.k = NewKeeper(s.stateStore, s.bankKeeper, s.paramsKeeper, moduleAddress)
	_, _, s.addr = testdata.KeyTestPubAddr()
	_, _, s.addr2 = testdata.KeyTestPubAddr()

	return s
}

// setupClassProjectBatch setups a class "C01", a project "C01-001", a batch "C01-001-20200101-20210101-01", and a
// supply/balance of "10.5" for both retired and tradable.
func (s baseSuite) setupClassProjectBatch(t gocuke.TestingT) (classId, projectId, batchDenom string) {
	var err error
	classId, projectId = "C01", "C01-001"
	start, end := &timestamppb.Timestamp{Seconds: 2}, &timestamppb.Timestamp{Seconds: 30}
	startTime, endTime := start.AsTime(), end.AsTime()
	batchDenom, err = core.FormatBatchDenom(projectId, 1, &startTime, &endTime)
	assert.NilError(t, err)

	cKey, err := s.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               classId,
		Admin:            s.addr,
		Metadata:         "",
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err)

	assert.NilError(t, s.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: cKey,
		Issuer:   s.addr,
	}))

	pKey, err := s.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Admin:        s.addr,
		Id:           projectId,
		ClassKey:     cKey,
		Jurisdiction: "US-OR",
		Metadata:     "",
	})
	assert.NilError(t, err)
	assert.NilError(t, s.stateStore.ProjectSequenceTable().Insert(s.ctx, &api.ProjectSequence{
		ClassKey:     cKey,
		NextSequence: 2,
	}))

	bKey, err := s.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		ProjectKey: pKey,
		Issuer:     s.addr,
		Denom:      batchDenom,
		Metadata:   "",
		StartDate:  start,
		EndDate:    end,
	})
	assert.NilError(t, err)
	assert.NilError(t, s.stateStore.BatchSequenceTable().Insert(s.ctx, &api.BatchSequence{
		ProjectKey:   pKey,
		NextSequence: 2,
	}))

	assert.NilError(t, s.stateStore.BatchSupplyTable().Insert(s.ctx, &api.BatchSupply{
		BatchKey:        bKey,
		TradableAmount:  "10.5",
		RetiredAmount:   "10.5",
		CancelledAmount: "",
	}))
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, &api.BatchBalance{
		BatchKey:       bKey,
		Address:        s.addr,
		TradableAmount: "10.5",
		RetiredAmount:  "10.5",
	}))
	return
}

// this is an example of how we will unit test the basket functionality with mocks
func TestKeeperExample(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	require.NotNil(t, s.k)
}
