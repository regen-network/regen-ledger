package ecocredit_test

import (
	"testing"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/stretchr/testify/require"
)

func TestGenesisDefaultParams(t *testing.T) {
	genesis := ecocredit.DefaultGenesisState()
	params := ecocredit.DefaultParams()
	require.Equal(t, params.String(), genesis.Params.String())
}
