package marketplace

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestQueryBuyOrders(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	assert.NilError(t, s.marketStore.BuyOrderTable().Insert(s.ctx, &api.BuyOrder{Buyer: s.addr}))
	assert.NilError(t, s.marketStore.BuyOrderTable().Insert(s.ctx, &api.BuyOrder{Buyer: s.addr}))

	res, err := s.k.BuyOrders(s.ctx, &marketplace.QueryBuyOrdersRequest{})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.BuyOrders))

	res, err = s.k.BuyOrders(s.ctx, &marketplace.QueryBuyOrdersRequest{Pagination: &query.PageRequest{Limit: 1, CountTotal: true}})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.BuyOrders))
	assert.Equal(t, uint64(2), res.Pagination.Total)
}

func TestQueryBuyOrdersByAddress(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, otherAddr := testdata.KeyTestPubAddr()
	_, _, noOrdersAddr := testdata.KeyTestPubAddr()

	assert.NilError(t, s.marketStore.BuyOrderTable().Insert(s.ctx, &api.BuyOrder{Buyer: s.addr}))
	assert.NilError(t, s.marketStore.BuyOrderTable().Insert(s.ctx, &api.BuyOrder{Buyer: otherAddr}))

	// valid queries
	res, err := s.k.BuyOrdersByAddress(s.ctx, &marketplace.QueryBuyOrdersByAddressRequest{Address: s.addr.String()})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.BuyOrders))
	assert.Equal(t, s.addr.String(), sdk.AccAddress(res.BuyOrders[0].Buyer).String())

	res, err = s.k.BuyOrdersByAddress(s.ctx, &marketplace.QueryBuyOrdersByAddressRequest{Address: otherAddr.String()})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.BuyOrders))
	assert.Equal(t, otherAddr.String(), sdk.AccAddress(res.BuyOrders[0].Buyer).String())

	res, err = s.k.BuyOrdersByAddress(s.ctx, &marketplace.QueryBuyOrdersByAddressRequest{Address: otherAddr.String(), Pagination: &query.PageRequest{CountTotal: true}})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.BuyOrders))
	assert.Equal(t, otherAddr.String(), sdk.AccAddress(res.BuyOrders[0].Buyer).String())
	assert.Equal(t, uint64(1), res.Pagination.Total)

	// empty slice for addr with no orders
	res, err = s.k.BuyOrdersByAddress(s.ctx, &marketplace.QueryBuyOrdersByAddressRequest{Address: noOrdersAddr.String()})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.BuyOrders))

	// error on bad address
	res, err = s.k.BuyOrdersByAddress(s.ctx, &marketplace.QueryBuyOrdersByAddressRequest{Address: "foobarasdfxxlck"})
	assert.ErrorContains(t, err, "decoding bech32 failed")
}
