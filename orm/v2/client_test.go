package orm

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/store/mem"
	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestClient(t *testing.T) {
	schema, err := BuildSchema(FileDescriptor(0, testpb.File__1_proto))
	require.NoError(t, err)
	clientConn := &ClientConn{schema}
	client := clientConn.Open(mem.NewStore())

	data := []proto.Message{
		&testpb.A{
			UINT32: 4,
			UINT64: 10,
			STRING: "abc",
		},
		&testpb.A{
			UINT32: 4,
			UINT64: 10,
			STRING: "foo",
		},
		&testpb.A{
			UINT32: 4,
			UINT64: 11,
			STRING: "abc",
		},
		&testpb.A{
			UINT32: 5,
			UINT64: 3,
			STRING: "foo",
		},
	}

	require.NoError(t, client.Save(data...))

	for i, x := range data {
		require.Truef(t, client.Has(x), "data[%d]", i)
	}

	it := client.List(&testpb.A{})
	defer it.Close()
	require.NotNil(t, it)
	var a1 testpb.A
	have, err := it.Next(&a1)
	require.True(t, have)
	require.NoError(t, err)
	have, err = it.Next(&a1)
	require.True(t, have)
	require.NoError(t, err)
	have, err = it.Next(&a1)
	require.True(t, have)
	require.NoError(t, err)
	have, err = it.Next(&a1)
	require.True(t, have)
	require.NoError(t, err)
	// no more elements
	have, err = it.Next(&a1)
	require.False(t, have)
	require.NoError(t, err)
}
