package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/regen-network/regen-ledger/x/group/simulation"
	"github.com/stretchr/testify/suite"

	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	regen "github.com/regen-network/regen-ledger/app"
	"github.com/regen-network/regen-ledger/x/group"
)

type SimTestSuite struct {
	suite.Suite

	ctx      sdk.Context
	app      *regen.RegenApp
	protoCdc *codec.ProtoCodec
}

func (suite *SimTestSuite) SetupTest() {
	app := regen.Setup(false)
	suite.app = app
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	suite.protoCdc = codec.NewProtoCodec(suite.app.InterfaceRegistry())
}

func (suite *SimTestSuite) TestWeightedOperations() {
	cdc := suite.app.AppCodec()
	appParams := make(simtypes.AppParams)

	weightedOps := simulation.WeightedOperations(appParams, cdc, suite.app.AccountKeeper,
		suite.app.BankKeeper, suite.protoCdc, nil,
	)

	s := rand.NewSource(1)
	r := rand.New(s)
	accs := suite.getTestingAccounts(r, 3)

	expected := []struct {
		weight     int
		opMsgRoute string
		opMsgName  string
	}{
		{simappparams.DefaultWeightMsgCreateValidator, group.ModuleName, simulation.TypeMsgCreateGroup},
		{simappparams.DefaultWeightMsgCreateValidator, group.ModuleName, simulation.TypeMsgCreateGroupAccount},
		{simappparams.DefaultWeightMsgCreateValidator, group.ModuleName, simulation.TypeMsgCreateProposal},
	}

	for i, w := range weightedOps {
		operationMsg, _, _ := w.Op()(r, suite.app.BaseApp, suite.ctx, accs, "")
		// the following checks are very much dependent from the ordering of the output given
		// by WeightedOperations. if the ordering in WeightedOperations changes some tests
		// will fail
		suite.Require().Equal(expected[i].weight, w.Weight(), "weight should be the same")
		suite.Require().Equal(expected[i].opMsgRoute, operationMsg.Route, "route should be the same")
		suite.Require().Equal(expected[i].opMsgName, operationMsg.Name, "operation Msg name should be the same")
	}
}

func (suite *SimTestSuite) getTestingAccounts(r *rand.Rand, n int) []simtypes.Account {
	accounts := simtypes.RandomAccounts(r, n)

	initAmt := sdk.TokensFromConsensusPower(200000)
	initCoins := sdk.NewCoins(sdk.NewCoin("foo", initAmt))

	// add coins to the accounts
	for _, account := range accounts {
		acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, account.Address)
		suite.app.AccountKeeper.SetAccount(suite.ctx, acc)
		err := suite.app.BankKeeper.SetBalances(suite.ctx, account.Address, initCoins)
		suite.Require().NoError(err)
	}

	return accounts
}

func (suite *SimTestSuite) TestSimulateCreateGroup() {
	// setup 1 account
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 1)

	// begin a new block
	suite.app.BeginBlock(abci.RequestBeginBlock{
		Header: tmproto.Header{
			Height:  suite.app.LastBlockHeight() + 1,
			AppHash: suite.app.LastCommitID().Hash,
		},
	})

	acc := accounts[0]

	// execute operation
	op := simulation.SimulateMsgCreateGroup(suite.app.AccountKeeper, suite.app.BankKeeper, suite.protoCdc)
	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, suite.ctx, accounts, "")
	suite.Require().NoError(err)

	var msg group.MsgCreateGroupRequest
	err = suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
	suite.Require().NoError(err)
	suite.Require().True(operationMsg.OK)
	suite.Require().Equal(acc.Address.String(), msg.Admin)
	suite.Require().Len(futureOperations, 0)

}

func (suite *SimTestSuite) TestSimulateCreateGroupAccount() {
	// setup 1 account
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 1)

	// begin a new block
	suite.app.BeginBlock(abci.RequestBeginBlock{
		Header: tmproto.Header{
			Height:  suite.app.LastBlockHeight() + 1,
			AppHash: suite.app.LastCommitID().Hash,
		},
	})

	acc := accounts[0]

	// execute operation
	op := simulation.SimulateMsgCreateGroupAccount(suite.app.AccountKeeper, suite.app.BankKeeper, suite.protoCdc, nil)
	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, suite.ctx, accounts, "")
	suite.Require().NoError(err)

	var msg group.MsgCreateGroupRequest
	err = suite.app.AppCodec().UnmarshalJSON(operationMsg.Msg, &msg)
	suite.Require().NoError(err)
	suite.Require().True(operationMsg.OK)
	suite.Require().Equal(acc.Address.String(), msg.Admin)
	suite.Require().Len(futureOperations, 0)

}

func TestSimTestSuite(t *testing.T) {
	suite.Run(t, new(SimTestSuite))
}
