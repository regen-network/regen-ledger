package marketplace

import (
	"testing"
)

func TestAllowDenom_Validate(t *testing.T) {
	type fields struct {
		Denom        string
		DisplayDenom string
		Exponent     uint32
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				Denom:        "uregen",
				DisplayDenom: "regen",
				Exponent:     18,
			},
		},
		{
			name: "invalid denom",
			fields: fields{
				Denom:        "u!f0",
				DisplayDenom: "regen",
				Exponent:     18,
			},
			wantErr: true,
		},
		{
			name: "no denom",
			fields: fields{
				DisplayDenom: "regen",
				Exponent:     18,
			},
			wantErr: true,
		},
		{
			name: "invalid display denom",
			fields: fields{
				Denom:        "uregen",
				DisplayDenom: "r!egen",
				Exponent:     18,
			},
			wantErr: true,
		},
		{
			name: "no display denom",
			fields: fields{
				Denom:    "uregen",
				Exponent: 18,
			},
			wantErr: true,
		},
		{
			name: "invalid exponent",
			fields: fields{
				Denom:        "uregen",
				DisplayDenom: "regen",
				Exponent:     20,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := AllowedDenom{
				BankDenom:    tt.fields.Denom,
				DisplayDenom: tt.fields.DisplayDenom,
				Exponent:     tt.fields.Exponent,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
