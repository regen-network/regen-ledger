package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/gocuke"
	"github.com/regen-network/regen-ledger/x/data"
)

type queryAttestorsSuite struct {
	*baseSuite
	contentEntry  *data.ContentEntry
	attestorEntry *data.AttestorEntry
	request       *data.QueryAttestorsByIRIRequest
	response      *data.QueryAttestorsByIRIResponse
	err           error
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
	s.attestorEntry = &data.AttestorEntry{}
	err := jsonpb.UnmarshalString(a.Content, s.attestorEntry)
	require.NoError(s.t, err)

	id, err := s.server.getOrCreateDataId(s.ctx, s.contentEntry.Iri)
	require.NoError(s.t, err)

	attestor, err := sdk.AccAddressFromBech32(s.attestorEntry.Attestor)
	require.NoError(s.t, err)

	err = s.server.stateStore.DataAttestorTable().Insert(s.ctx, &api.DataAttestor{
		Id:        id,
		Attestor:  attestor,
		Timestamp: types.GogoToProtobufTimestamp(s.attestorEntry.Timestamp),
	})
	require.NoError(s.t, err)
}

func (s *queryAttestorsSuite) TheQueryByIriRequest(a gocuke.DocString) {
	s.request = &data.QueryAttestorsByIRIRequest{}
	err := jsonpb.UnmarshalString(a.Content, s.request)
	require.NoError(s.t, err)
}

func (s *queryAttestorsSuite) TheQueryIsExecuted() {
	s.response, s.err = s.server.AttestorsByIRI(s.ctx, s.request)
}

func (s *queryAttestorsSuite) TheQueryByIriResponse(a gocuke.DocString) {
	res := &data.QueryAttestorsByIRIResponse{}
	err := jsonpb.UnmarshalString(a.Content, res)
	require.NoError(s.t, err)

	require.Equal(s.t, len(res.Attestors), len(s.response.Attestors))
	require.Equal(s.t, res.Attestors, s.response.Attestors)
}

func (s *queryAttestorsSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}
