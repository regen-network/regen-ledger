package testsuite

import (
	"encoding/json"
	"time"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func (s *GenesisTestSuite) TestInitExportGenesis() {
	require := s.Require()
	ctx := s.genesisCtx

	// Set the param set to empty values to properly test init
	var ecocreditParams core.Params
	s.paramSpace.SetParamSet(ctx.Context, &ecocreditParams)

	defaultParams := core.DefaultParams()
	paramsJSON, err := s.fixture.Codec().MarshalJSON(&defaultParams)
	require.NoError(err)

	classIssuersJSON := `[
	{"class_key":"1","issuer":"1ygCfmJaPVMIvVEcpx6r+2gpurM="},
	{"class_key":"1","issuer":"KoXfzfqe+V/9x7C4XjnqDFB2Tl4="},
	{"class_key":"2","issuer":"KoXfzfqe+V/9x7C4XjnqDFB2Tl4="},
	{"class_key":"2","issuer":"lEjmu9Vooa24qp9vCMIlXGrMZoU="}
	]`

	classInfoJSON := `[
		{"id":"BIO001","admin":"4A/V6LMEL2lZv9PZnkWSIDQzZM4=","metadata":"credit class metadata","credit_type_abbrev":"BIO"},
		{"id":"BIO02","admin":"HK9YDsBMN1hU8tjfLTNy+qjbqLE=","metadata":"credit class metadata","credit_type_abbrev":"BIO"}	
	]`

	projectInfoJSON := `[
		{"id":"C01-001","admin":"gPFuHL7Hn+uVYD6XOR00du3C/Xg=","class_key":"1","jurisdiction":"AQ","metadata":"project metadata"},
		{"id":"C01-002","admin":"CHkV2Tv6A7RXPJYTivVklbxXWP8=","class_key":"2","jurisdiction":"AQ","metadata":"project metadata"}	
	]`

	batchInfoJSON := `[
	{"issuer":"WCBEyNFP/N5RoS4h43AqkjC6zA8=","project_key":"1","denom":"BIO01-00000000-00000000-001","metadata":"batch metadata","start_date":null,"end_date":null,"issuance_date":"2022-04-08T10:40:10.774108141Z"},
	{"issuer":null,"project_key":"1","denom":"BIO02-00000000-00000000-001","metadata":"batch metadata","start_date":null,"end_date":null,"issuance_date":"2022-04-08T10:40:10.774108556Z"}
	]`

	batchBalancesJSON := `[
	{"address":"gydQIvR2RUi0N1RJnmgOLVSkcd4=","batch_key":"1","tradable":"90.003","retired":"9.997","escrowed":""}
	]`

	batchSupplyJSON := `[
		{"batch_key":"1","tradable_amount":"90.003","retired_amount":"9.997","cancelled_amount":""}
	]`

	classSeqJSON := `[{"credit_type_abbrev":"BIO","next_sequence":"3"}]`
	batchSeqJSON := `[{"project_key":"1","next_sequence":"3"}]`
	projectSeqJSON := `[{"class_key":"1","next_sequence":"3"}]`

	wrapper := map[string]json.RawMessage{}
	wrapper[gogoproto.MessageName(&core.Class{})] = []byte(classInfoJSON)
	wrapper[gogoproto.MessageName(&core.ClassIssuer{})] = []byte(classIssuersJSON)
	wrapper[gogoproto.MessageName(&core.Project{})] = []byte(projectInfoJSON)
	wrapper[gogoproto.MessageName(&core.Batch{})] = []byte(batchInfoJSON)
	wrapper[gogoproto.MessageName(&core.BatchBalance{})] = []byte(batchBalancesJSON)
	wrapper[gogoproto.MessageName(&core.BatchSupply{})] = []byte(batchSupplyJSON)
	wrapper[gogoproto.MessageName(&core.ClassSequence{})] = []byte(classSeqJSON)
	wrapper[gogoproto.MessageName(&core.BatchSequence{})] = []byte(batchSeqJSON)
	wrapper[gogoproto.MessageName(&core.ProjectSequence{})] = []byte(projectSeqJSON)
	wrapper[gogoproto.MessageName(&core.Params{})] = []byte(paramsJSON)

	bz, err := json.Marshal(wrapper)
	require.NoError(err)
	wrapper = map[string]json.RawMessage{}
	wrapper["ecocredit"] = bz

	_, err = s.fixture.InitGenesis(s.genesisCtx.Context, wrapper)
	require.NoError(err)

	exported := s.exportGenesisState(s.genesisCtx)
	require.NotNil(exported)

}

func (s *GenesisTestSuite) exportGenesisState(ctx types.Context) map[string]json.RawMessage {
	require := s.Require()
	exported, err := s.fixture.ExportGenesis(ctx.Context)
	require.NoError(err)

	var wrapper map[string]json.RawMessage
	err = json.Unmarshal(exported[ecocredit.ModuleName], &wrapper)
	require.NoError(err)

	return wrapper
}

type GenesisTestSuite struct {
	suite.Suite

	fixtureFactory testutil.FixtureFactory
	fixture        testutil.Fixture
	signers        []sdk.AccAddress

	paramSpace paramstypes.Subspace
	bankKeeper bankkeeper.Keeper

	genesisCtx types.Context
}

func NewGenesisTestSuite(fixtureFactory testutil.FixtureFactory, paramSpace paramstypes.Subspace, bankKeeper bankkeeper.BaseKeeper) *GenesisTestSuite {
	return &GenesisTestSuite{
		fixtureFactory: fixtureFactory,
		paramSpace:     paramSpace,
		bankKeeper:     bankKeeper,
	}
}

func (s *GenesisTestSuite) SetupSuite() {
	s.fixture = s.fixtureFactory.Setup()

	blockTime := time.Now().UTC()

	sdkCtx := s.fixture.Context().(types.Context).WithBlockTime(blockTime)
	s.genesisCtx = types.Context{Context: sdkCtx}

	s.signers = s.fixture.Signers()
	s.Require().GreaterOrEqual(len(s.signers), 8)
}
