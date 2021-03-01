package simulation_test

// import (
// 	"context"
// 	"math/rand"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/suite"
// 	abci "github.com/tendermint/tendermint/abci/types"
// 	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

// 	"github.com/cosmos/cosmos-sdk/baseapp"
// 	"github.com/cosmos/cosmos-sdk/codec"
// 	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
// 	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	"github.com/cosmos/cosmos-sdk/x/bank"
// 	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
// 	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
// 	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
// 	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

// 	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
// 	"github.com/regen-network/regen-ledger/testutil"
// 	"github.com/regen-network/regen-ledger/types"
// 	"github.com/regen-network/regen-ledger/types/module"
// 	"github.com/regen-network/regen-ledger/types/module/server"
// 	servermodule "github.com/regen-network/regen-ledger/types/module/server"
// 	"github.com/regen-network/regen-ledger/x/group"
// 	groupmodule "github.com/regen-network/regen-ledger/x/group/module"
// 	"github.com/regen-network/regen-ledger/x/group/simulation"
// )

// type Sim1TestSuite struct {
// 	suite.Suite

// 	fixtureFactory *servermodule.FixtureFactory
// 	fixture        testutil.Fixture
// 	baseApp        *baseapp.BaseApp

// 	ctx         context.Context
// 	sdkCtx      sdk.Context
// 	msgClient   group.MsgClient
// 	queryClient group.QueryClient

// 	accountKeeper authkeeper.AccountKeeper
// 	bankKeeper    bankkeeper.Keeper
// 	govKeeper     *govkeeper.Keeper
// 	cdc           codec.Marshaler
// 	blockTime     time.Time
// }

// func NewSim1TestSuite(ff *server.FixtureFactory, cdc codec.Marshaler, accKeeper authkeeper.AccountKeeper, bankKeeper bankkeeper.Keeper, govKeeper *govkeeper.Keeper) *Sim1TestSuite {
// 	return &Sim1TestSuite{
// 		accountKeeper:  accKeeper,
// 		fixtureFactory: ff,
// 		bankKeeper:     bankKeeper,
// 		govKeeper:      govKeeper,
// 		cdc:            cdc,
// 	}
// }

// func (s *Sim1TestSuite) SetupTest() {
// 	s.fixture = s.fixtureFactory.Setup()

// 	s.baseApp = s.fixtureFactory.BaseApp()
// 	s.ctx = s.fixture.Context()

// 	s.blockTime = time.Now().UTC()

// 	// TODO clean up once types.Context merged upstream into sdk.Context
// 	sdkCtx := s.ctx.(types.Context).WithBlockTime(s.blockTime)
// 	s.sdkCtx = sdkCtx
// 	s.ctx = types.Context{Context: sdkCtx}

// 	totalSupply := banktypes.NewSupply(sdk.NewCoins(sdk.NewInt64Coin("test", 400000000)))
// 	s.bankKeeper.SetSupply(sdkCtx, totalSupply)
// 	s.bankKeeper.SetParams(sdkCtx, banktypes.DefaultParams())

// 	s.msgClient = group.NewMsgClient(s.fixture.TxConn())
// 	s.queryClient = group.NewQueryClient(s.fixture.QueryConn())

// }

// func (s *Sim1TestSuite) TearDownSuite() {
// 	s.fixture.Teardown()
// }

// func (suite *Sim1TestSuite) TestWeightedOperations() {
// 	appParams := make(simtypes.AppParams)

// 	weightedOps := simulation.WeightedOperations(appParams, suite.cdc, suite.accountKeeper,
// 		suite.bankKeeper, suite.govKeeper, suite.queryClient,
// 	)

// 	s := rand.NewSource(1)
// 	r := rand.New(s)
// 	accs := suite.getTestingAccounts(r, 3)

