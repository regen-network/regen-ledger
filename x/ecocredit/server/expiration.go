package server

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// PruneOrders checks if there are any expired sell or buy orders and removes them from state.
func (s serverImpl) PruneOrders(ctx sdk.Context) error {

	blockTime := ctx.BlockTime().String()
	minTime := time.Time{}.String() // 0001-01-01 00:00:00 +0000 UTC

	sellOrdersIter, err := s.sellOrderByExpirationIndex.PrefixScan(ctx, minTime, blockTime)
	if err != nil {
		panic(err)
	}

	var sellOrders []*ecocredit.SellOrder
	_, err = orm.ReadAll(sellOrdersIter, &sellOrders)
	if err != nil {
		panic(err)
	}

	for _, order := range sellOrders {
		err := s.buyOrderTable.Delete(ctx, order.OrderId)
		if err != nil {
			panic(err)
		}
	}

	buyOrdersIter, err := s.buyOrderByExpirationIndex.PrefixScan(ctx, minTime, blockTime)
	if err != nil {
		panic(err)
	}

	var buyOrders []*ecocredit.BuyOrder
	_, err = orm.ReadAll(buyOrdersIter, &buyOrders)
	if err != nil {
		panic(err)
	}

	for _, order := range buyOrders {
		err := s.buyOrderTable.Delete(ctx, order.BuyOrderId)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
