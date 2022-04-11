package testsuite

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s *IntegrationTestSuite) TestGetClasses() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		expItems int
	}{
		{
			"invalid path",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/class", val.APIAddress),
			true,
			0,
		},
		{
			"valid query",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/classes", val.APIAddress),
			false,
			4,
		},
		{
			"valid query pagination",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/classes?pagination.limit=2", val.APIAddress),
			false,
			2,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var classes ecocredit.QueryClassesResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &classes)

			if tc.expErr {
				require.Error(err)
			} else {
				require.NoError(err)
				require.NotNil(classes.Classes)
				require.Len(classes.Classes, tc.expItems)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetClass() {
	val := s.network.Validators[0]

	testCases := []struct {
		name    string
		url     string
		expErr  bool
		errMsg  string
		classID string
	}{
		{
			"invalid path",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/class", val.APIAddress),
			true,
			"Not Implemented",
			"",
		},
		{
			"class not found",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/classes/%s", val.APIAddress, "C999"),
			true,
			"not found",
			"",
		},
		{
			"valid class-id",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/classes/%s", val.APIAddress, "C01"),
			false,
			"",
			"C01",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var class ecocredit.QueryClassInfoResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &class)

			if tc.expErr {
				require.Error(err)
			} else {
				require.NoError(err)
				require.NotNil(class.Info)
				require.Contains(class.Info.ClassId, tc.classID)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetBatches() {
	val := s.network.Validators[0]

	testCases := []struct {
		name       string
		url        string
		numBatches int
		expErr     bool
		errMsg     string
	}{
		{
			"invalid project-id",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/projects/%s/batches", val.APIAddress, "abc-d"),
			0,
			true,
			"invalid project id",
		},
		{
			"no batches found",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/projects/%s/batches", val.APIAddress, "P02"),
			0,
			false,
			"",
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/projects/%s/batches", val.APIAddress, "P01"),
			4,
			false,
			"",
		},
		{
			"valid request with pagination",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/projects/%s/batches?pagination.limit=2", val.APIAddress, "P01"),
			2,
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var batches ecocredit.QueryBatchesResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &batches)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(batches.Batches)
				require.Len(batches.Batches, tc.numBatches)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetBatch() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		url       string
		expErr    bool
		errMsg    string
		projectID string
	}{
		{
			"invalid batch denom",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/batches/%s", val.APIAddress, "C999"),
			true,
			"invalid denom",
			"",
		},
		{
			"no batches found",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/batches/%s", val.APIAddress, "A00-00000000-00000000-000"),
			true,
			"not found",
			"",
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/batches/%s", val.APIAddress, "C01-20210101-20210201-002"),
			false,
			"",
			"P01",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var batch ecocredit.QueryBatchInfoResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &batch)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(batch.Info)
				require.Equal(batch.Info.ProjectId, tc.projectID)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCreditTypes() {
	require := s.Require()
	val := s.network.Validators[0]

	url := fmt.Sprintf("%s/regen/ecocredit/v1alpha2/credit-types", val.APIAddress)
	resp, err := rest.GetRequest(url)
	require.NoError(err)

	var creditTypes ecocredit.QueryCreditTypesResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(resp, &creditTypes)

	require.NoError(err)
	require.Len(creditTypes.CreditTypes, 1)
	require.Equal(creditTypes.CreditTypes[0].Abbreviation, "C")
	require.Equal(creditTypes.CreditTypes[0].Name, "carbon")
}

func (s *IntegrationTestSuite) TestGetBalance() {
	val := s.network.Validators[0]

	testCases := []struct {
		name   string
		url    string
		expErr bool
		errMsg string
	}{
		{
			"invalid batch-denom",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/batches/%s/balance/%s", val.APIAddress, "abcd", val.Address.String()),
			true,
			"invalid denom",
		},
		{
			"invalid account address",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/batches/%s/balance/%s", val.APIAddress, "C01-20210101-20210201-001", "abcd"),
			true,
			"decoding bech32 failed",
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/batches/%s/balance/%s", val.APIAddress, "C01-20210101-20210201-002", val.Address.String()),
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var balance ecocredit.QueryBalanceResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &balance)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(balance)
				require.Equal(balance.TradableAmount, "100")
				require.Equal(balance.RetiredAmount, "0.000001")
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetSupply() {
	val := s.network.Validators[0]

	testCases := []struct {
		name   string
		url    string
		expErr bool
		errMsg string
	}{
		{
			"invalid batch-denom",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/batches/%s/supply", val.APIAddress, "abcd"),
			true,
			"invalid denom",
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha2/batches/%s/supply", val.APIAddress, "C01-20210101-20210201-001"),
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var supply ecocredit.QuerySupplyResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &supply)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(supply)
				require.Equal(supply.RetiredSupply, "0.000001")
				require.Equal(supply.TradableSupply, "100")
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGRPCQueryParams() {
	val := s.network.Validators[0]
	require := s.Require()

	resp, err := rest.GetRequest(fmt.Sprintf("%s/regen/ecocredit/v1alpha2/params", val.APIAddress))
	require.NoError(err)

	var params ecocredit.QueryParamsResponse
	require.NoError(val.ClientCtx.Codec.UnmarshalJSON(resp, &params))

	s.Require().Equal(ecocredit.DefaultParams(), *params.Params)
}

func (s *IntegrationTestSuite) TestGetSellOrder() {
	val := s.network.Validators[0]

	testCases := []struct {
		name   string
		url    string
		expErr bool
		errMsg string
	}{
		{
			"not found",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/sell-orders/id/%s", val.APIAddress, "99"),
			true,
			"not found",
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/sell-orders/id/%s", val.APIAddress, "1"),
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var sellOrder ecocredit.QuerySellOrderResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &sellOrder)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(sellOrder.SellOrder)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetSellOrders() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/sell-orders", val.APIAddress),
			false,
			"",
			3,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/sell-orders?pagination.limit=2", val.APIAddress),
			false,
			"",
			2,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var sellOrders ecocredit.QuerySellOrdersResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &sellOrders)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(sellOrders.SellOrders)
				require.Len(sellOrders.SellOrders, tc.expItems)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetSellOrdersByBatchDenom() {
	val := s.network.Validators[0]
	batchDenom := s.batchInfo.BatchDenom

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"invalid denom",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/sell-orders/batch-denom/%s", val.APIAddress, "foo"),
			true,
			"invalid denom",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/sell-orders/batch-denom/%s", val.APIAddress, batchDenom),
			false,
			"",
			3,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/sell-orders/batch-denom/%s?pagination.limit=2", val.APIAddress, batchDenom),
			false,
			"",
			2,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var sellOrders ecocredit.QuerySellOrdersByBatchDenomResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &sellOrders)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(sellOrders.SellOrders)
				require.Len(sellOrders.SellOrders, tc.expItems)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetSellOrdersByAddress() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"invalid address",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/sell-orders/address/%s", val.APIAddress, "abc"),
			true,
			"invalid request",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/sell-orders/address/%s", val.APIAddress, val.Address.String()),
			false,
			"",
			3,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/sell-orders/address/%s?pagination.limit=2", val.APIAddress, val.Address.String()),
			false,
			"",
			2,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var sellOrders ecocredit.QuerySellOrdersByAddressResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &sellOrders)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(sellOrders.SellOrders)
				require.Len(sellOrders.SellOrders, tc.expItems)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetBuyOrder() {
	val := s.network.Validators[0]

	testCases := []struct {
		name   string
		url    string
		expErr bool
		errMsg string
	}{
		{
			"not found",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/buy-orders/id/%s", val.APIAddress, "99"),
			true,
			"not found",
		},
		// TODO: filtered buy orders required #623
		//{
		//	"valid request",
		//	fmt.Sprintf("%s/regen/ecocredit/v1alpha1/buy-orders/id/%s", val.APIAddress, "1"),
		//	false,
		//	"",
		//},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var buyOrder ecocredit.QueryBuyOrderResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &buyOrder)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(buyOrder.BuyOrder)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetBuyOrders() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/buy-orders", val.APIAddress),
			false,
			"",
			3,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/buy-orders?pagination.limit=2", val.APIAddress),
			false,
			"",
			2,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var buyOrders ecocredit.QueryBuyOrdersResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &buyOrders)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(buyOrders.BuyOrders)
				// TODO: filtered buy orders required #623
				//require.Len(buyOrders.BuyOrders, tc.expItems)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetBuyOrdersByAddress() {
	val := s.network.Validators[0]
	addr := s.testAccount.String()

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"invalid address",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/buy-orders/address/%s", val.APIAddress, "abc"),
			true,
			"invalid request",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/buy-orders/address/%s", val.APIAddress, addr),
			false,
			"",
			3,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/buy-orders/address/%s?pagination.limit=2", val.APIAddress, addr),
			false,
			"",
			2,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var buyOrders ecocredit.QueryBuyOrdersByAddressResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &buyOrders)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(buyOrders.BuyOrders)
				// TODO: filtered buy orders required #623
				//require.Len(buyOrders.BuyOrders, tc.expItems)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetAllowedAskDenoms() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/ask-denoms", val.APIAddress),
			false,
			"",
			3,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/ask-denoms", val.APIAddress),
			false,
			"",
			2,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var askDenoms ecocredit.QueryAllowedAskDenomsResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &askDenoms)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(askDenoms.AskDenoms)
				// TODO: AllowAskDenom not yet implemented #624
				//require.Len(askDenoms.AskDenoms, tc.expItems)
			}
		})
	}
}
