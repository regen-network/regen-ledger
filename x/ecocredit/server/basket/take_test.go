package basket_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/golang/mock/gomock"
	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket/mocks"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTake(t *testing.T) {
	// prepare database
	db, err := ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{})
	stateStore, err := basketv1.NewStateStore(db)
	require.NoError(t, err)
	ctx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	// add some data
	fooBasketId, err := stateStore.BasketStore().InsertReturningID(ctx, &basketv1.Basket{
		BasketDenom:       "foo",
		DisableAutoRetire: false,
		CreditTypeName:    "C",
		Exponent:          6,
	})
	require.NoError(t, err)

	require.NoError(t, stateStore.BasketBalanceStore().Insert(ctx, &basketv1.BasketBalance{
		BasketId:       fooBasketId,
		BatchDenom:     "A",
		Balance:        "3.0",
		BatchStartDate: timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))

	require.NoError(t, stateStore.BasketBalanceStore().Insert(ctx, &basketv1.BasketBalance{
		BasketId:       fooBasketId,
		BatchDenom:     "B",
		Balance:        "5.0",
		BatchStartDate: timestamppb.New(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))

	barBasketId, err := stateStore.BasketStore().InsertReturningID(ctx, &basketv1.Basket{
		BasketDenom:       "bar",
		DisableAutoRetire: true,
		CreditTypeName:    "C",
		Exponent:          6,
	})
	require.NoError(t, err)

	// setup test keeper
	ctrl := gomock.NewController(t)
	require.NoError(t, err)
	bankKeeper := mocks.NewMockBankKeeper(ctrl)
	ecocreditKeeper := mocks.NewMockEcocreditKeeper(ctrl)
	k := basket.NewKeeper(db, ecocreditKeeper, bankKeeper)
}
