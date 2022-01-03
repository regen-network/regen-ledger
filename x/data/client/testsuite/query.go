package testsuite

import (
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/data/client"
)

func (s *IntegrationTestSuite) TestZQueryByIRICmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

	validIri := "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expIRI    string
	}{
		{
			name:      "missing args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name: "too many args",
			args: []string{
				"foo", "bar",
			},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name: "invalid iri",
			args: []string{
				"foo",
			},
			expErr:    true,
			expErrMsg: "key not found",
		},
		{
			name:   "valid",
			args:   []string{validIri},
			expErr: false,
			expIRI: validIri,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryByIRICmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res data.QueryByIRIResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

				s.Require().Equal(tc.expIRI, res.Entry.Iri)
				s.Require().NotNil(res.Entry.Hash)
				s.Require().NotNil(res.Entry.Timestamp)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestZQueryBySignerCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

	validIri := "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expIRIs   []string
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
			name:      "invalid signer",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "invalid bech32 string",
		},
		{
			name:    "valid",
			args:    []string{val.Address.String()},
			expErr:  false,
			expIRIs: []string{validIri},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryBySignerCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res data.QueryBySignerResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

				for i, entry := range res.Entries {
					s.Require().Equal(tc.expIRIs[i], entry.Iri)
					s.Require().NotNil(entry.Hash)
					s.Require().NotNil(entry.Timestamp)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestZQuerySignersCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

	validIri := "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expResp   []string
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
			name:      "invalid signer",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "key not found",
		},
		{
			name:    "valid",
			args:    []string{validIri},
			expErr:  false,
			expResp: []string{val.Address.String()},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QuerySignersCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res data.QuerySignersResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				s.Require().Equal(tc.expResp, res.Signers)
			}
		})
	}
}
