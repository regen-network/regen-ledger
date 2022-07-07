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

type msgUpdateProjectMetadata struct {
	t   gocuke.TestingT
	msg *MsgUpdateProjectMetadata
	err error
}

func TestMsgUpdateProjectMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateProjectMetadata{}).Path("./features/msg_update_project_metadata.feature").Run()
}

func (s *msgUpdateProjectMetadata) Before(t gocuke.TestingT) {
	s.t = t

	// TODO: move to init function in the root directory of the module #1243
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
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
