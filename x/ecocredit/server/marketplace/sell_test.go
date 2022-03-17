package marketplace

import (
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"

	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestSell_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	ask := sdk.NewInt64Coin("ufoo", 10)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).Times(2)

	balanceBefore, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	supplyBefore, err := s.coreStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	sellTime := time.Now()
	res, err := s.k.Sell(s.ctx, &v1.MsgSell{
		Owner: s.addr.String(),
		Orders: []*v1.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellTime},
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellTime},
		},
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.SellOrderIds))

	it, err := s.marketStore.SellOrderTable().List(s.ctx, marketApi.SellOrderSellerIndexKey{})
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

	assertCoinsEscrowed(t, balanceBefore, balanceAfter, supplyBefore, supplyAfter, math.NewDecFromInt64(20))
}

func TestSell_CreatesMarket(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	any := gomock.Any()
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	ask := sdk.NewInt64Coin("ubar", 10)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	testSellSetup(t, s, batchDenom, "ufoo", "foo", "C01", start, end, creditType)
	sellTime := time.Now()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).Times(1)

	// market shouldn't exist before sell call
	has, err := s.k.stateStore.MarketTable().HasByCreditTypeBankDenom(s.ctx, creditType.Abbreviation, ask.Denom)
	assert.NilError(t, err)
	assert.Equal(t, false, has)

	_, err = s.k.Sell(s.ctx, &v1.MsgSell{
		Owner: s.addr.String(),
		Orders: []*v1.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellTime},
		},
	})
	assert.NilError(t, err)

	// market should exist now
	has, err = s.k.stateStore.MarketTable().HasByCreditTypeBankDenom(s.ctx, creditType.Abbreviation, ask.Denom)
	assert.NilError(t, err)
	assert.Equal(t, true, has)
}

// TODO: add a check once params are refactored and the ask denom param is active - https://github.com/regen-network/regen-ledger/issues/624
func TestSell_Invalid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	any := gomock.Any()
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	ask := sdk.NewInt64Coin("ufoo", 10)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	sellTime := time.Now()

	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).Times(2)

	// invalid batch
	_, err := s.k.Sell(s.ctx, &v1.MsgSell{
		Owner: s.addr.String(),
		Orders: []*v1.MsgSell_Order{
			{BatchDenom: "foo-bar-baz-001", Quantity: "10", AskPrice: &ask, DisableAutoRetire: true, Expiration: &sellTime},
		},
	})
	assert.ErrorContains(t, err, "batch denom foo-bar-baz-001")

	// invalid balance
	_, err = s.k.Sell(s.ctx, &v1.MsgSell{
		Owner: s.addr.String(),
		Orders: []*v1.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10000000000", AskPrice: &ask, DisableAutoRetire: true, Expiration: &sellTime},
		},
	})
	assert.ErrorContains(t, err, "insufficient funds")

	// order expiration not in the future
	s.sdkCtx = s.sdkCtx.WithBlockTime(time.Now())
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
	invalidExpirationTime, err := time.Parse("2006-01-02", "1500-01-01")
	_, err = s.k.Sell(s.ctx, &v1.MsgSell{
		Owner: s.addr.String(),
		Orders: []*v1.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: true, Expiration: &invalidExpirationTime},
		},
	})
	assert.ErrorContains(t, err, "expiration must be in the future")
}

// assertCoinsEscrowed adds orderAmt to tradable, subtracts from escrowed in before balance/supply and checks that it is equal to after balance/supply.
func assertCoinsEscrowed(t *testing.T, balanceBefore, balanceAfter *ecocreditv1.BatchBalance, supplyBefore, supplyAfter *ecocreditv1.BatchSupply, orderAmt math.Dec) {
	decs, err := server.GetNonNegativeFixedDecs(6, balanceBefore.Tradable, balanceAfter.Tradable,
		balanceBefore.Escrowed, balanceAfter.Escrowed, supplyBefore.TradableAmount, supplyAfter.TradableAmount,
		supplyBefore.EscrowedAmount, supplyAfter.EscrowedAmount)
	assert.NilError(t, err)
	balBeforeTradable, balAfterTradable, balBeforeEscrowed, balAfterEscrowed, supBeforeTradable, supAfterTradable,
	supBeforeEscrowed, supAfterEscrowed := decs[0], decs[1], decs[2], decs[3], decs[4], decs[5], decs[6], decs[7]

	// check the resulting balance -> tradableBefore - orderAmt = tradableAfter
	calculatedTradable, err := balBeforeTradable.Sub(orderAmt)
	assert.NilError(t, err)
	assert.Check(t, calculatedTradable.Equal(balAfterTradable))

	// check the resulting escrow balance -> escrowedBefore + orderAmt = escrowedAfter
	calculatedEscrow, err := balBeforeEscrowed.Add(orderAmt)
	assert.NilError(t, err)
	assert.Check(t, calculatedEscrow.Equal(balAfterEscrowed), "calculated: %s, actual: %s", calculatedEscrow.String(), balAfterEscrowed.String())

	// check the resulting tradable supply -> tradableBefore - orderAmt = tradableAfter
	calculatedTSupply, err := supBeforeTradable.Sub(orderAmt)
	assert.NilError(t, err)
	assert.Check(t, calculatedTSupply.Equal(supAfterTradable))

	// check the resulting escrowed supply -> escrowedBefore + orderAmt = escrowedAfter
	calculatedESupply, err := supBeforeEscrowed.Add(orderAmt)
	assert.NilError(t, err)
	assert.Check(t, calculatedESupply.Equal(supAfterEscrowed))
}


func testSellSetup(t *testing.T, s *baseSuite, batchDenom, bankDenom, displayDenom, classId string, start, end *timestamppb.Timestamp, creditType ecocredit.CreditType) {
	assert.NilError(t, s.coreStore.BatchInfoTable().Insert(s.ctx, &ecocreditv1.BatchInfo{
		ProjectId:  1,
		BatchDenom: batchDenom,
		Metadata:   "",
		StartDate:  start,
		EndDate:    end,
	}))
	assert.NilError(t, s.coreStore.ClassInfoTable().Insert(s.ctx, &ecocreditv1.ClassInfo{
		Name:       classId,
		Admin:      s.addr,
		Metadata:   "",
		CreditType: creditType.Abbreviation,
	}))
	assert.NilError(t, s.marketStore.MarketTable().Insert(s.ctx, &marketApi.Market{
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
	assert.NilError(t, s.k.coreStore.BatchBalanceTable().Insert(s.ctx, &ecocreditv1.BatchBalance{
		Address:  s.addr,
		BatchId:  1,
		Tradable: "100",
		Retired:  "100",
	}))
	assert.NilError(t, s.k.coreStore.BatchSupplyTable().Insert(s.ctx, &ecocreditv1.BatchSupply{
		BatchId:  1,
		TradableAmount: "100",
		RetiredAmount:  "100",
	}))
}
