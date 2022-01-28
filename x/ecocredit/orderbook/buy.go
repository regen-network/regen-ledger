package orderbook

import (
	"context"
	"fmt"

	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (o OrderBook) OnInsertBuyOrder(ctx context.Context, buyOrder *marketplacev1beta1.BuyOrder) error {
	bidPrice, ok := sdk.NewIntFromString(buyOrder.BidPrice)
	if !ok {
		return ecocredit.ErrInvalidInteger.Wrapf("bid price: %d", buyOrder.BidPrice)
	}

	market, err := o.marketStore.MarketStore().Get(ctx, buyOrder.MarketId)
	if err != nil {
		return err
	}
	if market == nil {
		return ecocredit.ErrInvalidBuyOrder.Wrapf("market %d not found", buyOrder.MarketId)
	}

	bidPriceU64, err := IntPriceToUInt64(bidPrice, market.PrecisionModifier)
	if err != nil {
		return err
	}

	switch selection := buyOrder.Selection.Sum.(type) {
	case *marketplacev1beta1.BuyOrder_Selection_SellOrderId:
		return o.insertBuyOrderDirect(ctx, buyOrder, market, selection, bidPrice, bidPriceU64)
	case *marketplacev1beta1.BuyOrder_Selection_Filter:
		matcher := &buyOrderMatcher{
			OrderBook:   o,
			ctx:         ctx,
			buyOrder:    buyOrder,
			market:      market,
			selection:   selection,
			bidPrice:    bidPrice,
			bidPriceU64: bidPriceU64,
		}
		return matcher.insertBuyOrderFiltered()
	default:
		return fmt.Errorf("unexpected")
	}
}

func (o OrderBook) insertBuyOrderDirect(
	ctx context.Context,
	buyOrder *marketplacev1beta1.BuyOrder,
	market *marketplacev1beta1.Market,
	selection *marketplacev1beta1.BuyOrder_Selection_SellOrderId,
	bidPrice sdk.Int,
	bidPriceU64 uint64,
) error {
	sellOrder, err := o.marketStore.SellOrderStore().Get(ctx, selection.SellOrderId)
	if err != nil {
		return err
	}

	if sellOrder == nil {
		return ecocredit.ErrNotFound.Wrapf("sell order %d", selection.SellOrderId)
	}

	if sellOrder.MarketId != buyOrder.MarketId {
		return ecocredit.ErrInvalidBuyOrder.Wrapf("buy order market id %d does not match sell order market id %d",
			buyOrder.MarketId, sellOrder.MarketId)
	}

	askPrice, ok := sdk.NewIntFromString(sellOrder.AskPrice.Amount)
	if !ok {
		return ecocredit.ErrInvalidInteger.Wrapf("ask price: %d", sellOrder.AskPrice.Amount)
	}

	if bidPrice.LT(askPrice) {
		return ecocredit.ErrInvalidBuyOrder.Wrapf("bid price %d is below ask price %d")
	}

	askPriceU64, err := IntPriceToUInt64(askPrice, market.PrecisionModifier)
	if err != nil {
		return err
	}

	return o.memStore.BuyOrderSellOrderMatchStore().Insert(ctx, &orderbookv1beta1.BuyOrderSellOrderMatch{
		MarketId:           buyOrder.MarketId,
		BuyOrderId:         buyOrder.Id,
		SellOrderId:        sellOrder.Id,
		BidPriceComplement: ^bidPriceU64,
		AskPrice:           askPriceU64,
	})
}

const MaxNumberSelectors = 4

type buyOrderMatcher struct {
	OrderBook
	ctx         context.Context
	buyOrder    *marketplacev1beta1.BuyOrder
	market      *marketplacev1beta1.Market
	selection   *marketplacev1beta1.BuyOrder_Selection_Filter
	bidPrice    sdk.Int
	bidPriceU64 uint64
}

func (o buyOrderMatcher) insertBuyOrderFiltered() error {
	var numSelectors int
	for _, criteria := range o.selection.Filter.Or {
		for _, selector := range criteria.Or {
			numSelectors++
			if numSelectors > MaxNumberSelectors {
				return ecocredit.ErrInvalidBuyOrder.Wrapf("too many selectors, the limit is %d", MaxNumberSelectors)
			}

			err := o.processSelector(criteria, selector)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (o buyOrderMatcher) processSelector(
	criteria *marketplacev1beta1.Filter_Criteria,
	selector *marketplacev1beta1.Selector,
) error {
	selectorType := selector.SelectorType
	switch selectorType {
	case marketplacev1beta1.SelectorType_SELECTOR_TYPE_CLASS,
		marketplacev1beta1.SelectorType_SELECTOR_TYPE_PROJECT,
		marketplacev1beta1.SelectorType_SELECTOR_TYPE_BATCH:
		value, ok := selector.Value.(*marketplacev1beta1.Selector_Uint64Value)
		if !ok {
			return ecocredit.ErrInvalidBuyOrder.Wrapf("expected uint64_value for %s", selectorType)
		}
		uint64Value := value.Uint64Value
		err := o.memStore.UInt64SelectorBuyOrderStore().Insert(o.ctx, &orderbookv1beta1.UInt64SelectorBuyOrder{
			BuyOrderId:      o.buyOrder.Id,
			SelectorType:    selectorType,
			Value:           uint64Value,
			ProjectLocation: criteria.ProjectLocation,
			MinStartDate:    criteria.MinStartDate,
			MaxEndDate:      criteria.MaxEndDate,
		})
		if err != nil {
			return err
		}

		return o.matchByUInt64Selector(criteria, selector, uint64Value)
	default:
		return ecocredit.ErrInvalidBuyOrder.Wrapf("unknown selector type %s", selectorType)
	}
}

func (o buyOrderMatcher) onMatch(batch *ecocreditv1beta1.BatchInfo) error {
	it, err := o.marketStore.SellOrderStore().List(o.ctx,
		marketplacev1beta1.SellOrderBatchDenomIndexKey{}.WithBatchDenom(batch.BatchDenom))
	if err != nil {
		return err
	}

	for it.Next() {
		sellOrder, err := it.Value()
		if err != nil {
			return err
		}

		if sellOrder.MarketId != o.buyOrder.MarketId {
			return ecocredit.ErrInvalidBuyOrder.Wrapf("buy order market id %d does not match matching sell order market id %d",
				o.buyOrder.MarketId, sellOrder.MarketId)
		}

		askPrice, ok := sdk.NewIntFromString(sellOrder.AskPrice.Amount)
		if !ok {
			return ecocredit.ErrInvalidInteger.Wrapf("ask price: %d", sellOrder.AskPrice.Amount)
		}

		if o.bidPrice.LT(askPrice) {
			return ecocredit.ErrInvalidBuyOrder.Wrapf("bid price %d is below ask price %d")
		}

		askPriceU64, err := IntPriceToUInt64(askPrice, o.market.PrecisionModifier)
		if err != nil {
			return err
		}

		err = o.memStore.BuyOrderSellOrderMatchStore().Insert(o.ctx, &orderbookv1beta1.BuyOrderSellOrderMatch{
			MarketId:           o.buyOrder.MarketId,
			BuyOrderId:         o.buyOrder.Id,
			SellOrderId:        sellOrder.Id,
			BidPriceComplement: ^o.bidPriceU64,
			AskPrice:           askPriceU64,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (o buyOrderMatcher) matchByUInt64Selector(criteria *marketplacev1beta1.Filter_Criteria, selector *marketplacev1beta1.Selector, value uint64) error {
	switch selector.SelectorType {
	case marketplacev1beta1.SelectorType_SELECTOR_TYPE_CLASS:
		return o.matchByClassIdSelector(criteria, value)
	case marketplacev1beta1.SelectorType_SELECTOR_TYPE_PROJECT:
		return o.matchByProjectIdSelector(criteria, value)
	default:
		panic("TODO")
	}
}

func (o buyOrderMatcher) matchByClassIdSelector(criteria *marketplacev1beta1.Filter_Criteria, classId uint64) error {
	it, err := o.ecocreditStore.BatchInfoStore().List(
		o.ctx,
		ecocreditv1beta1.BatchInfoClassIdIndexKey{}.WithClassId(classId),
	)
	if err != nil {
		return err
	}

	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return err
		}

		if !matchLocation(batch, criteria) {
			continue
		}

		if !matchDates(batch, criteria) {
			continue
		}

		err = o.onMatch(batch)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o buyOrderMatcher) matchByProjectIdSelector(criteria *marketplacev1beta1.Filter_Criteria, projectId uint64) error {
	it, err := o.ecocreditStore.BatchInfoStore().List(
		o.ctx,
		ecocreditv1beta1.BatchInfoProjectIdIndexKey{}.WithProjectId(projectId),
	)
	if err != nil {
		return err
	}

	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return err
		}

		if criteria.ProjectLocation != batch.ProjectLocation {
			return ecocredit.ErrInvalidBuyOrder.Wrapf("project ID %d selected but criteria location %s != %s",
				projectId, criteria.ProjectLocation, batch.ProjectId,
			)
		}

		if !matchDates(batch, criteria) {
			continue
		}

		err = o.onMatch(batch)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o buyOrderMatcher) matchByBatchIdSelector(batchId uint64) error {
	batch, err := o.ecocreditStore.BatchInfoStore().Get(o.ctx, batchId)
	if err != nil {
		return err
	}

	panic("TODO")
}

func matchDates(batch *ecocreditv1beta1.BatchInfo, criteria *marketplacev1beta1.Filter_Criteria) bool {
	if criteria.MinStartDate != nil && batch.StartDate.AsTime().Before(criteria.MinStartDate.AsTime()) {
		return false
	}

	if criteria.MaxEndDate != nil && batch.EndDate.AsTime().After(criteria.MaxEndDate.AsTime()) {
		return false
	}

	return true
}

func matchLocation(batch *ecocreditv1beta1.BatchInfo, criteria *marketplacev1beta1.Filter_Criteria) bool {
	target := batch.ProjectLocation
	filter := criteria.ProjectLocation

	n := len(filter)

	// if the target is shorter than the filter than we know we don't have a match
	if len(target) < n {
		return false
	}

	// if the filter length is less than 2 we should match anything (the only filters less than 2 should be totally empty)
	if n < 2 {
		return true
	}

	// check if country matches
	if target[:2] != filter[:2] {
		return false
	}

	panic("TODO")
}
