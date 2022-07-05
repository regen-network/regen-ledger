package core

import (
	"testing"

	"github.com/regen-network/gocuke"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type cancel struct {
	*baseSuite
	alice sdk.AccAddress
	res   *core.MsgCancelResponse
	err   error
}

func TestCancel(t *testing.T) {
	gocuke.NewRunner(t, &cancel{}).Path("./features/msg_cancel.feature").Run()
}

func (s *cancel) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
}
