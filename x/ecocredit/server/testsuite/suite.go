package testsuite

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory testutil.FixtureFactory
	fixture        testutil.Fixture

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
	mockDist      *mocks.MockDistributionKeeper

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
	createClassFee = sdk.NewInt64Coin("stake", core.DefaultCreditClassFeeTokens.Int64())
)

func NewIntegrationTestSuite(fixtureFactory testutil.FixtureFactory, paramSpace paramstypes.Subspace, bankKeeper bankkeeper.BaseKeeper, accountKeeper authkeeper.AccountKeeper, distKeeper *mocks.MockDistributionKeeper) *IntegrationTestSuite {
	return &IntegrationTestSuite{
		fixtureFactory: fixtureFactory,
		paramSpace:     paramSpace,
		bankKeeper:     bankKeeper,
		accountKeeper:  accountKeeper,
		mockDist:       distKeeper,
	}
}

func (s *IntegrationTestSuite) SetupSuite() {
	sdk.SetCoinDenomRegex(func() string {
		return `[a-zA-Z][a-zA-Z0-9/:._-]{2,127}`
	})

	s.fixture = s.fixtureFactory.Setup()

	s.blockTime = time.Now().UTC()

	// TODO clean up once types.Context merged upstream into sdk.Context
	sdkCtx := s.fixture.Context().(types.Context).WithBlockTime(s.blockTime)
	s.sdkCtx, _ = sdkCtx.CacheContext()
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
	s.genesisCtx = types.Context{Context: sdkCtx}

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

// TODO: reimpl after full submodule integration
//func (s *IntegrationTestSuite) TestBasketScenario() {
//	require := s.Require()
//	user := s.signers[0]
//	user2 := s.signers[1]
//
//	// create a class and issue a batch
//	userTotalCreditBalance, err := math.NewDecFromString("1000000000000000")
//	require.NoError(err)
//	classId, batchDenom := s.createClassAndIssueBatch(user, user, "bazcredits", userTotalCreditBalance.String(), "2020-01-01", "2022-01-01")
//
//	// fund account to create a basket
//	balanceBefore := sdk.NewInt64Coin(basketFeeDenom, 30000)
//	s.fundAccount(user, sdk.NewCoins(balanceBefore))
//	s.mockDist.EXPECT().FundCommunityPool(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(interface{}, interface{}, interface{}) error {
//		err := s.bankKeeper.SendCoinsFromAccountToModule(s.sdkCtx, user, ecocredit.ModuleName, sdk.NewCoins(s.basketFee))
//		return err
//	})
//	// create a basket
//	res, err := s.basketServer.Create(s.ctx, &basket.MsgCreate{
//		Curator:           s.signers[0].String(),
//		Name:              "BASKET",
//		Exponent:          6,
//		DisableAutoRetire: true,
//		CreditTypeAbbrev:  "BAZ",
//		AllowedClasses:    []string{classId},
//		DateCriteria:      nil,
//		Fee:               sdk.NewCoins(s.basketFee),
//	})
//	require.NoError(err)
//	basketDenom := res.BasketDenom
//
//	// check it was created
//	qRes, err := s.basketServer.Baskets(s.ctx, &basket.QueryBasketsRequest{})
//	require.NoError(err)
//	require.Len(qRes.Baskets, 1)
//	require.Equal(qRes.Baskets[0].BasketDenom, basketDenom)
//
//	// assert the fee was paid - the fee mechanism was mocked, but we still call the same underlying SendFromAccountToModule
//	// function so the result is the same
//	balanceAfter := s.getUserBalance(user, basketFeeDenom)
//	require.Equal(balanceAfter.Add(s.basketFee), balanceBefore)
//
//	// put some BAZ credits in the basket
//	creditAmtDeposited := math.NewDecFromInt64(3)
//	pRes, err := s.basketServer.Put(s.ctx, &basket.MsgPut{
//		Owner:       user.String(),
//		BasketDenom: basketDenom,
//		Credits:     []*basket.BasketCredit{{BatchDenom: batchDenom, Amount: creditAmtDeposited.String()}},
//	})
//	require.NoError(err)
//	basketTokensReceived, err := math.NewPositiveDecFromString(pRes.AmountReceived)
//	require.NoError(err)
//
//	// make sure the bank actually has this balance for the user
//	basketBal := s.getUserBalance(user, basketDenom)
//	i64BT, err := basketTokensReceived.Int64()
//	require.NoError(err)
//	require.Equal(i64BT, basketBal.Amount.Int64())
//
//	// make sure the basket has the credits now.
//	basketBalance, err := s.basketServer.BasketBalance(s.ctx, &basket.QueryBasketBalanceRequest{
//		BasketDenom: basketDenom,
//		BatchDenom:  batchDenom,
//	})
//	require.NoError(err)
//	require.Equal(basketBalance.Balance, creditAmtDeposited.String())
//
//	// make sure user doesn't have any of that credit - should error out
//	userCreditBalance, err := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
//		Account:    user.String(),
//		BatchDenom: batchDenom,
//	})
//	require.NoError(err)
//
//	// make sure the core server is properly tracking the user balance
//	newUserTotal, err := userTotalCreditBalance.Sub(creditAmtDeposited)
//	require.NoError(err)
//	require.Equal(newUserTotal.String(), userCreditBalance.TradableAmount)
//
//	// send the basket coins to another account - user2
//	require.NoError(s.bankKeeper.SendCoins(s.sdkCtx, user, user2, sdk.NewCoins(sdk.NewInt64Coin(basketDenom, i64BT))))
//
//	// user2 can take all the credits from the basket
//	tRes, err := s.basketServer.Take(s.ctx, &basket.MsgTake{
//		Owner:              user2.String(),
//		BasketDenom:        basketDenom,
//		Amount:             basketTokensReceived.String(),
//		RetirementLocation: "US-NY",
//		RetireOnTake:       false,
//	})
//	require.NoError(err)
//	require.Equal(tRes.Credits[0].BatchDenom, batchDenom)
//	require.Equal(tRes.Credits[0].Amount, creditAmtDeposited.String())
//
//	// user shouldn't be able to take any since we sent our tokens to user2
//	noRes, err := s.basketServer.Take(s.ctx, &basket.MsgTake{
//		Owner:              user.String(),
//		BasketDenom:        basketDenom,
//		Amount:             basketTokensReceived.String(),
//		RetirementLocation: "US-NY",
//		RetireOnTake:       false,
//	})
//	require.Error(err)
//	require.Contains(err.Error(), sdkerrors.ErrInsufficientFunds.Error())
//	require.Nil(noRes)
//
//	// there should be nothing left in the basket
//	bRes, err := s.basketServer.BasketBalance(s.ctx, &basket.QueryBasketBalanceRequest{
//		BasketDenom: basketDenom,
//		BatchDenom:  batchDenom,
//	})
//	require.Error(err)
//	require.Contains(err.Error(), "not found")
//	require.Nil(bRes)
//
//	// basket token balance of user2 should be empty now
//	endBal := s.getUserBalance(user2, basketDenom)
//	require.True(endBal.Amount.Equal(sdk.NewInt(0)), "ending balance was %s, expected 0", endBal.Amount.String())
//
//	// check retire enabled basket
//
//	s.mockDist.EXPECT().FundCommunityPool(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(interface{}, interface{}, interface{}) error {
//		err := s.bankKeeper.SendCoinsFromAccountToModule(s.sdkCtx, user, ecocredit.ModuleName, sdk.NewCoins(s.basketFee))
//		return err
//	})
//	// create a retire enabled basket
//	resR, err := s.basketServer.Create(s.ctx, &basket.MsgCreate{
//		Curator:           s.signers[0].String(),
//		Name:              "RETIRE",
//		Exponent:          6,
//		DisableAutoRetire: false,
//		CreditTypeAbbrev:  "BAZ",
//		AllowedClasses:    []string{classId},
//		DateCriteria:      nil,
//		Fee:               sdk.NewCoins(s.basketFee),
//	})
//	require.NoError(err)
//	basketDenom = resR.BasketDenom
//
//	creditsToDeposit := math.NewDecFromInt64(3)
//
//	// put some credits in the basket
//	pRes, err = s.basketServer.Put(s.ctx, &basket.MsgPut{
//		Owner:       user.String(),
//		BasketDenom: basketDenom,
//		Credits:     []*basket.BasketCredit{{Amount: creditsToDeposit.String(), BatchDenom: batchDenom}},
//	})
//	require.NoError(err)
//
//	amountBasketCoins, err := math.NewDecFromString(pRes.AmountReceived)
//	require.NoError(err)
//
//	// take them out of the basket, retiring them
//	tRes, err = s.basketServer.Take(s.ctx, &basket.MsgTake{
//		Owner:              user.String(),
//		BasketDenom:        basketDenom,
//		Amount:             amountBasketCoins.String(),
//		RetirementLocation: "US-NY",
//		RetireOnTake:       true,
//	})
//	require.NoError(err)
//	require.Len(tRes.Credits, 1) // should only be one credit
//	require.Equal(creditsToDeposit.String(), tRes.Credits[0].Amount)
//
//	// check retired balance, should be equal to the amount we put in
//	cbRes, err := s.queryClient.Balance(s.ctx, &ecocredit.QueryBalanceRequest{
//		Account:    user.String(),
//		BatchDenom: batchDenom,
//	})
//	require.NoError(err)
//	require.Equal(creditsToDeposit.String(), cbRes.RetiredAmount)
//}

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
	start, err := time.Parse("2006-04-02", startStr)
	require.NoError(err)
	end, err := time.Parse("2006-04-02", endStr)
	require.NoError(err)
	pRes, err := s.msgClient.CreateProject(s.ctx, &core.MsgCreateProject{
		Issuer:          admin.String(),
		ClassId:         classId,
		Metadata:        "",
		ProjectLocation: "US-NY",
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

func (s *IntegrationTestSuite) createClass(admin, creditTypeAbbrev, metadata string, issuers []string) string {
	res, err := s.msgClient.CreateClass(s.ctx, &core.MsgCreateClass{Admin: admin, Issuers: issuers, Metadata: metadata, CreditTypeAbbrev: creditTypeAbbrev, Fee: &createClassFee})
	s.Require().NoError(err)
	return res.ClassId
}

func addDecimalString(t *testing.T, d1, d2 string) math.Dec {
	dec1, err := math.NewDecFromString(d1)
	require.NoError(t, err)
	dec2, err := math.NewDecFromString(d2)
	require.NoError(t, err)

	result, err := dec1.Add(dec2)
	require.NoError(t, err)
	return result
}
