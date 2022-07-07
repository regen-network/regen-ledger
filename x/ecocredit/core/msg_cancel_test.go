package core

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgCancel struct {
	t   gocuke.TestingT
	msg *MsgCancel
	err error
}

func TestMsgCancel(t *testing.T) {
	gocuke.NewRunner(t, &msgCancel{}).Path("./features/msg_cancel.feature").Run()
}

func (s *msgCancel) Before(t gocuke.TestingT) {
	s.t = t

	// TODO: move to init function in the root directory of the module #1243
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
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
