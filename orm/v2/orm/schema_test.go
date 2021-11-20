package orm_test

import (
	"testing"

	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
	"github.com/regen-network/regen-ledger/orm/v2/orm"
	"github.com/stretchr/testify/require"
)

func TestSchema(t *testing.T) {
	_, err := orm.BuildSchema(orm.FileDescriptor(0, testpb.File__1_proto))
	require.NoError(t, err)
}
