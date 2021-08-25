package orm_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/orm/testdata"
)

func TestPrimaryKeyTable(t *testing.T) {
	rapid.Check(t, rapid.Run(&primaryKeyMachine{}))
}

// primaryKeyMachine is a state machine model of the PrimaryKeyTable. The state
// is modelled as a map of strings to GroupMembers.
type primaryKeyMachine struct {
	ctx   orm.HasKVStore
	table *orm.PrimaryKeyTable
	state map[string]*testdata.GroupMember
}

// stateKeys gets all the keys in the model map
func (m *primaryKeyMachine) stateKeys() []string {
	keys := make([]string, len(m.state))

	i := 0
	for k := range m.state {
		keys[i] = k
		i++
	}

	return keys
}

// Generate a GroupMember that has a 50% chance of being a part of the existing
// state
func (m *primaryKeyMachine) genGroupMember() *rapid.Generator {
	genStateGroupMember := rapid.Custom(func(t *rapid.T) *testdata.GroupMember {
		pk := rapid.SampledFrom(m.stateKeys()).Draw(t, "key").(string)
		return m.state[pk]
	})

	if len(m.stateKeys()) == 0 {
		return genGroupMember
	} else {
		return rapid.OneOf(genGroupMember, genStateGroupMember)
	}
}

// Init creates a new instance of the state machine model by building the real
// table and making the empty model map
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

// Check that the real values match the state values. This is kind of overkill,
// because we should catch any discrepancies in a Has command.
func (m *primaryKeyMachine) Check(t *rapid.T) {
	iter, err := m.table.PrefixScan(m.ctx, nil, nil)
	require.NoError(t, err)

	for {
		var dest testdata.GroupMember
		rowID, err := iter.LoadNext(&dest)
		if err == orm.ErrIteratorDone {
			break
		} else {
			require.Equal(t, *m.state[string(rowID)], dest)
		}
	}
}

// Create is one of the model commands. It adds an object to the table, creating
// an error if it already exists.
func (m *primaryKeyMachine) Create(t *rapid.T) {
	g := genGroupMember.Draw(t, "g").(*testdata.GroupMember)
	pk := string(orm.PrimaryKey(g))

	t.Logf("pk: %v", pk)
	t.Logf("m.state: %v", m.state)

	err := m.table.Create(m.ctx, g)

	if m.state[pk] != nil {
		require.Error(t, err)
	} else {
		require.NoError(t, err)
		m.state[pk] = g
	}
}

// Save is one of the model commands. It updates the value at a given primary
// key and fails if that primary key doesn't already exist in the table.
func (m *primaryKeyMachine) Save(t *rapid.T) {
	gm := m.genGroupMember().Draw(t, "gm").(*testdata.GroupMember)

	// We can only really change the weight here, because Group and Member
	// both form part of the PrimaryKey
	newWeight := rapid.Uint64().Draw(t, "newWeight").(uint64)
	gm.Weight = newWeight

	// Perform the real Save
	err := m.table.Save(m.ctx, gm)

	if m.state[string(orm.PrimaryKey(gm))] == nil {
		// If there's no value in the model, we expect an error
		require.Error(t, err)
	} else {
		// If we have a value in the model, expect no error
		require.NoError(t, err)

		// Update the model with the new value
		m.state[string(orm.PrimaryKey(gm))] = gm
	}
}

// Delete is one of the model commands. It removes the object with the given
// primary key from the table and returns an error if that primary key doesn't
// already exist in the table.
func (m *primaryKeyMachine) Delete(t *rapid.T) {
	gm := m.genGroupMember().Draw(t, "gm").(*testdata.GroupMember)

	// Perform the real Delete
	err := m.table.Delete(m.ctx, gm)

	if m.state[string(orm.PrimaryKey(gm))] == nil {
		// If there's no value in the model, we expect an error
		require.Error(t, err)
	} else {
		// If we have a value in the model, expect no error
		require.NoError(t, err)

		// Delete the value from the model
		delete(m.state, string(orm.PrimaryKey(gm)))
	}
}

// Has is one of the model commands. It checks whether a key already exists in
// the table.
func (m *primaryKeyMachine) Has(t *rapid.T) {
	pk := orm.PrimaryKey(m.genGroupMember().Draw(t, "g").(*testdata.GroupMember))

	realHas := m.table.Has(m.ctx, pk)
	modelHas := m.state[string(pk)] != nil

	require.Equal(t, realHas, modelHas)
}

// GetOne is one of the model commands. It fetches an object from the table by
// its primary key and returns an error if that primary key isn't in the table.
func (m *primaryKeyMachine) GetOne(t *rapid.T) {
	pk := orm.PrimaryKey(m.genGroupMember().Draw(t, "gm").(*testdata.GroupMember))

	var gm testdata.GroupMember

	err := m.table.GetOne(m.ctx, pk, &gm)
	t.Logf("gm: %v", gm)

	if m.state[string(pk)] == nil {
		require.Error(t, err)
	} else {
		require.NoError(t, err)
		require.Equal(t, *m.state[string(pk)], gm)
	}
}

// genGroupMember generates a new group member. At the moment it doesn't
// generate empty strings for Group or Member.
var genGroupMember = rapid.Custom(func(t *rapid.T) *testdata.GroupMember {
	return &testdata.GroupMember{
		Group:  []byte(rapid.StringN(1, 100, 150).Draw(t, "group").(string)),
		Member: []byte(rapid.StringN(1, 100, 150).Draw(t, "member").(string)),
		Weight: rapid.Uint64().Draw(t, "weight").(uint64),
	}
})
