package testsuite

import (
	"fmt"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (s *IntegrationTestSuite) TestQueryClasses() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *core.QueryClassesRequest
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
			&core.QueryClassesRequest{},
			false,
			"",
		},
	}

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
		request   *core.QueryClassInfoRequest
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
			&core.QueryClassInfoRequest{
				ClassId: "",
			},
			true,
			"not found",
		},
		{
			"credit class not found",
			&core.QueryClassInfoRequest{
				ClassId: "123",
			},
			true,
			"not found",
		},
		{
			"valid test case",
			&core.QueryClassInfoRequest{
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
		request   *core.QueryBatchesRequest
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
			"empty project id",
			&core.QueryBatchesRequest{},
			true,
			"invalid project id",
		},
		{
			"valid test case",
			&core.QueryBatchesRequest{
				ProjectId: "P01",
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
		request   *core.QueryBatchInfoRequest
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
			&core.QueryBatchInfoRequest{},
			true,
			"invalid denom",
		},
		{
			"batch not found",
			&core.QueryBatchInfoRequest{
				BatchDenom: "A00-00000000-00000000-000",
			},
			true,
			"not found",
		},
		{
			"valid request",
			&core.QueryBatchInfoRequest{
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
		request   *core.QueryBalanceRequest
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
			&core.QueryBalanceRequest{},
			true,
			"invalid denom",
		},
		{
			"with address",
			&core.QueryBalanceRequest{
				Account: s.signers[0].String(),
			},
			true,
			"invalid denom",
		},
		{
			"invalid denom",
			&core.QueryBalanceRequest{
				Account:    s.signers[0].String(),
				BatchDenom: "invalid-batch",
			},
			true,
			"invalid denom",
		},
		{
			"valid request",
			&core.QueryBalanceRequest{
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
		request   *core.QueryCreditTypesRequest
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
			&core.QueryCreditTypesRequest{},
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

func (s *IntegrationTestSuite) TestSellOrderQuery() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *marketplace.QuerySellOrderRequest
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
			&marketplace.QuerySellOrderRequest{},
			true,
			"not found",
		},
		{
			"invalid order id",
			&marketplace.QuerySellOrderRequest{
				SellOrderId: 99,
			},
			true,
			"not found",
		},
		{
			"valid request",
			&marketplace.QuerySellOrderRequest{
				SellOrderId: 1,
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.marketServer.SellOrder(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestSellOrdersQuery() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *marketplace.QuerySellOrdersRequest
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
			&marketplace.QuerySellOrdersRequest{},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.marketServer.SellOrders(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestSellOrdersByAddressQuery() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *marketplace.QuerySellOrdersByAddressRequest
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
			&marketplace.QuerySellOrdersByAddressRequest{},
			true,
			"empty address string is not allowed",
		},
		{
			"valid request",
			&marketplace.QuerySellOrdersByAddressRequest{
				Address: s.signers[3].String(),
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.marketServer.SellOrdersByAddress(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestSellOrdersByBatchDenomQuery() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *marketplace.QuerySellOrdersByBatchDenomRequest
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
			&marketplace.QuerySellOrdersByBatchDenomRequest{},
			true,
			"invalid denom",
		},
		{
			"valid request",
			&marketplace.QuerySellOrdersByBatchDenomRequest{
				BatchDenom: "A00-00000000-00000000-000",
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.marketServer.SellOrdersByBatchDenom(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestBuyOrderQuery() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *marketplace.QueryBuyOrderRequest
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
			&marketplace.QueryBuyOrderRequest{},
			true,
			"not found",
		},
		{
			"invalid order id",
			&marketplace.QueryBuyOrderRequest{
				BuyOrderId: 99,
			},
			true,
			"not found",
		},
		// TODO: filtered buy orders required #623
		//{
		//	"valid request",
		//	&ecocredit.QueryBuyOrderRequest{
		//		BuyOrderId: 1,
		//	},
		//	false,
		//	"",
		//},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.marketServer.BuyOrder(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestBuyOrdersQuery() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *marketplace.QueryBuyOrdersRequest
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
			&marketplace.QueryBuyOrdersRequest{},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.marketServer.BuyOrders(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestBuyOrdersByAddressQuery() {
	require := s.Require()

	testCases := []struct {
		name      string
		request   *marketplace.QueryBuyOrdersByAddressRequest
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
			&marketplace.QueryBuyOrdersByAddressRequest{},
			true,
			"empty address string is not allowed",
		},
		{
			"valid request",
			&marketplace.QueryBuyOrdersByAddressRequest{
				Address: s.signers[3].String(),
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.name), func() {
			_, err := s.marketServer.BuyOrdersByAddress(s.ctx, tc.request)
			if tc.expectErr {
				require.Error(err)
				require.Contains(err.Error(), tc.errMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}
