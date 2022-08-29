package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgBridge struct {
	t   gocuke.TestingT
	msg *MsgBridge
	err error
}

func TestMsgBridge(t *testing.T) {
	gocuke.NewRunner(t, &msgBridge{}).Path("./features/msg_bridge.feature").Run()
}

func (s *msgBridge) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgBridge) TheMessage(a gocuke.DocString) {
	s.msg = &MsgBridge{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgBridge) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgBridge) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgBridge) ExpectNoError() {
	require.NoError(s.t, s.err)
}
