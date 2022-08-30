package keeper

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types/query"

	ecocreditApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

func TestSellOrdersByBatch(t *testing.T) {
	t.Parallel()
	s := setupBase(t, 2)
	s.testSellSetup(batchDenom, ask.Denom, ask.Denom[1:], classID, start, end, creditType)

	// make another batch
	otherDenom := "C01-19990101-20290101-001"
	assert.NilError(t, s.coreStore.BatchTable().Insert(s.ctx, &ecocreditApi.Batch{
		ProjectKey: 1,
		Denom:      otherDenom,
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))

	order1 := insertSellOrder(t, s, s.addrs[0], 1)
	order2 := insertSellOrder(t, s, s.addrs[0], 2)

	// query the first denom
	res, err := s.k.SellOrdersByBatch(s.ctx, &types.QuerySellOrdersByBatchRequest{
		BatchDenom: batchDenom,
		Pagination: &query.PageRequest{CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))
	assertOrderEqual(s.ctx, t, s.k, res.SellOrders[0], order1)
	assert.Equal(t, uint64(1), res.Pagination.Total)

	// query the second denom
	res, err = s.k.SellOrdersByBatch(s.ctx, &types.QuerySellOrdersByBatchRequest{
		BatchDenom: otherDenom,
		Pagination: &query.PageRequest{CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))
	assertOrderEqual(s.ctx, t, s.k, res.SellOrders[0], order2)
	assert.Equal(t, uint64(1), res.Pagination.Total)

	// bad denom should error
	_, err = s.k.SellOrdersByBatch(s.ctx, &types.QuerySellOrdersByBatchRequest{
		BatchDenom: "yikes!",
		Pagination: nil,
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
