package testsuite

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"

	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/data/client"
)

const outputFormat = "JSON"

func (s *IntegrationTestSuite) TestQueryAnchorByIRICmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

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
			name: "valid",
			args: []string{s.iri1},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryAnchorByIRICmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res data.QueryAnchorByIRIResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Anchor)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAnchorByHashCmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	bz, err := s.val.ClientCtx.Codec.MarshalJSON(s.hash1)
	require.NoError(err)

	filePath := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()

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
			name:      "invalid file path",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name: "valid",
			args: []string{filePath},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryAnchorByHashCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res data.QueryAnchorByHashResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Anchor)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAttestationsByAttestorCmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

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
			name: "valid",
			args: []string{s.addr1.String()},
		},
		{
			name: "valid with pagination",
			args: []string{
				s.addr1.String(),
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryAttestationsByAttestorCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res data.QueryAttestationsByAttestorResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Attestations)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.Attestations, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAttestationsByIRICmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

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
			name: "valid",
			args: []string{s.iri2},
		},
		{
			name: "valid with pagination",
			args: []string{
				s.iri2,
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryAttestationsByIRICmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res data.QueryAttestationsByIRIResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Attestations)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.Attestations, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAttestationsByHashCmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	bz, err := s.val.ClientCtx.Codec.MarshalJSON(s.hash1)
	require.NoError(err)

	filePath := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()

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
			name:      "invalid file path",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name: "valid",
			args: []string{filePath},
		},
		{
			name: "valid with pagination",
			args: []string{
				filePath,
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryAttestationsByHashCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res data.QueryAttestationsByHashResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Attestations)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.Attestations, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolverCmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

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
			name: "valid",
			args: []string{fmt.Sprint(s.resolverID)},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryResolverCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res data.QueryResolverResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Resolver)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolversByIRICmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

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
			name: "valid",
			args: []string{s.iri1},
		},
		{
			name: "valid with pagination",
			args: []string{
				s.iri1,
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryResolversByIRICmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err, out.String())
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res data.QueryResolversByIRIResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Resolvers)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.Resolvers, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolversByHashCmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	bz, err := s.val.ClientCtx.Codec.MarshalJSON(s.hash1)
	require.NoError(err)

	filePath := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()

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
			name:      "invalid file path",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name: "valid",
			args: []string{filePath},
		},
		{
			name: "valid with pagination",
			args: []string{
				filePath,
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryResolversByHashCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res data.QueryResolversByHashResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Resolvers)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.Resolvers, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolversByURLCmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

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
			name: "valid",
			args: []string{s.url},
		},
		{
			name: "valid with pagination",
			args: []string{
				s.url,
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s", flags.FlagCountTotal),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryResolversByURLCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err, out.String())
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res data.QueryResolversByURLResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Resolvers)

				if strings.Contains(tc.name, "pagination") {
					require.Len(res.Resolvers, 1)
					require.NotEmpty(res.Pagination)
					require.NotEmpty(res.Pagination.Total)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestConvertIRIToHashCmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

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
			name: "valid",
			args: []string{s.iri1},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.ConvertIRIToHashCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res data.ConvertIRIToHashResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.ContentHash)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestConvertHashToIRICmd() {
	require := s.Require()
	clientCtx := s.val.ClientCtx
	clientCtx.OutputFormat = outputFormat

	bz, err := s.val.ClientCtx.Codec.MarshalJSON(s.hash1)
	require.NoError(err)

	filePath := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()

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
			name:      "invalid file path",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name: "valid",
			args: []string{filePath},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.ConvertHashToIRICmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(err)
				require.Contains(out.String(), tc.expErrMsg)
			} else {
				require.NoError(err)

				var res data.ConvertHashToIRIResponse
				require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(res.Iri)
			}
		})
	}
}
