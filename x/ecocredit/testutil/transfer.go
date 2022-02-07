package testutil

import (
	"github.com/regen-network/regen-ledger/x/ecocredit/fill"
	"github.com/rs/zerolog"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/math"
)

type TestTransferManager struct {
	logger zerolog.Logger
}

func NewTestTransferManager(logger zerolog.Logger) *TestTransferManager {
	return &TestTransferManager{logger: logger}
}

func (t TestTransferManager) SendCoinsTo(denom string, amount sdk.Int, from, to sdk.AccAddress) error {
	t.logger.Printf("Transfer %s %s from %s -> %s",
		amount.String(), denom, from, to)

	return nil
}

func (t TestTransferManager) SendCreditsTo(batchId uint64, amount math.Dec, from, to sdk.AccAddress, retire bool) error {
	action := "Transfer"
	if retire {
		action = "Retire"
	}
	t.logger.Printf("%s %s credits from batch %d from %s -> %s",
		action, amount.String(), batchId, from, to)

	return nil
}

var _ fill.TransferManager = TestTransferManager{}
