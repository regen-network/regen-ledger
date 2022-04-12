package data

import (
	"encoding/json"
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
	runner := gocuke.NewRunner(t, &msgAttestSuite{}).Path("./features/msg_attest.feature")
	runner.Step(`the message "((?:[^\"]|\")*)"`, (*msgAttestSuite).TheMessage)
	runner.Run()
}

func (s *msgAttestSuite) Before(t gocuke.TestingT) {
	s.t = t
	s.msg = &MsgAttest{}
}

func (s *msgAttestSuite) TheMessage(a gocuke.DocString) {
	err := json.Unmarshal([]byte(a.Content), &s.msg)
	require.NoError(s.t, err)
}

func (s *msgAttestSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAttestSuite) ExpectTheError(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.EqualError(s.t, s.err, a)
	}
}
