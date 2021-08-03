package orm_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"pgregory.net/rapid"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/orm/testdata"
)

func TestPrimaryKeyTable(t *testing.T) {
	rapid.Check(t, rapid.Run(&primaryKeyMachine{}))
}

type primaryKeyMachine struct {
	ctx   orm.HasKVStore
	table *orm.PrimaryKeyTable
	state map[string]*testdata.GroupMember
}

func (m *primaryKeyMachine) Init(t *rapid.T) {
	// Create context
	m.ctx = orm.NewMockContext()

	// Create primary key table
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	storeKey := sdk.NewKVStoreKey("test")
	const testTablePrefix = iota
	table := orm.NewPrimaryKeyTableBuilder(
		testTablePrefix,
		storeKey,
		&testdata.GroupMember{},
		orm.Max255DynamicLengthIndexKeyCodec{},
		cdc,
	).Build()

	m.table = &table

	// Create model state
	m.state = make(map[string]*testdata.GroupMember)
}

func (*primaryKeyMachine) Check(t *rapid.T) {
	// What do we actually want to check here?
}

func (m *primaryKeyMachine) Create(t *rapid.T) {
	g := genGroupMember.Draw(t, "g").(*testdata.GroupMember)
	pk := string(g.PrimaryKey())

	t.Logf("pk: %v", pk)
	t.Logf("m.state: %v", m.state)

	err := m.table.Create(m.ctx, g)

	if m.state[pk] != nil && err == nil {
		t.Fatal("Create: Expected err when primary key exists")
	} else if m.state[pk] == nil && err != nil {
		t.Fatalf("Create: Unexpected error when key doesn't exists: %s", err.Error())
	}

	m.state[pk] = g
}

var genGroupMember = rapid.Custom(func(t *rapid.T) *testdata.GroupMember {
	return &testdata.GroupMember{
		Group:  []byte(rapid.String().Draw(t, "group").(string)),
		Member: []byte(rapid.String().Draw(t, "member").(string)),
		Weight: rapid.Uint64().Draw(t, "weight").(uint64),
	}
})
