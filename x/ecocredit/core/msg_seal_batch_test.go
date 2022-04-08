package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMsgMsgSealBatch(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	issuer := genAddress()

	msg := MsgSealBatch{}
	require.Error(msg.ValidateBasic(), "empty issuer")

	msg = MsgSealBatch{Issuer: "abc"}
	require.Error(msg.ValidateBasic(), "invalid issuer")

	msg = MsgSealBatch{Issuer: issuer}
	require.NoError(msg.ValidateBasic(), "valid issuer")
}
