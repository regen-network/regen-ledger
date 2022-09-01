package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func TestBalancesByBatch(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	addr1 := sdk.AccAddress("addr1")
	addr2 := sdk.AccAddress("addr2")

	// denom hasn't been inserted yet, so expect not found error
	denom := "C01-001-20200101-20220101-001"
	_, err := s.k.BalancesByBatch(s.ctx, &types.QueryBalancesByBatchRequest{BatchDenom: denom})
	require.ErrorContains(t, err, fmt.Sprintf("could not get batch with denom %s: not found", denom))

	id, err := s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{Denom: denom})
	require.NoError(t, err)

	bal1 := api.BatchBalance{
		BatchKey:       id,
		Address:        addr1,
		TradableAmount: "10",
		RetiredAmount:  "20",
		EscrowedAmount: "30",
	}
	bal2 := api.BatchBalance{
		BatchKey:       id,
		Address:        addr2,
		TradableAmount: "15",
		RetiredAmount:  "25",
		EscrowedAmount: "35",
	}
	bal3 := api.BatchBalance{
		BatchKey:       50,
		Address:        addr1,
		TradableAmount: "10",
		RetiredAmount:  "20",
		EscrowedAmount: "30",
	}
	require.NoError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, &bal1))
	require.NoError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, &bal2))
	require.NoError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, &bal3))

	expected := []*api.BatchBalance{&bal1, &bal2}
	res, err := s.k.BalancesByBatch(s.ctx, &types.QueryBalancesByBatchRequest{BatchDenom: denom, Pagination: &query.PageRequest{Limit: 10, CountTotal: true}})
	require.NoError(t, err)
	require.Len(t, res.Balances, len(expected))
	require.Equal(t, res.Pagination.Total, uint64(len(expected)))
	for i, bal := range res.Balances {
		s.assertBalanceEqual(bal, expected[i])
	}
}
