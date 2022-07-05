package core

import (
	"testing"

	"github.com/regen-network/gocuke"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type retire struct {
	*baseSuite
	alice sdk.AccAddress
	res   *core.MsgRetireResponse
	err   error
}

func TestRetire(t *testing.T) {
	gocuke.NewRunner(t, &retire{}).Path("./features/msg_retire.feature").Run()
}

func (s *retire) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
}
