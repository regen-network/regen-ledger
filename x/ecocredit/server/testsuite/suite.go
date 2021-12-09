package testsuite

import (
	"context"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/testutil"
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

	paramSpace    paramstypes.Subspace
	bankKeeper    bankkeeper.Keeper
	accountKeeper authkeeper.AccountKeeper

	genesisCtx types.Context
	blockTime  time.Time
}

func NewIntegrationTestSuite(fixtureFactory testutil.FixtureFactory, paramSpace paramstypes.Subspace, bankKeeper bankkeeper.BaseKeeper, accountKeeper authkeeper.AccountKeeper) *IntegrationTestSuite {
	return &IntegrationTestSuite{
		fixtureFactory: fixtureFactory,
		paramSpace:     paramSpace,
		bankKeeper:     bankKeeper,
		accountKeeper:  accountKeeper,
	}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.fixture = s.fixtureFactory.Setup()

	s.blockTime = time.Now().UTC()

	// TODO clean up once types.Context merged upstream into sdk.Context
	sdkCtx := s.fixture.Context().(types.Context).WithBlockTime(s.blockTime)
	s.sdkCtx, _ = sdkCtx.CacheContext()
	s.ctx = types.Context{Context: s.sdkCtx}
	s.genesisCtx = types.Context{Context: sdkCtx}

	ecocreditParams := ecocredit.DefaultParams()
	// Add biodiversity credit type for testing credit type sequence numbers
	ecocreditParams.CreditTypes = append(
		ecocreditParams.CreditTypes,
		&ecocredit.CreditType{
			Name:         "biodiversity",
			Abbreviation: "BIO",
			Unit:         "hectare",
			Precision:    6,
		},
	)
	s.paramSpace.SetParamSet(s.sdkCtx, &ecocreditParams)

	s.signers = s.fixture.Signers()
	s.Require().GreaterOrEqual(len(s.signers), 8)
	s.msgClient = ecocredit.NewMsgClient(s.fixture.TxConn())
	s.queryClient = ecocredit.NewQueryClient(s.fixture.QueryConn())
	s.paramsQueryClient = params.NewQueryClient(s.fixture.QueryConn())
}

func (s *IntegrationTestSuite) fundAccount(addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := s.bankKeeper.MintCoins(s.sdkCtx, minttypes.ModuleName, amounts); err != nil {
		return err
	}
	return s.bankKeeper.SendCoinsFromModuleToAccount(s.sdkCtx, minttypes.ModuleName, addr, amounts)
}

func (s *IntegrationTestSuite) TestUpdateClassAdmin() {
	admin := s.signers[0]
	issuer1 := s.signers[1].String()
	issuer2 := s.signers[2].String()
	newAdmin := s.signers[3].String()

	s.Require().NoError(s.fundAccount(admin, sdk.NewCoins(sdk.NewInt64Coin("stake", 4*ecocredit.DefaultCreditClassFeeTokens.Int64()))))
	createClsRes, err := s.msgClient.CreateClass(s.ctx, &ecocredit.MsgCreateClass{Admin: admin.String(), Issuers: []string{issuer1, issuer2}, Metadata: nil, CreditTypeName: "carbon"})
	s.Require().NoError(err)
	s.Require().NotNil(createClsRes)
	classID := createClsRes.ClassId

	testCases := []struct {
		name   string
		msg    ecocredit.MsgUpdateClassAdmin
		expErr bool
	}{
		{
			name:   "invalid: not admin",
			msg:    ecocredit.MsgUpdateClassAdmin{ClassId: classID, Admin: issuer1, NewAdmin: newAdmin},
			expErr: true,
		},
		{
			name:   "invalid: bad classID",
			msg:    ecocredit.MsgUpdateClassAdmin{ClassId: "foobarbaz", Admin: admin.String(), NewAdmin: newAdmin},
			expErr: true,
		},
		{
			name:   "valid",
			msg:    ecocredit.MsgUpdateClassAdmin{ClassId: classID, Admin: admin.String(), NewAdmin: newAdmin},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			updateRes, err := s.msgClient.UpdateClassAdmin(s.ctx, &tc.msg)
			if tc.expErr {
				s.Require().Error(err)
				return
			}

			s.Require().NoError(err)
			s.Require().NotNil(updateRes)

			res, err := s.queryClient.ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classID})
			s.Require().NoError(err)
			s.Require().NotNil(res)

			s.Require().Equal(res.Info.Admin, newAdmin)
		})
	}
}

