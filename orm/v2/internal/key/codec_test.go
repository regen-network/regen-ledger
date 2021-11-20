package key

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
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

		buf := &bytes.Buffer{}
		err := key.Codec.Encode(keyValues, buf, false)
		require.NoError(t, err)
		keyValues2, err := key.Codec.Decode(bytes.NewReader(buf.Bytes()))
		require.NoError(t, err)
		key.RequireValuesEqual(t, keyValues, keyValues2)
	})
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
	require.NoError(t, err)
	rapid.Check(t, func(t *rapid.T) {
		x := protoreflect.ValueOf(generator.Draw(t, fname))
		buf := &bytes.Buffer{}
		err = cdc.encode(x, buf, false)
		require.NoError(t, err)
		y, err := cdc.decode(bytes.NewReader(buf.Bytes()))
		require.NoError(t, err)
		require.Equal(t, x.Interface(), y.Interface())
	})
}
