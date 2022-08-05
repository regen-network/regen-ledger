package testsuite

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/testutil/rest"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/regen-network/regen-ledger/x/data"
)

const dataRoute = "regen/data/v1"

func (s *IntegrationTestSuite) TestQueryAnchorByIRI() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/anchor-by-iri/%s", s.val.APIAddress, dataRoute, s.iri1),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/anchors/iri/%s", s.val.APIAddress, dataRoute, s.iri1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res data.QueryAnchorByIRIResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Anchor)
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAnchorByHash() {
	require := s.Require()

	hash, err := json.Marshal(s.hash1)
	require.NoError(err)

	testCases := []struct {
		name string
		url  string
		body []byte
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/anchor-by-hash", s.val.APIAddress, dataRoute),
			[]byte(fmt.Sprintf(`{"content_hash": %s}`, hash)),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/anchors/hash", s.val.APIAddress, dataRoute),
			[]byte(fmt.Sprintf(`{"content_hash": %s}`, hash)),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.PostRequest(tc.url, "text/JSON", tc.body)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res data.QueryAnchorByHashResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Anchor)
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAttestationsByAttestor() {
	require := s.Require()

	pgn := "pagination.countTotal=true"
	// TODO: #1113
	// pgn := pagination.limit=1&pagination.countTotal=true

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf(
				"%s/%s/attestations-by-attestor/%s",
				s.val.APIAddress,
				dataRoute,
				s.addr1,
			),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/attestations-by-attestor/%s?%s",
				s.val.APIAddress,
				dataRoute,
				s.addr1,
				pgn,
			),
		},
		{
			"valid alternative",
			fmt.Sprintf(
				"%s/%s/attestations/attestor/%s",
				s.val.APIAddress,
				dataRoute,
				s.addr1,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res data.QueryAttestationsByAttestorResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Attestations)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Attestations, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			} else {
				require.Empty(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAttestationsByIRI() {
	require := s.Require()

	pgn := "pagination.limit=1&pagination.countTotal=true"

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf(
				"%s/%s/attestations-by-iri/%s",
				s.val.APIAddress,
				dataRoute,
				s.iri2,
			),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/attestations-by-iri/%s?%s",
				s.val.APIAddress,
				dataRoute,
				s.iri2,
				pgn,
			),
		},
		{
			"valid alternative",
			fmt.Sprintf(
				"%s/%s/attestations/iri/%s",
				s.val.APIAddress,
				dataRoute,
				s.iri2,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res data.QueryAttestationsByIRIResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Attestations)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Attestations, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			} else {
				require.Empty(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAttestationsByHash() {
	require := s.Require()

	hash, err := json.Marshal(s.hash1)
	require.NoError(err)

	pgn, err := json.Marshal(query.PageRequest{
		Limit:      1,
		CountTotal: true,
	})
	require.NoError(err)

	testCases := []struct {
		name string
		url  string
		body []byte
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/attestations-by-hash", s.val.APIAddress, dataRoute),
			[]byte(fmt.Sprintf(`{"content_hash": %s}`, hash)),
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/attestations-by-hash", s.val.APIAddress, dataRoute),
			[]byte(fmt.Sprintf(`{"content_hash": %s, "pagination": %s}`, hash, pgn)),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/attestations/hash", s.val.APIAddress, dataRoute),
			[]byte(fmt.Sprintf(`{"content_hash": %s}`, hash)),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.PostRequest(tc.url, "text/JSON", tc.body)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res data.QueryAttestationsByHashResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Attestations)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Attestations, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			} else {
				require.Empty(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolver() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/resolver/%d", s.val.APIAddress, dataRoute, s.resolverID),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/resolvers/%d", s.val.APIAddress, dataRoute, s.resolverID),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res data.QueryResolverResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Resolver)
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolversByIRI() {
	require := s.Require()

	pgn := "pagination.limit=1&pagination.countTotal=true"

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/resolvers-by-iri/%s", s.val.APIAddress, dataRoute, s.iri1),
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/resolvers-by-iri/%s?%s", s.val.APIAddress, dataRoute, s.iri1, pgn),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/resolvers/iri/%s", s.val.APIAddress, dataRoute, s.iri1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res data.QueryResolversByIRIResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Resolvers)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Resolvers, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			} else {
				require.Empty(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolversByHash() {
	require := s.Require()

	hash, err := json.Marshal(s.hash1)
	require.NoError(err)

	pgn, err := json.Marshal(query.PageRequest{
		Limit:      1,
		CountTotal: true,
	})
	require.NoError(err)

	testCases := []struct {
		name string
		url  string
		body []byte
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/resolvers-by-hash", s.val.APIAddress, dataRoute),
			[]byte(fmt.Sprintf(`{"content_hash": %s}`, hash)),
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/resolvers-by-hash", s.val.APIAddress, dataRoute),
			[]byte(fmt.Sprintf(`{"content_hash": %s, "pagination": %s}`, hash, pgn)),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/resolvers/hash", s.val.APIAddress, dataRoute),
			[]byte(fmt.Sprintf(`{"content_hash": %s}`, hash)),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.PostRequest(tc.url, "text/JSON", tc.body)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res data.QueryResolversByHashResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Resolvers)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Resolvers, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			} else {
				require.Empty(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolversByURL() {
	require := s.Require()

	pgn, err := json.Marshal(query.PageRequest{
		Limit:      1,
		CountTotal: true,
	})
	require.NoError(err)

	testCases := []struct {
		name string
		url  string
		body []byte
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/resolvers-by-url", s.val.APIAddress, dataRoute),
			[]byte(fmt.Sprintf(`{"url": "%s"}`, s.url)),
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/resolvers-by-url", s.val.APIAddress, dataRoute),
			[]byte(fmt.Sprintf(`{"url": "%s", "pagination": %s}`, s.url, pgn)),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/resolvers/url", s.val.APIAddress, dataRoute),
			[]byte(fmt.Sprintf(`{"url": "%s"}`, s.url)),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.PostRequest(tc.url, "text/JSON", tc.body)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res data.QueryResolversByURLResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Resolvers)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Resolvers, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			} else {
				require.Empty(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestConvertIRIToHash() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/convert-iri-to-hash/%s", s.val.APIAddress, dataRoute, s.iri1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res data.ConvertIRIToHashResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.ContentHash)
		})
	}
}

func (s *IntegrationTestSuite) TestConvertHashToIRI() {
	require := s.Require()

	hash, err := json.Marshal(s.hash1)
	require.NoError(err)

	testCases := []struct {
		name string
		url  string
		body []byte
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/convert-hash-to-iri", s.val.APIAddress, dataRoute),
			[]byte(fmt.Sprintf(`{"content_hash": %s}`, hash)),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.PostRequest(tc.url, "text/JSON", tc.body)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res data.ConvertHashToIRIResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Iri)
		})
	}
}
