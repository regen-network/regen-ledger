package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
)

type queryAttestorsSuite struct {
	*baseSuite
	contentEntry   *data.ContentEntry
	byIriRequest   *data.QueryAttestorsByIRIRequest
	byIriResponse  *data.QueryAttestorsByIRIResponse
	byHashRequest  *data.QueryAttestorsByHashRequest
	byHashResponse *data.QueryAttestorsByHashResponse
	err            error
}

func TestQueryAttestors(t *testing.T) {
	runner := gocuke.NewRunner(t, &queryAttestorsSuite{}).Path("./features/query_attestors.feature")
	runner.Step(`^the\s+content\s+hash\s+"((?:[^\"]|\")*)"`, (*queryAttestorsSuite).TheAttestorEntry)
	runner.Run()
}

func (s *queryAttestorsSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryAttestorsSuite) TheContentEntry(a gocuke.DocString) {
	s.contentEntry = &data.ContentEntry{}
	err := jsonpb.UnmarshalString(a.Content, s.contentEntry)
	require.NoError(s.t, err)

	id, err := s.server.getOrCreateDataId(s.ctx, s.contentEntry.Iri)
	require.NoError(s.t, err)

	s.server.stateStore.DataIDTable().Insert(s.ctx, &api.DataID{
		Id:  id,
		Iri: s.contentEntry.Iri,
	})
}

func (s *queryAttestorsSuite) TheAttestorEntry(a gocuke.DocString) {
	attestorEntry := &data.AttestorEntry{}
	err := jsonpb.UnmarshalString(a.Content, attestorEntry)
	require.NoError(s.t, err)

	id, err := s.server.getOrCreateDataId(s.ctx, s.contentEntry.Iri)
	require.NoError(s.t, err)

	attestor, err := sdk.AccAddressFromBech32(attestorEntry.Attestor)
	require.NoError(s.t, err)

	err = s.server.stateStore.DataAttestorTable().Insert(s.ctx, &api.DataAttestor{
		Id:        id,
		Attestor:  attestor,
		Timestamp: types.GogoToProtobufTimestamp(attestorEntry.Timestamp),
	})
	require.NoError(s.t, err)
}

func (s *queryAttestorsSuite) TheQueryByIriRequest(a gocuke.DocString) {
	s.byIriRequest = &data.QueryAttestorsByIRIRequest{}
	err := jsonpb.UnmarshalString(a.Content, s.byIriRequest)
	require.NoError(s.t, err)
}

func (s *queryAttestorsSuite) TheQueryByHashRequest(a gocuke.DocString) {
	s.byHashRequest = &data.QueryAttestorsByHashRequest{}
	err := jsonpb.UnmarshalString(a.Content, s.byHashRequest)
	require.NoError(s.t, err)
}

func (s *queryAttestorsSuite) TheQueryByIriIsExecuted() {
	s.byIriResponse, s.err = s.server.AttestorsByIRI(s.ctx, s.byIriRequest)
}

func (s *queryAttestorsSuite) TheQueryByHashIsExecuted() {
	s.byHashResponse, s.err = s.server.AttestorsByHash(s.ctx, s.byHashRequest)
}

func (s *queryAttestorsSuite) TheQueryByIriResponse(a gocuke.DocString) {
	res := &data.QueryAttestorsByIRIResponse{}
	err := jsonpb.UnmarshalString(a.Content, res)
	require.NoError(s.t, err)

	require.Equal(s.t, len(res.Attestors), len(s.byIriResponse.Attestors))
	for i, attestor := range res.Attestors {
		require.Equal(s.t, attestor, s.byIriResponse.Attestors[i])
	}
}

func (s *queryAttestorsSuite) TheQueryByHashResponse(a gocuke.DocString) {
	res := &data.QueryAttestorsByIRIResponse{}
	err := jsonpb.UnmarshalString(a.Content, res)
	require.NoError(s.t, err)

	require.Equal(s.t, len(res.Attestors), len(s.byHashResponse.Attestors))
	for i, attestor := range res.Attestors {
		require.Equal(s.t, attestor, s.byHashResponse.Attestors[i])
	}
}

func (s *queryAttestorsSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}
