package marketplace

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestSell_Prune(t *testing.T) {
	t.Parallel()
	s := setupBase(t, 1)
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)

	blockTime, err := types.ParseDate("block time", "2020-01-01")
	assert.NilError(t, err)
	expired, err := types.ParseDate("expiration", "2019-12-30")
	assert.NilError(t, err)
	notExpired, err := types.ParseDate("expiration", "2022-01-01")
	assert.NilError(t, err)

	res, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Seller: s.addrs[0].String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, Expiration: &expired},
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, Expiration: &notExpired},
		},
	})
	assert.NilError(t, err)

	// setup block time so the orders expire
	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// get the balance before pruning
	balBefore, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addrs[0], 1)
	assert.NilError(t, err)

	// prune the orders
	err = s.k.PruneSellOrders(s.ctx)
	assert.NilError(t, err)

	balAfter, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addrs[0], 1)
	assert.NilError(t, err)

	// we can reuse this function and pass the negated amount to get our desired behavior.
	assertCreditsEscrowed(t, balBefore, balAfter, math.NewDecFromInt64(-10))

	assert.Equal(t, 2, len(res.SellOrderIds))
	shouldBeExpired := res.SellOrderIds[0]
	shouldBeValid := res.SellOrderIds[1]

	_, err = s.marketStore.SellOrderTable().Get(s.ctx, shouldBeExpired)
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	_, err = s.marketStore.SellOrderTable().Get(s.ctx, shouldBeValid)
	assert.NilError(t, err)
}

// TestPrune_NilExpiration tests that sell orders with nil expirations are not deleted when PruneOrders is called.
func TestPrune_NilExpiration(t *testing.T) {
	t.Parallel()
	s := setupBase(t, 1)
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)

	blockTime, err := types.ParseDate("block time", "2020-01-01")
	assert.NilError(t, err)
	expired, err := types.ParseDate("expiration", "2010-01-01")
	assert.NilError(t, err)

	msg := &marketplace.MsgSell{
		Seller: s.addrs[0].String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "5", AskPrice: &ask, Expiration: nil},
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, Expiration: &expired},
		},
	}
	res, err := s.k.Sell(s.ctx, msg)
	assert.NilError(t, err)

	shouldExistOrder := res.SellOrderIds[0]
	shouldNotExistOrder := res.SellOrderIds[1]

	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	err = s.k.PruneSellOrders(s.ctx)
	assert.NilError(t, err)

	order, err := s.marketStore.SellOrderTable().Get(s.ctx, shouldExistOrder)
	assert.NilError(t, err)
	assert.Equal(t, "5", order.Quantity)

	_, err = s.marketStore.SellOrderTable().Get(s.ctx, shouldNotExistOrder)
	assert.ErrorIs(t, err, ormerrors.NotFound)
}
