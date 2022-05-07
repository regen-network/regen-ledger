package basket_test

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
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
	ecocreditApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	mocks2 "github.com/regen-network/regen-ledger/x/ecocredit/mocks"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
)

type baseSuite struct {
	t            gocuke.TestingT
	db           ormdb.ModuleDB
	ctx          context.Context
	k            basket.Keeper
	ctrl         *gomock.Controller
	addr         sdk.AccAddress
	stateStore   api.StateStore
	coreStore    ecocreditApi.StateStore
	bankKeeper   *mocks2.MockBankKeeper
	distKeeper   *mocks2.MockDistributionKeeper
	paramsKeeper *mocks2.MockParamKeeper
	storeKey     *sdk.KVStoreKey
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
	s.coreStore, err = ecocreditApi.NewStateStore(s.db)
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
	s.bankKeeper = mocks2.NewMockBankKeeper(s.ctrl)
	s.distKeeper = mocks2.NewMockDistributionKeeper(s.ctrl)
	s.paramsKeeper = mocks2.NewMockParamKeeper(s.ctrl)

	_, _, moduleAddress := testdata.KeyTestPubAddr()
	s.k = basket.NewKeeper(s.stateStore, s.coreStore, s.bankKeeper, s.distKeeper, s.paramsKeeper, moduleAddress)
	s.coreStore, err = ecoApi.NewStateStore(s.db)
	assert.NilError(t, err)
	assert.NilError(t, s.coreStore.CreditTypeTable().Insert(s.ctx, &ecocreditApi.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "metric ton C02",
		Precision:    6,
	}))
	_, _, s.addr = testdata.KeyTestPubAddr()

	return s
}
