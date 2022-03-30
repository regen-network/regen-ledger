package basket_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecocreditapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	basketserver "github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
)

func TestPut(t *testing.T) {
	basketDenom := "BASKET"
	basketDenom2 := "BASKET2"
	basketDenom3 := "BASKET3"
	classId := "C02"
	projectId := "P01"
	startDate, err := time.Parse("2006-01-02", "2020-01-01")
	require.NoError(t, err)
	endDate, err := time.Parse("2006-01-02", "2021-01-01")
	require.NoError(t, err)
	denom, err := ecocredit.FormatDenom(classId, 1, &startDate, &endDate)
	require.NoError(t, err)
	testClassInfo := ecocredit.ClassInfo{
		ClassId:  classId,
		Admin:    "somebody",
		Issuers:  nil,
		Metadata: nil,
		CreditType: &ecocredit.CreditType{
			Name:         "carbon",
			Abbreviation: "C",
			Unit:         "many carbons",
			Precision:    6,
		},
		NumBatches: 1,
	}
	classInfoRes := ecocredit.QueryClassInfoResponse{Info: &testClassInfo}
	testProjectInfo := ecocredit.ProjectInfo{
		ProjectId:       projectId,
		ClassId:         classId,
		Issuer:          "somebody",
		ProjectLocation: "US-NY",
		Metadata:        nil,
	}
	projectInfoRes := ecocredit.QueryProjectInfoResponse{Info: &testProjectInfo}
	testBatchInfo := ecocredit.BatchInfo{
		ProjectId:       projectId,
		BatchDenom:      denom,
		TotalAmount:     "1000000000000000000000000000",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &startDate,
		EndDate:         &endDate,
	}
	batchInfoRes := ecocredit.QueryBatchInfoResponse{Info: &testBatchInfo}

	s := setupBase(t)

	err = s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom:       basketDenom,
		Name:              basketDenom,
		DisableAutoRetire: true,
		CreditTypeAbbrev:  "C",
		DateCriteria:      &api.DateCriteria{MinStartDate: timestamppb.New(startDate)},
		Exponent:          6,
	})
	require.NoError(t, err)

	var dur time.Duration = 500000000000000000
	validStartDateWindow := startDate.Add(-dur)
	err = s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom:       basketDenom2,
		Name:              basketDenom2,
		DisableAutoRetire: true,
		CreditTypeAbbrev:  "C",
		DateCriteria:      &api.DateCriteria{StartDateWindow: durationpb.New(dur)},
		Exponent:          6,
	})
	require.NoError(t, err)
	validYearsInThePast := uint32(10)
	err = s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom:       basketDenom3,
		Name:              basketDenom3,
		DisableAutoRetire: true,
		CreditTypeAbbrev:  "C",
		DateCriteria:      &api.DateCriteria{YearsInThePast: validYearsInThePast},
		Exponent:          6,
	})
	require.NoError(t, err)
	basketDenomToId := make(map[string]uint64)
	basketDenomToId[basketDenom] = 1
	basketDenomToId[basketDenom2] = 2
	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: 1,
		ClassId:  classId,
	})
	require.NoError(t, err)
	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: 2,
		ClassId:  classId,
	})
	require.NoError(t, err)

	s.sdkCtx = s.sdkCtx.WithBlockTime(startDate)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	testCases := []struct {
		name                string
		startingBalance     string
		msg                 basket.MsgPut
		expectedCredits     []*basket.BasketCredit
		expectedBasketCoins string
		expectCalls         func()
		errMsg              string
	}{
		{
			name:            "valid case",
			startingBalance: "100000000",
			msg: basket.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedCredits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ProjectInfo(s.ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&classInfoRes, nil)

				coinAward := sdk.NewCoins(sdk.NewCoin(basketDenom, sdk.NewInt(2_000_000)))
				s.bankKeeper.EXPECT().
					MintCoins(s.sdkCtx, basket.BasketSubModuleName, coinAward).
					Return(nil)

				s.bankKeeper.EXPECT().
					SendCoinsFromModuleToAccount(s.sdkCtx, basket.BasketSubModuleName, s.addr, coinAward).
					Return(nil)
			},
		},
		{
			name:            "valid case - basket with existing balance",
			startingBalance: "100000000",
			msg: basket.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "1"}},
			},
			expectedCredits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "3"}},
			expectedBasketCoins: "1000000", // 1 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ProjectInfo(s.ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&classInfoRes, nil)

				coinAward := sdk.NewCoins(sdk.NewCoin(basketDenom, sdk.NewInt(1_000_000)))
				s.bankKeeper.EXPECT().
					MintCoins(s.sdkCtx, basket.BasketSubModuleName, coinAward).
					Return(nil)

				s.bankKeeper.EXPECT().
					SendCoinsFromModuleToAccount(s.sdkCtx, basket.BasketSubModuleName, s.addr, coinAward).
					Return(nil)
			},
		},
		{
			name:            "valid case - basket 2 with rolling window",
			startingBalance: "100000000",
			msg: basket.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom2,
				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedCredits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ProjectInfo(s.ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&classInfoRes, nil)

				coinAward := sdk.NewCoins(sdk.NewCoin(basketDenom2, sdk.NewInt(2_000_000)))
				s.bankKeeper.EXPECT().
					MintCoins(s.sdkCtx, basket.BasketSubModuleName, coinAward).
					Return(nil)

				s.bankKeeper.EXPECT().
					SendCoinsFromModuleToAccount(s.sdkCtx, basket.BasketSubModuleName, s.addr, coinAward).
					Return(nil)
			},
		},
		{
			name:            "insufficient funds",
			startingBalance: "1",
			msg: basket.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ProjectInfo(s.ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&classInfoRes, nil)

			},
			errMsg: basketserver.ErrInsufficientCredits.Error(),
		},
		{
			name:            "basket not found",
			startingBalance: "1",
			msg: basket.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: "FooBar",
				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
			},
			errMsg: ormerrors.NotFound.Error(),
		},
		{
			name:            "batch not found",
			startingBalance: "20",
			msg: basket.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket.BasketCredit{{BatchDenom: "FooBarBaz", Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: "FooBarBaz"}).
					Return(nil, ormerrors.NotFound)
			},
			errMsg: ormerrors.NotFound.Error(),
		},
		//{
		//	name:            "class not allowed",
		//	startingBalance: "100000000",
		//	msg: basket.MsgPut{
		//		Owner:       s.addr.String(),
		//		BasketDenom: basketDenom,
		//		Credits:     []*basket.BasketCredit{{BatchDenom: "blah", Amount: "2"}},
		//	},
		//	expectedBasketCoins: "2000000", // 2 million
		//	expectCalls: func() {
		//		badInfo := *batchInfoRes.Info
		//		badInfo.ClassId = "blah01"
		//		ecocreditKeeper.EXPECT().
		//			BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: "blah"}).
		//			Return(&ecocredit.QueryBatchInfoResponse{Info: &badInfo}, nil)
		//	},
		//	errMsg: "credit class blah01 is not allowed in this basket",
		//},
		{
			name:            "wrong credit type",
			startingBalance: "100000000",
			msg: basket.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ProjectInfo(s.ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				badClass := *classInfoRes.Info
				badClass.CreditType.Abbreviation = "FOO"
				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&ecocredit.QueryClassInfoResponse{Info: &badClass}, nil)
			},
			errMsg: "cannot use credit of type FOO in a basket that requires credit type C",
		},
		{
			name:            "batch out of time window",
			startingBalance: "100000000",
			msg: basket.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				badTime, err := time.Parse("2006-01-02", "1984-01-01")
				require.NoError(t, err)
				badTimeInfo := *batchInfoRes.Info
				badTimeInfo.StartDate = &badTime
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&ecocredit.QueryBatchInfoResponse{Info: &badTimeInfo}, nil)

			},
			errMsg: "cannot put a credit from a batch with start date",
		},
		{
			name:            "batch outside of start date window",
			startingBalance: "100000000",
			msg: basket.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom2,
				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				badTimeInfo := *batchInfoRes.Info
				bogusDur := time.Duration(999999999999999)
				badTime := validStartDateWindow.Add(-bogusDur)
				badTimeInfo.StartDate = &badTime
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&ecocredit.QueryBatchInfoResponse{Info: &badTimeInfo}, nil)

			},
			errMsg: "cannot put a credit from a batch with start date",
		},
		{
			name:            "batch outside of years in the past",
			startingBalance: "100000000",
			msg: basket.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: basketDenom3,
				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				badTimeInfo := *batchInfoRes.Info
				badYear := int(validYearsInThePast + 10)
				badTime := time.Date(badYear, 1, 1, 0, 0, 0, 0, time.UTC)
				badTimeInfo.StartDate = &badTime
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&ecocredit.QueryBatchInfoResponse{Info: &badTimeInfo}, nil)

			},
			errMsg: "cannot put a credit from a batch with start date",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.expectCalls()
			legacyStore := s.sdkCtx.KVStore(s.storeKey)
			tradKey := ecocredit.TradableBalanceKey(s.addr, ecocredit.BatchDenomT(denom))
			userFunds, err := math.NewDecFromString(tc.startingBalance)
			require.NoError(t, err)
			ecocredit.SetDecimal(legacyStore, tradKey, userFunds)
			res, err := s.k.Put(s.ctx, &tc.msg)
			if tc.errMsg == "" { //  no error!
				require.NoError(t, err)
				require.Equal(t, res.AmountReceived, tc.expectedBasketCoins)
				for _, credit := range tc.msg.Credits {
					assertUserSentCredits(t, userFunds, credit.Amount, tradKey, legacyStore)
				}
				for _, credit := range tc.expectedCredits {
					assertBasketHasCredits(t, s.ctx, credit, basketDenomToId[tc.msg.BasketDenom], s.stateStore.BasketBalanceTable())
				}
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errMsg)
			}
		})
	}
}

