package fill

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/types/math"
)

type Manager interface {
	Fill(ctx context.Context, market *marketplacev1beta1.Market, buyOrder *marketplacev1beta1.BuyOrder, sellOrder *marketplacev1beta1.SellOrder) (Status, error)
}

type Status int

const (
	NotFilled Status = iota
	BothFilled
	BuyFilled
	SellFilled
)

func (s Status) String() string {
	switch s {
	case BothFilled:
		return "BothFilled"
	case BuyFilled:
		return "BuyFilled"
	case SellFilled:
		return "SellFilled"
	default:
		return "undefined"
	}
}

type manager struct {
	marketplaceStore marketplacev1beta1.StateStore
	ecocreditStore   ecocreditv1beta1.StateStore
	bankBalances     map[string]sdk.Int
	transferManager  TransferManager
	logger           zerolog.Logger
}

func NewManager(db ormdb.ModuleDB, transferManager TransferManager, logger zerolog.Logger) (Manager, error) {
	mgr := &manager{transferManager: transferManager, logger: logger}

	var err error
	mgr.marketplaceStore, err = marketplacev1beta1.NewStateStore(db)
	if err != nil {
		return nil, err
	}

	mgr.ecocreditStore, err = ecocreditv1beta1.NewStateStore(db)
	if err != nil {
		return nil, err
	}

	return mgr, nil
}

type TransferManager interface {
	SendCoinsTo(denom string, amount sdk.Int, from, to sdk.AccAddress) error
	SendCreditsTo(batchId uint64, amount math.Dec, from, to sdk.AccAddress, retire bool) error
}

func (t manager) Fill(
	ctx context.Context,
	market *marketplacev1beta1.Market,
	buyOrder *marketplacev1beta1.BuyOrder,
	sellOrder *marketplacev1beta1.SellOrder,
) (Status, error) {
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
	var status Status
	if cmp < 0 {
		actualQuant = buyQuant
		status = BuyFilled

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
		status = BothFilled

		err = t.marketplaceStore.SellOrderStore().Delete(ctx, sellOrder)
		if err != nil {
			return 0, err
		}

		err = t.marketplaceStore.BuyOrderStore().Delete(ctx, buyOrder)
		if err != nil {
			return 0, err
		}
	} else {
		actualQuant = sellQuant
		status = BothFilled

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
	}

	retire := true
	if buyOrder.DisableAutoRetire {
		if !sellOrder.DisableAutoRetire {
			return 0, fmt.Errorf("disable auto-retire failed")
		}
		retire = false
	}
	err = t.transferManager.SendCreditsTo(sellOrder.BatchId, actualQuant, sellOrder.Seller, buyOrder.Buyer, retire)

	// TODO correct decimal precision
	payment, err := actualQuant.Mul(settlementPrice)
	if err != nil {
		return 0, err
	}

	paymentInt, err := payment.Int64()
	if err != nil {
		return 0, err
	}

	err = t.transferManager.SendCoinsTo(market.BankDenom, sdk.NewInt(paymentInt), buyOrder.Buyer, sellOrder.Seller)
	if err != nil {
		return 0, err
	}

	return status, nil
}

var _ Manager = &manager{}
