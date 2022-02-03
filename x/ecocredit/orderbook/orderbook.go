package orderbook

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/regen-network/regen-ledger/x/ecocredit/orderbook/fill"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
)

type orderbook struct {
	memStore         orderbookv1beta1.MemoryStore
	marketplaceStore marketplacev1beta1.StateStore
	ecocreditStore   ecocreditv1beta1.StateStore
	fillManager      fill.Manager
	logger           zerolog.Logger
}

func NewOrderBook(db ormdb.ModuleDB, fillManager fill.Manager, logger zerolog.Logger) (*orderbook, error) {
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
		fillManager:      fillManager,
		logger:           logger,
	}, nil
}

type OrderBook interface {
	// OnInsertBuyOrder
	OnInsertBuyOrder(ctx context.Context, buyOrder *marketplacev1beta1.BuyOrder) error

	// OnInsertSellOrder
	OnInsertSellOrder(ctx context.Context, sellOrder *marketplacev1beta1.SellOrder, batchInfo *ecocreditv1beta1.BatchInfo) error

	// ProcessBatch called in end blocker, can happen every block or at some epoch.
	ProcessBatch(ctx context.Context) error

	// Reload gets call on end blocker only when a node starts up.
	Reload(ctx context.Context) error
}
