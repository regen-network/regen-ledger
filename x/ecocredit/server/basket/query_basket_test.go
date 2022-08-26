package basket

import (
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func TestKeeper_Basket(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// add a basket
	batchDenom := "bar"
	err := s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom: testBasketDenom,
	})
	require.NoError(t, err)

	// query
	res, err := s.k.Basket(s.ctx, &baskettypes.QueryBasketRequest{
		BasketDenom: testBasketDenom,
	})
	require.NoError(t, err)
	require.Equal(t, testBasketDenom, res.Basket.BasketDenom)

	// bad query
	_, err = s.k.Basket(s.ctx, &baskettypes.QueryBasketRequest{
		BasketDenom: batchDenom,
	})
	require.Error(t, err)
}

func TestKeeper_BasketClasses(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// add a basket
	err := s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom: testBasketDenom,
	})
	require.NoError(t, err)

	// add a basket class
	classID := "C01"
	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: 1,
		ClassId:  classID,
	})
	require.NoError(t, err)

	// query
	res, err := s.k.Basket(s.ctx, &baskettypes.QueryBasketRequest{
		BasketDenom: testBasketDenom,
	})
	require.NoError(t, err)
	require.Equal(t, testBasketDenom, res.Basket.BasketDenom)
	require.Equal(t, []string{classID}, res.Classes)

	// query unknown basket
	_, err = s.k.Basket(s.ctx, &baskettypes.QueryBasketRequest{
		BasketDenom: "unknown",
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "basket unknown not found")
}
