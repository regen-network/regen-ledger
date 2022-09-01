package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgToggleClassAllowlist struct {
	t   gocuke.TestingT
	msg *MsgToggleCreditClassAllowlist
	err error
}

func TestMsgToggleClassAllowlist(t *testing.T) {
	gocuke.NewRunner(t, &msgToggleClassAllowlist{}).Path("./features/msg_toggle_class_allowlist.feature").Run()
}

func (s *msgToggleClassAllowlist) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgToggleClassAllowlist) TheMessage(a gocuke.DocString) {
	s.msg = &MsgToggleCreditClassAllowlist{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgToggleClassAllowlist) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgToggleClassAllowlist) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgToggleClassAllowlist) ExpectNoError() {
	require.NoError(s.t, s.err)
}
