package hasher

import (
	"fmt"
	"hash"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/blake2b"

	"github.com/cosmos/cosmos-sdk/store/mem"
	"github.com/tendermint/tendermint/libs/rand"
)

func TestHasher(t *testing.T) {
	// test default case with good params
	hasher, err := NewHasher()
	require.NoError(t, err)
	testHasher(t, hasher, 5)

	// test suboptimal case to trigger varint fallback
	hasher, err = NewHasherWithOptions(HashOptions{
		MinLength: 1,
		NewHash: func() hash.Hash {
			hash, err := blake2b.New(8, nil)
			if err != nil {
				panic(err) // an error should not occur creating a hash
			}
			return sixteenBitHash{
				hash,
			}
		},
	})
	require.NoError(t, err)
	testHasher(t, hasher, 5)
}

type sixteenBitHash struct {
	hash.Hash
}

func (h sixteenBitHash) Sum(b []byte) []byte {
	bz := h.Hash.Sum(b)
	// just return b + the first three bytes
	return bz[:len(b)+3]
}

func testHasher(t *testing.T, h Hasher, k int) {
	hasher := h.(hasher)
	store := mem.NewStore()
	n := int(math.Pow10(k))
	data := make([][]byte, n)
	values := map[string]bool{}
	ids := map[int][]byte{}
	totalCollisions := 0
	secondaryCollisions := 0

	for i := 0; i < n; i++ {
		var value []byte
		var valueStr string
		for {
			m := rand.Int31n(256)
			value = rand.Bytes(int(m))
			valueStr = fmt.Sprintf("%x", value)
			if !values[valueStr] {
				break
			}
		}
		data[i] = value
		values[valueStr] = true

		c := 0
		for ; ; c++ {
			id := hasher.CreateID(value, c)
			v := store.Get(id)
			if len(v) == 0 {
				ids[i] = id
				store.Set(id, value)
				break
			}
		}
		if c > 1 {
			totalCollisions++
		}
		if c > 2 {
			secondaryCollisions++
		}
	}

	t.Logf("total collisions: %d / %.0e, secondary collisions: %d, collision rate: %.4f%%", totalCollisions, float64(n), secondaryCollisions, float64(totalCollisions)/float64(n)*100.0)

	store = mem.NewStore()

	for i := 0; i < n; i++ {
		id := ids[i]
		value := data[i]

		for c := 0; ; c++ {
			newID := hasher.CreateID(value, c)
			v := store.Get(newID)
			if len(v) == 0 {
				store.Set(newID, value)
				require.Equal(t, id, newID)
				break
			}
		}
	}
}
