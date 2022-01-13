package lookup

import (
	"fmt"
	"hash"
	"hash/fnv"
	"math"
	"testing"

	"github.com/cosmos/cosmos-sdk/store/mem"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/libs/rand"
)

func TestTable(t *testing.T) {
	// test default case with good params
	table, err := NewTable(nil)
	require.NoError(t, err)
	testTable(t, table, 5)

	// test suboptimal case to trigger varint fallback
	table, err = NewTableWithOptions(TableOptions{
		MinLength: 1,
		NewHash: func() hash.Hash {
			return sixteenBitHash{
				fnv.New32(),
			}
		},
	})
	require.NoError(t, err)
	testTable(t, table, 5)
}

type sixteenBitHash struct {
	hash.Hash
}

func (h sixteenBitHash) Sum(b []byte) []byte {
	bz := h.Hash.Sum(b)
	// just return b + the first two bytes
	return bz[:len(b)+2]
}

func testTable(t *testing.T, tbl Table, k int) {
	table := tbl.(table)
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

		// id is nil before it gets set
		id := table.GetID(store, value)
		require.Empty(t, id)

		id, collisions := table.getOrCreateID(store, value, true)
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
		require.Equal(t, value, table.GetValue(store, id))
		newId := table.GetOrCreateID(store, value)
		// make sure getting an ID the second time returns the same ID
		require.Equal(t, id, newId)
		newId = table.GetID(store, value)
		// make sure the normal get method returns an ID
		require.Equal(t, id, newId)
	}
}

//
// NOTE: The tests and benchmarks below were used to find optimal parameter values.
// For now they are retained in source code in case parameters need to be revisisted
// in the future.
//
//func BenchmarkHash(b *testing.B) {
//	buf := rand.Bytes(128)
//	var res []byte
//	b.Run("32a", func(b *testing.B) {
//		for i := 0; i < b.N; i++ {
//			h := fnv.New32a()
//			_, _ = h.Write(buf)
//			res = h.Sum(nil)
//		}
//		b.StopTimer()
//		b.Logf("%x", res)
//	})
//	b.Run("64a", func(b *testing.B) {
//		for i := 0; i < b.N; i++ {
//			h := fnv.New64a()
//			_, _ = h.Write(buf)
//			res = h.Sum(nil)
//		}
//		b.StopTimer()
//		b.Logf("%x", res)
//	})
//	b.Run("128a", func(b *testing.B) {
//		for i := 0; i < b.N; i++ {
//			h := fnv.New128a()
//			_, _ = h.Write(buf)
//			res = h.Sum(nil)
//		}
//		b.StopTimer()
//		b.Logf("%x", res)
//	})
//	b.Run("blake2b-256", func(b *testing.B) {
//		for i := 0; i < b.N; i++ {
//			h := crypto.BLAKE2b_256.New()
//			_, _ = h.Write(buf)
//			res = h.Sum(nil)
//		}
//	})
//	b.Run("blake2b-64", func(b *testing.B) {
//		for i := 0; i < b.N; i++ {
//			h, err := blake2b.New(8, nil)
//			if err != nil {
//				panic(err)
//			}
//			_, _ = h.Write(buf)
//			res = h.Sum(nil)
//		}
//	})
//}

//func TestTableParams(t *testing.T) {
//	if !testing.Verbose() {
//		return
//	}
//
//	n := int(math.Pow10(8))
//	t.Logf("n = %d", n)
//	data := make(map[string][]byte, n)
//	randCollisions := 0
//	for i := 0; i < n; i++ {
//		value := rand.Bytes(32)
//		iri := fmt.Sprintf("regen:%s.rdf", base58.CheckEncode(value, 0))
//		if _, ok := data[iri]; ok {
//			randCollisions++
//		}
//		data[iri] = nil
//	}
//
//	t.Logf("PRNG collisions: %d", randCollisions)
//
//	for lo := 3; lo <= 5; lo++ {
//		for hi := lo + 1; hi <= 6; hi++ {
//			store := mem.NewStore()
//			var totalCollisions uint64
//			var secondaryCollisions uint64
//			var totalBytes int
//
//			for iri := range data {
//				id, collisions := getOrCreateIDForValue(store, []byte(iri), lo, hi)
//				totalCollisions += collisions
//				if collisions > 1 {
//					secondaryCollisions += collisions - 1
//				}
//				totalBytes += len(id)
//			}
//
//			t.Logf("lo %d hi %d totalCollisions %d totalBytes %d secondaryCollisions %d collisionRate %.3f%%", lo, hi, totalCollisions, totalBytes, secondaryCollisions, float64(totalCollisions)/float64(n)*100.0)
//		}
//	}
//}

