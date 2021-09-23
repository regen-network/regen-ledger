package testsuite

import (
	"fmt"

	"github.com/regen-network/regen-ledger/x/ecocredit"

	"github.com/cosmos/cosmos-sdk/types/rest"
)

func (s *IntegrationTestSuite) TestGRPCQueryParams() {
	val := s.network.Validators[0]
	require := s.Require()

	resp, err := rest.GetRequest(fmt.Sprintf("%s/regen/ecocredit/v1alpha1/params", val.APIAddress))
	require.NoError(err)

	var params ecocredit.QueryParamsResponse
	require.NoError(val.ClientCtx.Codec.UnmarshalJSON(resp, &params))

	exp := ecocredit.DefaultParams()
	s.Require().Equal(params.Params.AllowedClassCreators, exp.AllowedClassCreators)
	s.Require().Equal(params.Params.AllowlistEnabled, exp.AllowlistEnabled)
	s.Require().Equal(params.Params.CreditClassFee, exp.CreditClassFee)
	s.Require().Equal(params.Params.CreditTypes, exp.CreditTypes)
}
