package ecocredit

import (
	"testing"

	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
)

func TestGetDecimalsFromBalance(t *testing.T) {
	t.Parallel()
	bal := api.BatchBalance{
		Tradable: "15",
		Retired:  "10",
		Escrowed: "",
	}
	tradable, retired, escrowed, err := GetDecimalsFromBalance(&bal)
	assert.NilError(t, err)
	expectedTradable := math.NewDecFromInt64(15)
	expectedRetired := math.NewDecFromInt64(10)
	expectedEscrowed := math.NewDecFromInt64(0)

	assert.Check(t, expectedTradable.Equal(tradable))
	assert.Check(t, expectedRetired.Equal(retired))
	assert.Check(t, expectedEscrowed.Equal(escrowed))
}