// n = 10^6
// === RUN   TestTableParams
//    table_test.go:93: n = 1000000
//    table_test.go:104: PRNG collisions: 0
//    table_test.go:122: lo 3 hi 4 totalCollisions 29364 totalBytes 3029368 secondaryCollisions 1 collisionRate 2.936%
//    table_test.go:122: lo 3 hi 5 totalCollisions 29363 totalBytes 3058726 secondaryCollisions 0 collisionRate 2.936%
//    table_test.go:122: lo 3 hi 6 totalCollisions 29363 totalBytes 3088089 secondaryCollisions 0 collisionRate 2.936%
//    table_test.go:122: lo 3 hi 7 totalCollisions 29363 totalBytes 3117452 secondaryCollisions 0 collisionRate 2.936%
//    table_test.go:122: lo 3 hi 8 totalCollisions 29363 totalBytes 3146815 secondaryCollisions 0 collisionRate 2.936%
//    table_test.go:122: lo 4 hi 5 totalCollisions 119 totalBytes 4000119 secondaryCollisions 0 collisionRate 0.012%
//    table_test.go:122: lo 4 hi 6 totalCollisions 119 totalBytes 4000238 secondaryCollisions 0 collisionRate 0.012%
//    table_test.go:122: lo 4 hi 7 totalCollisions 119 totalBytes 4000357 secondaryCollisions 0 collisionRate 0.012%
//    table_test.go:122: lo 4 hi 8 totalCollisions 119 totalBytes 4000476 secondaryCollisions 0 collisionRate 0.012%
//    table_test.go:122: lo 5 hi 6 totalCollisions 0 totalBytes 5000000 secondaryCollisions 0 collisionRate 0.000%
//    table_test.go:122: lo 5 hi 7 totalCollisions 0 totalBytes 5000000 secondaryCollisions 0 collisionRate 0.000%
//    table_test.go:122: lo 5 hi 8 totalCollisions 0 totalBytes 5000000 secondaryCollisions 0 collisionRate 0.000%
//    table_test.go:122: lo 6 hi 7 totalCollisions 0 totalBytes 6000000 secondaryCollisions 0 collisionRate 0.000%
//    table_test.go:122: lo 6 hi 8 totalCollisions 0 totalBytes 6000000 secondaryCollisions 0 collisionRate 0.000%
//    table_test.go:122: lo 7 hi 8 totalCollisions 0 totalBytes 7000000 secondaryCollisions 0 collisionRate 0.000%
//--- PASS: TestTableParams (45.40s)

// n = 10^7
// === RUN   TestTableParams
//    table_test.go:93: n = 10000000
//    table_test.go:105: PRNG collisions: 0
//    table_test.go:123: lo 2 hi 3 totalCollisions 12372869 totalBytes 44564894 secondaryCollisions 2438405 collisionRate 123.729%
//    table_test.go:123: lo 2 hi 4 totalCollisions 9945920 totalBytes 39926208 secondaryCollisions 11456 collisionRate 99.459%
//    table_test.go:123: lo 2 hi 5 totalCollisions 9934502 totalBytes 49803544 secondaryCollisions 38 collisionRate 99.345%
//    table_test.go:123: lo 2 hi 6 totalCollisions 9934464 totalBytes 59737856 secondaryCollisions 0 collisionRate 99.345%
//    table_test.go:123: lo 3 hi 4 totalCollisions 2471316 totalBytes 32479556 secondaryCollisions 2060 collisionRate 24.713%
//    table_test.go:123: lo 3 hi 5 totalCollisions 2469266 totalBytes 34938552 secondaryCollisions 10 collisionRate 24.693%
//    table_test.go:123: lo 3 hi 6 totalCollisions 2469256 totalBytes 37407768 secondaryCollisions 0 collisionRate 24.693%
//    table_test.go:123: lo 3 hi 7 totalCollisions 2469256 totalBytes 39877024 secondaryCollisions 0 collisionRate 24.693%
//    table_test.go:123: lo 3 hi 8 totalCollisions 2469256 totalBytes 42346280 secondaryCollisions 0 collisionRate 24.693%
//    table_test.go:123: lo 4 hi 5 totalCollisions 11670 totalBytes 40011670 secondaryCollisions 0 collisionRate 0.117%
//    table_test.go:123: lo 4 hi 6 totalCollisions 11670 totalBytes 40023340 secondaryCollisions 0 collisionRate 0.117%
//    table_test.go:123: lo 4 hi 7 totalCollisions 11670 totalBytes 40035010 secondaryCollisions 0 collisionRate 0.117%
//    table_test.go:123: lo 4 hi 8 totalCollisions 11670 totalBytes 40046680 secondaryCollisions 0 collisionRate 0.117%
//    table_test.go:123: lo 5 hi 6 totalCollisions 55 totalBytes 50000055 secondaryCollisions 0 collisionRate 0.001%
//    table_test.go:123: lo 5 hi 7 totalCollisions 55 totalBytes 50000110 secondaryCollisions 0 collisionRate 0.001%
//    table_test.go:123: lo 5 hi 8 totalCollisions 55 totalBytes 50000165 secondaryCollisions 0 collisionRate 0.001%
//    table_test.go:123: lo 6 hi 7 totalCollisions 0 totalBytes 60000000 secondaryCollisions 0 collisionRate 0.000%
//    table_test.go:123: lo 6 hi 8 totalCollisions 0 totalBytes 60000000 secondaryCollisions 0 collisionRate 0.000%
//    table_test.go:123: lo 7 hi 8 totalCollisions 0 totalBytes 70000000 secondaryCollisions 0 collisionRate 0.000%
//--- PASS: TestTableParams (734.79s)

