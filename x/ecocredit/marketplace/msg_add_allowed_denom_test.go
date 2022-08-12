package marketplace

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAddAllowedDenomSuite struct {
	t   gocuke.TestingT
	msg *MsgAddAllowedDenom
	err error
}

func TestMsgAddAllowedDenomSuite(t *testing.T) {
	gocuke.NewRunner(t, &msgAddAllowedDenomSuite{}).Path("./features/msg_add_allowed_denom.feature").Run()
}

func (s *msgAddAllowedDenomSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgAddAllowedDenomSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgAddAllowedDenom{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgAddAllowedDenomSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAddAllowedDenomSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgAddAllowedDenomSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
