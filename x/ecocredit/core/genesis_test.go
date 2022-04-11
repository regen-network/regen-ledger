package core_test

import (
	"encoding/json"
	fmt "fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestValidateGenesis(t *testing.T) {
	x := `{
		"regen.ecocredit.v1.BatchBalance":[
			{"address":"gydQIvR2RUi0N1RJnmgOLVSkcd4=","batch_id":"1","tradable":"90.003","retired":"9.997","escrowed":""}
		],
		"regen.ecocredit.v1.BatchInfo":[
			{"issuer":"WCBEyNFP/N5RoS4h43AqkjC6zA8=","project_id":"1","batch_denom":"BIO01-00000000-00000000-001","metadata":"batch metadata","start_date":null,"end_date":null,"issuance_date":"2022-04-08T10:40:10.774108141Z"},
		    {"issuer":null,"project_id":"1","batch_denom":"BIO02-00000000-00000000-001","metadata":"batch metadata","start_date":null,"end_date":null,"issuance_date":"2022-04-08T10:40:10.774108556Z"}
		],
		"regen.ecocredit.v1.BatchSequence":[{"project_id":"P01","next_batch_id":"3"}],
		"regen.ecocredit.v1.BatchSupply":[{"batch_id":"1","tradable_amount":"90.003","retired_amount":"9.997","cancelled_amount":""}],
		"regen.ecocredit.v1.ClassInfo":[
			{"name":"BIO001","admin":"4A/V6LMEL2lZv9PZnkWSIDQzZM4=","metadata":"credit class metadata","credit_type":"BIO"},
		    {"name":"BIO02","admin":"HK9YDsBMN1hU8tjfLTNy+qjbqLE=","metadata":"credit class metadata","credit_type":"BIO"}
		],
		"regen.ecocredit.v1.ClassIssuer":[
			{"class_id":"1","issuer":"1ygCfmJaPVMIvVEcpx6r+2gpurM="},
			{"class_id":"1","issuer":"KoXfzfqe+V/9x7C4XjnqDFB2Tl4="},
			{"class_id":"2","issuer":"KoXfzfqe+V/9x7C4XjnqDFB2Tl4="},
			{"class_id":"2","issuer":"lEjmu9Vooa24qp9vCMIlXGrMZoU="}
		],
		"regen.ecocredit.v1.ClassSequence":[{"credit_type":"BIO","next_class_id":"3"}],
		"regen.ecocredit.v1.ProjectInfo":[
			{"name":"P01","admin":"gPFuHL7Hn+uVYD6XOR00du3C/Xg=","class_id":"1","project_location":"AQ","metadata":"project metadata"},
			{"name":"P02","admin":"CHkV2Tv6A7RXPJYTivVklbxXWP8=","class_id":"2","project_location":"AQ","metadata":"project metadata"}
		],
		"regen.ecocredit.v1.ProjectSequence":[{"class_id":"1","next_project_id":"3"}]}`

	params := core.Params{
		CreditTypes: []*core.CreditType{
			{
				Name:         "carbon",
				Abbreviation: "C",
				Unit:         "metric ton CO2 equivalent",
				Precision:    6,
			},
			{
				Abbreviation: "BIO",
				Name:         "biodiversity",
				Unit:         "ton",
				Precision:    6,
			},
		},
	}
	core.DefaultParams()
	err := core.ValidateGenesis(json.RawMessage(x), params)
	require.NoError(t, err)
}

