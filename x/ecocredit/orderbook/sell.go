package orderbook

import (
	"context"

	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"

	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"

	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
)

func (o OrderBook) OnInsertSellOrder(ctx context.Context, sellOrder *marketplacev1beta1.SellOrder, batchInfo *ecocreditv1beta1.BatchInfo) error {
	askPrice, ok := sdk.NewIntFromString(sellOrder.AskPrice)
	if !ok {
		return ecocredit.ErrInvalidInteger.Wrapf("ask price: %d", sellOrder.AskPrice)
	}

	market, err := o.marketplaceStore.MarketStore().Get(ctx, sellOrder.MarketId)
	if err != nil {
		return err
	}
	if market == nil {
		return ecocredit.ErrInvalidSellOrder.Wrapf("market %d not found", sellOrder.MarketId)
	}

	askPriceU64, err := IntPriceToUInt64(askPrice, market.PrecisionModifier)
	if err != nil {
		return err
	}

	matcher := &sellOrderMatcher{
		OrderBook:   OrderBook{},
		ctx:         ctx,
		sellOrder:   sellOrder,
		market:      market,
		askPrice:    askPrice,
		askPriceU64: askPriceU64,
		batchInfo:   batchInfo,
	}

	return matcher.match()
}

type sellOrderMatcher struct {
	OrderBook
	ctx         context.Context
	sellOrder   *marketplacev1beta1.SellOrder
	market      *marketplacev1beta1.Market
	askPrice    sdk.Int
	askPriceU64 uint64
	batchInfo   *ecocreditv1beta1.BatchInfo
}

func (m sellOrderMatcher) match() error {
	err := m.matchCreditClass()
	if err != nil {
		return err
	}

	err = m.matchProject()
	if err != nil {
		return err
	}

	return m.matchBatch()
}

func (m sellOrderMatcher) matchCreditClass() error {
	it, err := m.memStore.BuyOrderClassSelectorStore().List(m.ctx,
		orderbookv1beta1.BuyOrderClassSelectorClassIdIndexKey{}.WithClassId(m.batchInfo.ClassId),
	)
	if err != nil {
		return err
	}

	for it.Next() {
		selector, err := it.Value()
		if err != nil {
			return err
		}

		if !matchLocation(m.batchInfo, selector.ProjectLocation) {
			continue
		}

		if !matchDates(m.batchInfo, selector.MinStartDate, selector.MaxEndDate) {
			continue
		}

		err = m.onMatch(selector.BuyOrderId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m sellOrderMatcher) matchProject() error {
	it, err := m.memStore.BuyOrderProjectSelectorStore().List(m.ctx,
		orderbookv1beta1.BuyOrderProjectSelectorProjectIdIndexKey{}.WithProjectId(m.batchInfo.ProjectId),
	)
	if err != nil {
		return err
	}

	for it.Next() {
		selector, err := it.Value()
		if err != nil {
			return err
		}

		if !matchDates(m.batchInfo, selector.MinStartDate, selector.MaxEndDate) {
			continue
		}

		err = m.onMatch(selector.BuyOrderId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m sellOrderMatcher) matchBatch() error {
	it, err := m.memStore.BuyOrderBatchSelectorStore().List(m.ctx,
		orderbookv1beta1.BuyOrderBatchSelectorBatchIdIndexKey{}.WithBatchId(m.batchInfo.Id),
	)
	if err != nil {
		return err
	}

	for it.Next() {
		selector, err := it.Value()
		if err != nil {
			return err
		}

		err = m.onMatch(selector.BuyOrderId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m sellOrderMatcher) onMatch(buyOrderId uint64) error {
	buyOrder, err := m.marketplaceStore.BuyOrderStore().Get(m.ctx, buyOrderId)
	if err != nil {
		return err
	}

	marketId := m.sellOrder.MarketId
	if buyOrder.MarketId != marketId {
		return ecocredit.ErrInvalidSellOrder.Wrapf("buy order market id %d does not match matching sell order market id %d",
			buyOrder.MarketId, marketId)
	}

	bidPrice, ok := sdk.NewIntFromString(buyOrder.BidPrice)
	if !ok {
		return ecocredit.ErrInvalidInteger.Wrapf("bid price: %d", buyOrder.BidPrice)
	}

	if m.askPrice.GT(bidPrice) {
		return nil
	}

	bidPriceU64, err := IntPriceToUInt64(bidPrice, m.market.PrecisionModifier)
	if err != nil {
		return err
	}

	return m.memStore.BuyOrderSellOrderMatchStore().Insert(m.ctx,
		&orderbookv1beta1.BuyOrderSellOrderMatch{
			MarketId:           marketId,
			BuyOrderId:         buyOrderId,
			SellOrderId:        m.sellOrder.Id,
			BidPriceComplement: ^bidPriceU64,
			AskPrice:           m.askPriceU64,
		},
	)
}
