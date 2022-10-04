package v1

import (
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRemoveAllowedBridgeChain struct {
	t         gocuke.TestingT
	msg       *MsgRemoveAllowedBridgeChain
	err       error
	signBytes string
}

func TestMsgRemoveAllowedBridgeChain(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveAllowedBridgeChain{}).Path("./features/msg_remove_allowed_bridge_chain.feature").Run()
}

func (s *msgRemoveAllowedBridgeChain) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRemoveAllowedBridgeChain) TheMessage(a gocuke.DocString) {
	s.msg = &MsgRemoveAllowedBridgeChain{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRemoveAllowedBridgeChain) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRemoveAllowedBridgeChain) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgRemoveAllowedBridgeChain) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveAllowedBridgeChain) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgRemoveAllowedBridgeChain) ExpectTheSignBytes(expected gocuke.DocString) {
	require.Equal(s.t, strings.TrimSpace(expected.Content), s.signBytes)
}
