package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgCreateSuite struct {
	t         gocuke.TestingT
	msg       *MsgCreate
	err       error
	signBytes string
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

func (s *msgCreateSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgCreateSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
