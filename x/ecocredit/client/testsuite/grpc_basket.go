package testsuite

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

const basketRoute = "regen/ecocredit/basket/v1"

func (s *IntegrationTestSuite) TestQueryBasket() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/basket/%s", s.val.APIAddress, basketRoute, s.basketDenom),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/baskets/%s", s.val.APIAddress, basketRoute, s.basketDenom),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res basket.QueryBasketResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Basket) // deprecated
			require.NotEmpty(res.BasketInfo)
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBaskets() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/baskets", s.val.APIAddress, basketRoute),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/baskets?pagination.countTotal=true",
				// TODO: #1113
				// "%s/%s/sell-orders?pagination.limit=1&pagination.countTotal=true",
				s.val.APIAddress,
				basketRoute,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res basket.QueryBasketsResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Baskets) // deprecated
			require.NotEmpty(res.BasketsInfo)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Baskets, 1) // deprecated
				require.Len(res.BasketsInfo, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBasketBalance() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf(
				"%s/%s/basket-balance/%s/%s",
				s.val.APIAddress,
				basketRoute,
				s.basketDenom,
				s.batchDenom,
			),
		},
		{
			"valid alternative",
			fmt.Sprintf(
				"%s/%s/baskets/%s/balances/%s",
				s.val.APIAddress,
				basketRoute,
				s.basketDenom,
				s.batchDenom,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res basket.QueryBasketBalanceResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Balance)
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBasketBalances() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf(
				"%s/%s/basket-balances/%s",
				s.val.APIAddress,
				basketRoute,
				s.basketDenom,
			),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/basket-balances/%s?pagination.countTotal=true",
				// TODO: #1113
				// "%s/%s/sell-orders?pagination.limit=1&pagination.countTotal=true",
				s.val.APIAddress,
				basketRoute,
				s.basketDenom,
			),
		},
		{
			"valid alternative",
			fmt.Sprintf(
				"%s/%s/baskets/%s/balances",
				s.val.APIAddress,
				basketRoute,
				s.basketDenom,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res basket.QueryBasketBalancesResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Balances) // deprecated
			require.NotEmpty(res.BalancesInfo)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Balances, 1) // deprecated
				require.Len(res.BalancesInfo, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			}
		})
	}
}
