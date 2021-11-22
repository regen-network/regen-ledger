package ormschema_test

import (
	"testing"

	"github.com/regen-network/regen-ledger/orm/v2/model/ormschema"

	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
)

func TestSchema(t *testing.T) {
	_, err := ormschema.BuildSchema(ormschema.FileDescriptor(0, testpb.File__1_proto))
	assert.NilError(t, err)
}
