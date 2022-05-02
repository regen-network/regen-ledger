package marketplace

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp/cmpopts"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestSell_CancelOrder(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	expir := time.Now()
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.AllowedAskDenoms = []*core.AskDenom{{Denom: ask.Denom}}
	}).Times(1)

	balBefore, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	supBefore, err := s.coreStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	res, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, Expiration: &expir},
		},
	})
	assert.NilError(t, err)

	_, err = s.k.CancelSellOrder(s.ctx, &marketplace.MsgCancelSellOrder{SellOrderId: res.SellOrderIds[0], Seller: s.addr.String()})
	assert.NilError(t, err)

	balAfter, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	supAfter, err := s.coreStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	assert.DeepEqual(t, balBefore, balAfter, cmpopts.IgnoreUnexported(api.BatchBalance{}))
	assert.DeepEqual(t, supBefore, supAfter, cmpopts.IgnoreUnexported(api.BatchSupply{}))
}

func TestSell_CancelOrderInvalid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	expir := time.Now()
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)

	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.AllowedAskDenoms = []*core.AskDenom{{Denom: ask.Denom}}
	}).Times(1)

	_, _, otherAddr := testdata.KeyTestPubAddr()

	res, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, Expiration: &expir},
		},
	})
	assert.NilError(t, err)

	_, err = s.k.CancelSellOrder(s.ctx, &marketplace.MsgCancelSellOrder{
		Seller:      otherAddr.String(),
		SellOrderId: res.SellOrderIds[0],
	})
	assert.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())
}
