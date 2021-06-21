package testsuite

import (
	"encoding/json"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s *IntegrationTestSuite) TestInitExportGenesisDefaultParams() {
	require := s.Require()
	ctx := s.genesisCtx
	cdc := s.fixture.Codec()

	// Set the param set to empty values to properly test init
	var ecocreditParams ecocredit.Params
	s.paramSpace.SetParamSet(ctx.Context, &ecocreditParams)

	genesisState := ecocredit.DefaultGenesisState()
	genesisBytes, err := cdc.MarshalJSON(genesisState)
	require.NoError(err)

	genesisData := map[string]json.RawMessage{ecocredit.ModuleName: genesisBytes}
	_, err = s.fixture.InitGenesis(ctx.Context, genesisData)
	require.NoError(err)

	expGenesisData, err := s.fixture.ExportGenesis(ctx.Context)
	require.NoError(err)

	require.Equal(string(genesisData[ecocredit.ModuleName]), string(expGenesisData[ecocredit.ModuleName]))
}

func (s *IntegrationTestSuite) TestInitExportGenesis() {
	require := s.Require()
	ctx := s.genesisCtx
	designer1 := s.signers[0].String()
	designer2 := s.signers[1].String()
	issuer1 := s.signers[2].String()
	issuer2 := s.signers[3].String()
	issuer3 := s.signers[4].String()
	// designer3 := s.signers[5]
	// addr3 := s.signers[6].String()
	cdc := s.fixture.Codec()

	// Set the param set to empty values to properly test init
	var ecocreditParams ecocredit.Params
	s.paramSpace.SetParamSet(ctx.Context, &ecocreditParams)

	genesisState := ecocredit.DefaultGenesisState()
	genesisBytes, err := cdc.MarshalJSON(genesisState)
	require.NoError(err)

	genesisData := map[string]json.RawMessage{ecocredit.ModuleName: genesisBytes}
	_, err = s.fixture.InitGenesis(ctx.Context, genesisData)
	require.NoError(err)

	exported, err := s.fixture.ExportGenesis(ctx.Context)
	require.NoError(err)

	var exportedGenesisState ecocredit.GenesisState
	err = cdc.UnmarshalJSON(exported[ecocredit.ModuleName], &exportedGenesisState)
	require.NoError(err)

	require.Equal(genesisState.Params, exportedGenesisState.Params)

	classInfos := []*ecocredit.ClassInfo{
		{
			ClassId:  "1",
			Designer: designer1,
			Issuers:  []string{issuer1, issuer2},
			Metadata: []byte("credit class metadata"),
		},
		{
			ClassId:  "2",
			Designer: designer2,
			Issuers:  []string{issuer2, issuer3},
			Metadata: []byte("credit class metadata"),
		},
	}

	batchInfos := []*ecocredit.BatchInfo{
		{
			ClassId:    "1",
			BatchDenom: "1/3",
			Issuer:     issuer1,
			TotalUnits: "100",
			Metadata:   []byte("batch metadata"),
		}, {
			ClassId:    "2",
			BatchDenom: "1/4",
			Issuer:     issuer3,
			TotalUnits: "100",
			Metadata:   []byte("batch metadata"),
		},
	}

	genesisState = &ecocredit.GenesisState{
		Params:     ecocredit.DefaultParams(),
		ClassInfos: classInfos,
		BatchInfos: batchInfos,
		IdSeq:      3,
	}

	genesisBytes, err = cdc.MarshalJSON(genesisState)
	require.NoError(err)

	genesisData = map[string]json.RawMessage{ecocredit.ModuleName: genesisBytes}
	_, err = s.fixture.InitGenesis(ctx.Context, genesisData)
	require.NoError(err)

	for _, classInfo := range classInfos {
		res, err := s.queryClient.ClassInfo(ctx, &ecocredit.QueryClassInfoRequest{
			ClassId: classInfo.ClassId,
		})
		require.NoError(err)
		s.assetClassInfoEqual(res.Info, classInfo)
	}

	for _, batchInfo := range batchInfos {
		res, err := s.queryClient.BatchInfo(ctx, &ecocredit.QueryBatchInfoRequest{
			BatchDenom: batchInfo.BatchDenom,
		})
		require.NoError(err)
		s.assetBatchInfoEqual(res.Info, batchInfo)
	}

}

func (s IntegrationTestSuite) assetClassInfoEqual(q, other *ecocredit.ClassInfo) {
	require := s.Require()
	require.Equal(q.ClassId, other.ClassId)
	require.Equal(q.Designer, other.Designer)
	require.Equal(q.Issuers, other.Issuers)
	require.Equal(q.Metadata, other.Metadata)
}

func (s IntegrationTestSuite) assetBatchInfoEqual(q, other *ecocredit.BatchInfo) {
	require := s.Require()
	require.Equal(q.ClassId, other.ClassId)
	require.Equal(q.BatchDenom, other.BatchDenom)
	require.Equal(q.Issuer, other.Issuer)
	require.Equal(q.Metadata, other.Metadata)
	require.Equal(q.TotalUnits, other.TotalUnits)
}
