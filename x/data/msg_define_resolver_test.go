package data

import (
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
	gocuke.NewRunner(t, &msgDefineResolverSuite{}).Path("./features/msg_define_resolver.feature").Run()
}

func (s *msgDefineResolverSuite) Before(t gocuke.TestingT) {
	s.t = t
	s.msg = &MsgDefineResolver{}
}

func (s *msgDefineResolverSuite) AManagerOf(a string) {
	s.msg.Manager = a
}

func (s *msgDefineResolverSuite) AResolverUrlOf(a string) {
	s.msg.ResolverUrl = a
}

func (s *msgDefineResolverSuite) AValidManagerAddress() {
	s.msg.Manager = "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
}

func (s *msgDefineResolverSuite) AValidResolverUrl() {
	s.msg.ResolverUrl = "https://foo.bar"
}

func (s *msgDefineResolverSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgDefineResolverSuite) AnErrorOf(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.EqualError(s.t, s.err, a)
	}
}
