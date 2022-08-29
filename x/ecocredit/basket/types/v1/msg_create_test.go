package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgCreateSuite struct {
	t   gocuke.TestingT
	msg *MsgCreate
	err error
}

func TestMsgCreate(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateSuite{}).Path("./features/msg_create.feature").Run()
}

func (s *msgCreateSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgCreateSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgCreate{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCreateSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCreateSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgCreateSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
