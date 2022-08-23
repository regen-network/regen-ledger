package orderbook

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"

	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	orderbookv1alpha1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1alpha1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
)

// TODO: revisit when BuyOrder is reintroduced for marketplace order-book https://github.com/regen-network/regen-ledger/issues/505

type orderbook struct {
	memStore         orderbookv1alpha1.MemoryStore
	marketplaceStore marketplacev1.StateStore
	ecocreditStore   ecocreditv1.StateStore
}

// NewOrderBook creates a new OrderBook instance.
func NewOrderBook(db ormdb.ModuleDB) (OrderBook, error) {
	memStore, err := orderbookv1alpha1.NewMemoryStore(db)
	if err != nil {
		return nil, err
	}

	marketplaceStore, err := marketplacev1.NewStateStore(db)
	if err != nil {
		return nil, err
	}

	ecocreditStore, err := ecocreditv1.NewStateStore(db)
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
	// OnInsertBuyOrder(ctx context.Context, buyOrder *marketplacev1.BuyOrder) error

	// OnInsertSellOrder gets called whenever a sell order is inserted into the marketplace state.
	OnInsertSellOrder(ctx context.Context, sellOrder *marketplacev1.SellOrder, batch *ecocreditv1.Batch) error

	// ProcessBatch called in end blocker, can happen every block or at some epoch.
	ProcessBatch(ctx context.Context) error

	// Reload gets call on end blocker only when a node starts up.
	Reload(ctx context.Context) error
}

/*
func (o orderbook) OnInsertBuyOrder(ctx context.Context, buyOrder *marketplacev1.BuyOrder) error {
	return nil
}
*/

func (o orderbook) OnInsertSellOrder(ctx context.Context, sellOrder *marketplacev1.SellOrder, batch *ecocreditv1.Batch) error {
	return nil
}

func (o orderbook) ProcessBatch(ctx context.Context) error {
	return nil
}

func (o orderbook) Reload(ctx context.Context) error {
	return nil
}
