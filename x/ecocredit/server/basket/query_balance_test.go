package basket_test

import (
	"testing"

	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"

	"github.com/stretchr/testify/require"
)

func TestKeeper_BasketBalance(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// add a basket
	basketDenom := "foo"
	batchDenom := "bar"
	balance := "5.3"
	id, err := s.stateStore.BasketStore().InsertReturningID(s.ctx, &basketv1.Basket{
		BasketDenom: basketDenom,
	})
	require.NoError(t, err)

	// add a balance
	require.NoError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
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
	res, err = s.k.BasketBalance(s.ctx, &baskettypes.QueryBasketBalanceRequest{
		BasketDenom: batchDenom,
		BatchDenom:  basketDenom,
	})
	require.Error(t, err)
}
