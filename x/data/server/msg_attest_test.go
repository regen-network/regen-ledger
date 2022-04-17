package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/regen-network/regen-ledger/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/data"
)

type attestSuite struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	ch    *data.ContentHash
	err   error
}

func TestAttest(t *testing.T) {
	runner := gocuke.NewRunner(t, &attestSuite{}).Path("./features/attest.feature")
	runner.Step(`^the\s+content\s+hash\s+"((?:[^\"]|\")*)"`, (*attestSuite).TheContentHash)
	runner.Run()
}

func (s *attestSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.bob = s.addrs[1]
}

func (s *attestSuite) TheContentHash(a gocuke.DocString) {
	s.ch = &data.ContentHash{}
	err := jsonpb.UnmarshalString(a.Content, s.ch)
	require.NoError(s.t, err)
}

func (s *attestSuite) AliceHasAnchoredTheDataAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Anchor(s.ctx, &data.MsgAnchor{
		Sender:      s.alice.String(),
		ContentHash: s.ch,
	})
}

func (s *attestSuite) AliceHasAttestedToTheDataAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.alice.String(),
		ContentHashes: []*data.ContentHash_Graph{s.ch.Graph},
	})
}

func (s *attestSuite) AliceAttemptsToAttestToTheDataAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.alice.String(),
		ContentHashes: []*data.ContentHash_Graph{s.ch.Graph},
	})
}

func (s *attestSuite) BobAttemptsToAttestToTheDataAtBlockTime(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.ctx = sdk.WrapSDKContext(s.sdkCtx.WithBlockTime(blockTime))

	_, s.err = s.server.Attest(s.ctx, &data.MsgAttest{
		Attestor:      s.bob.String(),
		ContentHashes: []*data.ContentHash_Graph{s.ch.Graph},
	})
}

func (s *attestSuite) TheAnchorEntryExistsWithTimestamp(a string) {
	anchorTime, err := types.ParseDate("anchor timestamp", a)
	require.NoError(s.t, err)

	dataId := s.getDataId()

	dataAnchor, err := s.server.stateStore.DataAnchorTable().Get(s.ctx, dataId)
	require.NoError(s.t, err)
	require.Equal(s.t, anchorTime, dataAnchor.Timestamp.AsTime())
}

func (s *attestSuite) TheAttestorEntryForAliceExistsWithTimestamp(a string) {
	attestTime, err := types.ParseDate("attest timestamp", a)
	require.NoError(s.t, err)

	dataId := s.getDataId()

	dataAttestor, err := s.server.stateStore.DataAttestorTable().Get(s.ctx, dataId, s.alice)
	require.NoError(s.t, err)
	require.Equal(s.t, attestTime, dataAttestor.Timestamp.AsTime())
}

func (s *attestSuite) TheAttestorEntryForBobExistsWithTimestamp(a string) {
	attestTime, err := types.ParseDate("attest timestamp", a)
	require.NoError(s.t, err)

	dataId := s.getDataId()

	dataAttestor, err := s.server.stateStore.DataAttestorTable().Get(s.ctx, dataId, s.bob)
	require.NoError(s.t, err)
	require.Equal(s.t, attestTime, dataAttestor.Timestamp.AsTime())
}

func (s *attestSuite) getDataId() []byte {
	iri, err := s.ch.ToIRI()
	require.NoError(s.t, err)
	require.NotNil(s.t, iri)

	dataId, err := s.server.stateStore.DataIDTable().GetByIri(s.ctx, iri)
	require.NoError(s.t, err)
	require.NotNil(s.t, dataId)

	return dataId.Id
}
