package types

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestMetadata(t *testing.T) {
	validMd := "aGVsbG8gcmVnZW4h" // this is "hello regen!" in base64
	validMdDecoded := "hello regen!"
	invalidMd := "ha!lkj23tkm,dsf"

	bz, err := DecodeMetadata(validMd)
	assert.NilError(t, err)
	assert.Equal(t, validMdDecoded, string(bz))

	_, err = DecodeMetadata(invalidMd)
	assert.ErrorContains(t, err, "malformed")
}
