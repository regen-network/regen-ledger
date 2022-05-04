package core_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestValidateGenesis(t *testing.T) {
	x := `{
		"regen.ecocredit.v1.BatchBalance":[
			{
				"batch_key":"1",
				"address":"gydQIvR2RUi0N1RJnmgOLVSkcd4=",
				"tradable":"90.003",
				"retired":"9.997"
			}
		],
		"regen.ecocredit.v1.Batch":[
			{
				"issuer":"WCBEyNFP/N5RoS4h43AqkjC6zA8=",
				"project_key":"1",
				"denom":"BIO01-001-00000000-00000000-001",
				"start_date":"2021-04-08T10:40:10.774108556Z",
				"end_date":"2022-04-08T10:40:10.774108556Z"
			},
		    {
				"issuer":"gydQIvR2RUi0N1RJnmgOLVSkcd4=",
				"project_key":"1",
				"denom":"BIO02-001-00000000-00000000-001",
				"start_date":"2021-04-08T10:40:10.774108556Z",
				"end_date":"2022-04-08T10:40:10.774108556Z"
			}
		],
		"regen.ecocredit.v1.BatchSupply":[
			{
				"batch_key":"1",
				"tradable_amount":"90.003",
				"retired_amount":"9.997"
			}
		],
		"regen.ecocredit.v1.Class":[
			{
				"id":"BIO001",
				"admin":"4A/V6LMEL2lZv9PZnkWSIDQzZM4=",
				"credit_type_abbrev":"BIO"
			},
		    {
				"id":"BIO02",
				"admin":"HK9YDsBMN1hU8tjfLTNy+qjbqLE=",
				"credit_type_abbrev":"BIO"
			}
		],
		"regen.ecocredit.v1.Project":[
			{
				"id":"C01-001",
				"admin":"gPFuHL7Hn+uVYD6XOR00du3C/Xg=",
				"class_key":"1",
				"jurisdiction":"AQ"
			},
			{
				"id":"C01-002",
				"admin":"CHkV2Tv6A7RXPJYTivVklbxXWP8=",
				"class_key":"2",
				"jurisdiction":"AQ",
				"metadata":"project metadata"
			}
		],
		"regen.ecocredit.v1.CreditType": [
			{
				"name":"carbon",
				"abbreviation": "C",
				"unit":"metric ton C02 equivalent",
				"precision":6
			},
			{
				"name":"biodiversity",
				"abbreviation": "BIO",
				"unit":"acres",
				"precision":6
			}
		]
	}`

	params := core.Params{
		AllowlistEnabled: true,
	}
	err := core.ValidateGenesis(json.RawMessage(x), params)
	require.NoError(t, err)
}

