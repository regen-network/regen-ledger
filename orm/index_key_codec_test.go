package orm

// import (
// 	"strings"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestEncodeIndexKey(t *testing.T) {
// 	specs := map[string]struct {
// 		srcKey   []byte
// 		srcRowID RowID
// 		enc      IndexKeyCodec
// 		expKey   []byte
// 		expError bool
// 	}{
// 		"dynamic length example 1": {
// 			srcKey:   []byte{0x0, 0x1, 0x2},
// 			srcRowID: []byte{0x3, 0x4},
// 			enc:      Max255DynamicLengthIndexKeyCodec{},
// 			expKey:   []byte{0x3, 0x0, 0x1, 0x2, 0x3, 0x4},
// 		},
// 		"dynamic length example 2": {
// 			srcKey:   []byte{0x0, 0x1},
// 			srcRowID: []byte{0x2, 0x3, 0x4},
// 			enc:      Max255DynamicLengthIndexKeyCodec{},
// 			expKey:   []byte{0x2, 0x0, 0x1, 0x2, 0x3, 0x4},
// 		},
// 		"dynamic length max row ID": {
// 			srcKey:   []byte{0x0, 0x1},
// 			srcRowID: []byte(strings.Repeat("a", 255)),
// 			enc:      Max255DynamicLengthIndexKeyCodec{},
// 			expKey:   append([]byte{0x2, 0x0, 0x1}, []byte(strings.Repeat("a", 255))...),
// 		},
// 		"dynamic length errors with empty rowID": {
// 			srcKey:   []byte{0x0, 0x1},
// 			srcRowID: []byte{},
// 			enc:      Max255DynamicLengthIndexKeyCodec{},
// 			expError: true,
// 		},
// 		"dynamic length exceeds max searchable key": {
// 			srcKey:   []byte(strings.Repeat("a", 257)),
// 			srcRowID: []byte{0x0, 0x1},
// 			enc:      Max255DynamicLengthIndexKeyCodec{},
// 			expError: true,
// 		},
// 		"uint64 example": {
// 			srcKey:   []byte{0x0, 0x1, 0x2},
// 			srcRowID: []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
// 			enc:      FixLengthIndexKeys(3, 8),
// 			expKey:   []byte{0x0, 0x1, 0x2, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
// 		},
// 		"uint64 errors with empty rowID": {
// 			srcKey:   []byte{0x0, 0x1},
// 			srcRowID: []byte{},
// 			enc:      FixLengthIndexKeys(2, 8),
// 			expError: true,
// 		},
// 		"uint64 exceeds max bytes in rowID": {
// 			srcKey:   []byte{0x0, 0x1},
// 			srcRowID: []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9},
// 			enc:      FixLengthIndexKeys(2, 8),
// 			expError: true,
// 		},
// 	}
// 	for msg, spec := range specs {
// 		t.Run(msg, func(t *testing.T) {
// 			if spec.expError {
// 				_, err := spec.enc.BuildIndexKey(spec.srcKey, spec.srcRowID)
// 				require.Error(t, err)
// 			} else {
// 				got, err := spec.enc.BuildIndexKey(spec.srcKey, spec.srcRowID)
// 				require.NoError(t, err)
// 				assert.Equal(t, spec.expKey, got)
// 			}
// 		})
// 	}
// }
// func TestDecodeIndexKey(t *testing.T) {
// 	specs := map[string]struct {
// 		srcKey   []byte
// 		enc      IndexKeyCodec
// 		expRowID RowID
// 	}{
// 		"dynamic length example 1": {
// 			srcKey:   []byte{0x3, 0x0, 0x1, 0x2, 0x3, 0x4},
// 			enc:      Max255DynamicLengthIndexKeyCodec{},
// 			expRowID: []byte{0x3, 0x4},
// 		},
// 		"dynamic length example 2": {
// 			srcKey:   []byte{0x2, 0x0, 0x1, 0x2, 0x3, 0x4},
// 			enc:      Max255DynamicLengthIndexKeyCodec{},
// 			expRowID: []byte{0x2, 0x3, 0x4},
// 		},
// 		"dynamic length max row ID": {
// 			srcKey:   append([]byte{0x2, 0x0, 0x1}, []byte(strings.Repeat("a", 255))...),
// 			enc:      Max255DynamicLengthIndexKeyCodec{},
// 			expRowID: []byte(strings.Repeat("a", 255)),
// 		},
// 		"uint64 example": {
// 			srcKey:   []byte{0x0, 0x1, 0x2, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
// 			expRowID: []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
// 			enc:      FixLengthIndexKeys(9, 8),
// 		},
// 	}
// 	for msg, spec := range specs {
// 		t.Run(msg, func(t *testing.T) {
// 			gotRow := spec.enc.StripRowID(spec.srcKey)
// 			assert.Equal(t, spec.expRowID, gotRow)
// 		})
// 	}
// }
