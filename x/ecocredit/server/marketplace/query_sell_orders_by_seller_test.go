package marketplace

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestSellOrdersBySeller(t *testing.T) {
	t.Parallel()
	s := setupBase(t, 3)
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	otherAddr := s.addrs[1]
	noOrdersAddr := s.addrs[2]

	order1 := insertSellOrder(t, s, s.addrs[0], 1)
	order2 := insertSellOrder(t, s, otherAddr, 1)

	res, err := s.k.SellOrdersBySeller(s.ctx, &marketplace.QuerySellOrdersBySellerRequest{
		Seller:     s.addrs[0].String(),
		Pagination: &query.PageRequest{CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))
	assertOrderEqual(t, s.ctx, s.k, res.SellOrders[0], order1)
	assert.Equal(t, uint64(1), res.Pagination.Total)

	res, err = s.k.SellOrdersBySeller(s.ctx, &marketplace.QuerySellOrdersBySellerRequest{
		Seller:     otherAddr.String(),
		Pagination: &query.PageRequest{CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))
	assertOrderEqual(t, s.ctx, s.k, res.SellOrders[0], order2)
	assert.Equal(t, uint64(1), res.Pagination.Total)

	// addr with no sell orders should just return empty slice
	res, err = s.k.SellOrdersBySeller(s.ctx, &marketplace.QuerySellOrdersBySellerRequest{
		Seller:     noOrdersAddr.String(),
		Pagination: &query.PageRequest{CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.SellOrders))
	assert.Equal(t, uint64(0), res.Pagination.Total)

	// bad address should fail
	_, err = s.k.SellOrdersBySeller(s.ctx, &marketplace.QuerySellOrdersBySellerRequest{
		Seller:     "foobar1vlk23jrkl",
		Pagination: nil,
	})
	assert.ErrorContains(t, err, "decoding bech32 failed")
}
