package v1

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type suite struct {
	t   gocuke.TestingT
	msg *MsgRegisterAccount
	err error
}

func TestMsgRegisterAccount(t *testing.T) {
	gocuke.NewRunner(t, &suite{}).Path("./features/msg_register_account.feature").Run()
}

func (s *suite) Before(t gocuke.TestingT) {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
	s.t = t
}

func (s *suite) TheMessage(a gocuke.DocString) {
	var msg MsgRegisterAccount
	err := json.Unmarshal([]byte(a.Content), &msg)
	assert.NilError(s.t, err)
	s.msg = &msg
}

func (s *suite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *suite) ExpectNoError() {
	assert.NilError(s.t, s.err)
}

func (s *suite) ExpectTheError(a string) {
	assert.ErrorContains(s.t, s.err, a)
}
