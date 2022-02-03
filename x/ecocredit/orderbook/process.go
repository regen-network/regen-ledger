package orderbook

import (
	"context"
	"fmt"

	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
)

func (o orderbook) ProcessBatch(ctx context.Context) error {
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

func (o orderbook) processMarket(ctx context.Context, market *marketplacev1beta1.Market) error {
	it, err := o.memStore.BuyOrderSellOrderMatchStore().
		List(ctx,
			orderbookv1beta1.BuyOrderSellOrderMatchMarketIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey{}.
				WithMarketId(market.Id),
		)

	if err != nil {
		return err
	}

	for it.Next() {
		match, err := it.Value()
		if err != nil {
			return err
		}

		buyOrder, err := o.marketplaceStore.BuyOrderStore().Get(ctx, match.BuyOrderId)
		if err != nil {
			return err
		}
		if buyOrder == nil {
			return fmt.Errorf("unexpected missing buy order")
		}

		sellOrder, err := o.marketplaceStore.SellOrderStore().Get(ctx, match.SellOrderId)
		if err != nil {
			return err
		}
		if sellOrder == nil {
			return fmt.Errorf("unexpected missing sell order")
		}

		status, err := o.fillManager.Fill(ctx, market, buyOrder, sellOrder)
		if err != nil {
			return err
		}
		switch status {
		case BuyFilled:
			err = o.deleteBuyOrder(ctx, buyOrder)
			if err != nil {
				return err
			}
		case SellFilled:
			err = o.deleteSellOrder(ctx, sellOrder)
			if err != nil {
				return err
			}
		case BothFilled:
			err = o.deleteBuyOrder(ctx, buyOrder)
			if err != nil {
				return err
			}
			err = o.deleteSellOrder(ctx, sellOrder)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected status %d", status)
		}
	}

	return nil
}

func (o orderbook) deleteBuyOrder(ctx context.Context, buyOrder *marketplacev1beta1.BuyOrder) error {
	err := o.memStore.BuyOrderSellOrderMatchStore().DeleteBy(ctx,
		orderbookv1beta1.BuyOrderSellOrderMatchBuyOrderIdSellOrderIdIndexKey{}.WithBuyOrderId(buyOrder.Id))
	if err != nil {
		return err
	}

	err = o.memStore.BuyOrderClassSelectorStore().DeleteBy(ctx,
		orderbookv1beta1.BuyOrderClassSelectorBuyOrderIdClassIdIndexKey{}.WithBuyOrderId(buyOrder.Id))
	if err != nil {
		return err
	}

	err = o.memStore.BuyOrderProjectSelectorStore().DeleteBy(ctx,
		orderbookv1beta1.BuyOrderProjectSelectorBuyOrderIdProjectIdIndexKey{}.WithBuyOrderId(buyOrder.Id))
	if err != nil {
		return err
	}

	return o.memStore.BuyOrderBatchSelectorStore().DeleteBy(ctx,
		orderbookv1beta1.BuyOrderBatchSelectorBuyOrderIdBatchIdIndexKey{}.WithBuyOrderId(buyOrder.Id))
}

func (o orderbook) deleteSellOrder(ctx context.Context, sellOrder *marketplacev1beta1.SellOrder) error {
	return o.memStore.BuyOrderSellOrderMatchStore().DeleteBy(ctx,
		orderbookv1beta1.BuyOrderSellOrderMatchSellOrderIdIndexKey{}.WithSellOrderId(sellOrder.Id))
}
