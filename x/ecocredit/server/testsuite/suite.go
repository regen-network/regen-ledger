package testsuite

import (
	"context"

	"github.com/regen-network/regen-ledger/types/testutil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory testutil.FixtureFactory
	fixture        testutil.Fixture

	sdkCtx            sdk.Context
	ctx               context.Context
	msgClient         ecocredit.MsgClient
	queryClient       ecocredit.QueryClient
	paramsQueryClient params.QueryClient
	signers           []sdk.AccAddress

	paramSpace  paramstypes.Subspace
	bankKeeper  bankkeeper.Keeper
}

func NewIntegrationTestSuite(fixtureFactory testutil.FixtureFactory, paramSpace paramstypes.Subspace, bankKeeper bankkeeper.BaseKeeper) *IntegrationTestSuite {
	return &IntegrationTestSuite{
		fixtureFactory: fixtureFactory,
		paramSpace:     paramSpace,
		bankKeeper:     bankKeeper,
	}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.fixture = s.fixtureFactory.Setup()

	// TODO clean up once types.Context merged upstream into sdk.Context
	s.sdkCtx, _ = s.fixture.Context().(types.Context).CacheContext()
	s.ctx = types.Context{Context: s.sdkCtx}

	ecocreditParams := ecocredit.DefaultParams()
	s.paramSpace.SetParamSet(s.sdkCtx, &ecocreditParams)

	s.signers = s.fixture.Signers()
	s.Require().GreaterOrEqual(len(s.signers), 6)
	s.msgClient = ecocredit.NewMsgClient(s.fixture.TxConn())
	s.queryClient = ecocredit.NewQueryClient(s.fixture.QueryConn())
	s.paramsQueryClient = params.NewQueryClient(s.fixture.QueryConn())
}

func fundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}
	return bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, amounts)
}

