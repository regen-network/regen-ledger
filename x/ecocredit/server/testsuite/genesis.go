package testsuite

import (
	"encoding/json"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func (s *GenesisTestSuite) TestInitExportGenesis() {
	require := s.Require()
	ctx := s.genesisCtx

	// Set the param set to empty values to properly test init
	var ecocreditParams core.Params
	s.paramSpace.SetParamSet(ctx, &ecocreditParams)

	defaultParams := core.DefaultParams()
	paramsJSON, err := s.fixture.Codec().MarshalJSON(&defaultParams)
	require.NoError(err)

	classIssuers := []api.ClassIssuer{
		{ClassKey: 1, Issuer: sdk.AccAddress("addr1")},
		{ClassKey: 1, Issuer: sdk.AccAddress("addr2")},
		{ClassKey: 2, Issuer: sdk.AccAddress("addr2")},
		{ClassKey: 2, Issuer: sdk.AccAddress("addr3")},
	}
	classIssuersJSON, err := json.Marshal(classIssuers)
	require.NoError(err)

	classes := []api.Class{
		{Id: "BIO001", Admin: sdk.AccAddress("addr1"), Metadata: "metadata", CreditTypeAbbrev: "BIO"},
		{Id: "BIO002", Admin: sdk.AccAddress("addr2"), Metadata: "metadata", CreditTypeAbbrev: "BIO"},
	}
	classJSON, err := json.Marshal(classes)
	require.NoError(err)

	projects := []api.Project{
		{Id: "C01-001", Admin: sdk.AccAddress("addr1"), ClassKey: 1, Jurisdiction: "AQ", Metadata: "metadata"},
		{Id: "C01-002", Admin: sdk.AccAddress("addr2"), ClassKey: 2, Jurisdiction: "AQ", Metadata: "metadata"},
	}
	projectJSON, err := json.Marshal(projects)
	require.NoError(err)

	batches := []api.Batch{
		{Issuer: sdk.AccAddress("addr1"), ProjectKey: 1, Denom: "BIO01-00000000-00000000-001", Metadata: "metadata"},
		{Issuer: nil, ProjectKey: 1, Denom: "BIO02-0000000-0000000-001", Metadata: "metadata"},
	}
	batchJSON, err := json.Marshal(batches)
	require.NoError(err)

	balances := []api.BatchBalance{
		{Address: sdk.AccAddress("addr1"), BatchKey: 1, TradableAmount: "90.003", RetiredAmount: "9.997", EscrowedAmount: ""},
	}
	batchBalancesJSON, err := json.Marshal(balances)
	require.NoError(err)

	supply := []api.BatchSupply{
		{BatchKey: 1, TradableAmount: "90.003", RetiredAmount: "9.997", CancelledAmount: ""},
	}
	batchSupplyJSON, err := json.Marshal(supply)
	require.NoError(err)

	classSeq := []api.ClassSequence{{CreditTypeAbbrev: "BIO", NextSequence: 3}}
	classSeqJSON, err := json.Marshal(classSeq)
	require.NoError(err)

	batchSeq := []api.BatchSequence{{ProjectKey: 1, NextSequence: 3}}
	batchSeqJSON, err := json.Marshal(batchSeq)
	require.NoError(err)

	projectSeq := []api.ProjectSequence{{ClassKey: 1, NextSequence: 3}}
	projectSeqJSON, err := json.Marshal(projectSeq)
	require.NoError(err)

	wrapper := map[string]json.RawMessage{}
	wrapper[string(proto.MessageName(&api.Class{}))] = classJSON
	wrapper[string(proto.MessageName(&api.ClassIssuer{}))] = classIssuersJSON
	wrapper[string(proto.MessageName(&api.Project{}))] = projectJSON
	wrapper[string(proto.MessageName(&api.Batch{}))] = batchJSON
	wrapper[string(proto.MessageName(&api.BatchBalance{}))] = batchBalancesJSON
	wrapper[string(proto.MessageName(&api.BatchSupply{}))] = batchSupplyJSON
	wrapper[string(proto.MessageName(&api.ClassSequence{}))] = classSeqJSON
	wrapper[string(proto.MessageName(&api.BatchSequence{}))] = batchSeqJSON
	wrapper[string(proto.MessageName(&api.ProjectSequence{}))] = projectSeqJSON
	wrapper[string(proto.MessageName(&api.Params{}))] = paramsJSON

	bz, err := json.Marshal(wrapper)
	require.NoError(err)
	wrapper = map[string]json.RawMessage{}
	wrapper["ecocredit"] = bz

	_, err = s.fixture.InitGenesis(s.genesisCtx, wrapper)
	require.NoError(err)

	exported := s.exportGenesisState(s.genesisCtx)
	require.NotNil(exported)

}

func (s *GenesisTestSuite) exportGenesisState(ctx sdk.Context) map[string]json.RawMessage {
	require := s.Require()
	exported, err := s.fixture.ExportGenesis(ctx)
	require.NoError(err)

	var wrapper map[string]json.RawMessage
	err = json.Unmarshal(exported[ecocredit.ModuleName], &wrapper)
	require.NoError(err)

	return wrapper
}

type GenesisTestSuite struct {
	suite.Suite

	fixtureFactory testutil.Factory
	fixture        testutil.Fixture
	signers        []sdk.AccAddress

	paramSpace paramstypes.Subspace
	bankKeeper bankkeeper.Keeper

	genesisCtx sdk.Context
}

func NewGenesisTestSuite(fixtureFactory testutil.Factory, paramSpace paramstypes.Subspace, bankKeeper bankkeeper.BaseKeeper) *GenesisTestSuite {
	return &GenesisTestSuite{
		fixtureFactory: fixtureFactory,
		paramSpace:     paramSpace,
		bankKeeper:     bankKeeper,
	}
}

func (s *GenesisTestSuite) SetupSuite() {
	s.fixture = s.fixtureFactory.Setup()

	blockTime := time.Now().UTC()

	sdkCtx := sdk.UnwrapSDKContext(s.fixture.Context()).WithBlockTime(blockTime)
	s.genesisCtx = sdkCtx

	s.signers = s.fixture.Signers()
	s.Require().GreaterOrEqual(len(s.signers), 8)
}
