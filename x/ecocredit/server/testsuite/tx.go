package testsuite

import (
	"fmt"
	"github.com/regen-network/regen-ledger/types/math"
	server2 "github.com/regen-network/regen-ledger/x/ecocredit/server"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

//func (s *IntegrationTestSuite) TestScenarioCreateSellOrders() {
//	addr1 := s.signers[3].String()
//
//	// create credit class and issue credits to addr1
//	_, createBatchRes := s.createClassAndIssueBatch(addr1, "2.0")
//
//	askPrice1 := sdk.NewInt64Coin("stake", 1000000)
//	// TODO: Verify that AskPrice.Denom is in AllowAskDenom #624
//	//askPrice2 := sdk.NewInt64Coin("token", 1000000)
//
//	// create sell orders
//	testCases := []struct {
//		name    string
//		owner   string
//		orders  []*ecocredit.MsgSell_Order
//		expErr  string
//		wantErr bool
//	}{
//		{
//			name:  "insufficient credit balance - batch denom",
//			owner: addr1,
//			orders: []*ecocredit.MsgSell_Order{
//				{
//					BatchDenom:        "A00-00000000-00000000-000",
//					Quantity:          "1.0",
//					AskPrice:          &askPrice1,
//					DisableAutoRetire: true,
//				},
//				{
//					BatchDenom:        "A00-00000000-00000000-000",
//					Quantity:          "1.0",
//					AskPrice:          &askPrice1,
//					DisableAutoRetire: true,
//				},
//			},
//			expErr:  "insufficient credit balance",
//			wantErr: true,
//		},
//		{
//			name:  "insufficient credit balance - quantity",
//			owner: addr1,
//			orders: []*ecocredit.MsgSell_Order{
//				{
//					BatchDenom:        createBatchRes.BatchDenom,
//					Quantity:          "99",
//					AskPrice:          &askPrice1,
//					DisableAutoRetire: true,
//				},
//				{
//					BatchDenom:        createBatchRes.BatchDenom,
//					Quantity:          "99",
//					AskPrice:          &askPrice1,
//					DisableAutoRetire: true,
//				},
//			},
//			expErr:  "insufficient credit balance",
//			wantErr: true,
//		},
//		// TODO: Verify that AskPrice.Denom is in AllowAskDenom #624
//		//{
//		//	name: "denom not allowed",
//		//	owner: addr1,
//		//	orders: []*ecocredit.MsgSell_Order{
//		//		{
//		//			BatchDenom:        createBatchRes.BatchDenom,
//		//			Quantity:          "1.0",
//		//			AskPrice:          &askPrice2,
//		//			DisableAutoRetire: true,
//		//		},
//		//		{
//		//			BatchDenom:        createBatchRes.BatchDenom,
//		//			Quantity:          "1.0",
//		//			AskPrice:          &askPrice2,
//		//			DisableAutoRetire: true,
//		//		},
//		//	},
//		//	expErr: "denom not allowed",
//		//	wantErr: true,
//		//},
//		{
//			name:  "valid request",
//			owner: addr1,
//			orders: []*ecocredit.MsgSell_Order{
//				{
//					BatchDenom:        createBatchRes.BatchDenom,
//					Quantity:          "1.0",
//					AskPrice:          &askPrice1,
//					DisableAutoRetire: true,
//				},
//				{
//					BatchDenom:        createBatchRes.BatchDenom,
//					Quantity:          "1.0",
//					AskPrice:          &askPrice1,
//					DisableAutoRetire: true,
//				},
//			},
//			expErr:  "",
//			wantErr: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//
//		s.Run(tc.name, func() {
//			require := s.Require()
//
//			res, err := s.msgClient.Sell(s.ctx, &ecocredit.MsgSell{
//				Owner:  tc.owner,
//				Orders: tc.orders,
//			})
//
//			if tc.wantErr {
//				require.Error(err)
//				require.Contains(err.Error(), tc.expErr)
//			} else {
//				require.NoError(err)
//				require.NotNil(res.SellOrderIds)
//
//				// query first sell order
//				_, sellError1 := s.queryClient.SellOrder(s.ctx, &ecocredit.QuerySellOrderRequest{
//					SellOrderId: res.SellOrderIds[0],
//				})
//
//				// query second sell order
//				_, sellError2 := s.queryClient.SellOrder(s.ctx, &ecocredit.QuerySellOrderRequest{
//					SellOrderId: res.SellOrderIds[1],
//				})
//
//				require.NoError(sellError1)
//				require.NoError(sellError2)
//			}
//		})
//	}
//}
//
//func (s *IntegrationTestSuite) TestScenarioUpdateSellOrders() {
//	addr1 := s.signers[3].String()
//	addr2 := s.signers[4].String()
//
//	// create credit class and issue credits to addr1
//	_, createBatchRes := s.createClassAndIssueBatch(addr1, "2.0")
//
//	askPrice1 := sdk.NewInt64Coin("stake", 2000000)
//	// TODO: Verify that NewAskPrice.Denom is in AllowAskDenom #624
//	//askPrice2 := sdk.NewInt64Coin("token", 2000000)
//
//	// create sell order
//	sellRes, err := s.msgClient.Sell(s.ctx, &ecocredit.MsgSell{
//		Owner: addr1,
//		Orders: []*ecocredit.MsgSell_Order{
//			{
//				BatchDenom:        createBatchRes.BatchDenom,
//				Quantity:          "1.0",
//				AskPrice:          &askPrice1,
//				DisableAutoRetire: true,
//			},
//			{
//				BatchDenom:        createBatchRes.BatchDenom,
//				Quantity:          "1.0",
//				AskPrice:          &askPrice1,
//				DisableAutoRetire: true,
//			},
//		},
//	})
//	s.Require().NoError(err)
//
//	// update sell orders
//	testCases := []struct {
//		name    string
//		owner   string
//		updates []*ecocredit.MsgUpdateSellOrders_Update
//		expErr  string
//		wantErr bool
//	}{
//		{
//			name:  "invalid sell order",
//			owner: addr1,
//			updates: []*ecocredit.MsgUpdateSellOrders_Update{
//				{
//					SellOrderId:       99,
//					NewQuantity:       "1.0",
//					NewAskPrice:       &askPrice1,
//					DisableAutoRetire: true,
//				},
//				{
//					SellOrderId:       100,
//					NewQuantity:       "1.0",
//					NewAskPrice:       &askPrice1,
//					DisableAutoRetire: true,
//				},
//			},
//			expErr:  "invalid sell order",
//			wantErr: true,
//		},
//		{
//			name:  "unauthorized",
//			owner: addr2,
//			updates: []*ecocredit.MsgUpdateSellOrders_Update{
//				{
//					SellOrderId:       sellRes.SellOrderIds[0],
//					NewQuantity:       "1.0",
//					NewAskPrice:       &askPrice1,
//					DisableAutoRetire: true,
//				},
//				{
//					SellOrderId:       sellRes.SellOrderIds[1],
//					NewQuantity:       "1.0",
//					NewAskPrice:       &askPrice1,
//					DisableAutoRetire: true,
//				},
//			},
//			expErr:  "unauthorized",
//			wantErr: true,
//		},
//		{
//			name:  "insufficient credit balance",
//			owner: addr1,
//			updates: []*ecocredit.MsgUpdateSellOrders_Update{
//				{
//					SellOrderId:       sellRes.SellOrderIds[0],
//					NewQuantity:       "99",
//					NewAskPrice:       &askPrice1,
//					DisableAutoRetire: true,
//				},
//				{
//					SellOrderId:       sellRes.SellOrderIds[1],
//					NewQuantity:       "99",
//					NewAskPrice:       &askPrice1,
//					DisableAutoRetire: true,
//				},
//			},
//			expErr:  "insufficient credit balance",
//			wantErr: true,
//		},
//		// TODO: Verify that NewAskPrice.Denom is in AllowAskDenom #624
//		//{
//		//	name: "denom not allowed",
//		//	owner: addr1,
//		//	updates: []*ecocredit.MsgUpdateSellOrders_Update{
//		//		{
//		//			SellOrderId:       sellRes.SellOrderIds[0],
//		//			NewQuantity:       "1.0",
//		//			NewAskPrice:       &askPrice2,
//		//			DisableAutoRetire: true,
//		//		},
//		//		{
//		//			SellOrderId:       sellRes.SellOrderIds[1],
//		//			NewQuantity:       "1.0",
//		//			NewAskPrice:       &askPrice2,
//		//			DisableAutoRetire: true,
//		//		},
//		//	},
//		//	expErr: "denom not allowed",
//		//	wantErr: true,
//		//},
//		{
//			name:  "valid request",
//			owner: addr1,
//			updates: []*ecocredit.MsgUpdateSellOrders_Update{
//				{
//					SellOrderId:       sellRes.SellOrderIds[0],
//					NewQuantity:       "1.0",
//					NewAskPrice:       &askPrice1,
//					DisableAutoRetire: true,
//				},
//				{
//					SellOrderId:       sellRes.SellOrderIds[1],
//					NewQuantity:       "1.0",
//					NewAskPrice:       &askPrice1,
//					DisableAutoRetire: true,
//				},
//			},
//			expErr:  "",
//			wantErr: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//
//		s.Run(tc.name, func() {
//			require := s.Require()
//
//			_, err := s.msgClient.UpdateSellOrders(s.ctx, &ecocredit.MsgUpdateSellOrders{
//				Owner:   tc.owner,
//				Updates: tc.updates,
//			})
//
//			if tc.wantErr {
//				require.Error(err)
//				require.Contains(err.Error(), tc.expErr)
//			} else {
//				require.NoError(err)
//
//				// query first sell order
//				sellResponse1, sellError1 := s.queryClient.SellOrder(s.ctx, &ecocredit.QuerySellOrderRequest{
//					SellOrderId: tc.updates[0].SellOrderId,
//				})
//
//				// query second sell order
//				sellResponse2, sellError2 := s.queryClient.SellOrder(s.ctx, &ecocredit.QuerySellOrderRequest{
//					SellOrderId: tc.updates[1].SellOrderId,
//				})
//
//				require.NoError(sellError1)
//				require.NoError(sellError2)
//				require.Equal(tc.updates[0].NewAskPrice, sellResponse1.SellOrder.AskPrice)
//				require.Equal(tc.updates[1].NewAskPrice, sellResponse2.SellOrder.AskPrice)
//			}
//		})
//	}
//}
//
//func (s *IntegrationTestSuite) TestScenarioCreateBuyOrders() {
//	addr1 := s.signers[3]
//	addr2 := s.signers[4]
//
//	// create credit class and issue credits to addr1
//	_, createBatchRes := s.createClassAndIssueBatch(addr1.String(), "4.0")
//
//	bidPrice1 := sdk.NewInt64Coin("stake", 1000000)
//	bidPrice2 := sdk.NewInt64Coin("stake", 9999999)
//	// TODO: Verify that BidPrice.Denom is in AllowAskDenom #624
//	//bidPrice3 := sdk.NewInt64Coin("token", 1000000)
//
//	// fund buyer account
//	s.Require().NoError(s.fundAccount(addr2, sdk.NewCoins(sdk.NewInt64Coin("stake", 3000000))))
//
//	// create sell orders
//	sellRes, err := s.msgClient.Sell(s.ctx, &ecocredit.MsgSell{
//		Owner: addr1.String(),
//		Orders: []*ecocredit.MsgSell_Order{
//			{
//				BatchDenom:        createBatchRes.BatchDenom,
//				Quantity:          "1.0",
//				AskPrice:          &bidPrice1,
//				DisableAutoRetire: true,
//			},
//			{
//				BatchDenom:        createBatchRes.BatchDenom,
//				Quantity:          "1.0",
//				AskPrice:          &bidPrice1,
//				DisableAutoRetire: true,
//			},
//			{
//				BatchDenom:        createBatchRes.BatchDenom,
//				Quantity:          "1.0",
//				AskPrice:          &bidPrice1,
//				DisableAutoRetire: true,
//			},
//			{
//				BatchDenom:        createBatchRes.BatchDenom,
//				Quantity:          "1.0",
//				AskPrice:          &bidPrice1,
//				DisableAutoRetire: true,
//			},
//		},
//	})
//	s.Require().NoError(err)
//
//	// process buy orders
//	testCases := []struct {
//		name             string
//		buyer            string
//		orders           []*ecocredit.MsgBuy_Order
//		expErr           string
//		wantErr          bool
//		partial          bool
//		expCoinBalance   sdk.Coin
//		expCreditBalance *ecocredit.QueryBalanceResponse
//	}{
//		{
//			name:  "invalid sell order",
//			buyer: addr2.String(),
//			orders: []*ecocredit.MsgBuy_Order{
//				{
//					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: 99}},
//					Quantity:          "1.0",
//					BidPrice:          &bidPrice1,
//					DisableAutoRetire: true,
//				},
//				{
//					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: 100}},
//					Quantity:          "1.0",
//					BidPrice:          &bidPrice1,
//					DisableAutoRetire: true,
//				},
//			},
//			expErr:  "not found",
//			wantErr: true,
//		},
//		{
//			name:  "insufficient coin balance - quantity",
//			buyer: addr2.String(),
//			orders: []*ecocredit.MsgBuy_Order{
//				{
//					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[0]}},
//					Quantity:          "99.99",
//					BidPrice:          &bidPrice1,
//					DisableAutoRetire: true,
//				},
//				{
//					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[1]}},
//					Quantity:          "99.99",
//					BidPrice:          &bidPrice1,
//					DisableAutoRetire: true,
//				},
//			},
//			expErr:  "insufficient balance",
//			wantErr: true,
//		},
//		{
//			name:  "insufficient coin balance - bid price",
//			buyer: addr2.String(),
//			orders: []*ecocredit.MsgBuy_Order{
//				{
//					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[0]}},
//					Quantity:          "1.0",
//					BidPrice:          &bidPrice2,
//					DisableAutoRetire: true,
//				},
//				{
//					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[1]}},
//					Quantity:          "1.0",
//					BidPrice:          &bidPrice2,
//					DisableAutoRetire: true,
//				},
//			},
//			expErr:  "insufficient balance",
//			wantErr: true,
//		},
//		// TODO: Verify that BidPrice.Denom is in AllowAskDenom #624
//		//{
//		//	name: "denom not allowed",
//		//	buyer: addr2.String(),
//		//	orders: []*ecocredit.MsgBuy_Order{
//		//		{
//		//			Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[0]}},
//		//			Quantity:          "1.0",
//		//			BidPrice:          &bidPrice3,
//		//			DisableAutoRetire: true,
//		//		},
//		//		{
//		//			Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[1]}},
//		//			Quantity:          "1.0",
//		//			BidPrice:          &bidPrice3,
//		//			DisableAutoRetire: true,
//		//		},
//		//	},
//		//	expErr: "denom not allowed",
//		//	wantErr: true,
//		//},
//		{
//			name:  "valid request",
//			buyer: addr2.String(),
//			orders: []*ecocredit.MsgBuy_Order{
//				{
//					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[0]}},
//					Quantity:          "1.0",
//					BidPrice:          &bidPrice1,
//					DisableAutoRetire: true,
//				},
//				{
//					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[1]}},
//					Quantity:          "1.0",
//					BidPrice:          &bidPrice1,
//					DisableAutoRetire: true,
//				},
//			},
//			expErr:  "",
//			wantErr: false,
//			partial: false,
//			expCoinBalance: sdk.Coin{
//				Denom:  "stake",
//				Amount: sdk.NewInt(1000000),
//			},
//			expCreditBalance: &ecocredit.QueryBalanceResponse{TradableAmount: "2", RetiredAmount: "0"},
//		},
//		{
//			name:  "valid request - partial fill",
//			buyer: addr2.String(),
//			orders: []*ecocredit.MsgBuy_Order{
//				{
//					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[2]}},
//					Quantity:          "0.5",
//					BidPrice:          &bidPrice1,
//					DisableAutoRetire: true,
//				},
//				{
//					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[3]}},
//					Quantity:          "0.5",
//					BidPrice:          &bidPrice1,
//					DisableAutoRetire: true,
//				},
//			},
//			expErr:  "",
//			wantErr: false,
//			partial: true,
//			expCoinBalance: sdk.Coin{
//				Denom:  "stake",
//				Amount: sdk.NewInt(0),
//			},
//			expCreditBalance: &ecocredit.QueryBalanceResponse{TradableAmount: "3", RetiredAmount: "0"},
//		},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//
//		s.Run(tc.name, func() {
//			require := s.Require()
//
//			// get buyer coin balance before
//			coinBalanceBefore := s.bankKeeper.GetBalance(s.sdkCtx, addr2, "stake")
//
//			// get buyer credit balance before
//			creditBalanceBefore, _ := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
//				Account:    addr2.String(),
//				BatchDenom: createBatchRes.BatchDenom,
//			})
//
//			// process buy orders
//			res, err := s.msgClient.Buy(s.ctx, &ecocredit.MsgBuy{
//				Buyer:  tc.buyer,
//				Orders: tc.orders,
//			})
//
//			// get buyer coin balance after
//			coinBalanceAfter := s.bankKeeper.GetBalance(s.sdkCtx, addr2, "stake")
//
//			// get buyer credit balance after
//			creditBalanceAfter, _ := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
//				Account:    addr2.String(),
//				BatchDenom: createBatchRes.BatchDenom,
//			})
//
//			// query first sell order
//			sellResponse1, sellError1 := s.queryClient.SellOrder(s.ctx, &ecocredit.QuerySellOrderRequest{
//				SellOrderId: tc.orders[0].Selection.GetSellOrderId(),
//			})
//
//			// query second sell order
//			sellResponse2, sellError2 := s.queryClient.SellOrder(s.ctx, &ecocredit.QuerySellOrderRequest{
//				SellOrderId: tc.orders[1].Selection.GetSellOrderId(),
//			})
//
//			if tc.wantErr {
//				require.Error(err)
//				require.Contains(err.Error(), tc.expErr)
//				require.Equal(coinBalanceBefore, coinBalanceAfter)
//				require.Equal(creditBalanceBefore, creditBalanceAfter)
//			} else {
//				require.NoError(err)
//				require.NotNil(res.BuyOrderIds)
//
//				require.Equal(tc.expCoinBalance, coinBalanceAfter)
//				require.Equal(tc.expCreditBalance, creditBalanceAfter)
//
//				if tc.partial {
//					require.NotNil(sellResponse1)
//					require.NotNil(sellResponse2)
//					require.NoError(sellError1)
//					require.NoError(sellError2)
//				} else {
//					require.Nil(sellResponse1)
//					require.Nil(sellResponse2)
//					require.Error(sellError1)
//					require.Error(sellError2)
//				}
//			}
//		})
//	}
//}
//
//func (s *IntegrationTestSuite) TestScenarioAllowAskDenom() {
//	addr1 := s.signers[3].String()
//
//	// TODO: Verify governance module address for AllowAskDenom #624
//	//rootAddress := s.accountKeeper.GetModuleAddress(govtypes.ModuleName).String()
//
//	// add ask denom
//	testCases := []struct {
//		name         string
//		rootAddress  string
//		denom        string
//		displayDenom string
//		exponent     uint32
//		expErr       string
//		wantErr      bool
//	}{
//		{
//			name:         "unauthorized address",
//			rootAddress:  addr1,
//			denom:        "utoken",
//			displayDenom: "token",
//			exponent:     6,
//			expErr:       "unauthorized",
//			wantErr:      true,
//		},
//		// TODO: Verify governance module address for AllowAskDenom #624
//		//{
//		//	name: "valid request",
//		//	rootAddress: rootAddress,
//		//	denom: "utoken",
//		//	displayDenom: "token",
//		//	exponent: 6,
//		//	expErr: "",
//		//	wantErr: false,
//		//},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//
//		s.Run(tc.name, func() {
//			require := s.Require()
//
//			res, err := s.msgClient.AllowAskDenom(s.ctx, &ecocredit.MsgAllowAskDenom{
//				RootAddress:  tc.rootAddress,
//				Denom:        tc.denom,
//				DisplayDenom: tc.displayDenom,
//				Exponent:     tc.exponent,
//			})
//
//			if tc.wantErr {
//				require.Error(err)
//				require.Contains(err.Error(), tc.expErr)
//			} else {
//				require.NoError(err)
//				require.NotNil(res)
//			}
//		})
//	}
//}

func (s *IntegrationTestSuite) createClassAndIssueBatch(recipient string, tradableCredits string) (*ecocredit.MsgCreateClassResponse, *ecocredit.MsgCreateProjectResponse, *ecocredit.MsgCreateBatchResponse) {
	admin := s.signers[0]
	issuer1 := s.signers[1].String()
	issuer2 := s.signers[2].String()

	time1 := time.Now()
	time2 := time.Now()

	// fund admin account
	s.Require().NoError(s.fundAccount(admin, sdk.NewCoins(sdk.NewInt64Coin("stake", ecocredit.DefaultCreditClassFeeTokens.Int64()))))

	// create credit class
	createClassRes, err := s.msgClient.CreateClass(s.ctx, &ecocredit.MsgCreateClass{
		Admin:          admin.String(),
		Issuers:        []string{issuer1, issuer2},
		Metadata:       nil,
		CreditTypeName: "carbon",
	})
	s.Require().NoError(err)

	// create project
	projectRes, err := s.msgClient.CreateProject(s.ctx, &ecocredit.MsgCreateProject{
		ClassId:         createClassRes.ClassId,
		Issuer:          issuer1,
		Metadata:        []byte("metadata"),
		ProjectLocation: "AB",
	})
	s.Require().NoError(err)

	// create credit batch
	createBatchRes, err := s.msgClient.CreateBatch(s.ctx, &ecocredit.MsgCreateBatch{
		Issuer:    issuer1,
		ProjectId: projectRes.ProjectId,
		StartDate: &time1,
		EndDate:   &time2,
		Issuance: []*ecocredit.MsgCreateBatch_BatchIssuance{
			{
				Recipient:          recipient,
				TradableAmount:     tradableCredits,
				RetiredAmount:      "0",
				RetirementLocation: "YZ",
			},
		},
	})
	s.Require().NoError(err)

	return createClassRes, projectRes, createBatchRes
}

//func (s *IntegrationTestSuite) TestCreateBasket() {
//	server := s.msgClient
//	require := s.Require()
//	addr := s.signers[0]
//
//	class, project, batch := s.createClassAndIssueBatch(addr.String(), "9000000000000")
//
//	testCases := []struct {
//		name   string
//		msg    *ecocredit.MsgCreateBasket
//		expErr bool
//		errMsg string
//	}{
//		{
//			name: "valid - no basket criteria",
//			msg: &ecocredit.MsgCreateBasket{
//				Curator:           addr.String(),
//				Name:              "My Very Cool Basket",
//				DisplayName:       "COOL",
//				Exponent:          1,
//				BasketCriteria:    nil,
//				DisableAutoRetire: false,
//				AllowPicking:      false,
//			},
//		},
//		{
//			name: "valid - with all possible scalar criteria",
//			msg: &ecocredit.MsgCreateBasket{
//				Curator:     addr.String(),
//				Name:        "basket23095",
//				DisplayName: "basket23095",
//				Exponent:    1,
//				BasketCriteria: &ecocredit.Filter{
//					Sum: &ecocredit.Filter_And_{
//						And: &ecocredit.Filter_And{
//							Filters: []*ecocredit.Filter{
//								{Sum: &ecocredit.Filter_CreditTypeName{CreditTypeName: "carbon"}},
//								{Sum: &ecocredit.Filter_ClassId{ClassId: class.ClassId}},
//								{Sum: &ecocredit.Filter_BatchDenom{BatchDenom: batch.BatchDenom}},
//								{Sum: &ecocredit.Filter_ProjectId{ProjectId: project.ProjectId}},
//							}}}},
//				DisableAutoRetire: false,
//				AllowPicking:      false,
//			},
//		},
//		{
//			name: "invalid - basket with name already exists",
//			msg: &ecocredit.MsgCreateBasket{
//				Curator:           addr.String(),
//				Name:              "My Very Cool Basket",
//				DisplayName:       "COOL2",
//				Exponent:          1,
//				BasketCriteria:    nil,
//				DisableAutoRetire: false,
//				AllowPicking:      false,
//			},
//			expErr: true,
//			errMsg: "basket with name My Very Cool Basket already exists",
//		},
//		{
//			name: "invalid - bad credit type in filter",
//			msg: &ecocredit.MsgCreateBasket{
//				Curator:           addr.String(),
//				Name:              "basket203958",
//				DisplayName:       "basket203958",
//				Exponent:          1,
//				BasketCriteria:    &ecocredit.Filter{Sum: &ecocredit.Filter_CreditTypeName{CreditTypeName: "foobar"}},
//				DisableAutoRetire: false,
//				AllowPicking:      false,
//			},
//			expErr: true,
//			errMsg: "credit type foobar not found",
//		},
//		{
//			name: "invalid - bad class ID",
//			msg: &ecocredit.MsgCreateBasket{
//				Curator:           addr.String(),
//				Name:              "basket3",
//				DisplayName:       "bb3",
//				Exponent:          1,
//				BasketCriteria:    &ecocredit.Filter{Sum: &ecocredit.Filter_ClassId{ClassId: "Z15"}},
//				DisableAutoRetire: false,
//				AllowPicking:      false,
//			},
//			expErr: true,
//			errMsg: "credit class with id Z15 not found",
//		},
//		{
//			name: "invalid - bad batch denom",
//			msg: &ecocredit.MsgCreateBasket{
//				Curator:           addr.String(),
//				Name:              "basket4",
//				DisplayName:       "bb4",
//				Exponent:          1,
//				BasketCriteria:    &ecocredit.Filter{Sum: &ecocredit.Filter_BatchDenom{BatchDenom: "A00-00000000-00000000-000"}},
//				DisableAutoRetire: false,
//				AllowPicking:      false,
//			},
//			expErr: true,
//			errMsg: "batch with denom A00-00000000-00000000-000 not found",
//		},
//		{
//			name: "invalid - bad project ID",
//			msg: &ecocredit.MsgCreateBasket{
//				Curator:           addr.String(),
//				Name:              "basket5",
//				DisplayName:       "bb5",
//				Exponent:          1,
//				BasketCriteria:    &ecocredit.Filter{Sum: &ecocredit.Filter_ProjectId{ProjectId: "F00"}},
//				DisableAutoRetire: false,
//				AllowPicking:      false,
//			},
//			expErr: true,
//			errMsg: "project with id F00 not found",
//		},
//	}
//
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			res, err := server.CreateBasket(s.ctx, tc.msg)
//			if tc.expErr {
//				require.Error(err)
//				require.Nil(res)
//				require.Contains(err.Error(), tc.errMsg)
//			} else {
//				require.NoError(err)
//				require.NotNil(res)
//			}
//		})
//	}
//}
//
//func (s *IntegrationTestSuite) TestAddToBasket() {
//	require := s.Require()
//	server := s.msgClient
//	admin := s.signers[0]
//
//	class, _, batch := s.createClassAndIssueBatch(admin.String(), "1000000000")
//	class2, _, batch2 := s.createClassAndIssueBatch(admin.String(), "50000050500")
//
//	testCases := []struct {
//		name   string
//		basket *ecocredit.MsgCreateBasket
//		msg    *ecocredit.MsgAddToBasket
//		expErr bool
//		errMsg string
//	}{
//		{
//			name: "valid - simple basket 1 basket token : 1 basket credit",
//			basket: &ecocredit.MsgCreateBasket{
//				Curator:        admin.String(),
//				Name:           "FooBarBasket",
//				BasketCriteria: &ecocredit.Filter{Sum: &ecocredit.Filter_BatchDenom{BatchDenom: batch.BatchDenom}},
//				DisplayName:    "FBB",
//				Exponent:       1,
//			},
//			msg: &ecocredit.MsgAddToBasket{
//				Owner:       admin.String(),
//				BasketDenom: "FooBarBasket",
//				Credits:     []*ecocredit.BasketCredit{{BatchDenom: batch.BatchDenom, TradableAmount: "10"}},
//			},
//		},
//		{
//			name:   "invalid - insufficient credits",
//			basket: nil, // using the basket from previous test
//			msg: &ecocredit.MsgAddToBasket{
//				Owner:       admin.String(),
//				BasketDenom: "FooBarBasket",
//				Credits:     []*ecocredit.BasketCredit{{BatchDenom: batch.BatchDenom, TradableAmount: "1000000000000"}},
//			},
//			expErr: true,
//			errMsg: "insufficient credit balance",
//		},
//		{
//			name:   "invalid - does not match filter",
//			basket: nil, // using the basket from previous test
//			msg: &ecocredit.MsgAddToBasket{
//				Owner:       admin.String(),
//				BasketDenom: "FooBarBasket",
//				Credits:     []*ecocredit.BasketCredit{{BatchDenom: batch2.BatchDenom, TradableAmount: "10"}},
//			},
//			expErr: true,
//			errMsg: fmt.Sprintf("basket filter requires batch denom %s, but a credit with batch denom %s was given", batch.BatchDenom, batch2.BatchDenom),
//		},
//		{
//			name: "valid - OR filter",
//			basket: &ecocredit.MsgCreateBasket{
//				Curator:     admin.String(),
//				Name:        "bfzxed",
//				DisplayName: "bff",
//				Exponent:    5,
//				BasketCriteria: &ecocredit.Filter{Sum: &ecocredit.Filter_Or_{
//					Or: &ecocredit.Filter_Or{Filters: []*ecocredit.Filter{
//						{Sum: &ecocredit.Filter_ClassId{ClassId: class.ClassId}},
//						{Sum: &ecocredit.Filter_ClassId{ClassId: class2.ClassId}}}}}},
//				DisableAutoRetire: false,
//				AllowPicking:      false,
//			},
//			msg: &ecocredit.MsgAddToBasket{
//				Owner:       admin.String(),
//				BasketDenom: "bfzxed",
//				Credits: []*ecocredit.BasketCredit{
//					{BatchDenom: batch.BatchDenom, TradableAmount: "5"},
//					{BatchDenom: batch2.BatchDenom, TradableAmount: "5"}},
//			},
//		},
//		{
//			name: "invalid - basket not found",
//			msg: &ecocredit.MsgAddToBasket{
//				Owner:       admin.String(),
//				BasketDenom: "FooBarBaz",
//				Credits:     []*ecocredit.BasketCredit{{batch.BatchDenom, "2"}},
//			},
//			expErr: true,
//			errMsg: "basket FooBarBaz not found",
//		},
//		{
//			name: "invalid - batch not found",
//			msg: &ecocredit.MsgAddToBasket{
//				Owner:       admin.String(),
//				BasketDenom: "bfzxed",
//				Credits:     []*ecocredit.BasketCredit{{"Z99-00000000-00000000-000", "2"}},
//			},
//			expErr: true,
//			errMsg: "batch Z99-00000000-00000000-000 not found",
//		},
//	}
//
//	// this is an ugly hack to make tests reference previous baskets for querying purposes.
//	var lastBasket string
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			if tc.basket != nil {
//				res, err := server.CreateBasket(s.ctx, tc.basket)
//				require.NoError(err)
//				require.NotNil(res)
//				lastBasket = res.BasketAddress
//			}
//
//			res2, err := server.AddToBasket(s.ctx, tc.msg)
//			if tc.expErr {
//				require.Error(err)
//				require.Contains(err.Error(), tc.errMsg)
//			} else {
//				require.NoError(err)
//				require.NotNil(res2)
//
//				actualTokensBack, err := math.NewPositiveFixedDecFromString(res2.AmountReceived, 6)
//				require.NoError(err)
//
//				creditsDeposited := math.NewDecFromInt64(0)
//				for _, credit := range tc.msg.Credits {
//					dec, err := math.NewDecFromString(credit.TradableAmount)
//					require.NoError(err)
//					creditsDeposited, err = creditsDeposited.Add(dec)
//					require.NoError(err)
//				}
//
//				var val float64 = 10
//				for i := 1; uint32(i) <= tc.basket.Exponent; i++ {
//					val = math2.Pow(10, float64(i))
//				}
//				valStr := fmt.Sprintf("%f", val)
//				multiplierDec, err := math.NewDecFromString(valStr)
//				require.NoError(err)
//
//				tokensExpected, err := creditsDeposited.Mul(multiplierDec)
//				require.NoError(err)
//
//				// 0 == equals
//				require.Equal(0, tokensExpected.Cmp(actualTokensBack))
//
//				for _, c := range tc.msg.Credits {
//					balRes, err := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
//						Account:    lastBasket,
//						BatchDenom: c.BatchDenom,
//					})
//					require.NoError(err)
//
//					basketBalanceDec, err := math.NewDecFromString(balRes.TradableAmount)
//					require.NoError(err)
//
//					depositedAmountDec, err := math.NewDecFromString(c.TradableAmount)
//					require.NoError(err)
//					require.True(basketBalanceDec.Equal(depositedAmountDec))
//				}
//			}
//		})
//	}
//}
//
//func (s *IntegrationTestSuite) TestTakeFromBasketScenario() {
//	require := s.Require()
//	server := s.msgClient
//	admin := s.signers[0]
//
//	// create two batches
//	_, _, batch := s.createClassAndIssueBatch(admin.String(), "1000000000")
//	_, _, batch2 := s.createClassAndIssueBatch(admin.String(), "50000050500")
//
//	// create a basket with no criteria for simplicity
//	resBasket, err := server.CreateBasket(s.ctx, &ecocredit.MsgCreateBasket{
//		Curator:           admin.String(),
//		Name:              "testTakeFrom1",
//		DisplayName:       "ttf1",
//		Exponent:          1,
//		BasketCriteria:    nil,
//		DisableAutoRetire: true,
//		AllowPicking:      false,
//	})
//	require.NoError(err)
//	require.NotNil(resBasket)
//
//	basketDenom := resBasket.BasketDenom
//
//	// credits we are going to add to the basket
//	creditsAddedToBasket := []*ecocredit.BasketCredit{{BatchDenom: batch.BatchDenom, TradableAmount: "1"}, {BatchDenom: batch2.BatchDenom, TradableAmount: "1"}}
//
//	// add them to the basket - should pass
//	resAdd, err := server.AddToBasket(s.ctx, &ecocredit.MsgAddToBasket{
//		Owner:       admin.String(),
//		BasketDenom: basketDenom,
//		Credits:     creditsAddedToBasket,
//	})
//	require.NoError(err)
//	require.NotNil(resAdd)
//
//	// the basket exponent is 1 so -> 10^1 * creditDepositAmount = 10 * 2 = 20
//	expectedAmt := math.NewDecFromInt64(20)
//	amtReceived, err := math.NewDecFromString(resAdd.AmountReceived)
//	require.NoError(err)
//	require.True(expectedAmt.Equal(amtReceived))
//
//	// take credit from basket, should give us the first credit
//	resTake, err := server.TakeFromBasket(s.ctx, &ecocredit.MsgTakeFromBasket{
//		Owner:              admin.String(),
//		BasketDenom:        basketDenom,
//		Amount:             "1",
//		RetirementLocation: "",
//	})
//	require.NoError(err)
//	require.NotNil(resTake)
//
//	// it should take the oldest credit first, aka the first batch created
//	require.Equal(creditsAddedToBasket[0:1], resTake.Credits)
//
//	// check to see the credit as taken
//	queryRes, err := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
//		Account:    resBasket.BasketAddress,
//		BatchDenom: creditsAddedToBasket[0].BatchDenom,
//	})
//	require.NoError(err)
//	require.NotNil(queryRes)
//	require.Equal("0", queryRes.TradableAmount) // the first credit should be gone
//
//	// user should now have the credit
//	balRes, err := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
//		Account:    admin.String(),
//		BatchDenom: creditsAddedToBasket[0].BatchDenom,
//	})
//	require.NoError(err)
//	require.Equal(balRes.TradableAmount, "1000000000") // we minted 1000000000 to ourselves, deposited 1, and took it back.
//
//	// basket should still have the other credit left
//	queryRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
//		Account:    resBasket.BasketAddress,
//		BatchDenom: creditsAddedToBasket[1].BatchDenom,
//	})
//	require.NoError(err)
//	require.NotNil(queryRes)
//	require.Equal("1", queryRes.TradableAmount)
//
//	// user should now have 10 credits. swapping 1 = 10^1 * 1 = 10. 20 - 10 = 10.
//	basketTokenBalance := s.bankKeeper.GetBalance(s.sdkCtx, admin, basketDenom)
//	require.True(basketTokenBalance.Amount.Equal(sdk.NewInt(10)))
//
//	// trash some coins to check fail case
//	coins := sdk.NewCoins(sdk.NewCoin(basketDenom, sdk.NewInt(3)))
//	err = s.bankKeeper.SendCoinsFromAccountToModule(s.sdkCtx, admin, ecocredit.ModuleName, coins)
//	require.NoError(err)
//
//	// make sure we don't have enough to complete a swap
//	basketTokenBalance = s.bankKeeper.GetBalance(s.sdkCtx, admin, basketDenom)
//	require.True(basketTokenBalance.Amount.Equal(sdk.NewInt(7)))
//
//	// try to take again, but should fail cause of insufficient basket tokens, we need at least 10 for 1 ecocredit.
//	resTake2, err := server.TakeFromBasket(s.ctx, &ecocredit.MsgTakeFromBasket{
//		Owner:              admin.String(),
//		BasketDenom:        basketDenom,
//		Amount:             "1",
//		RetirementLocation: "",
//	})
//	require.Error(err)
//	require.Nil(resTake2)
//	require.Contains(err.Error(), "insufficient basket token balance, got: 7, needed at least: 10")
//
//	// get the tokens back so we can try to take again.
//	require.NoError(s.bankKeeper.SendCoinsFromModuleToAccount(s.sdkCtx, ecocredit.ModuleName, admin, coins))
//	balanceAfter := s.bankKeeper.GetBalance(s.sdkCtx, admin, basketDenom)
//	require.True(balanceAfter.Amount.Equal(sdk.NewInt(10)))
//
//	// try to take again, but ask for more than the basket has - should fail.
//	resTake3, err := server.TakeFromBasket(s.ctx, &ecocredit.MsgTakeFromBasket{
//		Owner:              admin.String(),
//		BasketDenom:        basketDenom,
//		Amount:             "25",
//		RetirementLocation: "",
//	})
//	require.Error(err)
//	require.Nil(resTake3)
//
//	// take the final credit from the basket
//	resTake4, err := server.TakeFromBasket(s.ctx, &ecocredit.MsgTakeFromBasket{
//		Owner:              admin.String(),
//		BasketDenom:        basketDenom,
//		Amount:             "1",
//		RetirementLocation: "",
//	})
//	require.NoError(err)
//	require.NotNil(resTake4)
//	require.Equal(resTake4.Credits, creditsAddedToBasket[1:])
//}

func (s *IntegrationTestSuite) TestPickFromBasket() {
	require := s.Require()
	server := s.msgClient
	admin := s.signers[0]

	// create two batches
	_, _, batch := s.createClassAndIssueBatch(admin.String(), "10000000000000")
	_, _, batch2 := s.createClassAndIssueBatch(admin.String(), "10000000000000")
	_, _, batchForErrors := s.createClassAndIssueBatch(admin.String(), "10000000000000")

	const exponent uint32 = 2
	retireBasket, err := server.CreateBasket(s.ctx, &ecocredit.MsgCreateBasket{
		Curator:           admin.String(),
		Name:              "TestPickBasket",
		DisplayName:       "TPB",
		Exponent:          exponent,
		BasketCriteria:    nil,
		DisableAutoRetire: false,
		AllowPicking:      true,
	})
	require.NoError(err)

	noRetireBasket, err := server.CreateBasket(s.ctx, &ecocredit.MsgCreateBasket{
		Curator:           admin.String(),
		Name:              "TestPickBasketII",
		DisplayName:       "TPBII",
		Exponent:          exponent,
		BasketCriteria:    nil,
		DisableAutoRetire: true,
		AllowPicking:      true,
	})
	require.NoError(err)

	testBasket, err := server.CreateBasket(s.ctx, &ecocredit.MsgCreateBasket{
		Curator:           admin.String(),
		Name:              "TestPickBasketIII",
		DisplayName:       "TPBIII",
		Exponent:          exponent,
		BasketCriteria:    nil,
		DisableAutoRetire: false,
		AllowPicking:      false,
	})
	require.NoError(err)

	denomToAddr := map[string]string{retireBasket.BasketDenom: retireBasket.BasketAddress, noRetireBasket.BasketDenom: noRetireBasket.BasketAddress, testBasket.BasketDenom: testBasket.BasketAddress}

	_, err = server.AddToBasket(s.ctx, &ecocredit.MsgAddToBasket{
		Owner:       admin.String(),
		BasketDenom: retireBasket.BasketDenom,
		Credits: []*ecocredit.BasketCredit{
			{BatchDenom: batch.BatchDenom, TradableAmount: "1000"},
			{BatchDenom: batch2.BatchDenom, TradableAmount: "1000"},
		},
	})
	require.NoError(err)

	_, err = server.AddToBasket(s.ctx, &ecocredit.MsgAddToBasket{
		Owner:       admin.String(),
		BasketDenom: noRetireBasket.BasketDenom,
		Credits: []*ecocredit.BasketCredit{
			{BatchDenom: batch.BatchDenom, TradableAmount: "1000"},
			{BatchDenom: batch2.BatchDenom, TradableAmount: "1000"},
		},
	})
	require.NoError(err)

	_, err = server.AddToBasket(s.ctx, &ecocredit.MsgAddToBasket{
		Owner:       admin.String(),
		BasketDenom: testBasket.BasketDenom,
		Credits: []*ecocredit.BasketCredit{
			{BatchDenom: batchForErrors.BatchDenom, TradableAmount: "1000"},
		},
	})
	require.NoError(err)

	testCases := []struct {
		name        string
		basketDenom string
		retire      bool
		msg         *ecocredit.MsgPickFromBasket
		expErr      bool
		errMsg      string
	}{
		{
			name:        "valid pick from a non-retire basket",
			basketDenom: noRetireBasket.BasketDenom,
			retire:      false,
			msg: &ecocredit.MsgPickFromBasket{
				Owner:              admin.String(),
				BasketDenom:        noRetireBasket.BasketDenom,
				Credits:            []*ecocredit.BasketCredit{{BatchDenom: batch.BatchDenom, TradableAmount: "3"}},
				RetirementLocation: "",
			},
		},
		{
			name:        "valid pick from a retire basket",
			basketDenom: retireBasket.BasketDenom,
			retire:      true,
			msg: &ecocredit.MsgPickFromBasket{
				Owner:              admin.String(),
				BasketDenom:        retireBasket.BasketDenom,
				Credits:            []*ecocredit.BasketCredit{{BatchDenom: batch.BatchDenom, TradableAmount: "3"}},
				RetirementLocation: "YZ",
			},
		},
		{
			name:        "invalid - insufficient funds",
			basketDenom: noRetireBasket.BasketDenom,
			retire:      false,
			msg: &ecocredit.MsgPickFromBasket{
				Owner:              admin.String(),
				BasketDenom:        noRetireBasket.BasketDenom,
				Credits:            []*ecocredit.BasketCredit{{BatchDenom: batch.BatchDenom, TradableAmount: "1000000000000000000000"}},
				RetirementLocation: "",
			},
			expErr: true,
			errMsg: fmt.Sprintf("requested 1000000000000000000000 credits but basket %s only has ", noRetireBasket.BasketDenom), // cut out the last part of the error to decrease complexity
		},
		{
			name:        "invalid - no picking",
			basketDenom: testBasket.BasketDenom,
			retire:      false,
			msg: &ecocredit.MsgPickFromBasket{
				Owner:              admin.String(),
				BasketDenom:        testBasket.BasketDenom,
				Credits:            []*ecocredit.BasketCredit{{BatchDenom: batch.BatchDenom, TradableAmount: "1"}},
				RetirementLocation: "",
			},
			expErr: true,
			errMsg: fmt.Sprintf("basket %s has auto-retirement enabled, but the request did not include a retirement location.", testBasket.BasketDenom), // cut out the last part of the error to decrease complexity
		},
		{
			name:        "valid partial credits tradable",
			basketDenom: noRetireBasket.BasketDenom,
			retire:      false,
			msg: &ecocredit.MsgPickFromBasket{
				Owner:              admin.String(),
				BasketDenom:        noRetireBasket.BasketDenom,
				Credits:            []*ecocredit.BasketCredit{{BatchDenom: batch.BatchDenom, TradableAmount: "3.250"}},
				RetirementLocation: "YZ",
			},
		},
		{
			name:        "valid partial credits retired",
			basketDenom: retireBasket.BasketDenom,
			retire:      true,
			msg: &ecocredit.MsgPickFromBasket{
				Owner:              admin.String(),
				BasketDenom:        retireBasket.BasketDenom,
				Credits:            []*ecocredit.BasketCredit{{BatchDenom: batch.BatchDenom, TradableAmount: "3.250"}},
				RetirementLocation: "YZ",
			},
		},
		{
			name:        "invalid - exceeds batch precision",
			basketDenom: noRetireBasket.BasketDenom,
			retire:      false,
			msg: &ecocredit.MsgPickFromBasket{
				Owner:              admin.String(),
				BasketDenom:        noRetireBasket.BasketDenom,
				Credits:            []*ecocredit.BasketCredit{{BatchDenom: batch.BatchDenom, TradableAmount: "3.25013587299818"}},
				RetirementLocation: "",
			},
			expErr: true,
			errMsg: "3.25013587299818 exceeds maximum decimal places",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {

			// basket balance
			basketBeforeTradable, _ := s.getCreditBalance(denomToAddr[tc.basketDenom], tc.msg.Credits[0].BatchDenom)

			// user balances
			userBeforeTradable, userBeforeRetired := s.getCreditBalance(admin.String(), tc.msg.Credits[0].BatchDenom)
			userTokenBalanceBefore := s.getBasketTokenBalance(admin, tc.basketDenom)

			// run the msg tx
			_, err = s.msgClient.PickFromBasket(s.ctx, tc.msg)

			if tc.expErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				for _, creditRequested := range tc.msg.Credits {
					require.NoError(err)
					amountTakenDec, _ := math.NewDecFromString(creditRequested.TradableAmount)

					// ensure the basket has balanceBefore - amountTaken.
					// [important: baskets only have tradable credits. they will never have retired balances]
					basketAfterTradable, _ := s.getCreditBalance(denomToAddr[tc.basketDenom], creditRequested.BatchDenom)
					newBalanceBasket, err := basketBeforeTradable.Sub(amountTakenDec)
					require.NoError(err)
					require.True(basketAfterTradable.Equal(newBalanceBasket), fmt.Sprintf("expected basket tradable to be %s, got: %s", basketAfterTradable.String(), newBalanceBasket.String()))

					// ensure the user now has balanceBefore + amountTaken (either retired or
					userAfterTradable, userAfterRetired := s.getCreditBalance(admin.String(), creditRequested.BatchDenom)
					if tc.retire {
						newBalanceUser, err := userBeforeRetired.Add(amountTakenDec)
						require.NoError(err)
						require.True(newBalanceUser.Equal(userAfterRetired))
					} else {
						newBalanceUser, err := userBeforeTradable.Add(amountTakenDec)
						require.NoError(err)
						require.True(newBalanceUser.Equal(userAfterTradable), fmt.Sprintf("expected user credit balance to be %s, got %s", newBalanceUser.String(), userAfterTradable.String()))
					}

					// ensure the user has tokenBalanceBefore - CalculateBasketTokens(amountTaken, exponent)
					userTokenBalanceAfter := s.getBasketTokenBalance(admin, tc.basketDenom)
					swapCost, err := server2.CalculateBasketTokens(amountTakenDec, exponent)
					require.NoError(err)
					expectedBalanceAfter, err := userTokenBalanceBefore.Sub(swapCost)
					require.NoError(err)
					require.True(expectedBalanceAfter.Equal(userTokenBalanceAfter))
				}
			}
		})
	}

	//// no retire basket
	//basket, err := server.CreateBasket(s.ctx, &ecocredit.MsgCreateBasket{
	//	Curator:           admin.String(),
	//	Name:              "TestPickFromBasket",
	//	DisplayName:       "TPFB1",
	//	Exponent:          3,
	//	BasketCriteria:    nil,
	//	DisableAutoRetire: false,
	//	AllowPicking:      true,
	//})
	//require.NoError(err)
	//require.NotNil(basket)
	//
	//originalDepositAmt := math.NewDecFromInt64(10)
	//addRes, err := server.AddToBasket(s.ctx, &ecocredit.MsgAddToBasket{
	//	Owner:       admin.String(),
	//	BasketDenom: basket.BasketDenom,
	//	Credits: []*ecocredit.BasketCredit{
	//		{BatchDenom: batch.BatchDenom, TradableAmount: "10"},
	//		{BatchDenom: batch2.BatchDenom, TradableAmount: "10"},
	//	},
	//})
	//require.NoError(err)
	//require.NotNil(addRes)
	//
	//pickAmount, err := math.NewDecFromString("0.5")
	//require.NoError(err)
	//pickRes, err := server.PickFromBasket(s.ctx, &ecocredit.MsgPickFromBasket{
	//	Owner:       admin.String(),
	//	BasketDenom: basket.BasketDenom,
	//	Credits: []*ecocredit.BasketCredit{
	//		{BatchDenom: batch.BatchDenom, TradableAmount: pickAmount.String()},
	//	},
	//	RetirementLocation: "YZ",
	//})
	//require.NoError(err)
	//require.NotNil(pickRes)
	//
	//res, err := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
	//	Account:    basket.BasketAddress,
	//	BatchDenom: batch.BatchDenom,
	//})
	////res, err := s.queryClient.BasketBatchBalance(s.ctx, &ecocredit.QueryBasketBatchBalanceRequest{
	////	BasketDenom: basket.BasketDenom,
	////	BatchDenom:  batch.BatchDenom,
	////})
	//require.NoError(err)
	//require.NotNil(res)
	//
	//basketBalDec, err := math.NewDecFromString(res.TradableAmount)
	//require.NoError(err)
	//
	//newBal, err := originalDepositAmt.Sub(pickAmount)
	//require.NoError(err)
	//
	//require.True(newBal.Equal(basketBalDec))
	//
	//res1, err := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
	//	Account:    admin.String(),
	//	BatchDenom: batch.BatchDenom,
	//})
	//require.NoError(err)
	//require.NotNil(res1)
	//fmt.Println(res1)
	//require.Equal(res1.RetiredAmount, pickAmount.String())

}

func (s *IntegrationTestSuite) getCreditBalance(addr string, denom string) (tradable, retired math.Dec) {
	res, err := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
		Account:    addr,
		BatchDenom: denom,
	})
	s.Require().NoError(err)

	tradable, _ = math.NewDecFromString(res.TradableAmount)
	retired, _ = math.NewDecFromString(res.RetiredAmount)

	return tradable, retired
}

func (s *IntegrationTestSuite) getBasketTokenBalance(addr sdk.AccAddress, denom string) math.Dec {
	coin := s.bankKeeper.GetBalance(s.sdkCtx, addr, denom)
	amtDec, err := math.NewDecFromString(coin.Amount.String())
	s.Require().NoError(err)
	return amtDec
}
