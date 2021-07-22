package ecocredit

import "testing"

func Test_validateCreditTypes(t *testing.T) {
	tests := []struct {
		name    string
		args    interface{}
		wantErr bool
	}{
		{
			name:    "valid credit types",
			args:    []*CreditType{{Type: "carbon", Units: "kg", Precision: 6}, {Type: "biodiversity", Units: "kg", Precision: 6}},
			wantErr: false,
		},
		{
			name:    "same names should fail regardless of stylization",
			args:    []*CreditType{{Type: "c a R b o N ", Units: "kg", Precision: 6}, {Type: "carbon", Units: "kg", Precision: 6}},
			wantErr: true,
		},
		{
			name:    "cant use precision other than 6",
			args:    []*CreditType{{Type: "c a R b o N ", Units: "kg", Precision: 0}},
			wantErr: true,
		},
		{
			name:    "cant use empty name",
			args:    []*CreditType{{Type: "", Units: "kg", Precision: 6}},
			wantErr: true,
		},
		{
			name:    "cant use empty units",
			args:    []*CreditType{{Type: "", Units: "", Precision: 6}},
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
