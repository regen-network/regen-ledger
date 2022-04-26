package marketplace

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestBuy_ValidTradable(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	userCoinBalance := sdk.NewInt64Coin("ufoo", 30)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	// make a sell order
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.CreditTypes = []*core.CreditType{&creditType}
		p.AllowedAskDenoms = []*core.AskDenom{{Denom: ask.Denom}}
	}).Times(3)
	sellExp := time.Now()
	res, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: true, Expiration: &sellExp},
		},
	})
	assert.NilError(t, err)
	sellOrderId := res.SellOrderIds[0]

	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(userCoinBalance).Times(1)
	// sell order ask price: 10ufoo, buy order of 3 credits -> 10 * 3 = 30ufoo
	s.bankKeeper.EXPECT().SendCoins(gmAny, gmAny, gmAny, sdk.Coins{sdk.NewInt64Coin("ufoo", 30)}).Return(nil).Times(1)

	purchaseAmt := math.NewDecFromInt64(3)
	_, err = buyDirect(s, buyerAddr.String(), sellOrderId, purchaseAmt.String(), &ask, true, "")
	assert.NilError(t, err)

	// sell order should now have quantity 10 - 3 -> 7
	sellOrder, err := s.marketStore.SellOrderTable().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, "7", sellOrder.Quantity)

	// buyer didn't have credits before, so they should now have 3 credits
	buyerBal, err := s.coreStore.BatchBalanceTable().Get(s.ctx, buyerAddr, 1)
	assert.NilError(t, err)
	tradableBalance, err := math.NewDecFromString(buyerBal.Tradable)
	assert.NilError(t, err)
	assert.Check(t, tradableBalance.Equal(purchaseAmt))
}

func TestBuy_ValidRetired(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	userBalance := sdk.NewInt64Coin("ufoo", 30)

	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)
	// make a sell order
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.CreditTypes = []*core.CreditType{&creditType}
		p.AllowedAskDenoms = []*core.AskDenom{{Denom: ask.Denom}}
	}).Times(3)
	sellExp := time.Now()
	res, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})
	assert.NilError(t, err)
	sellOrderId := res.SellOrderIds[0]

	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(userBalance).Times(1)
	s.bankKeeper.EXPECT().SendCoins(gmAny, gmAny, gmAny, gmAny).Return(nil).Times(1)

	purchaseAmt := math.NewDecFromInt64(3)
	_, err = buyDirect(s, buyerAddr.String(), sellOrderId, purchaseAmt.String(), &ask, false, "US-NY")
	assert.NilError(t, err)

	// sell order should now have quantity 10 - 3 -> 7
	sellOrder, err := s.marketStore.SellOrderTable().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, "7", sellOrder.Quantity)

	// buyer didn't have credits before, so they should now have 3 credits
	buyerBal, err := s.coreStore.BatchBalanceTable().Get(s.ctx, buyerAddr, 1)
	assert.NilError(t, err)
	retiredBalance, err := math.NewDecFromString(buyerBal.Retired)
	assert.NilError(t, err)
	assert.Check(t, retiredBalance.Equal(math.NewDecFromInt64(3)))
}

func TestBuy_OrderFilled(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	userBalance := sdk.NewInt64Coin("ufoo", 100)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)
	// make a sell order
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.CreditTypes = []*core.CreditType{&creditType}
		p.AllowedAskDenoms = []*core.AskDenom{{Denom: ask.Denom}}
	}).Times(3)
	sellExp := time.Now()
	res, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})
	assert.NilError(t, err)
	sellOrderId := res.SellOrderIds[0]

	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(userBalance).Times(1)
	s.bankKeeper.EXPECT().SendCoins(gmAny, gmAny, gmAny, gmAny).Return(nil).Times(1)

	purchaseAmt := math.NewDecFromInt64(10)
	_, err = buyDirect(s, buyerAddr.String(), sellOrderId, purchaseAmt.String(), &ask, false, "US-OR")
	assert.NilError(t, err)

	// order was filled, so sell order should no longer exist
	_, err = s.marketStore.SellOrderTable().Get(s.ctx, sellOrderId)
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
	buyerBal, err := s.coreStore.BatchBalanceTable().Get(s.ctx, buyerAddr, 1)
	assert.NilError(t, err)
	retiredBal, err := math.NewDecFromString(buyerBal.Retired)
	assert.NilError(t, err)
	assert.Check(t, retiredBal.Equal(math.NewDecFromInt64(10)))
}

