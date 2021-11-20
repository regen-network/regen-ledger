package key

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
	"pgregory.net/rapid"

	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
)

type keySpec struct {
	fieldName string
	gen       *rapid.Generator
}

var keySpecs = []keySpec{
	{
		"UINT32",
		rapid.Uint32(),
	},
	{
		"UINT64",
		rapid.Uint64(),
	},
	{
		"STRING",
		rapid.String().Filter(func(x string) bool {
			// filter out null terminators
			return strings.IndexByte(x, 0) < 0
		}),
	},
	{
		"BYTES",
		rapid.SliceOfN(rapid.Byte(), 0, 255),
	},
}

func TestPartCodec(t *testing.T) {
	for _, ks := range keySpecs {
		testKeyPartCodec(t, ks)
	}
}

func TestCodec(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		n := len(keySpecs)
		key := rapid.SliceOfN(rapid.IntRange(0, n-1), 1, n).Map(
			func(xs []int) []keySpec {
				var res []keySpec
				for _, x := range xs {
					res = append(res, keySpecs[x])
				}
				return res
			},
		).Draw(t, "key").([]keySpec)

		n = len(key)
		fields := make([]protoreflect.FieldDescriptor, n)
		keyValues := make([]protoreflect.Value, n)
		for i, k := range key {
			fields[i] = getTestField(k.fieldName)
			keyValues[i] = protoreflect.ValueOf(k.gen.Draw(t, fmt.Sprintf("keyValue[%d]", i)))
		}

		keyCdc, err := MakeCodec(fields, true)
		require.NoError(t, err)
		buf := &bytes.Buffer{}
		err = keyCdc.Encode(keyValues, buf, false)
		require.NoError(t, err)
		keyValues2, err := keyCdc.Decode(bytes.NewReader(buf.Bytes()))
		require.NoError(t, err)
		for i := 0; i < n; i++ {
			require.Equalf(t, keyValues[i].Interface(), keyValues2[i].Interface(), "values[%d]", i)
		}
	})
}

func testKeyPartCodec(t *testing.T, spec keySpec) {
	t.Run(fmt.Sprintf("%s %v", spec.fieldName, false), func(t *testing.T) {
		testKeyPartCodecNT(t, spec.fieldName, spec.gen, false)
	})
	t.Run(fmt.Sprintf("%s %v", spec.fieldName, true), func(t *testing.T) {
		testKeyPartCodecNT(t, spec.fieldName, spec.gen, true)
	})
}

func testKeyPartCodecNT(t *testing.T, fname string, generator *rapid.Generator, nonTerminal bool) {
	cdc, err := makeKeyPartCodec(fname, nonTerminal)
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

func makeKeyPartCodec(fname string, nonTerminal bool) (PartCodec, error) {
	return makePartCodec(getTestField(fname), nonTerminal)
}

func getTestField(fname string) protoreflect.FieldDescriptor {
	a := &testpb.A{}
	return GetFieldDescriptor(a.ProtoReflect().Descriptor(), fname)
}
