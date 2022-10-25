package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgSetClassCreatorAllowlist struct {
	t         gocuke.TestingT
	msg       *MsgSetClassCreatorAllowlist
	err       error
	signBytes string
}

func TestMsgSetClassCreatorAllowlist(t *testing.T) {
	gocuke.NewRunner(t, &msgSetClassCreatorAllowlist{}).Path("./features/msg_set_class_creator_allowlist.feature").Run()
}

func (s *msgSetClassCreatorAllowlist) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgSetClassCreatorAllowlist) TheMessage(a gocuke.DocString) {
	s.msg = &MsgSetClassCreatorAllowlist{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgSetClassCreatorAllowlist) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgSetClassCreatorAllowlist) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgSetClassCreatorAllowlist) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgSetClassCreatorAllowlist) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgSetClassCreatorAllowlist) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
