package basket

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgTakeSuite struct {
	t   gocuke.TestingT
	msg *MsgTake
	err error
}

func TestMsgTake(t *testing.T) {
	gocuke.NewRunner(t, &msgTakeSuite{}).Path("./features/msg_take.feature").Run()
}

func (s *msgTakeSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgTakeSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgTake{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgTakeSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgTakeSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgTakeSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
