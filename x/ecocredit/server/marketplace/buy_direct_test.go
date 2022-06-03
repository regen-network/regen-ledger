package marketplace

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestBuyDirect_ValidTradable(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	// make a sell order
	sellExp := time.Now()
	userCoinBalance := sdk.NewInt64Coin(validAskDenom, 30)
	sellOrderId := s.createSellOrder(&marketplace.MsgSell{
		Seller: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: true, Expiration: &sellExp},
		},
	})[0]

	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, batchDenom)
	assert.NilError(s.t, err)
	balBefore, supBefore := s.getBalanceAndSupply(batch.Key, buyerAddr)

	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(userCoinBalance).Times(1)
	// sell order ask price: 10, buy order of 3 credits -> 10 * 3 = 30
	s.bankKeeper.EXPECT().SendCoins(gmAny, gmAny, gmAny, sdk.Coins{userCoinBalance}).Return(nil).Times(1)
	purchaseAmt := math.NewDecFromInt64(3)
	order := &marketplace.MsgBuyDirect_Order{
		SellOrderId:       sellOrderId,
		Quantity:          purchaseAmt.String(),
		BidPrice:          &ask,
		DisableAutoRetire: true}
	err = buyDirectSingle(s, buyerAddr, order)
	assert.NilError(t, err)

	balAfter, supAfter := s.getBalanceAndSupply(batch.Key, buyerAddr)
	s.assertBalanceAndSupplyUpdated([]*marketplace.MsgBuyDirect_Order{order}, balBefore, balAfter, supBefore, supAfter)
}

func TestBuyDirect_ValidRetired(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	userBalance := sdk.NewInt64Coin(validAskDenom, 30)
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	// make a sell order
	sellExp := time.Now()
	sellOrderId := s.createSellOrder(&marketplace.MsgSell{
		Seller: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})[0]

	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, batchDenom)
	assert.NilError(s.t, err)
	balBefore, supBefore := s.getBalanceAndSupply(batch.Key, buyerAddr)

	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(userBalance).Times(1)
	s.bankKeeper.EXPECT().SendCoins(gmAny, gmAny, gmAny, gmAny).Return(nil).Times(1)

	purchaseAmt := math.NewDecFromInt64(3)
	order := &marketplace.MsgBuyDirect_Order{
		SellOrderId:            sellOrderId,
		Quantity:               purchaseAmt.String(),
		BidPrice:               &ask,
		DisableAutoRetire:      false,
		RetirementJurisdiction: "US-NY",
	}
	err = buyDirectSingle(s, buyerAddr, order)
	assert.NilError(t, err)

	balAfter, supAfter := s.getBalanceAndSupply(batch.Key, buyerAddr)
	s.assertBalanceAndSupplyUpdated([]*marketplace.MsgBuyDirect_Order{order}, balBefore, balAfter, supBefore, supAfter)
}

func TestBuyDirect_OrderFilled(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	userBalance := sdk.NewInt64Coin(validAskDenom, 100)
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	// make a sell order
	sellExp := time.Now()
	sellOrderId := s.createSellOrder(&marketplace.MsgSell{
		Seller: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})[0]

	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(userBalance).Times(1)
	s.bankKeeper.EXPECT().SendCoins(gmAny, gmAny, gmAny, gmAny).Return(nil).Times(1)

	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, batchDenom)
	assert.NilError(t, err)
	balBefore, supBefore := s.getBalanceAndSupply(batch.Key, buyerAddr)

	purchaseAmt := math.NewDecFromInt64(10)
	order := &marketplace.MsgBuyDirect_Order{
		SellOrderId:            sellOrderId,
		Quantity:               purchaseAmt.String(),
		BidPrice:               &ask,
		RetirementJurisdiction: "US-OR",
	}
	err = buyDirectSingle(s, buyerAddr, order)
	assert.NilError(t, err)

	balAfter, supAfter := s.getBalanceAndSupply(batch.Key, buyerAddr)

	// order was filled, so sell order should no longer exist
	_, err = s.marketStore.SellOrderTable().Get(s.ctx, sellOrderId)
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
	s.assertBalanceAndSupplyUpdated([]*marketplace.MsgBuyDirect_Order{order}, balBefore, balAfter, supBefore, supAfter)
}

