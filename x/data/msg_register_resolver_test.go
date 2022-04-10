package data

import (
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
	gocuke.NewRunner(t, &msgRegisterResolverSuite{}).Path("./features/msg_register_resolver.feature").Run()
}

func (s *msgRegisterResolverSuite) Before(t gocuke.TestingT) {
	s.t = t
	s.msg = &MsgRegisterResolver{}
}

func (s *msgRegisterResolverSuite) AValidManager() {
	s.msg.Manager = "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27"
}

func (s *msgRegisterResolverSuite) AValidResolverId() {
	s.msg.ResolverId = 1
}

func (s *msgRegisterResolverSuite) AValidListOfData() {
	s.msg.Data = []*ContentHash{
		{
			Raw: &ContentHash_Raw{
				Hash:            make([]byte, 32),
				DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			},
		},
	}
}

func (s *msgRegisterResolverSuite) AManagerOf(a string) {
	s.msg.Manager = a
}

func (s *msgRegisterResolverSuite) AResolverIdOf(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	s.msg.ResolverId = id
}

func (s *msgRegisterResolverSuite) AnEmptyListOfData() {
	s.msg.Data = []*ContentHash{}
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
