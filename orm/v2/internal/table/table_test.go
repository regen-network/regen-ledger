package table

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/store/mem"
	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"

	"github.com/regen-network/regen-ledger/orm/v2/internal/key"
	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
	"github.com/regen-network/regen-ledger/orm/v2/ormpb"
)

func TestBuildStore(t *testing.T) {
	a := &testpb.A{}
	msgDesc := a.ProtoReflect().Descriptor()
	_, err := BuildStore(nil, &ormpb.TableDescriptor{}, msgDesc)
	require.Error(t, err)
	_, err = BuildStore(nil, &ormpb.TableDescriptor{Id: 1}, msgDesc)
	require.Error(t, err)

	_, err = BuildStore(
		nil,
		&ormpb.TableDescriptor{Id: 1, PrimaryKey: &ormpb.PrimaryKeyDescriptor{
			Fields: "FOO",
		}},
		msgDesc,
	)
	require.Error(t, err)

	_, err = BuildStore(
		nil,
		&ormpb.TableDescriptor{Id: 1, PrimaryKey: &ormpb.PrimaryKeyDescriptor{
			Fields: "UINT32",
		}},
		msgDesc,
	)
	require.NoError(t, err)

	_, err = BuildStore(
		nil,
		&ormpb.TableDescriptor{Id: 1, PrimaryKey: &ormpb.PrimaryKeyDescriptor{
			Fields: "UINT32,UINT32",
		}},
		msgDesc,
	)
	require.Error(t, err, "duplicate")

	_, err = BuildStore(
		nil,
		&ormpb.TableDescriptor{Id: 1, PrimaryKey: &ormpb.PrimaryKeyDescriptor{
			Fields: "UINT32",
		}, Index: []*ormpb.SecondaryIndexDescriptor{
			{},
		}},
		msgDesc,
	)
	require.Error(t, err)

	_, err = BuildStore(
		nil,
		&ormpb.TableDescriptor{Id: 1, PrimaryKey: &ormpb.PrimaryKeyDescriptor{
			Fields: "UINT32",
		}, Index: []*ormpb.SecondaryIndexDescriptor{
			{Id: 1},
		}},
		msgDesc,
	)
	require.Error(t, err)

	_, err = BuildStore(
		nil,
		&ormpb.TableDescriptor{Id: 1, PrimaryKey: &ormpb.PrimaryKeyDescriptor{
			Fields: "UINT32",
		}, Index: []*ormpb.SecondaryIndexDescriptor{
			{Id: 1, Fields: "STRING"},
		}},
		msgDesc,
	)
	require.NoError(t, err)
}

func TestScenarios(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		pk := key.TestKeyGen.Draw(t, "pk").(key.TestKey)
		numIndexes := rapid.IntRange(0, 3).Draw(t, "num indexes").(int)
		var indexes []*ormpb.SecondaryIndexDescriptor
		for i := 0; i < numIndexes; i++ {
			k := key.TestKeyGen.Draw(t, "pk").(key.TestKey)
			indexes = append(indexes, &ormpb.SecondaryIndexDescriptor{
				Fields: k.Fields,
				Id:     uint32(i + 1),
			})
		}

		tableDesc := &ormpb.TableDescriptor{
			PrimaryKey: &ormpb.PrimaryKeyDescriptor{
				Fields: pk.Fields,
			},
			Index: indexes,
			Id:    1,
		}

		a1 := &testpb.A{}
		msgDesc := a1.ProtoReflect().Descriptor()
		st, err := BuildStore(nil, tableDesc, msgDesc)
		require.NoError(t, err)
		require.NotNil(t, st)

		// read empty
		kv := mem.NewStore()
		found, err := st.Read(kv, a1)
		require.False(t, found)
		require.NoError(t, err)

		pk1 := pk.Draw(t, "pk1")
		pk.Codec.SetValues(a1.ProtoReflect(), pk1)
		a1.MESSAGE = &testpb.B{X: "foo"}

		// create
		err = st.Save(kv, a1, store.SAVE_MODE_CREATE)
		require.NoError(t, err)

		// read
		var a2 testpb.A
		pk.Codec.SetValues(a2.ProtoReflect(), pk1)
		found, err = st.Read(kv, &a2)
		require.True(t, found)
		require.NoError(t, err)
		pk.RequireValuesEqual(t, pk1, pk.Codec.GetValues(a2.ProtoReflect()))
		require.NotNil(t, a2.MESSAGE)
		require.Equal(t, a1.MESSAGE.X, a2.MESSAGE.X)
	})
}

func TestAutoInc(t *testing.T) {
	tableDesc := &ormpb.TableDescriptor{
		PrimaryKey: &ormpb.PrimaryKeyDescriptor{
			Fields:        "UINT64",
			AutoIncrement: true,
		},
		Id: 1,
	}

	a1 := &testpb.A{}
	msgDesc := a1.ProtoReflect().Descriptor()
	st, err := BuildStore(nil, tableDesc, msgDesc)
	require.NoError(t, err)
	require.NotNil(t, st)

	// read empty
	kv := mem.NewStore()
	found, err := st.Read(kv, a1)
	require.False(t, found)
	require.NoError(t, err)

	// create fails
	a1.UINT64 = 10
	err = st.Save(kv, a1, store.SAVE_MODE_CREATE)
	require.Error(t, err)

	// create
	a1.UINT64 = 0
	a1.STRING = "foo"
	err = st.Save(kv, a1, store.SAVE_MODE_CREATE)
	require.NoError(t, err)
	require.Equal(t, uint64(1), a1.UINT64)

	// read
	a2 := &testpb.A{UINT64: a1.UINT64}
	found, err = st.Read(kv, a2)
	require.True(t, found)
	require.NoError(t, err)
	require.Equal(t, uint64(1), a2.UINT64)
	require.Equal(t, a1.STRING, a2.STRING)
}