func (s *IntegrationTestSuite) TestUpdateClassIssuers() {
	admin := s.signers[0]
	issuer1 := s.signers[1].String()
	issuer2 := s.signers[2].String()
	issuer3 := s.signers[3].String()

	s.Require().NoError(s.fundAccount(admin, sdk.NewCoins(sdk.NewInt64Coin("stake", 4*ecocredit.DefaultCreditClassFeeTokens.Int64()))))
	createClsRes, err := s.msgClient.CreateClass(s.ctx, &ecocredit.MsgCreateClass{Admin: admin.String(), Issuers: []string{issuer1}, Metadata: nil, CreditTypeName: "carbon"})
	s.Require().NoError(err)
	s.Require().NotNil(createClsRes)
	classID := createClsRes.ClassId

	testCases := []struct {
		name   string
		msg    ecocredit.MsgUpdateClassIssuers
		expErr bool
	}{
		{
			name:   "invalid: not admin",
			msg:    ecocredit.MsgUpdateClassIssuers{ClassId: classID, Admin: issuer1, Issuers: []string{issuer1}},
			expErr: true,
		},
		{
			name:   "invalid: bad classID",
			msg:    ecocredit.MsgUpdateClassIssuers{ClassId: "foobarbaz", Admin: admin.String(), Issuers: []string{}},
			expErr: true,
		},
		{
			name:   "valid",
			msg:    ecocredit.MsgUpdateClassIssuers{ClassId: classID, Admin: admin.String(), Issuers: []string{issuer2, issuer3}},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			updateRes, err := s.msgClient.UpdateClassIssuers(s.ctx, &tc.msg)
			if tc.expErr {
				s.Require().Error(err)
				return
			}

			s.Require().NoError(err)
			s.Require().NotNil(updateRes)

			res, err := s.queryClient.ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classID})
			s.Require().NoError(err)
			s.Require().NotNil(res)

			s.Require().Equal(res.Info.Issuers, tc.msg.Issuers)

		})
	}
}

func (s *IntegrationTestSuite) TestUpdateClassMetadata() {
	admin := s.signers[0]
	issuer1 := s.signers[3].String()

	s.Require().NoError(s.fundAccount(admin, sdk.NewCoins(sdk.NewInt64Coin("stake", 4*ecocredit.DefaultCreditClassFeeTokens.Int64()))))
	createClsRes, err := s.msgClient.CreateClass(s.ctx, &ecocredit.MsgCreateClass{Admin: admin.String(), Issuers: []string{issuer1}, Metadata: nil, CreditTypeName: "carbon"})
	s.Require().NoError(err)
	s.Require().NotNil(createClsRes)
	classID := createClsRes.ClassId

	testCases := []struct {
		name   string
		msg    ecocredit.MsgUpdateClassMetadata
		expErr bool
	}{
		{
			name:   "invalid: not admin",
			msg:    ecocredit.MsgUpdateClassMetadata{ClassId: classID, Admin: issuer1, Metadata: []byte("hello")},
			expErr: true,
		},
		{
			name:   "invalid: bad classID",
			msg:    ecocredit.MsgUpdateClassMetadata{ClassId: "foobarbaz", Admin: admin.String()},
			expErr: true,
		},
		{
			name:   "valid",
			msg:    ecocredit.MsgUpdateClassMetadata{ClassId: classID, Admin: admin.String(), Metadata: []byte("hello world")},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			updateRes, err := s.msgClient.UpdateClassMetadata(s.ctx, &tc.msg)
			if tc.expErr {
				s.Require().Error(err)
				return
			}

			s.Require().NoError(err)
			s.Require().NotNil(updateRes)

			res, err := s.queryClient.ClassInfo(s.ctx, &ecocredit.QueryClassInfoRequest{ClassId: classID})
			s.Require().NoError(err)
			s.Require().NotNil(res)

			s.Require().Equal(res.Info.Metadata, tc.msg.Metadata)

		})
	}
}

