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

type msgCreateBatch struct {
	t   gocuke.TestingT
	msg *MsgCreateBatch
	err error
}

func TestMsgCreateBatch(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateBatch{}).Path("./features/msg_create_batch.feature").Run()
}

func (s *msgCreateBatch) Before(t gocuke.TestingT) {
	s.t = t

	// TODO: move to init function in the root directory of the module #1243
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
}

func (s *msgCreateBatch) TheMessage(a gocuke.DocString) {
	s.msg = &MsgCreateBatch{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCreateBatch) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Metadata = strings.Repeat("x", int(length))
}

func (s *msgCreateBatch) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCreateBatch) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgCreateBatch) ExpectNoError() {
	require.NoError(s.t, s.err)
}
