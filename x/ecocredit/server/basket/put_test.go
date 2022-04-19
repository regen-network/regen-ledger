package basket_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	ecocreditapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func TestPut_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	batchDenom, classId := "C01-0000000-0000000-001", "C01"
	userStartingBalance, basketStartingBalance, amtToDeposit := math.NewDecFromInt64(10), math.NewDecFromInt64(0), math.NewDecFromInt64(3)
	insertBasket(t, s, "foo", "basket", "C", &api.DateCriteria{YearsInThePast: 3}, []string{classId})
	insertBatch(t, s, batchDenom, timestamppb.Now())
	insertBatchBalance(t, s, s.addr, 1, userStartingBalance.String())
	insertClassInfo(t, s, "C01", "C")
	s.bankKeeper.EXPECT().MintCoins(gmAny, gmAny, gmAny).Return(nil).Times(2)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gmAny, gmAny, gmAny, gmAny).Return(nil).Times(2)

	// can put 3 credits in basket
	res, err := s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom, Amount: amtToDeposit.String()},
		},
	})
	assert.NilError(t, err)
	assert.Equal(t, res.AmountReceived, "3000000") // 3 credits 10^6 * 3 = 3M
	assertCreditsDeposited(t, s, userStartingBalance, basketStartingBalance, amtToDeposit, s.addr, 1, 1, batchDenom)

	// can put 3 more credits in basket
	res, err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom, Amount: amtToDeposit.String()},
		},
	})
	assert.NilError(t, err)
	assert.Equal(t, res.AmountReceived, "3000000") // 3 credits 10^6 * 3
	assertCreditsDeposited(t, s, math.NewDecFromInt64(7), math.NewDecFromInt64(3), amtToDeposit, s.addr, 1, 1, batchDenom)
}

func TestPut_BasketNotFound(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	_, err := s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "FOO",
		Credits:     nil,
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}

func TestPut_BatchNotFound(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom, classId := "C01-0000000-0000000-001", "C01"
	userStartingBalance, _ := math.NewDecFromInt64(10), math.NewDecFromInt64(0)
	insertBasket(t, s, "foo", "basket", "C", &api.DateCriteria{YearsInThePast: 3}, []string{classId})

	// can put all credits in basket
	_, err := s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom, Amount: userStartingBalance.String()},
		},
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}

func TestPut_YearsIntoPast(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	batchDenom, classId := "C01-0000000-0000000-001", "C01"
	insertBasket(t, s, "foo", "basket", "C", &api.DateCriteria{YearsInThePast: 3}, []string{classId})
	currentTime, err := time.Parse("2006", "2020")
	assert.NilError(t, err)
	s.sdkCtx = s.sdkCtx.WithBlockTime(currentTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// too long ago should fail
	fourYearsAgo, err := time.Parse("2006", "2016")
	assert.NilError(t, err)
	insertBatch(t, s, batchDenom, timestamppb.New(fourYearsAgo))
	_, err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom, Amount: "10"},
		},
	})
	assert.ErrorContains(t, err, "basket that requires an earliest start date")

	// exactly 3 years into the past should work
	threeYearsAgo, err := time.Parse("2006", "2017")
	assert.NilError(t, err)
	otherBatchDenom := "C01-000000-0000000-002"
	insertBatch(t, s, otherBatchDenom, timestamppb.New(threeYearsAgo))
	insertBatchBalance(t, s, s.addr, 2, "10")
	insertClassInfo(t, s, "C01", "C")
	s.bankKeeper.EXPECT().MintCoins(gmAny, gmAny, gmAny).Return(nil).Times(1)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gmAny, gmAny, gmAny, gmAny).Return(nil).Times(1)
	_, err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: otherBatchDenom, Amount: "10"},
		},
	})
	assert.NilError(t, err)
}

func TestPut_MinStartDate(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	batchDenom, classId := "C01-0000000-0000000-001", "C01"
	currentTime, err := time.Parse("2006", "2020")
	assert.NilError(t, err)

	// make a basket with min start date as 2020
	insertBasket(t, s, "foo", "basket", "C", &api.DateCriteria{MinStartDate: timestamppb.New(currentTime)}, []string{classId})
	s.sdkCtx = s.sdkCtx.WithBlockTime(currentTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// make a batch 1 year before min start date
	_2019, err := time.Parse("2006", "2019")
	assert.NilError(t, err)
	insertBatch(t, s, batchDenom, timestamppb.New(_2019))
	_, err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom, Amount: "10"},
		},
	})
	assert.ErrorContains(t, err, "basket that requires an earliest start date")

	// should pass with batch start date == min start date
	otherBatchDenom := "C01-000000-0000000-002"
	insertBatch(t, s, otherBatchDenom, timestamppb.New(currentTime))
	insertBatchBalance(t, s, s.addr, 2, "10")
	insertClassInfo(t, s, "C01", "C")
	s.bankKeeper.EXPECT().MintCoins(gmAny, gmAny, gmAny).Return(nil).Times(1)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gmAny, gmAny, gmAny, gmAny).Return(nil).Times(1)
	_, err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: otherBatchDenom, Amount: "10"},
		},
	})
	assert.NilError(t, err)
}

