package data

import (
	"strconv"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAnchorSuite struct {
	t   gocuke.TestingT
	msg *MsgAnchor
	err error
}

func TestMsgAnchor(t *testing.T) {
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

func (s *msgAnchorSuite) AValidSenderAddress() {
	s.msg.Sender = "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
}

func (s *msgAnchorSuite) AValidContentHash() {
	s.msg.Hash = &ContentHash{
		Raw: &ContentHash_Raw{
			Hash:            make([]byte, 32),
			DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		},
	}
}

func (s *msgAnchorSuite) ARawContentHashOfBytesLengthAndDigestAlgorithm(a string, b string) {
	length, err := strconv.Atoi(a)
	require.NoError(s.t, err)

	digest, err := strconv.Atoi(b)
	require.NoError(s.t, err)

	s.msg.Hash = &ContentHash{
		Raw: &ContentHash_Raw{
			Hash:            make([]byte, length),
			DigestAlgorithm: DigestAlgorithm(digest),
		},
	}
}

func (s *msgAnchorSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAnchorSuite) AnErrorOf(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.EqualError(s.t, s.err, a)
	}
}
