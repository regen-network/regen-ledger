package keeper

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	baseapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

const (
	testClassID    = "C01"
	testBatchDenom = "C01-001-20200101-20210101-001"
)

type baseSuite struct {
	t           gocuke.TestingT
	db          ormdb.ModuleDB
	baseStore   baseapi.StateStore
	marketStore api.StateStore
	ctx         context.Context
	k           Keeper
	ctrl        *gomock.Controller
	addrs       []sdk.AccAddress
	bankKeeper  *mocks.MockBankKeeper
	storeKey    *storetypes.KVStoreKey
	sdkCtx      sdk.Context
}

func setupBase(t gocuke.TestingT, numAddresses int) *baseSuite {
	// prepare database
	s := &baseSuite{t: t}
	var err error
	s.db, err = ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	s.baseStore, err = baseapi.NewStateStore(s.db)
	assert.NilError(t, err)
	s.marketStore, err = api.NewStateStore(s.db)
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

	authority, err := sdk.AccAddressFromBech32("regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68")
	assert.NilError(s.t, err)
	s.k = NewKeeper(s.marketStore, s.baseStore, s.bankKeeper, authority)

	// set test accounts
	for i := 0; i < numAddresses; i++ {
		var _, _, addr = testdata.KeyTestPubAddr()
		s.addrs = append(s.addrs, addr)
	}

	return s
}

// assertCreditsEscrowed adds orderAmt to tradable, subtracts from escrowed in before balance/supply and checks that it is equal to after balance/supply.
func assertCreditsEscrowed(t gocuke.TestingT, balanceBefore, balanceAfter *baseapi.BatchBalance, orderAmt math.Dec) {
	decs, err := utils.GetNonNegativeFixedDecs(6, balanceBefore.TradableAmount, balanceAfter.TradableAmount,
		balanceBefore.EscrowedAmount, balanceAfter.EscrowedAmount)
	assert.NilError(t, err)

	balBeforeTradable, balAfterTradable, balBeforeEscrowed, balAfterEscrowed := decs[0], decs[1], decs[2], decs[3]

	// check the resulting balance -> tradableBefore - orderAmt = tradableAfter
	calculatedTradable, err := balBeforeTradable.Sub(orderAmt)
	assert.NilError(t, err)
	assert.Check(t, calculatedTradable.Equal(balAfterTradable))

	// check the resulting escrow balance -> escrowedBefore + orderAmt = escrowedAfter
	calculatedEscrow, err := balBeforeEscrowed.Add(orderAmt)
	assert.NilError(t, err)
	assert.Check(t, calculatedEscrow.Equal(balAfterEscrowed), "calculated: %s, actual: %s", calculatedEscrow.String(), balAfterEscrowed.String())
}

// testSellSetup sets up a batch, class, market, and issues a balance of 100 retired and tradable to the base suite's addr.
func (s *baseSuite) testSellSetup(batchDenom, bankDenom, displayDenom, classID string, start, end *timestamppb.Timestamp, creditType basetypes.CreditType) {
	assert.Check(s.t, len(s.addrs) > 0, "When calling `testSellSetup`, the base suite must have a non-empty `addrs`.")
	assert.NilError(s.t, s.baseStore.CreditTypeTable().Insert(s.ctx, &baseapi.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "metric ton C02",
		Precision:    6,
	}))

	assert.NilError(s.t, s.baseStore.ClassTable().Insert(s.ctx, &baseapi.Class{
		Id:               classID,
		Admin:            s.addrs[0],
		Metadata:         "",
		CreditTypeAbbrev: creditType.Abbreviation,
	}))
	assert.NilError(s.t, s.baseStore.BatchTable().Insert(s.ctx, &baseapi.Batch{
		ProjectKey: 1,
		Denom:      batchDenom,
		Metadata:   "",
		StartDate:  start,
		EndDate:    end,
	}))

	assert.NilError(s.t, s.marketStore.MarketTable().Insert(s.ctx, &api.Market{
		CreditTypeAbbrev:  creditType.Abbreviation,
		BankDenom:         bankDenom,
		PrecisionModifier: 0,
	}))
	assert.NilError(s.t, s.marketStore.AllowedDenomTable().Insert(s.ctx, &api.AllowedDenom{
		BankDenom:    bankDenom,
		DisplayDenom: displayDenom,
		Exponent:     1,
	}))
	assert.NilError(s.t, s.k.baseStore.BatchBalanceTable().Insert(s.ctx, &baseapi.BatchBalance{
		BatchKey:       1,
		Address:        s.addrs[0],
		TradableAmount: "100",
		RetiredAmount:  "100",
		EscrowedAmount: "0",
	}))
	assert.NilError(s.t, s.k.baseStore.BatchSupplyTable().Insert(s.ctx, &baseapi.BatchSupply{
		BatchKey:        1,
		TradableAmount:  "100",
		RetiredAmount:   "100",
		CancelledAmount: "0",
	}))
}
