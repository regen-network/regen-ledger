package basket_test

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
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	mocks2 "github.com/regen-network/regen-ledger/x/ecocredit/mocks"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
)

type baseSuite struct {
	t            *testing.T
	db           ormdb.ModuleDB
	stateStore   api.StateStore
	coreStore    ecoApi.StateStore
	ctx          context.Context
	k            basket.Keeper
	ctrl         *gomock.Controller
	addr         sdk.AccAddress
	bankKeeper   *mocks2.MockBankKeeper
	paramsKeeper *mocks2.MockParamKeeper
	sdkCtx       sdk.Context
	distKeeper   *mocks2.MockDistributionKeeper
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
	storeKey := sdk.NewKVStoreKey("test")
	cms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// setup test keeper
	s.ctrl = gomock.NewController(t)
	assert.NilError(t, err)
	s.bankKeeper = mocks2.NewMockBankKeeper(s.ctrl)
	s.distKeeper = mocks2.NewMockDistributionKeeper(s.ctrl)
	s.paramsKeeper = mocks2.NewMockParamKeeper(s.ctrl)
	s.k = basket.NewKeeper(s.db, s.bankKeeper, s.distKeeper, s.paramsKeeper)
	s.coreStore, err = ecoApi.NewStateStore(s.db)
	assert.NilError(t, err)
	_, _, s.addr = testdata.KeyTestPubAddr()

	return s
}

// this is an example of how we will unit test the basket functionality with mocks
func TestKeeperExample(t *testing.T) {
	s := setupBase(t)
	require.NotNil(t, s.k)
}
