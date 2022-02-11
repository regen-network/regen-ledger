package basket_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/regen-network/regen-ledger/x/ecocredit"

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

type suite struct {
	db              ormdb.ModuleDB
	stateStore      basketv1.StateStore
	ctx             context.Context
	k               basket.Keeper
	ctrl            *gomock.Controller
	acct            sdk.AccAddress
	bankKeeper      *mocks.MockBankKeeper
	ecocreditKeeper *mocks.MockEcocreditKeeper
	fooBasketId     uint64
	barBasketId     uint64
}

func setup(t *testing.T) *suite {
	// prepare database
	s := &suite{}
	var err error
	s.db, err = ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{})
	s.stateStore, err = basketv1.NewStateStore(s.db)
	assert.NilError(t, err)
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.ctx = sdk.WrapSDKContext(sdk.Context{}.WithContext(ormCtx))

	// add some data
	s.fooBasketId, err = s.stateStore.BasketStore().InsertReturningID(s.ctx, &basketv1.Basket{
		BasketDenom:       "foo",
		DisableAutoRetire: false,
		CreditTypeName:    "C",
		Exponent:          6,
	})
	assert.NilError(t, err)

	assert.NilError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
		BasketId:       s.fooBasketId,
		BatchDenom:     "C1",
		Balance:        "3.0",
		BatchStartDate: timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))

	assert.NilError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
		BasketId:       s.fooBasketId,
		BatchDenom:     "C2",
		Balance:        "5.0",
		BatchStartDate: timestamppb.New(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))

	s.barBasketId, err = s.stateStore.BasketStore().InsertReturningID(s.ctx, &basketv1.Basket{
		BasketDenom:       "bar",
		DisableAutoRetire: true,
		CreditTypeName:    "C",
		Exponent:          6,
	})
	assert.NilError(t, err)

	assert.NilError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
		BasketId:       s.barBasketId,
		BatchDenom:     "C3",
		Balance:        "7.0",
		BatchStartDate: timestamppb.New(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))

	assert.NilError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
		BasketId:       s.barBasketId,
		BatchDenom:     "C4",
		Balance:        "4.0",
		BatchStartDate: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))

	// setup test keeper
	s.ctrl = gomock.NewController(t)
	assert.NilError(t, err)
	s.bankKeeper = mocks.NewMockBankKeeper(s.ctrl)
	s.ecocreditKeeper = mocks.NewMockEcocreditKeeper(s.ctrl)
	sk := sdk.NewKVStoreKey("test")
	s.k = basket.NewKeeper(s.db, s.ecocreditKeeper, s.bankKeeper, sk, ecocredit.ModuleName)

	s.acct = sdk.AccAddress{0, 1, 2, 3, 4, 5}

	return s
}

func TestTakeMustRetire(t *testing.T) {
	t.Parallel()
	s := setup(t)

	// foo requires RetireOnTake
	_, err := s.k.Take(s.ctx, &baskettypes.MsgTake{
		Owner:              s.acct.String(),
		BasketDenom:        "foo",
		Amount:             "6.0",
		RetirementLocation: "",
		RetireOnTake:       false,
	})
	assert.ErrorIs(t, err, basket.ErrCantDisableRetire)
}

func TestTakeRetire(t *testing.T) {
	t.Parallel()
	s := setup(t)

	fooCoins := sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(6000000)))
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), s.acct, ecocredit.ModuleName, fooCoins)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), ecocredit.ModuleName, fooCoins)
	s.ecocreditKeeper.EXPECT().AddCreditBalance(gomock.Any(), s.acct, "C2", math.MatchEq(math.NewDecFromInt64(5)), true, "US")
	s.ecocreditKeeper.EXPECT().AddCreditBalance(gomock.Any(), s.acct, "C1", math.MatchEq(math.NewDecFromInt64(1)), true, "US")

	res, err := s.k.Take(s.ctx, &baskettypes.MsgTake{
		Owner:              s.acct.String(),
		BasketDenom:        "foo",
		Amount:             "6000000",
		RetirementLocation: "US",
		RetireOnTake:       true,
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Credits))
	assert.Equal(t, "C2", res.Credits[0].BatchDenom)
	assertDecStringEqual(t, "5.0", res.Credits[0].Amount)
	assert.Equal(t, "C1", res.Credits[1].BatchDenom)
	assertDecStringEqual(t, "1.0", res.Credits[1].Amount)
	found, err := s.stateStore.BasketBalanceStore().Has(s.ctx, s.fooBasketId, "C2")
	assert.NilError(t, err)
	assert.Assert(t, !found)
	balance, err := s.stateStore.BasketBalanceStore().Get(s.ctx, s.fooBasketId, "C1")
	assert.NilError(t, err)
	assertDecStringEqual(t, "2.0", balance.Balance)
}

func TestTakeTradable(t *testing.T) {
	t.Parallel()
	s := setup(t)

	barCoins := sdk.NewCoins(sdk.NewCoin("bar", sdk.NewInt(10000000)))
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), s.acct, ecocredit.ModuleName, barCoins)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), ecocredit.ModuleName, barCoins)
	s.ecocreditKeeper.EXPECT().AddCreditBalance(gomock.Any(), s.acct, "C3", math.MatchEq(math.NewDecFromInt64(7)), false, "")
	s.ecocreditKeeper.EXPECT().AddCreditBalance(gomock.Any(), s.acct, "C4", math.MatchEq(math.NewDecFromInt64(3)), false, "")

	res, err := s.k.Take(s.ctx, &baskettypes.MsgTake{
		Owner:        s.acct.String(),
		BasketDenom:  "bar",
		Amount:       "10000000",
		RetireOnTake: false,
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Credits))
	assert.Equal(t, "C3", res.Credits[0].BatchDenom)
	assertDecStringEqual(t, "7.0", res.Credits[0].Amount)
	assert.Equal(t, "C4", res.Credits[1].BatchDenom)
	assertDecStringEqual(t, "3.0", res.Credits[1].Amount)
	found, err := s.stateStore.BasketBalanceStore().Has(s.ctx, s.barBasketId, "C3")
	assert.NilError(t, err)
	assert.Assert(t, !found)
	balance, err := s.stateStore.BasketBalanceStore().Get(s.ctx, s.barBasketId, "C4")
	assert.NilError(t, err)
	assertDecStringEqual(t, "1.0", balance.Balance)
}

func assertDecStringEqual(t *testing.T, expected, actual string) {
	dx, err := math.NewDecFromString(expected)
	assert.NilError(t, err)
	dy, err := math.NewDecFromString(actual)
	assert.NilError(t, err)
	assert.Assert(t, 0 == dx.Cmp(dy), fmt.Sprintf("%s != %s", expected, actual))
}
