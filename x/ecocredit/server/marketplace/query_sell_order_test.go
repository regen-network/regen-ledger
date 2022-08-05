package marketplace

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestQuery_SellOrder(t *testing.T) {
	t.Parallel()
	s := setupBase(t, 1)
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)

	expiration, err := types.ParseDate("expiration", "2030-01-01")
	require.NoError(s.t, err)

	// make a sell order (with expiration)
	order1 := api.SellOrder{
		Seller:     s.addrs[0],
		BatchKey:   1,
		Quantity:   "15.32",
		MarketId:   1,
		AskAmount:  "100",
		Expiration: timestamppb.New(expiration),
	}
	id1, err := s.marketStore.SellOrderTable().InsertReturningID(s.ctx, &order1)
	require.NoError(t, err)

	// make a sell order (no expiration)
	order2 := api.SellOrder{
		Seller:    s.addrs[0],
		BatchKey:  1,
		Quantity:  "15.32",
		MarketId:  1,
		AskAmount: "100",
	}
	id2, err := s.marketStore.SellOrderTable().InsertReturningID(s.ctx, &order2)
	require.NoError(t, err)

	var gogoOrder1 marketplace.SellOrder
	require.NoError(t, ormutil.PulsarToGogoSlow(&order1, &gogoOrder1))

	var gogoOrder2 marketplace.SellOrder
	require.NoError(t, ormutil.PulsarToGogoSlow(&order2, &gogoOrder2))

	res1, err := s.k.SellOrder(s.ctx, &marketplace.QuerySellOrderRequest{SellOrderId: id1})
	require.NoError(t, err)
	require.Equal(t, s.addrs[0].String(), res1.SellOrder.Seller)
	require.Equal(t, batchDenom, res1.SellOrder.BatchDenom)
	require.Equal(t, order1.Quantity, res1.SellOrder.Quantity)
	require.Equal(t, ask.Denom, res1.SellOrder.AskDenom)
	require.Equal(t, order1.AskAmount, res1.SellOrder.AskAmount)
	require.Equal(t, order1.DisableAutoRetire, res1.SellOrder.DisableAutoRetire)
	require.Equal(t, types.ProtobufToGogoTimestamp(order1.Expiration), res1.SellOrder.Expiration)

	res2, err := s.k.SellOrder(s.ctx, &marketplace.QuerySellOrderRequest{SellOrderId: id2})
	require.NoError(t, err)
	require.True(t, res2.SellOrder.Expiration.Equal(nil))
	require.Equal(t, types.ProtobufToGogoTimestamp(order2.Expiration), res2.SellOrder.Expiration)

	// invalid order id should fail
	_, err = s.k.SellOrder(s.ctx, &marketplace.QuerySellOrderRequest{SellOrderId: 404})
	require.ErrorContains(t, err, ormerrors.NotFound.Error())
}
