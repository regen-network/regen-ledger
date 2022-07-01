package core

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/gocuke"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type mintBatchCredits struct {
	*baseSuite
	alice            sdk.AccAddress
	creditTypeAbbrev string
	classId          string
	classKey         uint64
	projectId        string
	res              *core.MsgMintBatchCreditsResponse
	err              error
}

func TestMintBatchCredits(t *testing.T) {
	gocuke.NewRunner(t, &mintBatchCredits{}).Path("./features/msg_mint_batch_credits.feature").Run()
}

func (s *mintBatchCredits) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.creditTypeAbbrev = "C"
	s.classId = "C01"
	s.projectId = "C01-001"
}
