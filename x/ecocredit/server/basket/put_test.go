package basket_test

import (
	"context"
	"testing"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

//func TestPut(t *testing.T) {
//	t.Parallel()
//
//	basketDenom := "BASKET"
//	basketDenom2 := "BASKET2"
//	basketDenom3 := "BASKET3"
//	classId := "C02"
//	projectId := "P01"
//
//	start, err := time.Parse("2006-01-02", "2020-01-01")
//	require.NoError(t, err)
//	startDate := timestamppb.New(start)
//	gogoStartDate := gogotypes.Timestamp{Seconds: startDate.Seconds, Nanos: startDate.Nanos}
//	end, err := time.Parse("2006-01-02", "2021-01-01")
//	require.NoError(t, err)
//	endDate := timestamppb.New(end)
//	gogoEndDate := gogotypes.Timestamp{Seconds: endDate.Seconds, Nanos: endDate.Nanos}
//	denom, err := ecocredit.FormatDenom(classId, 1, &start, &end)
//	require.NoError(t, err)
//	s := setupBase(t)
//	testClassInfo := core.ClassInfo{
//		Id:         1,
//		Name:       classId,
//		Admin:      sdk.AccAddress("somebody"),
//		CreditType: "C",
//	}
//	classInfoRes := core.QueryClassInfoResponse{Info: &testClassInfo}
//	testProjectInfo := core.ProjectInfo{
//		Id:              1,
//		Name:            projectId,
//		ClassId:         1,
//		Admin:           sdk.AccAddress("somebody"),
//		ProjectLocation: "US-NY",
//	}
//	projectInfoRes := core.QueryProjectInfoResponse{Info: &testProjectInfo}
//	testBatchInfo := core.BatchInfo{
//		Id:         1,
//		ProjectId:  1,
//		BatchDenom: denom,
//		StartDate:  &gogoStartDate,
//		EndDate:    &gogoEndDate,
//	}
//	batchInfoRes := core.QueryBatchInfoResponse{Info: &testBatchInfo}
//
//	err = s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
//		BasketDenom:       basketDenom,
//		Name:              basketDenom,
//		DisableAutoRetire: true,
//		CreditTypeAbbrev:  "C",
//		DateCriteria:      &api.DateCriteria{MinStartDate: timestamppb.New(startDate.AsTime())},
//		Exponent:          6,
//	})
//	require.NoError(t, err)
//
//	var dur time.Duration = 500000000000000000
//	validStartDateWindow := start.Add(-dur)
//	err = s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
//		BasketDenom:       basketDenom2,
//		Name:              basketDenom2,
//		DisableAutoRetire: true,
//		CreditTypeAbbrev:  "C",
//		DateCriteria:      &api.DateCriteria{StartDateWindow: durationpb.New(dur)},
//		Exponent:          6,
//	})
//	require.NoError(t, err)
//	validYearsInThePast := uint32(10)
//	err = s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
//		BasketDenom:       basketDenom3,
//		Name:              basketDenom3,
//		DisableAutoRetire: true,
//		CreditTypeAbbrev:  "C",
//		DateCriteria:      &api.DateCriteria{YearsInThePast: validYearsInThePast},
//		Exponent:          6,
//	})
//	require.NoError(t, err)
//	basketDenomToId := make(map[string]uint64)
//	basketDenomToId[basketDenom] = 1
//	basketDenomToId[basketDenom2] = 2
//	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
//		BasketId: 1,
//		ClassId:  classId,
//	})
//	require.NoError(t, err)
//	err = s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
//		BasketId: 2,
//		ClassId:  classId,
//	})
//	require.NoError(t, err)
//
//	s.sdkCtx = s.sdkCtx.WithBlockTime(startDate)
//	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
//
//	testCases := []struct {
//		name                string
//		startingBalance     string
//		msg                 basket.MsgPut
//		expectedCredits     []*basket.BasketCredit
//		expectedBasketCoins string
//		expectCalls         func()
//		errMsg              string
//	}{
//		{
//			name:            "valid case",
//			startingBalance: "100000000",
//			msg: basket.MsgPut{
//				Owner:       s.addr.String(),
//				BasketDenom: basketDenom,
//				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
//			},
//			expectedCredits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
//			expectedBasketCoins: "2000000", // 2 million
//			expectCalls: func() {
//				s.ecocreditKeeper.EXPECT().
//					BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: denom}).
//					Return(&batchInfoRes, nil)
//
//				s.ecocreditKeeper.EXPECT().
//					ProjectInfo(s.ctx, &core.QueryProjectInfoRequest{ProjectId: projectId}).
//					Return(&projectInfoRes, nil)
//
//				s.ecocreditKeeper.EXPECT().
//					ClassInfo(s.ctx, &core.QueryClassInfoRequest{ClassId: classId}).
//					Return(&classInfoRes, nil)
//
//				coinAward := sdk.NewCoins(sdk.NewCoin(basketDenom, sdk.NewInt(2_000_000)))
//				s.bankKeeper.EXPECT().
//					MintCoins(s.sdkCtx, basket.BasketSubModuleName, coinAward).
//					Return(nil)
//
//				s.bankKeeper.EXPECT().
//					SendCoinsFromModuleToAccount(s.sdkCtx, basket.BasketSubModuleName, s.addr, coinAward).
//					Return(nil)
//			},
//		},
//		{
//			name:            "valid case - basket with existing balance",
//			startingBalance: "100000000",
//			msg: basket.MsgPut{
//				Owner:       s.addr.String(),
//				BasketDenom: basketDenom,
//				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "1"}},
//			},
//			expectedCredits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "3"}},
//			expectedBasketCoins: "1000000", // 1 million
//			expectCalls: func() {
//				s.ecocreditKeeper.EXPECT().
//					BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: denom}).
//					Return(&batchInfoRes, nil)
//
//				s.ecocreditKeeper.EXPECT().
//					ProjectInfo(s.ctx, &core.QueryProjectInfoRequest{ProjectId: projectId}).
//					Return(&projectInfoRes, nil)
//
//				s.ecocreditKeeper.EXPECT().
//					ClassInfo(s.ctx, &core.QueryClassInfoRequest{ClassId: classId}).
//					Return(&classInfoRes, nil)
//
//				coinAward := sdk.NewCoins(sdk.NewCoin(basketDenom, sdk.NewInt(1_000_000)))
//				s.bankKeeper.EXPECT().
//					MintCoins(s.sdkCtx, basket.BasketSubModuleName, coinAward).
//					Return(nil)
//
//				s.bankKeeper.EXPECT().
//					SendCoinsFromModuleToAccount(s.sdkCtx, basket.BasketSubModuleName, s.addr, coinAward).
//					Return(nil)
//			},
//		},
//		{
//			name:            "valid case - basket 2 with rolling window",
//			startingBalance: "100000000",
//			msg: basket.MsgPut{
//				Owner:       s.addr.String(),
//				BasketDenom: basketDenom2,
//				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
//			},
//			expectedCredits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
//			expectedBasketCoins: "2000000", // 2 million
//			expectCalls: func() {
//				s.ecocreditKeeper.EXPECT().
//					BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: denom}).
//					Return(&batchInfoRes, nil)
//
//				s.ecocreditKeeper.EXPECT().
//					ProjectInfo(s.ctx, &core.QueryProjectInfoRequest{ProjectId: projectId}).
//					Return(&projectInfoRes, nil)
//
//				s.ecocreditKeeper.EXPECT().
//					ClassInfo(s.ctx, &core.QueryClassInfoRequest{ClassId: classId}).
//					Return(&classInfoRes, nil)
//
//				coinAward := sdk.NewCoins(sdk.NewCoin(basketDenom2, sdk.NewInt(2_000_000)))
//				s.bankKeeper.EXPECT().
//					MintCoins(s.sdkCtx, basket.BasketSubModuleName, coinAward).
//					Return(nil)
//
//				s.bankKeeper.EXPECT().
//					SendCoinsFromModuleToAccount(s.sdkCtx, basket.BasketSubModuleName, s.addr, coinAward).
//					Return(nil)
//			},
//		},
//		{
//			name:            "insufficient funds",
//			startingBalance: "1",
//			msg: basket.MsgPut{
//				Owner:       s.addr.String(),
//				BasketDenom: basketDenom,
//				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
//			},
//			expectedBasketCoins: "2000000", // 2 million
//			expectCalls: func() {
//				s.ecocreditKeeper.EXPECT().
//					BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: denom}).
//					Return(&batchInfoRes, nil)
//
//				s.ecocreditKeeper.EXPECT().
//					ProjectInfo(s.ctx, &core.QueryProjectInfoRequest{ProjectId: projectId}).
//					Return(&projectInfoRes, nil)
//
//				s.ecocreditKeeper.EXPECT().
//					ClassInfo(s.ctx, &core.QueryClassInfoRequest{ClassId: classId}).
//					Return(&classInfoRes, nil)
//
//			},
//			errMsg: basketserver.ErrInsufficientCredits.Error(),
//		},
//		{
//			name:            "basket not found",
//			startingBalance: "1",
//			msg: basket.MsgPut{
//				Owner:       s.addr.String(),
//				BasketDenom: "FooBar",
//				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
//			},
//			expectedBasketCoins: "2000000", // 2 million
//			expectCalls: func() {
//			},
//			errMsg: ormerrors.NotFound.Error(),
//		},
//		{
//			name:            "batch not found",
//			startingBalance: "20",
//			msg: basket.MsgPut{
//				Owner:       s.addr.String(),
//				BasketDenom: basketDenom,
//				Credits:     []*basket.BasketCredit{{BatchDenom: "FooBarBaz", Amount: "2"}},
//			},
//			expectedBasketCoins: "2000000", // 2 million
//			expectCalls: func() {
//				s.ecocreditKeeper.EXPECT().
//					BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: "FooBarBaz"}).
//					Return(nil, orm.ErrNotFound)
//			},
//			errMsg: orm.ErrNotFound.Error(),
//		},
//		//{
//		//	name:            "class not allowed",
//		//	startingBalance: "100000000",
//		//	msg: basket.MsgPut{
//		//		Owner:       s.addr.String(),
//		//		BasketDenom: basketDenom,
//		//		Credits:     []*basket.BasketCredit{{BatchDenom: "blah", Amount: "2"}},
//		//	},
//		//	expectedBasketCoins: "2000000", // 2 million
//		//	expectCalls: func() {
//		//		badInfo := *batchInfoRes.Info
//		//		badInfo.ClassId = "blah01"
//		//		ecocreditKeeper.EXPECT().
//		//			BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: "blah"}).
//		//			Return(&ecocredit.QueryBatchInfoResponse{Info: &badInfo}, nil)
//		//	},
//		//	errMsg: "credit class blah01 is not allowed in this basket",
//		//},
//		{
//			name:            "wrong credit type",
//			startingBalance: "100000000",
//			msg: basket.MsgPut{
//				Owner:       s.addr.String(),
//				BasketDenom: basketDenom,
//				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
//			},
//			expectedBasketCoins: "2000000", // 2 million
//			expectCalls: func() {
//				s.ecocreditKeeper.EXPECT().
//					BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: denom}).
//					Return(&batchInfoRes, nil)
//
//				s.ecocreditKeeper.EXPECT().
//					ProjectInfo(s.ctx, &core.QueryProjectInfoRequest{ProjectId: projectId}).
//					Return(&projectInfoRes, nil)
//
//				badClass := *classInfoRes.Info
//				badClass.CreditType = "FOO"
//				s.ecocreditKeeper.EXPECT().
//					ClassInfo(s.ctx, &core.QueryClassInfoRequest{ClassId: classId}).
//					Return(&core.QueryClassInfoResponse{Info: &badClass}, nil)
//			},
//			errMsg: "cannot use credit of type FOO in a basket that requires credit type C",
//		},
//		{
//			name:            "batch out of time window",
//			startingBalance: "100000000",
//			msg: basket.MsgPut{
//				Owner:       s.addr.String(),
//				BasketDenom: basketDenom,
//				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
//			},
//			expectedBasketCoins: "2000000", // 2 million
//			expectCalls: func() {
//				badTime, err := time.Parse("2006-01-02", "1984-01-01")
//				require.NoError(t, err)
//				badTimeInfo := *batchInfoRes.Info
//				gogoBadTime, err := gogotypes.TimestampProto(badTime)
//				assert.NilError(t, err)
//				badTimeInfo.StartDate = gogoBadTime
//				s.ecocreditKeeper.EXPECT().
//					BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: denom}).
//					Return(&core.QueryBatchInfoResponse{Info: &badTimeInfo}, nil)
//
//			},
//			errMsg: "cannot put a credit from a batch with start date",
//		},
//		{
//			name:            "batch outside of start date window",
//			startingBalance: "100000000",
//			msg: basket.MsgPut{
//				Owner:       s.addr.String(),
//				BasketDenom: basketDenom2,
//				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
//			},
//			expectedBasketCoins: "2000000", // 2 million
//			expectCalls: func() {
//				badTimeInfo := *batchInfoRes.Info
//				bogusDur := time.Duration(999999999999999)
//				badTime := validStartDateWindow.Add(-bogusDur)
//				gogoBadTime, err := gogotypes.TimestampProto(badTime)
//				assert.NilError(t, err)
//				badTimeInfo.StartDate = gogoBadTime
//				s.ecocreditKeeper.EXPECT().
//					BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: denom}).
//					Return(&core.QueryBatchInfoResponse{Info: &badTimeInfo}, nil)
//
//			},
//			errMsg: "cannot put a credit from a batch with start date",
//		},
//		{
//			name:            "batch outside of years in the past",
//			startingBalance: "100000000",
//			msg: basket.MsgPut{
//				Owner:       s.addr.String(),
//				BasketDenom: basketDenom3,
//				Credits:     []*basket.BasketCredit{{BatchDenom: denom, Amount: "2"}},
//			},
//			expectedBasketCoins: "2000000", // 2 million
//			expectCalls: func() {
//				badTimeInfo := *batchInfoRes.Info
//				badYear := int(validYearsInThePast + 10)
//				badTime := time.Date(badYear, 1, 1, 0, 0, 0, 0, time.UTC)
//				gogoBadTime, err := gogotypes.TimestampProto(badTime)
//				assert.NilError(t, err)
//				badTimeInfo.StartDate = gogoBadTime
//				s.ecocreditKeeper.EXPECT().
//					BatchInfo(s.ctx, &core.QueryBatchInfoRequest{BatchDenom: denom}).
//					Return(&core.QueryBatchInfoResponse{Info: &badTimeInfo}, nil)
//
//			},
//			errMsg: "cannot put a credit from a batch with start date",
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			tc.expectCalls()
//			userFunds, err := math.NewDecFromString(tc.startingBalance)
//			require.NoError(t, err)
//			s.k
//			ecocredit.SetDecimal(legacyStore, tradKey, userFunds)
//			res, err := s.k.Put(s.ctx, &tc.msg)
//			if tc.errMsg == "" { //  no error!
//				require.NoError(t, err)
//				require.Equal(t, res.AmountReceived, tc.expectedBasketCoins)
//				for _, credit := range tc.msg.Credits {
//					assertUserSentCredits(t, s.ctx)
//				}
//				for _, credit := range tc.expectedCredits {
//					assertBasketHasCredits(t, s.ctx, credit, basketDenomToId[tc.msg.BasketDenom], s.stateStore.BasketBalanceTable())
//				}
//			} else {
//				require.Error(t, err)
//				require.Contains(t, err.Error(), tc.errMsg)
//			}
//		})
//	}
//}

func TestPut_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom := "C01-20200101-20220101-001"
	basketDenom := setupPutTest(t, s, []string{"C01"}, "basket", "C", 6,
		&api.DateCriteria{YearsInThePast: 3})
	// should be able to put into basket
	gogoTs := gogotypes.TimestampNow()
	s.ecocreditKeeper.EXPECT().BatchInfo(gomock.Any(), gomock.Any()).Return(&core.QueryBatchInfoResponse{Info: &core.BatchInfo{
		Id:         1,
		BatchDenom: batchDenom,
		Metadata:   "",
		StartDate:  gogoTs,
		EndDate:    nil,
	}}, nil).Times(1)
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), "ecocredit-basket", gomock.Any()).Return(nil)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	res, err := s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: basketDenom,
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom, Amount: "10"},
		},
	})
	assert.NilError(t, err)
	// multiplier == 10^6 = 1,000,000
	// put 10 coins == 1,000,000 * 10 = 10M
	assert.Equal(t, "10000000", res.AmountReceived)
}

