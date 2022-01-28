package orderbook

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

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

	market, err := o.marketplaceStore.MarketStore().Get(ctx, buyOrder.MarketId)
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
	sellOrder, err := o.marketplaceStore.SellOrderStore().Get(ctx, selection.SellOrderId)
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

	askPrice, ok := sdk.NewIntFromString(sellOrder.AskPrice)
	if !ok {
		return ecocredit.ErrInvalidInteger.Wrapf("ask price: %d", sellOrder.AskPrice)
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
	n := len(o.selection.Filter.Or)
	if n > MaxNumberSelectors {
		return ecocredit.ErrInvalidBuyOrder.Wrapf("too many selectors, got %d, the limit is %d", n, MaxNumberSelectors)
	}
	var numSelectors int
	for _, criteria := range o.selection.Filter.Or {
		numSelectors++
		if numSelectors > MaxNumberSelectors {
		}

		err := o.processSelector(criteria)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o buyOrderMatcher) processSelector(criteria *marketplacev1beta1.Filter_Criteria) error {
	switch selector := criteria.Selector.(type) {
	case *marketplacev1beta1.Filter_Criteria_ClassSelector:
		return o.matchByClassIdSelector(selector.ClassSelector)
	case *marketplacev1beta1.Filter_Criteria_ProjectSelector:
		return o.matchByProjectIdSelector(selector.ProjectSelector)
	case *marketplacev1beta1.Filter_Criteria_BatchSelector:
		return o.matchByBatchIdSelector(selector.BatchSelector.BatchId)
	default:
		return ecocredit.ErrInvalidBuyOrder.Wrapf("unknown selector type %s", selector)
	}
}

func (o buyOrderMatcher) matchByClassIdSelector(selector *marketplacev1beta1.ClassSelector) error {
	err := o.memStore.BuyOrderClassSelectorStore().Insert(o.ctx,
		&orderbookv1beta1.BuyOrderClassSelector{
			BuyOrderId:      o.buyOrder.Id,
			ClassId:         selector.ClassId,
			ProjectLocation: selector.ProjectLocation,
			MinStartDate:    selector.MinStartDate,
			MaxEndDate:      selector.MaxEndDate,
		},
	)
	if err != nil {
		return err
	}

	it, err := o.ecocreditStore.BatchInfoStore().List(
		o.ctx,
		ecocreditv1beta1.BatchInfoClassIdIndexKey{}.WithClassId(selector.ClassId),
	)
	if err != nil {
		return err
	}

	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return err
		}

		if !matchLocation(batch, selector.ProjectLocation) {
			continue
		}

		if !matchDates(batch, selector.MinStartDate, selector.MaxEndDate) {
			continue
		}

		err = o.onMatch(batch)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o buyOrderMatcher) matchByProjectIdSelector(selector *marketplacev1beta1.ProjectSelector) error {
	err := o.memStore.BuyOrderProjectSelectorStore().Insert(o.ctx,
		&orderbookv1beta1.BuyOrderProjectSelector{
			BuyOrderId:   o.buyOrder.Id,
			ProjectId:    selector.ProjectId,
			MinStartDate: selector.MinStartDate,
			MaxEndDate:   selector.MaxEndDate,
		},
	)
	if err != nil {
		return err
	}

	it, err := o.ecocreditStore.BatchInfoStore().List(
		o.ctx,
		ecocreditv1beta1.BatchInfoProjectIdIndexKey{}.WithProjectId(selector.ProjectId),
	)
	if err != nil {
		return err
	}

	for it.Next() {
		batch, err := it.Value()
		if err != nil {
			return err
		}

		if !matchDates(batch, selector.MinStartDate, selector.MaxEndDate) {
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

	if batch == nil {
		return ecocredit.ErrInvalidBuyOrder.Wrapf("batch %d not found", batchId)
	}

	err = o.memStore.BuyOrderBatchSelectorStore().Insert(o.ctx,
		&orderbookv1beta1.BuyOrderBatchSelector{
			BuyOrderId: o.buyOrder.Id,
			BatchId:    batchId,
		})
	if err != nil {
		return err
	}

	return o.onMatch(batch)
}

func (o buyOrderMatcher) onMatch(batch *ecocreditv1beta1.BatchInfo) error {
	it, err := o.marketplaceStore.SellOrderStore().List(o.ctx,
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

		askPrice, ok := sdk.NewIntFromString(sellOrder.AskPrice)
		if !ok {
			return ecocredit.ErrInvalidInteger.Wrapf("ask price: %d", sellOrder.AskPrice)
		}

		if o.bidPrice.LT(askPrice) {
			continue
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

func matchDates(batch *ecocreditv1beta1.BatchInfo, minStart, maxEnd *timestamppb.Timestamp) bool {
	if minStart != nil && batch.StartDate.AsTime().Before(minStart.AsTime()) {
		return false
	}

	if maxEnd != nil && batch.EndDate.AsTime().After(maxEnd.AsTime()) {
		return false
	}

	return true
}

func matchLocation(batch *ecocreditv1beta1.BatchInfo, filter string) bool {
	target := batch.ProjectLocation

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
