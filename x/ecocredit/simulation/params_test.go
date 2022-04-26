package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation"
)

func TestParamChanges(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	bz1, err := json.Marshal([]*core.AskDenom{
		{
			Denom:        "stake",
			DisplayDenom: "stake",
			Exponent:     18,
		},
	})
	require.NoError(t, err)

	expected := []struct {
		composedKey string
		key         string
		simValue    string
		subspace    string
	}{
		{"ecocredit/CreditClassFee", "CreditClassFee", "[{\"denom\":\"stake\",\"amount\":\"6\"}]", "ecocredit"},
		{"ecocredit/AllowlistEnabled", "AllowlistEnabled", "true", "ecocredit"},
		{"ecocredit/AllowedClassCreators", "AllowedClassCreators", "[\"cosmos10z82e5ztmrm4pujgummvmr7aqjzwlp6ga9ams9\"]", "ecocredit"},
		{"ecocredit/AllowedAskDenoms", "AllowedAskDenoms", string(bz1), "ecocredit"},
	}

	paramChanges := simulation.ParamChanges()
	require.Len(t, paramChanges, 4)

	for i, p := range paramChanges {
		require.Equal(t, expected[i].composedKey, p.ComposedKey())
		require.Equal(t, expected[i].key, p.Key())
		require.Equal(t, expected[i].simValue, p.SimValue()(r))
		require.Equal(t, expected[i].subspace, p.Subspace())
	}
}
