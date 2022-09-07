package v1

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type registerAccountSuite struct {
	t   gocuke.TestingT
	msg *MsgRegisterAccount
	err error
}

func TestMsgRegisterAccount(t *testing.T) {
	gocuke.NewRunner(t, &registerAccountSuite{}).Path("./features/msg_register_account.feature").Run()
}

func (s *registerAccountSuite) Before(t gocuke.TestingT) {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
	s.t = t
}

func (s *registerAccountSuite) TheMessage(a gocuke.DocString) {
	var msg MsgRegisterAccount
	err := json.Unmarshal([]byte(a.Content), &msg)
	assert.NilError(s.t, err)
	s.msg = &msg
}

func (s *registerAccountSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *registerAccountSuite) ExpectNoError() {
	assert.NilError(s.t, s.err)
}

func (s *registerAccountSuite) ExpectTheError(a string) {
	assert.ErrorContains(s.t, s.err, a)
}
