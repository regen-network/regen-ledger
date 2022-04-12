package server

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/data"
)

type attestSuite struct {
	*baseSuite
	alice sdk.AccAddress
	ch    *data.ContentHash
	err   error
}

func TestAttest(t *testing.T) {
	runner := gocuke.NewRunner(t, &attestSuite{}).Path("./features/attest.feature")
	runner.Step(`the content hash "((?:[^\"]|\")*)"`, (*attestSuite).TheContentHash)
	runner.Run()
}

func (s *attestSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *attestSuite) AliceIsTheAttestor() {
	s.alice = s.addrs[0]
}

func (s *attestSuite) TheContentHash(a gocuke.DocString) {
	err := json.Unmarshal([]byte(a.Content), &s.ch)
	require.NoError(s.t, err)
}

func (s *attestSuite) AliceHasAnchoredTheDataAtBlockTime(a string) {
	blockTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender: s.alice.String(),
		Hash:   s.ch,
	})
}

func (s *attestSuite) AliceHasAttestedToTheDataAtBlockTime(a string) {
	blockTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Attest(s.ctx, &data.MsgAttest{
		Attestor: s.alice.String(),
		Hashes:   []*data.ContentHash_Graph{s.ch.GetGraph()},
	})
}

func (s *attestSuite) AliceAttemptsToAttestToTheDataAtBlockTime(a string) {
	blockTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Attest(s.ctx, &data.MsgAttest{
		Attestor: s.alice.String(),
		Hashes:   []*data.ContentHash_Graph{s.ch.GetGraph()},
	})
}

func (s *attestSuite) TheDataAnchorEntryExistsAndTheTimestampIsEqualTo(a string) {
	anchorTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, dataId.Id)
	require.NoError(s.t, err)
	require.Equal(s.t, anchorTime, dataAnchor.Timestamp.AsTime())
}

func (s *attestSuite) TheDataAttestorEntryExistsAndTheTimestampIsEqualTo(a string) {
	attestTime, err := time.Parse("2006-01-02", a)
	require.NoError(s.t, err)

	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	dataAttestor, err := s.server.stateStore.DataAttestorTable().Get(s.ctx, dataId.Id, s.alice)
	require.NoError(s.t, err)
	require.Equal(s.t, attestTime, dataAttestor.Timestamp.AsTime())
}