func TestGenesisValidate(t *testing.T) {
	defaultParams := core.DefaultParams()

	testCases := []struct {
		name        string
		gensisState func() json.RawMessage
		params      core.Params
		expectErr   bool
		errorMsg    string
	}{
		{
			"valid: no credit batches",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.ClassInfo": [{"id":1,"name":"1","admin":"0lxfU2Ca/sqly8hyRhD8/lNBrvM=","metadata":"meta-data","credit_type":"C"}]
				}`)
			},
			defaultParams,
			false,
			"",
		},
		{
			"invalid: credit type param",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.ClassInfo": [{"name":"1","admin":"v9PCozRRuFc5I5hdJOwD3k9WMOI=","metadata":"meta-data","credit_type":"C"}]					
					}`)
			},
			func() core.Params {
				p := core.DefaultParams()
				p.CreditTypes[0].Precision = 7
				return p
			}(),
			true,
			"invalid precision 7: precision is currently locked to 6: invalid request",
		},
		{
			"invalid: duplicate credit type",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.ClassInfo": [{"name":"1","admin":"OFX2S1F4zl9HmpAILrS4O6I7zEk=","metadata":"meta-data","credit_type":"C"}]					
					}`)
			},
			func() core.Params {
				p := core.DefaultParams()
				p.CreditTypes = []*core.CreditType{{
					Name:         "carbon",
					Abbreviation: "C",
					Unit:         "metric ton CO2 equivalent",
					Precision:    6,
				}, {
					Name:         "carbon",
					Abbreviation: "C",
					Unit:         "metric ton CO2 equivalent",
					Precision:    6,
				}}
				return p
			}(),
			true,
			"duplicate credit type name in request: carbon: invalid request",
		},
		{
			"invalid: bad addresses in allowlist",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.ClassInfo": [{"name":"1","admin":"OFX2S1F4zl9HmpAILrS4O6I7zEk=","metadata":"meta-data","credit_type":"C"}]					
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
			"invalid: type name does not match param name",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.ClassInfo": [{"name":"1","admin":"gm+Xr47EcefPFePZxYYL6WaK6V8=","metadata":"meta-data","credit_type":"F"}]
				}`)
			},
			defaultParams,
			true,
			formatCreditTypeParamError(core.CreditType{"badbadnotgood", "C", "metric ton CO2 equivalent", 6}).Error(),
		},
		{
			"invalid: non-existent abbreviation",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.ClassInfo": [{name":"1","admin":"gm+Xr47EcefPFePZxYYL6WaK6V8=","metadata":"meta-data","credit_type":"F"}]	
				}`)
			},
			defaultParams,
			true,
			"unknown credit type abbreviation: F: not found",
		},
		{
			"expect error: supply is missing",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.CreditTypes":[{"abbreviation":"C","name":"carbon","unit":"metric ton CO2 equivalent","precision":6}],
					"regen.ecocredit.v1.ClassInfo":[{"name":"1","admin":"PPUOsQeEHJyQV0ABQzU91iytr9s=","metadata":"meta-data","credit_type":"C"}],
					"regen.ecocredit.v1.ProjectInfo":[{"name":"01","admin":"PPUOsQeEHJyQV0ABQzU91iytr9s=","class_id":"1","project_location":"AQ","metadata":"meta-data"}],
					"regen.ecocredit.v1.BatchInfo":[{issuer":"PPUOsQeEHJyQV0ABQzU91iytr9s=","project_id":"1","batch_denom":"1/2","metadata":"meta-data","start_date":null,"end_date":null,"issuance_date":null}],
					regen.ecocredit.v1.BatchBalance":[{"address":"mAAyikSMAfVwmlW4BPV2Q6GmpHc=","batch_id":"1","tradable":"400.456","retired":"","escrowed":""}],
				}`)
			},
			defaultParams,
			true,
			"supply is not found for 1/2 credit batch: not found",
		},
		{
			"expect error: invalid supply",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.ClassInfo":[{"name":"1","admin":"PPUOsQeEHJyQV0ABQzU91iytr9s=","metadata":"meta-data","credit_type":"C"}],
					"regen.ecocredit.v1.ProjectInfo":[{"name":"01","admin":"PPUOsQeEHJyQV0ABQzU91iytr9s=","class_id":"1","project_location":"AQ","metadata":"meta-data"}],
					"regen.ecocredit.v1.BatchInfo":[{"issuer":"PPUOsQeEHJyQV0ABQzU91iytr9s=","project_id":"1","batch_denom":"1/2","metadata":"meta-data","start_date":null,"end_date":null,"issuance_date":null}],
					"regen.ecocredit.v1.BatchBalance":[{"address":"mAAyikSMAfVwmlW4BPV2Q6GmpHc=","batch_id":"1","tradable":"100","retired":"100","escrowed":""}],
					"regen.ecocredit.v1.BatchSupply":[{"batch_id":"1","tradable_amount":"10","retired_amount":"","cancelled_amount":""}]
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
				"regen.ecocredit.v1.ClassInfo":[{"name":"1","admin":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=","metadata":"meta-data","credit_type":"C"}],
				"regen.ecocredit.v1.ProjectInfo":[{"name":"01","admin":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=","class_id":"1","project_location":"AQ","metadata":"meta-data"}],
				"regen.ecocredit.v1.BatchInfo":[{"issuer":null,"project_id":"1","batch_denom":"1/2","metadata":"meta-data","start_date":null,"end_date":null,"issuance_date":null}],
				"regen.ecocredit.v1.BatchBalance":[{"address":"Ak5WDUYGfdv4gNMF500MFF86NWA=","batch_id":"1","tradable":"100.123","retired":"100.123","escrowed":""},{"address":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=","batch_id":"1","tradable":"100.123","retired":"100.123","escrowed":""}],
				"regen.ecocredit.v1.BatchSupply":[{"batch_id":"1","tradable_amount":"200.246","retired_amount":"200.246","cancelled_amount":""}]
			}`)
			},
			defaultParams,
			false,
			"",
		},
		{
			"valid test case, multiple classes",
			func() json.RawMessage {
				return json.RawMessage(`{
					"regen.ecocredit.v1.ClassInfo":[{"name":"1","admin":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=","metadata":"meta-data","credit_type":"C"},{"name":"2","admin":"Ak5WDUYGfdv4gNMF500MFF86NWA=","metadata":"meta-data","credit_type":"C"}],
					"regen.ecocredit.v1.ProjectInfo":[{"name":"01","admin":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=","class_id":"1","project_location":"AQ","metadata":"meta-data"},{"name":"03","admin":"Ak5WDUYGfdv4gNMF500MFF86NWA=","class_id":"2","project_location":"AQ","metadata":"meta-data"}],
					"regen.ecocredit.v1.BatchInfo":[{"issuer":null,"project_id":"1","batch_denom":"1/2","metadata":"meta-data","start_date":null,"end_date":null,"issuance_date":null},{"issuer":null,"project_id":"2","batch_denom":"2/2","metadata":"meta-data","start_date":null,"end_date":null,"issuance_date":null}],
					"regen.ecocredit.v1.BatchBalance":[{"address":"Ak5WDUYGfdv4gNMF500MFF86NWA=","batch_id":"1","tradable":"100.123","retired":"100.123","escrowed":""},{"address":"OfVGZ+vChK/1gQfbXZ6rxsz3QNQ=","batch_id":"2","tradable":"100.123","retired":"100.123","escrowed":""}],
					"regen.ecocredit.v1.BatchSupply":[{"batch_id":"1","tradable_amount":"100.123","retired_amount":"100.123","cancelled_amount":""},{"batch_id":"2","tradable_amount":"100.123","retired_amount":"100.123","cancelled_amount":""}]
				}`)
			},
			defaultParams,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := core.ValidateGenesis(tc.gensisState(), tc.params)
			if tc.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errorMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

var defaultCreditTypes = core.DefaultParams().CreditTypes

func formatCreditTypeParamError(ct core.CreditType) error {
	return fmt.Errorf("credit type %+v does not match param type %+v: invalid type", ct, *defaultCreditTypes[0])
}
