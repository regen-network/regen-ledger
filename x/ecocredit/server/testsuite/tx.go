package testsuite

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s *IntegrationTestSuite) TestScenarioCreateSellOrders() {
	addr1 := s.signers[3]

	// create credit class and issue credits to addr1
	_, batchDenom := s.createClassAndIssueBatch(addr1, addr1, "carbon", "2.0", "2020-01-01", "2022-01-01")

	askPrice1 := sdk.NewInt64Coin("stake", 1000000)
	// TODO: Verify that AskPrice.Denom is in AllowAskDenom #624
	//askPrice2 := sdk.NewInt64Coin("token", 1000000)

	expiration := time.Date(2030, 01, 01, 0, 0, 0, 0, time.UTC)
	invalidExpiration := time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)

	// create sell orders
	testCases := []struct {
		name    string
		owner   string
		orders  []*ecocredit.MsgSell_Order
		expErr  string
		wantErr bool
	}{
		{
			name:  "insufficient credit balance - batch denom",
			owner: addr1.String(),
			orders: []*ecocredit.MsgSell_Order{
				{
					BatchDenom:        "A00-00000000-00000000-000",
					Quantity:          "1.0",
					AskPrice:          &askPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
				{
					BatchDenom:        "A00-00000000-00000000-000",
					Quantity:          "1.0",
					AskPrice:          &askPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
			},
			expErr:  "insufficient credit balance",
			wantErr: true,
		},
		{
			name:  "insufficient credit balance - quantity",
			owner: addr1.String(),
			orders: []*ecocredit.MsgSell_Order{
				{
					BatchDenom:        batchDenom,
					Quantity:          "99",
					AskPrice:          &askPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
				{
					BatchDenom:        batchDenom,
					Quantity:          "99",
					AskPrice:          &askPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
			},
			expErr:  "insufficient credit balance",
			wantErr: true,
		},
		{
			name:  "invalid expiration",
			owner: addr1.String(),
			orders: []*ecocredit.MsgSell_Order{
				{
					BatchDenom:        batchDenom,
					Quantity:          "1.0",
					AskPrice:          &askPrice1,
					DisableAutoRetire: true,
					Expiration:        &invalidExpiration,
				},
				{
					BatchDenom:        batchDenom,
					Quantity:          "1.0",
					AskPrice:          &askPrice1,
					DisableAutoRetire: true,
					Expiration:        &invalidExpiration,
				},
			},
			expErr:  "expiration must be in the future",
			wantErr: true,
		},
		// TODO: Verify that AskPrice.Denom is in AllowAskDenom #624
		//{
		//	name: "denom not allowed",
		//	owner: addr1.String(),
		//	orders: []*ecocredit.MsgSell_Order{
		//		{
		//			BatchDenom:        batchDenom,
		//			Quantity:          "1.0",
		//			AskPrice:          &askPrice2,
		//			DisableAutoRetire: true,
		//			Expiration:        &expiration,
		//		},
		//		{
		//			BatchDenom:        batchDenom,
		//			Quantity:          "1.0",
		//			AskPrice:          &askPrice2,
		//			DisableAutoRetire: true,
		//			Expiration:        &expiration,
		//		},
		//	},
		//	expErr: "denom not allowed",
		//	wantErr: true,
		//},
		{
			name:  "valid request",
			owner: addr1.String(),
			orders: []*ecocredit.MsgSell_Order{
				{
					BatchDenom:        batchDenom,
					Quantity:          "1.0",
					AskPrice:          &askPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
				{
					BatchDenom:        batchDenom,
					Quantity:          "1.0",
					AskPrice:          &askPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
			},
			expErr:  "",
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			require := s.Require()

			res, err := s.msgClient.Sell(s.ctx, &ecocredit.MsgSell{
				Owner:  tc.owner,
				Orders: tc.orders,
			})

			if tc.wantErr {
				require.Error(err)
				require.Contains(err.Error(), tc.expErr)
			} else {
				require.NoError(err)
				require.NotNil(res.SellOrderIds)

				// query first sell order
				_, sellError1 := s.queryClient.SellOrder(s.ctx, &ecocredit.QuerySellOrderRequest{
					SellOrderId: res.SellOrderIds[0],
				})

				// query second sell order
				_, sellError2 := s.queryClient.SellOrder(s.ctx, &ecocredit.QuerySellOrderRequest{
					SellOrderId: res.SellOrderIds[1],
				})

				require.NoError(sellError1)
				require.NoError(sellError2)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestScenarioUpdateSellOrders() {
	addr1 := s.signers[3]
	addr2 := s.signers[4]

	// create credit class and issue credits to addr1
	_, batchDenom := s.createClassAndIssueBatch(addr1, addr1, "carbon", "2.0", "2020-01-01", "2022-01-01")

	askPrice1 := sdk.NewInt64Coin("stake", 2000000)
	// TODO: Verify that NewAskPrice.Denom is in AllowAskDenom #624
	//askPrice2 := sdk.NewInt64Coin("token", 2000000)

	expiration := time.Date(2030, 01, 01, 0, 0, 0, 0, time.UTC)
	invalidExpiration := time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)

	// create sell order
	sellRes, err := s.msgClient.Sell(s.ctx, &ecocredit.MsgSell{
		Owner: addr1.String(),
		Orders: []*ecocredit.MsgSell_Order{
			{
				BatchDenom:        batchDenom,
				Quantity:          "1.0",
				AskPrice:          &askPrice1,
				DisableAutoRetire: true,
				Expiration:        &expiration,
			},
			{
				BatchDenom:        batchDenom,
				Quantity:          "1.0",
				AskPrice:          &askPrice1,
				DisableAutoRetire: true,
				Expiration:        &expiration,
			},
		},
	})
	s.Require().NoError(err)

	// update sell orders
	testCases := []struct {
		name    string
		owner   string
		updates []*ecocredit.MsgUpdateSellOrders_Update
		expErr  string
		wantErr bool
	}{
		{
			name:  "invalid sell order",
			owner: addr1.String(),
			updates: []*ecocredit.MsgUpdateSellOrders_Update{
				{
					SellOrderId:       99,
					NewQuantity:       "1.0",
					NewAskPrice:       &askPrice1,
					DisableAutoRetire: true,
					NewExpiration:     &expiration,
				},
				{
					SellOrderId:       100,
					NewQuantity:       "1.0",
					NewAskPrice:       &askPrice1,
					DisableAutoRetire: true,
					NewExpiration:     &expiration,
				},
			},
			expErr:  "invalid sell order",
			wantErr: true,
		},
		{
			name:  "unauthorized",
			owner: addr2.String(),
			updates: []*ecocredit.MsgUpdateSellOrders_Update{
				{
					SellOrderId:       sellRes.SellOrderIds[0],
					NewQuantity:       "1.0",
					NewAskPrice:       &askPrice1,
					DisableAutoRetire: true,
					NewExpiration:     &expiration,
				},
				{
					SellOrderId:       sellRes.SellOrderIds[1],
					NewQuantity:       "1.0",
					NewAskPrice:       &askPrice1,
					DisableAutoRetire: true,
					NewExpiration:     &expiration,
				},
			},
			expErr:  "unauthorized",
			wantErr: true,
		},
		{
			name:  "insufficient credit balance",
			owner: addr1.String(),
			updates: []*ecocredit.MsgUpdateSellOrders_Update{
				{
					SellOrderId:       sellRes.SellOrderIds[0],
					NewQuantity:       "99",
					NewAskPrice:       &askPrice1,
					DisableAutoRetire: true,
					NewExpiration:     &expiration,
				},
				{
					SellOrderId:       sellRes.SellOrderIds[1],
					NewQuantity:       "99",
					NewAskPrice:       &askPrice1,
					DisableAutoRetire: true,
					NewExpiration:     &expiration,
				},
			},
			expErr:  "insufficient credit balance",
			wantErr: true,
		},
		{
			name:  "invalid expiration",
			owner: addr1.String(),
			updates: []*ecocredit.MsgUpdateSellOrders_Update{
				{
					SellOrderId:       sellRes.SellOrderIds[0],
					NewQuantity:       "1.0",
					NewAskPrice:       &askPrice1,
					DisableAutoRetire: true,
					NewExpiration:     &invalidExpiration,
				},
				{
					SellOrderId:       sellRes.SellOrderIds[1],
					NewQuantity:       "1.0",
					NewAskPrice:       &askPrice1,
					DisableAutoRetire: true,
					NewExpiration:     &invalidExpiration,
				},
			},
			expErr:  "expiration must be in the future",
			wantErr: true,
		},
		// TODO: Verify that NewAskPrice.Denom is in AllowAskDenom #624
		//{
		//	name: "denom not allowed",
		//	owner: addr1.String(),
		//	updates: []*ecocredit.MsgUpdateSellOrders_Update{
		//		{
		//			SellOrderId:       sellRes.SellOrderIds[0],
		//			NewQuantity:       "1.0",
		//			NewAskPrice:       &askPrice2,
		//			DisableAutoRetire: true,
		//			NewExpiration:     &expiration,
		//		},
		//		{
		//			SellOrderId:       sellRes.SellOrderIds[1],
		//			NewQuantity:       "1.0",
		//			NewAskPrice:       &askPrice2,
		//			DisableAutoRetire: true,
		//			NewExpiration:     &expiration,
		//		},
		//	},
		//	expErr: "denom not allowed",
		//	wantErr: true,
		//},
		{
			name:  "valid request",
			owner: addr1.String(),
			updates: []*ecocredit.MsgUpdateSellOrders_Update{
				{
					SellOrderId:       sellRes.SellOrderIds[0],
					NewQuantity:       "1.0",
					NewAskPrice:       &askPrice1,
					DisableAutoRetire: true,
					NewExpiration:     &expiration,
				},
				{
					SellOrderId:       sellRes.SellOrderIds[1],
					NewQuantity:       "1.0",
					NewAskPrice:       &askPrice1,
					DisableAutoRetire: true,
					NewExpiration:     &expiration,
				},
			},
			expErr:  "",
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			require := s.Require()

			_, err := s.msgClient.UpdateSellOrders(s.ctx, &ecocredit.MsgUpdateSellOrders{
				Owner:   tc.owner,
				Updates: tc.updates,
			})

			if tc.wantErr {
				require.Error(err)
				require.Contains(err.Error(), tc.expErr)
			} else {
				require.NoError(err)

				// query first sell order
				sellResponse1, sellError1 := s.queryClient.SellOrder(s.ctx, &ecocredit.QuerySellOrderRequest{
					SellOrderId: tc.updates[0].SellOrderId,
				})

				// query second sell order
				sellResponse2, sellError2 := s.queryClient.SellOrder(s.ctx, &ecocredit.QuerySellOrderRequest{
					SellOrderId: tc.updates[1].SellOrderId,
				})

				require.NoError(sellError1)
				require.NoError(sellError2)
				require.Equal(tc.updates[0].NewAskPrice, sellResponse1.SellOrder.AskPrice)
				require.Equal(tc.updates[1].NewAskPrice, sellResponse2.SellOrder.AskPrice)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestScenarioCreateBuyOrders() {
	addr1 := s.signers[3]
	addr2 := s.signers[4]

	// create credit class and issue credits to addr1
	_, batchDenom := s.createClassAndIssueBatch(addr1, addr1, "carbon", "6.0", "2020-01-01", "2022-01-01")

	bidPrice1 := sdk.NewInt64Coin("stake", 1000000)
	bidPrice2 := sdk.NewInt64Coin("stake", 9999999)
	// TODO: Verify that BidPrice.Denom is in AllowAskDenom #624
	//bidPrice3 := sdk.NewInt64Coin("token", 1000000)

	expiration := time.Date(2030, 01, 01, 0, 0, 0, 0, time.UTC)
	invalidExpiration := time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)

	// fund buyer account
	s.fundAccount(addr2, sdk.NewCoins(sdk.NewInt64Coin("stake", 5000000)))

	// create sell orders
	sellRes, err := s.msgClient.Sell(s.ctx, &ecocredit.MsgSell{
		Owner: addr1.String(),
		Orders: []*ecocredit.MsgSell_Order{
			{
				BatchDenom:        batchDenom,
				Quantity:          "1.0",
				AskPrice:          &bidPrice1,
				DisableAutoRetire: true,
				Expiration:        &expiration,
			},
			{
				BatchDenom:        batchDenom,
				Quantity:          "1.0",
				AskPrice:          &bidPrice1,
				DisableAutoRetire: true,
				Expiration:        &expiration,
			},
			{
				BatchDenom:        batchDenom,
				Quantity:          "1.0",
				AskPrice:          &bidPrice1,
				DisableAutoRetire: false,
			},
			{
				BatchDenom:        batchDenom,
				Quantity:          "1.0",
				AskPrice:          &bidPrice1,
				DisableAutoRetire: false,
			},
			{
				BatchDenom:        batchDenom,
				Quantity:          "1.0",
				AskPrice:          &bidPrice1,
				DisableAutoRetire: true,
				Expiration:        &expiration,
			},
			{
				BatchDenom:        batchDenom,
				Quantity:          "1.0",
				AskPrice:          &bidPrice1,
				DisableAutoRetire: true,
				Expiration:        &expiration,
			},
		},
	})
	s.Require().NoError(err)

	// process buy orders
	testCases := []struct {
		name             string
		buyer            string
		orders           []*ecocredit.MsgBuy_Order
		expErr           string
		wantErr          bool
		partial          bool
		expCoinBalance   sdk.Coin
		expCreditBalance *ecocredit.QueryBalanceResponse
	}{
		{
			name:  "invalid sell order",
			buyer: addr2.String(),
			orders: []*ecocredit.MsgBuy_Order{
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: 99}},
					Quantity:          "1.0",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: 100}},
					Quantity:          "1.0",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
			},
			expErr:  "not found",
			wantErr: true,
		},
		{
			name:  "insufficient coin balance - quantity",
			buyer: addr2.String(),
			orders: []*ecocredit.MsgBuy_Order{
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[0]}},
					Quantity:          "99.99",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[1]}},
					Quantity:          "99.99",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
			},
			expErr:  "insufficient balance",
			wantErr: true,
		},
		{
			name:  "insufficient coin balance - bid price",
			buyer: addr2.String(),
			orders: []*ecocredit.MsgBuy_Order{
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[0]}},
					Quantity:          "1.0",
					BidPrice:          &bidPrice2,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[1]}},
					Quantity:          "1.0",
					BidPrice:          &bidPrice2,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
			},
			expErr:  "insufficient balance",
			wantErr: true,
		},
		{
			name:  "invalid expiration",
			buyer: addr2.String(),
			orders: []*ecocredit.MsgBuy_Order{
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[0]}},
					Quantity:          "1.0",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: true,
					Expiration:        &invalidExpiration,
				},
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[1]}},
					Quantity:          "1.0",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: true,
					Expiration:        &invalidExpiration,
				},
			},
			expErr:  "expiration must be in the future",
			wantErr: true,
		},
		// TODO: Verify that BidPrice.Denom is in AllowAskDenom #624
		//{
		//	name: "denom not allowed",
		//	buyer: addr2.String(),
		//	orders: []*ecocredit.MsgBuy_Order{
		//		{
		//			Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[0]}},
		//			Quantity:          "1.0",
		//			BidPrice:          &bidPrice3,
		//			DisableAutoRetire: true,
		//			Expiration: &expiration,
		//		},
		//		{
		//			Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[1]}},
		//			Quantity:          "1.0",
		//			BidPrice:          &bidPrice3,
		//			DisableAutoRetire: true,
		//			Expiration: &expiration,
		//		},
		//	},
		//	expErr: "denom not allowed",
		//	wantErr: true,
		//},
		{
			name:  "auto-retire is required for sell order",
			buyer: addr2.String(),
			orders: []*ecocredit.MsgBuy_Order{
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[2]}},
					Quantity:          "1.0",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: true,
				},
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[3]}},
					Quantity:          "1.0",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: true,
				},
			},
			expErr:  "auto-retire is required for sell order",
			wantErr: true,
		},
		{
			name:  "retirement location is required for sell order",
			buyer: addr2.String(),
			orders: []*ecocredit.MsgBuy_Order{
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[2]}},
					Quantity:          "1.0",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: false,
				},
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[3]}},
					Quantity:          "1.0",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: false,
				},
			},
			expErr:  "retirement location is required for sell order",
			wantErr: true,
		},
		{
			name:  "valid request",
			buyer: addr2.String(),
			orders: []*ecocredit.MsgBuy_Order{
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[0]}},
					Quantity:          "1.0",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[1]}},
					Quantity:          "1.0",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
			},
			expErr:  "",
			wantErr: false,
			partial: false,
			expCoinBalance: sdk.Coin{
				Denom:  "stake",
				Amount: sdk.NewInt(3000000),
			},
			expCreditBalance: &ecocredit.QueryBalanceResponse{TradableAmount: "2", RetiredAmount: "0"},
		},
		{
			name:  "valid request - auto-retire",
			buyer: addr2.String(),
			orders: []*ecocredit.MsgBuy_Order{
				{
					Selection:          &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[2]}},
					Quantity:           "1.0",
					BidPrice:           &bidPrice1,
					DisableAutoRetire:  false,
					RetirementLocation: "AB",
				},
				{
					Selection:          &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[3]}},
					Quantity:           "1.0",
					BidPrice:           &bidPrice1,
					DisableAutoRetire:  false,
					RetirementLocation: "AB",
				},
			},
			expErr:  "",
			wantErr: false,
			partial: false,
			expCoinBalance: sdk.Coin{
				Denom:  "stake",
				Amount: sdk.NewInt(1000000),
			},
			expCreditBalance: &ecocredit.QueryBalanceResponse{TradableAmount: "2", RetiredAmount: "2"},
		},
		{
			name:  "valid request - partial fill",
			buyer: addr2.String(),
			orders: []*ecocredit.MsgBuy_Order{
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[4]}},
					Quantity:          "0.5",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
				{
					Selection:         &ecocredit.MsgBuy_Order_Selection{Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: sellRes.SellOrderIds[5]}},
					Quantity:          "0.5",
					BidPrice:          &bidPrice1,
					DisableAutoRetire: true,
					Expiration:        &expiration,
				},
			},
			expErr:  "",
			wantErr: false,
			partial: true,
			expCoinBalance: sdk.Coin{
				Denom:  "stake",
				Amount: sdk.NewInt(0),
			},
			expCreditBalance: &ecocredit.QueryBalanceResponse{TradableAmount: "3", RetiredAmount: "2"},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			require := s.Require()

			// get buyer coin balance before
			coinBalanceBefore := s.bankKeeper.GetBalance(s.sdkCtx, addr2, "stake")

			// get buyer credit balance before
			creditBalanceBefore, _ := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
				Account:    addr2.String(),
				BatchDenom: batchDenom,
			})

			// process buy orders
			res, err := s.msgClient.Buy(s.ctx, &ecocredit.MsgBuy{
				Buyer:  tc.buyer,
				Orders: tc.orders,
			})

			// get buyer coin balance after
			coinBalanceAfter := s.bankKeeper.GetBalance(s.sdkCtx, addr2, "stake")

			// get buyer credit balance after
			creditBalanceAfter, _ := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
				Account:    addr2.String(),
				BatchDenom: batchDenom,
			})

			// query first sell order
			sellResponse1, sellError1 := s.queryClient.SellOrder(s.ctx, &ecocredit.QuerySellOrderRequest{
				SellOrderId: tc.orders[0].Selection.GetSellOrderId(),
			})

			// query second sell order
			sellResponse2, sellError2 := s.queryClient.SellOrder(s.ctx, &ecocredit.QuerySellOrderRequest{
				SellOrderId: tc.orders[1].Selection.GetSellOrderId(),
			})

			if tc.wantErr {
				require.Error(err)
				require.Contains(err.Error(), tc.expErr)
				require.Equal(coinBalanceBefore, coinBalanceAfter)
				require.Equal(creditBalanceBefore, creditBalanceAfter)
			} else {
				require.NoError(err)
				require.NotNil(res.BuyOrderIds)

				require.Equal(tc.expCoinBalance, coinBalanceAfter)
				require.Equal(tc.expCreditBalance, creditBalanceAfter)

				if tc.partial {
					require.NotNil(sellResponse1)
					require.NotNil(sellResponse2)
					require.NoError(sellError1)
					require.NoError(sellError2)
				} else {
					require.Nil(sellResponse1)
					require.Nil(sellResponse2)
					require.Error(sellError1)
					require.Error(sellError2)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestScenarioAllowAskDenom() {
	addr1 := s.signers[3].String()

	// TODO: Verify governance module address for AllowAskDenom #624
	//rootAddress := s.accountKeeper.GetModuleAddress(govtypes.ModuleName).String()

	// add ask denom
	testCases := []struct {
		name         string
		rootAddress  string
		denom        string
		displayDenom string
		exponent     uint32
		expErr       string
		wantErr      bool
	}{
		{
			name:         "unauthorized address",
			rootAddress:  addr1,
			denom:        "utoken",
			displayDenom: "token",
			exponent:     6,
			expErr:       "unauthorized",
			wantErr:      true,
		},
		// TODO: Verify governance module address for AllowAskDenom #624
		//{
		//	name: "valid request",
		//	rootAddress: rootAddress,
		//	denom: "utoken",
		//	displayDenom: "token",
		//	exponent: 6,
		//	expErr: "",
		//	wantErr: false,
		//},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			require := s.Require()

			res, err := s.msgClient.AllowAskDenom(s.ctx, &ecocredit.MsgAllowAskDenom{
				RootAddress:  tc.rootAddress,
				Denom:        tc.denom,
				DisplayDenom: tc.displayDenom,
				Exponent:     tc.exponent,
			})

			if tc.wantErr {
				require.Error(err)
				require.Contains(err.Error(), tc.expErr)
			} else {
				require.NoError(err)
				require.NotNil(res)
			}
		})
	}
}
