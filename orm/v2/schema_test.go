package orm_test

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/orm/v2"
	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
)

func TestSchema(t *testing.T) {
	_, err := orm.BuildSchema(orm.FileDescriptor(0, testpb.File__1_proto))
	assert.NilError(t, err)
}
