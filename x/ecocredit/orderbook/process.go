package orderbook

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"

	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
)

func (o OrderBook) ProcessBatch(ctx context.Context) error {
	// for now do all processing synchronously, although in the future we can process
	// different markets in parallel
	marketIt, err := o.marketplaceStore.MarketStore().List(ctx, marketplacev1beta1.MarketIdIndexKey{})
	if err != nil {
		return err
	}

	for marketIt.Next() {
		market, err := marketIt.Value()
		if err != nil {
			return err
		}

		err = o.processMarket(ctx, market)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o OrderBook) processMarket(ctx context.Context, market *marketplacev1beta1.Market) error {
	buyOrderIterator, err := o.marketplaceStore.BuyOrderStore().List(ctx,
		marketplacev1beta1.BuyOrderMarketIdBidPriceU32IndexKey{}.WithMarketId(market.Id),
		ormlist.Reverse(),
	)
	if err != nil {
		return err
	}

	for buyOrderIterator.Next() {
		buyOrder, err := buyOrderIterator.Value()
		if err != nil {
			return err
		}

		switch selection := buyOrder.Selection.Sum.(type) {
		case *marketplacev1beta1.BuyOrder_Selection_SellOrderId:
			return o.processDirect(ctx, buyOrder, selection.SellOrderId)
		case *marketplacev1beta1.BuyOrder_Selection_Filter:
			panic("TODO")
		}

	}

	return nil
}

func (o OrderBook) processDirect(ctx context.Context, buyOrder *marketplacev1beta1.BuyOrder, sellOrderId uint64) error {
	sellOrder, err := o.marketplaceStore.SellOrderStore().Get(ctx, sellOrderId)
	if err != nil {
		return err
	}
	if sellOrder == nil {
		// TODO: delete buy order
		return ecocredit.ErrInvalidBuyOrder.Wrapf("can't find sell order %d", sellOrderId)
	}

	// TODO compare bid/ask price
	// settle

	return nil
}

func (o OrderBook) deleteBuyOrder(ctx context.Context, buyOrderId uint64) {
	//it, err := o.memStore.BuyOrderSellOrderMatchStore().List(ctx,
	//	orderbookv1beta1.BuyOrderSellOrderMatchBuyOrderIdSellOrderIdIndexKey{}.WithBuyOrderId(buyOrderId),
	//)
	//var toDelete orderbookv1beta1.BuyOrderSellOrderMatch
}
