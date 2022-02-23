package basket_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
)

type takeSuite struct {
	*baseSuite
	fooBasketId uint64
	barBasketId uint64
}

func setupTake(t *testing.T) *takeSuite {
	// prepare database
	s := &takeSuite{baseSuite: setupBase(t)}

	// add some data
	var err error
	s.fooBasketId, err = s.stateStore.BasketStore().InsertReturningID(s.ctx, &basketv1.Basket{
		BasketDenom:       "foo",
		Name:              "foo",
		DisableAutoRetire: false,
		CreditTypeAbbrev:  "C",
		Exponent:          6,
	})
	assert.NilError(t, err)

	assert.NilError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
		BasketId:       s.fooBasketId,
		BatchDenom:     "C1",
		Balance:        "3.0",
		BatchStartDate: timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))
	s.setTradableSupply("C1", "3.0")

	assert.NilError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
		BasketId:       s.fooBasketId,
		BatchDenom:     "C2",
		Balance:        "5.0",
		BatchStartDate: timestamppb.New(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))
	s.setTradableSupply("C2", "5.0")

	s.barBasketId, err = s.stateStore.BasketStore().InsertReturningID(s.ctx, &basketv1.Basket{
		BasketDenom:       "bar",
		Name:              "bar",
		DisableAutoRetire: true,
		CreditTypeAbbrev:  "C",
		Exponent:          6,
	})
	assert.NilError(t, err)

	assert.NilError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
		BasketId:       s.barBasketId,
		BatchDenom:     "C3",
		Balance:        "7.0",
		BatchStartDate: timestamppb.New(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))
	s.setTradableSupply("C3", "7.0")

	assert.NilError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
		BasketId:       s.barBasketId,
		BatchDenom:     "C4",
		Balance:        "4.0",
		BatchStartDate: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))
	s.setTradableSupply("C4", "4.0")

	return s
}

func TestTakeMustRetire(t *testing.T) {
	t.Parallel()
	s := setupTake(t)

	// foo requires RetireOnTake
	_, err := s.k.Take(s.ctx, &baskettypes.MsgTake{
		Owner:              s.addr.String(),
		BasketDenom:        "foo",
		Amount:             "6.0",
		RetirementLocation: "",
		RetireOnTake:       false,
	})
	assert.ErrorIs(t, err, basket.ErrCantDisableRetire)
}

func TestTakeRetire(t *testing.T) {
	t.Parallel()
	s := setupTake(t)

	fooCoins := sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(6000000)))
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), s.addr, baskettypes.BasketSubModuleName, fooCoins)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), baskettypes.BasketSubModuleName, fooCoins)

	res, err := s.k.Take(s.ctx, &baskettypes.MsgTake{
		Owner:              s.addr.String(),
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

	s.expectTradableBalance("C1", "0")
	s.expectTradableBalance("C2", "0")
	s.expectRetiredBalance("C1", "1")
	s.expectRetiredBalance("C2", "5")
	s.expectTradableSupply("C1", "2")
	s.expectTradableSupply("C2", "0")
	s.expectRetiredSupply("C1", "1")
	s.expectRetiredSupply("C2", "5")
}

func TestTakeTradable(t *testing.T) {
	t.Parallel()
	s := setupTake(t)

	barCoins := sdk.NewCoins(sdk.NewCoin("bar", sdk.NewInt(10000000)))
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), s.addr, baskettypes.BasketSubModuleName, barCoins)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), baskettypes.BasketSubModuleName, barCoins)

	res, err := s.k.Take(s.ctx, &baskettypes.MsgTake{
		Owner:        s.addr.String(),
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

	s.expectTradableBalance("C3", "7")
	s.expectTradableBalance("C4", "3")
	s.expectRetiredBalance("C3", "0")
	s.expectRetiredBalance("C4", "0")
	s.expectTradableSupply("C3", "7")
	s.expectTradableSupply("C4", "4")
	s.expectRetiredSupply("C3", "0")
	s.expectRetiredSupply("C4", "0")
}

func TestTakeTooMuchTradable(t *testing.T) {
	t.Parallel()
	s := setupTake(t)

	// Try to take more than what's in the basket, should error.
	amount := sdk.NewInt(99999999999)
	barCoins := sdk.NewCoins(sdk.NewCoin("bar", amount))
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), s.addr, baskettypes.BasketSubModuleName, barCoins)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), baskettypes.BasketSubModuleName, barCoins)

	_, err := s.k.Take(s.ctx, &baskettypes.MsgTake{
		Owner:        s.addr.String(),
		BasketDenom:  "bar",
		Amount:       amount.String(),
		RetireOnTake: false,
	})
	// IRL, the error below should throw earlier, on SendCoinsFromAccountToModule
	// We're just testing here that some error is thrown.
	assert.Error(t, err, "unexpected failure - balance invariant broken")
}