// n=10^8
// === RUN   TestTableParams
//    table_test.go:93: n = 100000000
//    table_test.go:105: PRNG collisions: 0
//    table_test.go:123: lo 3 hi 4 totalCollisions 84101626 totalBytes 387444850 secondaryCollisions 835806 collisionRate 84.102%

// n = 10^8
// === RUN   TestTableParams
//    table_test.go:105: n = 100000000
//    table_test.go:117: PRNG collisions: 0
//    table_test.go:135: lo 3 hi 4 totalCollisions 84099113 totalBytes 387431717 secondaryCollisions 833151 collisionRate 84.099%
//    table_test.go:135: lo 3 hi 5 totalCollisions 83269180 totalBytes 466544796 secondaryCollisions 3218 collisionRate 83.269%
//    table_test.go:135: lo 3 hi 6 totalCollisions 83265971 totalBytes 549797913 secondaryCollisions 9 collisionRate 83.266%
//    table_test.go:135: lo 4 hi 5 totalCollisions 1156004 totalBytes 401156088 secondaryCollisions 28 collisionRate 1.156%
//    table_test.go:135: lo 4 hi 6 totalCollisions 1155976 totalBytes 402311952 secondaryCollisions 0 collisionRate 1.156%
//    table_test.go:135: lo 5 hi 6 totalCollisions 4504 totalBytes 500004504 secondaryCollisions 0 collisionRate 0.005%

// BLAKE2b-64 n=2^6
// === RUN   TestTableParams
//    table_test.go:93: n = 1000000
//    table_test.go:105: PRNG collisions: 0
//    table_test.go:123: lo 3 hi 4 totalCollisions 29151 totalBytes 3029167 secondaryCollisions 4 collisionRate 2.915%
//    table_test.go:123: lo 3 hi 5 totalCollisions 29147 totalBytes 3058294 secondaryCollisions 0 collisionRate 2.915%
//    table_test.go:123: lo 3 hi 6 totalCollisions 29147 totalBytes 3087441 secondaryCollisions 0 collisionRate 2.915%
//    table_test.go:123: lo 4 hi 5 totalCollisions 114 totalBytes 4000114 secondaryCollisions 0 collisionRate 0.011%
//    table_test.go:123: lo 4 hi 6 totalCollisions 114 totalBytes 4000228 secondaryCollisions 0 collisionRate 0.011%
//    table_test.go:123: lo 5 hi 6 totalCollisions 0 totalBytes 5000000 secondaryCollisions 0 collisionRate 0.000%

// BLAKE2b-64 n=2^7
// === RUN   TestTableParams
//    table_test.go:105: n = 10000000
//    table_test.go:117: PRNG collisions: 0
//    table_test.go:135: lo 3 hi 4 totalCollisions 2469057 totalBytes 32477117 secondaryCollisions 2015 collisionRate 24.691%
//    table_test.go:135: lo 3 hi 5 totalCollisions 2467050 totalBytes 34934116 secondaryCollisions 8 collisionRate 24.671%
//    table_test.go:135: lo 3 hi 6 totalCollisions 2467042 totalBytes 37401126 secondaryCollisions 0 collisionRate 24.670%
//    table_test.go:135: lo 4 hi 5 totalCollisions 11637 totalBytes 40011637 secondaryCollisions 0 collisionRate 0.116%
//    table_test.go:135: lo 4 hi 6 totalCollisions 11637 totalBytes 40023274 secondaryCollisions 0 collisionRate 0.116%
//    table_test.go:135: lo 5 hi 6 totalCollisions 47 totalBytes 50000047 secondaryCollisions 0 collisionRate 0.000%
