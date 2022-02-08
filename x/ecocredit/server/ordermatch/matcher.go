package ordermatch

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
)

type matcher struct {
	memStore         orderbookv1beta1.MemoryStore
	marketplaceStore marketplacev1beta1.StateStore
	ecocreditStore   ecocreditv1beta1.StateStore
}

// NewMatcher creates a new Matcher instance.
func NewMatcher(db ormdb.ModuleDB) (Matcher, error) {
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

	return &matcher{
		memStore:         memStore,
		marketplaceStore: marketplaceStore,
		ecocreditStore:   ecocreditStore,
	}, nil
}

type Matcher interface {
	// OnInsertBuyOrder gets called whenever a buy order is inserted into the marketplace state.
	OnInsertBuyOrder(ctx context.Context, buyOrder *marketplacev1beta1.BuyOrder) error

	// OnInsertSellOrder gets called whenever a sell order is inserted into the marketplace state.
	OnInsertSellOrder(ctx context.Context, sellOrder *marketplacev1beta1.SellOrder, batchInfo *ecocreditv1beta1.BatchInfo) error

	// Reload gets call on end blocker only when a node starts up.
	Reload(ctx context.Context) error
}

func (o matcher) OnInsertSellOrder(context.Context, *marketplacev1beta1.SellOrder, *ecocreditv1beta1.BatchInfo) error {
	return fmt.Errorf("TODO")
}
