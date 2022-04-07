package server

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data"
)

type attestSuite struct {
	*baseSuite
	chg      *data.ContentHash_Graph
	attestor sdk.AccAddress
	err      error
}

func TestAttest(t *testing.T) {
	gocuke.NewRunner(t, &attestSuite{}).Path("./features/attest.feature").Run()
}

func (s *attestSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.attestor = s.addrs[0]
}

func (s *attestSuite) AGraphDataContentHash() {
	s.chg = &data.ContentHash_Graph{
		Hash:                      make([]byte, 32),
		DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
	}
}

func (s *attestSuite) TheDataHasBeenAnchored() {
	iri, err := s.chg.ToIRI()
	require.NoError(s.t, err)

	id := s.server.iriHasher.CreateID([]byte(iri), 0)

	err = s.server.stateStore.DataIDTable().Insert(s.ctx, &api.DataID{
		Id:  id,
		Iri: iri,
	})
	require.NoError(s.t, err)

	err = s.server.stateStore.DataAnchorTable().Insert(s.ctx, &api.DataAnchor{
		Id:        id,
		Timestamp: &timestamppb.Timestamp{},
	})
	require.NoError(s.t, err)
}

func (s *attestSuite) TheDataHasNotBeenAnchored() {
	// skip
}

func (s *attestSuite) AUserAttemptsToAttestToTheData() {
	_, s.err = s.server.Attest(s.ctx, &data.MsgAttest{
		Attestor: s.attestor.String(),
		Hashes:   []*data.ContentHash_Graph{s.chg},
	})
}

func (s *attestSuite) TheDataIsAttestedTo() {
	require.NoError(s.t, s.err)

	iri, err := s.chg.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)

	dataAttestor, err := s.server.stateStore.DataAttestorTable().Get(s.ctx, dataId.Id, s.attestor)
	require.NoError(s.t, err)
	require.Equal(s.t, s.attestor.Bytes(), dataAttestor.Attestor)
	require.Equal(s.t, s.sdkCtx.BlockTime(), dataAttestor.Timestamp.AsTime())
}

func (s *attestSuite) TheDataIsNotAttestedTo() {
	require.Error(s.t, s.err)
}
