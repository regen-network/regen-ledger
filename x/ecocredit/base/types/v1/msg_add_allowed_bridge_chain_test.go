package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAddAllowedBridgeChain struct {
	t   gocuke.TestingT
	msg *MsgAddAllowedBridgeChain
	err error
}

func TestMsgAddAllowedBridgeChain(t *testing.T) {
	gocuke.NewRunner(t, &msgAddAllowedBridgeChain{}).Path("./features/msg_add_allowed_bridge_chain.feature.feature").Run()
}

func (s *msgAddAllowedBridgeChain) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgAddAllowedBridgeChain) TheMessage(a gocuke.DocString) {
	s.msg = &MsgAddAllowedBridgeChain{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgAddAllowedBridgeChain) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAddAllowedBridgeChain) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgAddAllowedBridgeChain) ExpectNoError() {
	require.NoError(s.t, s.err)
}
