package key

import (
	"bytes"
	"fmt"
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
	"pgregory.net/rapid"
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
		rapid.String(),
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

func testKeyPartCodec(t *testing.T, spec keySpec) {
	t.Run(fmt.Sprintf("%s %v", spec.fieldName, false), func(t *testing.T) {
		testKeyPartCodecNT(t, spec.fieldName, spec.gen, false)
	})
	t.Run(fmt.Sprintf("%s %v", spec.fieldName, true), func(t *testing.T) {
		testKeyPartCodecNT(t, spec.fieldName, spec.gen, true)
	})
}

func testKeyPartCodecNT(t *testing.T, fname string, generator *rapid.Generator, nonTerminal bool) {
	enc, dec, err := makeKeyPartCodec(fname, nonTerminal)
	require.NoError(t, err)
	rapid.Check(t, func(t *rapid.T) {
		x := protoreflect.ValueOf(generator.Draw(t, fname))
		buf := &bytes.Buffer{}
		err = enc(x, buf, false)
		require.NoError(t, err)
		var y interface{}
		y, err = dec(bytes.NewReader(buf.Bytes()))
		require.NoError(t, err)
		require.Equal(t, x, y)
	})
}

func makeKeyPartCodec(fname string, nonTerminal bool) (keyPartEncoder, keyPartDecoder, error) {
	a := &testpb.A{}
	f := GetFieldDescriptor(a.ProtoReflect().Descriptor(), fname)
	enc, err := makeKeyPartEncoder(f, nonTerminal)
	if err != nil {
		return nil, nil, err
	}

	dec, err := makeKeyPartDecoder(f, nonTerminal)
	if err != nil {
		return nil, nil, err
	}

	return enc, dec, nil
}
