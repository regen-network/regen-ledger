package marketplace

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestSell_Prune(t *testing.T) {
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
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).Times(2)

	blockTime, err := time.Parse("2006-01-02", "2020-01-01")
	assert.NilError(t, err)
	expired, err := time.Parse("2006-01-02", "2019-12-30")
	assert.NilError(t, err)
	notExpired, err := time.Parse("2006-01-02", "2022-01-01")
	assert.NilError(t, err)

	res, err := s.k.Sell(s.ctx, &v1.MsgSell{
		Owner: s.addr.String(),
		Orders: []*v1.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, Expiration: &expired},
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, Expiration: &notExpired},
		},
	})
	assert.NilError(t, err)

	// setup block time so the orders expire
	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// get the balance before pruning
	balBefore, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	supBefore, err := s.coreStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	// prune the orders
	err = s.k.PruneSellOrders(s.ctx)
	assert.NilError(t, err)

	balAfter, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	supAfter, err := s.coreStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	// we can reuse this function and pass the negated amount to get our desired behavior.
	assertCreditsEscrowed(t, balBefore, balAfter, supBefore, supAfter, math.NewDecFromInt64(-10))

	assert.Equal(t, 2, len(res.SellOrderIds))
	shouldBeExpired := res.SellOrderIds[0]
	shouldBeValid := res.SellOrderIds[1]

	_, err = s.marketStore.SellOrderTable().Get(s.ctx, shouldBeExpired)
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	_, err = s.marketStore.SellOrderTable().Get(s.ctx, shouldBeValid)
	assert.NilError(t, err)
}
