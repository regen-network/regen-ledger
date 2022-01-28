package orderbook

import (
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
)

type OrderBook struct {
	memStore       orderbookv1beta1.MemoryStore
	marketStore    marketplacev1beta1.StateStore
	ecocreditStore ecocreditv1beta1.StateStore
}
