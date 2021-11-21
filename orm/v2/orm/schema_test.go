package orm_test

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
	"github.com/regen-network/regen-ledger/orm/v2/orm"
)

func TestSchema(t *testing.T) {
	_, err := orm.BuildSchema(orm.FileDescriptor(0, testpb.File__1_proto))
	assert.NilError(t, err)
}
