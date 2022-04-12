package data

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRegisterResolverSuite struct {
	t   gocuke.TestingT
	msg *MsgRegisterResolver
	err error
}

func TestMsgRegisterResolver(t *testing.T) {
	runner := gocuke.NewRunner(t, &msgRegisterResolverSuite{}).Path("./features/msg_register_resolver.feature")
	runner.Step(`a message of "((?:[^\"]|\")*)"`, (*msgRegisterResolverSuite).AMessageOf)
	runner.Run()
}

func (s *msgRegisterResolverSuite) Before(t gocuke.TestingT) {
	s.t = t
	s.msg = &MsgRegisterResolver{}
}

func (s *msgRegisterResolverSuite) AMessageOf(a gocuke.DocString) {
	err := json.Unmarshal([]byte(a.Content), &s.msg)
	require.NoError(s.t, err)
}

func (s *msgRegisterResolverSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRegisterResolverSuite) AnErrorOf(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.EqualError(s.t, s.err, a)
	}
}