func TestGenesisValidate(t *testing.T) {
	defaultParams := core.DefaultParams()

	testCases := []struct {
		id           string
		genesisState func() json.RawMessage
		params       core.Params
		expectErr    bool
		errorMsg     string
	}{
		{
			"valid: no credit batches",
			func() json.RawMessage {
				return json.RawMessage(`
				{
					"regen.ecocredit.v1.Class": [
						{
							"id":"C01",
							"admin":"0lxfU2Ca/sqly8hyRhD8/lNBrvM=",
							"credit_type_abbrev":"C"
						}
					],
					"regen.ecocredit.v1.CreditType": [
						{
							"name":"carbon",
							"abbreviation": "C",
							"unit":"metric ton C02 equivalent",
							"precision":6
						},
						{
							"name":"biodiversity",
							"abbreviation": "BIO",
							"unit":"acres",
							"precision":6
						}
				]}
`)
			},
			defaultParams,
			false,
			"",
		},
		{
			"invalid credit type abbreviation",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.CreditType": [
						{
							"name": "carbon",
							"abbreviation":"1234",
							"unit":"kg",
							"precision":"6"
						}
					],
					"regen.ecocredit.v1.Class": [
						{
							"id":"C01",
							"admin":"0lxfU2Ca/sqly8hyRhD8/lNBrvM=",
							"credit_type_abbrev":"C"
						}
					]}`)
			},
			defaultParams,
			true,
			"credit type abbreviation must be 1-3 uppercase latin letters",
		},
		{
			"invalid: bad addresses in allowlist",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.Class": [{
						"id":"C01",
						"admin":"OFX2S1F4zl9HmpAILrS4O6I7zEk=",
						"credit_type_abbrev":"C"
					}]					
				}`)
			},
			func() core.Params {
				p := core.DefaultParams()
				p.AllowlistEnabled = true
				p.AllowedClassCreators = []string{"-=!?#09)("}
				return p
			}(),
			true,
			"invalid creator address: decoding bech32 failed",
		},
		{
			"invalid: type id does not match param id",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.Class": [{
						"id":"C01",
						"admin":"gm+Xr47EcefPFePZxYYL6WaK6V8=",
						"credit_type_abbrev":"F"
					}]
				}`)
			},
			defaultParams,
			true,
			"credit type not exist",
		},
		{
			"invalid: non-existent abbreviation",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.Class": [{
						"id":"C01",
						"admin":"gm+Xr47EcefPFePZxYYL6WaK6V8=",
						"credit_type_abbrev":"F"
					}]	
				}`)
			},
			defaultParams,
			true,
			"credit type not exist for F abbreviation: not found",
		},
		{
			"expect error: supply is missing",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.Class":[{
						"id":"C01",
						"admin":"PPUOsQeEHJyQV0ABQzU91iytr9s=",
						"credit_type_abbrev":"C"
					}],
					"regen.ecocredit.v1.Project":[{
						"id":"C01-001",
						"admin":"PPUOsQeEHJyQV0ABQzU91iytr9s=",
						"class_key":"1",
						"jurisdiction":"AQ"
					}],
					"regen.ecocredit.v1.Batch":[{
						"issuer":"PPUOsQeEHJyQV0ABQzU91iytr9s=",
						"project_key":"1",
						"denom":"C01-001-00000000-00000000-001",
						"start_date":"2021-04-08T10:40:10.774108556Z",
						"end_date":"2022-04-08T10:40:10.774108556Z"
					}],
					"regen.ecocredit.v1.CreditType": [
						{
							"name":"carbon",
							"abbreviation": "C",
							"unit":"metric ton C02 equivalent",
							"precision":6
						},
						{
							"name":"biodiversity",
							"abbreviation": "BIO",
							"unit":"acres",
							"precision":6
						}
					],
					"regen.ecocredit.v1.BatchBalance":[{
						"batch_key":"1",
						"address":"mAAyikSMAfVwmlW4BPV2Q6GmpHc=",
						"tradable":"400.456"
					}]
				}`)
			},
			defaultParams,
			true,
			"supply is not found",
		},
		{
			"expect error: invalid supply",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.CreditType": [
						{
							"name":"carbon",
							"abbreviation": "C",
							"unit":"metric ton C02 equivalent",
							"precision":6
						},
						{
							"name":"biodiversity",
							"abbreviation": "BIO",
							"unit":"acres",
							"precision":6
						}
					],
					"regen.ecocredit.v1.Class":[{
						"id":"C01",
						"admin":"PPUOsQeEHJyQV0ABQzU91iytr9s=",
						"credit_type_abbrev":"C"
					}],
					"regen.ecocredit.v1.Project":[{
						"id":"C01-001",
						"admin":"PPUOsQeEHJyQV0ABQzU91iytr9s=",
						"class_key":"1",
						"jurisdiction":"AQ"
					}],
					"regen.ecocredit.v1.Batch":[{
						"issuer":"PPUOsQeEHJyQV0ABQzU91iytr9s=",
						"project_key":"1",
						"denom":"C01-001-00000000-00000000-001",
						"start_date":"2021-04-08T10:40:10.774108556Z",
						"end_date":"2022-04-08T10:40:10.774108556Z"
					}],
					"regen.ecocredit.v1.BatchBalance":[{
						"batch_key":"1",
						"address":"mAAyikSMAfVwmlW4BPV2Q6GmpHc=",
						"tradable":"100",
						"retired":"100"
					}],
					"regen.ecocredit.v1.BatchSupply":[{
						"batch_key":"1",
						"tradable_amount":"10"
					}]
				}`)
			},
			defaultParams,
			true,
			"supply is incorrect for 1 credit batch, expected 10, got 200: invalid coins",
		},
		{
			"valid test case",
			func() json.RawMessage {
				return json.RawMessage(`{
				"regen.ecocredit.v1.CreditType": [
					{
						"name":"carbon",
						"abbreviation": "C",
						"unit":"metric ton C02 equivalent",
						"precision":6
					},
					{
						"name":"biodiversity",
						"abbreviation": "BIO",
						"unit":"acres",
						"precision":6
					}
				],
				"regen.ecocredit.v1.Class":[{
					"id":"C01",
					"admin":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
					"credit_type_abbrev":"C"
				}],
				"regen.ecocredit.v1.Project":[{
					"id":"C01-001",
					"admin":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
					"class_key":"1",
					"jurisdiction":"AQ"
				}],
				"regen.ecocredit.v1.Batch":[{
					"issuer":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
					"project_key":"1",
					"denom":"C01-001-00000000-00000000-001",
					"start_date":"2021-04-08T10:40:10.774108556Z",
					"end_date":"2022-04-08T10:40:10.774108556Z"
				}],
				"regen.ecocredit.v1.BatchBalance":[
					{
						"batch_key":"1",
						"address":"Ak5WDUYGfdv4gNMF500MFF86NWA=",
						"tradable":"100.123",
						"retired":"100.123"
					},
					{
						"batch_key":"1",
						"address":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
						"tradable":"100.123",
						"retired":"100.123"
					}
				],
				"regen.ecocredit.v1.BatchSupply":[
					{
						"batch_key":"1",
						"tradable_amount":"200.246",
						"retired_amount":"200.246"
					}
				]
			}`)
			},
			defaultParams,
			false,
			"",
		},
		{
			"valid test case escrowed balance",
			func() json.RawMessage {
				return json.RawMessage(`{
				"regen.ecocredit.v1.CreditType": [
					{
						"name":"carbon",
						"abbreviation": "C",
						"unit":"metric ton C02 equivalent",
						"precision":6
					},
					{
						"name":"biodiversity",
						"abbreviation": "BIO",
						"unit":"acres",
						"precision":6
					}
				],
				"regen.ecocredit.v1.Class":[{
					"id":"C01",
					"admin":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
					"credit_type_abbrev":"C"
				}],
				"regen.ecocredit.v1.Project":[{
					"id":"C01-001",
					"admin":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
					"class_key":"1",
					"jurisdiction":"AQ"
				}],
				"regen.ecocredit.v1.Batch":[{
					"issuer":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
					"project_key":"1",
					"denom":"C01-001-00000000-00000000-001",
					"start_date":"2021-04-08T10:40:10.774108556Z",
					"end_date":"2022-04-08T10:40:10.774108556Z",
					"metadata":"meta-data"
				}],
				"regen.ecocredit.v1.BatchBalance":[
					{
						"batch_key":"1",
						"address":"Ak5WDUYGfdv4gNMF500MFF86NWA=",
						"tradable":"100.123",
						"retired":"100.123",
						"escrowed":"100.123"
					},
					{
						"batch_key":"1",
						"address":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
						"tradable":"100.123",
						"retired":"100.123"
					}
				],
				"regen.ecocredit.v1.BatchSupply":[
					{
						"batch_key":"1",
						"tradable_amount":"300.369",
						"retired_amount":"200.246"
					}
				]
			}`)
			},
			defaultParams,
			false,
			"",
		},
		{
			"valid test case, multiple classes",
			func() json.RawMessage {
				return json.RawMessage(`
				{
					"regen.ecocredit.v1.CreditType": [
						{
							"name":"carbon",
							"abbreviation": "C",
							"unit":"metric ton C02 equivalent",
							"precision":6
						},
						{
							"name":"biodiversity",
							"abbreviation": "BIO",
							"unit":"acres",
							"precision":6
						}
					],
					"regen.ecocredit.v1.Class": [
					  {
						"id": "C01",
						"admin": "OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
						"credit_type_abbrev": "C"
					  },
					  {
						"id": "C02",
						"admin": "Ak5WDUYGfdv4gNMF500MFF86NWA=",
						"credit_type_abbrev": "C"
					  }
					],
					"regen.ecocredit.v1.Project": [
					  {
						"id": "C01-001",
						"admin": "OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
						"class_key": "1",
						"jurisdiction":"AQ"
					  },
					  {
						"id": "C02-001",
						"admin": "Ak5WDUYGfdv4gNMF500MFF86NWA=",
						"class_key": "2",
						"jurisdiction":"AQ"
					  }
					],
					"regen.ecocredit.v1.Batch": [
					  {
						"issuer": "OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
						"project_key": "1",
						"denom":"C01-001-00000000-00000000-001",
						"start_date":"2021-04-08T10:40:10.774108556Z",
						"end_date":"2022-04-08T10:40:10.774108556Z",
						"metadata":"meta-data"
					  },
					  {
						"issuer": "OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
						"project_key": "2",
						"denom":"C01-002-00000000-00000000-001",
						"start_date":"2021-04-08T10:40:10.774108556Z",
						"end_date":"2022-04-08T10:40:10.774108556Z",
						"metadata":"meta-data"
					  }
					],
					"regen.ecocredit.v1.BatchBalance": [
					  {
						"batch_key": "1",
						"address": "Ak5WDUYGfdv4gNMF500MFF86NWA=",
						"tradable": "100.123",
						"retired": "100.123"
					  },
					  {
						"batch_key": "2",
						"address": "OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=",
						"tradable": "100.123",
						"retired": "100.123"
					  }
					],
					"regen.ecocredit.v1.BatchSupply": [
					  {
						"batch_key": "1",
						"tradable_amount": "100.123",
						"retired_amount": "100.123"
					  },
					  {
						"batch_key": "2",
						"tradable_amount": "100.123",
						"retired_amount": "100.123"
					  }
					]
				  }
				`)
			},
			defaultParams,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.id, func(t *testing.T) {
			err := core.ValidateGenesis(tc.genesisState(), tc.params)
			if tc.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errorMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
