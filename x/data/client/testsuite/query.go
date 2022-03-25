package testsuite

import (
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/data/client"
)

func (s *IntegrationTestSuite) TestQueryByIRICmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

	validIri := "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"

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

func (s *IntegrationTestSuite) TestQueryByAttestorCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

	validIris := []string{
		"regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf",
		"regen:13toVgfX85Ny2ZTVxNzuL7MUquwwF7vcMKSAdVw2bUpEaL7XCFnshuh.rdf",
	}

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
			name:      "invalid attestor",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "invalid bech32 string",
		},
		{
			name:    "valid",
			args:    []string{val.Address.String()},
			expErr:  false,
			expIRIs: validIris,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryByAttestorCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res data.QueryByAttestorResponse
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

func (s *IntegrationTestSuite) TestQueryAttestorsCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

	validIri := "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"

	acc1, err := val.ClientCtx.Keyring.Key("acc1")
	s.Require().NoError(err)

	acc2, err := val.ClientCtx.Keyring.Key("acc2")
	s.Require().NoError(err)

	testCases := []struct {
		name         string
		args         []string
		expErr       bool
		expErrMsg    string
		expAttestors []string
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
			name:      "invalid attestor",
			args:      []string{"foo"},
			expErr:    true,
			expErrMsg: "key not found",
		},
		{
			name:   "valid",
			args:   []string{validIri},
			expErr: false,
			expAttestors: []string{
				acc1.GetAddress().String(),
				acc2.GetAddress().String(),
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryAttestorsCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res data.QueryAttestorsResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

				for _, attestor := range tc.expAttestors {
					s.Require().Contains(res.Attestors, attestor)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolverInfoCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

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
			name:      "invalid url",
			args:      []string{"abcd"},
			expErr:    true,
			expErrMsg: "not found",
		},
		{
			name:   "valid",
			args:   []string{"http://foo.bar"},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryResolverInfoCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res data.QueryResolverInfoResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolversCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	clientCtx.OutputFormat = "JSON"

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
			name:      "invalid iri",
			args:      []string{"abcd"},
			expErr:    true,
			expErrMsg: "can't find",
		},
		{
			name:   "valid test",
			args:   []string{s.iri},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := client.QueryResolversCmd()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err, out.String())
				s.Require().Contains(out.String(), tc.expErrMsg)
			} else {
				s.Require().NoError(err, out.String())

				var res data.QueryResolversResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
			}
		})
	}
}
