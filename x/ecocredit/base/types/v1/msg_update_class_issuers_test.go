package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateClassIssuers struct {
	t   gocuke.TestingT
	msg *MsgUpdateClassIssuers
	err error
}

func TestMsgUpdateClassIssuers(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateClassIssuers{}).Path("./features/msg_update_class_issuers.feature").Run()
}

func (s *msgUpdateClassIssuers) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateClassIssuers) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateClassIssuers{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateClassIssuers) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateClassIssuers) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateClassIssuers) ExpectNoError() {
	require.NoError(s.t, s.err)
}
