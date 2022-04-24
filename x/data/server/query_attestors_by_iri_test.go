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

type queryAttestorsByIriSuite struct {
	*baseSuite
	ce  *data.ContentEntry
	req *data.QueryAttestorsByIRIRequest
	res *data.QueryAttestorsByIRIResponse
	err error
}

func TestQueryAttestorsByIri(t *testing.T) {
	gocuke.NewRunner(t, &queryAttestorsByIriSuite{}).Path("./features/query_attestors_by_iri.feature").Run()
}

func (s *queryAttestorsByIriSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryAttestorsByIriSuite) TheContentEntry(a gocuke.DocString) {
	s.ce = &data.ContentEntry{}
	err := jsonpb.UnmarshalString(a.Content, s.ce)
	require.NoError(s.t, err)

	id, err := s.server.getOrCreateDataId(s.ctx, s.ce.Iri)
	require.NoError(s.t, err)

	s.server.stateStore.DataIDTable().Insert(s.ctx, &api.DataID{
		Id:  id,
		Iri: s.ce.Iri,
	})
}

func (s *queryAttestorsByIriSuite) TheAttestorEntry(a gocuke.DocString) {
	attestorEntry := &data.AttestorEntry{}
	err := jsonpb.UnmarshalString(a.Content, attestorEntry)
	require.NoError(s.t, err)

	id, err := s.server.getOrCreateDataId(s.ctx, s.ce.Iri)
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

func (s *queryAttestorsByIriSuite) TheRequest(a gocuke.DocString) {
	s.req = &data.QueryAttestorsByIRIRequest{}
	err := jsonpb.UnmarshalString(a.Content, s.req)
	require.NoError(s.t, err)
}

func (s *queryAttestorsByIriSuite) TheRequestIsExecuted() {
	s.res, s.err = s.server.AttestorsByIRI(s.ctx, s.req)
}

func (s *queryAttestorsByIriSuite) TheResponse(a gocuke.DocString) {
	res := &data.QueryAttestorsByIRIResponse{}
	err := jsonpb.UnmarshalString(a.Content, res)
	require.NoError(s.t, err)

	require.Equal(s.t, len(res.Attestors), len(s.res.Attestors))
	for i, attestor := range res.Attestors {
		require.Equal(s.t, attestor, s.res.Attestors[i])
	}
}

func (s *queryAttestorsByIriSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}
