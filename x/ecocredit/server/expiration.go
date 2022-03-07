package server

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// PruneOrders checks if there are any expired sell or buy orders and removes them from state.
func (s serverImpl) PruneOrders(ctx sdk.Context) error {
	blockTime := uint64(ctx.BlockTime().Add(time.Nanosecond).UnixNano())
	minTime := uint64(0)

	sellOrdersIter, err := s.sellOrderByExpirationIndex.PrefixScan(ctx, minTime, blockTime)
	if err != nil {
		return err
	}
	defer sellOrdersIter.Close()

	var sellOrder ecocredit.SellOrder
	for {
		_, err := sellOrdersIter.LoadNext(&sellOrder)
		if err != nil {
			if orm.ErrIteratorDone.Is(err) {
				break
			}
			return err
		}
		err = s.sellOrderTable.Delete(ctx, sellOrder.OrderId)
		if err != nil {
			return err
		}
	}

	buyOrdersIter, err := s.buyOrderByExpirationIndex.PrefixScan(ctx, minTime, blockTime)
	if err != nil {
		return err
	}
	defer buyOrdersIter.Close()

	var buyOrder ecocredit.BuyOrder
	for {
		_, err := buyOrdersIter.LoadNext(&buyOrder)
		if err != nil {
			if orm.ErrIteratorDone.Is(err) {
				break
			}
			return err
		}
		err = s.buyOrderTable.Delete(ctx, buyOrder.BuyOrderId)
		if err != nil {
			return err
		}
	}

	return nil
}
