package core

import (
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgCancel struct {
	t   gocuke.TestingT
	msg *MsgCancel
	err error
}

func TestMsgCancel(t *testing.T) {
	gocuke.NewRunner(t, &msgCancel{}).Path("./features/msg_cancel.feature").Run()
}

func (s *msgCancel) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgCancel) TheMessage(a gocuke.DocString) {
	s.msg = &MsgCancel{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCancel) TheMessageIsValidated() {
	s.checkAndSetMockValues()

	s.err = s.msg.ValidateBasic()
}

func (s *msgCancel) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgCancel) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCancel) checkAndSetMockValues() {
	if strings.Contains(s.msg.Reason, "[mock-string-513]") {
		s.msg.Reason = strings.Repeat("x", 513)
	}
}
