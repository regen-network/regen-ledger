package data

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAnchorSuite struct {
	t   gocuke.TestingT
	msg *MsgAnchor
	err error
}

func TestAnchorMsg(t *testing.T) {
	gocuke.NewRunner(t, &msgAnchorSuite{}).Path("./features/msg_anchor.feature").Run()
}

func (s *msgAnchorSuite) Before(t gocuke.TestingT) {
	s.t = t
	s.msg = &MsgAnchor{}
}

func (s *msgAnchorSuite) ASenderOf(a string) {
	s.msg.Sender = a
}

func (s *msgAnchorSuite) AContentHashOf(a string) {
	s.msg.Hash = &ContentHash{
		Raw: &ContentHash_Raw{
			Hash:            make([]byte, 32),
			DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		},
	}
}

func (s *msgAnchorSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAnchorSuite) AnErrorOf(a string) {
	require.EqualError(s.t, s.err, a)
}
