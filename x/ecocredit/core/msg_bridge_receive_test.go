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

type msgBridgeReceive struct {
	t   gocuke.TestingT
	msg *MsgBridgeReceive
	err error
}

func TestMsgBridgeReceive(t *testing.T) {
	gocuke.NewRunner(t, &msgBridgeReceive{}).Path("./features/msg_bridge_receive.feature").Run()
}

func (s *msgBridgeReceive) Before(t gocuke.TestingT) {
	s.t = t

	// TODO: move to init function in the root directory of the module #1243
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
}

func (s *msgBridgeReceive) TheMessage(a gocuke.DocString) {
	s.msg = &MsgBridgeReceive{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgBridgeReceive) AProjectReferenceIdWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Project = &MsgBridgeReceive_Project{}
	s.msg.Project.ReferenceId = strings.Repeat("x", int(length))
}

func (s *msgBridgeReceive) ProjectMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Project.Metadata = strings.Repeat("x", int(length))
}

func (s *msgBridgeReceive) BatchMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Batch.Metadata = strings.Repeat("x", int(length))
}

func (s *msgBridgeReceive) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgBridgeReceive) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgBridgeReceive) ExpectNoError() {
	require.NoError(s.t, s.err)
}
