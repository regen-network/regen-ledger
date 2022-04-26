package marketplace

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestQuery_SellOrder(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)

	// make a sell order
	order := api.SellOrder{
		Seller:            s.addr,
		BatchId:           1,
		Quantity:          "15.32",
		MarketId:          1,
		AskPrice:          "100",
		DisableAutoRetire: false,
		Expiration:        nil,
		Maker:             false,
	}
	id, err := s.marketStore.SellOrderTable().InsertReturningID(s.ctx, &order)
	assert.NilError(t, err)

	var gogoOrder marketplace.SellOrder
	assert.NilError(t, ormutil.PulsarToGogoSlow(&order, &gogoOrder))

	res, err := s.k.SellOrder(s.ctx, &marketplace.QuerySellOrderRequest{SellOrderId: id})
	assert.NilError(t, err)
	assert.Equal(t, s.addr.String(), res.SellOrder.Seller)
	assert.Equal(t, batchDenom, res.SellOrder.BatchDenom)
	assert.Equal(t, order.Quantity, res.SellOrder.Quantity)
	assert.Equal(t, ask.Denom, res.SellOrder.AskDenom)
	assert.Equal(t, order.AskPrice, res.SellOrder.AskPrice)
	assert.Equal(t, order.DisableAutoRetire, res.SellOrder.DisableAutoRetire)
	assert.DeepEqual(t, types.ProtobufToGogoTimestamp(order.Expiration), res.SellOrder.Expiration)

	// invalid order id should fail
	_, err = s.k.SellOrder(s.ctx, &marketplace.QuerySellOrderRequest{SellOrderId: 404})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
