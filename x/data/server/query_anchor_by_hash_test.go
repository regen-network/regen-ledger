package server

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/x/data"
)

func TestQuery_AnchorByHash(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	id := []byte{0}
	ch := &data.ContentHash{Graph: &data.ContentHash_Graph{
		Hash:                      bytes.Repeat([]byte{0}, 32),
		DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
	}}
	iri, err := ch.ToIRI()
	require.NoError(t, err)

	// insert data id
	err = s.server.stateStore.DataIDTable().Insert(s.ctx, &api.DataID{
		Id:  id,
		Iri: iri,
	})
	require.NoError(t, err)

	timestamp := timestamppb.New(time.Now().UTC())

	// insert data anchor
	err = s.server.stateStore.DataAnchorTable().Insert(s.ctx, &api.DataAnchor{
		Id:        id,
		Timestamp: timestamp,
	})
	require.NoError(t, err)

	// query data anchor with valid content hash
	res, err := s.server.AnchorByHash(s.ctx, &data.QueryAnchorByHashRequest{
		ContentHash: ch,
	})
	require.NoError(t, err)
	require.Equal(t, iri, res.Anchor.Iri)
	require.Equal(t, ch, res.Anchor.ContentHash)
	require.Equal(t, timestamp.Seconds, res.Anchor.Timestamp.Seconds)
	require.Equal(t, timestamp.Nanos, res.Anchor.Timestamp.Nanos)

	// query data anchor with empty content hash
	_, err = s.server.AnchorByHash(s.ctx, &data.QueryAnchorByHashRequest{})
	require.EqualError(t, err, "content hash cannot be empty: invalid request")

	// query data anchor with invalid content hash
	_, err = s.server.AnchorByHash(s.ctx, &data.QueryAnchorByHashRequest{
		ContentHash: &data.ContentHash{},
	})
	require.EqualError(t, err, "invalid data.ContentHash: invalid request")

	// query data anchor with content hash that has not been anchored
	_, err = s.server.AnchorByHash(s.ctx, &data.QueryAnchorByHashRequest{
		ContentHash: &data.ContentHash{Graph: &data.ContentHash_Graph{
			Hash:                      bytes.Repeat([]byte{1}, 32),
			DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		}},
	})
	require.EqualError(t, err, "data record with content hash: not found")
}
