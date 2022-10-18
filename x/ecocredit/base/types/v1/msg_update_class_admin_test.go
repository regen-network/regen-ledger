package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateClassAdmin struct {
	t         gocuke.TestingT
	msg       *MsgUpdateClassAdmin
	err       error
	signBytes string
}

func TestMsgUpdateClassAdmin(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateClassAdmin{}).Path("./features/msg_update_class_admin.feature").Run()
}

func (s *msgUpdateClassAdmin) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateClassAdmin) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateClassAdmin{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateClassAdmin) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateClassAdmin) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateClassAdmin) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateClassAdmin) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgUpdateClassAdmin) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