func TestPut_Window(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	batchDenom, classId := "C01-0000000-0000000-001", "C01"
	currentTime, err := time.Parse("2006", "2020")
	assert.NilError(t, err)

	// make a basket with StartDateWindow of 1 year. with block time forced to 2020, this means only credits 2019 and up are accepted.
	insertBasket(t, s, "foo", "basket", "C", &api.DateCriteria{StartDateWindow: &durationpb.Duration{Seconds: 31560000}}, []string{classId})
	s.sdkCtx = s.sdkCtx.WithBlockTime(currentTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// make a batch 1 year before min start date
	_2018, err := time.Parse("2006", "2018")
	assert.NilError(t, err)
	insertBatch(t, s, batchDenom, timestamppb.New(_2018))
	_, err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom, Amount: "10"},
		},
	})
	assert.ErrorContains(t, err, "basket that requires an earliest start date")

	// should pass with batch start date exactly 1 year before block time. (1 year window).
	otherBatchDenom := "C01-000000-0000000-002"
	_2019, err := time.Parse("2006", "2019")
	assert.NilError(t, err)
	insertBatch(t, s, otherBatchDenom, timestamppb.New(_2019))
	insertBatchBalance(t, s, s.addr, 2, "10")
	insertClassInfo(t, s, "C01", "C")
	s.bankKeeper.EXPECT().MintCoins(gmAny, gmAny, gmAny).Return(nil).Times(1)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gmAny, gmAny, gmAny, gmAny).Return(nil).Times(1)
	_, err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: otherBatchDenom, Amount: "10"},
		},
	})
	assert.NilError(t, err)
}

func TestPut_InsufficientCredits(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom, classId := "C01-0000000-0000000-001", "C01"
	userStartingBalance := math.NewDecFromInt64(10)
	insertBasket(t, s, "foo", "basket", "C", &api.DateCriteria{YearsInThePast: 3}, []string{classId})
	insertBatch(t, s, batchDenom, timestamppb.Now())
	insertBatchBalance(t, s, s.addr, 1, userStartingBalance.String())
	insertClassInfo(t, s, classId, "C")

	// can put 3 credits in basket
	_, err := s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom, Amount: "1000000"},
		},
	})
	assert.ErrorContains(t, err, ecocredit.ErrInsufficientCredits.Error())
}

func TestPut_ClassNotAccepted(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom := "C01-0000000-0000000-001"
	userStartingBalance := math.NewDecFromInt64(10)
	insertBasket(t, s, "foo", "basket", "C", &api.DateCriteria{YearsInThePast: 3}, []string{"C02"})
	insertBatch(t, s, batchDenom, timestamppb.Now())
	insertBatchBalance(t, s, s.addr, 1, userStartingBalance.String())

	// can put 3 credits in basket
	_, err := s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom, Amount: "1000000"},
		},
	})
	assert.ErrorContains(t, err, "credit class C01 is not allowed in this basket")
}

func TestPut_BadCreditType(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom := "C01-0000000-0000000-001"
	userStartingBalance := math.NewDecFromInt64(10)
	insertBasket(t, s, "foo", "basket", "C", &api.DateCriteria{YearsInThePast: 3}, []string{"C01"})
	insertBatch(t, s, batchDenom, timestamppb.Now())
	insertBatchBalance(t, s, s.addr, 1, userStartingBalance.String())
	insertClassInfo(t, s, "C01", "F")

	_, err := s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom, Amount: userStartingBalance.String()},
		},
	})
	assert.ErrorContains(t, err, "basket requires credit type C but a credit with type F was given")
}

func insertBasket(t *testing.T, s *baseSuite, denom, name, ctAbbrev string, criteria *api.DateCriteria, classes []string) {
	assert.NilError(t, s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom:       denom,
		Name:              name,
		DisableAutoRetire: false,
		CreditTypeAbbrev:  ctAbbrev,
		DateCriteria:      criteria,
		Exponent:          6,
		Curator:           s.addr.String(),
	}))
	for _, class := range classes {
		assert.NilError(t, s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
			BasketId: 1,
			ClassId:  class,
		}))
	}
}

func insertBatchBalance(t *testing.T, s *baseSuite, user sdk.AccAddress, batchKey uint64, amount string) {
	assert.NilError(t, s.coreStore.BatchBalanceTable().Insert(s.ctx, &ecoApi.BatchBalance{
		BatchKey: batchKey,
		Address:  user,
		Tradable: amount,
		Retired:  "",
		Escrowed: "",
	}))
}

func insertClassInfo(t *testing.T, s *baseSuite, name, creditTypeAbb string) {
	assert.NilError(t, s.coreStore.ClassInfoTable().Insert(s.ctx, &ecoApi.ClassInfo{
		Id:               name,
		Admin:            s.addr,
		Metadata:         "",
		CreditTypeAbbrev: creditTypeAbb,
	}))
}

