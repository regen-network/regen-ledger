package eth

import (
	"testing"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

var genValidAddress = rapid.OneOf(
	rapid.StringMatching(RegexAddress),
)

var genInvalidAddress = rapid.OneOf(
	rapid.StringMatching(`0x[g-zG-Z]{40}`),
	rapid.StringMatching(`0x[0-9a-fA-F]{0,39}`),
	rapid.StringMatching(`0x[0-9a-fA-F]{41,}`),
)

func TestAddr(t *testing.T) {
	t.Run("TestValidAddress", rapid.MakeCheck(testValidAddress))
	t.Run("TestInvalidAddress", rapid.MakeCheck(testInvalidAddress))
}

func testValidAddress(t *rapid.T) {
	addr := genValidAddress.Draw(t, "addr").(string)
	require.True(t, IsValidAddress(addr))
}

func testInvalidAddress(t *rapid.T) {
	addr := genInvalidAddress.Draw(t, "addr").(string)
	require.False(t, IsValidAddress(addr))
}
