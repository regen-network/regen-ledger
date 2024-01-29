package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgBurnRegen struct {
	gocuke.TestingT
	msg *MsgBurnRegen
	err error
}

func TestMsgBurnRegen(t *testing.T) {
	gocuke.NewRunner(t, &msgBurnRegen{}).Path("./features/msg_burn_regen.feature").Run()
}

func (s *msgBurnRegen) TheMessage(a gocuke.DocString) {
	s.msg = &MsgBurnRegen{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s, err)
}

func (s *msgBurnRegen) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgBurnRegen) ExpectNoError() {
	require.NoError(s, s.err)
}

func (s *msgBurnRegen) ExpectErrorContains(a string) {
	require.ErrorContains(s, s.err, a)
}
