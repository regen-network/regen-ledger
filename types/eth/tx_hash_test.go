package eth

import (
	"testing"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

var genValidTxHash = rapid.OneOf(
	rapid.StringMatching(RegexTxHash),
)

var genInvalidTxHash = rapid.OneOf(
	rapid.StringMatching(`0x[g-zG-Z]{64}`),
	rapid.StringMatching(`0x[0-9a-fA-F]{0,63}`),
	rapid.StringMatching(`0x[0-9a-fA-F]{65,}`),
)

func TestTxHash(t *testing.T) {
	t.Run("TestValidTxHash", rapid.MakeCheck(testValidTxHash))
	t.Run("TestInvalidTxHash", rapid.MakeCheck(testInvalidTxHash))
}

func testValidTxHash(t *rapid.T) {
	txHash := genValidTxHash.Draw(t, "txHash").(string)
	require.True(t, IsValidTxHash(txHash))
}

func testInvalidTxHash(t *rapid.T) {
	txHash := genInvalidTxHash.Draw(t, "txHash").(string)
	require.False(t, IsValidTxHash(txHash))
}
