package basket_test

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	coreapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type putSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	bob              sdk.AccAddress
	classId          string
	creditTypeAbbrev string
	batchDenom       string
	basketDenom      string
	tradableCredits  string
	res              *basket.MsgPutResponse
	err              error
}

func TestPut(t *testing.T) {
	gocuke.NewRunner(t, &putSuite{}).Path("./features/msg_put.feature").Run()
}

func (s *putSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.bob = s.addrs[1]
	s.classId = "C01"
	s.creditTypeAbbrev = "C"
	s.batchDenom = "C01-20200101-20210101-001"
	s.basketDenom = "NCT"
	s.tradableCredits = "100"
}

func (s *putSuite) ABasket() {
	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      s.basketDenom,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: basketId,
		ClassId:  s.classId,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) ABasketWithExponent(a string) {
	exponent, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      s.basketDenom,
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Exponent:         uint32(exponent),
	})
	require.NoError(s.t, err)

	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: basketId,
		ClassId:  s.classId,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) ABasketWithDenom(a string) {
	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      a,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: basketId,
		ClassId:  s.classId,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) ABasketWithAllowedCreditClass(a string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(a)

	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      s.basketDenom,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: basketId,
		ClassId:  a,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) ABasketWithMinimumStartDate(a string) {
	minStartDate, err := types.ParseDate("start date", a)
	require.NoError(s.t, err)

	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      s.basketDenom,
		CreditTypeAbbrev: s.creditTypeAbbrev,
		DateCriteria: &api.DateCriteria{
			MinStartDate: timestamppb.New(minStartDate),
		},
	})
	require.NoError(s.t, err)

	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: basketId,
		ClassId:  s.classId,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) ABasketWithStartDateWindow(a string) {
	startDateWindow, err := strconv.ParseInt(a, 10, 32)
	require.NoError(s.t, err)

	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      s.basketDenom,
		CreditTypeAbbrev: s.creditTypeAbbrev,
		DateCriteria: &api.DateCriteria{
			StartDateWindow: &durationpb.Duration{
				Seconds: startDateWindow,
			},
		},
	})
	require.NoError(s.t, err)

	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: basketId,
		ClassId:  s.classId,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) ABasketWithYearsInThePast(a string) {
	yearsInThePast, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      s.basketDenom,
		CreditTypeAbbrev: s.creditTypeAbbrev,
		DateCriteria: &api.DateCriteria{
			YearsInThePast: uint32(yearsInThePast),
		},
	})
	require.NoError(s.t, err)

	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: basketId,
		ClassId:  s.classId,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) AliceOwnsCredits() {
	classId := core.GetClassIdFromBatchDenom(s.batchDenom)
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(classId)

	classKey, err := s.coreStore.ClassTable().InsertReturningID(s.ctx, &coreapi.Class{
		Id:               classId,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	projectKey, err := s.coreStore.ProjectTable().InsertReturningID(s.ctx, &coreapi.Project{
		ClassKey: classKey,
	})
	require.NoError(s.t, err)

	batchKey, err := s.coreStore.BatchTable().InsertReturningID(s.ctx, &coreapi.Batch{
		ProjectKey: projectKey,
		Denom:      s.batchDenom,
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, &coreapi.BatchBalance{
		BatchKey: batchKey,
		Address:  s.alice,
		Tradable: s.tradableCredits,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) AliceOwnsCreditAmount(a string) {
	classId := core.GetClassIdFromBatchDenom(s.batchDenom)
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(classId)

	classKey, err := s.coreStore.ClassTable().InsertReturningID(s.ctx, &coreapi.Class{
		Id:               classId,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	projectKey, err := s.coreStore.ProjectTable().InsertReturningID(s.ctx, &coreapi.Project{
		ClassKey: classKey,
	})
	require.NoError(s.t, err)

	batchKey, err := s.coreStore.BatchTable().InsertReturningID(s.ctx, &coreapi.Batch{
		ProjectKey: projectKey,
		Denom:      s.batchDenom,
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, &coreapi.BatchBalance{
		BatchKey: batchKey,
		Address:  s.alice,
		Tradable: a,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) AliceOwnsCreditsFromCreditBatch(a string) {
	classId := core.GetClassIdFromBatchDenom(a)
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(classId)

	classKey, err := s.coreStore.ClassTable().InsertReturningID(s.ctx, &coreapi.Class{
		Id:               classId,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	projectKey, err := s.coreStore.ProjectTable().InsertReturningID(s.ctx, &coreapi.Project{
		ClassKey: classKey,
	})
	require.NoError(s.t, err)

	batchKey, err := s.coreStore.BatchTable().InsertReturningID(s.ctx, &coreapi.Batch{
		ProjectKey: projectKey,
		Denom:      a,
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, &coreapi.BatchBalance{
		BatchKey: batchKey,
		Address:  s.alice,
		Tradable: s.tradableCredits,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) AliceOwnsCreditsWithStartDate(a string) {
	startDate, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	classKey, err := s.coreStore.ClassTable().InsertReturningID(s.ctx, &coreapi.Class{
		Id:               s.classId,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	pKey, err := s.coreStore.ProjectTable().InsertReturningID(s.ctx, &coreapi.Project{
		ClassKey: classKey,
	})
	require.NoError(s.t, err)

	batchKey, err := s.coreStore.BatchTable().InsertReturningID(s.ctx, &coreapi.Batch{
		ProjectKey: pKey,
		Denom:      s.batchDenom,
		StartDate:  timestamppb.New(startDate),
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, &coreapi.BatchBalance{
		BatchKey: batchKey,
		Address:  s.alice,
		Tradable: s.tradableCredits,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) TheBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
}

func (s *putSuite) AliceAttemptsToPutCreditsIntoBasket(a string) {
	amount, ok := sdk.NewIntFromString(s.tradableCredits)
	require.True(s.t, ok)

	coins := sdk.NewCoins(sdk.NewCoin(s.basketDenom, amount))

	s.bankKeeper.EXPECT().
		MintCoins(s.sdkCtx, basket.BasketSubModuleName, coins).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(s.sdkCtx, basket.BasketSubModuleName, s.alice, coins).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	_, s.err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.alice.String(),
		BasketDenom: a,
		Credits: []*basket.BasketCredit{
			{
				BatchDenom: s.batchDenom,
				Amount:     s.tradableCredits,
			},
		},
	})
}

func (s *putSuite) AliceAttemptsToPutCreditAmountIntoTheBasket(a string) {
	amount, ok := sdk.NewIntFromString(a)
	require.True(s.t, ok)

	coins := sdk.NewCoins(sdk.NewCoin(s.basketDenom, amount))

	s.bankKeeper.EXPECT().
		MintCoins(s.sdkCtx, basket.BasketSubModuleName, coins).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(s.sdkCtx, basket.BasketSubModuleName, s.alice, coins).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	_, s.err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.alice.String(),
		BasketDenom: s.basketDenom,
		Credits: []*basket.BasketCredit{
			{
				BatchDenom: s.batchDenom,
				Amount:     a,
			},
		},
	})
}

func (s *putSuite) AliceAttemptsToPutCreditAmountIntoTheBasketWithExponent(a string) {
	coins := s.calculateExpectedCoins(a)

	s.bankKeeper.EXPECT().
		MintCoins(s.sdkCtx, basket.BasketSubModuleName, coins).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(s.sdkCtx, basket.BasketSubModuleName, s.alice, coins).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.res, s.err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.alice.String(),
		BasketDenom: s.basketDenom,
		Credits: []*basket.BasketCredit{
			{
				BatchDenom: s.batchDenom,
				Amount:     a,
			},
		},
	})
}

func (s *putSuite) AliceAttemptsToPutCreditsFromCreditBatchIntoTheBasket(a string) {
	amount, ok := sdk.NewIntFromString(s.tradableCredits)
	require.True(s.t, ok)

	coins := sdk.NewCoins(sdk.NewCoin(s.basketDenom, amount))

	s.bankKeeper.EXPECT().
		MintCoins(s.sdkCtx, basket.BasketSubModuleName, coins).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(s.sdkCtx, basket.BasketSubModuleName, s.alice, coins).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	_, s.err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.alice.String(),
		BasketDenom: s.basketDenom,
		Credits: []*basket.BasketCredit{
			{
				BatchDenom: a,
				Amount:     s.tradableCredits,
			},
		},
	})
}

func (s *putSuite) BobAttemptsToPutCreditsFromCreditBatchIntoTheBasket(a string) {
	amount, ok := sdk.NewIntFromString(s.tradableCredits)
	require.True(s.t, ok)

	coins := sdk.NewCoins(sdk.NewCoin(s.basketDenom, amount))

	s.bankKeeper.EXPECT().
		MintCoins(s.sdkCtx, basket.BasketSubModuleName, coins).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(s.sdkCtx, basket.BasketSubModuleName, s.bob, coins).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	_, s.err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.bob.String(),
		BasketDenom: s.basketDenom,
		Credits: []*basket.BasketCredit{
			{
				BatchDenom: a,
				Amount:     s.tradableCredits,
			},
		},
	})
}

func (s *putSuite) AliceAttemptsToPutCreditsIntoTheBasket() {
	amount, ok := sdk.NewIntFromString(s.tradableCredits)
	require.True(s.t, ok)

	coins := sdk.NewCoins(sdk.NewCoin(s.basketDenom, amount))

	s.bankKeeper.EXPECT().
		MintCoins(s.sdkCtx, basket.BasketSubModuleName, coins).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(s.sdkCtx, basket.BasketSubModuleName, s.alice, coins).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	_, s.err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.alice.String(),
		BasketDenom: s.basketDenom,
		Credits: []*basket.BasketCredit{
			{
				BatchDenom: s.batchDenom,
				Amount:     s.tradableCredits,
			},
		},
	})
}

func (s *putSuite) TheBasketHasACreditBalanceWithAmount(a string) {
	basket, err := s.stateStore.BasketTable().GetByBasketDenom(s.ctx, s.basketDenom)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BasketBalanceTable().Get(s.ctx, basket.Id, s.batchDenom)
	require.NoError(s.t, err)

	require.Equal(s.t, a, balance.Balance)
}

func (s *putSuite) TheBasketTokenHasATotalSupplyWithAmount(a string) {
	basket, err := s.stateStore.BasketTable().GetByBasketDenom(s.ctx, s.basketDenom)
	require.NoError(s.t, err)

	amount, err := strconv.ParseInt(a, 10, 32)
	require.NoError(s.t, err)

	coin := sdk.NewInt64Coin(s.basketDenom, amount)

	s.bankKeeper.EXPECT().
		GetSupply(s.sdkCtx, basket.BasketDenom).
		Return(coin).
		Times(1)

	supply := s.bankKeeper.GetSupply(s.sdkCtx, s.basketDenom)
	require.Equal(s.t, coin, supply)
}

func (s *putSuite) AliceHasACreditBalanceWithAmount(a string) {
	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.alice, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, a, balance.Tradable)
}

func (s *putSuite) AliceHasABasketTokenBalanceWithAmount(a string) {
	basket, err := s.stateStore.BasketTable().GetByBasketDenom(s.ctx, s.basketDenom)
	require.NoError(s.t, err)

	amount, err := strconv.ParseInt(a, 10, 32)
	require.NoError(s.t, err)

	coin := sdk.NewInt64Coin(basket.BasketDenom, amount)

	s.bankKeeper.EXPECT().
		GetBalance(s.sdkCtx, s.alice, basket.BasketDenom).
		Return(coin).
		AnyTimes()

	balance := s.bankKeeper.GetBalance(s.sdkCtx, s.alice, basket.BasketDenom)
	require.Equal(s.t, coin, balance)
}

func (s *putSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *putSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *putSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *putSuite) ExpectTheResponse(a gocuke.DocString) {
	res := &basket.MsgPutResponse{}
	err := jsonpb.UnmarshalString(a.Content, res)
	require.NoError(s.t, err)

	require.Equal(s.t, res, s.res)
}

func (s *putSuite) calculateExpectedCoins(amount string) sdk.Coins {
	basket, err := s.stateStore.BasketTable().GetByBasketDenom(s.ctx, s.basketDenom)
	require.NoError(s.t, err)

	dec, err := math.NewPositiveFixedDecFromString(amount, basket.Exponent)
	if err != nil && strings.Contains(err.Error(), "exceeds maximum decimal places") {
		// expected coins irrelevant if amount exceeds maximum decimal places
		return sdk.Coins{}
	}
	require.NoError(s.t, err)

	tokenAmt, err := math.NewDecFinite(1, int32(basket.Exponent)).MulExact(dec)
	require.NoError(s.t, err)

	amtInt, err := tokenAmt.BigInt()
	require.NoError(s.t, err)

	return sdk.Coins{sdk.NewCoin(s.basketDenom, sdk.NewIntFromBigInt(amtInt))}
}
