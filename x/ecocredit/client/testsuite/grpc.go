package testsuite

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/testutil/rest"

	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

const baseRoute = "regen/ecocredit/v1"

func (s *IntegrationTestSuite) TestQueryClasses() {
	testCases := []struct {
		name      string
		url       string
		paginated bool
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/classes", s.val.APIAddress, baseRoute),
			false,
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/classes?pagination.limit=1", s.val.APIAddress, baseRoute),
			true,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res types.QueryClassesResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Classes)
			require.True(len(res.Classes) > 0)
			if tc.paginated {
				require.NotNil(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryClass() {
	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/class/%s", s.val.APIAddress, baseRoute, s.classID),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/classes/%s", s.val.APIAddress, baseRoute, s.classID),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res types.QueryClassResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Class)
			require.Equal(res.Class.Id, s.classID)
		})
	}
}

func (s *IntegrationTestSuite) TestQueryProject() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/project/%s", s.val.APIAddress, baseRoute, s.projectID),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/projects/%s", s.val.APIAddress, baseRoute, s.projectID),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res types.QueryProjectResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Project)
		})
	}
}

func (s *IntegrationTestSuite) TestQueryProjects() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/projects", s.val.APIAddress, baseRoute),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/projects?pagination.limit=1&pagination.countTotal=true",
				s.val.APIAddress,
				baseRoute,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res types.QueryProjectsResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Projects)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Projects, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryProjectsByClass() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/projects-by-class/%s", s.val.APIAddress, baseRoute, s.classID),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/projects-by-class/%s?pagination.limit=1&pagination.countTotal=true",
				s.val.APIAddress,
				baseRoute,
				s.classID,
			),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/projects/class/%s", s.val.APIAddress, baseRoute, s.classID),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/classes/%s/projects", s.val.APIAddress, baseRoute, s.classID),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res types.QueryProjectsByClassResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Projects)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Projects, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryProjectsByReferenceID() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf(
				"%s/%s/projects-by-reference-id/%s",
				s.val.APIAddress,
				baseRoute,
				s.projectReferenceID,
			),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/projects-by-reference-id/%s?pagination.limit=1&pagination.countTotal=true",
				s.val.APIAddress,
				baseRoute,
				s.projectReferenceID,
			),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/projects/reference-id/%s",
				s.val.APIAddress,
				baseRoute,
				s.projectReferenceID,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res types.QueryProjectsByReferenceIdResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Projects)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Projects, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatches() {
	testCases := []struct {
		name      string
		url       string
		paginated bool
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/batches", s.val.APIAddress, baseRoute),
			false,
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/batches?pagination.limit=2", s.val.APIAddress, baseRoute),
			true,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res types.QueryBatchesResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Batches)
			require.Greater(len(res.Batches), 0)
			if tc.paginated {
				require.NotNil(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchesByIssuer() {
	testCases := []struct {
		name      string
		url       string
		paginated bool
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/batches-by-issuer/%s", s.val.APIAddress, baseRoute, s.addr1),
			false,
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/batches-by-issuer/%s?pagination.limit=2", s.val.APIAddress, baseRoute, s.addr1),
			true,
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/batches/issuer/%s", s.val.APIAddress, baseRoute, s.addr1),
			false,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res types.QueryBatchesByIssuerResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Batches)
			require.Greater(len(res.Batches), 0)
			if tc.paginated {
				require.NotNil(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchesByClass() {
	testCases := []struct {
		name      string
		url       string
		paginated bool
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/batches-by-class/%s", s.val.APIAddress, baseRoute, s.classID),
			false,
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/batches-by-class/%s?pagination.limit=2", s.val.APIAddress, baseRoute, s.classID),
			true,
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/batches/class/%s", s.val.APIAddress, baseRoute, s.classID),
			false,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res types.QueryBatchesByClassResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Batches)
			require.Greater(len(res.Batches), 0)
			if tc.paginated {
				require.NotNil(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchesByProject() {
	testCases := []struct {
		name      string
		url       string
		paginated bool
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/batches-by-project/%s", s.val.APIAddress, baseRoute, s.projectID),
			false,
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/batches-by-project/%s?pagination.limit=2", s.val.APIAddress, baseRoute, s.projectID),
			true,
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/batches/project/%s", s.val.APIAddress, baseRoute, s.projectID),
			false,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res types.QueryBatchesResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Batches)
			require.Greater(len(res.Batches), 0)
			if tc.paginated {
				require.NotNil(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatch() {
	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/batch/%s", s.val.APIAddress, baseRoute, s.batchDenom),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/batches/%s", s.val.APIAddress, baseRoute, s.batchDenom),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res types.QueryBatchResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Batch)
			require.Equal(res.Batch.Denom, s.batchDenom)
		})
	}
}

func (s *IntegrationTestSuite) TestCreditTypes() {
	require := s.Require()

	url := fmt.Sprintf("%s/%s/credit-types", s.val.APIAddress, baseRoute)
	resp, err := rest.GetRequest(url)
	require.NoError(err)

	var res types.QueryCreditTypesResponse
	err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
	require.NoError(err)
	require.Greater(len(res.CreditTypes), 0)
}

func (s *IntegrationTestSuite) TestQueryBalance() {
	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/balance/%s/%s", s.val.APIAddress, baseRoute, s.batchDenom, s.addr1),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/batches/%s/balance/%s", s.val.APIAddress, baseRoute, s.batchDenom, s.addr1),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/balances/%s/batch/%s", s.val.APIAddress, baseRoute, s.addr1, s.batchDenom),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res types.QueryBalanceResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res)
			require.NotEmpty(res.Balance.TradableAmount)
			require.NotEmpty(res.Balance.RetiredAmount)
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySupply() {
	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/supply/%s", s.val.APIAddress, baseRoute, s.batchDenom),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/batches/%s/supply", s.val.APIAddress, baseRoute, s.batchDenom),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res types.QuerySupplyResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res)
			require.NotEmpty(res.RetiredAmount)
			require.NotEmpty(res.TradableAmount)
		})
	}
}

// TODO: #1363
// func (s *IntegrationTestSuite) TestQueryParams() {
// 	require := s.Require()

// 	resp, err := rest.GetRequest(fmt.Sprintf("%s/%s/params", s.val.APIAddress, baseRoute))
// 	require.NoError(err)

// 	var res types.QueryParamsResponse
// 	require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res))
// 	s.Require().Equal(types.DefaultParams(), *res.Params)
// }

func (s *IntegrationTestSuite) TestCreditType() {
	require := s.Require()

	url := fmt.Sprintf("%s/%s/credit-types/%s", s.val.APIAddress, baseRoute, "C")
	resp, err := rest.GetRequest(url)
	require.NoError(err)

	var res types.QueryCreditTypeResponse
	err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
	require.NoError(err)
	require.Equal("C", res.CreditType.Abbreviation)
	require.Equal(uint32(6), res.CreditType.Precision)
}

func (s *IntegrationTestSuite) TestClassCreatorAllowlist() {
	require := s.Require()

	url := fmt.Sprintf("%s/%s/class-creator-allowlist", s.val.APIAddress, baseRoute)
	resp, err := rest.GetRequest(url)
	require.NoError(err)

	var res types.QueryClassCreatorAllowlistResponse
	err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
	require.NoError(err)
	require.Equal(false, res.Enabled)
}

func (s *IntegrationTestSuite) TestAllBalances() {
	require := s.Require()

	url := fmt.Sprintf("%s/%s/all-balances?pagination.countTotal=true", s.val.APIAddress, baseRoute)
	resp, err := rest.GetRequest(url)
	require.NoError(err)

	var res types.QueryAllBalancesResponse
	err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
	require.NoError(err)
	require.NotEmpty(res.Balances)
	require.NotZero(res.Pagination.Total)

	url = fmt.Sprintf("%s/%s/balances?pagination.countTotal=true", s.val.APIAddress, baseRoute)
	resp, err = rest.GetRequest(url)
	require.NoError(err)

	res = types.QueryAllBalancesResponse{}
	err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
	require.NoError(err)
	require.NotEmpty(res.Balances)
	require.NotZero(res.Pagination.Total)
}

func (s *IntegrationTestSuite) TestBalancesByBatch() {
	require := s.Require()

	checkQuery := func(url string) {
		resp, err := rest.GetRequest(url)
		require.NoError(err)
		var res types.QueryBalancesByBatchResponse
		err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
		require.NoError(err)
		require.NotEmpty(res.Balances)
		require.NotZero(res.Pagination.Total)
	}

	url := fmt.Sprintf("%s/%s/balances-by-batch/%s?pagination.countTotal=true", s.val.APIAddress, baseRoute, s.batchDenom)
	checkQuery(url)

	url = fmt.Sprintf("%s/%s/batches/%s/balances?pagination.countTotal=true", s.val.APIAddress, baseRoute, s.batchDenom)
	checkQuery(url)
}