func insertBatch(t *testing.T, s *baseSuite, batchDenom string, startDate *timestamppb.Timestamp) {
	assert.NilError(t, s.coreStore.BatchInfoTable().Insert(s.ctx, &ecoApi.BatchInfo{
		ProjectKey: 1,
		BatchDenom: batchDenom,
		Metadata:   "",
		StartDate:  startDate,
		EndDate:    nil,
	}))
}

func assertCreditsDeposited(t *testing.T, s *baseSuite, startingUserBalance, startingBasketBalance, amountPut math.Dec, user sdk.AccAddress, batchKey, basketId uint64, batchDenom string) {
	userBal, err := s.coreStore.BatchBalanceTable().Get(s.ctx, user, batchKey)
	assert.NilError(t, err)
	userTradable, err := math.NewDecFromString(userBal.Tradable)
	assert.NilError(t, err)

	basketBal, err := s.stateStore.BasketBalanceTable().Get(s.ctx, basketId, batchDenom)
	assert.NilError(t, err)
	basketBalAmt, err := math.NewDecFromString(basketBal.Balance)
	assert.NilError(t, err)

	expectedUserBal, err := startingUserBalance.Sub(amountPut)
	assert.NilError(t, err)

	expectedBasketBalance, err := startingBasketBalance.Add(amountPut)
	assert.NilError(t, err)

	assert.Check(t, expectedUserBal.Equal(userTradable))
	assert.Check(t, expectedBasketBalance.Equal(basketBalAmt))
}

type putSuite struct {
	*baseSuite
	basketDenom     string
	classId         string
	creditType      string
	batchDenom      string
	batchStartDate  *timestamppb.Timestamp
	tradableCredits string
	err             error
}

func TestPutDate(t *testing.T) {
	gocuke.NewRunner(t, &putSuite{}).Path("../../features/basket/put_date.feature").Run()
}

func (s *putSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.tradableCredits = "5"
	s.classId = "batch"
	s.creditType = "C"
}

func (s *putSuite) ACurrentBlockTimestampOf(a string) {
	blockTime, err := time.Parse("2006-01-02", a)
	assert.NilError(s.t, err)

	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
}

func (s *putSuite) ABasketWithDateCriteriaYearsIntoThePastOf(a string) {
	yearsInThePast, err := strconv.ParseUint(a, 10, 32)
	assert.NilError(s.t, err)

	s.basketDenom = "basket-" + a

	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      s.basketDenom,
		CreditTypeAbbrev: s.creditType,
		DateCriteria:     &api.DateCriteria{YearsInThePast: uint32(yearsInThePast)},
	})
	assert.NilError(s.t, err)

	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: basketId,
		ClassId:  s.classId,
	})
	assert.NilError(s.t, err)
}

func (s *putSuite) AUserOwnsCreditsFromABatchWithStartDateOf(a string) {
	startDate, err := time.Parse("2006-01-02", a)
	assert.NilError(s.t, err)

	s.batchDenom = "batch-" + a
	s.batchStartDate = timestamppb.New(startDate)

	key, err := s.coreStore.ClassInfoTable().InsertReturningID(s.ctx, &ecocreditapi.ClassInfo{
		Id:               s.classId,
		CreditTypeAbbrev: s.creditType,
	})
	assert.NilError(s.t, err)

	key, err = s.coreStore.ProjectInfoTable().InsertReturningID(s.ctx, &ecocreditapi.ProjectInfo{ClassKey: key})
	assert.NilError(s.t, err)

	key, err = s.coreStore.BatchInfoTable().InsertReturningID(s.ctx, &ecocreditapi.BatchInfo{
		ProjectKey: 1,
		BatchDenom: s.batchDenom,
		StartDate:  s.batchStartDate,
	})
	assert.NilError(s.t, err)

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, &ecocreditapi.BatchBalance{
		BatchKey: key,
		Address:  s.addr,
		Tradable: s.tradableCredits,
	})
	assert.NilError(s.t, err)
}

func (s *putSuite) TheUserAttemptsToPutTheCreditsIntoTheBasket() {
	gmAny := gomock.Any()
	tokenInt, _ := sdk.NewIntFromString(s.tradableCredits)
	tokenAmount := sdk.NewCoins(sdk.NewCoin(s.basketDenom, tokenInt))

	s.bankKeeper.EXPECT().
		MintCoins(gmAny, basket.BasketSubModuleName, tokenAmount).
		Return(nil).AnyTimes() // only called if valid start date

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(gmAny, basket.BasketSubModuleName, s.addr, tokenAmount).
		Return(nil).AnyTimes() // only called if valid start date

	_, s.err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: s.basketDenom,
		Credits: []*basket.BasketCredit{
			{
				BatchDenom: s.batchDenom,
				Amount:     s.tradableCredits,
			},
		},
	})
}

func (s *putSuite) TheCreditsArePutIntoTheBasket() {
	assert.ErrorIs(s.t, s.err, nil)
}

func (s *putSuite) TheCreditsAreNotPutIntoTheBasket() {
	assert.ErrorContains(s.t, s.err, "cannot put a credit from a batch with start date")
}
