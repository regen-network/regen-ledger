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
	runner.Step(`hashes of "((?:[^\"]|\")*)"`, (*msgAttestSuite).AGraphContentHashOf)
	runner.Run()
}

func (s *msgAttestSuite) Before(t gocuke.TestingT) {
	s.t = t
	s.msg = &MsgAttest{}
}

func (s *msgAttestSuite) AnAttestorOf(a string) {
	s.msg.Attestor = a
}

func (s *msgAttestSuite) AGraphContentHashOf(a string) {
	if a == "" {
		s.msg.Hashes = nil
	} else {
		var hashes []*ContentHash_Graph
		err := json.Unmarshal([]byte(a), &hashes)
		require.NoError(s.t, err)

		s.msg.Hashes = hashes
	}
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
