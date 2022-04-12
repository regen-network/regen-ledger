package data

import (
	"encoding/json"
	"strconv"
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
	runner.Step(`data of "((?:[^\"]|\")*)"`, (*msgRegisterResolverSuite).ContentHashesOf)
	runner.Run()
}

func (s *msgRegisterResolverSuite) Before(t gocuke.TestingT) {
	s.t = t
	s.msg = &MsgRegisterResolver{}
}

func (s *msgRegisterResolverSuite) AManagerOf(a string) {
	s.msg.Manager = a
}

func (s *msgRegisterResolverSuite) AResolverIdOf(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	s.msg.ResolverId = id
}

func (s *msgRegisterResolverSuite) ContentHashesOf(a string) {
	if a == "" {
		s.msg.Data = nil
	} else {
		var data []*ContentHash
		err := json.Unmarshal([]byte(a), &data)
		require.NoError(s.t, err)

		s.msg.Data = data
	}
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
