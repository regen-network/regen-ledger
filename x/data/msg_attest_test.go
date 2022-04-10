package data

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAttestSuite struct {
	t   gocuke.TestingT
	msg *MsgAttest
	err error
}

func TestMsgAttest(t *testing.T) {
	gocuke.NewRunner(t, &msgAttestSuite{}).Path("./features/msg_attest.feature").Run()
}

func (s *msgAttestSuite) Before(t gocuke.TestingT) {
	s.t = t
	s.msg = &MsgAttest{}
}

func (s *msgAttestSuite) AValidAttestor() {
	s.msg.Attestor = "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
}

func (s *msgAttestSuite) AValidContentHash() {
	s.msg.Hashes = []*ContentHash_Graph{
		{
			Hash:                      make([]byte, 32),
			DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		},
	}
}

func (s *msgAttestSuite) AnAttestorOf(a string) {
	s.msg.Attestor = a
}

func (s *msgAttestSuite) AnEmptyListOfContentHashes() {
	s.msg.Hashes = []*ContentHash_Graph{}
}

func (s *msgAttestSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAttestSuite) AnErrorOf(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.EqualError(s.t, s.err, a)
	}
}
