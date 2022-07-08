package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgMintBatchCredits struct {
	t   gocuke.TestingT
	msg *MsgMintBatchCredits
	err error
}

func TestMsgMintBatchCredits(t *testing.T) {
	gocuke.NewRunner(t, &msgMintBatchCredits{}).Path("./features/msg_mint_batch_credits.feature").Run()
}

func (s *msgMintBatchCredits) Before(t gocuke.TestingT) {
	s.t = t

	// TODO: move to init function in the root directory of the module #1243
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
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
