package basket_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	coreapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func TestKeeper_BasketBalance(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// add a basket
	basketDenom := testBasketDenom
	batchDenom := "bar"
	balance := "5.3"
	id, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom: basketDenom,
	})
	require.NoError(t, err)

	err = s.coreStore.BatchTable().Insert(s.ctx, &coreapi.Batch{
		Denom: batchDenom,
	})
	require.NoError(t, err)

	// add a balance
	require.NoError(t, s.stateStore.BasketBalanceTable().Insert(s.ctx, &api.BasketBalance{
		BasketId:   id,
		BatchDenom: batchDenom,
		Balance:    balance,
	}))

	// query
	res, err := s.k.BasketBalance(s.ctx, &baskettypes.QueryBasketBalanceRequest{
		BasketDenom: basketDenom,
		BatchDenom:  batchDenom,
	})
	require.NoError(t, err)
	require.Equal(t, balance, res.Balance)

	// bad query
	_, err = s.k.BasketBalance(s.ctx, &baskettypes.QueryBasketBalanceRequest{
		BasketDenom: batchDenom,
		BatchDenom:  basketDenom,
	})
	require.Error(t, err)

	// add another basket
	basketDenom = "foo1"
	basketName := "foo1.bar"
	err = s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom: basketDenom,
		Name:        basketName,
	})
	require.NoError(t, err)

	// expect empty basket balance
	res, err = s.k.BasketBalance(s.ctx, &baskettypes.QueryBasketBalanceRequest{
		BasketDenom: basketDenom,
		BatchDenom:  batchDenom,
	})
	require.NoError(t, err)
	require.Equal(t, res.Balance, "0")
}
