package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgTakeSuite struct {
	t         gocuke.TestingT
	msg       *MsgTake
	err       error
	signBytes string
}

func TestMsgTake(t *testing.T) {
	gocuke.NewRunner(t, &msgTakeSuite{}).Path("./features/msg_take.feature").Run()
}

func (s *msgTakeSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgTakeSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgTake{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgTakeSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgTakeSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgTakeSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgTakeSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgTakeSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
