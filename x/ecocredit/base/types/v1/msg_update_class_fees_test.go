package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateClassFee struct {
	t   gocuke.TestingT
	msg *MsgUpdateClassFee
	err error
}

func TestMsgUpdateClassFee(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateClassFee{}).Path("./features/msg_update_class_fee.feature").Run()
}

func (s *msgUpdateClassFee) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateClassFee) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateClassFee{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateClassFee) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateClassFee) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateClassFee) ExpectNoError() {
	require.NoError(s.t, s.err)
}
