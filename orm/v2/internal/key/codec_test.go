package key

import (
	"bytes"
	"fmt"
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
	"gotest.tools/v3/assert"
	"pgregory.net/rapid"
)

func TestPartCodec(t *testing.T) {
	for _, ks := range TestKeyPartSpecs {
		testKeyPartCodec(t, ks)
	}
}

func TestCodec(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		key := TestKeyGen.Draw(t, "key").(TestKey)
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

func assertEncDecKey(t *rapid.T, key TestKey, keyValues []protoreflect.Value) []byte {
	buf := &bytes.Buffer{}
	err := key.Codec.Encode(keyValues, buf)
	assert.NilError(t, err)
	keyValues2, err := key.Codec.Decode(bytes.NewReader(buf.Bytes()))
	assert.NilError(t, err)
	key.RequireValuesEqual(t, keyValues, keyValues2)
	return buf.Bytes()
}

func testKeyPartCodec(t *testing.T, spec TestKeyPartSpec) {
	t.Run(fmt.Sprintf("%s %v", spec.fieldName, false), func(t *testing.T) {
		testKeyPartCodecNT(t, spec.fieldName, spec.gen, false)
	})
	t.Run(fmt.Sprintf("%s %v", spec.fieldName, true), func(t *testing.T) {
		testKeyPartCodecNT(t, spec.fieldName, spec.gen, true)
	})
}

func testKeyPartCodecNT(t *testing.T, fname string, generator *rapid.Generator, nonTerminal bool) {
	cdc, err := MakeTestPartCodec(fname, nonTerminal)
	assert.NilError(t, err)
	rapid.Check(t, func(t *rapid.T) {
		x := protoreflect.ValueOf(generator.Draw(t, fname))
		bz1 := assertEncDecPart(t, x, cdc)
		if cdc.IsOrdered() {
			y := protoreflect.ValueOf(generator.Draw(t, fname+"2"))
			bz2 := assertEncDecPart(t, y, cdc)
			assert.Equal(t, cdc.Compare(x, y), bytes.Compare(bz1, bz2))
		}
	})
}

func assertEncDecPart(t *rapid.T, x protoreflect.Value, cdc PartCodec) []byte {
	buf := &bytes.Buffer{}
	err := cdc.Encode(x, buf)
	assert.NilError(t, err)
	y, err := cdc.Decode(bytes.NewReader(buf.Bytes()))
	assert.NilError(t, err)
	assert.Equal(t, 0, cdc.Compare(x, y))
	return buf.Bytes()
}
