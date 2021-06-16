package testsuite

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s *IntegrationTestSuite) TestInitExportGenesis() {
	require := s.Require()
	ctx := s.sdkCtx

	// Export the default param set
	exported, err := s.fixture.ExportGenesis(ctx)
	require.NoError(err)

	// Set the param set to empty values to properly test init
	var ecocreditParams ecocredit.Params
	s.paramSpace.SetParamSet(ctx, &ecocreditParams)

	// Init the genesis data from the exported data
	_, err = s.fixture.InitGenesis(ctx, exported)
	require.NoError(err)

	// Check that the param set is correctly updated to the default params
	s.paramSpace.GetParamSet(ctx, &ecocreditParams)
	require.NoError(err)
	require.Equal(ecocredit.DefaultParams().CreditClassFee, ecocreditParams.CreditClassFee)
}
