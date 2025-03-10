package v1

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgCancel struct {
	t         gocuke.TestingT
	msg       *MsgCancel
	err       error
	signBytes string
}

func TestMsgCancel(t *testing.T) {
	gocuke.NewRunner(t, &msgCancel{}).Path("./features/msg_cancel.feature").Run()
}

func (s *msgCancel) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgCancel) TheMessage(a gocuke.DocString) {
	s.msg = &MsgCancel{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCancel) AReasonWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Reason = strings.Repeat("x", int(length))
}

func (s *msgCancel) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCancel) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgCancel) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCancel) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgCancel) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