func (s *IntegrationTestSuite) TestScenario() {
	admin := s.signers[0]
	issuer1 := s.signers[1].String()
	issuer2 := s.signers[2].String()
	addr1 := s.signers[3].String()
	addr2 := s.signers[4].String()
	addr3 := s.signers[5].String()
	addr4 := s.signers[6].String()
	addr5 := s.signers[7].String()

	// create class with insufficient funds and it should fail
	createClsRes, err := s.msgClient.CreateClass(s.ctx, &ecocredit.MsgCreateClass{
		Admin:          admin.String(),
		Issuers:        []string{issuer1, issuer2},
		Metadata:       nil,
		CreditTypeName: "carbon",
	})
	s.Require().Error(err)
	s.Require().Nil(createClsRes)

	// create class with sufficient funds and it should succeed
	s.Require().NoError(s.fundAccount(admin, sdk.NewCoins(sdk.NewInt64Coin("stake", 4*ecocredit.DefaultCreditClassFeeTokens.Int64()))))

	// Run multiple tests to test the CreditTypeSeqs
	createClassTestCases := []struct {
		creditType      string
		expectedClassID string
	}{
		{
			creditType:      "carbon",
			expectedClassID: "C01",
		},
		{
			creditType:      "biodiversity",
			expectedClassID: "BIO04",
		},
		{
			creditType:      "biodiversity",
			expectedClassID: "BIO05",
		},
		{
			creditType:      "carbon",
			expectedClassID: "C02",
		},
	}

	for _, tc := range createClassTestCases {
		createClsRes, err = s.msgClient.CreateClass(s.ctx, &ecocredit.MsgCreateClass{
			Admin:          admin.String(),
			Issuers:        []string{issuer1, issuer2},
			Metadata:       nil,
			CreditTypeName: tc.creditType,
		})
		s.Require().NoError(err)
		s.Require().NotNil(createClsRes)

		s.Require().Equal(tc.expectedClassID, createClsRes.ClassId)
	}

	// Use first test class for remainder of tests
	clsID := createClassTestCases[0].expectedClassID

	// admin should have no funds remaining
	s.Require().Equal(s.bankKeeper.GetBalance(s.sdkCtx, admin, "stake"), sdk.NewInt64Coin("stake", 0))

	// create batch
	t0, t1, t2 := "10.37", "1007.3869", "100"
	tSupply0 := "1117.7569"
	r0, r1, r2 := "4.286", "10000.45899", "0"
	rSupply0 := "10004.74499"

	time1 := time.Now()
	time2 := time.Now()

	// Batch creation should succeed with StartDate before EndDate, and valid data
	createBatchRes, err := s.msgClient.CreateBatch(s.ctx, &ecocredit.MsgCreateBatch{
		Issuer:          issuer1,
		ClassId:         clsID,
		StartDate:       &time1,
		EndDate:         &time2,
		ProjectLocation: "AB",
		Issuance: []*ecocredit.MsgCreateBatch_BatchIssuance{
			{
				Recipient:          addr1,
				TradableAmount:     t0,
				RetiredAmount:      r0,
				RetirementLocation: "GB",
			},
			{
				Recipient:          addr2,
				TradableAmount:     t1,
				RetiredAmount:      r1,
				RetirementLocation: "BF",
			},
			{
				Recipient:          addr4,
				TradableAmount:     t2,
				RetiredAmount:      r2,
				RetirementLocation: "",
			},
			{
				Recipient:          addr5,
				RetirementLocation: "",
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
	s.Require().Equal(t0, queryBalanceRes.TradableAmount)
	s.Require().Equal(r0, queryBalanceRes.RetiredAmount)

	queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
		Account:    addr2,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal(t1, queryBalanceRes.TradableAmount)
	s.Require().Equal(r1, queryBalanceRes.RetiredAmount)

	queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
		Account:    addr4,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal(t2, queryBalanceRes.TradableAmount)
	s.Require().Equal(r2, queryBalanceRes.RetiredAmount)

	// if we didn't issue tradable or retired balances, they'll be default to zero.
	queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
		Account:    addr5,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal("0", queryBalanceRes.TradableAmount)
	s.Require().Equal("0", queryBalanceRes.RetiredAmount)

	// query supply
	querySupplyRes, err := s.queryClient.Supply(s.ctx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
	s.Require().NoError(err)
	s.Require().NotNil(querySupplyRes)
	s.Require().Equal(tSupply0, querySupplyRes.TradableSupply)
	s.Require().Equal(rSupply0, querySupplyRes.RetiredSupply)

	// cancel credits
	cancelCases := []struct {
		name               string
		holder             string
		toCancel           string
		expectErr          bool
		expTradable        string
		expTradableSupply  string
		expRetired         string
		expTotalAmount     string
		expAmountCancelled string
		expErrMessage      string
	}{
		{
			name:          "can't cancel more credits than are tradable",
			holder:        addr4,
			toCancel:      "101",
			expectErr:     true,
			expErrMessage: "insufficient credit balance",
		},
		{
			name:          "can't cancel with a higher precision than the credit type",
			holder:        addr4,
			toCancel:      "0.1234567",
			expectErr:     true,
			expErrMessage: "exceeds maximum decimal places",
		},
		{
			name:          "can't cancel no credits",
			holder:        addr4,
			toCancel:      "0",
			expectErr:     true,
			expErrMessage: "expected a positive decimal",
		},
		{
			name:               "can cancel a small amount of credits",
			holder:             addr4,
			toCancel:           "2.0002",
			expectErr:          false,
			expTradable:        "97.9998",
			expTradableSupply:  "1115.7567",
			expRetired:         "0",
			expTotalAmount:     "11120.50169",
			expAmountCancelled: "2.0002",
		},
		{
			name:               "can cancel all remaining credits",
			holder:             addr4,
			toCancel:           "97.9998",
			expectErr:          false,
			expTradable:        "0",
			expTradableSupply:  "1017.7569",
			expRetired:         "0",
			expTotalAmount:     "11022.50189",
			expAmountCancelled: "100.0000",
		},
		{
			name:          "can't cancel anymore credits",
			holder:        addr4,
			toCancel:      "1",
			expectErr:     true,
			expErrMessage: "insufficient credit balance",
		},
		{
			name:               "can cancel from account with positive retired balance",
			holder:             addr1,
			toCancel:           "1",
			expectErr:          false,
			expTradable:        "9.37",
			expTradableSupply:  "1016.7569",
			expRetired:         "4.286",
			expTotalAmount:     "11021.50189",
			expAmountCancelled: "101.0000",
		},
	}

	for _, tc := range cancelCases {
		s.Run(tc.name, func() {
			_, err := s.msgClient.Cancel(s.ctx, &ecocredit.MsgCancel{
				Holder: tc.holder,
				Credits: []*ecocredit.MsgCancel_CancelCredits{
					{
						BatchDenom: batchDenom,
						Amount:     tc.toCancel,
					},
				},
			})

			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMessage)
			} else {
				s.Require().NoError(err)

				// query balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
					Account:    tc.holder,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradable, queryBalanceRes.TradableAmount)
				s.Require().Equal(tc.expRetired, queryBalanceRes.RetiredAmount)

				// query supply
				querySupplyRes, err = s.queryClient.Supply(s.ctx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.Require().Equal(tc.expTradableSupply, querySupplyRes.TradableSupply)
				s.Require().Equal(rSupply0, querySupplyRes.RetiredSupply)

				// query batchInfo
				queryBatchInfoRes, err := s.queryClient.BatchInfo(s.ctx, &ecocredit.QueryBatchInfoRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(queryBatchInfoRes)
				s.Require().Equal(tc.expTotalAmount, queryBatchInfoRes.Info.TotalAmount)
				s.Require().Equal(tc.expAmountCancelled, queryBatchInfoRes.Info.AmountCancelled)
			}
		})
	}

	// retire credits
	retireCases := []struct {
		name               string
		toRetire           string
		retirementLocation string
		expectErr          bool
		expTradable        string
		expRetired         string
		expTradableSupply  string
		expRetiredSupply   string
		expErrMessage      string
	}{
		{
			name:               "cannot retire more credits than are tradable",
			toRetire:           "10.371",
			retirementLocation: "AF",
			expectErr:          true,
			expErrMessage:      "insufficient credit balance",
		},
		{
			name:               "can't use more precision than the credit type allows (6)",
			toRetire:           "10.00000001",
			retirementLocation: "AF",
			expectErr:          true,
			expErrMessage:      "exceeds maximum decimal places",
		},
		{
			name:               "can't retire to an invalid country",
			toRetire:           "0.0001",
			retirementLocation: "ZZZ",
			expectErr:          true,
			expErrMessage:      "Invalid location",
		},
		{
			name:               "can't retire to an invalid region",
			toRetire:           "0.0001",
			retirementLocation: "AF-ZZZZ",
			expectErr:          true,
			expErrMessage:      "Invalid location",
		},
		{
			name:               "can't retire to an invalid postal code",
			toRetire:           "0.0001",
			retirementLocation: "AF-BDS 0123456789012345678901234567890123456789012345678901234567890123456789",
			expectErr:          true,
			expErrMessage:      "Invalid location",
		},
		{
			name:               "can't retire without a location",
			toRetire:           "0.0001",
			retirementLocation: "",
			expectErr:          true,
			expErrMessage:      "Invalid location",
		},
		{
			name:               "can retire a small amount of credits",
			toRetire:           "0.0001",
			retirementLocation: "AF",
			expectErr:          false,
			expTradable:        "9.3699",
			expRetired:         "4.2861",
			expTradableSupply:  "1016.7568",
			expRetiredSupply:   "10004.74509",
		},
		{
			name:               "can retire more credits",
			toRetire:           "9",
			retirementLocation: "AF-BDS",
			expectErr:          false,
			expTradable:        "0.3699",
			expRetired:         "13.2861",
			expTradableSupply:  "1007.7568",
			expRetiredSupply:   "10013.74509",
		},
		{
			name:               "can retire all credits",
			toRetire:           "0.3699",
			retirementLocation: "AF-BDS 12345",
			expectErr:          false,
			expTradable:        "0",
			expRetired:         "13.656",
			expTradableSupply:  "1007.3869",
			expRetiredSupply:   "10014.11499",
		},
		{
			name:               "can't retire any more credits",
			toRetire:           "1",
			retirementLocation: "AF-BDS",
			expectErr:          true,
			expErrMessage:      "insufficient credit balance",
		},
	}

	for _, tc := range retireCases {
		tc := tc
		s.Run(tc.name, func() {
			_, err := s.msgClient.Retire(s.ctx, &ecocredit.MsgRetire{
				Holder: addr1,
				Credits: []*ecocredit.MsgRetire_RetireCredits{
					{
						BatchDenom: batchDenom,
						Amount:     tc.toRetire,
					},
				},
				Location: tc.retirementLocation,
			})

			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMessage)
			} else {
				s.Require().NoError(err)

				// query balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
					Account:    addr1,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradable, queryBalanceRes.TradableAmount)
				s.Require().Equal(tc.expRetired, queryBalanceRes.RetiredAmount)

				// query supply
				querySupplyRes, err = s.queryClient.Supply(s.ctx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.Require().Equal(tc.expTradableSupply, querySupplyRes.TradableSupply)
				s.Require().Equal(tc.expRetiredSupply, querySupplyRes.RetiredSupply)
			}
		})
	}

	sendCases := []struct {
		name                 string
		sendTradable         string
		sendRetired          string
		retirementLocation   string
		expectErr            bool
		expTradableSender    string
		expRetiredSender     string
		expTradableRecipient string
		expRetiredRecipient  string
		expTradableSupply    string
		expRetiredSupply     string
		expErrMessage        string
	}{
		{
			name:               "can't send an amount with more decimal places than allowed precision (6)",
			sendTradable:       "2.123456789",
			sendRetired:        "10.123456789",
			retirementLocation: "AF",
			expectErr:          true,
			expErrMessage:      "exceeds maximum decimal places",
		},
		{
			name:               "can't send more tradable than is tradable",
			sendTradable:       "2000",
			sendRetired:        "10",
			retirementLocation: "AF",
			expectErr:          true,
			expErrMessage:      "insufficient credit balance",
		},
		{
			name:               "can't send more retired than is tradable",
			sendTradable:       "10",
			sendRetired:        "2000",
			retirementLocation: "AF",
			expectErr:          true,
			expErrMessage:      "insufficient credit balance",
		},
		{
			name:               "can't send to an invalid country",
			sendTradable:       "10",
			sendRetired:        "20",
			retirementLocation: "ZZZ",
			expectErr:          true,
			expErrMessage:      "Invalid location",
		},
		{
			name:               "can't send to an invalid region",
			sendTradable:       "10",
			sendRetired:        "20",
			retirementLocation: "AF-ZZZZ",
			expectErr:          true,
			expErrMessage:      "Invalid location",
		},
		{
			name:               "can't send to an invalid postal code",
			sendTradable:       "10",
			sendRetired:        "20",
			retirementLocation: "AF-BDS 0123456789012345678901234567890123456789012345678901234567890123456789",
			expectErr:          true,
			expErrMessage:      "Invalid location",
		},
		{
			name:                 "can send some",
			sendTradable:         "10",
			sendRetired:          "20",
			retirementLocation:   "AF",
			expectErr:            false,
			expTradableSender:    "977.3869",
			expRetiredSender:     "10000.45899",
			expTradableRecipient: "10",
			expRetiredRecipient:  "20",
			expTradableSupply:    "987.3869",
			expRetiredSupply:     "10034.11499",
		},
		{
			name:                 "can send with no retirement location",
			sendTradable:         "10",
			sendRetired:          "0",
			retirementLocation:   "",
			expectErr:            false,
			expTradableSender:    "967.3869",
			expRetiredSender:     "10000.45899",
			expTradableRecipient: "20",
			expRetiredRecipient:  "20",
			expTradableSupply:    "987.3869",
			expRetiredSupply:     "10034.11499",
		},
		{
			name:                 "can send all tradable",
			sendTradable:         "67.3869",
			sendRetired:          "900",
			retirementLocation:   "AF",
			expectErr:            false,
			expTradableSender:    "0",
			expRetiredSender:     "10000.45899",
			expTradableRecipient: "87.3869",
			expRetiredRecipient:  "920",
			expTradableSupply:    "87.3869",
			expRetiredSupply:     "10934.11499",
		},
		{
			name:               "can't send any more",
			sendTradable:       "1",
			sendRetired:        "1",
			expectErr:          true,
			retirementLocation: "AF",
			expErrMessage:      "insufficient credit balance",
		},
	}

	for _, tc := range sendCases {
		tc := tc
		s.Run(tc.name, func() {
			_, err := s.msgClient.Send(s.ctx, &ecocredit.MsgSend{
				Sender:    addr2,
				Recipient: addr3,
				Credits: []*ecocredit.MsgSend_SendCredits{
					{
						BatchDenom:         batchDenom,
						TradableAmount:     tc.sendTradable,
						RetiredAmount:      tc.sendRetired,
						RetirementLocation: tc.retirementLocation,
					},
				},
			})

			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMessage)
			} else {
				s.Require().NoError(err)

				// query sender balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
					Account:    addr2,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradableSender, queryBalanceRes.TradableAmount)
				s.Require().Equal(tc.expRetiredSender, queryBalanceRes.RetiredAmount)

				// query recipient balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
					Account:    addr3,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradableRecipient, queryBalanceRes.TradableAmount)
				s.Require().Equal(tc.expRetiredRecipient, queryBalanceRes.RetiredAmount)

				// query supply
				querySupplyRes, err = s.queryClient.Supply(s.ctx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.Require().Equal(tc.expTradableSupply, querySupplyRes.TradableSupply)
				s.Require().Equal(tc.expRetiredSupply, querySupplyRes.RetiredSupply)
			}
		})
	}

	/****   TEST ALLOWLIST CREDIT CREATORS   ****/
	allowlistCases := []struct {
		name             string
		creatorAcc       sdk.AccAddress
		allowlist        []string
		allowlistEnabled bool
		wantErr          bool
	}{
		{
			name:             "valid allowlist and enabled",
			allowlist:        []string{s.signers[0].String()},
			creatorAcc:       s.signers[0],
			allowlistEnabled: true,
			wantErr:          false,
		},
		{
			name:             "valid multi addrs in allowlist",
			allowlist:        []string{s.signers[0].String(), s.signers[1].String(), s.signers[2].String()},
			creatorAcc:       s.signers[0],
			allowlistEnabled: true,
			wantErr:          false,
		},
		{
			name:             "creator is not part of the allowlist",
			allowlist:        []string{s.signers[0].String()},
			creatorAcc:       s.signers[1],
			allowlistEnabled: true,
			wantErr:          true,
		},
		{
			name:             "valid allowlist but disabled - anyone can create credits",
			allowlist:        []string{s.signers[0].String()},
			creatorAcc:       s.signers[0],
			allowlistEnabled: false,
			wantErr:          false,
		},
		{
			name:             "empty and enabled allowlist - nobody can create credits",
			allowlist:        []string{},
			creatorAcc:       s.signers[0],
			allowlistEnabled: true,
			wantErr:          true,
		},
	}

	for _, tc := range allowlistCases {
		tc := tc
		s.Run(tc.name, func() {
			s.paramSpace.Set(s.sdkCtx, ecocredit.KeyAllowedClassCreators, tc.allowlist)
			s.paramSpace.Set(s.sdkCtx, ecocredit.KeyAllowlistEnabled, tc.allowlistEnabled)

			// fund the creator account
			s.Require().NoError(s.fundAccount(tc.creatorAcc, sdk.NewCoins(sdk.NewCoin("stake", ecocredit.DefaultCreditClassFeeTokens))))

			createClsRes, err = s.msgClient.CreateClass(s.ctx, &ecocredit.MsgCreateClass{
				Admin:          tc.creatorAcc.String(),
				Issuers:        []string{issuer1, issuer2},
				CreditTypeName: "carbon",
				Metadata:       nil,
			})
			if tc.wantErr {
				s.Require().Error(err)
				s.Require().Nil(createClsRes)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(createClsRes)
			}
		})
	}

	// Disable credit class allowlist for credit type tests
	s.paramSpace.Set(s.sdkCtx, ecocredit.KeyAllowlistEnabled, false)

	/****   TEST CREDIT TYPES   ****/
	creditTypeCases := []struct {
		name        string
		creditTypes []*ecocredit.CreditType
		msg         ecocredit.MsgCreateClass
		wantErr     bool
	}{
		{
			name: "valid eco credit creation",
			creditTypes: []*ecocredit.CreditType{
				{Name: "carbon", Abbreviation: "C", Unit: "metric ton CO2 equivalent", Precision: 3},
			},
			msg: ecocredit.MsgCreateClass{
				Admin:          s.signers[0].String(),
				Issuers:        []string{s.signers[1].String(), s.signers[2].String()},
				Metadata:       nil,
				CreditTypeName: "carbon",
			},
			wantErr: false,
		},
		{
			name: "invalid request - not a valid credit type",
			creditTypes: []*ecocredit.CreditType{
				{Name: "carbon", Abbreviation: "C", Unit: "metric ton CO2 equivalent", Precision: 3},
			},
			msg: ecocredit.MsgCreateClass{
				Admin:          s.signers[0].String(),
				Issuers:        []string{s.signers[1].String(), s.signers[2].String()},
				Metadata:       nil,
				CreditTypeName: "biodiversity",
			},
			wantErr: true,
		},
		{
			name: "request with strange font should be valid",
			creditTypes: []*ecocredit.CreditType{
				{Name: "carbon", Abbreviation: "C", Unit: "metric ton CO2 equivalent", Precision: 3},
			},
			msg: ecocredit.MsgCreateClass{
				Admin:          s.signers[0].String(),
				Issuers:        []string{s.signers[1].String(), s.signers[2].String()},
				Metadata:       nil,
				CreditTypeName: "cArBoN",
			},
			wantErr: false,
		},
		{
			name:        "empty credit types should error",
			creditTypes: []*ecocredit.CreditType{},
			msg: ecocredit.MsgCreateClass{
				Admin:          s.signers[0].String(),
				Issuers:        []string{s.signers[1].String(), s.signers[2].String()},
				Metadata:       nil,
				CreditTypeName: "carbon",
			},
			wantErr: true,
		},
	}

	for _, tc := range creditTypeCases {
		tc := tc

		s.Run(tc.name, func() {
			require := s.Require()
			s.paramSpace.Set(s.sdkCtx, ecocredit.KeyCreditTypes, tc.creditTypes)
			admin, err := sdk.AccAddressFromBech32(tc.msg.Admin)
			require.NoError(err)

			// fund the admin account so tx will go through
			s.Require().NoError(s.fundAccount(admin, sdk.NewCoins(sdk.NewCoin("stake", ecocredit.DefaultCreditClassFeeTokens))))
			res, err := s.msgClient.CreateClass(s.ctx, &tc.msg)
			if tc.wantErr {
				require.Error(err)
				require.Nil(res)
			} else {
				require.NoError(err)
				require.NotNil(res)
			}
		})
	}

	// reset the space to avoid corrupting other tests
	s.paramSpace.Set(s.sdkCtx, ecocredit.KeyCreditTypes, ecocredit.DefaultParams().CreditTypes)

	askPrice := sdk.NewInt64Coin("stake", 1)
	expectedSellOrderIds := []uint64{1, 2}
	createSellOrder, err := s.msgClient.Sell(s.ctx, &ecocredit.MsgSell{
		Owner: addr3,
		Orders: []*ecocredit.MsgSell_Order{
			{
				BatchDenom:        batchDenom,
				Quantity:          "1.0",
				AskPrice:          &askPrice,
				DisableAutoRetire: true,
			},
			{
				BatchDenom:        batchDenom,
				Quantity:          "1.0",
				AskPrice:          &askPrice,
				DisableAutoRetire: true,
			},
		},
	})
	s.Require().Nil(err)
	s.Require().Equal(expectedSellOrderIds, createSellOrder.SellOrderIds)

	expectedBuyOrderIds := []uint64{1}
	selection := &ecocredit.MsgBuy_Order_Selection{
		Sum: &ecocredit.MsgBuy_Order_Selection_SellOrderId{SellOrderId: 2},
	}
	createBuyOrder, err := s.msgClient.Buy(s.ctx, &ecocredit.MsgBuy{
		Buyer: admin.String(),
		Orders: []*ecocredit.MsgBuy_Order{
			{
				Selection:          selection,
				Quantity:           "1.0",
				BidPrice:           &askPrice,
				DisableAutoRetire:  true,
				DisablePartialFill: true,
			},
		},
	})
	s.Require().Nil(err)
	s.Require().Equal(expectedBuyOrderIds, createBuyOrder.BuyOrderIds)
}
