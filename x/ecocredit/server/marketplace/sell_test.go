package marketplace

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

func TestSell_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.CreditTypes = []*core.CreditType{&creditType}
		p.AllowedAskDenoms = []*core.AskDenom{{Denom: ask.Denom}}
	}).Times(4)

	balanceBefore, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	supplyBefore, err := s.coreStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	sellTime := time.Now()
	res, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellTime},
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellTime},
		},
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.SellOrderIds))

	it, err := s.marketStore.SellOrderTable().List(s.ctx, api.SellOrderSellerIndexKey{})
	assert.NilError(t, err)
	count := 0
	for it.Next() {
		val, err := it.Value()
		assert.NilError(t, err)
		assert.Equal(t, "10", val.Quantity)
		assert.Equal(t, ask.Amount.String(), val.AskPrice)
		count++
	}
	assert.Equal(t, 2, count)

	balanceAfter, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	supplyAfter, err := s.coreStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	assertCreditsEscrowed(t, balanceBefore, balanceAfter, supplyBefore, supplyAfter, math.NewDecFromInt64(20))
}

func TestSell_CreatesMarket(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	testSellSetup(t, s, batchDenom, "ufoo", "foo", "C01", start, end, creditType)
	sellTime := time.Now()
	newCoin := sdk.NewInt64Coin("ubaz", 10)
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.CreditTypes = []*core.CreditType{&creditType}
		p.AllowedAskDenoms = []*core.AskDenom{{Denom: newCoin.Denom}}
	}).Times(2)

	// market shouldn't exist before sell call
	has, err := s.k.stateStore.MarketTable().HasByCreditTypeBankDenom(s.ctx, creditType.Abbreviation, newCoin.Denom)
	assert.NilError(t, err)
	assert.Equal(t, false, has)

	_, err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &newCoin, DisableAutoRetire: false, Expiration: &sellTime},
		},
	})
	assert.NilError(t, err)

	// market should exist now
	has, err = s.k.stateStore.MarketTable().HasByCreditTypeBankDenom(s.ctx, creditType.Abbreviation, newCoin.Denom)
	assert.NilError(t, err)
	assert.Equal(t, true, has)
}

// TODO: add a check once params are refactored and the ask denom param is active - https://github.com/regen-network/regen-ledger/issues/624
func TestSell_Invalid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	sellTime := time.Now()

	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.CreditTypes = []*core.CreditType{&creditType}
		p.AllowedAskDenoms = []*core.AskDenom{{Denom: ask.Denom}}
	}).Times(2)

	// invalid batch
	_, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: "foo-bar-baz-001", Quantity: "10", AskPrice: &ask, DisableAutoRetire: true, Expiration: &sellTime},
		},
	})
	assert.ErrorContains(t, err, "batch denom foo-bar-baz-001")

	// invalid balance
	_, err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10000000000", AskPrice: &ask, DisableAutoRetire: true, Expiration: &sellTime},
		},
	})
	assert.ErrorContains(t, err, "insufficient funds")

	// order expiration not in the future
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
	invalidExpirationTime, err := time.Parse("2006-01-02", "1500-01-01")
	_, err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: true, Expiration: &invalidExpirationTime},
		},
	})
	assert.ErrorContains(t, err, "expiration must be in the future")
}

func TestSell_InvalidDenom(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.CreditTypes = []*core.CreditType{&creditType}
		p.AllowedAskDenoms = []*core.AskDenom{{Denom: ask.Denom}}
	}).Times(2)

	sellTime := time.Now()
	invalidAsk := sdk.NewInt64Coin("ubar", 10)
	_, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &invalidAsk, DisableAutoRetire: false, Expiration: &sellTime},
		},
	})
	assert.ErrorContains(t, err, "ubar is not allowed to be used in sell orders")
}

// assertCreditsEscrowed adds orderAmt to tradable, subtracts from escrowed in before balance/supply and checks that it is equal to after balance/supply.
func assertCreditsEscrowed(t *testing.T, balanceBefore, balanceAfter *ecoApi.BatchBalance, supplyBefore, supplyAfter *ecoApi.BatchSupply, orderAmt math.Dec) {
	decs, err := utils.GetNonNegativeFixedDecs(6, balanceBefore.Tradable, balanceAfter.Tradable,
		balanceBefore.Escrowed, balanceAfter.Escrowed, supplyBefore.TradableAmount, supplyAfter.TradableAmount)
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
func testSellSetup(t *testing.T, s *baseSuite, batchDenom, bankDenom, displayDenom, classId string, start, end *timestamppb.Timestamp, creditType core.CreditType) {
	assert.NilError(t, s.coreStore.BatchTable().Insert(s.ctx, &ecoApi.Batch{
		ProjectKey: 1,
		Denom:      batchDenom,
		Metadata:   "",
		StartDate:  start,
		EndDate:    end,
	}))
	assert.NilError(t, s.coreStore.ClassTable().Insert(s.ctx, &ecoApi.Class{
		Id:               classId,
		Admin:            s.addr,
		Metadata:         "",
		CreditTypeAbbrev: creditType.Abbreviation,
	}))
	assert.NilError(t, s.marketStore.MarketTable().Insert(s.ctx, &api.Market{
		CreditType:        creditType.Abbreviation,
		BankDenom:         bankDenom,
		PrecisionModifier: 0,
	}))
	// TODO: awaiting param refactor https://github.com/regen-network/regen-ledger/issues/624
	//assert.NilError(t, s.marketStore.AllowedDenomTable().Insert(s.ctx, &marketApi.AllowedDenom{
	//	BankDenom:    bankDenom,
	//	DisplayDenom: displayDenom,
	//	Exponent:     1,
	//}))
	assert.NilError(t, s.k.coreStore.BatchBalanceTable().Insert(s.ctx, &ecoApi.BatchBalance{
		BatchKey: 1,
		Address:  s.addr,
		Tradable: "100",
		Retired:  "100",
		Escrowed: "0",
	}))
	assert.NilError(t, s.k.coreStore.BatchSupplyTable().Insert(s.ctx, &ecoApi.BatchSupply{
		BatchKey:       1,
		TradableAmount: "100",
		RetiredAmount:  "100",
	}))
}
