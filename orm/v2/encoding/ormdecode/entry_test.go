package ormdecode

import (
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/testing/protocmp"
)

func Test1(t *testing.T) {
	p1 := PrimaryKeyEntry{
		Key:   []protoreflect.Value{protoreflect.ValueOfUint32(2)},
		Value: &testpb.A{UINT32: 5},
	}
	p2 := PrimaryKeyEntry{
		Key:   []protoreflect.Value{protoreflect.ValueOfUint32(2)},
		Value: &testpb.A{UINT32: 5},
	}
	assert.Equal(t, p1, p2, protocmp.Transform())
	t.Logf("%v", p1)
}
