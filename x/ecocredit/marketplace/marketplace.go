package marketplace

import (
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
)

type Marketplace struct {
	marketplaceStore marketplacev1beta1.StateStore
	ecocreditStore   ecocreditv1beta1.StateStore
	orderbook        *orderbook.orderbook
}

var _ marketplacev1beta1.MsgServer = Marketplace{}

func (receiver Marketplace) MsgBuy() {

}
