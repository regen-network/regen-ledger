package testsuite

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory testutil.FixtureFactory
	fixture        testutil.Fixture

	codec             *codec.ProtoCodec
	sdkCtx            sdk.Context
	ctx               context.Context
	msgClient         core.MsgClient
	marketServer      marketServer
	basketServer      basketServer
	queryClient       core.QueryClient
	paramsQueryClient params.QueryClient
	signers           []sdk.AccAddress
	basketFee         sdk.Coin

	paramSpace    paramstypes.Subspace
	bankKeeper    bankkeeper.Keeper
	accountKeeper authkeeper.AccountKeeper

	genesisCtx types.Context
	blockTime  time.Time
}

type marketServer struct {
	marketplace.QueryClient
	marketplace.MsgClient
}

type basketServer struct {
	basket.QueryClient
	basket.MsgClient
}

var (
	createClassFee = sdk.Coin{Denom: sdk.DefaultBondDenom, Amount: core.DefaultCreditClassFee}
)

func NewIntegrationTestSuite(fixtureFactory testutil.FixtureFactory, paramSpace paramstypes.Subspace, bankKeeper bankkeeper.BaseKeeper, accountKeeper authkeeper.AccountKeeper) *IntegrationTestSuite {
	return &IntegrationTestSuite{
		fixtureFactory: fixtureFactory,
		paramSpace:     paramSpace,
		bankKeeper:     bankKeeper,
		accountKeeper:  accountKeeper,
	}
}

func (s *IntegrationTestSuite) SetupSuite() {
	// TODO: remove after updating to cosmos-sdk v0.46 #857
	sdk.SetCoinDenomRegex(func() string {
		return types.CoinDenomRegex
	})

	s.fixture = s.fixtureFactory.Setup()

	s.codec = s.fixture.Codec()

	s.blockTime = time.Now().UTC()

	// TODO clean up once types.Context merged upstream into sdk.Context
	sdkCtx := s.fixture.Context().(types.Context).WithBlockTime(s.blockTime)
	s.sdkCtx, _ = sdkCtx.CacheContext()
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
	s.genesisCtx = types.Context{Context: sdkCtx}

	_, err := s.fixture.InitGenesis(s.sdkCtx, map[string]json.RawMessage{ecocredit.ModuleName: s.ecocreditGenesis()})
	s.Require().NoError(err)

	ecocreditParams := core.DefaultParams()
	s.basketFee = sdk.NewInt64Coin("bfee", 20)
	ecocreditParams.BasketFee = sdk.NewCoins(s.basketFee)
	s.paramSpace.SetParamSet(s.sdkCtx, &ecocreditParams)

	s.signers = s.fixture.Signers()
	s.Require().GreaterOrEqual(len(s.signers), 8)
	s.basketServer = basketServer{basket.NewQueryClient(s.fixture.QueryConn()), basket.NewMsgClient(s.fixture.TxConn())}
	s.marketServer = marketServer{marketplace.NewQueryClient(s.fixture.QueryConn()), marketplace.NewMsgClient(s.fixture.TxConn())}
	s.msgClient = core.NewMsgClient(s.fixture.TxConn())
	s.queryClient = core.NewQueryClient(s.fixture.QueryConn())
	s.paramsQueryClient = params.NewQueryClient(s.fixture.QueryConn())
}

func (s *IntegrationTestSuite) ecocreditGenesis() json.RawMessage {
	// setup temporary mem db
	db := dbm.NewMemDB()
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()
	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      db,
	})
	modDB, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	s.Require().NoError(err)
	ormCtx := ormtable.WrapContextDefault(backend)
	ss, err := api.NewStateStore(modDB)
	s.Require().NoError(err)
	ms, err := marketApi.NewStateStore(modDB)
	s.Require().NoError(err)

	err = ms.AllowedDenomTable().Insert(ormCtx, &marketApi.AllowedDenom{
		BankDenom:    sdk.DefaultBondDenom,
		DisplayDenom: sdk.DefaultBondDenom,
	})
	s.Require().NoError(err)

	err = ss.CreditTypeTable().Insert(ormCtx, &api.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "metric ton C02",
		Precision:    6,
	})
	s.Require().NoError(err)

	// export genesis into target
	target := ormjson.NewRawMessageTarget()
	err = modDB.ExportJSON(ormCtx, target)
	s.Require().NoError(err)

	// merge the params into the json target
	coreParams := core.DefaultParams()
	err = core.MergeParamsIntoTarget(s.codec, &coreParams, target)
	s.Require().NoError(err)

	// get raw json from target
	ecoJsn, err := target.JSON()
	s.Require().NoError(err)

	// set the module genesis
	return ecoJsn
}

