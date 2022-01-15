package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"time"
)

// PruneOrders checks if there are any expired sell or buy orders and removes them from state.
func (s serverImpl) PruneOrders(ctx sdk.Context) error {
	blockTime := uint64(ctx.BlockTime().Add(time.Nanosecond).UnixNano())
	minTime := uint64(0)

	sellOrdersIter, err := s.sellOrderByExpirationIndex.PrefixScan(ctx, minTime, blockTime)
	if err != nil {
		return err
	}

	var sellOrders []*ecocredit.SellOrder
	_, err = orm.ReadAll(sellOrdersIter, &sellOrders)
	if err != nil {
		return err
	}

	for _, order := range sellOrders {
		err := s.sellOrderTable.Delete(ctx, order.OrderId)
		if err != nil {
			return err
		}
	}

	buyOrdersIter, err := s.buyOrderByExpirationIndex.PrefixScan(ctx, minTime, blockTime)
	if err != nil {
		return err
	}

	var buyOrders []*ecocredit.BuyOrder
	_, err = orm.ReadAll(buyOrdersIter, &buyOrders)
	if err != nil {
		return err
	}

	for _, order := range buyOrders {
		err := s.buyOrderTable.Delete(ctx, order.BuyOrderId)
		if err != nil {
			return err
		}
	}

	return nil
}
