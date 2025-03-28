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

type msgCreateClass struct {
	t         gocuke.TestingT
	msg       *MsgCreateClass
	err       error
	signBytes string
}

func TestMsgCreateClass(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateClass{}).Path("./features/msg_create_class.feature").Run()
}

func (s *msgCreateClass) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgCreateClass) TheMessage(a gocuke.DocString) {
	s.msg = &MsgCreateClass{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCreateClass) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Metadata = strings.Repeat("x", int(length))
}

func (s *msgCreateClass) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCreateClass) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgCreateClass) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreateClass) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgCreateClass) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
