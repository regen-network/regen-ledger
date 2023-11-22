package testsuite

import (
	"encoding/json"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"

	basev1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/testutil/fixture"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
)

func (s *GenesisTestSuite) TestInitExportGenesis() {
	require := s.Require()

	classIssuers := []baseapi.ClassIssuer{
		{ClassKey: 1, Issuer: sdk.AccAddress("addr1")},
		{ClassKey: 1, Issuer: sdk.AccAddress("addr2")},
		{ClassKey: 2, Issuer: sdk.AccAddress("addr2")},
		{ClassKey: 2, Issuer: sdk.AccAddress("addr3")},
	}
	classIssuersJSON, err := json.Marshal(classIssuers)
	require.NoError(err)

	classes := []baseapi.Class{
		{Id: "BIO001", Admin: sdk.AccAddress("addr1"), Metadata: "metadata", CreditTypeAbbrev: "BIO"},
		{Id: "BIO002", Admin: sdk.AccAddress("addr2"), Metadata: "metadata", CreditTypeAbbrev: "BIO"},
	}
	classJSON, err := json.Marshal(classes)
	require.NoError(err)

	projects := []baseapi.Project{
		{Id: "C01-001", Admin: sdk.AccAddress("addr1"), ClassKey: 1, Jurisdiction: "AQ", Metadata: "metadata"},
		{Id: "C01-002", Admin: sdk.AccAddress("addr2"), ClassKey: 2, Jurisdiction: "AQ", Metadata: "metadata"},
	}
	projectJSON, err := json.Marshal(projects)
	require.NoError(err)

	batches := []baseapi.Batch{
		{Issuer: sdk.AccAddress("addr1"), ProjectKey: 1, Denom: "BIO01-00000000-00000000-001", Metadata: "metadata"},
		{Issuer: nil, ProjectKey: 1, Denom: "BIO02-0000000-0000000-001", Metadata: "metadata"},
	}
	batchJSON, err := json.Marshal(batches)
	require.NoError(err)

	balances := []baseapi.BatchBalance{
		{Address: sdk.AccAddress("addr1"), BatchKey: 1, TradableAmount: "90.003", RetiredAmount: "9.997", EscrowedAmount: ""},
	}
	batchBalancesJSON, err := json.Marshal(balances)
	require.NoError(err)

	supply := []baseapi.BatchSupply{
		{BatchKey: 1, TradableAmount: "90.003", RetiredAmount: "9.997", CancelledAmount: ""},
	}
	batchSupplyJSON, err := json.Marshal(supply)
	require.NoError(err)

	classSeq := []baseapi.ClassSequence{{CreditTypeAbbrev: "BIO", NextSequence: 3}}
	classSeqJSON, err := json.Marshal(classSeq)
	require.NoError(err)

	batchSeq := []baseapi.BatchSequence{{ProjectKey: 1, NextSequence: 3}}
	batchSeqJSON, err := json.Marshal(batchSeq)
	require.NoError(err)

	projectSeq := []baseapi.ProjectSequence{{ClassKey: 1, NextSequence: 3}}
	projectSeqJSON, err := json.Marshal(projectSeq)
	require.NoError(err)

	classAllowlistSettingJSON, err := json.Marshal(baseapi.ClassCreatorAllowlist{
		Enabled: true,
	})
	require.NoError(err)

	allowedClassCreatorsJSON, err := json.Marshal([]baseapi.AllowedClassCreator{
		{
			Address: sdk.AccAddress("creator1"),
		},
		{
			Address: sdk.AccAddress("creator2"),
		},
	})
	require.NoError(err)

	classFeeJSON, err := json.Marshal(baseapi.ClassFee{
		Fee: &basev1beta1.Coin{
			Denom:  "stake",
			Amount: "1000",
		},
	})
	require.NoError(err)

	wrapper := map[string]json.RawMessage{}
	wrapper[string(proto.MessageName(&baseapi.Class{}))] = classJSON
	wrapper[string(proto.MessageName(&baseapi.ClassIssuer{}))] = classIssuersJSON
	wrapper[string(proto.MessageName(&baseapi.Project{}))] = projectJSON
	wrapper[string(proto.MessageName(&baseapi.Batch{}))] = batchJSON
	wrapper[string(proto.MessageName(&baseapi.BatchBalance{}))] = batchBalancesJSON
	wrapper[string(proto.MessageName(&baseapi.BatchSupply{}))] = batchSupplyJSON
	wrapper[string(proto.MessageName(&baseapi.ClassSequence{}))] = classSeqJSON
	wrapper[string(proto.MessageName(&baseapi.BatchSequence{}))] = batchSeqJSON
	wrapper[string(proto.MessageName(&baseapi.ProjectSequence{}))] = projectSeqJSON
	wrapper[string(proto.MessageName(&baseapi.ClassCreatorAllowlist{}))] = classAllowlistSettingJSON
	wrapper[string(proto.MessageName(&baseapi.AllowedClassCreator{}))] = allowedClassCreatorsJSON
	wrapper[string(proto.MessageName(&baseapi.ClassFee{}))] = classFeeJSON

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

	fixtureFactory fixture.Factory
	fixture        fixture.Fixture
	signers        []sdk.AccAddress

	bankKeeper bankkeeper.Keeper

	genesisCtx sdk.Context
}

func NewGenesisTestSuite(fixtureFactory fixture.Factory, bankKeeper bankkeeper.BaseKeeper) *GenesisTestSuite {
	return &GenesisTestSuite{
		fixtureFactory: fixtureFactory,
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