func assertBasketHasCredits(t *testing.T, ctx context.Context, credit *basket.BasketCredit, basketID uint64, balanceTable api.BasketBalanceTable) {
	bal, err := balanceTable.Get(ctx, basketID, credit.BatchDenom)
	require.NoError(t, err)
	require.Equal(t, bal.Balance, credit.Amount)
}

func assertUserSentCredits(t *testing.T, oldBalance math.Dec, amountSent string, balanceKey []byte, store types.KVStore) {
	amtSent, err := math.NewDecFromString(amountSent)
	require.NoError(t, err)
	currentBalance, err := ecocredit.GetDecimal(store, balanceKey)
	require.NoError(t, err)

	checkBalance, err := currentBalance.Add(amtSent)
	require.NoError(t, err)

	require.True(t, checkBalance.Equal(oldBalance))
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
	s.classId = "C01"
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

	err = s.ecocreditStore.BatchInfoTable().Insert(s.ctx, &ecocreditapi.BatchInfo{
		BatchDenom: s.batchDenom,
		StartDate:  s.batchStartDate,
	})
	assert.NilError(s.t, err)

	// TODO: migrate from legacy tradable balance store to batch balance table

	//batch, err := s.ecocreditStore.BatchInfoTable().GetByBatchDenom(s.ctx, s.batchDenom)
	//assert.NilError(s.t, err)
	//
	//err = s.ecocreditStore.BatchBalanceTable().Insert(s.ctx, &ecocreditapi.BatchBalance{
	//	Address:  s.addr,
	//	BatchId:  batch.Id,
	//	Tradable: s.tradableCredits,
	//})
	//assert.NilError(s.t, err)

	legacyStore := s.sdkCtx.KVStore(s.storeKey)
	tradKey := ecocredit.TradableBalanceKey(s.addr, ecocredit.BatchDenomT(s.batchDenom))
	userFunds, err := math.NewDecFromString(s.tradableCredits)
	require.NoError(s.t, err)

	ecocredit.SetDecimal(legacyStore, tradKey, userFunds)
}

