package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/testutil"
)

func TestMsgMsgSealBatch(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	issuer := testutil.GenAddress()

	msg := MsgSealBatch{}
	require.Error(msg.ValidateBasic(), "empty issuer")

	msg = MsgSealBatch{Issuer: "abc"}
	require.Error(msg.ValidateBasic(), "invalid issuer")

	msg = MsgSealBatch{Issuer: "abc", BatchDenom: "ABC"}
	require.Error(msg.ValidateBasic(), "invalid denom")

	msg = MsgSealBatch{Issuer: issuer, BatchDenom: batchDenom}
	require.NoError(msg.ValidateBasic(), "valid issuer")
}
