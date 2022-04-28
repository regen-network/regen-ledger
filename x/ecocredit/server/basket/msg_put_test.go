package basket_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	coreapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type putSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	classId          string
	creditTypeAbbrev string
	batchDenom       string
	basketDenom      string
	tradableCredits  string
	err              error
}

func TestPut(t *testing.T) {
	gocuke.NewRunner(t, &putSuite{}).Path("./features/msg_put.feature").Run()
}

func (s *putSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr

	// values used if not specified
	s.classId = "C01"
	s.creditTypeAbbrev = "C"
	s.batchDenom = "C01-20200101-20210101-001"
	s.basketDenom = "NCT"
	s.tradableCredits = "100"
}

func (s *putSuite) ABasketWithDenom(a string) {
	id, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      a,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: id,
		ClassId:  s.classId,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) ABasketWithDenomAndAllowedCreditClass(a string, b string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(b)

	id, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      a,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: id,
		ClassId:  b,
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

func (s *putSuite) ACreditBatchWithDenom(a string) {
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

	err = s.coreStore.BatchTable().Insert(s.ctx, &coreapi.Batch{
		ProjectKey: projectKey,
		Denom:      a,
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

func (s *putSuite) AliceOwnsCreditAmountFromCreditBatch(a string, b string) {
	classId := core.GetClassIdFromBatchDenom(b)
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
		Denom:      b,
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, &coreapi.BatchBalance{
		BatchKey: batchKey,
		Address:  s.alice,
		Tradable: a,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) AliceOwnsCreditsWithStartDate(a string) {
	startDate, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	key, err := s.coreStore.ClassTable().InsertReturningID(s.ctx, &coreapi.Class{
		Id:               s.classId,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	key, err = s.coreStore.ProjectTable().InsertReturningID(s.ctx, &coreapi.Project{
		ClassKey: key,
	})
	require.NoError(s.t, err)

	key, err = s.coreStore.BatchTable().InsertReturningID(s.ctx, &coreapi.Batch{
		ProjectKey: 1,
		Denom:      s.batchDenom,
		StartDate:  timestamppb.New(startDate),
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, &coreapi.BatchBalance{
		BatchKey: key,
		Address:  s.addr,
		Tradable: s.tradableCredits,
	})
	require.NoError(s.t, err)
}

func (s *putSuite) TheBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))
}

func (s *putSuite) AliceAttemptsToPutCreditsIntoBasket(a string) {
	gmAny := gomock.Any()

	s.bankKeeper.EXPECT().
		MintCoins(gmAny, basket.BasketSubModuleName, gmAny).
		Return(nil).AnyTimes()

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(gmAny, basket.BasketSubModuleName, s.addr, gmAny).
		Return(nil).AnyTimes()

	_, s.err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.alice.String(),
		BasketDenom: a,
	})
}

func (s *putSuite) AliceAttemptsToPutCreditAmountIntoBasket(a string, b string) {
	gmAny := gomock.Any()

	s.bankKeeper.EXPECT().
		MintCoins(gmAny, basket.BasketSubModuleName, gmAny).
		Return(nil).AnyTimes()

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(gmAny, basket.BasketSubModuleName, s.addr, gmAny).
		Return(nil).AnyTimes()

	_, s.err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.alice.String(),
		BasketDenom: b,
		Credits: []*basket.BasketCredit{
			{
				BatchDenom: s.batchDenom,
				Amount:     a,
			},
		},
	})
}

func (s *putSuite) AliceAttemptsToPutCreditsFromCreditBatchIntoBasket(a string, b string) {
	gmAny := gomock.Any()

	s.bankKeeper.EXPECT().
		MintCoins(gmAny, basket.BasketSubModuleName, gmAny).
		Return(nil).AnyTimes()

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(gmAny, basket.BasketSubModuleName, s.addr, gmAny).
		Return(nil).AnyTimes()

	_, s.err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.alice.String(),
		BasketDenom: b,
		Credits: []*basket.BasketCredit{
			{
				BatchDenom: a,
				Amount:     s.tradableCredits,
			},
		},
	})
}

func (s *putSuite) AliceAttemptsToPutTheCreditsIntoTheBasket() {
	gmAny := gomock.Any()
	tokenInt, _ := sdk.NewIntFromString(s.tradableCredits)
	tokenAmount := sdk.NewCoins(sdk.NewCoin(s.basketDenom, tokenInt))

	s.bankKeeper.EXPECT().
		MintCoins(gmAny, basket.BasketSubModuleName, tokenAmount).
		Return(nil).AnyTimes()

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(gmAny, basket.BasketSubModuleName, s.addr, tokenAmount).
		Return(nil).AnyTimes()

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

func (s *putSuite) TheBasketHasABalanceOfCreditAmount(a string, b string) {
	basket, err := s.stateStore.BasketTable().GetByBasketDenom(s.ctx, a)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BasketBalanceTable().Get(s.ctx, basket.Id, s.batchDenom)
	require.NoError(s.t, err)

	require.Equal(s.t, b, balance.Balance)
}

func (s *putSuite) AliceHasABalanceOfCreditAmount(a string) {
	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.alice, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, a, balance.Tradable)
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
