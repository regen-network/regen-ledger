package lookup

import (
	"crypto"
	"hash/fnv"
	"testing"

	_ "golang.org/x/crypto/blake2b"

	"github.com/tendermint/tendermint/libs/rand"
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
