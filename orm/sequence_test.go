package orm

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSequenceUniqueConstraint(t *testing.T) {
	storeKey := sdk.NewKVStoreKey("test")
	ctx := NewMockContext()
	seq := NewSequence(storeKey, 0x1)

	err := seq.InitVal(ctx, 2)
	require.NoError(t, err)
	err = seq.InitVal(ctx, 3)
	require.True(t, ErrUniqueConstraint.Is(err))
}

func TestSequenceIncrements(t *testing.T) {
	storeKey := sdk.NewKVStoreKey("test")
	ctx := NewMockContext()

	seq := NewSequence(storeKey, 0x1)
	var i uint64
	for i = 1; i < 10; i++ {
		autoID := seq.NextVal(ctx)
		assert.Equal(t, i, autoID)
		assert.Equal(t, i, seq.CurVal(ctx))
	}
	// and persisted
	seq = NewSequence(storeKey, 0x1)
	assert.Equal(t, uint64(9), seq.CurVal(ctx))
}
