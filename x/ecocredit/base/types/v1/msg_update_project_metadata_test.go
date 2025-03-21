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

type msgUpdateProjectMetadata struct {
	t         gocuke.TestingT
	msg       *MsgUpdateProjectMetadata
	err       error
	signBytes string
}

func TestMsgUpdateProjectMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateProjectMetadata{}).Path("./features/msg_update_project_metadata.feature").Run()
}

func (s *msgUpdateProjectMetadata) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateProjectMetadata) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateProjectMetadata{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateProjectMetadata) NewMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.NewMetadata = strings.Repeat("x", int(length))
}

func (s *msgUpdateProjectMetadata) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateProjectMetadata) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateProjectMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateProjectMetadata) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgUpdateProjectMetadata) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
