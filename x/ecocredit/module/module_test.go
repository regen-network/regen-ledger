package module

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

func TestHybridORMLegacyGenesis(t *testing.T) {
	m := Module{}
	cdc := codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
	bz := m.DefaultGenesis(cdc)
	require.NotNil(t, bz)
	require.NoError(t, m.ValidateGenesis(cdc, nil, bz))

	require.NoError(t, m.ValidateGenesis(cdc, nil, []byte(`
	{
	  "regen.ecocredit.v1alpha1.GenesisState":{
	    "params":{
	      "allowlist_enabled":true
	    }
	  },
	  "regen.ecocredit.basket.v1.BasketBalance":[]
	}`)))

	require.NoError(t, m.ValidateGenesis(cdc, nil, []byte(`
	{
	  "regen.ecocredit.basket.v1.Basket":[{
	    "basket_denom":"foo"
	  }],
	  "regen.ecocredit.basket.v1.BasketBalance":[]
	}`)))

	require.NoError(t, m.ValidateGenesis(cdc, nil, []byte(`
{
  "regen.ecocredit.v1alpha1.GenesisState":{
    "params":{
      "allowlist_enabled":true
    }
  },
  "regen.ecocredit.basket.v1.Basket":[{
    "basket_denom":"foo"
  }],
  "regen.ecocredit.basket.v1.BasketBalance":[{
    "batch_denom":"foo",
    "balance":"123",
    "basket_id":1
  }]
}`)))

}