func (s *IntegrationTestSuite) TestBasketScenario() {
	require := s.Require()
	user := s.signers[0]
	user2 := s.signers[1]

	// create a class and issue a batch
	userTotalCreditBalance, err := math.NewDecFromString("1000000000000000")
	require.NoError(err)
	classId, batchDenom := s.createClassAndIssueBatch(user, user, "C", userTotalCreditBalance.String(), "2020-01-01", "2022-01-01")

	// fund account to create a basket
	balanceBefore := sdk.NewInt64Coin(s.basketFee.Denom, 30000)
	s.fundAccount(user, sdk.NewCoins(balanceBefore))

	// create a basket
	res, err := s.basketServer.Create(s.ctx, &basket.MsgCreate{
		Curator:           s.signers[0].String(),
		Name:              "BASKET",
		Exponent:          6,
		DisableAutoRetire: true,
		CreditTypeAbbrev:  "C",
		AllowedClasses:    []string{classId},
		DateCriteria:      nil,
		Fee:               sdk.NewCoins(s.basketFee),
	})
	require.NoError(err)
	basketDenom := res.BasketDenom

	// check it was created
	qRes, err := s.basketServer.Baskets(s.ctx, &basket.QueryBasketsRequest{})
	require.NoError(err)
	require.Len(qRes.Baskets, 1)
	require.Equal(qRes.Baskets[0].BasketDenom, basketDenom)

	// assert the fee was paid - the fee mechanism was mocked, but we still call the same underlying SendFromAccountToModule
	// function so the result is the same
	balanceAfter := s.getUserBalance(user, s.basketFee.Denom)
	require.Equal(balanceAfter.Add(s.basketFee), balanceBefore)

	// put some BAZ credits in the basket
	creditAmtDeposited := math.NewDecFromInt64(3)
	pRes, err := s.basketServer.Put(s.ctx, &basket.MsgPut{
		Owner:       user.String(),
		BasketDenom: basketDenom,
		Credits:     []*basket.BasketCredit{{BatchDenom: batchDenom, Amount: creditAmtDeposited.String()}},
	})
	require.NoError(err)
	basketTokensReceived, err := math.NewPositiveDecFromString(pRes.AmountReceived)
	require.NoError(err)

	// make sure the bank actually has this balance for the user
	basketBal := s.getUserBalance(user, basketDenom)
	i64BT, err := basketTokensReceived.Int64()
	require.NoError(err)
	require.Equal(i64BT, basketBal.Amount.Int64())

	// make sure the basket has the credits now.
	basketBalance, err := s.basketServer.BasketBalance(s.ctx, &basket.QueryBasketBalanceRequest{
		BasketDenom: basketDenom,
		BatchDenom:  batchDenom,
	})
	require.NoError(err)
	require.Equal(basketBalance.Balance, creditAmtDeposited.String())

	// make sure user doesn't have any of that credit - should error out
	userCreditBalance, err := s.queryClient.Balance(s.ctx, &core.QueryBalanceRequest{
		Address:    user.String(),
		BatchDenom: batchDenom,
	})
	require.NoError(err)

	// make sure the core server is properly tracking the user balance
	newUserTotal, err := userTotalCreditBalance.Sub(creditAmtDeposited)
	require.NoError(err)
	require.Equal(newUserTotal.String(), userCreditBalance.Balance.TradableAmount)

	// send the basket coins to another account - user2
	require.NoError(s.bankKeeper.SendCoins(s.sdkCtx, user, user2, sdk.NewCoins(sdk.NewInt64Coin(basketDenom, i64BT))))

	// user2 can take all the credits from the basket
	tRes, err := s.basketServer.Take(s.ctx, &basket.MsgTake{
		Owner:                  user2.String(),
		BasketDenom:            basketDenom,
		Amount:                 basketTokensReceived.String(),
		RetirementJurisdiction: "US-NY",
		RetireOnTake:           false,
	})
	require.NoError(err)
	require.Equal(tRes.Credits[0].BatchDenom, batchDenom)
	require.Equal(tRes.Credits[0].Amount, creditAmtDeposited.String())

	// user shouldn't be able to take any since we sent our tokens to user2
	noRes, err := s.basketServer.Take(s.ctx, &basket.MsgTake{
		Owner:                  user.String(),
		BasketDenom:            basketDenom,
		Amount:                 basketTokensReceived.String(),
		RetirementJurisdiction: "US-NY",
		RetireOnTake:           false,
	})
	require.Error(err)
	require.Contains(err.Error(), sdkerrors.ErrInsufficientFunds.Error())
	require.Nil(noRes)

	// there should be nothing left in the basket
	bRes, err := s.basketServer.BasketBalance(s.ctx, &basket.QueryBasketBalanceRequest{
		BasketDenom: basketDenom,
		BatchDenom:  batchDenom,
	})
	require.Error(err)
	require.Contains(err.Error(), "not found")
	require.Nil(bRes)

	// basket token balance of user2 should be empty now
	endBal := s.getUserBalance(user2, basketDenom)
	require.True(endBal.Amount.Equal(sdk.NewInt(0)), "ending balance was %s, expected 0", endBal.Amount.String())

	// create a retire enabled basket
	resR, err := s.basketServer.Create(s.ctx, &basket.MsgCreate{
		Curator:           s.signers[0].String(),
		Name:              "RETIRE",
		Exponent:          6,
		DisableAutoRetire: false,
		CreditTypeAbbrev:  "C",
		AllowedClasses:    []string{classId},
		DateCriteria:      nil,
		Fee:               sdk.NewCoins(s.basketFee),
	})
	require.NoError(err)
	basketDenom = resR.BasketDenom

	creditsToDeposit := math.NewDecFromInt64(3)

	// put some credits in the basket
	pRes, err = s.basketServer.Put(s.ctx, &basket.MsgPut{
		Owner:       user.String(),
		BasketDenom: basketDenom,
		Credits:     []*basket.BasketCredit{{Amount: creditsToDeposit.String(), BatchDenom: batchDenom}},
	})
	require.NoError(err)

	amountBasketCoins, err := math.NewDecFromString(pRes.AmountReceived)
	require.NoError(err)

	// take them out of the basket, retiring them
	tRes, err = s.basketServer.Take(s.ctx, &basket.MsgTake{
		Owner:                  user.String(),
		BasketDenom:            basketDenom,
		Amount:                 amountBasketCoins.String(),
		RetirementJurisdiction: "US-NY",
		RetireOnTake:           true,
	})
	require.NoError(err)
	require.Len(tRes.Credits, 1) // should only be one credit
	require.Equal(creditsToDeposit.String(), tRes.Credits[0].Amount)

	// check retired balance, should be equal to the amount we put in
	cbRes, err := s.queryClient.Balance(s.ctx, &core.QueryBalanceRequest{
		Address:    user.String(),
		BatchDenom: batchDenom,
	})
	require.NoError(err)
	require.Equal(creditsToDeposit.String(), cbRes.Balance.RetiredAmount)
}

