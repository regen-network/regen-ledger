package basket_test

import (
	"context"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	basket2 "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type putSuite struct {
	*baseSuite
	basketDenom          string
	basketDenom2         string
	classId              string
	startDate            time.Time
	endDate              time.Time
	classInfoRes         ecocredit.QueryClassInfoResponse
	batchInfoRes         ecocredit.QueryBatchInfoResponse
	basketBalanceTbl     ormtable.Table
	validStartDateWindow time.Time
	denom                string
	basketDenomToId      map[string]uint64
}

func setupPut(t *testing.T) *putSuite {
	s := &putSuite{baseSuite: setupBase(t)}

	s.basketDenom = "BASKET"
	s.basketDenom2 = "BASKET2"
	s.classId = "C02"
	var err error
	s.startDate, err = time.Parse("2006-01-02", "2020-01-01")
	require.NoError(t, err)
	s.endDate, err = time.Parse("2006-01-02", "2021-01-01")
	require.NoError(t, err)
	s.denom, err = ecocredit.FormatDenom(s.classId, 1, &s.startDate, &s.endDate)
	require.NoError(t, err)
	testClassInfo := ecocredit.ClassInfo{
		ClassId:  s.classId,
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

	s.classInfoRes = ecocredit.QueryClassInfoResponse{Info: &testClassInfo}
	testBatchInfo := ecocredit.BatchInfo{
		ClassId:         s.classId,
		BatchDenom:      s.denom,
		Issuer:          "somebody",
		TotalAmount:     "1000000000000000000000000000",
		Metadata:        nil,
		AmountCancelled: "0",
		StartDate:       &s.startDate,
		EndDate:         &s.endDate,
		ProjectLocation: "US-NY",
	}

	s.batchInfoRes = ecocredit.QueryBatchInfoResponse{Info: &testBatchInfo}

	basketTbl := s.db.GetTable(&basketv1.Basket{})
	err = basketTbl.Insert(s.ctx, &basketv1.Basket{
		BasketDenom:       s.basketDenom,
		DisableAutoRetire: true,
		CreditTypeName:    "carbon",
		DateCriteria:      &basketv1.DateCriteria{Sum: &basketv1.DateCriteria_MinStartDate{MinStartDate: timestamppb.New(s.startDate)}},
		Exponent:          6,
	})
	require.NoError(t, err)
	s.basketBalanceTbl = s.db.GetTable(&basketv1.BasketBalance{})
	var dur time.Duration = 500000000000000000
	s.validStartDateWindow = s.startDate.Add(-dur)
	err = basketTbl.Insert(s.ctx, &basketv1.Basket{
		BasketDenom:       s.basketDenom2,
		DisableAutoRetire: true,
		CreditTypeName:    "carbon",
		DateCriteria:      &basketv1.DateCriteria{Sum: &basketv1.DateCriteria_StartDateWindow{StartDateWindow: durationpb.New(dur)}},
		Exponent:          6,
	})
	require.NoError(t, err)
	s.basketDenomToId = make(map[string]uint64)
	s.basketDenomToId[s.basketDenom] = 1
	s.basketDenomToId[s.basketDenom2] = 2
	bsktClsTbl := s.db.GetTable(&basketv1.BasketClass{})
	err = bsktClsTbl.Insert(s.ctx, &basketv1.BasketClass{
		BasketId: 1,
		ClassId:  s.classId,
	})
	require.NoError(t, err)
	err = bsktClsTbl.Insert(s.ctx, &basketv1.BasketClass{
		BasketId: 2,
		ClassId:  s.classId,
	})
	require.NoError(t, err)

	return s
}

func TestPut(t *testing.T) {
	s := setupPut(t)

	testCases := []struct {
		name                string
		startingBalance     string
		msg                 basket2.MsgPut
		expectedBasketCoins string
		expectCalls         func()
		errMsg              string
	}{
		{
			name:            "valid case",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: s.basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: s.denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: s.denom}).
					Return(&s.batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: s.classId}).
					Return(&s.classInfoRes, nil)

				coinAward := sdk.NewCoins(sdk.NewCoin(s.basketDenom, sdk.NewInt(2_000_000)))
				s.bankKeeper.EXPECT().
					MintCoins(s.sdkCtx, basket2.BasketSubModuleName, coinAward).
					Return(nil)

				s.bankKeeper.EXPECT().
					SendCoinsFromModuleToAccount(s.sdkCtx, basket2.BasketSubModuleName, s.addr, coinAward).
					Return(nil)
			},
		},
		{
			name:            "valid case - basket 2 with rolling window",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: s.basketDenom2,
				Credits:     []*basket2.BasketCredit{{BatchDenom: s.denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: s.denom}).
					Return(&s.batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: s.classId}).
					Return(&s.classInfoRes, nil)

				coinAward := sdk.NewCoins(sdk.NewCoin(s.basketDenom2, sdk.NewInt(2_000_000)))
				s.bankKeeper.EXPECT().
					MintCoins(s.sdkCtx, basket2.BasketSubModuleName, coinAward).
					Return(nil)

				s.bankKeeper.EXPECT().
					SendCoinsFromModuleToAccount(s.sdkCtx, basket2.BasketSubModuleName, s.addr, coinAward).
					Return(nil)
			},
		},
		{
			name:            "insufficient funds",
			startingBalance: "1",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: s.basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: s.denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: s.denom}).
					Return(&s.batchInfoRes, nil)

				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: s.classId}).
					Return(&s.classInfoRes, nil)

			},
			errMsg: "insufficient funds",
		},
		{
			name:            "basket not found",
			startingBalance: "1",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: "FooBar",
				Credits:     []*basket2.BasketCredit{{BatchDenom: s.denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
			},
			errMsg: ormerrors.NotFound.Error(),
		},
		{
			name:            "batch not found",
			startingBalance: "20",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: s.basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: "FooBarBaz", Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: "FooBarBaz"}).
					Return(nil, orm.ErrNotFound)
			},
			errMsg: orm.ErrNotFound.Error(),
		},
		{
			name:            "class not allowed",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: s.basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: "blah", Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				badInfo := *s.batchInfoRes.Info
				badInfo.ClassId = "blah01"
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: "blah"}).
					Return(&ecocredit.QueryBatchInfoResponse{Info: &badInfo}, nil)
			},
			errMsg: "credit class blah01 is not allowed in this basket",
		},
		{
			name:            "wrong credit type",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: s.basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: s.denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: s.denom}).
					Return(&s.batchInfoRes, nil)

				badClass := *s.classInfoRes.Info
				badClass.CreditType.Name = "BadType"
				s.ecocreditKeeper.EXPECT().
					ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: s.classId}).
					Return(&ecocredit.QueryClassInfoResponse{Info: &badClass}, nil)
			},
			errMsg: "cannot use credit of type BadType in a basket that requires credit type carbon",
		},
		{
			name:            "batch out of time window",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: s.basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: s.denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				badTime, err := time.Parse("2006-01-02", "1984-01-01")
				require.NoError(t, err)
				badTimeInfo := *s.batchInfoRes.Info
				badTimeInfo.StartDate = &badTime
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: s.denom}).
					Return(&ecocredit.QueryBatchInfoResponse{Info: &badTimeInfo}, nil)

			},
			errMsg: "cannot put a credit from a batch with start date",
		},
		{
			name:            "batch outside of rolling time window",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       s.addr.String(),
				BasketDenom: s.basketDenom2,
				Credits:     []*basket2.BasketCredit{{BatchDenom: s.denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				badTimeInfo := *s.batchInfoRes.Info
				bogusDur := time.Duration(999999999999999)
				badTime := s.validStartDateWindow.Add(-bogusDur)
				badTimeInfo.StartDate = &badTime
				s.ecocreditKeeper.EXPECT().
					BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: s.denom}).
					Return(&ecocredit.QueryBatchInfoResponse{Info: &badTimeInfo}, nil)

			},
			errMsg: "cannot put a credit from a batch with start date",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.expectCalls()
			legacyStore := s.sdkCtx.KVStore(s.storeKey)
			tradKey := ecocredit.TradableBalanceKey(s.addr, ecocredit.BatchDenomT(s.denom))
			userFunds, err := math.NewDecFromString(tc.startingBalance)
			require.NoError(t, err)
			ecocredit.SetDecimal(legacyStore, tradKey, userFunds)
			res, err := s.k.Put(s.ctx, &tc.msg)
			if tc.errMsg == "" { //  no error!
				require.NoError(t, err)
				require.Equal(t, res.AmountReceived, tc.expectedBasketCoins)
				for _, credit := range tc.msg.Credits {
					assertUserSentCredits(t, userFunds, credit.Amount, tradKey, legacyStore)
					assertBasketHasCredits(t, s.ctx, credit, s.basketDenomToId[tc.msg.BasketDenom], s.basketBalanceTbl)
				}
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errMsg)
			}
		})
	}
}

func assertBasketHasCredits(t *testing.T, ctx context.Context, credit *basket2.BasketCredit, basketID uint64, basketBalTbl ormtable.Table) {
	basketBal := basketv1.BasketBalance{
		BasketId:       basketID,
		BatchDenom:     credit.BatchDenom,
		Balance:        "",
		BatchStartDate: nil,
	}
	found, err := basketBalTbl.Get(ctx, &basketBal)
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, basketBal.Balance, credit.Amount)
}

func assertUserSentCredits(t *testing.T, oldBalance math.Dec, amountSent string, balanceKey []byte, store types.KVStore) {
	amtSent, err := math.NewDecFromString(amountSent)
	require.NoError(t, err)
	currentBalance, err := ecocredit.GetDecimal(store, balanceKey)
	require.NoError(t, err)

	checkBalance, err := currentBalance.Add(amtSent)
	require.NoError(t, err)

	require.True(t, checkBalance.IsEqual(oldBalance))
}
