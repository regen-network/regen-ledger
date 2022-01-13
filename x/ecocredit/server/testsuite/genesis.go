package testsuite

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/stretchr/testify/suite"
)

func (s *IntegrationTestSuite) TestInitExportGenesis() {
	require := s.Require()
	ctx := s.genesisCtx
	admin1 := s.signers[0]
	admin2 := s.signers[1].String()
	issuer1 := s.signers[2].String()
	issuer2 := s.signers[3].String()
	addr1 := s.signers[4].String()

	// Set the param set to empty values to properly test init
	var ecocreditParams ecocredit.Params
	s.paramSpace.SetParamSet(ctx.Context, &ecocreditParams)

	classInfo := []*ecocredit.ClassInfo{
		{
			ClassId:  "BIO01",
			Admin:    admin1.String(),
			Issuers:  []string{issuer1, issuer2},
			Metadata: []byte("credit class metadata"),
		},
		{
			ClassId:  "BIO02",
			Admin:    admin2,
			Issuers:  []string{issuer2, addr1},
			Metadata: []byte("credit class metadata"),
		},
	}

	projectInfo := []*ecocredit.ProjectInfo{
		{
			ProjectId:       "P01",
			ClassId:         "BIO01",
			Issuer:          issuer1,
			ProjectLocation: "AQ",
			Metadata:        []byte("project metadata"),
		},
		{
			ProjectId:       "P02",
			ClassId:         "BIO02",
			Issuer:          issuer2,
			ProjectLocation: "AQ",
			Metadata:        []byte("project metadata"),
		},
	}

	batchInfo := []*ecocredit.BatchInfo{
		{
			ProjectId:   "P01",
			BatchDenom:  "BIO01-00000000-00000000-001",
			TotalAmount: "100",
			Metadata:    []byte("batch metadata"),
		}, {
			ProjectId:   "P02",
			BatchDenom:  "BIO02-00000000-00000000-001",
			TotalAmount: "100",
			Metadata:    []byte("batch metadata"),
		},
	}

	balances := []*ecocredit.Balance{
		{
			Address:         addr1,
			BatchDenom:      "BIO01-00000000-00000000-001",
			TradableBalance: "90.003",
			RetiredBalance:  "9.997",
		},
	}

	supplies := []*ecocredit.Supply{
		{
			BatchDenom:     "BIO01-00000000-00000000-001",
			TradableSupply: "90.003",
			RetiredSupply:  "9.997",
		},
	}

	sequences := []*ecocredit.CreditTypeSeq{
		{
			Abbreviation: "BIO",
			SeqNumber:    3,
		},
	}

	genesisState := &ecocredit.GenesisState{
		Params:        ecocredit.DefaultParams(),
		Sequences:     sequences,
		ClassInfo:     classInfo,
		BatchInfo:     batchInfo,
		Balances:      balances,
		Supplies:      supplies,
		ProjectInfo:   projectInfo,
		ProjectSeqNum: 2,
	}
	require.NoError(s.initGenesisState(ctx, genesisState))

	exportedGenesisState := s.exportGenesisState(ctx)
	require.Equal(genesisState.Params, exportedGenesisState.Params)
	require.Equal(genesisState.Sequences, exportedGenesisState.Sequences)

	for _, info := range classInfo {
		res, err := s.queryClient.ClassInfo(ctx, &ecocredit.QueryClassInfoRequest{
			ClassId: info.ClassId,
		})
		require.NoError(err)
		s.assetClassInfoEqual(res.Info, info)
	}

	for _, info := range projectInfo {
		res, err := s.queryClient.ProjectInfo(ctx, &ecocredit.QueryProjectInfoRequest{
			ProjectId: info.ProjectId,
		})
		require.NoError(err)
		s.assetProjectInfoEqual(res.Info, info)
	}

	for _, info := range batchInfo {
		res, err := s.queryClient.BatchInfo(ctx, &ecocredit.QueryBatchInfoRequest{
			BatchDenom: info.BatchDenom,
		})
		require.NoError(err)
		s.assetBatchInfoEqual(res.Info, info)
	}

	for _, balance := range balances {
		res, err := s.queryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{
			Account:    balance.Address,
			BatchDenom: balance.BatchDenom,
		})
		require.NoError(err)
		require.NotNil(res)

		require.Equal(res.TradableAmount, balance.TradableBalance)
		require.Equal(res.RetiredAmount, balance.RetiredBalance)
	}

	for _, supply := range supplies {
		res, err := s.queryClient.Supply(ctx, &ecocredit.QuerySupplyRequest{
			BatchDenom: supply.BatchDenom,
		})
		require.NoError(err)
		require.NotNil(res)
		tSupply, err := math.NewNonNegativeDecFromString(res.TradableSupply)
		require.NoError(err)
		rSupply, err := math.NewNonNegativeDecFromString(res.RetiredSupply)
		require.NoError(err)
		require.Equal(tSupply.String(), supply.TradableSupply)
		require.Equal(rSupply.String(), supply.RetiredSupply)
	}

	exported := s.exportGenesisState(ctx)
	require.Equal(genesisState.Sequences, exportedGenesisState.Sequences)
	require.Equal(genesisState.Params, exported.Params)
	require.Equal(genesisState.ClassInfo, exported.ClassInfo)
	require.Equal(genesisState.BatchInfo, exported.BatchInfo)
	require.Equal(genesisState.Balances, exported.Balances)
	require.Equal(genesisState.Supplies, exported.Supplies)

}

