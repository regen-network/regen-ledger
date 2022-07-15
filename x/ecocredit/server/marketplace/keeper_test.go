package marketplace

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

var (
	gmAny = gomock.Any()
)

type baseSuite struct {
	t            gocuke.TestingT
	db           ormdb.ModuleDB
	coreStore    ecoApi.StateStore
	marketStore  api.StateStore
	ctx          context.Context
	k            Keeper
	ctrl         *gomock.Controller
	addrs        []sdk.AccAddress
	bankKeeper   *mocks.MockBankKeeper
	paramsKeeper *mocks.MockParamKeeper
	storeKey     *sdk.KVStoreKey
	sdkCtx       sdk.Context
}

func setupBase(t gocuke.TestingT, numAddresses int) *baseSuite {
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

	// setup test keeper
	s.ctrl = gomock.NewController(t)
	assert.NilError(t, err)
	s.bankKeeper = mocks.NewMockBankKeeper(s.ctrl)
	s.paramsKeeper = mocks.NewMockParamKeeper(s.ctrl)
	s.k = NewKeeper(s.marketStore, s.coreStore, s.bankKeeper, s.paramsKeeper)

	// set test accounts
	for i := 0; i < numAddresses; i++ {
		var _, _, addr = testdata.KeyTestPubAddr()
		s.addrs = append(s.addrs, addr)
	}

	return s
}

func (s *baseSuite) createSellOrder(msg *marketplace.MsgSell) []uint64 {
	res, err := s.k.Sell(s.ctx, msg)
	assert.NilError(s.t, err)
	return res.SellOrderIds
}

func (s *baseSuite) getBalanceAndSupply(batchKey uint64, addr sdk.AccAddress) (*ecoApi.BatchBalance, *ecoApi.BatchSupply) {
	bal, err := utils.GetBalance(s.ctx, s.coreStore.BatchBalanceTable(), addr, batchKey)
	assert.NilError(s.t, err)
	sup, err := s.coreStore.BatchSupplyTable().Get(s.ctx, batchKey)
	assert.NilError(s.t, err)
	return bal, sup
}

func (s *baseSuite) assertBalanceAndSupplyUpdated(orders []*marketplace.MsgBuyDirect_Order, b1, b2 *ecoApi.BatchBalance, s1, s2 *ecoApi.BatchSupply) {
	purchaseTradable := math.NewDecFromInt64(0)
	purchaseRetired := math.NewDecFromInt64(0)
	for _, order := range orders {
		if order.DisableAutoRetire {
			qty, err := math.NewDecFromString(order.Quantity)
			assert.NilError(s.t, err)
			purchaseTradable, err = purchaseTradable.Add(qty)
			assert.NilError(s.t, err)
		} else {
			qty, err := math.NewDecFromString(order.Quantity)
			assert.NilError(s.t, err)
			purchaseRetired, err = purchaseRetired.Add(qty)
			assert.NilError(s.t, err)
		}
	}

	balBeforeT, balBeforeR, _ := extractBalanceDecs(s.t, b1)
	balAfterT, balAfterR, _ := extractBalanceDecs(s.t, b2)

	supBeforeT, supBeforeR, _ := extractSupplyDecs(s.t, s1)
	supAfterT, supAfterR, _ := extractSupplyDecs(s.t, s2)

	expectedTradableBal, err := balBeforeT.Add(purchaseTradable)
	assert.NilError(s.t, err)
	expectedRetiredBal, err := balBeforeR.Add(purchaseRetired)
	assert.NilError(s.t, err)
	assert.Check(s.t, balAfterT.Equal(expectedTradableBal))
	assert.Check(s.t, balAfterR.Equal(expectedRetiredBal))

	expectedTradableSup, err := supBeforeT.Sub(purchaseRetired)
	assert.NilError(s.t, err)
	expectedRetiredSup, err := supBeforeR.Add(purchaseRetired)
	assert.NilError(s.t, err)

	assert.Check(s.t, expectedTradableSup.Equal(supAfterT))
	assert.Check(s.t, expectedRetiredSup.Equal(supAfterR))

}

func extractBalanceDecs(t gocuke.TestingT, b *ecoApi.BatchBalance) (tradable, retired, escrowed math.Dec) {
	decs, err := utils.GetNonNegativeFixedDecs(6, b.TradableAmount, b.RetiredAmount, b.EscrowedAmount)
	assert.NilError(t, err)
	return decs[0], decs[1], decs[2]
}

func extractSupplyDecs(t gocuke.TestingT, s *ecoApi.BatchSupply) (tradable, retired, cancelled math.Dec) {
	decs, err := utils.GetNonNegativeFixedDecs(6, s.TradableAmount, s.RetiredAmount, s.CancelledAmount)
	assert.NilError(t, err)
	return decs[0], decs[1], decs[2]
}

// assertCreditsEscrowed adds orderAmt to tradable, subtracts from escrowed in before balance/supply and checks that it is equal to after balance/supply.
func assertCreditsEscrowed(t gocuke.TestingT, balanceBefore, balanceAfter *ecoApi.BatchBalance, orderAmt math.Dec) {
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
func (s *baseSuite) testSellSetup(batchDenom, bankDenom, displayDenom, classId string, start, end *timestamppb.Timestamp, creditType core.CreditType) {
	assert.Check(s.t, len(s.addrs) > 0, "When calling `testSellSetup`, the base suite must have a non-empty `addrs`.")
	assert.NilError(s.t, s.coreStore.CreditTypeTable().Insert(s.ctx, &ecoApi.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "metric ton C02",
		Precision:    6,
	}))

	assert.NilError(s.t, s.coreStore.ClassTable().Insert(s.ctx, &ecoApi.Class{
		Id:               classId,
		Admin:            s.addrs[0],
		Metadata:         "",
		CreditTypeAbbrev: creditType.Abbreviation,
	}))
	assert.NilError(s.t, s.coreStore.BatchTable().Insert(s.ctx, &ecoApi.Batch{
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
	assert.NilError(s.t, s.k.coreStore.BatchBalanceTable().Insert(s.ctx, &ecoApi.BatchBalance{
		BatchKey:       1,
		Address:        s.addrs[0],
		TradableAmount: "100",
		RetiredAmount:  "100",
		EscrowedAmount: "0",
	}))
	assert.NilError(s.t, s.k.coreStore.BatchSupplyTable().Insert(s.ctx, &ecoApi.BatchSupply{
		BatchKey:        1,
		TradableAmount:  "100",
		RetiredAmount:   "100",
		CancelledAmount: "0",
	}))
}
