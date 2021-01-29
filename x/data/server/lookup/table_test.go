package lookup

import (
	"crypto"
	"hash/fnv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/store/mem"
	"github.com/tendermint/tendermint/libs/rand"
	_ "golang.org/x/crypto/blake2b"
)

func BenchmarkHash(b *testing.B) {
	buf := rand.Bytes(128)
	var res []byte
	b.Run("32a", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := fnv.New32a()
			_, _ = h.Write(buf)
			res = h.Sum(nil)
		}
		b.StopTimer()
		b.Logf("%x", res)
	})
	b.Run("64a", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := fnv.New64a()
			_, _ = h.Write(buf)
			res = h.Sum(nil)
		}
		b.StopTimer()
		b.Logf("%x", res)
	})
	b.Run("128a", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := fnv.New128a()
			_, _ = h.Write(buf)
			res = h.Sum(nil)
		}
		b.StopTimer()
		b.Logf("%x", res)
	})
	b.Run("blake2b", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := crypto.BLAKE2b_256.New()
			_, _ = h.Write(buf)
			res = h.Sum(nil)
		}
	})
}

func TestTable(t *testing.T) {
	store := mem.NewStore()
	n := 1000000
	data := make([][]byte, n)
	ids := map[int][]byte{}
	for i := 0; i < n; i++ {
		m := rand.Int31n(1000)
		value := rand.Bytes(int(m))
		data[i] = value
		id := GetOrCreateIDForValue(store, value)
		ids[i] = id
	}
	for i := 0; i < n; i++ {
		id := ids[i]
		value := data[i]
		require.Equal(t, value, store.Get(id))
		newId := GetOrCreateIDForValue(store, value)
		require.Equal(t, id, newId)
	}
}
