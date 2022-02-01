package orderbook

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"

	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"

	"github.com/regen-network/regen-ledger/x/ecocredit"

	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
)

func (o OrderBook) OnInsertSellOrder(ctx context.Context, sellOrder *marketplacev1beta1.SellOrder, batchInfo *ecocreditv1beta1.BatchInfo) error {
	batchInfo, err := o.ecocreditStore.BatchInfoStore().GetByBatchDenom(ctx, sellOrder.BatchDenom)
	if err != nil {
		return err
	}

	if batchInfo == nil {
		return ecocredit.ErrInvalidSellOrder.Wrapf("can't find batch %s", sellOrder.BatchDenom)
	}

	project, err := o.ecocreditStore.ProjectInfoStore().Get(ctx, batchInfo.ProjectId)
	if err != nil {
		return err
	}

	if project == nil {
		return ecocredit.ErrInvalidSellOrder.Wrapf("can't find project %s", batchInfo.ProjectId)
	}

	startDays, err := timestampToDays(batchInfo.StartDate)
	if err != nil {
		return err
	}

	endDays, err := timestampToDays(batchInfo.EndDate)
	if err != nil {
		return err
	}

	return o.memStore.SellOrderStore().Insert(ctx,
		&orderbookv1beta1.SellOrder{
			SellOrderId: sellOrder.Id,
			MarketId:    sellOrder.MarketId,
			AskPriceU32: sellOrder.AskPriceU32,
			ClassId:     project.ClassId,
			ProjectId:   batchInfo.ProjectId,
			BatchId:     batchInfo.Id,
			Location:    project.ProjectLocation,
			StartDate:   startDays,
			EndDate:     endDays,
		},
	)
}

func timestampToDays(timestamp *timestamppb.Timestamp) (uint32, error) {
	panic("TODO")
}
