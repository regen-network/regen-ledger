package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRemoveClassCreator struct {
	t         gocuke.TestingT
	msg       *MsgRemoveClassCreator
	err       error
	signBytes string
}

func TestMsgRemoveClassCreators(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveClassCreator{}).Path("./features/msg_remove_class_creator.feature").Run()
}

func (s *msgRemoveClassCreator) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRemoveClassCreator) TheMessage(a gocuke.DocString) {
	s.msg = &MsgRemoveClassCreator{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRemoveClassCreator) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRemoveClassCreator) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgRemoveClassCreator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveClassCreator) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgRemoveClassCreator) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