// 	expected := []struct {
// 		weight     int
// 		opMsgRoute string
// 		opMsgName  string
// 	}{
// 		{simappparams.DefaultWeightMsgCreateValidator, group.ModuleName, group.TypeMsgCreateGroup},
// 		{simappparams.DefaultWeightMsgCreateValidator, group.ModuleName, group.TypeMsgCreateGroupAccount},
// 		{simappparams.DefaultWeightMsgCreateValidator, group.ModuleName, group.TypeMsgCreateProposal},
// 	}

// 	for i, w := range weightedOps {
// 		operationMsg, _, _ := w.Op()(r, suite.baseApp, suite.sdkCtx, accs, "")
// 		// the following checks are very much dependent from the ordering of the output given
// 		// by WeightedOperations. if the ordering in WeightedOperations changes some tests
// 		// will fail
// 		suite.Require().Equal(expected[i].weight, w.Weight(), "weight should be the same")
// 		suite.Require().Equal(expected[i].opMsgRoute, operationMsg.Route, "route should be the same")
// 		suite.Require().Equal(expected[i].opMsgName, operationMsg.Name, "operation Msg name should be the same")
// 	}
// }

// func (suite *Sim1TestSuite) getTestingAccounts(r *rand.Rand, n int) []simtypes.Account {
// 	accounts := simtypes.RandomAccounts(r, n)

// 	initAmt := sdk.TokensFromConsensusPower(200000)
// 	initCoins := sdk.NewCoins(sdk.NewCoin("foo", initAmt))

// 	// add coins to the accounts
// 	for _, account := range accounts {
// 		acc := suite.accountKeeper.NewAccountWithAddress(suite.sdkCtx, account.Address)
// 		suite.accountKeeper.SetAccount(suite.sdkCtx, acc)
// 		err := suite.bankKeeper.SetBalances(suite.sdkCtx, account.Address, initCoins)
// 		suite.Require().NoError(err)
// 	}

// 	return accounts
// }

// func (suite *Sim1TestSuite) TestSimulateCreateGroup() {
// 	// setup 1 account
// 	s := rand.NewSource(1)
// 	r := rand.New(s)
// 	accounts := suite.getTestingAccounts(r, 1)

// 	// begin a new block
// 	suite.baseApp.BeginBlock(abci.RequestBeginBlock{
// 		Header: tmproto.Header{
// 			Height:  suite.baseApp.LastBlockHeight() + 1,
// 			AppHash: suite.baseApp.LastCommitID().Hash,
// 		},
// 	})

// 	acc := accounts[0]

// 	// execute operation
// 	op := simulation.SimulateMsgCreateGroup(suite.accountKeeper, suite.bankKeeper)
// 	operationMsg, futureOperations, err := op(r, suite.baseApp, suite.sdkCtx, accounts, "")
// 	suite.Require().NoError(err)

// 	var msg group.MsgCreateGroupRequest
// 	err = group.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)
// 	suite.Require().NoError(err)
// 	suite.Require().True(operationMsg.OK)
// 	suite.Require().Equal(acc.Address.String(), msg.Admin)
// 	suite.Require().Len(futureOperations, 0)

// }

// func (suite *Sim1TestSuite) TestSimulateCreateGroupAccount() {
// 	// setup 1 account
// 	s := rand.NewSource(1)
// 	r := rand.New(s)
// 	accounts := suite.getTestingAccounts(r, 1)

// 	// begin a new block
// 	suite.baseApp.BeginBlock(abci.RequestBeginBlock{
// 		Header: tmproto.Header{
// 			Height:  suite.baseApp.LastBlockHeight() + 1,
// 			AppHash: suite.baseApp.LastCommitID().Hash,
// 		},
// 	})

// 	acc := accounts[0]

// 	// execute operation
// 	op := simulation.SimulateMsgCreateGroupAccount(suite.accountKeeper, suite.bankKeeper, nil)
// 	operationMsg, futureOperations, err := op(r, suite.baseApp, suite.sdkCtx, accounts, "")
// 	suite.Require().NoError(err)

