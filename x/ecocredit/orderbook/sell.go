package orderbook

import (
	"context"

	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
)

func (o OrderBook) OnInsertSellOrder(ctx context.Context, order *marketplacev1beta1.SellOrder) error {

}
