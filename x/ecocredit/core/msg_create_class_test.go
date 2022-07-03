package core

import (
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgCreateClass struct {
	t   gocuke.TestingT
	msg *MsgCreateClass
	err error
}

func TestMsgCreateClass(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateClass{}).Path("./features/msg_create_class.feature").Run()
}

func (s *msgCreateClass) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgCreateClass) TheMessage(a gocuke.DocString) {
	s.msg = &MsgCreateClass{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCreateClass) TheMessageIsValidated() {
	s.checkAndSetMockValues()

	s.err = s.msg.ValidateBasic()
}

func (s *msgCreateClass) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgCreateClass) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreateClass) checkAndSetMockValues() {
	if strings.Contains(s.msg.Metadata, "[mock-string-257]") {
		s.msg.Metadata = strings.Repeat("x", 257)
	}
}
