package basket_test

import (
	"testing"
	"time"

	"github.com/regen-network/regen-ledger/types/math"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket/mocks"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"
)

func TestTake(t *testing.T) {
	// prepare database
	db, err := ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{})
	stateStore, err := basketv1.NewStateStore(db)
	assert.NilError(t, err)
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	ctx := sdk.WrapSDKContext(sdk.Context{}.WithContext(ormCtx))

	// add some data
	fooBasketId, err := stateStore.BasketStore().InsertReturningID(ctx, &basketv1.Basket{
		BasketDenom:       "foo",
		DisableAutoRetire: false,
		CreditTypeName:    "C",
		Exponent:          6,
	})
	assert.NilError(t, err)

	assert.NilError(t, stateStore.BasketBalanceStore().Insert(ctx, &basketv1.BasketBalance{
		BasketId:       fooBasketId,
		BatchDenom:     "C1",
		Balance:        "3.0",
		BatchStartDate: timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))

	assert.NilError(t, stateStore.BasketBalanceStore().Insert(ctx, &basketv1.BasketBalance{
		BasketId:       fooBasketId,
		BatchDenom:     "C2",
		Balance:        "5.0",
		BatchStartDate: timestamppb.New(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))

	barBasketId, err := stateStore.BasketStore().InsertReturningID(ctx, &basketv1.Basket{
		BasketDenom:       "bar",
		DisableAutoRetire: true,
		CreditTypeName:    "C",
		Exponent:          6,
	})
	assert.NilError(t, err)

	assert.NilError(t, stateStore.BasketBalanceStore().Insert(ctx, &basketv1.BasketBalance{
		BasketId:       barBasketId,
		BatchDenom:     "C3",
		Balance:        "7.0",
		BatchStartDate: timestamppb.New(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))

	assert.NilError(t, stateStore.BasketBalanceStore().Insert(ctx, &basketv1.BasketBalance{
		BasketId:       barBasketId,
		BatchDenom:     "C4",
		Balance:        "4.0",
		BatchStartDate: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))

	// setup test keeper
	ctrl := gomock.NewController(t)
	assert.NilError(t, err)
	bankKeeper := mocks.NewMockBankKeeper(ctrl)
	ecocreditKeeper := mocks.NewMockEcocreditKeeper(ctrl)
	moduleAccountName := "basket"
	k := basket.NewKeeper(db, ecocreditKeeper, bankKeeper, moduleAccountName)

	acct := sdk.AccAddress{0, 1, 2, 3, 4, 5}

	// foo requires RetireOnTake
	_, err = k.Take(ctx, &baskettypes.MsgTake{
		Owner:              acct.String(),
		BasketDenom:        "foo",
		Amount:             "6.0",
		RetirementLocation: "",
		RetireOnTake:       false,
	})
	assert.ErrorIs(t, err, basket.ErrCantDisableRetire)

	fooCoins := sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(6000000)))
	bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), acct, moduleAccountName, fooCoins)
	bankKeeper.EXPECT().BurnCoins(gomock.Any(), moduleAccountName, fooCoins)
	ecocreditKeeper.EXPECT().AddCreditBalance(gomock.Any(), acct, "C2", math.MatchEq(math.NewDecFromInt64(5)), true, "US")
	ecocreditKeeper.EXPECT().AddCreditBalance(gomock.Any(), acct, "C1", math.MatchEq(math.NewDecFromInt64(1)), true, "US")

	res, err := k.Take(ctx, &baskettypes.MsgTake{
		Owner:              acct.String(),
		BasketDenom:        "foo",
		Amount:             "6000000",
		RetirementLocation: "US",
		RetireOnTake:       true,
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Credits))
	assert.Equal(t, "C2", res.Credits[0].BatchDenom)
	assert.Assert(t, "5.0", res.Credits[0].Amount)
	assert.Equal(t, "C1", res.Credits[1].BatchDenom)
	assert.Equal(t, "1.0", res.Credits[1].Amount)
}
