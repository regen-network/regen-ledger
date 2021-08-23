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
			"key must not be nil",
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
				ClassId: "4",
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
			"class id cannot be empty",
		},
		{
			"valid test case",
			&ecocredit.QueryBatchesRequest{
				ClassId: "4",
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
			"key must not be nil",
		},
		{
			"batch not found",
			&ecocredit.QueryBatchInfoRequest{
				BatchDenom: "invalid-batch",
			},
			true,
			"not found",
		},
		{
			"valid testcase",
			&ecocredit.QueryBatchInfoRequest{
				BatchDenom: "4",
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
			"empty address string is not allowed",
		},
		{
			"with address",
			&ecocredit.QueryBalanceRequest{
				Account: s.signers[0].String(),
			},
			false,
			"",
		},
		{
			"invalid denom",
			&ecocredit.QueryBalanceRequest{
				Account:    s.signers[0].String(),
				BatchDenom: "invalid-batch",
			},
			false,
			"",
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
