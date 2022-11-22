package v1

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateBatchMetadata struct {
	t         gocuke.TestingT
	msg       *MsgUpdateBatchMetadata
	err       error
	signBytes string
}

func TestMsgUpdateBatchMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateBatchMetadata{}).Path("./features/msg_update_batch_metadata.feature").Run()
}

func (s *msgUpdateBatchMetadata) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateBatchMetadata) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateBatchMetadata{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateBatchMetadata) NewMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.NewMetadata = strings.Repeat("x", int(length))
}

func (s *msgUpdateBatchMetadata) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateBatchMetadata) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateBatchMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateBatchMetadata) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgUpdateBatchMetadata) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
