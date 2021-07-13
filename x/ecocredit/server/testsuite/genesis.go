package testsuite

import (
	"encoding/json"

	"github.com/regen-network/regen-ledger/types"
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

	tradableBalances := []*ecocredit.Balance{
		{
			Address:    addr1,
			BatchDenom: "4/6",
			Balance:    "90.003",
		},
	}

	retiredBalances := []*ecocredit.Balance{
		{
			Address:    addr1,
			BatchDenom: "4/6",
			Balance:    "9.997",
		},
	}

	tradableSupplies := []*ecocredit.Supply{
		{
			BatchDenom: "4/6",
			Supply:     "90.003",
		},
	}

	retiredSupplies := []*ecocredit.Supply{
		{
			BatchDenom: "4/6",
			Supply:     "9.997",
		},
	}

	precisions := []*ecocredit.Precision{
		{
			BatchDenom:       "4/6",
			MaxDecimalPlaces: 5,
		},
	}

	genesisState := &ecocredit.GenesisState{
		Params:           ecocredit.DefaultParams(),
		IdSeq:            7,
		ClassInfo:        classInfo,
		BatchInfo:        batchInfo,
		Precisions:       precisions,
		TradableBalances: tradableBalances,
		RetiredBalances:  retiredBalances,
		TradableSupplies: tradableSupplies,
		RetiredSupplies:  retiredSupplies,
	}
	s.initGenesisState(ctx, genesisState)

	exportedGenesisState := s.exportGenesisState(ctx)
	require.Equal(genesisState.Params, exportedGenesisState.Params)
	require.Equal(genesisState.IdSeq, exportedGenesisState.IdSeq)

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

	for i, tradableBalance := range tradableBalances {
		res, err := s.queryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{
			Account:    tradableBalance.Address,
			BatchDenom: tradableBalance.BatchDenom,
		})
		require.NoError(err)
		require.NotNil(res)
		s.assertTradableBalanceEqual(res, tradableBalance)
		s.assertRetiredBalanceEqual(res, *retiredBalances[i])
	}

	for i, supply := range tradableSupplies {
		res, err := s.queryClient.Supply(ctx, &ecocredit.QuerySupplyRequest{
			BatchDenom: supply.BatchDenom,
		})
		require.NoError(err)
		require.NotNil(res)
		require.Equal(res.TradableSupply, supply.Supply)
		require.Equal(res.RetiredSupply, retiredSupplies[i].Supply)
	}

	for _, precision := range precisions {
		res, err := s.queryClient.Precision(ctx, &ecocredit.QueryPrecisionRequest{
			BatchDenom: precision.BatchDenom,
		})
		require.NoError(err)
		require.NotNil(res)
		require.Equal(precision.MaxDecimalPlaces, res.MaxDecimalPlaces)
	}

	exported := s.exportGenesisState(ctx)
	require.Equal(uint64(7), exported.IdSeq)
	require.Equal(genesisState.Params, exported.Params)
	require.Equal(genesisState.ClassInfo, exported.ClassInfo)
	require.Equal(genesisState.BatchInfo, exported.BatchInfo)
	require.Equal(genesisState.RetiredBalances, exported.RetiredBalances)
	require.Equal(genesisState.TradableBalances, exported.TradableBalances)
	require.Equal(genesisState.TradableSupplies, exported.TradableSupplies)
	require.Equal(genesisState.RetiredSupplies, exported.RetiredSupplies)
	require.Equal(genesisState.Precisions, exported.Precisions)
}

func (s *IntegrationTestSuite) assertTradableBalanceEqual(res *ecocredit.QueryBalanceResponse, balance *ecocredit.Balance) {
	require := s.Require()
	require.Equal(balance.Balance, res.TradableAmount)
}

func (s *IntegrationTestSuite) assertRetiredBalanceEqual(res *ecocredit.QueryBalanceResponse, balance ecocredit.Balance) {
	require := s.Require()
	require.Equal(balance.Balance, res.RetiredAmount)
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

func (s *IntegrationTestSuite) initGenesisState(ctx types.Context, genesisState *ecocredit.GenesisState) {
	cdc := s.fixture.Codec()
	require := s.Require()
	genesisBytes, err := cdc.MarshalJSON(genesisState)
	require.NoError(err)

	genesisData := map[string]json.RawMessage{ecocredit.ModuleName: genesisBytes}
	_, err = s.fixture.InitGenesis(ctx.Context, genesisData)
	require.NoError(err)
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
