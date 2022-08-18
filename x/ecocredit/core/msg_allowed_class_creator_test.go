package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAllowedClassCreator struct {
	t   gocuke.TestingT
	msg *MsgAllowedClassCreator
	err error
}

func TestMsgAllowedClassCreator(t *testing.T) {
	gocuke.NewRunner(t, &msgAllowedClassCreator{}).Path("./features/msg_allowed_class_creator.feature").Run()
}

func (s *msgAllowedClassCreator) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgAllowedClassCreator) TheMessage(a gocuke.DocString) {
	s.msg = &MsgAllowedClassCreator{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgAllowedClassCreator) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAllowedClassCreator) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgAllowedClassCreator) ExpectNoError() {
	require.NoError(s.t, s.err)
}
