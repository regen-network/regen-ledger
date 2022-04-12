package data

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgDefineResolverSuite struct {
	t   gocuke.TestingT
	msg *MsgDefineResolver
	err error
}

func TestMsgDefineResolver(t *testing.T) {
	runner := gocuke.NewRunner(t, &msgDefineResolverSuite{}).Path("./features/msg_define_resolver.feature")
	runner.Step(`the message "((?:[^\"]|\")*)"`, (*msgDefineResolverSuite).TheMessage)
	runner.Run()
}

func (s *msgDefineResolverSuite) Before(t gocuke.TestingT) {
	s.t = t
	s.msg = &MsgDefineResolver{}
}

func (s *msgDefineResolverSuite) TheMessage(a gocuke.DocString) {
	err := json.Unmarshal([]byte(a.Content), &s.msg)
	require.NoError(s.t, err)
}

func (s *msgDefineResolverSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgDefineResolverSuite) ExpectTheError(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.EqualError(s.t, s.err, a)
	}
}
