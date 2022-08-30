package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateClassFees struct {
	t   gocuke.TestingT
	msg *MsgUpdateClassFees
	err error
}

func TestMsgUpdateClassFees(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateClassFees{}).Path("./features/msg_update_class_fees.feature").Run()
}

func (s *msgUpdateClassFees) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateClassFees) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateClassFees{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateClassFees) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateClassFees) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateClassFees) ExpectNoError() {
	require.NoError(s.t, s.err)
}
