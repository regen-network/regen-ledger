package v1

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type submitTxSuite struct {
	t   gocuke.TestingT
	msg *MsgSubmitTx
	err error
}

func TestMsgSubmitTx(t *testing.T) {
	gocuke.NewRunner(t, &submitTxSuite{}).Path("./features/msg_submit_tx.feature").Run()
}

func (s *submitTxSuite) Before(t gocuke.TestingT) {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
	s.t = t
}

func (s *submitTxSuite) ExpectTheError(a string) {
	assert.ErrorContains(s.t, s.err, a)
}

func (s *submitTxSuite) TheMessage(a gocuke.DocString) {
	var msg MsgSubmitTx
	err := json.Unmarshal([]byte(a.Content), &msg)
	assert.NilError(s.t, err)
	s.msg = &msg
}

func (s *submitTxSuite) AValidTxForMsg() {
	msg, err := types.NewAnyWithValue(&MsgRegisterAccount{})
	assert.NilError(s.t, err)
	s.msg.Msg = msg
}

func (s *submitTxSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *submitTxSuite) ExpectNoError() {
	assert.NilError(s.t, s.err)
}
