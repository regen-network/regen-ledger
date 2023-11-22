package data

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAnchorSuite struct {
	t         gocuke.TestingT
	msg       *MsgAnchor
	err       error
	signBytes string
}

func TestMsgAnchor(t *testing.T) {
	gocuke.NewRunner(t, &msgAnchorSuite{}).Path("./features/msg_anchor.feature").Run()
}

func (s *msgAnchorSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgAnchorSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgAnchor{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgAnchorSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAnchorSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgAnchorSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgAnchorSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgAnchorSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