// 	var msg group.MsgCreateGroupRequest
// 	err = group.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)
// 	suite.Require().NoError(err)
// 	suite.Require().True(operationMsg.OK)
// 	suite.Require().Equal(acc.Address.String(), msg.Admin)
// 	suite.Require().Len(futureOperations, 0)

// }

// func TestSim1TestSuite(t *testing.T) {
// 	ff := server.NewFixtureFactory(t, 6)
// 	cdc := ff.Codec()
// 	// Setting up bank keeper
// 	banktypes.RegisterInterfaces(cdc.InterfaceRegistry())
// 	authtypes.RegisterInterfaces(cdc.InterfaceRegistry())
// 	govtypes.RegisterInterfaces(cdc.InterfaceRegistry())

// 	group.RegisterTypes(cdc.InterfaceRegistry())

// 	paramsKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
// 	authKey := sdk.NewKVStoreKey(authtypes.StoreKey)
// 	bankKey := sdk.NewKVStoreKey(banktypes.StoreKey)
// 	// govKey := sdk.NewKVStoreKey(govtypes.StoreKey)
// 	// // groupKey := sdk.NewKVStoreKey(group.StoreKey)
// 	// stakingKey := sdk.NewKVStoreKey(stakingtypes.StoreKey)

// 	tkey := sdk.NewTransientStoreKey(paramstypes.TStoreKey)
// 	amino := codec.NewLegacyAmino()

// 	authSubspace := paramstypes.NewSubspace(cdc, amino, paramsKey, tkey, authtypes.ModuleName)
// 	bankSubspace := paramstypes.NewSubspace(cdc, amino, paramsKey, tkey, banktypes.ModuleName)
// 	// govSubspace := paramstypes.NewSubspace(cdc, amino, paramsKey, tkey, govtypes.ModuleName)
// 	// stakingSubspace := paramstypes.NewSubspace(cdc, amino, paramsKey, tkey, stakingtypes.ModuleName)
// 	// groupSubspace := paramstypes.NewSubspace(cdc, amino, paramsKey, tkey, group.ModuleName)

// 	maccPerms := map[string][]string{
// 		govtypes.ModuleName: {authtypes.Burner},
// 	}
// 	accountKeeper := authkeeper.NewAccountKeeper(
// 		cdc, authKey, authSubspace, authtypes.ProtoBaseAccount, maccPerms,
// 	)

// 	bankKeeper := bankkeeper.NewBaseKeeper(
// 		cdc, bankKey, accountKeeper, bankSubspace, map[string]bool{},
// 	)

// 	// stakingKeeper := stakingkeeper.NewKeeper(cdc, stakingKey, accountKeeper, bankKeeper, stakingSubspace)

// 	// register the proposal types
// 	// govRouter := govtypes.NewRouter()
// 	// govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler)
// 	// govKeeper := govkeeper.NewKeeper(
// 	// 	cdc, govKey, govSubspace, accountKeeper, bankKeeper, stakingKeeper, govRouter,
// 	// )

// 	baseApp := ff.BaseApp()

// 	baseApp.Router().AddRoute(sdk.NewRoute(banktypes.ModuleName, bank.NewHandler(bankKeeper)))
// 	baseApp.MountStore(tkey, sdk.StoreTypeTransient)
// 	baseApp.MountStore(paramsKey, sdk.StoreTypeIAVL)
// 	baseApp.MountStore(authKey, sdk.StoreTypeIAVL)
// 	baseApp.MountStore(bankKey, sdk.StoreTypeIAVL)

// 	ff.SetModules([]module.Module{groupmodule.Module{AccountKeeper: accountKeeper, BankKeeper: bankKeeper}})

// 	s := NewSim1TestSuite(ff, cdc, accountKeeper, bankKeeper, nil)

// 	suite.Run(t, s)
// }