func (s *IntegrationTestSuite) createClassAndIssueBatch(admin, recipient sdk.AccAddress, creditTypeAbbrev, tradableAmount, startStr, endStr string) (string, string) {
	require := s.Require()
	// fund the account so this doesn't fail
	s.fundAccount(admin, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 20000000)))

	cRes, err := s.msgClient.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            admin.String(),
		Issuers:          []string{admin.String()},
		Metadata:         "",
		CreditTypeAbbrev: creditTypeAbbrev,
		Fee:              &createClassFee,
	})
	require.NoError(err)
	classId := cRes.ClassId
	start, err := types.ParseDate("start date", startStr)
	require.NoError(err)
	end, err := types.ParseDate("end date", endStr)
	require.NoError(err)
	pRes, err := s.msgClient.CreateProject(s.ctx, &core.MsgCreateProject{
		Admin:        admin.String(),
		ClassId:      classId,
		Metadata:     "",
		Jurisdiction: "US-NY",
	})
	require.NoError(err)
	bRes, err := s.msgClient.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    admin.String(),
		ProjectId: pRes.ProjectId,
		Issuance:  []*core.BatchIssuance{{Recipient: recipient.String(), TradableAmount: tradableAmount}},
		Metadata:  "",
		StartDate: &start,
		EndDate:   &end,
	})
	require.NoError(err)
	batchDenom := bRes.BatchDenom
	return classId, batchDenom
}