func (s *putSuite) TheUserAttemptsToPutTheCreditsIntoTheBasket() {
	startDate := time.Unix(s.batchStartDate.Seconds, int64(s.batchStartDate.Nanos))

	s.ecocreditKeeper.EXPECT().
		BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: s.batchDenom}).
		Return(&ecocredit.QueryBatchInfoResponse{Info: &ecocredit.BatchInfo{
			BatchDenom: s.batchDenom,
			StartDate:  &startDate,
		}}, nil).AnyTimes() // only called if valid start date

	s.ecocreditKeeper.EXPECT().
		ProjectInfo(s.ctx, &ecocredit.QueryProjectInfoRequest{}).
		Return(&ecocredit.QueryProjectInfoResponse{Info: &ecocredit.ProjectInfo{
			ClassId: s.classId,
		}}, nil).AnyTimes() // only called if valid start date

	s.ecocreditKeeper.EXPECT().
		ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: s.classId}).
		Return(&ecocredit.QueryClassInfoResponse{Info: &ecocredit.ClassInfo{
			CreditType: &ecocredit.CreditType{Abbreviation: s.creditType},
		}}, nil).AnyTimes() // only called if valid start date

	any := gomock.Any()
	tokenInt, _ := sdk.NewIntFromString(s.tradableCredits)
	tokenAmount := sdk.NewCoins(sdk.NewCoin(s.basketDenom, tokenInt))

	s.bankKeeper.EXPECT().
		MintCoins(any, basket.BasketSubModuleName, tokenAmount).
		Return(nil).AnyTimes() // only called if valid start date

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(any, basket.BasketSubModuleName, s.addr, tokenAmount).
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
