package testsuite

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/regen-network/regen-ledger/x/data"
)

func (s *IntegrationTestSuite) TestQueryAnchorByIRI() {
	val := s.network.Validators[0]

	iri := "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"

	testCases := []struct {
		name   string
		url    string
		expErr bool
		errMsg string
	}{
		{
			"invalid IRI",
			fmt.Sprintf("%s/regen/data/v1/anchor-by-iri/%s", val.APIAddress, "foo"),
			true,
			"not found",
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/anchor-by-iri/%s", val.APIAddress, iri),
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res data.QueryAnchorByIRIResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(bz, &res)

			if tc.expErr {
				require.Error(err)
			} else {
				require.NoError(err)
				require.NotNil(res.Anchor)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAnchorsByAttestor() {
	val := s.network.Validators[0]

	acc1, err := val.ClientCtx.Keyring.Key("acc1")
	s.Require().NoError(err)

	addr := acc1.GetAddress().String()

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"invalid attestor",
			fmt.Sprintf("%s/regen/data/v1/anchors-by-attestor/%s", val.APIAddress, "foo"),
			true,
			"invalid bech32 string",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/anchors-by-attestor/%s", val.APIAddress, addr),
			false,
			"",
			2,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/data/v1/anchors-by-attestor/%s?pagination.limit=1", val.APIAddress, addr),
			false,
			"",
			1,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res data.QueryAnchorsByAttestorResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(bz, &res)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(bz), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(res.Anchors)
				require.Len(res.Anchors, tc.expItems)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestConvertIRIToHash() {
	val := s.network.Validators[0]

	iri := "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"

	testCases := []struct {
		name   string
		url    string
		expErr bool
		errMsg string
	}{
		{
			"invalid IRI",
			fmt.Sprintf("%s/regen/data/v1/iri-to-hash/%s", val.APIAddress, "foo"),
			true,
			"invalid IRI",
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/iri-to-hash/%s", val.APIAddress, iri),
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res data.ConvertIRIToHashResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(bz, &res)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(bz), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(res.ContentHash)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestConvertRawHashToIRI() {
	val := s.network.Validators[0]

	iri, ch := s.createIRIAndRawHash([]byte("xyzabc123"))

	encodedHash := encodeBase64Bytes(ch.Raw.Hash)

	testCases := []struct {
		name   string
		url    string
		expErr bool
		errMsg string
	}{
		{
			"empty hash",
			fmt.Sprintf(
				"%s/regen/data/v1/raw-hash-to-iri?digest_algorithm=%s",
				val.APIAddress,
				ch.Raw.DigestAlgorithm, // enum 1
			),
			true,
			"hash cannot be empty",
		},
		{
			"invalid hash",
			fmt.Sprintf(
				"%s/regen/data/v1/raw-hash-to-iri?hash=%s&digest_algorithm=%s",
				val.APIAddress,
				"foo",
				ch.Raw.DigestAlgorithm, // enum 1
			),
			true,
			"failed to decode base64 string",
		},
		{
			"unspecified digest algorithm",
			fmt.Sprintf(
				"%s/regen/data/v1/raw-hash-to-iri?hash=%s",
				val.APIAddress,
				encodedHash, // base64 encoded string
			),
			true,
			"digest algorithm cannot be unspecified",
		},
		{
			"invalid digest algorithm",
			fmt.Sprintf(
				"%s/regen/data/v1/raw-hash-to-iri?hash=%s&digest_algorithm=%s",
				val.APIAddress,
				encodedHash, // base64 encoded string
				"foo",
			),
			true,
			"foo is not a valid data.DigestAlgorithm",
		},
		{
			"invalid media type",
			fmt.Sprintf(
				"%s/regen/data/v1/raw-hash-to-iri?hash=%s&digest_algorithm=%d&media_type=%s",
				val.APIAddress,
				encodedHash,            // base64 encoded string
				ch.Raw.DigestAlgorithm, // enum 1
				"foo",
			),
			true,
			"foo is not a valid data.RawMediaType",
		},
		{
			"valid request",
			fmt.Sprintf(
				"%s/regen/data/v1/raw-hash-to-iri?hash=%s&digest_algorithm=%d&media_type=%d",
				val.APIAddress,
				encodedHash,            // base64 encoded string
				ch.Raw.DigestAlgorithm, // enum 1
				ch.Raw.MediaType,       // enum 0
			),
			false,
			"",
		},
		{
			"valid request enums as strings",
			fmt.Sprintf(
				"%s/regen/data/v1/raw-hash-to-iri?hash=%s&digest_algorithm=%s&media_type=%s",
				val.APIAddress,
				encodedHash,                    // base64 encoded string
				"DIGEST_ALGORITHM_BLAKE2B_256", // enum 1
				"RAW_MEDIA_TYPE_UNSPECIFIED",   // enum 1
			),
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res data.ConvertRawHashToIRIResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(bz, &res)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(bz), tc.errMsg)
			} else {
				require.NoError(err)
				require.Equal(iri, res.Iri)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestConvertGraphHashToIRI() {
	val := s.network.Validators[0]

	iri, ch := s.createIRIAndGraphHash([]byte("xyzabc123"))

	encodedHash := encodeBase64Bytes(ch.Graph.Hash)

	testCases := []struct {
		name   string
		url    string
		expErr bool
		errMsg string
	}{
		{
			"empty hash",
			fmt.Sprintf(
				"%s/regen/data/v1/graph-hash-to-iri?digest_algorithm=%d&canonicalization_algorithm=%d",
				val.APIAddress,
				ch.Graph.DigestAlgorithm,           // enum 1
				ch.Graph.CanonicalizationAlgorithm, // enum 1
			),
			true,
			"hash cannot be empty",
		},
		{
			"invalid hash",
			fmt.Sprintf(
				"%s/regen/data/v1/graph-hash-to-iri?hash=%s&digest_algorithm=%d&canonicalization_algorithm=%d",
				val.APIAddress,
				"foo",
				ch.Graph.DigestAlgorithm,           // enum 1
				ch.Graph.CanonicalizationAlgorithm, // enum 1
			),
			true,
			"failed to decode base64 string",
		},
		{
			"unspecified digest algorithm",
			fmt.Sprintf(
				"%s/regen/data/v1/graph-hash-to-iri?hash=%s&canonicalization_algorithm=%d",
				val.APIAddress,
				encodedHash,                        // base64 encoded string
				ch.Graph.CanonicalizationAlgorithm, // enum 1
			),
			true,
			"digest algorithm cannot be unspecified",
		},
		{
			"invalid digest algorithm",
			fmt.Sprintf(
				"%s/regen/data/v1/graph-hash-to-iri?hash=%s&digest_algorithm=%s&canonicalization_algorithm=%d",
				val.APIAddress,
				encodedHash, // base64 encoded string
				"foo",
				ch.Graph.CanonicalizationAlgorithm, // enum 1
			),
			true,
			"foo is not a valid data.DigestAlgorithm",
		},
		{
			"unspecified canonicalization algorithm",
			fmt.Sprintf(
				"%s/regen/data/v1/graph-hash-to-iri?hash=%s&digest_algorithm=%s",
				val.APIAddress,
				encodedHash,              // base64 encoded string
				ch.Graph.DigestAlgorithm, // enum 1
			),
			true,
			"canonicalization algorithm cannot be unspecified",
		},
		{
			"invalid canonicalization algorithm",
			fmt.Sprintf(
				"%s/regen/data/v1/graph-hash-to-iri?hash=%s&digest_algorithm=%s&canonicalization_algorithm=%s",
				val.APIAddress,
				encodedHash,              // base64 encoded string
				ch.Graph.DigestAlgorithm, // enum 1
				"foo",
			),
			true,
			"foo is not a valid data.GraphCanonicalizationAlgorithm",
		},
		{
			"valid request",
			fmt.Sprintf(
				"%s/regen/data/v1/graph-hash-to-iri?hash=%s&digest_algorithm=%d&canonicalization_algorithm=%d",
				val.APIAddress,
				encodedHash,                        // base64 encoded string
				ch.Graph.DigestAlgorithm,           // enum 1
				ch.Graph.CanonicalizationAlgorithm, // enum 1
			),
			false,
			"",
		},
		{
			"valid request enums as strings",
			fmt.Sprintf(
				"%s/regen/data/v1/graph-hash-to-iri?hash=%s&digest_algorithm=%s&canonicalization_algorithm=%s",
				val.APIAddress,
				encodedHash,                    // base64 encoded string
				"DIGEST_ALGORITHM_BLAKE2B_256", // enum 1
				"GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015", // enum 1
			),
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res data.ConvertGraphHashToIRIResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(bz, &res)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(bz), tc.errMsg)
			} else {
				require.NoError(err)
				require.Equal(iri, res.Iri)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAttestorsByIRI() {
	val := s.network.Validators[0]

	iri := "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"invalid attestor",
			fmt.Sprintf("%s/regen/data/v1/attestors-by-iri/%s", val.APIAddress, "foo"),
			true,
			"not found",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/attestors-by-iri/%s", val.APIAddress, iri),
			false,
			"",
			2,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/data/v1/attestors-by-iri/%s?pagination.limit=1", val.APIAddress, iri),
			false,
			"",
			1,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res data.QueryAttestorsByIRIResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(bz, &res)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(bz), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(res.Attestors)
				require.Len(res.Attestors, tc.expItems)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolver() {
	val := s.network.Validators[0]

	testCases := []struct {
		name   string
		url    string
		expErr bool
		errMsg string
	}{
		{
			"not found",
			fmt.Sprintf("%s/regen/data/v1/resolver/%d", val.APIAddress, 404),
			true,
			"not found",
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/resolver/%d", val.APIAddress, s.resolverID),
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res data.QueryResolverResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(bz, &res)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(bz), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(res.Resolver.Url)
				require.NotNil(res.Resolver.Manager)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolversByIri() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"not found",
			fmt.Sprintf("%s/regen/data/v1/resolvers-by-iri/%s", val.APIAddress, "foo"),
			true,
			"not found",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/resolvers-by-iri/%s", val.APIAddress, s.iri),
			false,
			"",
			2,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/data/v1/resolvers-by-iri/%s?pagination.limit=1", val.APIAddress, s.iri),
			false,
			"",
			1,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res data.QueryResolversByIRIResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(bz, &res)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(bz), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(res.Resolvers)
				require.Len(res.Resolvers, tc.expItems)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolversByURL() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"empty url",
			fmt.Sprintf("%s/regen/data/v1/resolvers-by-url", val.APIAddress),
			true,
			"url cannot be empty",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/resolvers-by-url?url=%s", val.APIAddress, s.url),
			false,
			"",
			2,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/data/v1/resolvers-by-url?url=%s&pagination.limit=1", val.APIAddress, s.url),
			false,
			"",
			1,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res data.QueryResolversByURLResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(bz, &res)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(bz), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(res.Resolvers)
				require.Len(res.Resolvers, tc.expItems)
			}
		})
	}
}

func encodeBase64Bytes(bz []byte) string {
	// encode base64 bytes to base64 string
	str := base64.StdEncoding.EncodeToString(bz)
	// replace all instances of "+" with "%2b"
	return strings.Replace(str, "+", "%2b", -1)
}
