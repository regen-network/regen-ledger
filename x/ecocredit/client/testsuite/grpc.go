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
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/class", val.APIAddress),
			true,
			0,
		},
		{
			"valid query",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/classes", val.APIAddress),
			false,
			4,
		},
		{
			"valid query pagination",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/classes?pagination.limit=2", val.APIAddress),
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
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/class", val.APIAddress),
			true,
			"Not Implemented",
			"",
		},
		{
			"class not found",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/classes/%s", val.APIAddress, "C999"),
			true,
			"not found",
			"",
		},
		{
			"valid class-id",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/classes/%s", val.APIAddress, "C01"),
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
	}{
		{
			"valid request",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/batches", val.APIAddress),
			3,
		},
		{
			"valid request with pagination",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/batches?pagination.limit=2", val.APIAddress),
			2,
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

			require.NoError(err)
			require.NotNil(batches.Batches)
			require.Len(batches.Batches, tc.numBatches)
		})
	}
}

func (s *IntegrationTestSuite) TestGetBatch() {
	val := s.network.Validators[0]

	testCases := []struct {
		name       string
		url        string
		expErr     bool
		errMsg     string
		batchDenom string
	}{
		{
			"batch not found",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/batches/%s", val.APIAddress, "C999"),
			true,
			"not found",
			"",
		},
		{
			"valid test",
			fmt.Sprintf("%s/regen/ecocredit/v1alpha1/batches/%s", val.APIAddress, "C01"),
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

			var batch ecocredit.QueryBatchInfoResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &batch)

			if tc.expErr {
				require.Error(err)
			} else {
				require.NoError(err)
				require.NotNil(batch.Info)
				require.Contains(batch.Info.BatchDenom, tc.batchDenom)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCreditTypes() {
	require := s.Require()
	val := s.network.Validators[0]

	url := fmt.Sprintf("%s/regen/ecocredit/v1alpha1/credit-types", val.APIAddress)
	resp, err := rest.GetRequest(url)
	require.NoError(err)

	var creditTypes ecocredit.QueryCreditTypesResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(resp, &creditTypes)

	require.NoError(err)
	require.Equal(creditTypes.String(), `{[name:"carbon" abbreviation:"C" unit:"metric ton CO2 equivalent" precision:6 ]}`)
}

func (s *IntegrationTestSuite) TestGetBalance() {
	// val := s.network.Validators[0]

}
