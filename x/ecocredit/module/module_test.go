package module

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

func TestValidateGenesis(t *testing.T) {
	m := Module{}
	cdc := codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

	// default genesis
	bz := m.DefaultGenesis(cdc)
	require.NotNil(t, bz)
	require.NoError(t, m.ValidateGenesis(cdc, nil, bz))

	// empty genesis state
	require.NoError(t, m.ValidateGenesis(cdc, nil, []byte(`{}`)))

	// genesis state with few orm tables data
	require.NoError(t, m.ValidateGenesis(cdc, nil, []byte(`
	{
		"regen.ecocredit.v1.CreditType": [
			{
				"name":"carbon",
				"abbreviation":"C",
				"precision":6,
				"unit":"kg"
			}
		],
		"regen.ecocredit.v1.Class": [
			2,
			{
				"key":1,
				"id":"C01",
				"admin":"abcd",
				"credit_type_abbrev":"C"
			}
		],
		"regen.ecocredit.basket.v1.Basket": [
			2,
			{
				"id":1,
				"basket_denom":"eco.uC.hello",
				"name":"abcd",
				"credit_type_abbrev":"C",
				"curator":"regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s"
			}
		],
		"regen.ecocredit.basket.v1.BasketBalance": []
	}`)))
}
