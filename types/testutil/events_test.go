package testutil

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/abci/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
)

func TestUtils_MatchFields(t *testing.T) {
	// happy case
	event := api.EventRetire{
		Owner:        "foo",
		BatchDenom:   "bar",
		Amount:       "baz",
		Jurisdiction: "qux",
	}
	se := sdk.Event{
		Attributes: []types.EventAttribute{
			{Key: []byte("owner"), Value: []byte("foo")},
			{Key: []byte("batch_denom"), Value: []byte("bar")},
			{Key: []byte("amount"), Value: []byte("baz")},
			{Key: []byte("jurisdiction"), Value: []byte("qux")},
		},
	}
	err := MatchEvent(&event, se)
	require.NoError(t, err)

	// wrong value case
	event = api.EventRetire{
		Owner:        "foo",
		BatchDenom:   "bar",
		Amount:       "baz",
		Jurisdiction: "nope",
	}
	err = MatchEvent(&event, se)
	require.ErrorContains(t, err, "expected nope, got qux for field Jurisdiction")

	// mismatch in amount of fields
	se.Attributes = append(se.Attributes, types.EventAttribute{Key: []byte("hello"), Value: []byte("world")})
	event.Jurisdiction = "qux"

	err = MatchEvent(&event, se)
	require.ErrorContains(t, err, "emitted event has 5 attributes, expected 4")

	// event has no key to match the field case
	se.Attributes = se.Attributes[:len(se.Attributes)-3]

	err = MatchEvent(&event, se)
	require.ErrorContains(t, err, "event has no attribute 'amount'")
}

func TestUtils_MatchFields_Nested(t *testing.T) {
	event := api.EventMintBatchCredits{
		BatchDenom: "C01-20010101-20050505-001",
		OriginTx: &api.OriginTx{
			Id:       "foo",
			Source:   "bar",
			Contract: "baz",
			Note:     "qux",
		},
	}

	// nested structs just get json marshalled into the value
	bz, err := json.Marshal(event.OriginTx)
	require.NoError(t, err)

	se := sdk.Event{
		Attributes: []types.EventAttribute{
			{Key: []byte("batch_denom"), Value: []byte("C01-20010101-20050505-001")},
			{Key: []byte("origin_tx"), Value: bz},
		},
	}

	err = MatchEvent(&event, se)
	require.NoError(t, err)
}
