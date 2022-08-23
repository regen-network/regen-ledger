package core

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"

	dbm "github.com/tendermint/tm-db"

	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
)

const (
	testClassID    = "C01"
	testProjectID  = "C01-001"
	testBatchDenom = "C01-001-20200101-20210101-001"
)

type baseSuite struct {
	t          gocuke.TestingT
	db         ormdb.ModuleDB
	stateStore api.StateStore
	ctx        context.Context
	k          Keeper
	ctrl       *gomock.Controller
	addr       sdk.AccAddress
	addr2      sdk.AccAddress
	bankKeeper *mocks.MockBankKeeper
	storeKey   *storetypes.KVStoreKey
	sdkCtx     sdk.Context
	authority  sdk.AccAddress
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

	_, _, moduleAddress := testdata.KeyTestPubAddr()
	s.authority, err = sdk.AccAddressFromBech32("regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68")
	require.NoError(t, err)

	basketStore, err := basketapi.NewStateStore(s.db)
	assert.NilError(t, err)

	s.k = NewKeeper(s.stateStore, s.bankKeeper, moduleAddress, basketStore, s.authority)
	_, _, s.addr = testdata.KeyTestPubAddr()
	_, _, s.addr2 = testdata.KeyTestPubAddr()

	return s
}

// this is an example of how we will unit test the basket functionality with mocks
func TestKeeperExample(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	require.NotNil(t, s.k)
}
