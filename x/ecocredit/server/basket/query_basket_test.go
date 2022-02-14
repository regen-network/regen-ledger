package basket_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func TestKeeper_Basket(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// add a basket
	basketDenom := "foo"
	batchDenom := "bar"
	err := s.stateStore.BasketStore().Insert(s.ctx, &basketv1.Basket{
		BasketDenom: basketDenom,
	})
	require.NoError(t, err)

	// query
	res, err := s.k.Basket(s.ctx, &baskettypes.QueryBasketRequest{
		BasketDenom: basketDenom,
	})
	require.NoError(t, err)
	require.Equal(t, basketDenom, res.Basket.BasketDenom)

	// bad query
	res, err = s.k.Basket(s.ctx, &baskettypes.QueryBasketRequest{
		BasketDenom: batchDenom,
	})
	require.Error(t, err)
}
