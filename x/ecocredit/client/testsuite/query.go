package testsuite

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/rand"
	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	coreclient "github.com/regen-network/regen-ledger/x/ecocredit/client"
	marketplaceclient "github.com/regen-network/regen-ledger/x/ecocredit/client/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (s *IntegrationTestSuite) TestQueryClassesCmd() {
	val := s.network.Validators[0]
	val2 := s.network.Validators[1]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
		Admin:            val.Address.String(),
		Issuers:          []string{val.Address.String()},
		Metadata:         "metadata",
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &ecocredit.DefaultParams().CreditClassFee[0],
	})
	s.Require().NoError(err)
	classId2, err := s.createClass(clientCtx, &core.MsgCreateClass{
		Admin:            val.Address.String(),
		Issuers:          []string{val.Address.String(), val2.Address.String()},
		Metadata:         "metadata2",
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &ecocredit.DefaultParams().CreditClassFee[0],
	})
	s.Require().NoError(err)
	classIds := [2]string{classId, classId2}

	testCases := []struct {
		name               string
		args               []string
		expectErr          bool
		expectedErrMsg     string
		expectedAmtClasses int
	}{
		{
			name:           "too many args",
			args:           []string{"abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 0 arg(s), received 1",
		},
		{
			name:               "no pagination flags",
			args:               []string{},
			expectErr:          false,
			expectedAmtClasses: -1,
		},
		{
			name: "limit 1",
			args: []string{
				fmt.Sprintf("--%s=1", flags.FlagLimit),
			},
			expectErr:          false,
			expectedAmtClasses: 1,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.QueryClassesCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryClassesResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

				if tc.expectedAmtClasses > 0 {
					s.Require().Len(res.Classes, tc.expectedAmtClasses)
				} else {
					resClassIds := make([]string, len(res.Classes))
					for i, cls := range res.Classes {
						resClassIds[i] = cls.Id
					}
					for _, id := range classIds {
						s.Require().Contains(resClassIds, id)
					}
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryClassInfoCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	class := &core.MsgCreateClass{
		Admin:            val.Address.String(),
		Issuers:          []string{val.Address.String()},
		Metadata:         "hi",
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &core.DefaultParams().CreditClassFee[0],
	}
	classId, err := s.createClass(clientCtx, class)
	s.Require().NoError(err)

	testCases := []struct {
		name              string
		args              []string
		expectErr         bool
		expectedErrMsg    string
		expectedClassInfo *core.ClassInfo
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:      "valid credit class",
			args:      []string{classId},
			expectErr: false,
			expectedClassInfo: &core.ClassInfo{
				Id:               classId,
				Admin:            val.Address,
				Metadata:         class.Metadata,
				CreditTypeAbbrev: class.CreditTypeAbbrev,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.QueryClassInfoCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryClassInfoResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				tc.expectedClassInfo.Key = res.Info.Key // force the db id's to be equal as we cannot know this beforehand.
				s.Require().Equal(tc.expectedClassInfo, res.Info)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchesCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	_, projectName, batchDenom := s.createClassProjectBatch(clientCtx, val.Address.String())

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name: "count",
			args: []string{
				projectName,
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.QueryBatchesCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryBatchesResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().True(len(res.Batches) > 0)
				s.Require().NotNil(res.Pagination)
				s.Require().True(res.Pagination.Total > 0)
				denoms := make([]string, len(res.Batches))
				for i, batch := range res.Batches {
					denoms[i] = batch.Denom
				}
				s.Require().Contains(denoms, batchDenom)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchInfoCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, val.Address.String())

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:      "valid credit batch",
			args:      []string{batchDenom},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.QueryBatchInfoCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryBatchInfoResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(res.Info.Denom, batchDenom)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBalanceCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, val.Address.String())

	testCases := []struct {
		name                   string
		args                   []string
		expectErr              bool
		expectedErrMsg         string
		expectedTradableAmount string
		expectedRetiredAmount  string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 2 arg(s), received 3",
		},
		{
			name:                   "valid",
			args:                   []string{batchDenom, val.Address.String()},
			expectErr:              false,
			expectedTradableAmount: "100",
			expectedRetiredAmount:  "0.000001",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.QueryBalanceCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryBalanceResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().True(sdk.AccAddress(res.Balance.Address).Equals(val.Address))
				s.Require().NotEmpty(res.Balance.Tradable)
				s.Require().NotEmpty(res.Balance.Retired)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySupplyCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, val.Address.String())

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:           "missing args",
			args:           []string{},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:           "too many args",
			args:           []string{"abcde", "abcde"},
			expectErr:      true,
			expectedErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:      "valid credit batch",
			args:      []string{batchDenom},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.QuerySupplyCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QuerySupplyResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().NotEmpty(res.TradableSupply)
				s.Require().NotEmpty(res.RetiredSupply)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCreditTypesCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	creditTypes := core.DefaultParams().CreditTypes
	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			name:           "valid",
			args:           []string{},
			expectErr:      false,
			expectedErrMsg: "",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.QueryCreditTypesCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expectedErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryCreditTypesResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				assert.DeepEqual(s.T(), res.CreditTypes, creditTypes)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryParamsCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	require := s.Require()

	cmd := coreclient.QueryParamsCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, []string{})
	require.NoError(err)

	var params core.QueryParamsResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &params))
	require.NoError(err)

	require.Equal(core.DefaultParams(), *params.Params)
}

func (s *IntegrationTestSuite) TestQuerySellOrderCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, val.Address.String())
	validAsk := sdk.NewInt64Coin(core.DefaultParams().AllowedAskDenoms[0].Denom, 10)
	expiration, err := types.ParseDate("expiration", "2050-03-11")
	s.Require().NoError(err)
	orderIds, err := s.createSellOrder(clientCtx, &marketplace.MsgSell{
		Owner: val.Address.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
		},
	})
	s.Require().NoError(err)

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "missing args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:      "valid",
			args:      []string{fmt.Sprintf("%d", orderIds[0])},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := marketplaceclient.QuerySellOrderCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res marketplace.QuerySellOrderResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().True(sdk.AccAddress(res.SellOrder.Seller).Equals(val.Address))
				s.Require().Equal(res.SellOrder.Quantity, "10")
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrdersCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, val.Address.String())
	validAsk := sdk.NewInt64Coin(core.DefaultParams().AllowedAskDenoms[0].Denom, 10)
	expiration, err := types.ParseDate("expiration", "2050-03-11")
	s.Require().NoError(err)

	_, err = s.createSellOrder(clientCtx, &marketplace.MsgSell{
		Owner: val.Address.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: batchDenom, Quantity: "3", AskPrice: &validAsk, Expiration: &expiration},
		},
	})
	s.Require().NoError(err)
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOrders []*ecocredit.SellOrder
	}{
		{
			name:      "too many args",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "Error: accepts 0 arg(s), received 1",
		},
		{
			name:      "valid",
			args:      []string{fmt.Sprintf("--%s", flags.FlagCountTotal)},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := marketplaceclient.QuerySellOrdersCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res marketplace.QuerySellOrdersResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().NotNil(res.Pagination)
				s.Require().True(res.Pagination.Total > 1)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrdersByAddressCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, val.Address.String())
	validAsk := sdk.NewInt64Coin(core.DefaultParams().AllowedAskDenoms[0].Denom, 10)
	expiration, err := types.ParseDate("expiration", "2050-03-11")
	s.Require().NoError(err)
	_, err = s.createSellOrder(clientCtx, &marketplace.MsgSell{
		Owner: val.Address.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: batchDenom, Quantity: "3", AskPrice: &validAsk, Expiration: &expiration},
		},
	})
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "missing args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:      "valid",
			args:      []string{val.Address.String(), fmt.Sprintf("--%s", flags.FlagCountTotal)},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := marketplaceclient.QuerySellOrdersByAddressCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res marketplace.QuerySellOrdersByAddressResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().NotNil(res.Pagination)
				s.Require().True(res.Pagination.Total > 1)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrdersByBatchDenomCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	_, _, batchDenom := s.createClassProjectBatch(clientCtx, val.Address.String())
	validAsk := sdk.NewInt64Coin(core.DefaultParams().AllowedAskDenoms[0].Denom, 10)
	expiration, err := types.ParseDate("expiration", "2050-03-11")
	s.Require().NoError(err)

	_, err = s.createSellOrder(clientCtx, &marketplace.MsgSell{
		Owner: val.Address.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: batchDenom, Quantity: "3", AskPrice: &validAsk, Expiration: &expiration},
		},
	})
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "missing args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:      "valid",
			args:      []string{batchDenom, fmt.Sprintf("--%s", flags.FlagCountTotal)},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := marketplaceclient.QuerySellOrdersByBatchDenomCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res marketplace.QuerySellOrdersByBatchDenomResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().NotNil(res.Pagination)
				s.Require().Len(res.SellOrders, 2)
				s.Require().Equal(uint64(2), res.Pagination.Total)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryProjectsCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
		Admin:            val.Address.String(),
		Issuers:          []string{val.Address.String()},
		Metadata:         "foo",
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &core.DefaultParams().CreditClassFee[0],
	})
	s.Require().NoError(err)
	pID, err := s.createProject(clientCtx, &core.MsgCreateProject{
		Issuer:              val.Address.String(),
		ClassId:             classId,
		Metadata:            "foo",
		ProjectJurisdiction: "US-OR",
		ProjectId:           rand.Str(3),
	})
	s.Require().NoError(err)
	pID2, err := s.createProject(clientCtx, &core.MsgCreateProject{
		Issuer:              val.Address.String(),
		ClassId:             classId,
		Metadata:            "foo",
		ProjectJurisdiction: "US-OR",
		ProjectId:           rand.Str(3),
	})
	s.Require().NoError(err)
	projectIds := [2]string{pID, pID2}
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expLen    int
	}{
		{
			name:      "no args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:   "valid query",
			args:   []string{classId, fmt.Sprintf("--%s", flags.FlagCountTotal)},
			expErr: false,
			expLen: 2,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.QueryProjectsCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res core.QueryProjectsResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Len(res.Projects, tc.expLen)
				s.Require().Equal(res.Pagination.Total, uint64(2))
				for _, project := range res.Projects {
					s.Require().Contains(projectIds, project.Id)
				}
			}
		})
	}

}

func (s *IntegrationTestSuite) TestQueryProjectInfoCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	require := s.Require()
	_, projectId, _ := s.createClassProjectBatch(clientCtx, val.Address.String())

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "no args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name:      "valid query",
			args:      []string{projectId},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.QueryProjectInfoCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err, out.String())

				var res core.QueryProjectInfoResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.Equal(projectId, res.Info.Id)
			}
		})
	}

}
