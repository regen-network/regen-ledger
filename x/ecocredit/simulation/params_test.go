package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation"
	"github.com/stretchr/testify/require"
)

func TestParamChanges(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	bz, err := json.Marshal(ecocredit.DefaultParams().CreditTypes)
	require.NoError(t, err)

	expected := []struct {
		composedKey string
		key         string
		simValue    string
		subspace    string
	}{
		{"ecocredit/CreditClassFee", "CreditClassFee", "[{\"denom\":\"stake\",\"amount\":\"2\"}]", "ecocredit"},
		{"ecocredit/AllowlistEnabled", "AllowlistEnabled", "true", "ecocredit"},
		{"ecocredit/AllowedClassDesigners", "AllowedClassDesigners", "[\"cosmos10z82e5ztmrm4pujgummvmr7aqjzwlp6ga9ams9\",\"cosmos1s8q96nnww8q8wp7h2ulqq5nzglj9mgqfvf2g5j\",\"cosmos1s7wmf52cnq7vdu7xe23gwtm73rfg2kjqp7z3lk\"]", "ecocredit"},
		{"ecocredit/CreditTypes", "CreditTypes", string(bz), "ecocredit"},
	}

	paramChanges := simulation.ParamChanges(r)

	require.Len(t, paramChanges, 4)

	for i, p := range paramChanges {
		require.Equal(t, expected[i].composedKey, p.ComposedKey())
		require.Equal(t, expected[i].key, p.Key())
		require.Equal(t, expected[i].simValue, p.SimValue()(r))
		require.Equal(t, expected[i].subspace, p.Subspace())
	}
}
