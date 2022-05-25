package marketplace

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestSell_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)

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
	s.testSellSetup(batchDenom, "ufoo", "foo", "C01", start, end, creditType)
	sellTime := time.Now()
	newCoin := sdk.NewInt64Coin("ubaz", 10)
	assert.NilError(t, s.marketStore.AllowedDenomTable().Insert(s.ctx, &api.AllowedDenom{
		BankDenom:    newCoin.Denom,
		DisplayDenom: newCoin.Denom,
		Exponent:     18,
	}))

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

func TestSell_Invalid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	sellTime := time.Now()

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
	invalidExpirationTime, err := types.ParseDate("expiration", "1500-01-01")
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
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)

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
