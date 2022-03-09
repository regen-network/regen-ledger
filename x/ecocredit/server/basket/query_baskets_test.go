package basket_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/stretchr/testify/require"
)

func TestQueryBaskets(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// add some baskets
	require.NoError(t, s.stateStore.BasketStore().Insert(s.ctx, &basketv1.Basket{
		BasketDenom: "foo", Name: "foo",
	}))
	require.NoError(t, s.stateStore.BasketStore().Insert(s.ctx, &basketv1.Basket{
		BasketDenom: "bar", Name: "bar",
	}))
	require.NoError(t, s.stateStore.BasketStore().Insert(s.ctx, &basketv1.Basket{
		BasketDenom: "baz", Name: "baz",
	}))

	// query all
	res, err := s.k.Baskets(s.ctx, &baskettypes.QueryBasketsRequest{})
	require.NoError(t, err)
	require.Len(t, res.Baskets, 3)
	require.Equal(t, "foo", res.Baskets[0].BasketDenom)
	require.Equal(t, "bar", res.Baskets[1].BasketDenom)
	require.Equal(t, "baz", res.Baskets[2].BasketDenom)

	// paginate
	res, err = s.k.Baskets(s.ctx, &baskettypes.QueryBasketsRequest{
		Pagination: &query.PageRequest{
			Limit:      2,
			CountTotal: true,
			Reverse:    true,
		},
	})
	require.NoError(t, err)
	require.Len(t, res.Baskets, 2)
	require.Equal(t, "baz", res.Baskets[0].BasketDenom)
	require.Equal(t, "bar", res.Baskets[1].BasketDenom)
}
