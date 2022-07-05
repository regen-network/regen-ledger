package core

import (
	"testing"

	"github.com/regen-network/gocuke"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type send struct {
	*baseSuite
	alice sdk.AccAddress
	res   *core.MsgSendResponse
	err   error
}

func TestSend(t *testing.T) {
	gocuke.NewRunner(t, &send{}).Path("./features/msg_send.feature").Run()
}

func (s *send) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
}