func TestTakeAllTradable(t *testing.T) {
	t.Parallel()
	s := setupTake(t)

	barCoins := sdk.NewCoins(sdk.NewCoin("bar", sdk.NewInt(11000000))) // == 7C4 + 4C4
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), s.addr, baskettypes.BasketSubModuleName, barCoins)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), baskettypes.BasketSubModuleName, barCoins)

	res, err := s.k.Take(s.ctx, &baskettypes.MsgTake{
		Owner:        s.addr.String(),
		BasketDenom:  "bar",
		Amount:       "11000000",
		RetireOnTake: false,
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Credits))
	assert.Equal(t, "C3", res.Credits[0].BatchDenom)
	assertDecStringEqual(t, "7.0", res.Credits[0].Amount)
	assert.Equal(t, "C4", res.Credits[1].BatchDenom)
	assertDecStringEqual(t, "4.0", res.Credits[1].Amount)
	found, err := s.stateStore.BasketBalanceStore().Has(s.ctx, s.barBasketId, "C3")
	assert.NilError(t, err)
	assert.Assert(t, !found)
	_, err = s.stateStore.BasketBalanceStore().Get(s.ctx, s.barBasketId, "C4")
	assert.ErrorIs(t, err, ormerrors.NotFound)

	s.expectTradableBalance("C3", "7")
	s.expectTradableBalance("C4", "4")
	s.expectRetiredBalance("C3", "0")
	s.expectRetiredBalance("C4", "0")
	s.expectTradableSupply("C3", "7")
	s.expectTradableSupply("C4", "4")
	s.expectRetiredSupply("C3", "0")
	s.expectRetiredSupply("C4", "0")
}

func assertDecStringEqual(t *testing.T, expected, actual string) {
	dx, err := math.NewDecFromString(expected)
	assert.NilError(t, err)
	dy, err := math.NewDecFromString(actual)
	assert.NilError(t, err)
	assert.Assert(t, 0 == dx.Cmp(dy), fmt.Sprintf("%s != %s", expected, actual))
}

func (s takeSuite) expectTradableBalance(batchDenom string, expected string) {
	kvStore := s.sdkCtx.KVStore(s.storeKey)
	bal, err := ecocredit.GetDecimal(kvStore, ecocredit.TradableBalanceKey(s.addr, ecocredit.BatchDenomT(batchDenom)))
	assert.NilError(s.t, err)
	s.expectDec(expected, bal)
}

func (s takeSuite) expectRetiredBalance(batchDenom string, expected string) {
	kvStore := s.sdkCtx.KVStore(s.storeKey)
	bal, err := ecocredit.GetDecimal(kvStore, ecocredit.RetiredBalanceKey(s.addr, ecocredit.BatchDenomT(batchDenom)))
	assert.NilError(s.t, err)
	s.expectDec(expected, bal)
}

func (s takeSuite) expectTradableSupply(batchDenom string, expected string) {
	kvStore := s.sdkCtx.KVStore(s.storeKey)
	bal, err := ecocredit.GetDecimal(kvStore, ecocredit.TradableSupplyKey(ecocredit.BatchDenomT(batchDenom)))
	assert.NilError(s.t, err)
	s.expectDec(expected, bal)
}

func (s takeSuite) expectRetiredSupply(batchDenom string, expected string) {
	kvStore := s.sdkCtx.KVStore(s.storeKey)
	bal, err := ecocredit.GetDecimal(kvStore, ecocredit.RetiredSupplyKey(ecocredit.BatchDenomT(batchDenom)))
	assert.NilError(s.t, err)
	s.expectDec(expected, bal)
}

func (s takeSuite) setTradableSupply(batchDenom string, amount string) {
	kvStore := s.sdkCtx.KVStore(s.storeKey)
	dec, err := math.NewDecFromString(amount)
	assert.NilError(s.t, err)
	ecocredit.SetDecimal(kvStore, ecocredit.TradableSupplyKey(ecocredit.BatchDenomT(batchDenom)), dec)
}

func (s takeSuite) expectDec(expected string, actual math.Dec) {
	dec, err := math.NewDecFromString(expected)
	assert.NilError(s.t, err)
	assert.Assert(s.t, actual.Cmp(dec) == 0)
}
