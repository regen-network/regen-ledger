package basket_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func TestQueryBalances(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// add some baskets
	basketDenom := "foo"
	batchDenoms := []string{"bar", "baz", "qux"}
	require.NoError(t, s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom: basketDenom,
	}))
	require.NoError(t, s.stateStore.BasketBalanceTable().Insert(s.ctx, &api.BasketBalance{
		BasketId:       1,
		BatchDenom:     batchDenoms[0],
		Balance:        "100.50",
		BatchStartDate: nil,
	}))
	require.NoError(t, s.stateStore.BasketBalanceTable().Insert(s.ctx, &api.BasketBalance{
		BasketId:       1,
		BatchDenom:     batchDenoms[1],
		Balance:        "4.20",
		BatchStartDate: nil,
	}))
	require.NoError(t, s.stateStore.BasketBalanceTable().Insert(s.ctx, &api.BasketBalance{
		BasketId:       1,
		BatchDenom:     batchDenoms[2],
		Balance:        "6.10",
		BatchStartDate: nil,
	}))

	// setup test keeper

	// query all
	res, err := s.k.BasketBalances(s.ctx, &baskettypes.QueryBasketBalancesRequest{BasketDenom: basketDenom})
	require.NoError(t, err)
	require.Len(t, res.Balances, 3)
	require.Equal(t, "100.50", res.Balances[0].Balance)
	require.Equal(t, "4.20", res.Balances[1].Balance)
	require.Equal(t, "6.10", res.Balances[2].Balance)

	// paginate
	res, err = s.k.BasketBalances(s.ctx, &baskettypes.QueryBasketBalancesRequest{
		BasketDenom: basketDenom,
		Pagination: &query.PageRequest{
			Limit:      2,
			CountTotal: true,
			Reverse:    true,
		},
	})
	require.NoError(t, err)
	require.Equal(t, res.Pagination.Total, uint64(3))
	require.Len(t, res.Balances, 2)
	require.Equal(t, "6.10", res.Balances[0].Balance)
	require.Equal(t, "4.20", res.Balances[1].Balance)

	// bad query
	_, err = s.k.BasketBalances(s.ctx, &baskettypes.QueryBasketBalancesRequest{BasketDenom: "nope"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "not found")
}