func (s *IntegrationTestSuite) exportGenesisState(ctx types.Context) ecocredit.GenesisState {
	require := s.Require()
	cdc := s.fixture.Codec()
	exported, err := s.fixture.ExportGenesis(ctx.Context)
	require.NoError(err)

	var exportedGenesisState ecocredit.GenesisState
	err = cdc.UnmarshalJSON(exported[ecocredit.ModuleName], &exportedGenesisState)
	require.NoError(err)

	return exportedGenesisState
}

func (s *IntegrationTestSuite) initGenesisState(ctx types.Context, genesisState *ecocredit.GenesisState) error {
	cdc := s.fixture.Codec()
	require := s.Require()
	genesisBytes, err := cdc.MarshalJSON(genesisState)
	require.NoError(err)

	genesisData := map[string]json.RawMessage{ecocredit.ModuleName: genesisBytes}
	_, err = s.fixture.InitGenesis(ctx.Context, genesisData)
	return err
}

func (s *IntegrationTestSuite) assetClassInfoEqual(q, other *ecocredit.ClassInfo) {
	require := s.Require()
	require.Equal(q.ClassId, other.ClassId)
	require.Equal(q.Admin, other.Admin)
	require.Equal(q.Issuers, other.Issuers)
	require.Equal(q.Metadata, other.Metadata)
}

func (s *IntegrationTestSuite) assetProjectInfoEqual(q, other *ecocredit.ProjectInfo) {
	require := s.Require()
	require.Equal(q, other)
}

func (s *IntegrationTestSuite) assetBatchInfoEqual(q, other *ecocredit.BatchInfo) {
	require := s.Require()
	require.Equal(q.ProjectId, other.ProjectId)
	require.Equal(q.BatchDenom, other.BatchDenom)
	require.Equal(q.Metadata, other.Metadata)
	require.Equal(q.TotalAmount, other.TotalAmount)
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

func (s *GenesisTestSuite) TestInvalidGenesis() {
	require := s.Require()

	ctx := s.genesisCtx
	admin1 := s.signers[0]
	admin2 := s.signers[1].String()
	issuer1 := s.signers[2].String()
	issuer2 := s.signers[3].String()
	addr1 := s.signers[4].String()

	// Set the param set to empty values to properly test init
	var ecocreditParams ecocredit.Params
	s.paramSpace.SetParamSet(ctx.Context, &ecocreditParams)

	classInfo := []*ecocredit.ClassInfo{
		{
			ClassId:  "BIO01",
			Admin:    admin1.String(),
			Issuers:  []string{issuer1, issuer2},
			Metadata: []byte("credit class metadata"),
		},
		{
			ClassId:  "BIO02",
			Admin:    admin2,
			Issuers:  []string{issuer2, addr1},
			Metadata: []byte("credit class metadata"),
		},
	}

	projectInfo := []*ecocredit.ProjectInfo{
		{
			ProjectId:       "P01",
			ClassId:         "BIO01",
			Issuer:          issuer1,
			ProjectLocation: "AQ",
			Metadata:        []byte("project metadata"),
		},
		{
			ProjectId:       "P02",
			ClassId:         "BIO02",
			Issuer:          issuer2,
			ProjectLocation: "AQ",
			Metadata:        []byte("project metadata"),
		},
	}

	batchInfo := []*ecocredit.BatchInfo{
		{
			ProjectId:   "P01",
			BatchDenom:  "BIO01-00000000-00000000-001",
			TotalAmount: "100",
			Metadata:    []byte("batch metadata"),
		}, {
			ProjectId:   "P02",
			BatchDenom:  "BIO02-00000000-00000000-001",
			TotalAmount: "100",
			Metadata:    []byte("batch metadata"),
		},
	}

	balances := []*ecocredit.Balance{
		{
			Address:         addr1,
			BatchDenom:      "BIO01-00000000-00000000-001",
			TradableBalance: "90.003",
			RetiredBalance:  "9.997",
		},
	}

	supplies := []*ecocredit.Supply{
		{
			BatchDenom:     "BIO01-00000000-00000000-001",
			TradableSupply: "101.000",
			RetiredSupply:  "9.997",
		},
	}

	sequences := []*ecocredit.CreditTypeSeq{
		{
			Abbreviation: "BIO",
			SeqNumber:    3,
		},
	}

	genesisState := &ecocredit.GenesisState{
		Params:        ecocredit.DefaultParams(),
		Sequences:     sequences,
		ClassInfo:     classInfo,
		BatchInfo:     batchInfo,
		Balances:      balances,
		Supplies:      supplies,
		ProjectInfo:   projectInfo,
		ProjectSeqNum: 2,
	}
	cdc := s.fixture.Codec()
	genesisBytes, err := cdc.MarshalJSON(genesisState)
	require.NoError(err)

	genesisData := map[string]json.RawMessage{ecocredit.ModuleName: genesisBytes}
	_, err = s.fixture.InitGenesis(ctx.Context, genesisData)

	require.Error(err)
	require.Contains(err.Error(), "supply is incorrect for BIO01-00000000-00000000-001 credit batch")

}
