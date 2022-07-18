package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgSealBatch struct {
	t   gocuke.TestingT
	msg *MsgSealBatch
	err error
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
