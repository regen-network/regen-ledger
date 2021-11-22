package ormkey_test

import (
	"bytes"
	"testing"

	"github.com/regen-network/regen-ledger/orm/v2/internal/testutil"

	"google.golang.org/protobuf/reflect/protoreflect"
	"gotest.tools/v3/assert"
	"pgregory.net/rapid"
)

func TestCodec(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		key := testutil.TestKeyGen.Draw(t, "key").(testutil.TestKey)
		keyValues := key.Draw(t, "values")

		bz1 := assertEncDecKey(t, key, keyValues)

		if key.Codec.IsFullyOrdered() {
			// check if ordered keys have ordered encodings
			keyValues2 := key.Draw(t, "values2")
			bz2 := assertEncDecKey(t, key, keyValues2)
			// bytes comparison should equal comparison of values
			assert.Equal(t, key.Codec.CompareValues(keyValues, keyValues2), bytes.Compare(bz1, bz2))
		}
	})
}

func assertEncDecKey(t *rapid.T, key testutil.TestKey, keyValues []protoreflect.Value) []byte {
	buf := &bytes.Buffer{}
	err := key.Codec.EncodeWriter(keyValues, buf)
	assert.NilError(t, err)
	keyValues2, err := key.Codec.Decode(bytes.NewReader(buf.Bytes()))
	assert.NilError(t, err)
	key.RequireValuesEqual(t, keyValues, keyValues2)
	return buf.Bytes()
}