func (s *IntegrationTestSuite) TestScenario() {
	admin := s.signers[0]
	issuer1 := s.signers[1].String()
	issuer2 := s.signers[2].String()
	addr1 := s.signers[3].String()
	addr2 := s.signers[4].String()
	acc3 := s.signers[5]
	addr3 := acc3.String()
	addr4 := s.signers[6].String()
	acc5 := s.signers[7]
	addr5 := acc5.String()

	// create class with insufficient funds and it should fail
	createClsRes, err := s.msgClient.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            admin.String(),
		Issuers:          []string{issuer1, issuer2},
		Metadata:         "",
		CreditTypeAbbrev: "C",
		Fee:              nil,
	})
	s.Require().Error(err)
	s.Require().Nil(createClsRes)

	// create class with sufficient funds and it should succeed
	s.fundAccount(admin, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 4*core.DefaultCreditClassFee.Int64())))
	adminBalanceBefore := s.bankKeeper.GetBalance(s.sdkCtx, admin, sdk.DefaultBondDenom)

	createClsRes, err = s.msgClient.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            admin.String(),
		Issuers:          []string{issuer1, issuer2},
		Metadata:         "",
		CreditTypeAbbrev: "C",
		Fee:              &createClassFee,
	})
	s.Require().NoError(err)
	classId := createClsRes.ClassId

	adminBalanceAfter := s.bankKeeper.GetBalance(s.sdkCtx, admin, sdk.DefaultBondDenom)
	expectedBalance := adminBalanceAfter.Add(createClassFee)
	s.Require().True(adminBalanceBefore.Equal(expectedBalance), "actual balance: %v \t expected: %v", adminBalanceAfter, expectedBalance)

	// create project
	createProjectRes, err := s.msgClient.CreateProject(s.ctx, &core.MsgCreateProject{
		ClassId:      classId,
		Admin:        issuer1,
		Metadata:     "metadata",
		Jurisdiction: "AQ",
	})
	s.Require().NoError(err)
	s.Require().NotNil(createProjectRes)
	s.Require().Equal("C02-001", createProjectRes.ProjectId)
	projectId := createProjectRes.ProjectId

	// create batch
	t0, t1, t2 := "10.37", "1007.3869", "100"
	tSupply0 := "1117.7569"
	r0, r1, r2 := "4.286", "10000.45899", "0"
	rSupply0 := "10004.74499"

	time1 := time.Now()
	time2 := time.Now()

	// Batch creation should succeed with StartDate before EndDate, and valid data
	createBatchRes, err := s.msgClient.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    issuer1,
		ProjectId: projectId,
		StartDate: &time1,
		EndDate:   &time2,
		Issuance: []*core.BatchIssuance{
			{
				Recipient:              addr1,
				TradableAmount:         t0,
				RetiredAmount:          r0,
				RetirementJurisdiction: "GB",
			},
			{
				Recipient:              addr2,
				TradableAmount:         t1,
				RetiredAmount:          r1,
				RetirementJurisdiction: "BF",
			},
			{
				Recipient:              addr4,
				TradableAmount:         t2,
				RetiredAmount:          r2,
				RetirementJurisdiction: "",
			},
			{
				Recipient:              addr5,
				RetirementJurisdiction: "",
			},
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(createBatchRes)

	batchDenom := createBatchRes.BatchDenom
	s.Require().NotEmpty(batchDenom)

	// query balances
	queryBalanceRes, err := s.queryClient.Balance(s.ctx, &core.QueryBalanceRequest{
		Address:    addr1,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal(t0, queryBalanceRes.Balance.TradableAmount)
	s.Require().Equal(r0, queryBalanceRes.Balance.RetiredAmount)

	queryBalanceRes, err = s.queryClient.Balance(s.ctx, &core.QueryBalanceRequest{
		Address:    addr2,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal(t1, queryBalanceRes.Balance.TradableAmount)
	s.Require().Equal(r1, queryBalanceRes.Balance.RetiredAmount)

	queryBalanceRes, err = s.queryClient.Balance(s.ctx, &core.QueryBalanceRequest{
		Address:    addr4,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal(t2, queryBalanceRes.Balance.TradableAmount)
	s.Require().Equal(r2, queryBalanceRes.Balance.RetiredAmount)

	// if we didn't issue tradable or retired balances, they'll be default to zero.
	queryBalanceRes, err = s.queryClient.Balance(s.ctx, &core.QueryBalanceRequest{
		Address:    addr5,
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal("0", queryBalanceRes.Balance.TradableAmount)
	s.Require().Equal("0", queryBalanceRes.Balance.RetiredAmount)

	// query supply
	querySupplyRes, err := s.queryClient.Supply(s.ctx, &core.QuerySupplyRequest{BatchDenom: batchDenom})
	s.Require().NoError(err)
	s.Require().NotNil(querySupplyRes)
	s.Require().Equal(tSupply0, querySupplyRes.TradableAmount)
	s.Require().Equal(rSupply0, querySupplyRes.RetiredAmount)

	// cancel credits
	cancelCases := []struct {
		name               string
		owner              string
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
			owner:         addr4,
			toCancel:      "101",
			expectErr:     true,
			expErrMessage: "insufficient funds",
		},
		{
			name:          "can't cancel with a higher precision than the credit type",
			owner:         addr4,
			toCancel:      "0.1234567",
			expectErr:     true,
			expErrMessage: "exceeds maximum decimal places",
		},
		{
			name:          "can't cancel no credits",
			owner:         addr4,
			toCancel:      "0",
			expectErr:     true,
			expErrMessage: "expected a positive decimal",
		},
		{
			name:               "can cancel a small amount of credits",
			owner:              addr4,
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
			owner:              addr4,
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
			owner:         addr4,
			toCancel:      "1",
			expectErr:     true,
			expErrMessage: "insufficient funds",
		},
		{
			name:               "can cancel from account with positive retired balance",
			owner:              addr1,
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
			_, err := s.msgClient.Cancel(s.ctx, &core.MsgCancel{
				Owner: tc.owner,
				Credits: []*core.Credits{
					{
						BatchDenom: batchDenom,
						Amount:     tc.toCancel,
					},
				},
				Reason: "bridging assets to another chain",
			})

			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMessage)
			} else {
				s.Require().NoError(err)

				// query balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &core.QueryBalanceRequest{
					Address:    tc.owner,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.assertDecStrEqual(tc.expTradable, queryBalanceRes.Balance.TradableAmount)
				s.assertDecStrEqual(tc.expRetired, queryBalanceRes.Balance.RetiredAmount)

				// query supply
				querySupplyRes, err = s.queryClient.Supply(s.ctx, &core.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.assertDecStrEqual(tc.expTradableSupply, querySupplyRes.TradableAmount)
				s.assertDecStrEqual(rSupply0, querySupplyRes.RetiredAmount)
				s.assertDecStrEqual(tc.expAmountCancelled, querySupplyRes.CancelledAmount)

				// query batch
				queryBatchRes, err := s.queryClient.Batch(s.ctx, &core.QueryBatchRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(queryBatchRes)
			}
		})
	}

	// retire credits
	retireCases := []struct {
		name              string
		toRetire          string
		jurisdiction      string
		expectErr         bool
		expTradable       string
		expRetired        string
		expTradableSupply string
		expRetiredSupply  string
		expErrMessage     string
	}{
		{
			name:          "cannot retire more credits than are tradable",
			toRetire:      "10.371",
			jurisdiction:  "AF",
			expectErr:     true,
			expErrMessage: "insufficient funds",
		},
		{
			name:          "can't use more precision than the credit type allows (6)",
			toRetire:      "10.00000001",
			jurisdiction:  "AF",
			expectErr:     true,
			expErrMessage: "exceeds maximum decimal places",
		},
		{
			name:          "can't retire to an invalid country",
			toRetire:      "0.0001",
			jurisdiction:  "ZZZ",
			expectErr:     true,
			expErrMessage: "invalid jurisdiction",
		},
		{
			name:          "can't retire to an invalid region",
			toRetire:      "0.0001",
			jurisdiction:  "AF-ZZZZ",
			expectErr:     true,
			expErrMessage: "invalid jurisdiction",
		},
		{
			name:          "can't retire to an invalid postal code",
			toRetire:      "0.0001",
			jurisdiction:  "AF-BDS 0123456789012345678901234567890123456789012345678901234567890123456789",
			expectErr:     true,
			expErrMessage: "invalid jurisdiction",
		},
		{
			name:          "can't retire without a jurisdiction",
			toRetire:      "0.0001",
			jurisdiction:  "",
			expectErr:     true,
			expErrMessage: "invalid jurisdiction",
		},
		{
			name:              "can retire a small amount of credits",
			toRetire:          "0.0001",
			jurisdiction:      "AF",
			expectErr:         false,
			expTradable:       "9.3699",
			expRetired:        "4.2861",
			expTradableSupply: "1016.7568",
			expRetiredSupply:  "10004.74509",
		},
		{
			name:              "can retire more credits",
			toRetire:          "9",
			jurisdiction:      "AF-BDS",
			expectErr:         false,
			expTradable:       "0.3699",
			expRetired:        "13.2861",
			expTradableSupply: "1007.7568",
			expRetiredSupply:  "10013.74509",
		},
		{
			name:              "can retire all credits",
			toRetire:          "0.3699",
			jurisdiction:      "AF-BDS 12345",
			expectErr:         false,
			expTradable:       "0",
			expRetired:        "13.656",
			expTradableSupply: "1007.3869",
			expRetiredSupply:  "10014.11499",
		},
		{
			name:          "can't retire any more credits",
			toRetire:      "1",
			jurisdiction:  "AF-BDS",
			expectErr:     true,
			expErrMessage: "insufficient funds",
		},
	}

	for _, tc := range retireCases {
		tc := tc
		s.Run(tc.name, func() {
			_, err := s.msgClient.Retire(s.ctx, &core.MsgRetire{
				Owner: addr1,
				Credits: []*core.Credits{
					{
						BatchDenom: batchDenom,
						Amount:     tc.toRetire,
					},
				},
				Jurisdiction: tc.jurisdiction,
			})

			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMessage)
			} else {
				s.Require().NoError(err)

				// query balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &core.QueryBalanceRequest{
					Address:    addr1,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.assertDecStrEqual(tc.expTradable, queryBalanceRes.Balance.TradableAmount)
				s.assertDecStrEqual(tc.expRetired, queryBalanceRes.Balance.RetiredAmount)

				// query supply
				querySupplyRes, err = s.queryClient.Supply(s.ctx, &core.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.assertDecStrEqual(tc.expTradableSupply, querySupplyRes.TradableAmount)
				s.assertDecStrEqual(tc.expRetiredSupply, querySupplyRes.RetiredAmount)
			}
		})
	}

	sendCases := []struct {
		name                 string
		sendTradable         string
		sendRetired          string
		jurisdiction         string
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
			name:          "can't send an amount with more decimal places than allowed precision (6)",
			sendTradable:  "2.123456789",
			sendRetired:   "10.123456789",
			jurisdiction:  "AF",
			expectErr:     true,
			expErrMessage: "exceeds maximum decimal places",
		},
		{
			name:          "can't send more tradable than is tradable",
			sendTradable:  "2000",
			sendRetired:   "10",
			jurisdiction:  "AF",
			expectErr:     true,
			expErrMessage: "insufficient funds",
		},
		{
			name:          "can't send more retired than is tradable",
			sendTradable:  "10",
			sendRetired:   "2000",
			jurisdiction:  "AF",
			expectErr:     true,
			expErrMessage: "insufficient funds",
		},
		{
			name:          "can't send to an invalid country",
			sendTradable:  "10",
			sendRetired:   "20",
			jurisdiction:  "ZZZ",
			expectErr:     true,
			expErrMessage: "invalid jurisdiction",
		},
		{
			name:          "can't send to an invalid region",
			sendTradable:  "10",
			sendRetired:   "20",
			jurisdiction:  "AF-ZZZZ",
			expectErr:     true,
			expErrMessage: "invalid jurisdiction",
		},
		{
			name:          "can't send to an invalid postal code",
			sendTradable:  "10",
			sendRetired:   "20",
			jurisdiction:  "AF-BDS 0123456789012345678901234567890123456789012345678901234567890123456789",
			expectErr:     true,
			expErrMessage: "invalid jurisdiction",
		},
		{
			name:                 "can send some",
			sendTradable:         "10",
			sendRetired:          "20",
			jurisdiction:         "AF",
			expectErr:            false,
			expTradableSender:    "977.3869",
			expRetiredSender:     "10000.45899",
			expTradableRecipient: "10",
			expRetiredRecipient:  "20",
			expTradableSupply:    "987.3869",
			expRetiredSupply:     "10034.11499",
		},
		{
			name:                 "can send with no retirement jurisdiction",
			sendTradable:         "10",
			sendRetired:          "0",
			jurisdiction:         "",
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
			jurisdiction:         "AF",
			expectErr:            false,
			expTradableSender:    "0",
			expRetiredSender:     "10000.45899",
			expTradableRecipient: "87.3869",
			expRetiredRecipient:  "920",
			expTradableSupply:    "87.3869",
			expRetiredSupply:     "10934.11499",
		},
		{
			name:          "can't send any more",
			sendTradable:  "1",
			sendRetired:   "1",
			expectErr:     true,
			jurisdiction:  "AF",
			expErrMessage: "insufficient funds",
		},
	}

	for _, tc := range sendCases {
		tc := tc
		s.Run(tc.name, func() {
			_, err := s.msgClient.Send(s.ctx, &core.MsgSend{
				Sender:    addr2,
				Recipient: addr3,
				Credits: []*core.MsgSend_SendCredits{
					{
						BatchDenom:             batchDenom,
						TradableAmount:         tc.sendTradable,
						RetiredAmount:          tc.sendRetired,
						RetirementJurisdiction: tc.jurisdiction,
					},
				},
			})

			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrMessage)
			} else {
				s.Require().NoError(err)

				// query sender balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &core.QueryBalanceRequest{
					Address:    addr2,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.assertDecStrEqual(tc.expTradableSender, queryBalanceRes.Balance.TradableAmount)
				s.assertDecStrEqual(tc.expRetiredSender, queryBalanceRes.Balance.RetiredAmount)

				// query recipient balance
				queryBalanceRes, err = s.queryClient.Balance(s.ctx, &core.QueryBalanceRequest{
					Address:    addr3,
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.assertDecStrEqual(tc.expTradableRecipient, queryBalanceRes.Balance.TradableAmount)
				s.assertDecStrEqual(tc.expRetiredRecipient, queryBalanceRes.Balance.RetiredAmount)

				// query supply
				querySupplyRes, err = s.queryClient.Supply(s.ctx, &core.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.assertDecStrEqual(tc.expTradableSupply, querySupplyRes.TradableAmount)
				s.assertDecStrEqual(tc.expRetiredSupply, querySupplyRes.RetiredAmount)
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
		errMsg           string
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
			errMsg:           "not allowed",
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
			errMsg:           "not allowed",
		},
	}

	for _, tc := range allowlistCases {
		tc := tc
		s.Run(tc.name, func() {
			s.paramSpace.Set(s.sdkCtx, core.KeyAllowedClassCreators, tc.allowlist)
			s.paramSpace.Set(s.sdkCtx, core.KeyAllowlistEnabled, tc.allowlistEnabled)

			// fund the creator account
			s.fundAccount(tc.creatorAcc, sdk.NewCoins(sdk.NewCoin("stake", core.DefaultCreditClassFee)))

			createClsRes, err = s.msgClient.CreateClass(s.ctx, &core.MsgCreateClass{
				Admin:            tc.creatorAcc.String(),
				Issuers:          []string{issuer1, issuer2},
				CreditTypeAbbrev: "C",
				Metadata:         "",
				Fee:              &createClassFee,
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

	coinPrice := sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000)
	expiration := time.Date(2030, 01, 01, 0, 0, 0, 0, time.UTC)
	expectedSellOrderIds := []uint64{1, 2}

	sellerAcc := acc3
	order1Qty, order2Qty := "10.54321", "15.54321"
	order1QtyDec, err := math.NewDecFromString(order1Qty)
	s.Require().NoError(err)
	order2QtyDec, err := math.NewDecFromString(order2Qty)
	s.Require().NoError(err)
	createSellOrder, err := s.marketServer.Sell(s.ctx, &marketplace.MsgSell{
		Seller: sellerAcc.String(),
		Orders: []*marketplace.MsgSell_Order{
			{
				BatchDenom:        batchDenom,
				Quantity:          order1Qty,
				AskPrice:          &coinPrice,
				DisableAutoRetire: false,
				Expiration:        &expiration,
			},
			{
				BatchDenom:        batchDenom,
				Quantity:          order2Qty,
				AskPrice:          &coinPrice,
				DisableAutoRetire: true,
				Expiration:        &expiration,
			},
		},
	})
	s.Require().Nil(err)
	s.Require().Equal(expectedSellOrderIds, createSellOrder.SellOrderIds)
	orderId1 := createSellOrder.SellOrderIds[0]
	orderId2 := createSellOrder.SellOrderIds[1]

	// now we buy these orders
	buyerAcc := acc5
	// (10.54321 + 15.54321) * 1000000 = 26,086,420
	expectedTotalCost := sdk.NewInt64Coin(coinPrice.Denom, 26_086_420)
	// this is the exact amount it should cost to purchase both orders
	s.fundAccount(buyerAcc, sdk.Coins{expectedTotalCost})

	buyerAccBefore := s.getAccountInfo(buyerAcc, batchDenom, coinPrice.Denom)
	sellerAccBefore := s.getAccountInfo(sellerAcc, batchDenom, coinPrice.Denom)
	_, err = s.marketServer.BuyDirect(s.ctx, &marketplace.MsgBuyDirect{
		Buyer: buyerAcc.String(),
		Orders: []*marketplace.MsgBuyDirect_Order{
			{
				SellOrderId:            orderId1,
				Quantity:               order1Qty,
				BidPrice:               &coinPrice,
				DisableAutoRetire:      false,
				RetirementJurisdiction: "US-OR",
			},
			{
				SellOrderId:       orderId2,
				Quantity:          order2Qty,
				BidPrice:          &coinPrice,
				DisableAutoRetire: true,
			},
		},
	})
	s.Require().NoError(err)
	buyerAccAfter := s.getAccountInfo(buyerAcc, batchDenom, coinPrice.Denom)
	sellerAccAfter := s.getAccountInfo(sellerAcc, batchDenom, coinPrice.Denom)

	s.assertSellerBalancesUpdated(sellerAccBefore, sellerAccAfter, order2QtyDec, order1QtyDec, expectedTotalCost)
	s.assertBuyerBalancesUpdated(buyerAccBefore, buyerAccAfter, order2QtyDec, order1QtyDec, expectedTotalCost)
}

type accountInfo struct {
	address                     sdk.AccAddress
	tradable, retired, escrowed math.Dec
	bankBalance                 sdk.Coin
}

func (s *IntegrationTestSuite) assertSellerBalancesUpdated(accBefore, accAfter accountInfo, tradable, retired math.Dec, totalCost sdk.Coin) {
	expectedEscrowed := accBefore.escrowed // account before the order was bought. should have tradable + retired in escrow

	// subtract the tradable+retired amounts from escrow
	var err error
	expectedEscrowed, err = expectedEscrowed.Sub(tradable)
	s.Require().NoError(err)
	expectedEscrowed, err = expectedEscrowed.Sub(retired)
	s.Require().NoError(err)

	s.Require().Equal(expectedEscrowed.String(), accAfter.escrowed.String())
	s.Require().Equal(accBefore.tradable.String(), accAfter.tradable.String())
	s.Require().Equal(accBefore.retired.String(), accAfter.retired.String())

	expectedBankBal := accBefore.bankBalance
	expectedBankBal = expectedBankBal.Add(totalCost)
	s.Require().Equal(expectedBankBal, accAfter.bankBalance)
}

func (s *IntegrationTestSuite) assertBuyerBalancesUpdated(accBefore, accAfter accountInfo, tradable, retired math.Dec, totalCost sdk.Coin) {

	expectedTradable := accBefore.tradable
	expectedRetired := accBefore.retired

	var err error
	expectedTradable, err = expectedTradable.Add(tradable)
	s.Require().NoError(err)
	expectedRetired, err = expectedRetired.Add(retired)
	s.Require().NoError(err)

	s.Require().True(accBefore.escrowed.Equal(accAfter.escrowed))
	s.Require().True(expectedTradable.Equal(accAfter.tradable), fmt.Sprintf("expected %v got %v", expectedTradable, accAfter.tradable))
	s.Require().True(expectedRetired.Equal(accAfter.retired), fmt.Sprintf("expected %v got %v", expectedRetired, accAfter.retired))

	expectedBankBal := accBefore.bankBalance
	expectedBankBal = expectedBankBal.Sub(totalCost)
	s.Require().True(expectedBankBal.Equal(accAfter.bankBalance))
}

func (s *IntegrationTestSuite) getUserBalance(addr sdk.AccAddress, denom string) sdk.Coin {
	require := s.Require()
	bRes, err := s.bankKeeper.Balance(s.ctx, &banktypes.QueryBalanceRequest{
		Address: addr.String(),
		Denom:   denom,
	})
	require.NoError(err)
	return *bRes.Balance
}

func (s *IntegrationTestSuite) fundAccount(addr sdk.AccAddress, amounts sdk.Coins) {
	err := s.bankKeeper.MintCoins(s.sdkCtx, minttypes.ModuleName, amounts)
	s.Require().NoError(err)
	err = s.bankKeeper.SendCoinsFromModuleToAccount(s.sdkCtx, minttypes.ModuleName, addr, amounts)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) assertDecStrEqual(d1, d2 string) {
	dec1, err := math.NewDecFromString(d1)
	s.Require().NoError(err)
	dec2, err := math.NewDecFromString(d2)
	s.Require().NoError(err)
	s.Require().True(dec1.Equal(dec2), "%v does not equal %v", dec1, dec2)
}

func (s *IntegrationTestSuite) createClass(admin, creditTypeAbbrev, metadata string, issuers []string) string {
	res, err := s.msgClient.CreateClass(s.ctx, &core.MsgCreateClass{Admin: admin, Issuers: issuers, Metadata: metadata, CreditTypeAbbrev: creditTypeAbbrev, Fee: &createClassFee})
	s.Require().NoError(err)
	return res.ClassId
}

func (s *IntegrationTestSuite) getAccountInfo(addr sdk.AccAddress, batchDenom, bankDenom string) accountInfo {
	coinBalance := s.getUserBalance(addr, bankDenom)
	bal := s.getUserBatchBalance(addr, batchDenom)
	t, r, e := s.getDecimalsFromBalance(bal)
	return accountInfo{
		address:     addr,
		tradable:    t,
		retired:     r,
		escrowed:    e,
		bankBalance: coinBalance,
	}
}

func (s *IntegrationTestSuite) getUserBatchBalance(addr sdk.AccAddress, denom string) *core.BatchBalanceInfo {
	bal, err := s.queryClient.Balance(s.ctx, &core.QueryBalanceRequest{
		Address:    addr.String(),
		BatchDenom: denom,
	})
	s.Require().NoError(err)
	return bal.Balance
}

func (s *IntegrationTestSuite) getDecimalsFromBalance(bal *core.BatchBalanceInfo) (tradable, retired, escrowed math.Dec) {
	decs, err := utils.GetNonNegativeFixedDecs(6, bal.TradableAmount, bal.RetiredAmount, bal.EscrowedAmount)
	s.Require().NoError(err)
	return decs[0], decs[1], decs[2]
}
