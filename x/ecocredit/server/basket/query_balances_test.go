package basket_test

import (
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/mock/gomock"
	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueryBalances(t *testing.T) {
	// prepare database
	db, err := ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{})
	stateStore, err := basketv1.NewStateStore(db)
	require.NoError(t, err)
	ctx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	// add some baskets
	basketDenom := "foo"
	batchDenoms := []string{"bar", "baz", "qux"}
	require.NoError(t, stateStore.BasketStore().Insert(ctx, &basketv1.Basket{
		BasketDenom: basketDenom,
	}))
	require.NoError(t, stateStore.BasketBalanceStore().Insert(ctx, &basketv1.BasketBalance{
		BasketId:       1,
		BatchDenom:     batchDenoms[0],
		Balance:        "100.50",
		BatchStartDate: nil,
	}))
	require.NoError(t, stateStore.BasketBalanceStore().Insert(ctx, &basketv1.BasketBalance{
		BasketId:       1,
		BatchDenom:     batchDenoms[1],
		Balance:        "4.20",
		BatchStartDate: nil,
	}))
	require.NoError(t, stateStore.BasketBalanceStore().Insert(ctx, &basketv1.BasketBalance{
		BasketId:       1,
		BatchDenom:     batchDenoms[2],
		Balance:        "6.10",
		BatchStartDate: nil,
	}))


	// setup test keeper
	ctrl := gomock.NewController(t)
	require.NoError(t, err)
	bankKeeper := mocks.NewMockBankKeeper(ctrl)
	ecocreditKeeper := mocks.NewMockEcocreditKeeper(ctrl)
	k := basket.NewKeeper(db, ecocreditKeeper, bankKeeper)

	// query all
	res, err := k.BasketBalances(ctx, &baskettypes.QueryBasketBalancesRequest{BasketDenom: basketDenom})
	require.NoError(t, err)
	require.Len(t, res.Balances, 3)
	require.Equal(t, "100.50", res.Balances[0].Balance)
	require.Equal(t, "4.20", res.Balances[1].Balance)
	require.Equal(t, "6.10", res.Balances[2].Balance)


	// paginate
	res, err = k.BasketBalances(ctx, &baskettypes.QueryBasketBalancesRequest{
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
	res, err = k.BasketBalances(ctx, &baskettypes.QueryBasketBalancesRequest{BasketDenom: "nope"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "not found")
}
