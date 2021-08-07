package ecocredit

import "testing"

func Test_validateCreditTypes(t *testing.T) {
	tests := []struct {
		name    string
		args    interface{}
		wantErr bool
	}{
		{
			name: "valid credit types",
			args: []*CreditType{
				{Name: "carbon", Abbreviation: "C", Units: "tons", Precision: 6},
				{Name: "biodiversity", Abbreviation: "BIO", Units: "mi", Precision: 6},
			},
			wantErr: false,
		},
		{
			name: "wrong type",
			args: []*ClassInfo{
				{
					ClassId:    "foo",
					Designer:   "0xdeadbeef",
					Issuers:    []string{"not", "an", "address"},
					Metadata:   nil,
					CreditType: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "cant have duplicate names",
			args: []*CreditType{
				{Name: "carbon", Abbreviation: "C", Units: "tons", Precision: 6},
				{Name: "carbon", Abbreviation: "CAR", Units: "tons", Precision: 6},
			},
			wantErr: true,
		},
		{
			name:    "cant use non-normalized credit type name",
			args:    []*CreditType{{Name: "biODiVerSitY", Abbreviation: "BIO", Units: "tons", Precision: 6}},
			wantErr: true,
		},
		{
			name:    "cant use empty name",
			args:    []*CreditType{{Name: "", Abbreviation: "C", Units: "tons", Precision: 6}},
			wantErr: true,
		},
		{
			name: "cant have duplicate abbreviations",
			args: []*CreditType{
				{Name: "carbon", Abbreviation: "C", Units: "tons", Precision: 6},
				{Name: "carbonic acid", Abbreviation: "C", Units: "tons", Precision: 6},
			},
			wantErr: true,
		},
		{
			name:    "cant use empty abbreviation",
			args:    []*CreditType{{Name: "carbon", Units: "tons", Precision: 6}},
			wantErr: true,
		},
		{
			name:    "cant use lowercase abbreviation",
			args:    []*CreditType{{Name: "carbon", Abbreviation: "c", Units: "tons", Precision: 6}},
			wantErr: true,
		},
		{
			name:    "cant use longer than 3 letter abbreviation",
			args:    []*CreditType{{Name: "carbon", Abbreviation: "CARB", Units: "tons", Precision: 6}},
			wantErr: true,
		},
		{
			name:    "cant use precision other than 6",
			args:    []*CreditType{{Name: "carbon", Abbreviation: "C", Units: "tons", Precision: 0}},
			wantErr: true,
		},
		{
			name:    "cant use empty units",
			args:    []*CreditType{{Name: "carbon", Abbreviation: "C", Units: "", Precision: 6}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreditTypes(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("validateCreditTypes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
