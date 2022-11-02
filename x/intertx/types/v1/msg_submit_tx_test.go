package v1

import (
	"testing"

	"github.com/regen-network/gocuke"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type submitTxSuite struct {
	t     gocuke.TestingT
	msg   *MsgSubmitTx
	codec *codec.ProtoCodec
	err   error
}

func TestMsgSubmitTx(t *testing.T) {
	gocuke.NewRunner(t, &submitTxSuite{}).Path("./features/msg_submit_tx.feature").Run()
}

func (s *submitTxSuite) Before(t gocuke.TestingT) {
	s.t = t

	// register bank interfaces
	ir := types.NewInterfaceRegistry()
	RegisterTypes(ir)
	banktypes.RegisterInterfaces(ir)
	s.codec = codec.NewProtoCodec(ir)
}

func (s *submitTxSuite) ExpectTheError(a string) {
	assert.ErrorContains(s.t, s.err, a)
}

func (s *submitTxSuite) TheMessage(a gocuke.DocString) {
	var msg MsgSubmitTx
	err := s.codec.UnmarshalJSON([]byte(a.Content), &msg)
	assert.NilError(s.t, err, "you may be receiving an error due to testing an inner msg that is not yet "+
		"registered in the interface registry for these tests. Please refer to the 'Before' step of this test suite and "+
		"add the msg type you would like to test to the interface registry")
	s.msg = &msg
}

func (s *submitTxSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *submitTxSuite) ExpectNoError() {
	assert.NilError(s.t, s.err)
}
