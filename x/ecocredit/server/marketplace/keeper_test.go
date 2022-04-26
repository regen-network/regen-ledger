package marketplace

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
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

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
)

type baseSuite struct {
	t            *testing.T
	db           ormdb.ModuleDB
	coreStore    ecoApi.StateStore
	marketStore  api.StateStore
	ctx          context.Context
	k            Keeper
	ctrl         *gomock.Controller
	addr         sdk.AccAddress
	bankKeeper   *mocks.MockBankKeeper
	paramsKeeper *mocks.MockParamKeeper
	storeKey     *sdk.KVStoreKey
	sdkCtx       sdk.Context
}

func setupBase(t *testing.T) *baseSuite {
	// prepare database
	s := &baseSuite{t: t}
	var err error
	s.db, err = ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	s.coreStore, err = ecoApi.NewStateStore(s.db)
	assert.NilError(t, err)
	s.marketStore, err = api.NewStateStore(s.db)
	assert.NilError(t, err)

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	s.storeKey = sdk.NewKVStoreKey("test")
	cms.MountStoreWithDB(s.storeKey, sdk.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	assert.NilError(t, s.coreStore.CreditTypeTable().Insert(s.ctx, &ecoApi.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "metric ton C02",
		Precision:    6,
	}))

	// setup test keeper
	s.ctrl = gomock.NewController(t)
	assert.NilError(t, err)
	s.bankKeeper = mocks.NewMockBankKeeper(s.ctrl)
	s.paramsKeeper = mocks.NewMockParamKeeper(s.ctrl)
	s.k = NewKeeper(s.marketStore, s.coreStore, s.bankKeeper, s.paramsKeeper)
	_, _, s.addr = testdata.KeyTestPubAddr()
	return s
}
