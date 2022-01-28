package orderbook

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"

	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
)

type OrderBook struct {
	memStore       orderbookv1beta1.MemoryStore
	marketStore    marketplacev1beta1.StateStore
	ecocreditStore ecocreditv1beta1.StateStore
}

func (o OrderBook) OnInsertBuyOrder(ctx context.Context, buyOrder *marketplacev1beta1.BuyOrder) error {
	bidPrice, ok := sdk.NewIntFromString(buyOrder.BidPrice.Amount)
	if !ok {
		return ecocredit.ErrInvalidInteger.Wrapf("bid price: %d", buyOrder.BidPrice.Amount)
	}

	switch selection := buyOrder.Selection.Sum.(type) {
	case *marketplacev1beta1.BuyOrder_Selection_SellOrderId:
		sellOrder, err := o.marketStore.SellOrderStore().Get(ctx, selection.SellOrderId)
		if err != nil {
			return err
		}

		if sellOrder == nil {
			return ecocredit.ErrNotFound.Wrapf("sell order %d", selection.SellOrderId)
		}

		askPrice, ok := sdk.NewIntFromString(sellOrder.AskPrice.Amount)
		if !ok {
			return ecocredit.ErrInvalidInteger.Wrapf("ask price: %d", sellOrder.AskPrice.Amount)
		}

		if bidPrice.LT(askPrice) {
			return ecocredit.ErrInvalidBuyOrder.Wrapf("bid price %d is below ask price %d")
		}

		batch, err := o.ecocreditStore.BatchInfoStore().GetByBatchDenom(ctx, sellOrder.BatchDenom)
		if err != nil {
			return err
		}

		if batch == nil {
			return ecocredit.ErrNotFound.Wrapf("batch %s", sellOrder.BatchDenom)
		}

		market, err := o.marketStore.MarketStore().GetByCreditTypeBankDenom(ctx)

		o.memStore.BuyOrderSellOrderMatchStore().Insert(ctx)
	case *marketplacev1beta1.BuyOrder_Selection_Filter:
	}
}

func (o OrderBook) OnInsertSellOrder(ctx context.Context, order *marketplacev1beta1.SellOrder) error {

}

func (o OrderBook) ProcessBatch(ctx context.Context) {

}
