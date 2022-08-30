package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRetire struct {
	t   gocuke.TestingT
	msg *MsgRetire
	err error
}

func TestMsgRetire(t *testing.T) {
	gocuke.NewRunner(t, &msgRetire{}).Path("./features/msg_retire.feature").Run()
}

func (s *msgRetire) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRetire) TheMessage(a gocuke.DocString) {
	s.msg = &MsgRetire{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRetire) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRetire) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgRetire) ExpectNoError() {
	require.NoError(s.t, s.err)
}
