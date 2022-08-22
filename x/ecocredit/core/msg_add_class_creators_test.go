package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAddClassCreators struct {
	t   gocuke.TestingT
	msg *MsgAddClassCreators
	err error
}

func TestMsgAddClassCreators(t *testing.T) {
	gocuke.NewRunner(t, &msgAddClassCreators{}).Path("./features/msg_add_class_creators.feature").Run()
}

func (s *msgAddClassCreators) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgAddClassCreators) TheMessage(a gocuke.DocString) {
	s.msg = &MsgAddClassCreators{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgAddClassCreators) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAddClassCreators) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgAddClassCreators) ExpectNoError() {
	require.NoError(s.t, s.err)
}
