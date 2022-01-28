package orderbook

import (
	"context"

	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
)

type OrderBook struct {
	memStore       orderbookv1beta1.MemoryStore
	marketStore    marketplacev1beta1.StateStore
	ecocreditStore ecocreditv1beta1.StateStore
}

func (o OrderBook) OnInsertBuyOrder(ctx context.Context, order *marketplacev1beta1.BuyOrder) error {

}

func (o OrderBook) OnInsertSellOrder(ctx context.Context, order *marketplacev1beta1.SellOrder) error {

}

func (o OrderBook) ProcessBatch(ctx context.Context) {

}