func (s *IntegrationTestSuite) TestScenario() {
	designer := s.signers[0]
	issuer1 := s.signers[1].String()
	issuer2 := s.signers[2].String()
	addr1 := s.signers[3].String()
	addr2 := s.signers[4].String()
	addr3 := s.signers[5].String()

	// create class with insufficient funds and it should fail
	createClsRes, err := s.msgClient.CreateClass(s.ctx, &ecocredit.MsgCreateClassRequest{
		Designer: designer.String(),
		Issuers:  []string{issuer1, issuer2},
		Metadata: nil,
	})
	s.Require().Error(err)
	s.Require().Nil(createClsRes)

	// create class with sufficient funds and it should succeed
	s.Require().NoError(fundAccount(s.bankKeeper, s.sdkCtx, designer, sdk.NewCoins(sdk.NewInt64Coin("stake", 10000))))

	createClsRes, err = s.msgClient.CreateClass(s.ctx, &ecocredit.MsgCreateClassRequest{
		Designer: designer.String(),
		Issuers:  []string{issuer1, issuer2},
		Metadata: nil,
	})
	s.Require().NoError(err)
	s.Require().NotNil(createClsRes)

	clsID := createClsRes.ClassId
	s.Require().NotEmpty(clsID)

	// designer should have no funds remaining
	s.Require().Equal(s.bankKeeper.GetBalance(s.sdkCtx, designer, "stake"), sdk.NewInt64Coin("stake", 0))

	// create batch
	t0, t1 := "10.37", "1007.3869"
	tSupply0 := "1017.7569"
	r0, r1 := "4.286", "10000.4589902"
	rSupply0 := "10004.7449902"

	createBatchRes, err := s.msgClient.CreateBatch(s.ctx, &ecocredit.MsgCreateBatchRequest{
		Issuer:  issuer1,
		ClassId: clsID,
		Issuance: []*ecocredit.MsgCreateBatchRequest_BatchIssuance{
			{
				Recipient:     addr1,
				TradableUnits: t0,
				RetiredUnits:  r0,
			},
			{
				Recipient:     addr2,
				TradableUnits: t1,
				RetiredUnits:  r1,
			},
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(createBatchRes)

	batchDenom := createBatchRes.BatchDenom
	s.Require().NotEmpty(batchDenom)

	// query balances
	queryBalanceRes, err := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
		Account:    addr1,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal(t0, queryBalanceRes.TradableUnits)
	s.Require().Equal(r0, queryBalanceRes.RetiredUnits)

	queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
		Account:    addr2,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal(t1, queryBalanceRes.TradableUnits)
	s.Require().Equal(r1, queryBalanceRes.RetiredUnits)

	// query supply
	querySupplyRes, err := s.queryClient.Supply(s.ctx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
	s.Require().NoError(err)
	s.Require().NotNil(querySupplyRes)
	s.Require().Equal(tSupply0, querySupplyRes.TradableSupply)
	s.Require().Equal(rSupply0, querySupplyRes.RetiredSupply)

	// retire credits
	retireCases := []struct {
		name               string
		toRetire           string
		expectErr          bool
		expTradeable       string
		expRetired         string
		expTradeableSupply string
		expRetiredSupply   string
	}{
		{
			name:      "cannot retire more credits than are tradeable",
			toRetire:  "10.371",
			expectErr: true,
		},
		{
			name:      "can't use more than 7 decimal places",
			toRetire:  "10.00000001",
			expectErr: true,
		},
		{
			name:               "can retire a small amount of credits",
			toRetire:           "0.0001",
			expectErr:          false,
			expTradeable:       "10.3699",
			expRetired:         "4.2861",
			expTradeableSupply: "1017.7568",
			expRetiredSupply:   "10004.7450902",
		},
		{
			name:               "can retire more credits",
			toRetire:           "10",
			expectErr:          false,
			expTradeable:       "0.3699",
			expRetired:         "14.2861",
			expTradeableSupply: "1007.7568",
			expRetiredSupply:   "10014.7450902",
		},
		{
			name:               "can retire all credits",
			toRetire:           "0.3699",
			expectErr:          false,
			expTradeable:       "0",
			expRetired:         "14.656",
			expTradeableSupply: "1007.3869",
			expRetiredSupply:   "10015.1149902",
		},
		{
			name:      "can't retire any more credits",
			toRetire:  "1",
			expectErr: true,
		},
	}

	for _, tc := range retireCases {
		tc := tc
		s.Run(tc.name, func() {
			_, err := s.msgClient.Retire(s.ctx, &ecocredit.MsgRetireRequest{
				Holder: addr1,
				Credits: []*ecocredit.MsgRetireRequest_RetireUnits{
					{
						BatchDenom: batchDenom,
						Units:      tc.toRetire,
					},
				},
			})

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				// query balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
					Account:    addr1,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradeable, queryBalanceRes.TradableUnits)
				s.Require().Equal(tc.expRetired, queryBalanceRes.RetiredUnits)

				// query supply
				querySupplyRes, err = s.queryClient.Supply(s.ctx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.Require().Equal(tc.expTradeableSupply, querySupplyRes.TradableSupply)
				s.Require().Equal(tc.expRetiredSupply, querySupplyRes.RetiredSupply)
			}
		})
	}

	sendCases := []struct {
		name                  string
		sendTradeable         string
		sendRetired           string
		expectErr             bool
		expTradeableSender    string
		expRetiredSender      string
		expTradeableRecipient string
		expRetiredRecipient   string
		expTradeableSupply    string
		expRetiredSupply      string
	}{
		{
			name:          "can't send more tradeable than is tradeable",
			sendTradeable: "2000",
			sendRetired:   "10",
			expectErr:     true,
		},
		{
			name:          "can't send more retired than is tradeable",
			sendTradeable: "10",
			sendRetired:   "2000",
			expectErr:     true,
		},
		{
			name:                  "can send some",
			sendTradeable:         "10",
			sendRetired:           "20",
			expectErr:             false,
			expTradeableSender:    "977.3869",
			expRetiredSender:      "10000.4589902",
			expTradeableRecipient: "10",
			expRetiredRecipient:   "20",
			expTradeableSupply:    "987.3869",
			expRetiredSupply:      "10035.1149902",
		},
		{
			name:                  "can send all tradeable",
			sendTradeable:         "77.3869",
			sendRetired:           "900",
			expectErr:             false,
			expTradeableSender:    "0",
			expRetiredSender:      "10000.4589902",
			expTradeableRecipient: "87.3869",
			expRetiredRecipient:   "920",
			expTradeableSupply:    "87.3869",
			expRetiredSupply:      "10935.1149902",
		},
		{
			name:          "can't send any more",
			sendTradeable: "1",
			sendRetired:   "1",
			expectErr:     true,
		},
	}

	for _, tc := range sendCases {
		tc := tc
		s.Run(tc.name, func() {
			_, err := s.msgClient.Send(s.ctx, &ecocredit.MsgSendRequest{
				Sender:    addr2,
				Recipient: addr3,
				Credits: []*ecocredit.MsgSendRequest_SendUnits{
					{
						BatchDenom:    batchDenom,
						TradableUnits: tc.sendTradeable,
						RetiredUnits:  tc.sendRetired,
					},
				},
			})

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				// query sender balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
					Account:    addr2,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradeableSender, queryBalanceRes.TradableUnits)
				s.Require().Equal(tc.expRetiredSender, queryBalanceRes.RetiredUnits)

				// query recipient balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
					Account:    addr3,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradeableRecipient, queryBalanceRes.TradableUnits)
				s.Require().Equal(tc.expRetiredRecipient, queryBalanceRes.RetiredUnits)

				// query supply
				querySupplyRes, err = s.queryClient.Supply(s.ctx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.Require().Equal(tc.expTradeableSupply, querySupplyRes.TradableSupply)
				s.Require().Equal(tc.expRetiredSupply, querySupplyRes.RetiredSupply)
			}
		})
	}

	/****   TEST SET PRECISION   ****/
	precisionCases := []struct {
		name string
		msg  ecocredit.MsgSetPrecisionRequest
		ok   bool
	}{
		{
			"can NOT decrease the decimals", ecocredit.MsgSetPrecisionRequest{
				Issuer: issuer1, BatchDenom: batchDenom, MaxDecimalPlaces: 2},
			false,
		}, {
			"can NOT set to the same value", ecocredit.MsgSetPrecisionRequest{
				Issuer: issuer1, BatchDenom: batchDenom, MaxDecimalPlaces: 7},
			false,
		}, {
			"can increase", ecocredit.MsgSetPrecisionRequest{
				Issuer: issuer1, BatchDenom: batchDenom, MaxDecimalPlaces: 8},
			true,
		}, {
			"can NOT change precision of not existing denom", ecocredit.MsgSetPrecisionRequest{
				Issuer: issuer1, BatchDenom: "not/existing", MaxDecimalPlaces: 1},
			false,
		},
	}
	require := s.Require()
	for _, tc := range precisionCases {
		tc := tc
		s.Run(tc.name, func() {
			_, err := s.msgClient.SetPrecision(s.ctx, &tc.msg)

			if !tc.ok {
				require.Error(err)
			} else {
				require.NoError(err)
				res, err := s.queryClient.Precision(s.ctx,
					&ecocredit.QueryPrecisionRequest{
						BatchDenom: tc.msg.BatchDenom})
				require.NoError(err)
				require.Equal(tc.msg.MaxDecimalPlaces, res.MaxDecimalPlaces)
			}
		})
	}
}
