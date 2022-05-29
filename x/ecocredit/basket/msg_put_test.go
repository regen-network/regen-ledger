package basket

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgPutSuite struct {
	t   gocuke.TestingT
	msg *MsgPut
	err error
}

func TestMsgPut(t *testing.T) {
	gocuke.NewRunner(t, &msgPutSuite{}).Path("./features/msg_put.feature").Run()
}

func (s *msgPutSuite) Before(t gocuke.TestingT) {
	s.t = t

	// set the denom regex for basket coins
	sdk.SetCoinDenomRegex(func() string {
		return RegexBasketDenom
	})
}

func (s *msgPutSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgPut{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgPutSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgPutSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgPutSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
