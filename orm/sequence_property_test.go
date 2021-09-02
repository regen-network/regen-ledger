package orm_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"

	"github.com/regen-network/regen-ledger/orm"
)

func TestSequence(t *testing.T) {
	rapid.Check(t, rapid.Run(&sequenceMachine{}))
}

// sequenceMachine is a state machine model of Sequence. It simply uses a uint64
// as the model of the sequence.
type sequenceMachine struct {
	ctx   orm.HasKVStore
	seq   *orm.Sequence
	state uint64
}

// Init sets up the real Sequence, including choosing a random initial value,
// and intialises the model state
func (m *sequenceMachine) Init(t *rapid.T) {
	// Create context
	m.ctx = orm.NewMockContext()

	// Create primary key table
	storeKey := sdk.NewKVStoreKey("test")
	seq := orm.NewSequence(storeKey, 0x1)
	m.seq = &seq

	// Choose initial sequence value
	initSeqVal := rapid.Uint64().Draw(t, "initSeqVal").(uint64)
	err := m.seq.InitVal(m.ctx, initSeqVal)
	require.NoError(t, err)

	// Create model state
	m.state = initSeqVal
}

// Check does nothing, because all our invariants are captured in the commands
func (m *sequenceMachine) Check(t *rapid.T) {}

// NextVal is one of the model commands. It checks that the next value of the
// sequence matches the model and increments the model state.
func (m *sequenceMachine) NextVal(t *rapid.T) {
	// Check that the next value in the sequence matches the model
	require.Equal(t, m.state+1, m.seq.NextVal(m.ctx))

	// Increment the model state
	m.state++
}

// CurVal is one of the model commands. It checks that the current value of the
// sequence matches the model.
func (m *sequenceMachine) CurVal(t *rapid.T) {
	// Check the the current value matches the model
	require.Equal(t, m.state, m.seq.CurVal(m.ctx))
}

// PeekNextVal is one of the model commands. It checks that the next value of
// the sequence matches the model without modifying the state.
func (m *sequenceMachine) PeekNextVal(t *rapid.T) {
	// Check that the next value in the sequence matches the model
	require.Equal(t, m.state+1, m.seq.PeekNextVal(m.ctx))
}
