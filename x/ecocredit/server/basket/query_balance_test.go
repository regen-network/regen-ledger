package basket_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/golang/mock/gomock"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket/mocks"
	"github.com/stretchr/testify/require"
)

func TestKeeper_BasketBalance(t *testing.T) {
	// prepare database
	db, err := ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{})
	stateStore, err := basketv1.NewStateStore(db)
	require.NoError(t, err)
	ctx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	// setup test keeper
	ctrl := gomock.NewController(t)
	require.NoError(t, err)
	bankKeeper := mocks.NewMockBankKeeper(ctrl)
	ecocreditKeeper := mocks.NewMockEcocreditKeeper(ctrl)
	sk := sdk.NewKVStoreKey("test")
	k := basket.NewKeeper(db, ecocreditKeeper, bankKeeper, sk)

	// add a basket
	basketDenom := "foo"
	batchDenom := "bar"
	balance := "5.3"
	id, err := stateStore.BasketStore().InsertReturningID(ctx, &basketv1.Basket{
		BasketDenom: basketDenom,
	})
	require.NoError(t, err)

	// add a balance
	require.NoError(t, stateStore.BasketBalanceStore().Insert(ctx, &basketv1.BasketBalance{
		BasketId:   id,
		BatchDenom: batchDenom,
		Balance:    balance,
	}))

	// query
	res, err := k.BasketBalance(ctx, &baskettypes.QueryBasketBalanceRequest{
		BasketDenom: basketDenom,
		BatchDenom:  batchDenom,
	})
	require.NoError(t, err)
	require.Equal(t, balance, res.Balance)

	// bad query
	res, err = k.BasketBalance(ctx, &baskettypes.QueryBasketBalanceRequest{
		BasketDenom: batchDenom,
		BatchDenom:  basketDenom,
	})
	require.Error(t, err)
}
