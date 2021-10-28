package testsuite

import (
	"fmt"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s *IntegrationTestSuite) TestQueryClasses() {

	testCases := []struct {
		name      string
		request   *ecocredit.QueryClassesRequest
		expectErr bool
		errMsg    string
	}{
		{
			"nil request",
			nil,
			true,
			"empty request",
		},
		{
			"valid request",
			&ecocredit.QueryClassesRequest{},
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.queryClient.Classes(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryClassInfo() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *ecocredit.QueryClassInfoRequest
		expectErr bool
		errMsg    string
	}{
		{
			"nil request",
			nil,
			true,
			"empty request",
		},
		{
			"empty class-ID",
			&ecocredit.QueryClassInfoRequest{
				ClassId: "",
			},
			true,
			"not found",
		},
		{
			"credit class not found",
			&ecocredit.QueryClassInfoRequest{
				ClassId: "123",
			},
			true,
			"not found",
		},
		{
			"valid test case",
			&ecocredit.QueryClassInfoRequest{
				ClassId: "BIO01",
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.queryClient.ClassInfo(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatches() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *ecocredit.QueryBatchesRequest
		expectErr bool
		errMsg    string
	}{
		{
			"nil request",
			nil,
			true,
			"empty request",
		},
		{
			"empty class id",
			&ecocredit.QueryBatchesRequest{},
			true,
			"class ID didn't match the format",
		},
		{
			"valid test case",
			&ecocredit.QueryBatchesRequest{
				ClassId: "C04",
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.queryClient.Batches(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchInfo() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *ecocredit.QueryBatchInfoRequest
		expectErr bool
		errMsg    string
	}{
		{
			"nil request",
			nil,
			true,
			"empty request",
		},
		{
			"empty batch denom",
			&ecocredit.QueryBatchInfoRequest{},
			true,
			"invalid denom",
		},
		{
			"batch not found",
			&ecocredit.QueryBatchInfoRequest{
				BatchDenom: "A00-00000000-00000000-000",
			},
			true,
			"not found",
		},
		{
			"valid testcase",
			&ecocredit.QueryBatchInfoRequest{
				BatchDenom: "BIO01-00000000-00000000-001",
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.queryClient.BatchInfo(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestBalanceQuery() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *ecocredit.QueryBalanceRequest
		expectErr bool
		errMsg    string
	}{
		{
			"nil request",
			nil,
			true,
			"empty request",
		},
		{
			"empty request",
			&ecocredit.QueryBalanceRequest{},
			true,
			"invalid denom",
		},
		{
			"with address",
			&ecocredit.QueryBalanceRequest{
				Account: s.signers[0].String(),
			},
			true,
			"invalid denom",
		},
		{
			"invalid denom",
			&ecocredit.QueryBalanceRequest{
				Account:    s.signers[0].String(),
				BatchDenom: "invalid-batch",
			},
			true,
			"invalid denom",
		},
		{
			"valid testcase",
			&ecocredit.QueryBalanceRequest{
				BatchDenom: "C01-20210823-20210823-001",
				Account:    s.signers[3].String(),
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.queryClient.Balance(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCreditTypeQuery() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *ecocredit.QueryCreditTypesRequest
		expectErr bool
		errMsg    string
	}{
		{
			"nil request",
			nil,
			false,
			"",
		},
		{
			"empty request",
			&ecocredit.QueryCreditTypesRequest{},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.queryClient.CreditTypes(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}
