package testutil

import (
	"context"
	"fmt"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/orderbook"
)

type TestFillManager struct {
	log              *log.Logger
	marketplaceStore marketplacev1beta1.StateStore
	ecocreditStore   ecocreditv1beta1.StateStore
	bankBalances     map[string]sdk.Int
}

func (t TestFillManager) Fill(
	ctx context.Context,
	market *marketplacev1beta1.Market,
	buyOrder *marketplacev1beta1.BuyOrder,
	sellOrder *marketplacev1beta1.SellOrder,
) (orderbook.FillStatus, error) {
	buyQuant, err := math.NewPositiveDecFromString(buyOrder.Quantity)
	if err != nil {
		return 0, err
	}

	sellQuant, err := math.NewPositiveDecFromString(sellOrder.Quantity)
	if err != nil {
		return 0, err
	}

	settlementPrice, err := math.NewPositiveDecFromString(buyOrder.BidPrice)
	if err != nil {
		return 0, err
	}

	cmp := buyQuant.Cmp(sellQuant)

	var actualQuant math.Dec
	var status orderbook.FillStatus
	if cmp < 0 {
		actualQuant = buyQuant
		status = orderbook.BuyFilled

		newSellQuant, err := sellQuant.Sub(buyQuant)
		if err != nil {
			return 0, err
		}
		sellOrder.Quantity = newSellQuant.String()
		err = t.marketplaceStore.SellOrderStore().Update(ctx, sellOrder)
		if err != nil {
			return 0, err
		}

		err = t.marketplaceStore.BuyOrderStore().Delete(ctx, buyOrder)
		if err != nil {
			return 0, err
		}
	} else if cmp == 0 {
		actualQuant = buyQuant
		status = orderbook.BothFilled

		err = t.marketplaceStore.SellOrderStore().Delete(ctx, sellOrder)
		if err != nil {
			return 0, err
		}

		err = t.marketplaceStore.BuyOrderStore().Delete(ctx, buyOrder)
		if err != nil {
			return 0, err
		}
		return orderbook.BothFilled, nil
	} else {
		actualQuant = sellQuant
		status = orderbook.BothFilled

		err = t.marketplaceStore.SellOrderStore().Delete(ctx, sellOrder)
		if err != nil {
			return 0, err
		}

		newBuyQuant, err := buyQuant.Sub(sellQuant)
		if err != nil {
			return 0, err
		}
		buyOrder.Quantity = newBuyQuant.String()
		err = t.marketplaceStore.BuyOrderStore().Update(ctx, buyOrder)
		if err != nil {
			return 0, err
		}

		return orderbook.SellFilled, nil
	}

	// fill buy order 100%
	// discard remaining sell order matches
	// delete buy order from
	// 	buy order table
	//	buy order selector indexes
	var action string
	if buyOrder.DisableAutoRetire {
		if !sellOrder.DisableAutoRetire {
			return 0, fmt.Errorf("disable auto-retire failed")
		}
		action = "Transfer"
	} else {
		action = "Retire"
	}
	t.log.Printf("%s %s credits from batch %d from %s -> %s",
		action, actualQuant.String(), sellOrder.BatchId, sellOrder.Seller, buyOrder.Buyer)

	// TODO correct decimal precision
	payment, err := actualQuant.Mul(settlementPrice)
	if err != nil {
		return 0, err
	}

	t.log.Printf("Transfer %s %s from %s -> %s",
		payment.String(), market.BankDenom, buyOrder.Buyer, sellOrder.Seller)

	return status, nil
}

var _ orderbook.FillManager = &TestFillManager{}
