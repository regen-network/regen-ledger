package marketplace

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
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
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	ask := sdk.NewInt64Coin("ufoo", 10)
	userCoinBalance := sdk.NewInt64Coin("ufoo", 30)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	// make a sell order
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).Times(2)
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
	_, err = s.k.Buy(s.ctx, &marketplace.MsgBuy{
		Buyer: buyerAddr.String(),
		Orders: []*marketplace.MsgBuy_Order{
			{Selection: &marketplace.MsgBuy_Order_Selection{Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellOrderId}},
				Quantity: purchaseAmt.String(), BidPrice: &ask, DisableAutoRetire: true, Expiration: &sellExp},
		},
	})
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
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	ask := sdk.NewInt64Coin("ufoo", 10)
	userBalance := sdk.NewInt64Coin("ufoo", 30)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	// make a sell order
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).Times(2)
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
	_, err = s.k.Buy(s.ctx, &marketplace.MsgBuy{
		Buyer: buyerAddr.String(),
		Orders: []*marketplace.MsgBuy_Order{
			{Selection: &marketplace.MsgBuy_Order_Selection{Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellOrderId}},
				Quantity: purchaseAmt.String(), BidPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})
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
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	ask := sdk.NewInt64Coin("ufoo", 10)
	userBalance := sdk.NewInt64Coin("ufoo", 100)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	// make a sell order
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).Times(2)
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
	_, err = s.k.Buy(s.ctx, &marketplace.MsgBuy{
		Buyer: buyerAddr.String(),
		Orders: []*marketplace.MsgBuy_Order{
			{Selection: &marketplace.MsgBuy_Order_Selection{Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellOrderId}},
				Quantity: purchaseAmt.String(), BidPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})
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
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	ask := sdk.NewInt64Coin("ufoo", 10)
	userBalance := sdk.NewInt64Coin("ufoo", 150)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	// make a sell order
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
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
	_, err = s.k.Buy(s.ctx, &marketplace.MsgBuy{
		Buyer: buyerAddr.String(),
		Orders: []*marketplace.MsgBuy_Order{
			{Selection: &marketplace.MsgBuy_Order_Selection{Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: 532}},
				Quantity: "10", BidPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// exceeds decimal precision
	_, err = s.k.Buy(s.ctx, &marketplace.MsgBuy{
		Buyer: buyerAddr.String(),
		Orders: []*marketplace.MsgBuy_Order{
			{Selection: &marketplace.MsgBuy_Order_Selection{Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellOrderId}},
				Quantity: "10.23958230958", BidPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})
	assert.ErrorContains(t, err, "exceeds maximum decimal places")

	// mismatch auto retire settings
	_, err = s.k.Buy(s.ctx, &marketplace.MsgBuy{
		Buyer: buyerAddr.String(),
		Orders: []*marketplace.MsgBuy_Order{
			{Selection: &marketplace.MsgBuy_Order_Selection{Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellOrderId}},
				Quantity: "10", BidPrice: &ask, DisableAutoRetire: true, Expiration: &sellExp},
		},
	})
	assert.ErrorContains(t, err, "cannot disable auto retire")

	// cannot buy more credits than available
	_, err = s.k.Buy(s.ctx, &marketplace.MsgBuy{
		Buyer: buyerAddr.String(),
		Orders: []*marketplace.MsgBuy_Order{
			{Selection: &marketplace.MsgBuy_Order_Selection{Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellOrderId}},
				Quantity: "11", BidPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})
	assert.ErrorContains(t, err, "cannot purchase 11 credits from a sell order that has 10 credits")

	// mismatchDenom
	wrongDenom := sdk.NewInt64Coin("ubar", 10)
	_, err = s.k.Buy(s.ctx, &marketplace.MsgBuy{
		Buyer: buyerAddr.String(),
		Orders: []*marketplace.MsgBuy_Order{
			{Selection: &marketplace.MsgBuy_Order_Selection{Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellOrderId}},
				Quantity: "10", BidPrice: &wrongDenom, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})
	assert.ErrorContains(t, err, "bid price denom does not match ask price denom")

	// bidding more than in the bank
	inBank := sdk.NewInt64Coin("ufoo", 10)
	biddingWith := sdk.NewInt64Coin("ufoo", 100)
	s.bankKeeper.EXPECT().GetBalance(gmAny, gmAny, gmAny).Return(inBank).Times(1)
	_, err = s.k.Buy(s.ctx, &marketplace.MsgBuy{
		Buyer: buyerAddr.String(),
		Orders: []*marketplace.MsgBuy_Order{
			{Selection: &marketplace.MsgBuy_Order_Selection{Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellOrderId}},
				Quantity: "10", BidPrice: &biddingWith, DisableAutoRetire: false, Expiration: &sellExp},
		},
	})
	assert.ErrorContains(t, err, sdkerrors.ErrInsufficientFunds.Error())
}
