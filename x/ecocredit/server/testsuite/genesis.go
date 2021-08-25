package testsuite

import (
	"encoding/json"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s *IntegrationTestSuite) TestInitExportGenesis() {
	require := s.Require()
	ctx := s.genesisCtx
	designer1 := s.signers[0]
	designer2 := s.signers[1].String()
	issuer1 := s.signers[2].String()
	issuer2 := s.signers[3].String()
	addr1 := s.signers[4].String()

	// Set the param set to empty values to properly test init
	var ecocreditParams ecocredit.Params
	s.paramSpace.SetParamSet(ctx.Context, &ecocreditParams)

	classInfo := []*ecocredit.ClassInfo{
		{
			ClassId:  "4",
			Designer: designer1.String(),
			Issuers:  []string{issuer1, issuer2},
			Metadata: []byte("credit class metadata"),
		},
		{
			ClassId:  "5",
			Designer: designer2,
			Issuers:  []string{issuer2, addr1},
			Metadata: []byte("credit class metadata"),
		},
	}

	batchInfo := []*ecocredit.BatchInfo{
		{
			ClassId:     "4",
			BatchDenom:  "4/6",
			Issuer:      issuer1,
			TotalAmount: "100",
			Metadata:    []byte("batch metadata"),
		}, {
			ClassId:     "5",
			BatchDenom:  "5/7",
			Issuer:      addr1,
			TotalAmount: "100",
			Metadata:    []byte("batch metadata"),
		},
	}

	balances := []*ecocredit.Balance{
		{
			Address:         addr1,
			BatchDenom:      "4/6",
			TradableBalance: "90.003",
			RetiredBalance:  "9.997",
		},
	}

	supplies := []*ecocredit.Supply{
		{
			BatchDenom:     "4/6",
			TradableSupply: "90.003",
			RetiredSupply:  "9.997",
		},
	}

	sequences := []*ecocredit.CreditTypeSeq{
		{
			Abbreviation: "BIO",
			SeqNumber:    0,
		},
	}

	genesisState := &ecocredit.GenesisState{
		Params:    ecocredit.DefaultParams(),
		Sequences: sequences,
		ClassInfo: classInfo,
		BatchInfo: batchInfo,
		Balances:  balances,
		Supplies:  supplies,
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

	// invalid supply
	genesisState.Supplies = []*ecocredit.Supply{
		{
			BatchDenom:     "4/6",
			TradableSupply: "101.000",
		},
	}

	err := s.initGenesisState(ctx, genesisState)
	require.Error(err)
	require.Contains(err.Error(), "supply is incorrect for 4/6 credit batch")

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
	require.Equal(q.Designer, other.Designer)
	require.Equal(q.Issuers, other.Issuers)
	require.Equal(q.Metadata, other.Metadata)
}

func (s *IntegrationTestSuite) assetBatchInfoEqual(q, other *ecocredit.BatchInfo) {
	require := s.Require()
	require.Equal(q.ClassId, other.ClassId)
	require.Equal(q.BatchDenom, other.BatchDenom)
	require.Equal(q.Issuer, other.Issuer)
	require.Equal(q.Metadata, other.Metadata)
	require.Equal(q.TotalAmount, other.TotalAmount)
}
