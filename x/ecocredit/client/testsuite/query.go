package testsuite

import (
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	coreclient "github.com/regen-network/regen-ledger/x/ecocredit/client"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func (s *IntegrationTestSuite) TestQueryClassesCmd() {
	val := s.network.Validators[0]
	val2 := s.network.Validators[1]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"
	require := s.Require()

	classId, err := s.createClass(clientCtx, &core.MsgCreateClass{
		Admin:            val.Address.String(),
		Issuers:          []string{val.Address.String(), val2.Address.String()},
		Metadata:         "metadata",
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &ecocredit.DefaultParams().CreditClassFee[0],
	})
	require.NoError(err)

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedErrMsg string
		numItems       int
	}{
		{
			name:           "no pagination flags",
			args:           []string{classId},
			expectErr:      false,
			expectedErrMsg: "",
			numItems:       -1,
		},
		{
			name:           "pagination limit 1",
			args:           []string{classId, "--limit=1"},
			expectErr:      false,
			expectedErrMsg: "",
			numItems:       1,
		},
		{
			name:           "class not found",
			args:           []string{"Z100"},
			expectErr:      true,
			expectedErrMsg: "not found",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := coreclient.QueryClassIssuersCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				require.Error(err)
				require.Contains(out.String(), tc.expectedErrMsg)
			} else {
				require.NoError(err, out.String())
			}
		})
	}
}
