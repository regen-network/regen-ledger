package v1

import (
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgMintBatchCredits struct {
	t         gocuke.TestingT
	msg       *MsgMintBatchCredits
	err       error
	signBytes string
}

func TestMsgMintBatchCredits(t *testing.T) {
	gocuke.NewRunner(t, &msgMintBatchCredits{}).Path("./features/msg_mint_batch_credits.feature").Run()
}

func (s *msgMintBatchCredits) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgMintBatchCredits) TheMessage(a gocuke.DocString) {
	s.msg = &MsgMintBatchCredits{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgMintBatchCredits) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgMintBatchCredits) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgMintBatchCredits) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgMintBatchCredits) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgMintBatchCredits) ExpectTheSignBytes(expected gocuke.DocString) {
	require.Equal(s.t, strings.TrimSpace(expected.Content), s.signBytes)
}
