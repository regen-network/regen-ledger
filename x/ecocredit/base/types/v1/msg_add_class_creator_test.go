package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAddClassCreator struct {
	t         gocuke.TestingT
	msg       *MsgAddClassCreator
	err       error
	signBytes string
}

func TestMsgAddClassCreator(t *testing.T) {
	gocuke.NewRunner(t, &msgAddClassCreator{}).Path("./features/msg_add_class_creator.feature").Run()
}

func (s *msgAddClassCreator) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgAddClassCreator) TheMessage(a gocuke.DocString) {
	s.msg = &MsgAddClassCreator{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgAddClassCreator) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAddClassCreator) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgAddClassCreator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgAddClassCreator) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgAddClassCreator) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
