package lookup

import (
	"crypto"
	"hash/fnv"
	"math"
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
	n := int(math.Pow10(5))
	data := make([][]byte, n)
	ids := map[int][]byte{}
	var totalCollisions uint64
	var secondaryCollisions uint64

	for i := 0; i < n; i++ {
		m := rand.Int31n(256)
		value := rand.Bytes(int(m))
		data[i] = value
		id, collisions := getOrCreateIDForValue(store, value)
		totalCollisions += collisions
		if collisions > 1 {
			secondaryCollisions += collisions - 1
		}
		ids[i] = id
	}

	t.Logf("total collisions: %d / %.0e, secondary collisions: %d, collision rate: %.4f%%", totalCollisions, float64(n), secondaryCollisions, float64(totalCollisions)/float64(n)*100.0)

	for i := 0; i < n; i++ {
		id := ids[i]
		value := data[i]
		// make sure stored value at id is expected value
		require.Equal(t, value, store.Get(id))
		newId := GetOrCreateIDForValue(store, value)
		// make sure getting an ID the second time returns the same ID
		require.Equal(t, id, newId)
	}
}
