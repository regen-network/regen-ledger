package orderbook

import (
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
)

type OrderBook struct {
	memStore         orderbookv1beta1.MemoryStore
	marketplaceStore marketplacev1beta1.StateStore
	ecocreditStore   ecocreditv1beta1.StateStore
}

func NewOrderBook(db ormdb.ModuleDB) (*OrderBook, error) {
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

	return &OrderBook{
		memStore:         memStore,
		marketplaceStore: marketplaceStore,
		ecocreditStore:   ecocreditStore,
	}, nil
}
