package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

func TestMsgAddCreditType_ValidateBasic(t *testing.T) {
	type fields struct {
		RootAddress string
		CreditTypes []*CreditType
	}
	_, _, addr := testdata.KeyTestPubAddr()
	tests := []struct {
		name   string
		fields fields
		errMsg string
	}{
		{
			name: "valid",
			fields: fields{
				RootAddress: addr.String(),
				CreditTypes: []*CreditType{{Abbreviation: "FOO", Name: "FooType", Unit: "tons", Precision: 31}},
			},
		},
		{
			name: "bad addr",
			fields: fields{
				RootAddress: "bad address",
				CreditTypes: []*CreditType{{Abbreviation: "FOO", Name: "FooType", Unit: "tons", Precision: 31}},
			},
			errMsg: "decoding bech32 failed",
		},
		{
			name: "no credit types",
			fields: fields{
				RootAddress: addr.String(),
			},
			errMsg: "credit types cannot be empty",
		},
		{
			name: "invalid abbreviation",
			fields: fields{
				RootAddress: addr.String(),
				CreditTypes: []*CreditType{{Abbreviation: "this is not good", Name: "FooType", Unit: "tons", Precision: 31}},
			},
			errMsg: "credit type abbreviation must be 1-3 uppercase latin letters",
		},
		{
			name: "no units",
			fields: fields{
				RootAddress: addr.String(),
				CreditTypes: []*CreditType{{Abbreviation: "FOO", Name: "FooType", Precision: 31}},
			},
			errMsg: "unit of measurement is required",
		},
		{
			name: "no name",
			fields: fields{
				RootAddress: addr.String(),
				CreditTypes: []*CreditType{{Abbreviation: "FOO", Unit: "tons", Precision: 31}},
			},
			errMsg: "name is required",
		},
		{
			name: "invalid precision",
			fields: fields{
				RootAddress: addr.String(),
				CreditTypes: []*CreditType{{Abbreviation: "FOO", Name: "foo type", Unit: "tons", Precision: 0}},
			},
			errMsg: "precision must be greater than 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MsgAddCreditType{
				CreditTypes: tt.fields.CreditTypes,
				RootAddress: tt.fields.RootAddress,
			}
			err := m.ValidateBasic()
			if len(tt.errMsg) == 0 {
				assert.NilError(t, err)
			} else {
				assert.ErrorContains(t, err, tt.errMsg)
			}
		})
	}
}