func TestBuyDirect_Invalid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)
	userBalance := sdk.NewInt64Coin(validAskDenom, 150)

	// make a sell order
	sellExp := time.Now()
	sellOrderId := s.createSellOrder(&marketplace.MsgSell{
		Seller: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})[0]

	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(userBalance).Times(1)

	// sell order not found
	err := buyDirectSingle(s, buyerAddr, &marketplace.MsgBuyDirect_Order{
		SellOrderId:            532,
		Quantity:               "10",
		BidPrice:               &ask,
		RetirementJurisdiction: "US-CA"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// exceeds decimal precision
	err = buyDirectSingle(s, buyerAddr, &marketplace.MsgBuyDirect_Order{
		SellOrderId:            sellOrderId,
		Quantity:               "10.3235235235",
		BidPrice:               &ask,
		RetirementJurisdiction: "US-CA"})
	assert.ErrorContains(t, err, "exceeds maximum decimal places")

	// mismatch auto retire settings
	err = buyDirectSingle(s, buyerAddr, &marketplace.MsgBuyDirect_Order{
		SellOrderId:       sellOrderId,
		Quantity:          "10",
		BidPrice:          &ask,
		DisableAutoRetire: true})
	assert.ErrorContains(t, err, "cannot disable auto retire")

	// cannot buy more credits than available
	err = buyDirectSingle(s, buyerAddr, &marketplace.MsgBuyDirect_Order{
		SellOrderId:            sellOrderId,
		Quantity:               "11",
		BidPrice:               &ask,
		RetirementJurisdiction: "US-WA"})
	assert.ErrorContains(t, err, "cannot purchase 11 credits from a sell order that has 10 credits")

	// mismatchDenom
	wrongDenom := sdk.NewInt64Coin("ubar", 10)
	err = buyDirectSingle(s, buyerAddr, &marketplace.MsgBuyDirect_Order{
		SellOrderId:            sellOrderId,
		Quantity:               "10",
		BidPrice:               &wrongDenom,
		RetirementJurisdiction: "US-CO"})
	assert.ErrorContains(t, err, "bid price denom does not match ask price denom")

	// bidding more than in the bank
	inBank := sdk.NewInt64Coin(validAskDenom, 10)
	biddingWith := sdk.NewInt64Coin(validAskDenom, 100)
	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(inBank).Times(1)
	err = buyDirectSingle(s, buyerAddr, &marketplace.MsgBuyDirect_Order{
		SellOrderId:            sellOrderId,
		Quantity:               "10",
		BidPrice:               &biddingWith,
		RetirementJurisdiction: "US-NV"})
	assert.ErrorContains(t, err, sdkerrors.ErrInsufficientFunds.Error())
}

func TestBuyDirect_Decimal(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	userCoinBalance := sdk.NewInt64Coin(validAskDenom, 50)
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)
	// make a sell order
	sellExp := time.Now()
	sellOrderId := s.createSellOrder(&marketplace.MsgSell{
		Seller: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{
				BatchDenom:        batchDenom,
				Quantity:          "10",
				AskPrice:          &ask,
				DisableAutoRetire: true,
				Expiration:        &sellExp,
			},
		},
	})[0]

	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, batchDenom)
	assert.NilError(s.t, err)
	balBefore, supBefore := s.getBalanceAndSupply(batch.Key, buyerAddr)

	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(userCoinBalance).Times(1)

	purchaseAmt := "3.985321"
	expectedCost := sdk.NewInt64Coin(validAskDenom, 39)
	// sell order ask price: 10, buy order of 3.215 credits -> 10 * 3.215 = 32.15
	s.bankKeeper.EXPECT().SendCoins(gmAny, gmAny, gmAny, sdk.Coins{expectedCost}).Return(nil).Times(1)

	err = buyDirectSingle(s, buyerAddr, &marketplace.MsgBuyDirect_Order{
		SellOrderId:       sellOrderId,
		Quantity:          purchaseAmt,
		BidPrice:          &ask,
		DisableAutoRetire: true})
	assert.NilError(t, err)

	balAfter, supAfter := s.getBalanceAndSupply(batch.Key, buyerAddr)

	s.assertBalanceAndSupplyUpdated([]*marketplace.MsgBuyDirect_Order{{
		Quantity:          purchaseAmt,
		DisableAutoRetire: true,
	}}, balBefore, balAfter, supBefore, supAfter)

}

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
		Seller: s.addr.String(),
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
