package orderbook

import (
	"context"

	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
)

func (o OrderBook) Reload(ctx context.Context) error {
	it, err := o.marketplaceStore.BuyOrderStore().List(ctx, marketplacev1beta1.BuyOrderPrimaryKey{})
	if err != nil {
		return err
	}

	for it.Next() {
		buyOrder, err := it.Value()
		if err != nil {
			return err
		}

		err = o.OnInsertBuyOrder(ctx, buyOrder)
		if err != nil {
			return err
		}
	}

	return nil
}
