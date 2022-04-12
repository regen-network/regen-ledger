package data

import (
	"encoding/json"
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
	runner := gocuke.NewRunner(t, &msgAnchorSuite{}).Path("./features/msg_anchor.feature")
	runner.Step(`a message of "((?:[^\"]|\")*)"`, (*msgAnchorSuite).AMessageOf)
	runner.Run()
}

func (s *msgAnchorSuite) Before(t gocuke.TestingT) {
	s.t = t
	s.msg = &MsgAnchor{}
}

func (s *msgAnchorSuite) AMessageOf(a gocuke.DocString) {
	err := json.Unmarshal([]byte(a.Content), &s.msg)
	require.NoError(s.t, err)
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
