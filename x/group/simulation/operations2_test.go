package simulation_test

// import (
// 	"context"
// 	"math/rand"
// 	"testing"
// 	"time"

// 	"github.com/cosmos/cosmos-sdk/baseapp"
// 	"github.com/cosmos/cosmos-sdk/codec"
// 	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
// 	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
// 	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
// 	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
// 	regen "github.com/regen-network/regen-ledger/app"
// 	"github.com/regen-network/regen-ledger/testutil"
// 	"github.com/regen-network/regen-ledger/types"
// 	"github.com/regen-network/regen-ledger/types/module"
// 	"github.com/regen-network/regen-ledger/types/module/server"
// 	servermodule "github.com/regen-network/regen-ledger/types/module/server"
// 	"github.com/regen-network/regen-ledger/x/group"
// 	groupmodule "github.com/regen-network/regen-ledger/x/group/module"
// 	"github.com/regen-network/regen-ledger/x/group/simulation"
// 	"github.com/stretchr/testify/suite"
// 	dbm "github.com/tendermint/tm-db"
// )

// type Sim2TestSuite struct {
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
// 	govKeeper     govkeeper.Keeper
// 	cdc           codec.Marshaler
// 	blockTime     time.Time
// }

// func (s *Sim2TestSuite) SetupTest() {
// 	s.fixture = s.fixtureFactory.Setup()

// 	s.baseApp = s.fixtureFactory.BaseApp()
// 	s.ctx = s.fixture.Context()

// 	s.blockTime = time.Now().UTC()

// 	// TODO clean up once types.Context merged upstream into sdk.Context
// 	sdkCtx := s.ctx.(types.Context).WithBlockTime(s.blockTime)
// 	s.sdkCtx = sdkCtx
// 	s.ctx = types.Context{Context: sdkCtx}
// 	s.msgClient = group.NewMsgClient(s.fixture.TxConn())
// 	s.queryClient = group.NewQueryClient(s.fixture.QueryConn())

// }

// func (s *Sim2TestSuite) TearDownSuite() {
// 	s.fixture.Teardown()
// }

// func (suite *Sim2TestSuite) getTestingAccounts(r *rand.Rand, n int) []simtypes.Account {
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

// func (suite *Sim2TestSuite) TestWeightedOperations() {
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

// func TestSim2TestSuite(t *testing.T) {
// 	ff := server.NewFixtureFactory(t, 6)
// 	// cdc := ff.Codec()
// 	app := regen.Setup(false)
// 	cdc := app.AppCodec()

// 	baseApp := baseapp.NewBaseApp("test", app.Logger(), dbm.NewMemDB(), nil)
// 	ff.SetBaseApp(baseApp)

// 	ff.SetModules([]module.Module{groupmodule.Module{AccountKeeper: app.AccountKeeper, BankKeeper: app.BankKeeper, GovKeeper: app.GovKeeper}})

// 	s := NewSim2TestSuite(ff, cdc, app.AccountKeeper, app.BankKeeper, app.GovKeeper)

// 	suite.Run(t, s)
// }

// func NewSim2TestSuite(ff *server.FixtureFactory, cdc codec.Marshaler, accKeeper authkeeper.AccountKeeper, bankKeeper bankkeeper.Keeper, govKeeper govkeeper.Keeper) *Sim2TestSuite {
// 	return &Sim2TestSuite{
// 		accountKeeper:  accKeeper,
// 		fixtureFactory: ff,
// 		bankKeeper:     bankKeeper,
// 		govKeeper:      govKeeper,
// 		cdc:            cdc,
// 	}
// }
