package ormvalue_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormvalue"

	"github.com/regen-network/regen-ledger/orm/v2/internal/testutil"

	"google.golang.org/protobuf/reflect/protoreflect"
	"gotest.tools/assert"
	"pgregory.net/rapid"
)

func TestPartCodec(t *testing.T) {
	for _, ks := range testutil.TestKeyPartSpecs {
		testKeyPartCodec(t, ks)
	}
}
func testKeyPartCodec(t *testing.T, spec testutil.TestKeyPartSpec) {
	t.Run(fmt.Sprintf("%s %v", spec.FieldName, false), func(t *testing.T) {
		testKeyPartCodecNT(t, spec.FieldName, spec.Gen, false)
	})
	t.Run(fmt.Sprintf("%s %v", spec.FieldName, true), func(t *testing.T) {
		testKeyPartCodecNT(t, spec.FieldName, spec.Gen, true)
	})
}

func testKeyPartCodecNT(t *testing.T, fname string, generator *rapid.Generator, nonTerminal bool) {
	cdc, err := testutil.MakeTestPartCodec(fname, nonTerminal)
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

func assertEncDecPart(t *rapid.T, x protoreflect.Value, cdc ormvalue.Codec) []byte {
	buf := &bytes.Buffer{}
	err := cdc.Encode(x, buf)
	assert.NilError(t, err)
	y, err := cdc.Decode(bytes.NewReader(buf.Bytes()))
	assert.NilError(t, err)
	assert.Equal(t, 0, cdc.Compare(x, y))
	return buf.Bytes()
}