func TestBuy_Invalid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	userBalance := sdk.NewInt64Coin("ufoo", 150)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)
	// make a sell order
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.CreditTypes = []*core.CreditType{&creditType}
		p.AllowedAskDenoms = []*core.AskDenom{{Denom: ask.Denom}}
	}).AnyTimes()
	sellExp := time.Now()
	res, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})
	assert.NilError(t, err)
	sellOrderId := res.SellOrderIds[0]

	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(userBalance).Times(1)

	// sell order not found
	_, err = buyDirect(s, buyerAddr.String(), 532, "10", &ask, false, "US-CA")
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// exceeds decimal precision
	_, err = buyDirect(s, buyerAddr.String(), sellOrderId, "10.3235235235", &ask, false, "US-OR")
	assert.ErrorContains(t, err, "exceeds maximum decimal places")

	// mismatch auto retire settings
	_, err = buyDirect(s, buyerAddr.String(), sellOrderId, "10", &ask, true, "")
	assert.ErrorContains(t, err, "cannot disable auto retire")

	// cannot buy more credits than available
	_, err = buyDirect(s, buyerAddr.String(), sellOrderId, "11", &ask, false, "US-WA")
	assert.ErrorContains(t, err, "cannot purchase 11 credits from a sell order that has 10 credits")

	// mismatchDenom
	wrongDenom := sdk.NewInt64Coin("ubar", 10)
	_, err = buyDirect(s, buyerAddr.String(), sellOrderId, "10", &wrongDenom, false, "US-CO")
	assert.ErrorContains(t, err, "bid price denom does not match ask price denom")

	// bidding more than in the bank
	inBank := sdk.NewInt64Coin("ufoo", 10)
	biddingWith := sdk.NewInt64Coin("ufoo", 100)
	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(inBank).Times(1)
	_, err = buyDirect(s, buyerAddr.String(), sellOrderId, "10", &biddingWith, false, "US-NV")
	assert.ErrorContains(t, err, sdkerrors.ErrInsufficientFunds.Error())
}

func TestBuy_Decimal(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	userCoinBalance := sdk.NewInt64Coin("ufoo", 50)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	// make a sell order
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.CreditTypes = []*core.CreditType{&creditType}
		p.AllowedAskDenoms = []*core.AskDenom{{Denom: "ufoo"}}
	}).Times(3)
	sellExp := time.Now()
	res, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: true, Expiration: &sellExp},
		},
	})
	assert.NilError(t, err)
	sellOrderId := res.SellOrderIds[0]

	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(userCoinBalance).Times(1)

	purchaseAmt := "3.985321"
	expectedCost := sdk.NewInt64Coin("ufoo", 39)
	// sell order ask price: 10ufoo, buy order of 3.215 credits -> 10 * 3.215 = 32.15
	s.bankKeeper.EXPECT().SendCoins(gmAny, gmAny, gmAny, sdk.Coins{expectedCost}).Return(nil).Times(1)

	_, err = buyDirect(s, buyerAddr.String(), sellOrderId, purchaseAmt, &ask, true, "")
	assert.NilError(t, err)
}

func buyDirect(s *baseSuite, buyer string, sellOrderId uint64, qty string, pricePerCredit *sdk.Coin, disableAutoRetire bool,
	jurisdiction string) (*marketplace.MsgBuyDirectResponse, error) {
	return s.k.BuyDirect(s.ctx, &marketplace.MsgBuyDirect{
		Buyer: buyer,
		Orders: []*marketplace.MsgBuyDirect_Order{
			{
				SellOrderId:            sellOrderId,
				Quantity:               qty,
				BidPrice:               pricePerCredit,
				DisableAutoRetire:      disableAutoRetire,
				RetirementJurisdiction: jurisdiction,
			},
		},
	})
}
