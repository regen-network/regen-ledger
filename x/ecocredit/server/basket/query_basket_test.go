package basket_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket/mocks"
)

func TestKeeper_Basket(t *testing.T) {
	ctx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	// prepare database
	db, err := ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{})
	stateStore, err := basketv1.NewStateStore(db)
	require.NoError(t, err)

	// setup test keeper
	ctrl := gomock.NewController(t)
	require.NoError(t, err)
	bankKeeper := mocks.NewMockBankKeeper(ctrl)
	ecocreditKeeper := mocks.NewMockEcocreditKeeper(ctrl)
	k := basket.NewKeeper(db, ecocreditKeeper, bankKeeper)

	// add a basket
	basketDenom := "foo"
	batchDenom := "bar"
	_, err = stateStore.BasketStore().InsertReturningID(ctx, &basketv1.Basket{
		BasketDenom: basketDenom,
	})
	require.NoError(t, err)

	// query
	res, err := k.Basket(ctx, &baskettypes.QueryBasketRequest{
		BasketDenom: basketDenom,
	})
	require.NoError(t, err)
	require.Equal(t, basketDenom, res.Basket.BasketDenom)

	// bad query
	res, err = k.Basket(ctx, &baskettypes.QueryBasketRequest{
		BasketDenom: batchDenom,
	})
	require.Error(t, err)
}
