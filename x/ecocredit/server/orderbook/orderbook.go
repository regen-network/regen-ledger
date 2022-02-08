package orderbook

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
)

type orderbook struct {
	memStore         orderbookv1beta1.MemoryStore
	marketplaceStore marketplacev1beta1.StateStore
	ecocreditStore   ecocreditv1beta1.StateStore
}

// NewOrderBook creates a new OrderBook instance.
func NewOrderBook(db ormdb.ModuleDB) (OrderBook, error) {
	memStore, err := orderbookv1beta1.NewMemoryStore(db)
	if err != nil {
		return nil, err
	}

	marketplaceStore, err := marketplacev1beta1.NewStateStore(db)
	if err != nil {
		return nil, err
	}

	ecocreditStore, err := ecocreditv1beta1.NewStateStore(db)
	if err != nil {
		return nil, err
	}

	return &orderbook{
		memStore:         memStore,
		marketplaceStore: marketplaceStore,
		ecocreditStore:   ecocreditStore,
	}, nil
}

type OrderBook interface {
	// OnInsertBuyOrder gets called whenever a buy order is inserted into the marketplace state.
	OnInsertBuyOrder(ctx context.Context, buyOrder *marketplacev1beta1.BuyOrder) error

	// OnInsertSellOrder gets called whenever a sell order is inserted into the marketplace state.
	OnInsertSellOrder(ctx context.Context, sellOrder *marketplacev1beta1.SellOrder, batchInfo *ecocreditv1beta1.BatchInfo) error

	// ProcessBatch called in end blocker, can happen every block or at some epoch.
	ProcessBatch(ctx context.Context) error

	// Reload gets call on end blocker only when a node starts up.
	Reload(ctx context.Context) error
}

func (o orderbook) OnInsertBuyOrder(ctx context.Context, buyOrder *marketplacev1beta1.BuyOrder) error {
	return nil
}

func (o orderbook) OnInsertSellOrder(ctx context.Context, sellOrder *marketplacev1beta1.SellOrder, batchInfo *ecocreditv1beta1.BatchInfo) error {
	return nil
}

func (o orderbook) ProcessBatch(ctx context.Context) error {
	return nil
}

func (o orderbook) Reload(ctx context.Context) error {
	return nil
}
