package basket_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	basket2 "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	mocks2 "github.com/regen-network/regen-ledger/x/ecocredit/mocks"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket/mocks"
)

func TestPut(t *testing.T) {
	basketDenom := "BASKET"
	basketDenom2 := "BASKET2"
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

	ctx := context.Background()
	_, _, addr := testdata.KeyTestPubAddr()
	ctrl := gomock.NewController(t)
	b := ormtest.NewMemoryBackend()
	db, err := ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{
		GetBackend: func(ctx context.Context) (ormtable.Backend, error) {
			return b, nil
		},
		GetReadBackend: func(ctx context.Context) (ormtable.ReadBackend, error) {
			return b, nil
		},
	})
	require.NoError(t, err)
	basketTbl := db.GetTable(&api.Basket{})
	err = basketTbl.Insert(ctx, &api.Basket{
		BasketDenom:       basketDenom,
		Name:              basketDenom,
		DisableAutoRetire: true,
		CreditTypeAbbrev:  "C",
		DateCriteria:      &api.DateCriteria{MinStartDate: timestamppb.New(startDate)},
		Exponent:          6,
	})
	require.NoError(t, err)
	basketBalanceTbl := db.GetTable(&api.BasketBalance{})
	var dur time.Duration = 500000000000000000
	validStartDateWindow := startDate.Add(-dur)
	err = basketTbl.Insert(ctx, &api.Basket{
		BasketDenom:       basketDenom2,
		Name:              basketDenom2,
		DisableAutoRetire: true,
		CreditTypeAbbrev:  "C",
		DateCriteria:      &api.DateCriteria{StartDateWindow: durationpb.New(dur)},
		Exponent:          6,
	})
	require.NoError(t, err)
	basketDenomToId := make(map[string]uint64)
	basketDenomToId[basketDenom] = 1
	basketDenomToId[basketDenom2] = 2
	bsktClsTbl := db.GetTable(&api.BasketClass{})
	err = bsktClsTbl.Insert(ctx, &api.BasketClass{
		BasketId: 1,
		ClassId:  classId,
	})
	require.NoError(t, err)
	err = bsktClsTbl.Insert(ctx, &api.BasketClass{
		BasketId: 2,
		ClassId:  classId,
	})
	require.NoError(t, err)

	bankKeeper := mocks2.NewMockBankKeeper(ctrl)
	ecocreditKeeper := mocks.NewMockEcocreditKeeper(ctrl)
	distKeeper := mocks2.NewMockDistributionKeeper(ctrl)
	sk := sdk.NewKVStoreKey("test")
	k := basket.NewKeeper(db, ecocreditKeeper, bankKeeper, distKeeper, sk)
	require.NotNil(t, k)

	sdkCtx := sdkContextForStoreKey(sk).WithContext(ctx).WithBlockTime(startDate)
	ctx = sdk.WrapSDKContext(sdkCtx)
	sdkCtx = ctx.Value(sdk.SdkContextKey).(sdk.Context)

	testCases := []struct {
		name                string
		startingBalance     string
		msg                 basket2.MsgPut
		expectedCredits     []*basket2.BasketCredit
		expectedBasketCoins string
		expectCalls         func()
		errMsg              string
	}{
		{
			name:            "valid case",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedCredits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				ecocreditKeeper.EXPECT().
					BatchInfo(ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				ecocreditKeeper.EXPECT().
					ProjectInfo(ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				ecocreditKeeper.EXPECT().
					ClassInfo(ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&classInfoRes, nil)

				coinAward := sdk.NewCoins(sdk.NewCoin(basketDenom, sdk.NewInt(2_000_000)))
				bankKeeper.EXPECT().
					MintCoins(sdkCtx, basket2.BasketSubModuleName, coinAward).
					Return(nil)

				bankKeeper.EXPECT().
					SendCoinsFromModuleToAccount(sdkCtx, basket2.BasketSubModuleName, addr, coinAward).
					Return(nil)
			},
		},
		{
			name:            "valid case - basket with existing balance",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "1"}},
			},
			expectedCredits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "3"}},
			expectedBasketCoins: "1000000", // 1 million
			expectCalls: func() {
				ecocreditKeeper.EXPECT().
					BatchInfo(ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				ecocreditKeeper.EXPECT().
					ProjectInfo(ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				ecocreditKeeper.EXPECT().
					ClassInfo(ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&classInfoRes, nil)

				coinAward := sdk.NewCoins(sdk.NewCoin(basketDenom, sdk.NewInt(1_000_000)))
				bankKeeper.EXPECT().
					MintCoins(sdkCtx, basket2.BasketSubModuleName, coinAward).
					Return(nil)

				bankKeeper.EXPECT().
					SendCoinsFromModuleToAccount(sdkCtx, basket2.BasketSubModuleName, addr, coinAward).
					Return(nil)
			},
		},
		{
			name:            "valid case - basket 2 with rolling window",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       addr.String(),
				BasketDenom: basketDenom2,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedCredits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				ecocreditKeeper.EXPECT().
					BatchInfo(ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				ecocreditKeeper.EXPECT().
					ProjectInfo(ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				ecocreditKeeper.EXPECT().
					ClassInfo(ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&classInfoRes, nil)

				coinAward := sdk.NewCoins(sdk.NewCoin(basketDenom2, sdk.NewInt(2_000_000)))
				bankKeeper.EXPECT().
					MintCoins(sdkCtx, basket2.BasketSubModuleName, coinAward).
					Return(nil)

				bankKeeper.EXPECT().
					SendCoinsFromModuleToAccount(sdkCtx, basket2.BasketSubModuleName, addr, coinAward).
					Return(nil)
			},
		},
		{
			name:            "insufficient funds",
			startingBalance: "1",
			msg: basket2.MsgPut{
				Owner:       addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				ecocreditKeeper.EXPECT().
					BatchInfo(ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				ecocreditKeeper.EXPECT().
					ProjectInfo(ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				ecocreditKeeper.EXPECT().
					ClassInfo(ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&classInfoRes, nil)

			},
			errMsg: basket.ErrInsufficientCredits.Error(),
		},
		{
			name:            "basket not found",
			startingBalance: "1",
			msg: basket2.MsgPut{
				Owner:       addr.String(),
				BasketDenom: "FooBar",
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
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
				Owner:       addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: "FooBarBaz", Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				ecocreditKeeper.EXPECT().
					BatchInfo(ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: "FooBarBaz"}).
					Return(nil, orm.ErrNotFound)
			},
			errMsg: orm.ErrNotFound.Error(),
		},
		//{
		//	name:            "class not allowed",
		//	startingBalance: "100000000",
		//	msg: basket2.MsgPut{
		//		Owner:       addr.String(),
		//		BasketDenom: basketDenom,
		//		Credits:     []*basket2.BasketCredit{{BatchDenom: "blah", Amount: "2"}},
		//	},
		//	expectedBasketCoins: "2000000", // 2 million
		//	expectCalls: func() {
		//		badInfo := *batchInfoRes.Info
		//		badInfo.ClassId = "blah01"
		//		ecocreditKeeper.EXPECT().
		//			BatchInfo(ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: "blah"}).
		//			Return(&ecocredit.QueryBatchInfoResponse{Info: &badInfo}, nil)
		//	},
		//	errMsg: "credit class blah01 is not allowed in this basket",
		//},
		{
			name:            "wrong credit type",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				ecocreditKeeper.EXPECT().
					BatchInfo(ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&batchInfoRes, nil)

				ecocreditKeeper.EXPECT().
					ProjectInfo(ctx, &ecocredit.QueryProjectInfoRequest{ProjectId: projectId}).
					Return(&projectInfoRes, nil)

				badClass := *classInfoRes.Info
				badClass.CreditType.Abbreviation = "FOO"
				ecocreditKeeper.EXPECT().
					ClassInfo(ctx, &ecocredit.QueryClassInfoRequest{ClassId: classId}).
					Return(&ecocredit.QueryClassInfoResponse{Info: &badClass}, nil)
			},
			errMsg: "cannot use credit of type FOO in a basket that requires credit type C",
		},
		{
			name:            "batch out of time window",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       addr.String(),
				BasketDenom: basketDenom,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				badTime, err := time.Parse("2006-01-02", "1984-01-01")
				require.NoError(t, err)
				badTimeInfo := *batchInfoRes.Info
				badTimeInfo.StartDate = &badTime
				ecocreditKeeper.EXPECT().
					BatchInfo(ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&ecocredit.QueryBatchInfoResponse{Info: &badTimeInfo}, nil)

			},
			errMsg: "cannot put a credit from a batch with start date",
		},
		{
			name:            "batch outside of rolling time window",
			startingBalance: "100000000",
			msg: basket2.MsgPut{
				Owner:       addr.String(),
				BasketDenom: basketDenom2,
				Credits:     []*basket2.BasketCredit{{BatchDenom: denom, Amount: "2"}},
			},
			expectedBasketCoins: "2000000", // 2 million
			expectCalls: func() {
				badTimeInfo := *batchInfoRes.Info
				bogusDur := time.Duration(999999999999999)
				badTime := validStartDateWindow.Add(-bogusDur)
				badTimeInfo.StartDate = &badTime
				ecocreditKeeper.EXPECT().
					BatchInfo(ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: denom}).
					Return(&ecocredit.QueryBatchInfoResponse{Info: &badTimeInfo}, nil)

			},
			errMsg: "cannot put a credit from a batch with start date",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.expectCalls()
			legacyStore := sdkCtx.KVStore(sk)
			tradKey := ecocredit.TradableBalanceKey(addr, ecocredit.BatchDenomT(denom))
			userFunds, err := math.NewDecFromString(tc.startingBalance)
			require.NoError(t, err)
			ecocredit.SetDecimal(legacyStore, tradKey, userFunds)
			res, err := k.Put(ctx, &tc.msg)
			if tc.errMsg == "" { //  no error!
				require.NoError(t, err)
				require.Equal(t, res.AmountReceived, tc.expectedBasketCoins)
				for _, credit := range tc.msg.Credits {
					assertUserSentCredits(t, userFunds, credit.Amount, tradKey, legacyStore)
				}
				for _, credit := range tc.expectedCredits {
					assertBasketHasCredits(t, ctx, credit, basketDenomToId[tc.msg.BasketDenom], basketBalanceTbl)
				}
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errMsg)
			}
		})
	}
}

func assertBasketHasCredits(t *testing.T, ctx context.Context, credit *basket2.BasketCredit, basketID uint64, basketBalTbl ormtable.Table) {
	basketBal := api.BasketBalance{
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

	require.True(t, checkBalance.Equal(oldBalance))
}

func sdkContextForStoreKey(key *types.KVStoreKey) sdk.Context {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}
	return sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
}
