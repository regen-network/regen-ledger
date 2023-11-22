package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgPutSuite struct {
	t         gocuke.TestingT
	msg       *MsgPut
	err       error
	signBytes string
}

func TestMsgPut(t *testing.T) {
	gocuke.NewRunner(t, &msgPutSuite{}).Path("./features/msg_put.feature").Run()
}

func (s *msgPutSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgPutSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgPut{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgPutSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgPutSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgPutSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgPutSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgPutSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
