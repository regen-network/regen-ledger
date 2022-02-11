package basket_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"

	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/golang/mock/gomock"
	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket/mocks"
	"github.com/stretchr/testify/require"
)

func TestQueryBaskets(t *testing.T) {
	// prepare database
	db, err := ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{})
	stateStore, err := basketv1.NewStateStore(db)
	require.NoError(t, err)
	ctx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	// add some baskets
	require.NoError(t, stateStore.BasketStore().Insert(ctx, &basketv1.Basket{
		BasketDenom: "foo",
	}))
	require.NoError(t, stateStore.BasketStore().Insert(ctx, &basketv1.Basket{
		BasketDenom: "bar",
	}))
	require.NoError(t, stateStore.BasketStore().Insert(ctx, &basketv1.Basket{
		BasketDenom: "baz",
	}))

	// setup test keeper
	ctrl := gomock.NewController(t)
	require.NoError(t, err)
	bankKeeper := mocks.NewMockBankKeeper(ctrl)
	ecocreditKeeper := mocks.NewMockEcocreditKeeper(ctrl)
	sk := sdk.NewKVStoreKey("test")
	k := basket.NewKeeper(db, ecocreditKeeper, bankKeeper, sk, ecocredit.ModuleName)

	// query all
	res, err := k.Baskets(ctx, &baskettypes.QueryBasketsRequest{})
	require.NoError(t, err)
	require.Len(t, res.Baskets, 3)
	require.Equal(t, "foo", res.Baskets[0].BasketDenom)
	require.Equal(t, "bar", res.Baskets[1].BasketDenom)
	require.Equal(t, "baz", res.Baskets[2].BasketDenom)

	// paginate
	res, err = k.Baskets(ctx, &baskettypes.QueryBasketsRequest{
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
