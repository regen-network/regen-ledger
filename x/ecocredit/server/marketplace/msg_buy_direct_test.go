package marketplace

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestBuyDirect_MultipleOrders(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	userCoinBalance := sdk.NewInt64Coin(ask.Denom, 1000000)
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)
	// make a sell order
	sellExp := time.Now()

	purchaseAmt1 := "12.3531"
	purchaseAmt2 := "15.39201"
	sellOrderIds := s.createSellOrder(&marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{
				BatchDenom:        batchDenom,
				Quantity:          purchaseAmt1,
				AskPrice:          &ask,
				DisableAutoRetire: true,
				Expiration:        &sellExp,
			},
			{
				BatchDenom:        batchDenom,
				Quantity:          purchaseAmt2,
				AskPrice:          &ask,
				DisableAutoRetire: false,
				Expiration:        &sellExp,
			},
		},
	})

	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, batchDenom)
	assert.NilError(s.t, err)
	balBefore, supBefore := s.getBalanceAndSupply(batch.Key, buyerAddr)

	orders := []*marketplace.MsgBuyDirect_Order{
		{sellOrderIds[0], purchaseAmt1, &ask, true, ""},
		{sellOrderIds[1], purchaseAmt2, &ask, false, ""},
	}

	s.bankKeeper.EXPECT().GetBalance(gmAny, buyerAddr, ask.Denom).Return(userCoinBalance).Times(2)

	// first order is 12.23531 * 10ufoo = 123ufoo
	cost := sdk.Coins{sdk.NewInt64Coin(ask.Denom, 123)}
	s.bankKeeper.EXPECT().SendCoins(gmAny, buyerAddr, s.addr, cost).Return(nil).Times(1)
	// second order is 15.39201 * 10ufoo = 153ufoo
	cost2 := sdk.Coins{sdk.NewInt64Coin(ask.Denom, 153)}
	s.bankKeeper.EXPECT().SendCoins(gmAny, buyerAddr, s.addr, cost2).Return(nil).Times(1)
	_, err = s.k.BuyDirect(s.ctx, &marketplace.MsgBuyDirect{
		Buyer:  buyerAddr.String(),
		Orders: orders,
	})

	balAfter, supAfter := s.getBalanceAndSupply(batch.Key, buyerAddr)
	s.assertBalanceAndSupplyUpdated(orders, balBefore, balAfter, supBefore, supAfter)

}
