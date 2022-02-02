package orderbook

import (
	"context"
	"fmt"

	"github.com/regen-network/regen-ledger/types/math"

	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
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

		sellOrder, err := o.marketplaceStore.BuyOrderStore().Get(ctx, match.SellOrderId)
		if err != nil {
			return err
		}
		if sellOrder == nil {
			return fmt.Errorf("unexpected missing sell order")
		}

		buyQuant, err := math.NewPositiveDecFromString(buyOrder.Quantity)
		if err != nil {
			return err
		}

		sellQuant, err := math.NewPositiveDecFromString(sellOrder.Quantity)
		if err != nil {
			return err
		}

		cmp := buyQuant.Cmp(sellQuant)
		if cmp < 0 {
			// fill buy order 100%
			// discard remaining sell order matches
			// delete buy order from
			// 	buy order table
			//	buy order selector indexes
		} else if cmp == 0 {

		} else {

		}
	}

	return nil
}

func (o OrderBook) deleteBuyOrder(ctx context.Context, buyOrderId uint64) error {
	it, err := o.memStore.BuyOrderSellOrderMatchStore().List(ctx,
		orderbookv1beta1.BuyOrderSellOrderMatchBuyOrderIdSellOrderIdIndexKey{}.WithBuyOrderId(buyOrderId),
	)
	if err != nil {
		return err
	}

	var toDelete []*orderbookv1beta1.BuyOrderSellOrderMatch
	for it.Next() {
		match, err := it.Value()
		if err != nil {
			return err
		}
		toDelete = append(toDelete, match)
	}
	it.Close()

	for _, match := range toDelete {
		err := o.memStore.BuyOrderSellOrderMatchStore().Delete(ctx, match)
		if err != nil {
			return err
		}
	}

	return nil
}