func TestPut_BasketNotFound(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, err := s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: "foo",
		Credits:     nil,
	})
	assert.ErrorContains(t, err, sdkerrors.ErrNotFound.Error())
}

func TestPut_BatchNotFound(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	basketDenom := setupPutTest(t, s, []string{"C01"}, "foo", "C", 6, &api.DateCriteria{YearsInThePast: 10})
	s.ecocreditKeeper.EXPECT().BatchInfo(gomock.Any(), gomock.Any()).Return(nil, ormerrors.NotFound)
	_, err := s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: basketDenom,
		Credits: []*basket.BasketCredit{
			{BatchDenom: "C05-00000000-00000000-001", Amount: "15"},
		},
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}

func TestPut_Criteria(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	basketDenom := setupPutTest(t, s, []string{"C01"}, "foo", "C", 6, &api.DateCriteria{YearsInThePast: 10})
	batchDenom := "C01-00000000-00000000-0001"
	s.ecocreditKeeper.EXPECT().BatchInfo(gomock.Any(), gomock.Any()).Return(&core.QueryBatchInfoResponse{
		Info: &core.BatchInfo{}},
		nil)
	res, err := s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addr.String(),
		BasketDenom: basketDenom,
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom, Amount: "10"},
		},
	})
}

func setupPutTest(t *testing.T, s *baseSuite, basketClasses []string, basketName, ctAbbrev string, exp uint32, dateCriteria *api.DateCriteria) string {
	denom, _, err := basket.BasketDenom(basketName, ctAbbrev, exp)
	assert.NilError(t, err)
	assert.NilError(t, s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom:      denom,
		Name:             basketName,
		CreditTypeAbbrev: ctAbbrev,
		DateCriteria:     dateCriteria,
		Exponent:         exp,
		Curator:          "",
	}))
	for _, class := range basketClasses {
		assert.NilError(t, s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{BasketId: 1, ClassId: class}))
	}
	return denom
}

func assertBasketHasCredits(t *testing.T, ctx context.Context, credit *basket.BasketCredit, basketID uint64, balanceTable api.BasketBalanceTable) {
	bal, err := balanceTable.Get(ctx, basketID, credit.BatchDenom)
	require.NoError(t, err)
	require.Equal(t, bal.Balance, credit.Amount)
}

func assertUserSentCredits(t *testing.T, ctx context.Context, oldBalance math.Dec, amountSent string, user sdk.AccAddress, store ecoApi.BatchBalanceTable, batchId uint64) {
	userBal, err := store.Get(ctx, user, batchId)
	assert.NilError(t, err)
	amtSent, err := math.NewDecFromString(amountSent)
	require.NoError(t, err)

	tradable, err := math.NewDecFromString(userBal.Tradable)
	assert.NilError(t, err)

	expectedBalance, err := tradable.Add(amtSent)
	assert.NilError(t, err)

	require.True(t, expectedBalance.Equal(oldBalance))
}
