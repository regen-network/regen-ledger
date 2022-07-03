package core

import (
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgBridgeReceive struct {
	t   gocuke.TestingT
	msg *MsgBridgeReceive
	err error
}

func TestMsgBridgeReceive(t *testing.T) {
	gocuke.NewRunner(t, &msgBridgeReceive{}).Path("./features/msg_bridge_receive.feature").Run()
}

func (s *msgBridgeReceive) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgBridgeReceive) TheMessage(a gocuke.DocString) {
	s.msg = &MsgBridgeReceive{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgBridgeReceive) TheMessageIsValidated() {
	s.checkAndSetMockValues()

	s.err = s.msg.ValidateBasic()
}

func (s *msgBridgeReceive) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgBridgeReceive) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgBridgeReceive) checkAndSetMockValues() {
	if s.msg.Project != nil {
		if strings.Contains(s.msg.Project.ReferenceId, "[mock-string-33]") {
			s.msg.Project.ReferenceId = strings.Repeat("x", 33)
		}
		if strings.Contains(s.msg.Project.Metadata, "[mock-string-257]") {
			s.msg.Project.Metadata = strings.Repeat("x", 257)
		}
	}
	if s.msg.Batch != nil {
		if strings.Contains(s.msg.Batch.Metadata, "[mock-string-257]") {
			s.msg.Batch.Metadata = strings.Repeat("x", 257)
		}
	}
}
