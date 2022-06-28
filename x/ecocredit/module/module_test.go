package module

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"
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
  "regen.ecocredit.basket.v1.BasketBalance":[]
}`)))

	require.NoError(t, m.ValidateGenesis(cdc, nil, []byte(`
{
	"regen.ecocredit.basket.v1.Basket": [
		2,
		{
			"basket_denom": "eco.uC.rNCT",
			"credit_type_abbrev": "C",
			"date_criteria": {
				"start_date_window": "315576000s"
			},
			"exponent": 6,
			"id": "1",
			"name": "rNCT"
		},
		{
			"basket_denom": "eco.uC.NCT",
			"credit_type_abbrev": "C",
			"date_criteria": {
				"start_date_window": "315576000s"
			},
			"exponent": 6,
			"id": "2",
			"name": "NCT"
		}
	],
	"regen.ecocredit.basket.v1.BasketBalance": [
		{
			"balance": "11947.698895",
			"basket_id": "1",
			"batch_denom": "C01-20180101-20200101-001",
			"batch_start_date": "2018-01-01T00:00:00Z"
		},
		{
			"balance": "10",
			"basket_id": "2",
			"batch_denom": "C01-20190101-20191231-009",
			"batch_start_date": "2019-01-01T00:00:00Z"
		}
	],
	"regen.ecocredit.basket.v1.BasketClass": [
		{
			"basket_id": "1",
			"class_id": "C01"
		},
		{
			"basket_id": "1",
			"class_id": "C02"
		},
		{
			"basket_id": "2",
			"class_id": "C01"
		}
	],
	"regen.ecocredit.v1alpha1.GenesisState": {
		"balances": [
			{
				"address": "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
				"batch_denom": "C01-20180101-20200101-001",
				"retired_balance": "64.301105",
				"tradable_balance": "69753"
			},
			{
				"address": "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
				"batch_denom": "C01-20190101-20191231-009",
				"retired_balance": "",
				"tradable_balance": "51"
			},
			{
				"address": "regen1pg97ufrzeu5hutcks95jzzex9fn7yuyew4da70",
				"batch_denom": "C01-20180101-20200101-001",
				"retired_balance": "",
				"tradable_balance": "1000"
			},
			{
				"address": "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
				"batch_denom": "C01-20190101-20210101-002",
				"retired_balance": "20",
				"tradable_balance": "98732"
			},
			{
				"address": "regen1lhzh0dl7hqq4xgfhps67wrlruejzxcnq5z2sk0",
				"batch_denom": "C01-20180101-20200101-001",
				"retired_balance": "",
				"tradable_balance": "5000"
			},
			{
				"address": "regen1rhvx7kpxtd0em0t2u87xu5a3xc9gs7evvr5v29",
				"batch_denom": "C01-20180101-20200101-001",
				"retired_balance": "13",
				"tradable_balance": "204"
			},
			{
				"address": "regen1sv6a7ry6nrls84z0w5lauae4mgmw3kh2mg97ht",
				"batch_denom": "C01-20190101-20210101-002",
				"retired_balance": "20",
				"tradable_balance": "999"
			},
			{
				"address": "regen1sv6a7ry6nrls84z0w5lauae4mgmw3kh2mg97ht",
				"batch_denom": "C01-20180101-20200101-001",
				"retired_balance": "10",
				"tradable_balance": "1000"
			},
			{
				"address": "regen15h5eszss2wtavw2x73f66jqp00sh4kewltwppt",
				"batch_denom": "C01-20180101-20200101-001",
				"retired_balance": "",
				"tradable_balance": "10000"
			},
			{
				"address": "regen1ql2tzktyc5clwgzp43g60khdrsecl9n8vfe70u",
				"batch_denom": "C01-20180101-20200101-001",
				"retired_balance": "7",
				"tradable_balance": "996"
			},
			{
				"address": "regen1rhvx7kpxtd0em0t2u87xu5a3xc9gs7evvr5v29",
				"batch_denom": "C01-20190101-20210101-002",
				"retired_balance": "23",
				"tradable_balance": "206"
			},
			{
				"address": "regen1axf5u6wj0tut6vs5w46krycs00k9lwmxyr7l9v",
				"batch_denom": "C01-20180101-20200101-001",
				"retired_balance": "",
				"tradable_balance": "1"
			},
			{
				"address": "regen1gtlfmszmhv3jnlqx6smt9n6rcwsfydrhznqyk9",
				"batch_denom": "C01-20180101-20200101-001",
				"retired_balance": "4",
				"tradable_balance": ""
			}
		],
		"batch_info": [
			{
				"amount_cancelled": "0",
				"batch_denom": "C01-20180101-20200101-001",
				"class_id": "C01",
				"end_date": "2020-01-01T00:00:00Z",
				"issuer": "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
				"metadata": null,
				"project_location": "AU-NSW 2453",
				"start_date": "2018-01-01T00:00:00Z",
				"total_amount": "100000"
			},
			{
				"amount_cancelled": "0",
				"batch_denom": "C01-20190101-20191231-009",
				"class_id": "C01",
				"end_date": "2019-12-31T00:00:00Z",
				"issuer": "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
				"metadata": "cmVnZW46MTN0b1ZnRjg0a1F3U1gxMURkaERhc1l0TUZVMVliNnFRd1F2dHYxcnZIOHBmNUU4VVR5YWpDWC5yZGY=",
				"project_location": "KE",
				"start_date": "2019-01-01T00:00:00Z",
				"total_amount": "61"
			},
			{
				"amount_cancelled": "0",
				"batch_denom": "C01-20190101-20210101-002",
				"class_id": "C01",
				"end_date": "2021-01-01T00:00:00Z",
				"issuer": "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
				"metadata": null,
				"project_location": "US",
				"start_date": "2019-01-01T00:00:00Z",
				"total_amount": "100000"
			}
		],
		"class_info": [
			{
				"admin": "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
				"class_id": "C01",
				"credit_type": {
					"abbreviation": "C",
					"name": "carbon",
					"precision": 6,
					"unit": "metric ton CO2 equivalent"
				},
				"issuers": [
					"regen1wjul39t07ds68xasfc4mw8yszwappktmypuuch",
					"regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46"
				],
				"metadata": "cmVnZW46MTN0b1ZnbzVDQ21Ra1BKRHdMZWd0ZjRVMWVzVzVycnRXcHdxRTZuU2RwMWhhOVc4OFJmdWY1TS5yZGY=",
				"num_batches": "9"
			},
			{
				"admin": "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
				"class_id": "C02",
				"credit_type": {
					"abbreviation": "C",
					"name": "carbon",
					"precision": 6,
					"unit": "metric ton CO2 equivalent"
				},
				"issuers": [
					"regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46"
				],
				"metadata": "cmVnZW46YUhSMGNEb3ZMM0psWjJWdUxtNWxkSGR2Y21zdlZrTlRMVk4wWVc1a1lYSmsucmRm",
				"num_batches": "2"
			}
		],
		"params": {
			"allowed_class_creators": [],
			"allowlist_enabled": false,
			"basket_creation_fee": [
				{
					"amount": "20000000",
					"denom": "uregen"
				}
			],
			"credit_class_fee": [
				{
					"amount": "20000000",
					"denom": "uregen"
				}
			],
			"credit_types": [
				{
					"abbreviation": "C",
					"name": "carbon",
					"precision": 6,
					"unit": "metric ton CO2 equivalent"
				}
			]
		},
		"sequences": [
			{
				"abbreviation": "C",
				"seq_number": "3"
			}
		],
		"supplies": [
			{
				"batch_denom": "C01-20180101-20200101-001",
				"retired_supply": "98.301105",
				"tradable_supply": "99901.698895"
			},
			{
				"batch_denom": "C01-20190101-20210101-002",
				"retired_supply": "63",
				"tradable_supply": "99937"
			},
			{
				"batch_denom": "C01-20190101-20191231-009",
				"retired_supply": "",
				"tradable_supply": "61"
			}
		]
	}
}
`)))
}
