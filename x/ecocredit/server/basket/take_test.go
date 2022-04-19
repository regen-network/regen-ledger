package basket_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
)

type takeSuite struct {
	*baseSuite
	fooBasketId uint64
	barBasketId uint64
	denomToId   map[string]uint64
}

func setupTake(t *testing.T) *takeSuite {
	// prepare database
	s := &takeSuite{baseSuite: setupBase(t)}

	// add some data
	var err error
	s.fooBasketId, err = s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:       "foo",
		Name:              "foo",
		DisableAutoRetire: false,
		CreditTypeAbbrev:  "C",
		Exponent:          6,
	})
	assert.NilError(t, err)

	assert.NilError(t, s.stateStore.BasketBalanceTable().Insert(s.ctx, &api.BasketBalance{
		BasketId:       s.fooBasketId,
		BatchDenom:     "C1-00000000-0000000-001",
		Balance:        "3.0",
		BatchStartDate: timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))
	s.setTradableSupply(1, "3.0")

	assert.NilError(t, s.stateStore.BasketBalanceTable().Insert(s.ctx, &api.BasketBalance{
		BasketId:       s.fooBasketId,
		BatchDenom:     "C2-00000000-0000000-001",
		Balance:        "5.0",
		BatchStartDate: timestamppb.New(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))
	s.setTradableSupply(2, "5.0")

	s.barBasketId, err = s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:       "bar",
		Name:              "bar",
		DisableAutoRetire: true,
		CreditTypeAbbrev:  "C",
		Exponent:          6,
	})
	assert.NilError(t, err)

	assert.NilError(t, s.stateStore.BasketBalanceTable().Insert(s.ctx, &api.BasketBalance{
		BasketId:       s.barBasketId,
		BatchDenom:     "C3-00000000-0000000-001",
		Balance:        "7.0",
		BatchStartDate: timestamppb.New(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))
	s.setTradableSupply(3, "7.0")

	assert.NilError(t, s.stateStore.BasketBalanceTable().Insert(s.ctx, &api.BasketBalance{
		BasketId:       s.barBasketId,
		BatchDenom:     "C4-00000000-0000000-001",
		Balance:        "4.0",
		BatchStartDate: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))
	s.setTradableSupply(4, "4.0")

	batchDenoms := []string{"C1-00000000-0000000-001", "C2-00000000-0000000-001", "C3-00000000-0000000-001", "C4-00000000-0000000-001"}
	for _, denom := range batchDenoms {
		assert.NilError(t, s.coreStore.BatchInfoTable().Insert(s.ctx, &ecoApi.BatchInfo{
			ProjectKey: 1,
			BatchDenom: denom,
		}))
	}
	s.denomToId = map[string]uint64{"C1-00000000-0000000-001": 1, "C2-00000000-0000000-001": 2, "C3-00000000-0000000-001": 3, "C4-00000000-0000000-001": 4}
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
	assert.Equal(t, "C2-00000000-0000000-001", res.Credits[0].BatchDenom)
	assertDecStringEqual(t, "5.0", res.Credits[0].Amount)
	assert.Equal(t, "C1-00000000-0000000-001", res.Credits[1].BatchDenom)
	assertDecStringEqual(t, "1.0", res.Credits[1].Amount)
	found, err := s.stateStore.BasketBalanceTable().Has(s.ctx, s.fooBasketId, "C2-00000000-0000000-001")
	assert.NilError(t, err)
	assert.Assert(t, !found)
	balance, err := s.stateStore.BasketBalanceTable().Get(s.ctx, s.fooBasketId, "C1-00000000-0000000-001")
	assert.NilError(t, err)
	assertDecStringEqual(t, "2.0", balance.Balance)

	s.expectTradableBalance("C1-00000000-0000000-001", "0")
	s.expectTradableBalance("C2-00000000-0000000-001", "0")
	s.expectRetiredBalance("C1-00000000-0000000-001", "1")
	s.expectRetiredBalance("C2-00000000-0000000-001", "5")
	s.expectTradableSupply("C1-00000000-0000000-001", "2")
	s.expectTradableSupply("C2-00000000-0000000-001", "0")
	s.expectRetiredSupply("C1-00000000-0000000-001", "1")
	s.expectRetiredSupply("C2-00000000-0000000-001", "5")
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
	assert.Equal(t, "C3-00000000-0000000-001", res.Credits[0].BatchDenom)
	assertDecStringEqual(t, "7.0", res.Credits[0].Amount)
	assert.Equal(t, "C4-00000000-0000000-001", res.Credits[1].BatchDenom)
	assertDecStringEqual(t, "3.0", res.Credits[1].Amount)
	found, err := s.stateStore.BasketBalanceTable().Has(s.ctx, s.barBasketId, "C3-00000000-0000000-001")
	assert.NilError(t, err)
	assert.Assert(t, !found)
	balance, err := s.stateStore.BasketBalanceTable().Get(s.ctx, s.barBasketId, "C4-00000000-0000000-001")
	assert.NilError(t, err)
	assertDecStringEqual(t, "1.0", balance.Balance)

	s.expectTradableBalance("C3-00000000-0000000-001", "7")
	s.expectTradableBalance("C4-00000000-0000000-001", "3")
	s.expectRetiredBalance("C3-00000000-0000000-001", "0")
	s.expectRetiredBalance("C4-00000000-0000000-001", "0")
	s.expectTradableSupply("C3-00000000-0000000-001", "7")
	s.expectTradableSupply("C4-00000000-0000000-001", "4")
	s.expectRetiredSupply("C3-00000000-0000000-001", "0")
	s.expectRetiredSupply("C4-00000000-0000000-001", "0")
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
	assert.Equal(t, "C3-00000000-0000000-001", res.Credits[0].BatchDenom)
	assertDecStringEqual(t, "7.0", res.Credits[0].Amount)
	assert.Equal(t, "C4-00000000-0000000-001", res.Credits[1].BatchDenom)
	assertDecStringEqual(t, "4.0", res.Credits[1].Amount)
	found, err := s.stateStore.BasketBalanceTable().Has(s.ctx, s.barBasketId, "C3-00000000-0000000-001")
	assert.NilError(t, err)
	assert.Assert(t, !found)
	_, err = s.stateStore.BasketBalanceTable().Get(s.ctx, s.barBasketId, "C4-00000000-0000000-001")
	assert.ErrorIs(t, err, ormerrors.NotFound)

	s.expectTradableBalance("C3-00000000-0000000-001", "7")
	s.expectTradableBalance("C4-00000000-0000000-001", "4")
	s.expectRetiredBalance("C3-00000000-0000000-001", "0")
	s.expectRetiredBalance("C4-00000000-0000000-001", "0")
	s.expectTradableSupply("C3-00000000-0000000-001", "7")
	s.expectTradableSupply("C4-00000000-0000000-001", "4")
	s.expectRetiredSupply("C3-00000000-0000000-001", "0")
	s.expectRetiredSupply("C4-00000000-0000000-001", "0")
}

func assertDecStringEqual(t *testing.T, expected, actual string) {
	dx, err := math.NewDecFromString(expected)
	assert.NilError(t, err)
	dy, err := math.NewDecFromString(actual)
	assert.NilError(t, err)
	assert.Assert(t, 0 == dx.Cmp(dy), fmt.Sprintf("%s != %s", expected, actual))
}

func (s takeSuite) expectTradableBalance(batchDenom string, expected string) {
	bal := s.getUserBalance(batchDenom)
	s.expectDec(expected, bal.Tradable)
}

func (s takeSuite) expectRetiredBalance(batchDenom string, expected string) {
	bal := s.getUserBalance(batchDenom)
	s.expectDec(expected, bal.Retired)
}

func (s takeSuite) expectTradableSupply(batchDenom string, expected string) {
	supply := s.getSupply(batchDenom)
	s.expectDec(expected, supply.TradableAmount)
}

func (s takeSuite) expectRetiredSupply(batchDenom string, expected string) {
	supply := s.getSupply(batchDenom)
	s.expectDec(expected, supply.RetiredAmount)
}

func (s takeSuite) getSupply(batchDenom string) *ecoApi.BatchSupply {
	id := s.denomToId[batchDenom]
	supply, err := s.coreStore.BatchSupplyTable().Get(s.ctx, id)
	if ormerrors.IsNotFound(err) {
		supply = &ecoApi.BatchSupply{
			BatchKey:        id,
			TradableAmount:  "0",
			RetiredAmount:   "0",
			CancelledAmount: "0",
		}
	} else {
		assert.NilError(s.t, err)
	}
	return supply
}

func (s takeSuite) setTradableSupply(batchId uint64, amount string) {
	assert.NilError(s.t, s.coreStore.BatchSupplyTable().Insert(s.ctx, &ecoApi.BatchSupply{
		BatchKey:        batchId,
		TradableAmount:  amount,
		RetiredAmount:   "0",
		CancelledAmount: "0",
	}))
}

func (s takeSuite) getUserBalance(batchDenom string) *ecoApi.BatchBalance {
	id := s.denomToId[batchDenom]
	bal, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addr, id)
	if ormerrors.IsNotFound(err) {
		bal = &ecoApi.BatchBalance{
			BatchKey: id,
			Address:  s.addr,
			Tradable: "0",
			Retired:  "0",
			Escrowed: "0",
		}
	} else {
		assert.NilError(s.t, err)
	}
	return bal
}

func (s takeSuite) expectDec(expected, actual string) {
	actualDec, err := math.NewDecFromString(actual)
	assert.NilError(s.t, err)

	expectedDec, err := math.NewDecFromString(expected)
	assert.NilError(s.t, err)

	assert.Assert(s.t, actualDec.Cmp(expectedDec) == 0)
}
