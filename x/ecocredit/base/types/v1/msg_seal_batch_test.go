package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgSealBatch struct {
	t         gocuke.TestingT
	msg       *MsgSealBatch
	err       error
	signBytes string
}

func TestMsgSealBatch(t *testing.T) {
	gocuke.NewRunner(t, &msgSealBatch{}).Path("./features/msg_seal_batch.feature").Run()
}

func (s *msgSealBatch) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgSealBatch) TheMessage(a gocuke.DocString) {
	s.msg = &MsgSealBatch{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgSealBatch) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgSealBatch) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgSealBatch) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgSealBatch) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgSealBatch) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
