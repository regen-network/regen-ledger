package marketplace

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

func TestBuy_ValidTradable(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	ask := sdk.NewInt64Coin("ufoo", 10)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	marketTestSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	// make a sell order
	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).AnyTimes()
	sellExp := time.Now()
	res, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, DisableAutoRetire: true, Expiration: &sellExp},
		},
	})
	assert.NilError(t, err)
	sellOrderId := res.SellOrderIds[0]

	s.bankKeeper.EXPECT().HasBalance(any, any, any).Return(true).Times(1)
	s.bankKeeper.EXPECT().SendCoins(any, any, any, any).Return(nil).Times(1)

	_, err = s.k.Buy(s.ctx, &marketplace.MsgBuy{
		Buyer: buyerAddr.String(),
		Orders: []*marketplace.MsgBuy_Order{
			{Selection: &marketplace.MsgBuy_Order_Selection{Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellOrderId}},
				Quantity: "3", BidPrice: &ask, DisableAutoRetire: true, Expiration: &sellExp},
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
	assert.Check(t, tradableBalance.Equal(math.NewDecFromInt64(3)))
}

func TestBuy_ValidRetired(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, buyerAddr := testdata.KeyTestPubAddr()
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	ask := sdk.NewInt64Coin("ufoo", 10)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	marketTestSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	// make a sell order
	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
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

	s.bankKeeper.EXPECT().HasBalance(any, any, any).Return(true).Times(1)
	s.bankKeeper.EXPECT().SendCoins(any, any, any, any).Return(nil).Times(1)

	_, err = s.k.Buy(s.ctx, &marketplace.MsgBuy{
		Buyer: buyerAddr.String(),
		Orders: []*marketplace.MsgBuy_Order{
			{Selection: &marketplace.MsgBuy_Order_Selection{Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellOrderId}},
				Quantity: "3", BidPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
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
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	marketTestSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	// make a sell order
	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
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

	s.bankKeeper.EXPECT().HasBalance(any, any, any).Return(true).Times(1)
	s.bankKeeper.EXPECT().SendCoins(any, any, any, any).Return(nil).Times(1)

	_, err = s.k.Buy(s.ctx, &marketplace.MsgBuy{
		Buyer: buyerAddr.String(),
		Orders: []*marketplace.MsgBuy_Order{
			{Selection: &marketplace.MsgBuy_Order_Selection{Sum: &marketplace.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellOrderId}},
				Quantity: "10", BidPrice: &ask, DisableAutoRetire: false, Expiration: &sellExp},
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
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	marketTestSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)
	// make a sell order
	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
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

	s.bankKeeper.EXPECT().HasBalance(any, any, any).Return(true).Times(3)

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
	assert.ErrorContains(t, err, "auto-retire mismatch")

}
