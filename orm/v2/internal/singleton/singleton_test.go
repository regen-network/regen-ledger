package singleton

import (
	"testing"

	"github.com/regen-network/regen-ledger/orm/v2/internal/list"
	"github.com/regen-network/regen-ledger/orm/v2/ormpb"

	"github.com/cosmos/cosmos-sdk/store/mem"
	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
	"github.com/stretchr/testify/require"
)

func TestSingleton(t *testing.T) {
	_, err := BuildStore(nil, &ormpb.SingletonDescriptor{Id: 0})
	require.Error(t, err)

	store, err := BuildStore(nil, &ormpb.SingletonDescriptor{Id: 1})
	require.NoError(t, err)

	kv := mem.NewStore()
	b1 := &testpb.B{X: "abc"}

	// read empty
	found, err := store.Read(kv, b1)
	require.False(t, found)
	require.NoError(t, err)

	// create
	err = store.Create(kv, b1)
	require.NoError(t, err)

	// read
	var b2 testpb.B
	found, err = store.Read(kv, &b2)
	require.True(t, found)
	require.NoError(t, err)
	require.Equal(t, b1.X, b2.X)

	// create a second time fails
	b1.X = "def"
	err = store.Create(kv, b1)
	require.Error(t, err)

	// save succeeds
	err = store.Save(kv, b1)
	require.NoError(t, err)

	// read
	found, err = store.Read(kv, &b2)
	require.True(t, found)
	require.NoError(t, err)
	require.Equal(t, b1.X, b2.X)

	// iterator just returns one value always
	it := store.List(kv, &list.Options{})
	require.NotNil(t, it)
	found, err = it.Next(&b2)
	require.True(t, found)
	require.NoError(t, err)
	require.Equal(t, b1.X, b2.X)
	found, err = it.Next(&b2)
	require.False(t, found)
	require.NoError(t, err)
	found, err = it.Next(&b2) // next always does the same thing
	require.False(t, found)
	require.NoError(t, err)

	// delete
	err = store.Delete(kv, b1)
	require.NoError(t, err)
	err = store.Delete(kv, b1)
	require.NoError(t, err) // deleting twice is a no-op

	// can't read
	found, err = store.Read(kv, b1)
	require.False(t, found)
	require.NoError(t, err)
}
