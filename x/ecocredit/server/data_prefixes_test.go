package server

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var addr = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())

func TestKeys(t *testing.T) {
	batchDenom := batchDenomT("testing-denom")

	// tradable-balance-key
	key := TradableBalanceKey(addr, batchDenom)
	a, d := ParseBalanceKey(key)
	require.Equal(t, a, addr)
	require.Equal(t, d, batchDenom)

	// tradable-supply-key
	key = TradableSupplyKey(batchDenom)
	d = ParseSupplyKey(key)
	require.Equal(t, a, addr)
	require.Equal(t, d, batchDenom)

	// retired-balance-key
	key = RetiredBalanceKey(addr, batchDenom)
	a, d = ParseBalanceKey(key)
	require.Equal(t, a, addr)
	require.Equal(t, d, batchDenom)

	// retired-supply-key
	key = RetiredSupplyKey(batchDenom)
	d = ParseSupplyKey(key)
	require.Equal(t, a, addr)
	require.Equal(t, d, batchDenom)

}
